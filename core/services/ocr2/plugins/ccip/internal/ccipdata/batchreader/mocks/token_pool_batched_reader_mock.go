// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	ccip "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"

	mock "github.com/stretchr/testify/mock"
)

// TokenPoolBatchedReader is an autogenerated mock type for the TokenPoolBatchedReader type
type TokenPoolBatchedReader struct {
	mock.Mock
}

type TokenPoolBatchedReader_Expecter struct {
	mock *mock.Mock
}

func (_m *TokenPoolBatchedReader) EXPECT() *TokenPoolBatchedReader_Expecter {
	return &TokenPoolBatchedReader_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *TokenPoolBatchedReader) Close() error {
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

// TokenPoolBatchedReader_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type TokenPoolBatchedReader_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *TokenPoolBatchedReader_Expecter) Close() *TokenPoolBatchedReader_Close_Call {
	return &TokenPoolBatchedReader_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *TokenPoolBatchedReader_Close_Call) Run(run func()) *TokenPoolBatchedReader_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TokenPoolBatchedReader_Close_Call) Return(_a0 error) *TokenPoolBatchedReader_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TokenPoolBatchedReader_Close_Call) RunAndReturn(run func() error) *TokenPoolBatchedReader_Close_Call {
	_c.Call.Return(run)
	return _c
}

// GetInboundTokenPoolRateLimits provides a mock function with given fields: ctx, tokenPoolReaders
func (_m *TokenPoolBatchedReader) GetInboundTokenPoolRateLimits(ctx context.Context, tokenPoolReaders []ccip.Address) ([]ccip.TokenBucketRateLimit, error) {
	ret := _m.Called(ctx, tokenPoolReaders)

	if len(ret) == 0 {
		panic("no return value specified for GetInboundTokenPoolRateLimits")
	}

	var r0 []ccip.TokenBucketRateLimit
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []ccip.Address) ([]ccip.TokenBucketRateLimit, error)); ok {
		return rf(ctx, tokenPoolReaders)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []ccip.Address) []ccip.TokenBucketRateLimit); ok {
		r0 = rf(ctx, tokenPoolReaders)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ccip.TokenBucketRateLimit)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []ccip.Address) error); ok {
		r1 = rf(ctx, tokenPoolReaders)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TokenPoolBatchedReader_GetInboundTokenPoolRateLimits_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInboundTokenPoolRateLimits'
type TokenPoolBatchedReader_GetInboundTokenPoolRateLimits_Call struct {
	*mock.Call
}

// GetInboundTokenPoolRateLimits is a helper method to define mock.On call
//   - ctx context.Context
//   - tokenPoolReaders []ccip.Address
func (_e *TokenPoolBatchedReader_Expecter) GetInboundTokenPoolRateLimits(ctx interface{}, tokenPoolReaders interface{}) *TokenPoolBatchedReader_GetInboundTokenPoolRateLimits_Call {
	return &TokenPoolBatchedReader_GetInboundTokenPoolRateLimits_Call{Call: _e.mock.On("GetInboundTokenPoolRateLimits", ctx, tokenPoolReaders)}
}

func (_c *TokenPoolBatchedReader_GetInboundTokenPoolRateLimits_Call) Run(run func(ctx context.Context, tokenPoolReaders []ccip.Address)) *TokenPoolBatchedReader_GetInboundTokenPoolRateLimits_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]ccip.Address))
	})
	return _c
}

func (_c *TokenPoolBatchedReader_GetInboundTokenPoolRateLimits_Call) Return(_a0 []ccip.TokenBucketRateLimit, _a1 error) *TokenPoolBatchedReader_GetInboundTokenPoolRateLimits_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TokenPoolBatchedReader_GetInboundTokenPoolRateLimits_Call) RunAndReturn(run func(context.Context, []ccip.Address) ([]ccip.TokenBucketRateLimit, error)) *TokenPoolBatchedReader_GetInboundTokenPoolRateLimits_Call {
	_c.Call.Return(run)
	return _c
}

// NewTokenPoolBatchedReader creates a new instance of TokenPoolBatchedReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTokenPoolBatchedReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *TokenPoolBatchedReader {
	mock := &TokenPoolBatchedReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
