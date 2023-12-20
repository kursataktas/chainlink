// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	common "github.com/ethereum/go-ethereum/common"
	ccipdata "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"

	context "context"

	evm_2_evm_offramp "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp"

	mock "github.com/stretchr/testify/mock"

	pg "github.com/smartcontractkit/chainlink/v2/core/services/pg"

	prices "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/prices"
)

// OffRampReader is an autogenerated mock type for the OffRampReader type
type OffRampReader struct {
	mock.Mock
}

// Address provides a mock function with given fields:
func (_m *OffRampReader) Address() common.Address {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Address")
	}

	var r0 common.Address
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	return r0
}

// ChangeConfig provides a mock function with given fields: onchainConfig, offchainConfig
func (_m *OffRampReader) ChangeConfig(onchainConfig []byte, offchainConfig []byte) (common.Address, common.Address, error) {
	ret := _m.Called(onchainConfig, offchainConfig)

	if len(ret) == 0 {
		panic("no return value specified for ChangeConfig")
	}

	var r0 common.Address
	var r1 common.Address
	var r2 error
	if rf, ok := ret.Get(0).(func([]byte, []byte) (common.Address, common.Address, error)); ok {
		return rf(onchainConfig, offchainConfig)
	}
	if rf, ok := ret.Get(0).(func([]byte, []byte) common.Address); ok {
		r0 = rf(onchainConfig, offchainConfig)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte, []byte) common.Address); ok {
		r1 = rf(onchainConfig, offchainConfig)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(common.Address)
		}
	}

	if rf, ok := ret.Get(2).(func([]byte, []byte) error); ok {
		r2 = rf(onchainConfig, offchainConfig)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Close provides a mock function with given fields: qopts
func (_m *OffRampReader) Close(qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...pg.QOpt) error); ok {
		r0 = rf(qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CurrentRateLimiterState provides a mock function with given fields: ctx
func (_m *OffRampReader) CurrentRateLimiterState(ctx context.Context) (evm_2_evm_offramp.RateLimiterTokenBucket, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CurrentRateLimiterState")
	}

	var r0 evm_2_evm_offramp.RateLimiterTokenBucket
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (evm_2_evm_offramp.RateLimiterTokenBucket, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) evm_2_evm_offramp.RateLimiterTokenBucket); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(evm_2_evm_offramp.RateLimiterTokenBucket)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecodeExecutionReport provides a mock function with given fields: report
func (_m *OffRampReader) DecodeExecutionReport(report []byte) (ccipdata.ExecReport, error) {
	ret := _m.Called(report)

	if len(ret) == 0 {
		panic("no return value specified for DecodeExecutionReport")
	}

	var r0 ccipdata.ExecReport
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (ccipdata.ExecReport, error)); ok {
		return rf(report)
	}
	if rf, ok := ret.Get(0).(func([]byte) ccipdata.ExecReport); ok {
		r0 = rf(report)
	} else {
		r0 = ret.Get(0).(ccipdata.ExecReport)
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(report)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EncodeExecutionReport provides a mock function with given fields: report
func (_m *OffRampReader) EncodeExecutionReport(report ccipdata.ExecReport) ([]byte, error) {
	ret := _m.Called(report)

	if len(ret) == 0 {
		panic("no return value specified for EncodeExecutionReport")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(ccipdata.ExecReport) ([]byte, error)); ok {
		return rf(report)
	}
	if rf, ok := ret.Get(0).(func(ccipdata.ExecReport) []byte); ok {
		r0 = rf(report)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(ccipdata.ExecReport) error); ok {
		r1 = rf(report)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GasPriceEstimator provides a mock function with given fields:
func (_m *OffRampReader) GasPriceEstimator() prices.GasPriceEstimatorExec {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GasPriceEstimator")
	}

	var r0 prices.GasPriceEstimatorExec
	if rf, ok := ret.Get(0).(func() prices.GasPriceEstimatorExec); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(prices.GasPriceEstimatorExec)
		}
	}

	return r0
}

// GetExecutionState provides a mock function with given fields: ctx, sequenceNumber
func (_m *OffRampReader) GetExecutionState(ctx context.Context, sequenceNumber uint64) (uint8, error) {
	ret := _m.Called(ctx, sequenceNumber)

	if len(ret) == 0 {
		panic("no return value specified for GetExecutionState")
	}

	var r0 uint8
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (uint8, error)); ok {
		return rf(ctx, sequenceNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) uint8); ok {
		r0 = rf(ctx, sequenceNumber)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, sequenceNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExecutionStateChangesBetweenSeqNums provides a mock function with given fields: ctx, seqNumMin, seqNumMax, confs
func (_m *OffRampReader) GetExecutionStateChangesBetweenSeqNums(ctx context.Context, seqNumMin uint64, seqNumMax uint64, confs int) ([]ccipdata.Event[ccipdata.ExecutionStateChanged], error) {
	ret := _m.Called(ctx, seqNumMin, seqNumMax, confs)

	if len(ret) == 0 {
		panic("no return value specified for GetExecutionStateChangesBetweenSeqNums")
	}

	var r0 []ccipdata.Event[ccipdata.ExecutionStateChanged]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, int) ([]ccipdata.Event[ccipdata.ExecutionStateChanged], error)); ok {
		return rf(ctx, seqNumMin, seqNumMax, confs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, int) []ccipdata.Event[ccipdata.ExecutionStateChanged]); ok {
		r0 = rf(ctx, seqNumMin, seqNumMax, confs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ccipdata.Event[ccipdata.ExecutionStateChanged])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, int) error); ok {
		r1 = rf(ctx, seqNumMin, seqNumMax, confs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSenderNonce provides a mock function with given fields: ctx, sender
func (_m *OffRampReader) GetSenderNonce(ctx context.Context, sender common.Address) (uint64, error) {
	ret := _m.Called(ctx, sender)

	if len(ret) == 0 {
		panic("no return value specified for GetSenderNonce")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) (uint64, error)); ok {
		return rf(ctx, sender)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) uint64); ok {
		r0 = rf(ctx, sender)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address) error); ok {
		r1 = rf(ctx, sender)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSourceToDestTokensMapping provides a mock function with given fields: ctx
func (_m *OffRampReader) GetSourceToDestTokensMapping(ctx context.Context) (map[common.Address]common.Address, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetSourceToDestTokensMapping")
	}

	var r0 map[common.Address]common.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (map[common.Address]common.Address, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) map[common.Address]common.Address); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[common.Address]common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStaticConfig provides a mock function with given fields: ctx
