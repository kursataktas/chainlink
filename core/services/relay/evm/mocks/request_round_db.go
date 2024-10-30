// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	evm "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm"
	mock "github.com/stretchr/testify/mock"

	ocr2aggregator "github.com/smartcontractkit/libocr/gethwrappers2/ocr2aggregator"

	sqlutil "github.com/smartcontractkit/chainlink-common/pkg/sqlutil"
)

// RequestRoundDB is an autogenerated mock type for the RequestRoundDB type
type RequestRoundDB struct {
	mock.Mock
}

type RequestRoundDB_Expecter struct {
	mock *mock.Mock
}

func (_m *RequestRoundDB) EXPECT() *RequestRoundDB_Expecter {
	return &RequestRoundDB_Expecter{mock: &_m.Mock}
}

// LoadLatestRoundRequested provides a mock function with given fields: _a0
func (_m *RequestRoundDB) LoadLatestRoundRequested(_a0 context.Context) (ocr2aggregator.OCR2AggregatorRoundRequested, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for LoadLatestRoundRequested")
	}

	var r0 ocr2aggregator.OCR2AggregatorRoundRequested
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (ocr2aggregator.OCR2AggregatorRoundRequested, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) ocr2aggregator.OCR2AggregatorRoundRequested); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(ocr2aggregator.OCR2AggregatorRoundRequested)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestRoundDB_LoadLatestRoundRequested_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadLatestRoundRequested'
type RequestRoundDB_LoadLatestRoundRequested_Call struct {
	*mock.Call
}

// LoadLatestRoundRequested is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *RequestRoundDB_Expecter) LoadLatestRoundRequested(_a0 interface{}) *RequestRoundDB_LoadLatestRoundRequested_Call {
	return &RequestRoundDB_LoadLatestRoundRequested_Call{Call: _e.mock.On("LoadLatestRoundRequested", _a0)}
}

func (_c *RequestRoundDB_LoadLatestRoundRequested_Call) Run(run func(_a0 context.Context)) *RequestRoundDB_LoadLatestRoundRequested_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *RequestRoundDB_LoadLatestRoundRequested_Call) Return(rr ocr2aggregator.OCR2AggregatorRoundRequested, err error) *RequestRoundDB_LoadLatestRoundRequested_Call {
	_c.Call.Return(rr, err)
	return _c
}

func (_c *RequestRoundDB_LoadLatestRoundRequested_Call) RunAndReturn(run func(context.Context) (ocr2aggregator.OCR2AggregatorRoundRequested, error)) *RequestRoundDB_LoadLatestRoundRequested_Call {
	_c.Call.Return(run)
	return _c
}

// SaveLatestRoundRequested provides a mock function with given fields: ctx, rr
func (_m *RequestRoundDB) SaveLatestRoundRequested(ctx context.Context, rr ocr2aggregator.OCR2AggregatorRoundRequested) error {
	ret := _m.Called(ctx, rr)

	if len(ret) == 0 {
		panic("no return value specified for SaveLatestRoundRequested")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ocr2aggregator.OCR2AggregatorRoundRequested) error); ok {
		r0 = rf(ctx, rr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RequestRoundDB_SaveLatestRoundRequested_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveLatestRoundRequested'
type RequestRoundDB_SaveLatestRoundRequested_Call struct {
	*mock.Call
}

// SaveLatestRoundRequested is a helper method to define mock.On call
//   - ctx context.Context
//   - rr ocr2aggregator.OCR2AggregatorRoundRequested
func (_e *RequestRoundDB_Expecter) SaveLatestRoundRequested(ctx interface{}, rr interface{}) *RequestRoundDB_SaveLatestRoundRequested_Call {
	return &RequestRoundDB_SaveLatestRoundRequested_Call{Call: _e.mock.On("SaveLatestRoundRequested", ctx, rr)}
}

func (_c *RequestRoundDB_SaveLatestRoundRequested_Call) Run(run func(ctx context.Context, rr ocr2aggregator.OCR2AggregatorRoundRequested)) *RequestRoundDB_SaveLatestRoundRequested_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ocr2aggregator.OCR2AggregatorRoundRequested))
	})
	return _c
}

func (_c *RequestRoundDB_SaveLatestRoundRequested_Call) Return(_a0 error) *RequestRoundDB_SaveLatestRoundRequested_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RequestRoundDB_SaveLatestRoundRequested_Call) RunAndReturn(run func(context.Context, ocr2aggregator.OCR2AggregatorRoundRequested) error) *RequestRoundDB_SaveLatestRoundRequested_Call {
	_c.Call.Return(run)
	return _c
}

// WithDataSource provides a mock function with given fields: _a0
func (_m *RequestRoundDB) WithDataSource(_a0 sqlutil.DataSource) evm.RequestRoundDB {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for WithDataSource")
	}

	var r0 evm.RequestRoundDB
	if rf, ok := ret.Get(0).(func(sqlutil.DataSource) evm.RequestRoundDB); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(evm.RequestRoundDB)
		}
	}

	return r0
}

// RequestRoundDB_WithDataSource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithDataSource'
type RequestRoundDB_WithDataSource_Call struct {
	*mock.Call
}

// WithDataSource is a helper method to define mock.On call
//   - _a0 sqlutil.DataSource
func (_e *RequestRoundDB_Expecter) WithDataSource(_a0 interface{}) *RequestRoundDB_WithDataSource_Call {
	return &RequestRoundDB_WithDataSource_Call{Call: _e.mock.On("WithDataSource", _a0)}
}

func (_c *RequestRoundDB_WithDataSource_Call) Run(run func(_a0 sqlutil.DataSource)) *RequestRoundDB_WithDataSource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(sqlutil.DataSource))
	})
	return _c
}

func (_c *RequestRoundDB_WithDataSource_Call) Return(_a0 evm.RequestRoundDB) *RequestRoundDB_WithDataSource_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RequestRoundDB_WithDataSource_Call) RunAndReturn(run func(sqlutil.DataSource) evm.RequestRoundDB) *RequestRoundDB_WithDataSource_Call {
	_c.Call.Return(run)
	return _c
}

// NewRequestRoundDB creates a new instance of RequestRoundDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRequestRoundDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *RequestRoundDB {
	mock := &RequestRoundDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
