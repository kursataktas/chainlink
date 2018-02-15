package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/smartcontractkit/chainlink/logger"
	"github.com/smartcontractkit/chainlink/store"
	"github.com/smartcontractkit/chainlink/store/models"
	"github.com/smartcontractkit/chainlink/utils"
	"go.uber.org/multierr"
)

const (
	EventTopicSignature = iota
	EventTopicRequestID
	EventTopicJobID
)

const RunLogTopic = "0x06f4bf36b4e011a5c499cef1113c2d166800ce4013f6c2509cab1a0e92b83fb2"

// NotificationListener contains fields for the pointer of the store and
// a channel to the EthNotification (as the field 'logs').
type NotificationListener struct {
	Store             *store.Store
	subscriptions     []*rpc.ClientSubscription
	logNotifications  chan types.Log
	headNotifications chan types.Header
	errors            chan error
	mutex             sync.Mutex
}

// Start obtains the jobs from the store and begins execution
// of the jobs' given runs.
func (nl *NotificationListener) Start() error {
	nl.errors = make(chan error)
	nl.logNotifications = make(chan types.Log)
	nl.headNotifications = make(chan types.Header)

	if err := nl.subscribeToNewHeads(); err != nil {
		return err
	}

	jobs, err := nl.Store.Jobs()
	if err != nil {
		return err
	}
	if err := nl.subscribeToInitiators(jobs); err != nil {
		return err
	}

	go nl.listenToSubscriptionErrors()
	go nl.listenToNewHeads()
	go nl.listenToLogs()
	return nil
}

// Stop gracefully closes its access to the store's EthNotifications.
func (nl *NotificationListener) Stop() error {
	if nl.logNotifications != nil {
		nl.unsubscribe()
		close(nl.errors)
		close(nl.logNotifications)
	}
	return nil
}

// AddJob looks for "runlog" and "ethlog" Initiators for a given job
// and watches the Ethereum blockchain for the addresses in the job.
func (nl *NotificationListener) AddJob(job models.Job) error {
	var addresses []common.Address
	for _, initr := range job.InitiatorsFor(models.InitiatorEthLog, models.InitiatorRunLog) {
		logger.Debugw(fmt.Sprintf("Listening for logs from address %v", initr.Address.String()))
		addresses = append(addresses, initr.Address)
	}

	if len(addresses) == 0 {
		return nil
	}

	sub, err := nl.Store.TxManager.SubscribeToLogs(nl.logNotifications, addresses)
	if err != nil {
		return err
	}
	nl.addSubscription(sub)
	return nil
}

func (nl *NotificationListener) subscribeToNewHeads() error {
	sub, err := nl.Store.TxManager.SubscribeToNewHeads(nl.headNotifications)
	if err != nil {
		return err
	}
	nl.addSubscription(sub)
	return nil
}

func (nl *NotificationListener) subscribeToInitiators(jobs []models.Job) error {
	var err error
	for _, j := range jobs {
		err = multierr.Append(err, nl.AddJob(j))
	}
	return err
}

func (nl *NotificationListener) listenToNewHeads() {
	for range nl.headNotifications {
		pendingRuns, err := nl.Store.PendingJobRuns()
		if err != nil {
			logger.Error(err.Error())
		}
		for _, jr := range pendingRuns {
			if err := ExecuteRun(&jr, nl.Store, models.JSON{}); err != nil {
				logger.Error(err.Error())
			}
		}
	}
}

func (nl *NotificationListener) addSubscription(sub *rpc.ClientSubscription) {
	nl.mutex.Lock()
	defer nl.mutex.Unlock()
	nl.subscriptions = append(nl.subscriptions, sub)
	go func() {
		nl.errors <- (<-sub.Err())
	}()
}

func (nl *NotificationListener) unsubscribe() {
	nl.mutex.Lock()
	defer nl.mutex.Unlock()
	for _, sub := range nl.subscriptions {
		if sub.Err() != nil {
			sub.Unsubscribe()
		}
	}
}

func (nl *NotificationListener) listenToSubscriptionErrors() {
	for err := range nl.errors {
		logger.Errorw("Error in log subscription", "err", err)
	}
}

func (nl *NotificationListener) listenToLogs() {
	for el := range nl.logNotifications {
		if err := nl.receiveLog(el); err != nil {
			logger.Errorw(err.Error())
		}
	}
}