func (_m *OffRampReader) GetStaticConfig(ctx context.Context) (ccipdata.OffRampStaticConfig, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetStaticConfig")
	}

	var r0 ccipdata.OffRampStaticConfig
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (ccipdata.OffRampStaticConfig, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) ccipdata.OffRampStaticConfig); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(ccipdata.OffRampStaticConfig)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTokenPoolsRateLimits provides a mock function with given fields: ctx, poolAddresses
func (_m *OffRampReader) GetTokenPoolsRateLimits(ctx context.Context, poolAddresses []common.Address) ([]ccipdata.TokenBucketRateLimit, error) {
	ret := _m.Called(ctx, poolAddresses)

	if len(ret) == 0 {
		panic("no return value specified for GetTokenPoolsRateLimits")
	}

	var r0 []ccipdata.TokenBucketRateLimit
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []common.Address) ([]ccipdata.TokenBucketRateLimit, error)); ok {
		return rf(ctx, poolAddresses)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []common.Address) []ccipdata.TokenBucketRateLimit); ok {
		r0 = rf(ctx, poolAddresses)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ccipdata.TokenBucketRateLimit)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []common.Address) error); ok {
		r1 = rf(ctx, poolAddresses)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTokens provides a mock function with given fields: ctx
func (_m *OffRampReader) GetTokens(ctx context.Context) (ccipdata.OffRampTokens, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetTokens")
	}

	var r0 ccipdata.OffRampTokens
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (ccipdata.OffRampTokens, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) ccipdata.OffRampTokens); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(ccipdata.OffRampTokens)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OffchainConfig provides a mock function with given fields:
func (_m *OffRampReader) OffchainConfig() ccipdata.ExecOffchainConfig {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for OffchainConfig")
	}

	var r0 ccipdata.ExecOffchainConfig
	if rf, ok := ret.Get(0).(func() ccipdata.ExecOffchainConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(ccipdata.ExecOffchainConfig)
	}

	return r0
}

// OnchainConfig provides a mock function with given fields:
func (_m *OffRampReader) OnchainConfig() ccipdata.ExecOnchainConfig {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for OnchainConfig")
	}

	var r0 ccipdata.ExecOnchainConfig
	if rf, ok := ret.Get(0).(func() ccipdata.ExecOnchainConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(ccipdata.ExecOnchainConfig)
	}

	return r0
}

// RegisterFilters provides a mock function with given fields: qopts
func (_m *OffRampReader) RegisterFilters(qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RegisterFilters")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...pg.QOpt) error); ok {
		r0 = rf(qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewOffRampReader creates a new instance of OffRampReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOffRampReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *OffRampReader {
	mock := &OffRampReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
