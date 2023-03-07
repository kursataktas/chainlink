// Code generated by mockery v2.21.2. DO NOT EDIT.

package mocks

import (
	context "context"

	synchronization "github.com/smartcontractkit/chainlink/core/services/synchronization"
	mock "github.com/stretchr/testify/mock"
)

// TelemetryIngressBatchClient is an autogenerated mock type for the TelemetryIngressBatchClient type
type TelemetryIngressBatchClient struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *TelemetryIngressBatchClient) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HealthReport provides a mock function with given fields:
func (_m *TelemetryIngressBatchClient) HealthReport() map[string]error {
	ret := _m.Called()

	var r0 map[string]error
	if rf, ok := ret.Get(0).(func() map[string]error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]error)
		}
	}

	return r0
}

// Healthy provides a mock function with given fields:
func (_m *TelemetryIngressBatchClient) Healthy() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *TelemetryIngressBatchClient) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Ready provides a mock function with given fields:
func (_m *TelemetryIngressBatchClient) Ready() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Send provides a mock function with given fields: _a0
func (_m *TelemetryIngressBatchClient) Send(_a0 synchronization.TelemPayload) {
	_m.Called(_a0)
}

// Start provides a mock function with given fields: _a0
func (_m *TelemetryIngressBatchClient) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTelemetryIngressBatchClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewTelemetryIngressBatchClient creates a new instance of TelemetryIngressBatchClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTelemetryIngressBatchClient(t mockConstructorTestingTNewTelemetryIngressBatchClient) *TelemetryIngressBatchClient {
	mock := &TelemetryIngressBatchClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
