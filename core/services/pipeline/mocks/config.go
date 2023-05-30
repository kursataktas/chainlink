// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	models "github.com/smartcontractkit/chainlink/v2/core/store/models"
	mock "github.com/stretchr/testify/mock"

	time "time"

	url "net/url"
)

// Config is an autogenerated mock type for the Config type
type Config struct {
	mock.Mock
}

// BridgeCacheTTL provides a mock function with given fields:
func (_m *Config) BridgeCacheTTL() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// BridgeResponseURL provides a mock function with given fields:
func (_m *Config) BridgeResponseURL() *url.URL {
	ret := _m.Called()

	var r0 *url.URL
	if rf, ok := ret.Get(0).(func() *url.URL); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*url.URL)
		}
	}

	return r0
}

// DatabaseURL provides a mock function with given fields:
func (_m *Config) DatabaseURL() url.URL {
	ret := _m.Called()

	var r0 url.URL
	if rf, ok := ret.Get(0).(func() url.URL); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(url.URL)
	}

	return r0
}

// DefaultHTTPLimit provides a mock function with given fields:
func (_m *Config) DefaultHTTPLimit() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// DefaultHTTPTimeout provides a mock function with given fields:
func (_m *Config) DefaultHTTPTimeout() models.Duration {
	ret := _m.Called()

	var r0 models.Duration
	if rf, ok := ret.Get(0).(func() models.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.Duration)
	}

	return r0
}

// JobPipelineMaxRunDuration provides a mock function with given fields:
func (_m *Config) JobPipelineMaxRunDuration() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// JobPipelineReaperInterval provides a mock function with given fields:
func (_m *Config) JobPipelineReaperInterval() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// JobPipelineReaperThreshold provides a mock function with given fields:
func (_m *Config) JobPipelineReaperThreshold() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

type mockConstructorTestingTNewConfig interface {
	mock.TestingT
	Cleanup(func())
}

// NewConfig creates a new instance of Config. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConfig(t mockConstructorTestingTNewConfig) *Config {
	mock := &Config{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
