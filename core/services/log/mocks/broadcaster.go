// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"

	gorm "gorm.io/gorm"

	log "github.com/smartcontractkit/chainlink/core/services/log"
	mock "github.com/stretchr/testify/mock"

	models "github.com/smartcontractkit/chainlink/core/store/models"

	null "github.com/smartcontractkit/chainlink/core/null"
)

// Broadcaster is an autogenerated mock type for the Broadcaster type
type Broadcaster struct {
	mock.Mock
}

// AddDependents provides a mock function with given fields: n
func (_m *Broadcaster) AddDependents(n int) {
	_m.Called(n)
}

// AwaitDependents provides a mock function with given fields:
func (_m *Broadcaster) AwaitDependents() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *Broadcaster) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Connect provides a mock function with given fields: head
func (_m *Broadcaster) Connect(head *models.Head) error {
	ret := _m.Called(head)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Head) error); ok {
		r0 = rf(head)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DependentReady provides a mock function with given fields:
func (_m *Broadcaster) DependentReady() {
	_m.Called()
}

// Healthy provides a mock function with given fields:
func (_m *Broadcaster) Healthy() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsConnected provides a mock function with given fields:
func (_m *Broadcaster) IsConnected() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// LatestHeadNumber provides a mock function with given fields:
func (_m *Broadcaster) LatestHeadNumber() null.Int64 {
	ret := _m.Called()

	var r0 null.Int64
	if rf, ok := ret.Get(0).(func() null.Int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(null.Int64)
	}

	return r0
}

// MarkConsumed provides a mock function with given fields: db, lb
func (_m *Broadcaster) MarkConsumed(db *gorm.DB, lb log.Broadcast) error {
	ret := _m.Called(db, lb)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, log.Broadcast) error); ok {
		r0 = rf(db, lb)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnNewLongestChain provides a mock function with given fields: ctx, head
func (_m *Broadcaster) OnNewLongestChain(ctx context.Context, head models.Head) {
	_m.Called(ctx, head)
}

// Ready provides a mock function with given fields:
func (_m *Broadcaster) Ready() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Register provides a mock function with given fields: listener, opts
func (_m *Broadcaster) Register(listener log.Listener, opts log.ListenerOpts) func() {
	ret := _m.Called(listener, opts)

	var r0 func()
	if rf, ok := ret.Get(0).(func(log.Listener, log.ListenerOpts) func()); ok {
		r0 = rf(listener, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func())
		}
	}

	return r0
}

// Restart provides a mock function with given fields: number
func (_m *Broadcaster) Restart(number int64) {
	_m.Called(number)
}

// Start provides a mock function with given fields:
func (_m *Broadcaster) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TrackedAddressesCount provides a mock function with given fields:
func (_m *Broadcaster) TrackedAddressesCount() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// WasAlreadyConsumed provides a mock function with given fields: db, lb
func (_m *Broadcaster) WasAlreadyConsumed(db *gorm.DB, lb log.Broadcast) (bool, error) {
	ret := _m.Called(db, lb)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*gorm.DB, log.Broadcast) bool); ok {
		r0 = rf(db, lb)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, log.Broadcast) error); ok {
		r1 = rf(db, lb)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
