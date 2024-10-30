// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	fluxmonitorv2 "github.com/smartcontractkit/chainlink/v2/core/services/fluxmonitorv2"

	mock "github.com/stretchr/testify/mock"

	sqlutil "github.com/smartcontractkit/chainlink-common/pkg/sqlutil"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

type ORM_Expecter struct {
	mock *mock.Mock
}

func (_m *ORM) EXPECT() *ORM_Expecter {
	return &ORM_Expecter{mock: &_m.Mock}
}

// CountFluxMonitorRoundStats provides a mock function with given fields: ctx
func (_m *ORM) CountFluxMonitorRoundStats(ctx context.Context) (int, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CountFluxMonitorRoundStats")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ORM_CountFluxMonitorRoundStats_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CountFluxMonitorRoundStats'
type ORM_CountFluxMonitorRoundStats_Call struct {
	*mock.Call
}

// CountFluxMonitorRoundStats is a helper method to define mock.On call
//   - ctx context.Context
func (_e *ORM_Expecter) CountFluxMonitorRoundStats(ctx interface{}) *ORM_CountFluxMonitorRoundStats_Call {
	return &ORM_CountFluxMonitorRoundStats_Call{Call: _e.mock.On("CountFluxMonitorRoundStats", ctx)}
}

func (_c *ORM_CountFluxMonitorRoundStats_Call) Run(run func(ctx context.Context)) *ORM_CountFluxMonitorRoundStats_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ORM_CountFluxMonitorRoundStats_Call) Return(count int, err error) *ORM_CountFluxMonitorRoundStats_Call {
	_c.Call.Return(count, err)
	return _c
}

func (_c *ORM_CountFluxMonitorRoundStats_Call) RunAndReturn(run func(context.Context) (int, error)) *ORM_CountFluxMonitorRoundStats_Call {
	_c.Call.Return(run)
	return _c
}

// CreateEthTransaction provides a mock function with given fields: ctx, fromAddress, toAddress, payload, gasLimit, idempotencyKey
func (_m *ORM) CreateEthTransaction(ctx context.Context, fromAddress common.Address, toAddress common.Address, payload []byte, gasLimit uint64, idempotencyKey *string) error {
	ret := _m.Called(ctx, fromAddress, toAddress, payload, gasLimit, idempotencyKey)

	if len(ret) == 0 {
		panic("no return value specified for CreateEthTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, common.Address, []byte, uint64, *string) error); ok {
		r0 = rf(ctx, fromAddress, toAddress, payload, gasLimit, idempotencyKey)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ORM_CreateEthTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateEthTransaction'
type ORM_CreateEthTransaction_Call struct {
	*mock.Call
}

// CreateEthTransaction is a helper method to define mock.On call
//   - ctx context.Context
//   - fromAddress common.Address
//   - toAddress common.Address
//   - payload []byte
//   - gasLimit uint64
//   - idempotencyKey *string
func (_e *ORM_Expecter) CreateEthTransaction(ctx interface{}, fromAddress interface{}, toAddress interface{}, payload interface{}, gasLimit interface{}, idempotencyKey interface{}) *ORM_CreateEthTransaction_Call {
	return &ORM_CreateEthTransaction_Call{Call: _e.mock.On("CreateEthTransaction", ctx, fromAddress, toAddress, payload, gasLimit, idempotencyKey)}
}

func (_c *ORM_CreateEthTransaction_Call) Run(run func(ctx context.Context, fromAddress common.Address, toAddress common.Address, payload []byte, gasLimit uint64, idempotencyKey *string)) *ORM_CreateEthTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Address), args[2].(common.Address), args[3].([]byte), args[4].(uint64), args[5].(*string))
	})
	return _c
}

func (_c *ORM_CreateEthTransaction_Call) Return(_a0 error) *ORM_CreateEthTransaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ORM_CreateEthTransaction_Call) RunAndReturn(run func(context.Context, common.Address, common.Address, []byte, uint64, *string) error) *ORM_CreateEthTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteFluxMonitorRoundsBackThrough provides a mock function with given fields: ctx, aggregator, roundID
func (_m *ORM) DeleteFluxMonitorRoundsBackThrough(ctx context.Context, aggregator common.Address, roundID uint32) error {
	ret := _m.Called(ctx, aggregator, roundID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteFluxMonitorRoundsBackThrough")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, uint32) error); ok {
		r0 = rf(ctx, aggregator, roundID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ORM_DeleteFluxMonitorRoundsBackThrough_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteFluxMonitorRoundsBackThrough'
type ORM_DeleteFluxMonitorRoundsBackThrough_Call struct {
	*mock.Call
}

// DeleteFluxMonitorRoundsBackThrough is a helper method to define mock.On call
//   - ctx context.Context
//   - aggregator common.Address
//   - roundID uint32
func (_e *ORM_Expecter) DeleteFluxMonitorRoundsBackThrough(ctx interface{}, aggregator interface{}, roundID interface{}) *ORM_DeleteFluxMonitorRoundsBackThrough_Call {
	return &ORM_DeleteFluxMonitorRoundsBackThrough_Call{Call: _e.mock.On("DeleteFluxMonitorRoundsBackThrough", ctx, aggregator, roundID)}
}

func (_c *ORM_DeleteFluxMonitorRoundsBackThrough_Call) Run(run func(ctx context.Context, aggregator common.Address, roundID uint32)) *ORM_DeleteFluxMonitorRoundsBackThrough_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Address), args[2].(uint32))
	})
	return _c
}

