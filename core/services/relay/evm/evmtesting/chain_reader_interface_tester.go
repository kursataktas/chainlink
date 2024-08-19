package evmtesting

import (
	"context"
	"encoding/json"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/jmoiron/sqlx"
	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/smartcontractkit/chainlink-common/pkg/codec"
	clcommontypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	. "github.com/smartcontractkit/chainlink-common/pkg/types/interfacetests" //nolint common practice to import test mods with .

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/headtracker"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	evmtxmgr "github.com/smartcontractkit/chainlink/v2/core/chains/evm/txmgr"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/chain_reader_tester"
	_ "github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest" // force binding for tx type
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

const (
	triggerWithDynamicTopic        = "TriggeredEventWithDynamicTopic"
	triggerWithAllTopics           = "TriggeredWithFourTopics"
	triggerWithAllTopicsWithHashed = "TriggeredWithFourTopicsWithHashed"
	finalityDepth                  = 4
)

type EVMChainReaderInterfaceTesterHelper[T TestingT[T]] interface {
	Init(t T)
	Client(t T) client.Client
	Commit()
	Backend() bind.ContractBackend
	ChainID() *big.Int
	Context(t T) context.Context
	NewSqlxDB(t T) *sqlx.DB
	MaxWaitTimeForEvents() time.Duration
	GasPriceBufferPercent() int64
	Accounts(t T) []*bind.TransactOpts
	TXM(T, client.Client) evmtxmgr.TxManager
	// To enable the historical wrappers required for Simulated Backend tests.
	ChainReaderEVMClient(t T, ctx context.Context, ht logpoller.HeadTracker, conf types.ChainReaderConfig) client.Client
	WrappedChainWriter(cw clcommontypes.ChainWriter, client client.Client) clcommontypes.ChainWriter
}

type EVMChainReaderInterfaceTester[T TestingT[T]] struct {
	Helper            EVMChainReaderInterfaceTesterHelper[T]
	client            client.Client
	address           string
	address2          string
	contractTesters   map[string]*chain_reader_tester.ChainReaderTester
	chainReaderConfig types.ChainReaderConfig
	chainWriterConfig types.ChainWriterConfig
	deployerAuth      *bind.TransactOpts
	senderAuth        *bind.TransactOpts
	dirtyContracts    bool
	evmTest           *chain_reader_tester.ChainReaderTester
	cr                evm.ChainReaderService
	cw                evm.ChainWriterService
	txm               evmtxmgr.TxManager
	gasEstimator      gas.EvmFeeEstimator
}

func (it *EVMChainReaderInterfaceTester[T]) Setup(t T) {
	t.Cleanup(func() {
		// DB may be closed by the test already, ignore errors
		if it.cr != nil {
			_ = it.cr.Close()
		}
		it.cr = nil

		if it.cw != nil {
			_ = it.cw.Close()
		}
		it.cw = nil

		it.contractTesters = nil
	})

	// can re-use the same chain for tests, just make new contract for each test
	if it.client != nil {
		it.deployNewContracts(t)
		return
	}

	// Need to seperate accounts to ensure the nonce doesn't get misaligned after the
	// contract deployments.
	accounts := it.Helper.Accounts(t)
	it.deployerAuth = accounts[0]
	it.senderAuth = accounts[1]

	testStruct := CreateTestStruct[T](0, it)

	methodTakingLatestParamsReturningTestStructConfig := types.ChainReaderDefinition{
		ChainSpecificName: "getElementAtIndex",
		OutputModifications: codec.ModifiersConfig{
			&codec.RenameModifierConfig{Fields: map[string]string{"NestedStruct.Inner.IntVal": "I"}},
		},
	}

	it.chainReaderConfig = types.ChainReaderConfig{
		Contracts: map[string]types.ChainContractReader{
			AnyContractName: {
				ContractABI: chain_reader_tester.ChainReaderTesterMetaData.ABI,
				ContractPollingFilter: types.ContractPollingFilter{
					GenericEventNames: []string{EventName, EventWithFilterName, triggerWithAllTopicsWithHashed},
				},
				Configs: map[string]*types.ChainReaderDefinition{
					MethodTakingLatestParamsReturningTestStruct: &methodTakingLatestParamsReturningTestStructConfig,
					MethodReturningAlterableUint64: {
						ChainSpecificName: "getAlterablePrimitiveValue",
					},
					MethodReturningUint64: {
						ChainSpecificName: "getPrimitiveValue",
					},
					MethodReturningUint64Slice: {
						ChainSpecificName: "getSliceValue",
					},
					EventName: {
						ChainSpecificName: "Triggered",
						ReadType:          types.Event,
						OutputModifications: codec.ModifiersConfig{
							&codec.RenameModifierConfig{Fields: map[string]string{"NestedStruct.Inner.IntVal": "I"}},
						},
					},
					EventWithFilterName: {
						ChainSpecificName: "Triggered",
						ReadType:          types.Event,
						EventDefinitions:  &types.EventDefinitions{InputFields: []string{"Field"}},
					},
					triggerWithDynamicTopic: {
						ChainSpecificName: triggerWithDynamicTopic,
						ReadType:          types.Event,
						EventDefinitions: &types.EventDefinitions{
							InputFields: []string{"fieldHash"},
							// No specific reason for filter being defined here instead of on contract level, this is just for test case variety.
							PollingFilter: &types.PollingFilter{},
						},
						InputModifications: codec.ModifiersConfig{
							&codec.RenameModifierConfig{Fields: map[string]string{"FieldHash": "Field"}},
						},
					},
					triggerWithAllTopics: {
						ChainSpecificName: triggerWithAllTopics,
						ReadType:          types.Event,
						EventDefinitions: &types.EventDefinitions{
							InputFields:   []string{"Field1", "Field2", "Field3"},
							PollingFilter: &types.PollingFilter{},
						},
						// This doesn't have to be here, since the defalt mapping would work, but is left as an example.
						// Keys which are string float values(confidence levels) are chain agnostic and should be reused across chains.
						// These float values can map to different finality concepts across chains.
						ConfidenceConfirmations: map[string]int{"0.0": int(evmtypes.Unconfirmed), "1.0": int(evmtypes.Finalized)},
					},
					triggerWithAllTopicsWithHashed: {
						ChainSpecificName: triggerWithAllTopicsWithHashed,
						ReadType:          types.Event,
						EventDefinitions: &types.EventDefinitions{
							InputFields: []string{"Field1", "Field2", "Field3"},
						},
					},
					MethodReturningSeenStruct: {
						ChainSpecificName: "returnSeen",
						InputModifications: codec.ModifiersConfig{
							&codec.HardCodeModifierConfig{
								OnChainValues: map[string]any{
									"BigField": testStruct.BigField.String(),
									"Account":  hexutil.Encode(testStruct.Account),
								},
							},
							&codec.RenameModifierConfig{Fields: map[string]string{"NestedStruct.Inner.IntVal": "I"}},
						},
						OutputModifications: codec.ModifiersConfig{
							&codec.HardCodeModifierConfig{OffChainValues: map[string]any{"ExtraField": AnyExtraValue}},
							&codec.RenameModifierConfig{Fields: map[string]string{"NestedStruct.Inner.IntVal": "I"}},
						},
					},
				},
			},
			AnySecondContractName: {
				ContractABI: chain_reader_tester.ChainReaderTesterMetaData.ABI,
				Configs: map[string]*types.ChainReaderDefinition{
					MethodTakingLatestParamsReturningTestStruct: &methodTakingLatestParamsReturningTestStructConfig,
					MethodReturningUint64: {
						ChainSpecificName: "getDifferentPrimitiveValue",
					},
				},
			},
		},
	}
	it.GetChainReader(t)
	it.txm = it.Helper.TXM(t, it.client)

	it.chainWriterConfig = types.ChainWriterConfig{
		Contracts: map[string]*types.ContractConfig{
			AnyContractName: {
				ContractABI: chain_reader_tester.ChainReaderTesterMetaData.ABI,
				Configs: map[string]*types.ChainWriterDefinition{
					"addTestStruct": {
						ChainSpecificName: "addTestStruct",
						FromAddress:       it.Helper.Accounts(t)[1].From,
						GasLimit:          2_000_000,
						Checker:           "simulate",
						InputModifications: codec.ModifiersConfig{
							&codec.RenameModifierConfig{Fields: map[string]string{"NestedStruct.Inner.IntVal": "I"}},
						},
					},
					"setAlterablePrimitiveValue": {
						ChainSpecificName: "setAlterablePrimitiveValue",
						FromAddress:       it.Helper.Accounts(t)[1].From,
						GasLimit:          2_000_000,
						Checker:           "simulate",
					},
					"triggerEvent": {
						ChainSpecificName: "triggerEvent",
						FromAddress:       it.Helper.Accounts(t)[1].From,
						GasLimit:          2_000_000,
						Checker:           "simulate",
						InputModifications: codec.ModifiersConfig{
							&codec.RenameModifierConfig{Fields: map[string]string{"NestedStruct.Inner.IntVal": "I"}},
						},
					},
					"triggerEventWithDynamicTopic": {
						ChainSpecificName: "triggerEventWithDynamicTopic",
						FromAddress:       it.Helper.Accounts(t)[1].From,
						GasLimit:          2_000_000,
						Checker:           "simulate",
					},
					"triggerWithFourTopics": {
						ChainSpecificName: "triggerWithFourTopics",
						FromAddress:       it.Helper.Accounts(t)[1].From,
						GasLimit:          2_000_000,
						Checker:           "simulate",
					},
					"triggerWithFourTopicsWithHashed": {
						ChainSpecificName: "triggerWithFourTopicsWithHashed",
						FromAddress:       it.Helper.Accounts(t)[1].From,
						GasLimit:          2_000_000,
						Checker:           "simulate",
					},
				},
			},
			AnySecondContractName: {
				ContractABI: chain_reader_tester.ChainReaderTesterMetaData.ABI,
				Configs: map[string]*types.ChainWriterDefinition{
					"addTestStruct": {
						ChainSpecificName: "addTestStruct",
						FromAddress:       it.Helper.Accounts(t)[1].From,
						GasLimit:          2_000_000,
						Checker:           "simulate",
						InputModifications: codec.ModifiersConfig{
							&codec.RenameModifierConfig{Fields: map[string]string{"NestedStruct.Inner.IntVal": "I"}},
						},
					},
				},
			},
		},
		MaxGasPrice: assets.NewWei(big.NewInt(1000000000000000000)),
	}
	it.deployNewContracts(t)
}

func (it *EVMChainReaderInterfaceTester[T]) Name() string {
	return "EVM"
}

func (it *EVMChainReaderInterfaceTester[T]) GetAccountBytes(i int) []byte {
	account := [20]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	account[i%20] += byte(i)
	account[(i+3)%20] += byte(i + 3)
	return account[:]
}

func (it *EVMChainReaderInterfaceTester[T]) GetChainReader(t T) clcommontypes.ContractReader {
	ctx := it.Helper.Context(t)
	if it.cr != nil {
		return it.cr
	}

	lggr := logger.NullLogger
	db := it.Helper.NewSqlxDB(t)
	lpOpts := logpoller.Opts{
		PollPeriod:               time.Millisecond,
		FinalityDepth:            finalityDepth,
		BackfillBatchSize:        1,
		RpcBatchSize:             1,
		KeepFinalizedBlocksDepth: 10000,
	}
	ht := headtracker.NewSimulatedHeadTracker(it.Helper.Client(t), lpOpts.UseFinalityTag, lpOpts.FinalityDepth)
	lp := logpoller.NewLogPoller(logpoller.NewORM(it.Helper.ChainID(), db, lggr), it.Helper.Client(t), lggr, ht, lpOpts)
	require.NoError(t, lp.Start(ctx))

	// encode and decode the config to ensure the test covers type issues
	confBytes, err := json.Marshal(it.chainReaderConfig)
	require.NoError(t, err)

	conf, err := types.ChainReaderConfigFromBytes(confBytes)
	require.NoError(t, err)

	cwh := it.Helper.ChainReaderEVMClient(t, ctx, ht, conf)
	it.client = cwh

	cr, err := evm.NewChainReaderService(ctx, lggr, lp, ht, it.client, conf)
	require.NoError(t, err)
	require.NoError(t, cr.Start(ctx))
	it.cr = cr
	return cr
}

func (it *EVMChainReaderInterfaceTester[T]) SetBatchLatestValues(t T, batchCallEntry BatchCallEntry) {
	nameToAddress := make(map[string]string)
	boundContracts := it.GetBindings(t)
	for _, bc := range boundContracts {
		nameToAddress[bc.Name] = bc.Address
	}

	for contractName, contractBatch := range batchCallEntry {
		require.Contains(t, nameToAddress, contractName)
		for _, readEntry := range contractBatch {
			val, isOk := readEntry.ReturnValue.(*TestStruct)
			if !isOk {
				require.Fail(t, "expected *TestStruct for contract: %s read: %s, but received %T", contractName, readEntry.Name, readEntry.ReturnValue)
			}
			it.sendTxWithTestStruct(t, nameToAddress[contractName], val, (*chain_reader_tester.ChainReaderTesterTransactor).AddTestStruct)
		}
	}
}

func (it *EVMChainReaderInterfaceTester[T]) GetChainWriter(t T) clcommontypes.ChainWriter {
	ctx := it.Helper.Context(t)
	if it.cw != nil {
		return it.cw
	}

	cw, err := evm.NewChainWriterService(logger.NullLogger, it.client, it.txm, it.gasEstimator, it.chainWriterConfig)
	require.NoError(t, err)
	it.cw = it.Helper.WrappedChainWriter(cw, it.client)

	require.NoError(t, err)
	require.NoError(t, cw.Start(ctx))
	return it.cw
}

func (it *EVMChainReaderInterfaceTester[T]) GetBindings(_ T) []clcommontypes.BoundContract {
	return []clcommontypes.BoundContract{
		{Name: AnyContractName, Address: it.address},
		{Name: AnySecondContractName, Address: it.address2},
	}
}

type uintFn = func(*chain_reader_tester.ChainReaderTesterTransactor, *bind.TransactOpts, uint64) (*gethtypes.Transaction, error)

type testStructFn = func(*chain_reader_tester.ChainReaderTesterTransactor, *bind.TransactOpts, int32, string, uint8, [32]uint8, common.Address, []common.Address, *big.Int, chain_reader_tester.MidLevelTestStruct) (*gethtypes.Transaction, error)

func (it *EVMChainReaderInterfaceTester[T]) sendTxWithTestStruct(t T, contractAddress string, testStruct *TestStruct, fn testStructFn) {
	tx, err := fn(
		&it.contractTesters[contractAddress].ChainReaderTesterTransactor,
		it.GetAuthWithGasSet(t),
		*testStruct.Field,
		testStruct.DifferentField,
		uint8(testStruct.OracleID),
		OracleIdsToBytes(testStruct.OracleIDs),
		common.Address(testStruct.Account),
		ConvertAccounts(testStruct.Accounts),
		testStruct.BigField,
		MidToInternalType(testStruct.NestedStruct),
	)
	require.NoError(t, err)
	it.Helper.Commit()
	it.IncNonce()
	it.AwaitTx(t, tx)
}

func (it *EVMChainReaderInterfaceTester[T]) GetAuthWithGasSet(t T) *bind.TransactOpts {
	gasPrice, err := it.client.SuggestGasPrice(it.Helper.Context(t))
	require.NoError(t, err)
	extra := new(big.Int).Mul(gasPrice, big.NewInt(it.Helper.GasPriceBufferPercent()))
	extra = extra.Div(extra, big.NewInt(100))
	it.deployerAuth.GasPrice = gasPrice.Add(gasPrice, extra)
	return it.deployerAuth
}

func (it *EVMChainReaderInterfaceTester[T]) IncNonce() {
	if it.deployerAuth.Nonce == nil {
		it.deployerAuth.Nonce = big.NewInt(1)
	} else {
		it.deployerAuth.Nonce = it.deployerAuth.Nonce.Add(it.deployerAuth.Nonce, big.NewInt(1))
	}
}

func (it *EVMChainReaderInterfaceTester[T]) AwaitTx(t T, tx *gethtypes.Transaction) {
	ctx := it.Helper.Context(t)
	receipt, err := bind.WaitMined(ctx, it.client, tx)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)
}

