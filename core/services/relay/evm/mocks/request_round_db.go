// Code generated by mockery v2.21.2. DO NOT EDIT.

package mocks

import (
	pg "github.com/smartcontractkit/chainlink/core/services/pg"
	ocr2aggregator "github.com/smartcontractkit/libocr/gethwrappers2/ocr2aggregator"
	mock "github.com/stretchr/testify/mock"
)

// RequestRoundDB is an autogenerated mock type for the RequestRoundDB type
type RequestRoundDB struct {
	mock.Mock
}

// LoadLatestRoundRequested provides a mock function with given fields:
func (_m *RequestRoundDB) LoadLatestRoundRequested() (ocr2aggregator.OCR2AggregatorRoundRequested, error) {
	ret := _m.Called()

	var r0 ocr2aggregator.OCR2AggregatorRoundRequested
	var r1 error
	if rf, ok := ret.Get(0).(func() (ocr2aggregator.OCR2AggregatorRoundRequested, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() ocr2aggregator.OCR2AggregatorRoundRequested); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(ocr2aggregator.OCR2AggregatorRoundRequested)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveLatestRoundRequested provides a mock function with given fields: tx, rr
func (_m *RequestRoundDB) SaveLatestRoundRequested(tx pg.Queryer, rr ocr2aggregator.OCR2AggregatorRoundRequested) error {
	ret := _m.Called(tx, rr)

	var r0 error
	if rf, ok := ret.Get(0).(func(pg.Queryer, ocr2aggregator.OCR2AggregatorRoundRequested) error); ok {
		r0 = rf(tx, rr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRequestRoundDB interface {
	mock.TestingT
	Cleanup(func())
}

// NewRequestRoundDB creates a new instance of RequestRoundDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRequestRoundDB(t mockConstructorTestingTNewRequestRoundDB) *RequestRoundDB {
	mock := &RequestRoundDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
