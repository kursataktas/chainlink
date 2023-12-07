// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	ethkey "github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ethkey"

	job "github.com/smartcontractkit/chainlink/v2/core/services/job"

	mock "github.com/stretchr/testify/mock"

	pg "github.com/smartcontractkit/chainlink/v2/core/services/pg"

	pipeline "github.com/smartcontractkit/chainlink/v2/core/services/pipeline"

	utils "github.com/smartcontractkit/chainlink/v2/core/utils"

	uuid "github.com/google/uuid"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

// AssertBridgesExist provides a mock function with given fields: p
func (_m *ORM) AssertBridgesExist(p pipeline.Pipeline) error {
	ret := _m.Called(p)

	if len(ret) == 0 {
		panic("no return value specified for AssertBridgesExist")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(pipeline.Pipeline) error); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *ORM) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CountPipelineRunsByJobID provides a mock function with given fields: jobID
func (_m *ORM) CountPipelineRunsByJobID(jobID int32) (int32, error) {
	ret := _m.Called(jobID)

	if len(ret) == 0 {
		panic("no return value specified for CountPipelineRunsByJobID")
	}

	var r0 int32
	var r1 error
	if rf, ok := ret.Get(0).(func(int32) (int32, error)); ok {
		return rf(jobID)
	}
	if rf, ok := ret.Get(0).(func(int32) int32); ok {
		r0 = rf(jobID)
	} else {
		r0 = ret.Get(0).(int32)
	}

	if rf, ok := ret.Get(1).(func(int32) error); ok {
		r1 = rf(jobID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateJob provides a mock function with given fields: jb, qopts
func (_m *ORM) CreateJob(jb *job.Job, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, jb)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateJob")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*job.Job, ...pg.QOpt) error); ok {
		r0 = rf(jb, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteJob provides a mock function with given fields: id, qopts
func (_m *ORM) DeleteJob(id int32, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteJob")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int32, ...pg.QOpt) error); ok {
		r0 = rf(id, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DismissError provides a mock function with given fields: ctx, errorID
func (_m *ORM) DismissError(ctx context.Context, errorID int64) error {
	ret := _m.Called(ctx, errorID)

	if len(ret) == 0 {
		panic("no return value specified for DismissError")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, errorID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindJob provides a mock function with given fields: ctx, id
func (_m *ORM) FindJob(ctx context.Context, id int32) (job.Job, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindJob")
	}

	var r0 job.Job
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) (job.Job, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32) job.Job); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(job.Job)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindJobByExternalJobID provides a mock function with given fields: _a0, qopts
func (_m *ORM) FindJobByExternalJobID(_a0 uuid.UUID, qopts ...pg.QOpt) (job.Job, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FindJobByExternalJobID")
	}

	var r0 job.Job
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, ...pg.QOpt) (job.Job, error)); ok {
		return rf(_a0, qopts...)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID, ...pg.QOpt) job.Job); ok {
		r0 = rf(_a0, qopts...)
	} else {
		r0 = ret.Get(0).(job.Job)
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID, ...pg.QOpt) error); ok {
		r1 = rf(_a0, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindJobIDByAddress provides a mock function with given fields: address, evmChainID, qopts
func (_m *ORM) FindJobIDByAddress(address ethkey.EIP55Address, evmChainID *utils.Big, qopts ...pg.QOpt) (int32, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, address, evmChainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FindJobIDByAddress")
	}

	var r0 int32
	var r1 error
	if rf, ok := ret.Get(0).(func(ethkey.EIP55Address, *utils.Big, ...pg.QOpt) (int32, error)); ok {
		return rf(address, evmChainID, qopts...)
	}
	if rf, ok := ret.Get(0).(func(ethkey.EIP55Address, *utils.Big, ...pg.QOpt) int32); ok {
		r0 = rf(address, evmChainID, qopts...)
	} else {
		r0 = ret.Get(0).(int32)
	}

	if rf, ok := ret.Get(1).(func(ethkey.EIP55Address, *utils.Big, ...pg.QOpt) error); ok {
		r1 = rf(address, evmChainID, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindJobIDsWithBridge provides a mock function with given fields: name
func (_m *ORM) FindJobIDsWithBridge(name string) ([]int32, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for FindJobIDsWithBridge")
	}

	var r0 []int32
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]int32, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) []int32); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int32)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindJobTx provides a mock function with given fields: ctx, id
func (_m *ORM) FindJobTx(ctx context.Context, id int32) (job.Job, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindJobTx")
	}

	var r0 job.Job
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) (job.Job, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32) job.Job); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(job.Job)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindJobWithoutSpecErrors provides a mock function with given fields: id
func (_m *ORM) FindJobWithoutSpecErrors(id int32) (job.Job, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindJobWithoutSpecErrors")
	}

	var r0 job.Job
	var r1 error
	if rf, ok := ret.Get(0).(func(int32) (job.Job, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int32) job.Job); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(job.Job)
	}

	if rf, ok := ret.Get(1).(func(int32) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindJobs provides a mock function with given fields: offset, limit
func (_m *ORM) FindJobs(offset int, limit int) ([]job.Job, int, error) {
	ret := _m.Called(offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for FindJobs")
	}

	var r0 []job.Job
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int) ([]job.Job, int, error)); ok {
		return rf(offset, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int) []job.Job); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]job.Job)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) int); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindJobsByPipelineSpecIDs provides a mock function with given fields: ids
