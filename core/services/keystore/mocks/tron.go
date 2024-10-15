// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	tronkey "github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/tronkey"
)

// Tron is an autogenerated mock type for the Tron type
type Tron struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, key
func (_m *Tron) Add(ctx context.Context, key tronkey.Key) error {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, tronkey.Key) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: ctx
func (_m *Tron) Create(ctx context.Context) (tronkey.Key, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 tronkey.Key
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (tronkey.Key, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) tronkey.Key); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(tronkey.Key)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Tron) Delete(ctx context.Context, id string) (tronkey.Key, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 tronkey.Key
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (tronkey.Key, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) tronkey.Key); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(tronkey.Key)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnsureKey provides a mock function with given fields: ctx
func (_m *Tron) EnsureKey(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for EnsureKey")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Export provides a mock function with given fields: id, password
func (_m *Tron) Export(id string, password string) ([]byte, error) {
	ret := _m.Called(id, password)

	if len(ret) == 0 {
		panic("no return value specified for Export")
	}

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
func (_m *Tron) Get(id string) (tronkey.Key, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 tronkey.Key
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (tronkey.Key, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) tronkey.Key); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(tronkey.Key)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *Tron) GetAll() ([]tronkey.Key, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []tronkey.Key
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]tronkey.Key, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []tronkey.Key); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]tronkey.Key)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Import provides a mock function with given fields: ctx, keyJSON, password
func (_m *Tron) Import(ctx context.Context, keyJSON []byte, password string) (tronkey.Key, error) {
	ret := _m.Called(ctx, keyJSON, password)

	if len(ret) == 0 {
		panic("no return value specified for Import")
	}

	var r0 tronkey.Key
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, string) (tronkey.Key, error)); ok {
		return rf(ctx, keyJSON, password)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte, string) tronkey.Key); ok {
		r0 = rf(ctx, keyJSON, password)
	} else {
		r0 = ret.Get(0).(tronkey.Key)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte, string) error); ok {
		r1 = rf(ctx, keyJSON, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Sign provides a mock function with given fields: ctx, id, msg
func (_m *Tron) Sign(ctx context.Context, id string, msg []byte) ([]byte, error) {
	ret := _m.Called(ctx, id, msg)

	if len(ret) == 0 {
		panic("no return value specified for Sign")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte) ([]byte, error)); ok {
		return rf(ctx, id, msg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte) []byte); ok {
		r0 = rf(ctx, id, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []byte) error); ok {
		r1 = rf(ctx, id, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTron creates a new instance of Tron. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTron(t interface {
	mock.TestingT
	Cleanup(func())
}) *Tron {
	mock := &Tron{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