func (it *EVMChainReaderInterfaceTester[T]) deployNewContracts(t T) {
	// First test deploy both contracts, otherwise only deploy contracts if cleanup decides that we need to.
	if it.address == "" || it.contractTesters == nil {
		it.contractTesters = make(map[string]*chain_reader_tester.ChainReaderTester, 2)
		address, ts1 := it.deployNewContract(t)
		address2, ts2 := it.deployNewContract(t)
		it.address, it.address2 = address, address2
		it.contractTesters[it.address] = ts1
		it.contractTesters[it.address2] = ts2
	}
}

func (it *EVMChainReaderInterfaceTester[T]) deployNewContract(t T) (string, *chain_reader_tester.ChainReaderTester) {
	// 105528 was in the error: gas too low: have 0, want 105528
	// Not sure if there's a better way to get it.
	it.deployerAuth.GasLimit = 10552800

	address, tx, ts, err := chain_reader_tester.DeployChainReaderTester(it.GetAuthWithGasSet(t), it.Helper.Backend())
	require.NoError(t, err)
	it.Helper.Commit()

	it.IncNonce()
	it.AwaitTx(t, tx)
	return address.String(), ts
}

func (it *EVMChainReaderInterfaceTester[T]) MaxWaitTimeForEvents() time.Duration {
	return it.Helper.MaxWaitTimeForEvents()
}

