// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	statuschecker "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/statuschecker"
	mock "github.com/stretchr/testify/mock"
)

// TransactionStatusChecker is an autogenerated mock type for the TransactionStatusChecker type
type TransactionStatusChecker struct {
	mock.Mock
}

// QueryTxStatuses provides a mock function with given fields: ctx, msgID
func (_m *TransactionStatusChecker) CheckMessageStatus(ctx context.Context, msgID string) ([]statuschecker.TransactionStatus, int, error) {
	ret := _m.Called(ctx, msgID)

	if len(ret) == 0 {
		panic("no return value specified for QueryTxStatuses")
	}

	var r0 []statuschecker.TransactionStatus
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]statuschecker.TransactionStatus, int, error)); ok {
		return rf(ctx, msgID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []statuschecker.TransactionStatus); ok {
		r0 = rf(ctx, msgID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]statuschecker.TransactionStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) int); ok {
		r1 = rf(ctx, msgID)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, msgID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewTransactionStatusChecker creates a new instance of TransactionStatusChecker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTransactionStatusChecker(t interface {
	mock.TestingT
	Cleanup(func())
}) *TransactionStatusChecker {
	mock := &TransactionStatusChecker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