func (_m *ORM) FindJobsByPipelineSpecIDs(ids []int32) ([]job.Job, error) {
	ret := _m.Called(ids)

	if len(ret) == 0 {
		panic("no return value specified for FindJobsByPipelineSpecIDs")
	}

	var r0 []job.Job
	var r1 error
	if rf, ok := ret.Get(0).(func([]int32) ([]job.Job, error)); ok {
		return rf(ids)
	}
	if rf, ok := ret.Get(0).(func([]int32) []job.Job); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]job.Job)
		}
	}

	if rf, ok := ret.Get(1).(func([]int32) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOCR2JobIDByAddress provides a mock function with given fields: contractID, feedID, qopts
func (_m *ORM) FindOCR2JobIDByAddress(contractID string, feedID *common.Hash, qopts ...pg.QOpt) (int32, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, contractID, feedID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FindOCR2JobIDByAddress")
	}

	var r0 int32
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *common.Hash, ...pg.QOpt) (int32, error)); ok {
		return rf(contractID, feedID, qopts...)
	}
	if rf, ok := ret.Get(0).(func(string, *common.Hash, ...pg.QOpt) int32); ok {
		r0 = rf(contractID, feedID, qopts...)
	} else {
		r0 = ret.Get(0).(int32)
	}

	if rf, ok := ret.Get(1).(func(string, *common.Hash, ...pg.QOpt) error); ok {
		r1 = rf(contractID, feedID, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindPipelineRunByID provides a mock function with given fields: id
func (_m *ORM) FindPipelineRunByID(id int64) (pipeline.Run, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindPipelineRunByID")
	}

	var r0 pipeline.Run
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (pipeline.Run, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int64) pipeline.Run); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(pipeline.Run)
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindPipelineRunIDsByJobID provides a mock function with given fields: jobID, offset, limit
func (_m *ORM) FindPipelineRunIDsByJobID(jobID int32, offset int, limit int) ([]int64, error) {
	ret := _m.Called(jobID, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for FindPipelineRunIDsByJobID")
	}

	var r0 []int64
	var r1 error
	if rf, ok := ret.Get(0).(func(int32, int, int) ([]int64, error)); ok {
		return rf(jobID, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(int32, int, int) []int64); ok {
		r0 = rf(jobID, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	if rf, ok := ret.Get(1).(func(int32, int, int) error); ok {
		r1 = rf(jobID, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindPipelineRunsByIDs provides a mock function with given fields: ids
func (_m *ORM) FindPipelineRunsByIDs(ids []int64) ([]pipeline.Run, error) {
	ret := _m.Called(ids)

	if len(ret) == 0 {
		panic("no return value specified for FindPipelineRunsByIDs")
	}

	var r0 []pipeline.Run
	var r1 error
	if rf, ok := ret.Get(0).(func([]int64) ([]pipeline.Run, error)); ok {
		return rf(ids)
	}
	if rf, ok := ret.Get(0).(func([]int64) []pipeline.Run); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pipeline.Run)
		}
	}

	if rf, ok := ret.Get(1).(func([]int64) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindSpecError provides a mock function with given fields: id, qopts
func (_m *ORM) FindSpecError(id int64, qopts ...pg.QOpt) (job.SpecError, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FindSpecError")
	}

	var r0 job.SpecError
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, ...pg.QOpt) (job.SpecError, error)); ok {
		return rf(id, qopts...)
	}
	if rf, ok := ret.Get(0).(func(int64, ...pg.QOpt) job.SpecError); ok {
		r0 = rf(id, qopts...)
	} else {
		r0 = ret.Get(0).(job.SpecError)
	}

	if rf, ok := ret.Get(1).(func(int64, ...pg.QOpt) error); ok {
		r1 = rf(id, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindSpecErrorsByJobIDs provides a mock function with given fields: ids, qopts
func (_m *ORM) FindSpecErrorsByJobIDs(ids []int32, qopts ...pg.QOpt) ([]job.SpecError, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ids)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FindSpecErrorsByJobIDs")
	}

	var r0 []job.SpecError
	var r1 error
	if rf, ok := ret.Get(0).(func([]int32, ...pg.QOpt) ([]job.SpecError, error)); ok {
		return rf(ids, qopts...)
	}
	if rf, ok := ret.Get(0).(func([]int32, ...pg.QOpt) []job.SpecError); ok {
		r0 = rf(ids, qopts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]job.SpecError)
		}
	}

	if rf, ok := ret.Get(1).(func([]int32, ...pg.QOpt) error); ok {
		r1 = rf(ids, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTaskResultByRunIDAndTaskName provides a mock function with given fields: runID, taskName, qopts
func (_m *ORM) FindTaskResultByRunIDAndTaskName(runID int64, taskName string, qopts ...pg.QOpt) ([]byte, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, runID, taskName)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FindTaskResultByRunIDAndTaskName")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, string, ...pg.QOpt) ([]byte, error)); ok {
		return rf(runID, taskName, qopts...)
	}
	if rf, ok := ret.Get(0).(func(int64, string, ...pg.QOpt) []byte); ok {
		r0 = rf(runID, taskName, qopts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(int64, string, ...pg.QOpt) error); ok {
		r1 = rf(runID, taskName, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertJob provides a mock function with given fields: _a0, qopts
func (_m *ORM) InsertJob(_a0 *job.Job, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for InsertJob")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*job.Job, ...pg.QOpt) error); ok {
		r0 = rf(_a0, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertWebhookSpec provides a mock function with given fields: webhookSpec, qopts
func (_m *ORM) InsertWebhookSpec(webhookSpec *job.WebhookSpec, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, webhookSpec)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for InsertWebhookSpec")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*job.WebhookSpec, ...pg.QOpt) error); ok {
		r0 = rf(webhookSpec, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PipelineRuns provides a mock function with given fields: jobID, offset, size
func (_m *ORM) PipelineRuns(jobID *int32, offset int, size int) ([]pipeline.Run, int, error) {
	ret := _m.Called(jobID, offset, size)

	if len(ret) == 0 {
		panic("no return value specified for PipelineRuns")
	}

	var r0 []pipeline.Run
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(*int32, int, int) ([]pipeline.Run, int, error)); ok {
		return rf(jobID, offset, size)
	}
	if rf, ok := ret.Get(0).(func(*int32, int, int) []pipeline.Run); ok {
		r0 = rf(jobID, offset, size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pipeline.Run)
		}
	}

	if rf, ok := ret.Get(1).(func(*int32, int, int) int); ok {
		r1 = rf(jobID, offset, size)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(*int32, int, int) error); ok {
		r2 = rf(jobID, offset, size)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RecordError provides a mock function with given fields: jobID, description, qopts
func (_m *ORM) RecordError(jobID int32, description string, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, jobID, description)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RecordError")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int32, string, ...pg.QOpt) error); ok {
		r0 = rf(jobID, description, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TryRecordError provides a mock function with given fields: jobID, description, qopts
func (_m *ORM) TryRecordError(jobID int32, description string, qopts ...pg.QOpt) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, jobID, description)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// NewORM creates a new instance of ORM. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewORM(t interface {
	mock.TestingT
	Cleanup(func())
}) *ORM {
	mock := &ORM{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