func OracleIdsToBytes(oracleIDs [32]commontypes.OracleID) [32]byte {
	convertedIds := [32]byte{}
	for i, id := range oracleIDs {
		convertedIds[i] = byte(id)
	}
	return convertedIds
}

func ConvertAccounts(accounts [][]byte) []common.Address {
	convertedAccounts := make([]common.Address, len(accounts))
	for i, a := range accounts {
		convertedAccounts[i] = common.Address(a)
	}
	return convertedAccounts
}

func ToInternalType(testStruct TestStruct) chain_reader_tester.TestStruct {
	return chain_reader_tester.TestStruct{
		Field:          *testStruct.Field,
		DifferentField: testStruct.DifferentField,
		OracleId:       byte(testStruct.OracleID),
		OracleIds:      OracleIdsToBytes(testStruct.OracleIDs),
		Account:        common.Address(testStruct.Account),
		Accounts:       ConvertAccounts(testStruct.Accounts),
		BigField:       testStruct.BigField,
		NestedStruct:   MidToInternalType(testStruct.NestedStruct),
	}
}

func MidToInternalType(m MidLevelTestStruct) chain_reader_tester.MidLevelTestStruct {
	return chain_reader_tester.MidLevelTestStruct{
		FixedBytes: m.FixedBytes,
		Inner: chain_reader_tester.InnerTestStruct{
			IntVal: int64(m.Inner.I),
			S:      m.Inner.S,
		},
	}
}