func (_c *ORM_DeleteFluxMonitorRoundsBackThrough_Call) Return(_a0 error) *ORM_DeleteFluxMonitorRoundsBackThrough_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ORM_DeleteFluxMonitorRoundsBackThrough_Call) RunAndReturn(run func(context.Context, common.Address, uint32) error) *ORM_DeleteFluxMonitorRoundsBackThrough_Call {
	_c.Call.Return(run)
	return _c
}

// FindOrCreateFluxMonitorRoundStats provides a mock function with given fields: ctx, aggregator, roundID, newRoundLogs
func (_m *ORM) FindOrCreateFluxMonitorRoundStats(ctx context.Context, aggregator common.Address, roundID uint32, newRoundLogs uint) (fluxmonitorv2.FluxMonitorRoundStatsV2, error) {
	ret := _m.Called(ctx, aggregator, roundID, newRoundLogs)

	if len(ret) == 0 {
		panic("no return value specified for FindOrCreateFluxMonitorRoundStats")
	}

	var r0 fluxmonitorv2.FluxMonitorRoundStatsV2
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, uint32, uint) (fluxmonitorv2.FluxMonitorRoundStatsV2, error)); ok {
		return rf(ctx, aggregator, roundID, newRoundLogs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, uint32, uint) fluxmonitorv2.FluxMonitorRoundStatsV2); ok {
		r0 = rf(ctx, aggregator, roundID, newRoundLogs)
	} else {
		r0 = ret.Get(0).(fluxmonitorv2.FluxMonitorRoundStatsV2)
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, uint32, uint) error); ok {
		r1 = rf(ctx, aggregator, roundID, newRoundLogs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ORM_FindOrCreateFluxMonitorRoundStats_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindOrCreateFluxMonitorRoundStats'
type ORM_FindOrCreateFluxMonitorRoundStats_Call struct {
	*mock.Call
}

// FindOrCreateFluxMonitorRoundStats is a helper method to define mock.On call
//   - ctx context.Context
//   - aggregator common.Address
//   - roundID uint32
//   - newRoundLogs uint
func (_e *ORM_Expecter) FindOrCreateFluxMonitorRoundStats(ctx interface{}, aggregator interface{}, roundID interface{}, newRoundLogs interface{}) *ORM_FindOrCreateFluxMonitorRoundStats_Call {
	return &ORM_FindOrCreateFluxMonitorRoundStats_Call{Call: _e.mock.On("FindOrCreateFluxMonitorRoundStats", ctx, aggregator, roundID, newRoundLogs)}
}

func (_c *ORM_FindOrCreateFluxMonitorRoundStats_Call) Run(run func(ctx context.Context, aggregator common.Address, roundID uint32, newRoundLogs uint)) *ORM_FindOrCreateFluxMonitorRoundStats_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Address), args[2].(uint32), args[3].(uint))
	})
	return _c
}

func (_c *ORM_FindOrCreateFluxMonitorRoundStats_Call) Return(_a0 fluxmonitorv2.FluxMonitorRoundStatsV2, _a1 error) *ORM_FindOrCreateFluxMonitorRoundStats_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ORM_FindOrCreateFluxMonitorRoundStats_Call) RunAndReturn(run func(context.Context, common.Address, uint32, uint) (fluxmonitorv2.FluxMonitorRoundStatsV2, error)) *ORM_FindOrCreateFluxMonitorRoundStats_Call {
	_c.Call.Return(run)
	return _c
}

