// Code generated by mockery v2.46.3. DO NOT EDIT.

package prices

import (
	context "context"
	big "math/big"

	mock "github.com/stretchr/testify/mock"
)

// MockGasPriceEstimatorCommit is an autogenerated mock type for the GasPriceEstimatorCommit type
type MockGasPriceEstimatorCommit struct {
	mock.Mock
}

type MockGasPriceEstimatorCommit_Expecter struct {
	mock *mock.Mock
}

func (_m *MockGasPriceEstimatorCommit) EXPECT() *MockGasPriceEstimatorCommit_Expecter {
	return &MockGasPriceEstimatorCommit_Expecter{mock: &_m.Mock}
}

// DenoteInUSD provides a mock function with given fields: ctx, p, wrappedNativePrice
func (_m *MockGasPriceEstimatorCommit) DenoteInUSD(ctx context.Context, p *big.Int, wrappedNativePrice *big.Int) (*big.Int, error) {
	ret := _m.Called(ctx, p, wrappedNativePrice)

	if len(ret) == 0 {
		panic("no return value specified for DenoteInUSD")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int, *big.Int) (*big.Int, error)); ok {
		return rf(ctx, p, wrappedNativePrice)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int, *big.Int) *big.Int); ok {
		r0 = rf(ctx, p, wrappedNativePrice)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Int, *big.Int) error); ok {
		r1 = rf(ctx, p, wrappedNativePrice)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGasPriceEstimatorCommit_DenoteInUSD_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DenoteInUSD'
type MockGasPriceEstimatorCommit_DenoteInUSD_Call struct {
	*mock.Call
}

// DenoteInUSD is a helper method to define mock.On call
//   - ctx context.Context
//   - p *big.Int
//   - wrappedNativePrice *big.Int
func (_e *MockGasPriceEstimatorCommit_Expecter) DenoteInUSD(ctx interface{}, p interface{}, wrappedNativePrice interface{}) *MockGasPriceEstimatorCommit_DenoteInUSD_Call {
	return &MockGasPriceEstimatorCommit_DenoteInUSD_Call{Call: _e.mock.On("DenoteInUSD", ctx, p, wrappedNativePrice)}
}

func (_c *MockGasPriceEstimatorCommit_DenoteInUSD_Call) Run(run func(ctx context.Context, p *big.Int, wrappedNativePrice *big.Int)) *MockGasPriceEstimatorCommit_DenoteInUSD_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*big.Int), args[2].(*big.Int))
	})
	return _c
}

func (_c *MockGasPriceEstimatorCommit_DenoteInUSD_Call) Return(_a0 *big.Int, _a1 error) *MockGasPriceEstimatorCommit_DenoteInUSD_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGasPriceEstimatorCommit_DenoteInUSD_Call) RunAndReturn(run func(context.Context, *big.Int, *big.Int) (*big.Int, error)) *MockGasPriceEstimatorCommit_DenoteInUSD_Call {
	_c.Call.Return(run)
	return _c
}

// Deviates provides a mock function with given fields: ctx, p1, p2
func (_m *MockGasPriceEstimatorCommit) Deviates(ctx context.Context, p1 *big.Int, p2 *big.Int) (bool, error) {
	ret := _m.Called(ctx, p1, p2)

	if len(ret) == 0 {
		panic("no return value specified for Deviates")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int, *big.Int) (bool, error)); ok {
		return rf(ctx, p1, p2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int, *big.Int) bool); ok {
		r0 = rf(ctx, p1, p2)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Int, *big.Int) error); ok {
		r1 = rf(ctx, p1, p2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGasPriceEstimatorCommit_Deviates_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Deviates'
type MockGasPriceEstimatorCommit_Deviates_Call struct {
	*mock.Call
}

// Deviates is a helper method to define mock.On call
//   - ctx context.Context
//   - p1 *big.Int
//   - p2 *big.Int
func (_e *MockGasPriceEstimatorCommit_Expecter) Deviates(ctx interface{}, p1 interface{}, p2 interface{}) *MockGasPriceEstimatorCommit_Deviates_Call {
	return &MockGasPriceEstimatorCommit_Deviates_Call{Call: _e.mock.On("Deviates", ctx, p1, p2)}
}

func (_c *MockGasPriceEstimatorCommit_Deviates_Call) Run(run func(ctx context.Context, p1 *big.Int, p2 *big.Int)) *MockGasPriceEstimatorCommit_Deviates_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*big.Int), args[2].(*big.Int))
	})
	return _c
}

