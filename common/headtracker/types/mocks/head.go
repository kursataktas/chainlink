// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	txmgrtypes "github.com/smartcontractkit/chainlink/v2/common/txmgr/types"
	types "github.com/smartcontractkit/chainlink/v2/common/types"
	mock "github.com/stretchr/testify/mock"
)

// Head is an autogenerated mock type for the Head type
type Head[H types.Head[BLOCK_HASH], BLOCK_HASH types.Hashable, CHAIN_ID txmgrtypes.ID] struct {
	mock.Mock
}

// BlockHash provides a mock function with given fields:
func (_m *Head[H, BLOCK_HASH, CHAIN_ID]) BlockHash() BLOCK_HASH {
	ret := _m.Called()

	var r0 BLOCK_HASH
	if rf, ok := ret.Get(0).(func() BLOCK_HASH); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(BLOCK_HASH)
	}

	return r0
}

// BlockNumber provides a mock function with given fields:
func (_m *Head[H, BLOCK_HASH, CHAIN_ID]) BlockNumber() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// ChainId provides a mock function with given fields:
func (_m *Head[H, BLOCK_HASH, CHAIN_ID]) ChainId() CHAIN_ID {
	ret := _m.Called()

	var r0 CHAIN_ID
	if rf, ok := ret.Get(0).(func() CHAIN_ID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(CHAIN_ID)
	}

	return r0
}

// ChainLength provides a mock function with given fields:
func (_m *Head[H, BLOCK_HASH, CHAIN_ID]) ChainLength() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// EarliestHeadInChain provides a mock function with given fields:
func (_m *Head[H, BLOCK_HASH, CHAIN_ID]) EarliestHeadInChain() types.Head[BLOCK_HASH] {
	ret := _m.Called()

	var r0 types.Head[BLOCK_HASH]
	if rf, ok := ret.Get(0).(func() types.Head[BLOCK_HASH]); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Head[BLOCK_HASH])
		}
	}

	return r0
}

// Equals provides a mock function with given fields: _a0
func (_m *Head[H, BLOCK_HASH, CHAIN_ID]) Equals(_a0 H) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(H) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetParent provides a mock function with given fields:
func (_m *Head[H, BLOCK_HASH, CHAIN_ID]) GetParent() types.Head[BLOCK_HASH] {
	ret := _m.Called()

	var r0 types.Head[BLOCK_HASH]
	if rf, ok := ret.Get(0).(func() types.Head[BLOCK_HASH]); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Head[BLOCK_HASH])
		}
	}

	return r0
}

// HasChainID provides a mock function with given fields:
func (_m *Head[H, BLOCK_HASH, CHAIN_ID]) HasChainID() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// HashAtHeight provides a mock function with given fields: blockNum
func (_m *Head[H, BLOCK_HASH, CHAIN_ID]) HashAtHeight(blockNum int64) BLOCK_HASH {
	ret := _m.Called(blockNum)

	var r0 BLOCK_HASH
	if rf, ok := ret.Get(0).(func(int64) BLOCK_HASH); ok {
		r0 = rf(blockNum)
	} else {
		r0 = ret.Get(0).(BLOCK_HASH)
	}

	return r0
}

// IsSameChain provides a mock function with given fields: _a0
func (_m *Head[H, BLOCK_HASH, CHAIN_ID]) IsSameChain(_a0 CHAIN_ID) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(CHAIN_ID) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewHead interface {
	mock.TestingT
	Cleanup(func())
}

// NewHead creates a new instance of Head. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHead[H types.Head[BLOCK_HASH], BLOCK_HASH types.Hashable, CHAIN_ID txmgrtypes.ID](t mockConstructorTestingTNewHead) *Head[H, BLOCK_HASH, CHAIN_ID] {
	mock := &Head[H, BLOCK_HASH, CHAIN_ID]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