// MostRecentFluxMonitorRoundID provides a mock function with given fields: ctx, aggregator
func (_m *ORM) MostRecentFluxMonitorRoundID(ctx context.Context, aggregator common.Address) (uint32, error) {
	ret := _m.Called(ctx, aggregator)

	if len(ret) == 0 {
		panic("no return value specified for MostRecentFluxMonitorRoundID")
	}

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) (uint32, error)); ok {
		return rf(ctx, aggregator)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) uint32); ok {
		r0 = rf(ctx, aggregator)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address) error); ok {
		r1 = rf(ctx, aggregator)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ORM_MostRecentFluxMonitorRoundID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MostRecentFluxMonitorRoundID'
type ORM_MostRecentFluxMonitorRoundID_Call struct {
	*mock.Call
}

// MostRecentFluxMonitorRoundID is a helper method to define mock.On call
//   - ctx context.Context
//   - aggregator common.Address
func (_e *ORM_Expecter) MostRecentFluxMonitorRoundID(ctx interface{}, aggregator interface{}) *ORM_MostRecentFluxMonitorRoundID_Call {
	return &ORM_MostRecentFluxMonitorRoundID_Call{Call: _e.mock.On("MostRecentFluxMonitorRoundID", ctx, aggregator)}
}

func (_c *ORM_MostRecentFluxMonitorRoundID_Call) Run(run func(ctx context.Context, aggregator common.Address)) *ORM_MostRecentFluxMonitorRoundID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Address))
	})
	return _c
}

func (_c *ORM_MostRecentFluxMonitorRoundID_Call) Return(_a0 uint32, _a1 error) *ORM_MostRecentFluxMonitorRoundID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ORM_MostRecentFluxMonitorRoundID_Call) RunAndReturn(run func(context.Context, common.Address) (uint32, error)) *ORM_MostRecentFluxMonitorRoundID_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateFluxMonitorRoundStats provides a mock function with given fields: ctx, aggregator, roundID, runID, newRoundLogsAddition
func (_m *ORM) UpdateFluxMonitorRoundStats(ctx context.Context, aggregator common.Address, roundID uint32, runID int64, newRoundLogsAddition uint) error {
	ret := _m.Called(ctx, aggregator, roundID, runID, newRoundLogsAddition)

	if len(ret) == 0 {
		panic("no return value specified for UpdateFluxMonitorRoundStats")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, uint32, int64, uint) error); ok {
		r0 = rf(ctx, aggregator, roundID, runID, newRoundLogsAddition)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ORM_UpdateFluxMonitorRoundStats_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateFluxMonitorRoundStats'
type ORM_UpdateFluxMonitorRoundStats_Call struct {
	*mock.Call
}

// UpdateFluxMonitorRoundStats is a helper method to define mock.On call
//   - ctx context.Context
//   - aggregator common.Address
//   - roundID uint32
//   - runID int64
//   - newRoundLogsAddition uint
func (_e *ORM_Expecter) UpdateFluxMonitorRoundStats(ctx interface{}, aggregator interface{}, roundID interface{}, runID interface{}, newRoundLogsAddition interface{}) *ORM_UpdateFluxMonitorRoundStats_Call {
	return &ORM_UpdateFluxMonitorRoundStats_Call{Call: _e.mock.On("UpdateFluxMonitorRoundStats", ctx, aggregator, roundID, runID, newRoundLogsAddition)}
}

func (_c *ORM_UpdateFluxMonitorRoundStats_Call) Run(run func(ctx context.Context, aggregator common.Address, roundID uint32, runID int64, newRoundLogsAddition uint)) *ORM_UpdateFluxMonitorRoundStats_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Address), args[2].(uint32), args[3].(int64), args[4].(uint))
	})
	return _c
}

func (_c *ORM_UpdateFluxMonitorRoundStats_Call) Return(_a0 error) *ORM_UpdateFluxMonitorRoundStats_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ORM_UpdateFluxMonitorRoundStats_Call) RunAndReturn(run func(context.Context, common.Address, uint32, int64, uint) error) *ORM_UpdateFluxMonitorRoundStats_Call {
	_c.Call.Return(run)
	return _c
}

// WithDataSource provides a mock function with given fields: _a0
func (_m *ORM) WithDataSource(_a0 sqlutil.DataSource) fluxmonitorv2.ORM {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for WithDataSource")
	}

	var r0 fluxmonitorv2.ORM
	if rf, ok := ret.Get(0).(func(sqlutil.DataSource) fluxmonitorv2.ORM); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fluxmonitorv2.ORM)
		}
	}

	return r0
}

// ORM_WithDataSource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithDataSource'
type ORM_WithDataSource_Call struct {
	*mock.Call
}

// WithDataSource is a helper method to define mock.On call
//   - _a0 sqlutil.DataSource
func (_e *ORM_Expecter) WithDataSource(_a0 interface{}) *ORM_WithDataSource_Call {
	return &ORM_WithDataSource_Call{Call: _e.mock.On("WithDataSource", _a0)}
}

func (_c *ORM_WithDataSource_Call) Run(run func(_a0 sqlutil.DataSource)) *ORM_WithDataSource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(sqlutil.DataSource))
	})
	return _c
}

func (_c *ORM_WithDataSource_Call) Return(_a0 fluxmonitorv2.ORM) *ORM_WithDataSource_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ORM_WithDataSource_Call) RunAndReturn(run func(sqlutil.DataSource) fluxmonitorv2.ORM) *ORM_WithDataSource_Call {
	_c.Call.Return(run)
	return _c
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