func (nl *NotificationListener) receiveLog(el types.Log) error {
	var merr error
	msg := fmt.Sprintf("Received log from %v", el.Address.String())
	logger.Debugw(msg, "log", el)
	for _, initr := range nl.initrsWithLogAndAddress(el.Address) {
		job, err := nl.Store.FindJob(initr.JobID)
		if err != nil {
			msg := fmt.Sprintf("Error initiating job from log: %v", err)
			logger.Errorw(msg, "job", initr.JobID, "initiator", initr.ID)
			merr = multierr.Append(merr, err)
			continue
		}

		input, err := FormatLogJSON(initr, el)
		if err != nil {
			logger.Errorw(err.Error(), "job", initr.JobID, "initiator", initr.ID)
			merr = multierr.Append(merr, err)
			continue
		}

		_, err = BeginRun(job, nl.Store, input)
		merr = multierr.Append(merr, err)
	}
	return merr
}

// FormatLogJSON uses the Initiator to decide how to format the EventLog
// as a JSON object.
func FormatLogJSON(initr models.Initiator, el types.Log) (models.JSON, error) {
	if initr.Type == models.InitiatorEthLog {
		return ethLogJSON(el)
	} else if initr.Type == models.InitiatorRunLog {
		out, err := runLogJSON(el)
		return out, err
	}
	return models.JSON{}, fmt.Errorf("no supported initiator type was found")
}

// make our own types.Log for better serialization
func ethLogJSON(el types.Log) (models.JSON, error) {
	var out models.JSON
	b, err := json.Marshal(el)
	if err != nil {
		return out, err
	}
	return out, json.Unmarshal(b, &out)
}

func runLogJSON(el types.Log) (models.JSON, error) {
	js, err := decodeABIToJSON(el.Data)
	if err != nil {
		return js, err
	}

	js, err = js.Add("address", el.Address.String())
	if err != nil {
		return js, err
	}

	js, err = js.Add("dataPrefix", el.Topics[EventTopicRequestID].String())
	if err != nil {
		return js, err
	}

	return js.Add("functionSelector", "76005c26")
}

func (nl *NotificationListener) initrsWithLogAndAddress(address common.Address) []models.Initiator {
	initrs := []models.Initiator{}
	query := nl.Store.Select(q.Or(
		q.And(q.Eq("Address", address), q.Re("Type", models.InitiatorRunLog)),
		q.And(q.Eq("Address", address), q.Re("Type", models.InitiatorEthLog)),
	))
	if err := query.Find(&initrs); err != nil {
		msg := fmt.Sprintf("Initiating job from log: %v", err)
		logger.Errorw(msg, "address", address.String())
	}
	return initrs
}

func decodeABIToJSON(data hexutil.Bytes) (models.JSON, error) {
	varLocationSize := 32
	varLengthSize := 32
	var js models.JSON
	hex := []byte(string([]byte(data)[varLocationSize+varLengthSize:]))
	return js, json.Unmarshal(bytes.TrimRight(hex, "\x00"), &js)
}

// InitiatorsForLog returns all of the Initiators relevant to a log.
func InitiatorsForLog(store *store.Store, log types.Log) ([]models.Initiator, error) {
	initrs, merr := ethLogInitrsForAddress(store, log.Address)
	if isRunLog(log) {
		rlInitrs, err := ethLogInitrsForRunLog(store, log)
		initrs = append(initrs, rlInitrs...)
		merr = multierr.Append(merr, err)
	}

	return initrs, merr
}

func ethLogInitrsForAddress(store *store.Store, address common.Address) ([]models.Initiator, error) {
	query := store.Select(q.And(q.Eq("Address", address), q.Re("Type", models.InitiatorEthLog)))
	initrs := []models.Initiator{}
	return initrs, allowNotFoundError(query.Find(&initrs))
}

func ethLogInitrsForRunLog(store *store.Store, log types.Log) ([]models.Initiator, error) {
	initrs := []models.Initiator{}
	if !isRunLog(log) {
		return initrs, nil
	}
	jobID, err := jobIDFromLog(log)
	if err != nil {
		return initrs, err
	}

	query := store.Select(q.And(q.Eq("JobID", jobID), q.Re("Type", models.InitiatorRunLog)))
	if err = query.Find(&initrs); allowNotFoundError(err) != nil {
		return initrs, err
	}
	return initrsForAddress(initrs, log.Address), nil
}

func allowNotFoundError(err error) error {
	if err == storm.ErrNotFound {
		return nil
	}
	return err
}

func isRunLog(log types.Log) bool {
	if len(log.Topics) == 3 && log.Topics[0] == common.StringToHash(RunLogTopic) {
		return true
	}
	return false
}

func jobIDFromLog(log types.Log) (string, error) {
	return utils.HexToString(log.Topics[EventTopicJobID].Hex())
}

func initrsForAddress(initrs []models.Initiator, addr common.Address) []models.Initiator {
	good := []models.Initiator{}
	for _, initr := range initrs {
		if utils.EmptyAddress(initr.Address) || initr.Address == addr {
			good = append(good, initr)
		}
	}
	return good
}
