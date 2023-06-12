// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"

	ethkey "github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ethkey"

	mock "github.com/stretchr/testify/mock"

	pg "github.com/smartcontractkit/chainlink/v2/core/services/pg"

	types "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
)

// Eth is an autogenerated mock type for the Eth type
type Eth struct {
	mock.Mock
}

// CheckEnabled provides a mock function with given fields: address, chainID
func (_m *Eth) CheckEnabled(address common.Address, chainID *big.Int) error {
	ret := _m.Called(address, chainID)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.Address, *big.Int) error); ok {
		r0 = rf(address, chainID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: chainIDs
func (_m *Eth) Create(chainIDs ...*big.Int) (ethkey.KeyV2, error) {
	_va := make([]interface{}, len(chainIDs))
	for _i := range chainIDs {
		_va[_i] = chainIDs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 ethkey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func(...*big.Int) (ethkey.KeyV2, error)); ok {
		return rf(chainIDs...)
	}
	if rf, ok := ret.Get(0).(func(...*big.Int) ethkey.KeyV2); ok {
		r0 = rf(chainIDs...)
	} else {
		r0 = ret.Get(0).(ethkey.KeyV2)
	}

	if rf, ok := ret.Get(1).(func(...*big.Int) error); ok {
		r1 = rf(chainIDs...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Eth) Delete(id string) (ethkey.KeyV2, error) {
	ret := _m.Called(id)

	var r0 ethkey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (ethkey.KeyV2, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) ethkey.KeyV2); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(ethkey.KeyV2)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Disable provides a mock function with given fields: address, chainID, qopts
func (_m *Eth) Disable(address common.Address, chainID *big.Int, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, address, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.Address, *big.Int, ...pg.QOpt) error); ok {
		r0 = rf(address, chainID, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Enable provides a mock function with given fields: address, chainID, qopts
func (_m *Eth) Enable(address common.Address, chainID *big.Int, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, address, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.Address, *big.Int, ...pg.QOpt) error); ok {
		r0 = rf(address, chainID, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnabledAddressesForChain provides a mock function with given fields: chainID
func (_m *Eth) EnabledAddressesForChain(chainID *big.Int) ([]common.Address, error) {
	ret := _m.Called(chainID)

	var r0 []common.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(*big.Int) ([]common.Address, error)); ok {
		return rf(chainID)
	}
	if rf, ok := ret.Get(0).(func(*big.Int) []common.Address); ok {
		r0 = rf(chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(*big.Int) error); ok {
		r1 = rf(chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnabledKeysForChain provides a mock function with given fields: chainID
func (_m *Eth) EnabledKeysForChain(chainID *big.Int) ([]ethkey.KeyV2, error) {
	ret := _m.Called(chainID)

	var r0 []ethkey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func(*big.Int) ([]ethkey.KeyV2, error)); ok {
		return rf(chainID)
	}
	if rf, ok := ret.Get(0).(func(*big.Int) []ethkey.KeyV2); ok {
		r0 = rf(chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ethkey.KeyV2)
		}
	}

	if rf, ok := ret.Get(1).(func(*big.Int) error); ok {
		r1 = rf(chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnsureKeys provides a mock function with given fields: chainIDs
func (_m *Eth) EnsureKeys(chainIDs ...*big.Int) error {
	_va := make([]interface{}, len(chainIDs))
	for _i := range chainIDs {
		_va[_i] = chainIDs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...*big.Int) error); ok {
		r0 = rf(chainIDs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Export provides a mock function with given fields: id, password
func (_m *Eth) Export(id string, password string) ([]byte, error) {
	ret := _m.Called(id, password)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]byte, error)); ok {
		return rf(id, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) []byte); ok {
		r0 = rf(id, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *Eth) Get(id string) (ethkey.KeyV2, error) {
	ret := _m.Called(id)

	var r0 ethkey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (ethkey.KeyV2, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) ethkey.KeyV2); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(ethkey.KeyV2)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *Eth) GetAll() ([]ethkey.KeyV2, error) {
	ret := _m.Called()

	var r0 []ethkey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]ethkey.KeyV2, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []ethkey.KeyV2); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ethkey.KeyV2)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoundRobinAddress provides a mock function with given fields: chainID, addresses
func (_m *Eth) GetRoundRobinAddress(chainID *big.Int, addresses ...common.Address) (common.Address, error) {
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

// GetState provides a mock function with given fields: id, chainID
func (_m *Eth) GetState(id string, chainID *big.Int) (ethkey.State, error) {
	ret := _m.Called(id, chainID)

	var r0 ethkey.State
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *big.Int) (ethkey.State, error)); ok {
		return rf(id, chainID)
	}
	if rf, ok := ret.Get(0).(func(string, *big.Int) ethkey.State); ok {
		r0 = rf(id, chainID)
	} else {
		r0 = ret.Get(0).(ethkey.State)
	}

	if rf, ok := ret.Get(1).(func(string, *big.Int) error); ok {
		r1 = rf(id, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStatesForChain provides a mock function with given fields: chainID
func (_m *Eth) GetStatesForChain(chainID *big.Int) ([]ethkey.State, error) {
	ret := _m.Called(chainID)

	var r0 []ethkey.State
	var r1 error
	if rf, ok := ret.Get(0).(func(*big.Int) ([]ethkey.State, error)); ok {
		return rf(chainID)
	}
	if rf, ok := ret.Get(0).(func(*big.Int) []ethkey.State); ok {
		r0 = rf(chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ethkey.State)
		}
	}

	if rf, ok := ret.Get(1).(func(*big.Int) error); ok {
		r1 = rf(chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStatesForKeys provides a mock function with given fields: _a0
func (_m *Eth) GetStatesForKeys(_a0 []ethkey.KeyV2) ([]ethkey.State, error) {
	ret := _m.Called(_a0)

	var r0 []ethkey.State
	var r1 error
	if rf, ok := ret.Get(0).(func([]ethkey.KeyV2) ([]ethkey.State, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func([]ethkey.KeyV2) []ethkey.State); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ethkey.State)
		}
	}

	if rf, ok := ret.Get(1).(func([]ethkey.KeyV2) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Import provides a mock function with given fields: keyJSON, password, chainIDs
func (_m *Eth) Import(keyJSON []byte, password string, chainIDs ...*big.Int) (ethkey.KeyV2, error) {
	_va := make([]interface{}, len(chainIDs))
	for _i := range chainIDs {
		_va[_i] = chainIDs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, keyJSON, password)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 ethkey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, string, ...*big.Int) (ethkey.KeyV2, error)); ok {
		return rf(keyJSON, password, chainIDs...)
	}
	if rf, ok := ret.Get(0).(func([]byte, string, ...*big.Int) ethkey.KeyV2); ok {
		r0 = rf(keyJSON, password, chainIDs...)
	} else {
		r0 = ret.Get(0).(ethkey.KeyV2)
	}

	if rf, ok := ret.Get(1).(func([]byte, string, ...*big.Int) error); ok {
		r1 = rf(keyJSON, password, chainIDs...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IncrementNextSequence provides a mock function with given fields: address, chainID, currentNonce, qopts
func (_m *Eth) IncrementNextSequence(address common.Address, chainID *big.Int, currentNonce types.Nonce, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, address, chainID, currentNonce)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.Address, *big.Int, types.Nonce, ...pg.QOpt) error); ok {
		r0 = rf(address, chainID, currentNonce, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NextSequence provides a mock function with given fields: address, chainID, qopts
func (_m *Eth) NextSequence(address common.Address, chainID *big.Int, qopts ...pg.QOpt) (types.Nonce, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, address, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 types.Nonce
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address, *big.Int, ...pg.QOpt) (types.Nonce, error)); ok {
		return rf(address, chainID, qopts...)
	}
	if rf, ok := ret.Get(0).(func(common.Address, *big.Int, ...pg.QOpt) types.Nonce); ok {
		r0 = rf(address, chainID, qopts...)
	} else {
		r0 = ret.Get(0).(types.Nonce)
	}

	if rf, ok := ret.Get(1).(func(common.Address, *big.Int, ...pg.QOpt) error); ok {
		r1 = rf(address, chainID, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Reset provides a mock function with given fields: address, chainID, nonce, qopts
func (_m *Eth) Reset(address common.Address, chainID *big.Int, nonce int64, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, address, chainID, nonce)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.Address, *big.Int, int64, ...pg.QOpt) error); ok {
		r0 = rf(address, chainID, nonce, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SignTx provides a mock function with given fields: fromAddress, tx, chainID
func (_m *Eth) SignTx(fromAddress common.Address, tx *coretypes.Transaction, chainID *big.Int) (*coretypes.Transaction, error) {
	ret := _m.Called(fromAddress, tx, chainID)

	var r0 *coretypes.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address, *coretypes.Transaction, *big.Int) (*coretypes.Transaction, error)); ok {
		return rf(fromAddress, tx, chainID)
	}
	if rf, ok := ret.Get(0).(func(common.Address, *coretypes.Transaction, *big.Int) *coretypes.Transaction); ok {
		r0 = rf(fromAddress, tx, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Address, *coretypes.Transaction, *big.Int) error); ok {
		r1 = rf(fromAddress, tx, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribeToKeyChanges provides a mock function with given fields:
func (_m *Eth) SubscribeToKeyChanges() (chan struct{}, func()) {
	ret := _m.Called()

	var r0 chan struct{}
	var r1 func()
	if rf, ok := ret.Get(0).(func() (chan struct{}, func())); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan struct{})
		}
	}

	if rf, ok := ret.Get(1).(func() func()); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(func())
		}
	}

	return r0, r1
}

// XXXTestingOnlyAdd provides a mock function with given fields: key
func (_m *Eth) XXXTestingOnlyAdd(key ethkey.KeyV2) {
	_m.Called(key)
}

// XXXTestingOnlySetState provides a mock function with given fields: _a0
func (_m *Eth) XXXTestingOnlySetState(_a0 ethkey.State) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewEth interface {
	mock.TestingT
	Cleanup(func())
}

// NewEth creates a new instance of Eth. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEth(t mockConstructorTestingTNewEth) *Eth {
	mock := &Eth{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
