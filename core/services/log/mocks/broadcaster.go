// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import (
	log "github.com/smartcontractkit/chainlink/core/services/log"
	mock "github.com/stretchr/testify/mock"
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

// DependentReady provides a mock function with given fields:
func (_m *Broadcaster) DependentReady() {
	_m.Called()
}

// Register provides a mock function with given fields: listener, opts
func (_m *Broadcaster) Register(listener log.Listener, opts log.ListenerOpts) (bool, func()) {
	ret := _m.Called(listener, opts)

	var r0 bool
	if rf, ok := ret.Get(0).(func(log.Listener, log.ListenerOpts) bool); ok {
		r0 = rf(listener, opts)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 func()
	if rf, ok := ret.Get(1).(func(log.Listener, log.ListenerOpts) func()); ok {
		r1 = rf(listener, opts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(func())
		}
	}

	return r0, r1
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

// Stop provides a mock function with given fields:
func (_m *Broadcaster) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
