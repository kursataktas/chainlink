// Code generated by mockery v2.21.2. DO NOT EDIT.

package mocks

import (
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"
	mock "github.com/stretchr/testify/mock"
)

// GethKeyStore is an autogenerated mock type for the GethKeyStore type
type GethKeyStore struct {
	mock.Mock
}

// GetRoundRobinAddress provides a mock function with given fields: chainID, addresses
func (_m *GethKeyStore) GetRoundRobinAddress(chainID *big.Int, addresses ...common.Address) (common.Address, error) {
	_va := make([]interface{}, len(addresses))
	for _i := range addresses {
		_va[_i] = addresses[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 common.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(*big.Int, ...common.Address) (common.Address, error)); ok {
		return rf(chainID, addresses...)
	}
	if rf, ok := ret.Get(0).(func(*big.Int, ...common.Address) common.Address); ok {
		r0 = rf(chainID, addresses...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(*big.Int, ...common.Address) error); ok {
		r1 = rf(chainID, addresses...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewGethKeyStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewGethKeyStore creates a new instance of GethKeyStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGethKeyStore(t mockConstructorTestingTNewGethKeyStore) *GethKeyStore {
	mock := &GethKeyStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
