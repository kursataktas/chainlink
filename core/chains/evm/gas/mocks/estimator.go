// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	gas "github.com/smartcontractkit/chainlink/core/chains/evm/gas"
	mock "github.com/stretchr/testify/mock"

	testing "testing"

	types "github.com/smartcontractkit/chainlink/core/chains/evm/types"
)

// Estimator is an autogenerated mock type for the Estimator type
type Estimator struct {
	mock.Mock
}

// BumpDynamicFee provides a mock function with given fields: original, gasLimit
func (_m *Estimator) BumpDynamicFee(original gas.DynamicFee, gasLimit uint64) (gas.DynamicFee, uint64, error) {
	ret := _m.Called(original, gasLimit)

	var r0 gas.DynamicFee
	if rf, ok := ret.Get(0).(func(gas.DynamicFee, uint64) gas.DynamicFee); ok {
		r0 = rf(original, gasLimit)
	} else {
		r0 = ret.Get(0).(gas.DynamicFee)
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(gas.DynamicFee, uint64) uint64); ok {
		r1 = rf(original, gasLimit)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(gas.DynamicFee, uint64) error); ok {
		r2 = rf(original, gasLimit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// BumpLegacyGas provides a mock function with given fields: originalGasPrice, gasLimit
func (_m *Estimator) BumpLegacyGas(originalGasPrice *big.Int, gasLimit uint64) (*big.Int, uint64, error) {
	ret := _m.Called(originalGasPrice, gasLimit)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*big.Int, uint64) *big.Int); ok {
		r0 = rf(originalGasPrice, gasLimit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(*big.Int, uint64) uint64); ok {
		r1 = rf(originalGasPrice, gasLimit)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*big.Int, uint64) error); ok {
		r2 = rf(originalGasPrice, gasLimit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Close provides a mock function with given fields:
func (_m *Estimator) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDynamicFee provides a mock function with given fields: gasLimit
func (_m *Estimator) GetDynamicFee(gasLimit uint64) (gas.DynamicFee, uint64, error) {
	ret := _m.Called(gasLimit)

	var r0 gas.DynamicFee
	if rf, ok := ret.Get(0).(func(uint64) gas.DynamicFee); ok {
		r0 = rf(gasLimit)
	} else {
		r0 = ret.Get(0).(gas.DynamicFee)
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(uint64) uint64); ok {
		r1 = rf(gasLimit)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(uint64) error); ok {
		r2 = rf(gasLimit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetLegacyGas provides a mock function with given fields: calldata, gasLimit, opts
func (_m *Estimator) GetLegacyGas(calldata []byte, gasLimit uint64, opts ...gas.Opt) (*big.Int, uint64, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, calldata, gasLimit)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func([]byte, uint64, ...gas.Opt) *big.Int); ok {
		r0 = rf(calldata, gasLimit, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func([]byte, uint64, ...gas.Opt) uint64); ok {
		r1 = rf(calldata, gasLimit, opts...)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func([]byte, uint64, ...gas.Opt) error); ok {
		r2 = rf(calldata, gasLimit, opts...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// OnNewLongestChain provides a mock function with given fields: _a0, _a1
func (_m *Estimator) OnNewLongestChain(_a0 context.Context, _a1 *types.Head) {
	_m.Called(_a0, _a1)
}

// Start provides a mock function with given fields: _a0
func (_m *Estimator) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewEstimator creates a new instance of Estimator. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewEstimator(t testing.TB) *Estimator {
	mock := &Estimator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
