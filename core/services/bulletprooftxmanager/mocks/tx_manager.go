// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	common "github.com/ethereum/go-ethereum/common"
	bulletprooftxmanager "github.com/smartcontractkit/chainlink/core/services/bulletprooftxmanager"

	context "context"

	gas "github.com/smartcontractkit/chainlink/core/services/gas"

	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	models "github.com/smartcontractkit/chainlink/core/store/models"
)

// TxManager is an autogenerated mock type for the TxManager type
type TxManager struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *TxManager) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateEthTransaction provides a mock function with given fields: db, fromAddress, toAddress, payload, gasLimit, meta, strategy
func (_m *TxManager) CreateEthTransaction(db *gorm.DB, fromAddress common.Address, toAddress common.Address, payload []byte, gasLimit uint64, meta interface{}, strategy bulletprooftxmanager.TxStrategy) (bulletprooftxmanager.EthTx, error) {
	ret := _m.Called(db, fromAddress, toAddress, payload, gasLimit, meta, strategy)

	var r0 bulletprooftxmanager.EthTx
	if rf, ok := ret.Get(0).(func(*gorm.DB, common.Address, common.Address, []byte, uint64, interface{}, bulletprooftxmanager.TxStrategy) bulletprooftxmanager.EthTx); ok {
		r0 = rf(db, fromAddress, toAddress, payload, gasLimit, meta, strategy)
	} else {
		r0 = ret.Get(0).(bulletprooftxmanager.EthTx)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, common.Address, common.Address, []byte, uint64, interface{}, bulletprooftxmanager.TxStrategy) error); ok {
		r1 = rf(db, fromAddress, toAddress, payload, gasLimit, meta, strategy)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGasEstimator provides a mock function with given fields:
func (_m *TxManager) GetGasEstimator() gas.Estimator {
	ret := _m.Called()

	var r0 gas.Estimator
	if rf, ok := ret.Get(0).(func() gas.Estimator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gas.Estimator)
		}
	}

	return r0
}

// Healthy provides a mock function with given fields:
func (_m *TxManager) Healthy() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnNewLongestChain provides a mock function with given fields: ctx, head
func (_m *TxManager) OnNewLongestChain(ctx context.Context, head models.Head) {
	_m.Called(ctx, head)
}

// Ready provides a mock function with given fields:
func (_m *TxManager) Ready() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields:
func (_m *TxManager) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Trigger provides a mock function with given fields: addr
func (_m *TxManager) Trigger(addr common.Address) {
	_m.Called(addr)
}
