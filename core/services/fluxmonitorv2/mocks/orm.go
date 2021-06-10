// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	common "github.com/ethereum/go-ethereum/common"
	fluxmonitorv2 "github.com/smartcontractkit/chainlink/core/services/fluxmonitorv2"
	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

// CreateEthTransaction provides a mock function with given fields: db, fromAddress, toAddress, payload, gasLimit, maxUnconfirmedTransactions
func (_m *ORM) CreateEthTransaction(db *gorm.DB, fromAddress common.Address, toAddress common.Address, payload []byte, gasLimit uint64, maxUnconfirmedTransactions uint64) error {
	ret := _m.Called(db, fromAddress, toAddress, payload, gasLimit, maxUnconfirmedTransactions)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, common.Address, common.Address, []byte, uint64, uint64) error); ok {
		r0 = rf(db, fromAddress, toAddress, payload, gasLimit, maxUnconfirmedTransactions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteFluxMonitorRoundsBackThrough provides a mock function with given fields: aggregator, roundID
func (_m *ORM) DeleteFluxMonitorRoundsBackThrough(aggregator common.Address, roundID uint32) error {
	ret := _m.Called(aggregator, roundID)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.Address, uint32) error); ok {
		r0 = rf(aggregator, roundID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindOrCreateFluxMonitorRoundStats provides a mock function with given fields: aggregator, roundID
func (_m *ORM) FindOrCreateFluxMonitorRoundStats(aggregator common.Address, roundID uint32) (fluxmonitorv2.FluxMonitorRoundStatsV2, error) {
	ret := _m.Called(aggregator, roundID)

	var r0 fluxmonitorv2.FluxMonitorRoundStatsV2
	if rf, ok := ret.Get(0).(func(common.Address, uint32) fluxmonitorv2.FluxMonitorRoundStatsV2); ok {
		r0 = rf(aggregator, roundID)
	} else {
		r0 = ret.Get(0).(fluxmonitorv2.FluxMonitorRoundStatsV2)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, uint32) error); ok {
		r1 = rf(aggregator, roundID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MostRecentFluxMonitorRoundID provides a mock function with given fields: aggregator
func (_m *ORM) MostRecentFluxMonitorRoundID(aggregator common.Address) (uint32, error) {
	ret := _m.Called(aggregator)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(common.Address) uint32); ok {
		r0 = rf(aggregator)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(aggregator)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateFluxMonitorRoundStats provides a mock function with given fields: db, aggregator, roundID, runID
func (_m *ORM) UpdateFluxMonitorRoundStats(db *gorm.DB, aggregator common.Address, roundID uint32, runID int64) error {
	ret := _m.Called(db, aggregator, roundID, runID)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, common.Address, uint32, int64) error); ok {
		r0 = rf(db, aggregator, roundID, runID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