func (_c *MockGasPriceEstimatorCommit_Deviates_Call) Return(_a0 bool, _a1 error) *MockGasPriceEstimatorCommit_Deviates_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGasPriceEstimatorCommit_Deviates_Call) RunAndReturn(run func(context.Context, *big.Int, *big.Int) (bool, error)) *MockGasPriceEstimatorCommit_Deviates_Call {
	_c.Call.Return(run)
	return _c
}

// GetGasPrice provides a mock function with given fields: ctx
func (_m *MockGasPriceEstimatorCommit) GetGasPrice(ctx context.Context) (*big.Int, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetGasPrice")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*big.Int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGasPriceEstimatorCommit_GetGasPrice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGasPrice'
type MockGasPriceEstimatorCommit_GetGasPrice_Call struct {
	*mock.Call
}

// GetGasPrice is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockGasPriceEstimatorCommit_Expecter) GetGasPrice(ctx interface{}) *MockGasPriceEstimatorCommit_GetGasPrice_Call {
	return &MockGasPriceEstimatorCommit_GetGasPrice_Call{Call: _e.mock.On("GetGasPrice", ctx)}
}

func (_c *MockGasPriceEstimatorCommit_GetGasPrice_Call) Run(run func(ctx context.Context)) *MockGasPriceEstimatorCommit_GetGasPrice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockGasPriceEstimatorCommit_GetGasPrice_Call) Return(_a0 *big.Int, _a1 error) *MockGasPriceEstimatorCommit_GetGasPrice_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGasPriceEstimatorCommit_GetGasPrice_Call) RunAndReturn(run func(context.Context) (*big.Int, error)) *MockGasPriceEstimatorCommit_GetGasPrice_Call {
	_c.Call.Return(run)
	return _c
}

// Median provides a mock function with given fields: ctx, gasPrices
func (_m *MockGasPriceEstimatorCommit) Median(ctx context.Context, gasPrices []*big.Int) (*big.Int, error) {
	ret := _m.Called(ctx, gasPrices)

	if len(ret) == 0 {
		panic("no return value specified for Median")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []*big.Int) (*big.Int, error)); ok {
		return rf(ctx, gasPrices)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []*big.Int) *big.Int); ok {
		r0 = rf(ctx, gasPrices)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []*big.Int) error); ok {
		r1 = rf(ctx, gasPrices)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGasPriceEstimatorCommit_Median_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Median'
type MockGasPriceEstimatorCommit_Median_Call struct {
	*mock.Call
}

// Median is a helper method to define mock.On call
//   - ctx context.Context
//   - gasPrices []*big.Int
func (_e *MockGasPriceEstimatorCommit_Expecter) Median(ctx interface{}, gasPrices interface{}) *MockGasPriceEstimatorCommit_Median_Call {
	return &MockGasPriceEstimatorCommit_Median_Call{Call: _e.mock.On("Median", ctx, gasPrices)}
}

func (_c *MockGasPriceEstimatorCommit_Median_Call) Run(run func(ctx context.Context, gasPrices []*big.Int)) *MockGasPriceEstimatorCommit_Median_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]*big.Int))
	})
	return _c
}

func (_c *MockGasPriceEstimatorCommit_Median_Call) Return(_a0 *big.Int, _a1 error) *MockGasPriceEstimatorCommit_Median_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGasPriceEstimatorCommit_Median_Call) RunAndReturn(run func(context.Context, []*big.Int) (*big.Int, error)) *MockGasPriceEstimatorCommit_Median_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockGasPriceEstimatorCommit creates a new instance of MockGasPriceEstimatorCommit. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGasPriceEstimatorCommit(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGasPriceEstimatorCommit {
	mock := &MockGasPriceEstimatorCommit{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
