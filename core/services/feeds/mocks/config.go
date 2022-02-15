// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	time "time"

	models "github.com/smartcontractkit/chainlink/core/store/models"
	mock "github.com/stretchr/testify/mock"
)

// Config is an autogenerated mock type for the Config type
type Config struct {
	mock.Mock
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

// Dev provides a mock function with given fields:
func (_m *Config) Dev() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// FeatureOffchainReporting provides a mock function with given fields:
func (_m *Config) FeatureOffchainReporting() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// LogSQL provides a mock function with given fields:
func (_m *Config) LogSQL() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// OCRBlockchainTimeout provides a mock function with given fields:
func (_m *Config) OCRBlockchainTimeout() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// OCRContractConfirmations provides a mock function with given fields:
func (_m *Config) OCRContractConfirmations() uint16 {
	ret := _m.Called()

	var r0 uint16
	if rf, ok := ret.Get(0).(func() uint16); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// OCRContractPollInterval provides a mock function with given fields:
func (_m *Config) OCRContractPollInterval() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// OCRContractSubscribeInterval provides a mock function with given fields:
func (_m *Config) OCRContractSubscribeInterval() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// OCRContractTransmitterTransmitTimeout provides a mock function with given fields:
func (_m *Config) OCRContractTransmitterTransmitTimeout() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// OCRDatabaseTimeout provides a mock function with given fields:
func (_m *Config) OCRDatabaseTimeout() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// OCRObservationGracePeriod provides a mock function with given fields:
func (_m *Config) OCRObservationGracePeriod() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// OCRObservationTimeout provides a mock function with given fields:
func (_m *Config) OCRObservationTimeout() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}
