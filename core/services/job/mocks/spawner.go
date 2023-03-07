// Code generated by mockery v2.21.2. DO NOT EDIT.

package mocks

import (
	context "context"

	job "github.com/smartcontractkit/chainlink/core/services/job"
	mock "github.com/stretchr/testify/mock"

	pg "github.com/smartcontractkit/chainlink/core/services/pg"
)

// Spawner is an autogenerated mock type for the Spawner type
type Spawner struct {
	mock.Mock
}

// ActiveJobs provides a mock function with given fields:
func (_m *Spawner) ActiveJobs() map[int32]job.Job {
	ret := _m.Called()

	var r0 map[int32]job.Job
	if rf, ok := ret.Get(0).(func() map[int32]job.Job); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[int32]job.Job)
		}
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *Spawner) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateJob provides a mock function with given fields: jb, qopts
func (_m *Spawner) CreateJob(jb *job.Job, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, jb)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*job.Job, ...pg.QOpt) error); ok {
		r0 = rf(jb, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteJob provides a mock function with given fields: jobID, qopts
func (_m *Spawner) DeleteJob(jobID int32, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, jobID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(int32, ...pg.QOpt) error); ok {
		r0 = rf(jobID, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HealthReport provides a mock function with given fields:
func (_m *Spawner) HealthReport() map[string]error {
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
func (_m *Spawner) Healthy() error {
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
func (_m *Spawner) Name() string {
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
func (_m *Spawner) Ready() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *Spawner) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartService provides a mock function with given fields: ctx, spec
func (_m *Spawner) StartService(ctx context.Context, spec job.Job) error {
	ret := _m.Called(ctx, spec)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, job.Job) error); ok {
		r0 = rf(ctx, spec)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewSpawner interface {
	mock.TestingT
	Cleanup(func())
}

// NewSpawner creates a new instance of Spawner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSpawner(t mockConstructorTestingTNewSpawner) *Spawner {
	mock := &Spawner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
