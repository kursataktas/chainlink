// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	p2ptypes "github.com/smartcontractkit/chainlink/v2/core/services/p2p/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/smartcontractkit/chainlink/v2/core/capabilities/remote/types"
)

// Dispatcher is an autogenerated mock type for the Dispatcher type
type Dispatcher struct {
	mock.Mock
}

type Dispatcher_Expecter struct {
	mock *mock.Mock
}

func (_m *Dispatcher) EXPECT() *Dispatcher_Expecter {
	return &Dispatcher_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *Dispatcher) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Dispatcher_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type Dispatcher_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *Dispatcher_Expecter) Close() *Dispatcher_Close_Call {
	return &Dispatcher_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *Dispatcher_Close_Call) Run(run func()) *Dispatcher_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Dispatcher_Close_Call) Return(_a0 error) *Dispatcher_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Dispatcher_Close_Call) RunAndReturn(run func() error) *Dispatcher_Close_Call {
	_c.Call.Return(run)
	return _c
}

// HealthReport provides a mock function with given fields:
func (_m *Dispatcher) HealthReport() map[string]error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HealthReport")
	}

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

// Dispatcher_HealthReport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HealthReport'
type Dispatcher_HealthReport_Call struct {
	*mock.Call
}

// HealthReport is a helper method to define mock.On call
func (_e *Dispatcher_Expecter) HealthReport() *Dispatcher_HealthReport_Call {
	return &Dispatcher_HealthReport_Call{Call: _e.mock.On("HealthReport")}
}

func (_c *Dispatcher_HealthReport_Call) Run(run func()) *Dispatcher_HealthReport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Dispatcher_HealthReport_Call) Return(_a0 map[string]error) *Dispatcher_HealthReport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Dispatcher_HealthReport_Call) RunAndReturn(run func() map[string]error) *Dispatcher_HealthReport_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *Dispatcher) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Dispatcher_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type Dispatcher_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *Dispatcher_Expecter) Name() *Dispatcher_Name_Call {
	return &Dispatcher_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *Dispatcher_Name_Call) Run(run func()) *Dispatcher_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Dispatcher_Name_Call) Return(_a0 string) *Dispatcher_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Dispatcher_Name_Call) RunAndReturn(run func() string) *Dispatcher_Name_Call {
	_c.Call.Return(run)
	return _c
}

// Ready provides a mock function with given fields:
func (_m *Dispatcher) Ready() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Ready")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Dispatcher_Ready_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ready'
type Dispatcher_Ready_Call struct {
	*mock.Call
}

// Ready is a helper method to define mock.On call
func (_e *Dispatcher_Expecter) Ready() *Dispatcher_Ready_Call {
	return &Dispatcher_Ready_Call{Call: _e.mock.On("Ready")}
}

func (_c *Dispatcher_Ready_Call) Run(run func()) *Dispatcher_Ready_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Dispatcher_Ready_Call) Return(_a0 error) *Dispatcher_Ready_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Dispatcher_Ready_Call) RunAndReturn(run func() error) *Dispatcher_Ready_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveReceiver provides a mock function with given fields: capabilityId, donId
func (_m *Dispatcher) RemoveReceiver(capabilityId string, donId uint32) {
	_m.Called(capabilityId, donId)
}

// Dispatcher_RemoveReceiver_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveReceiver'
type Dispatcher_RemoveReceiver_Call struct {
	*mock.Call
}

// RemoveReceiver is a helper method to define mock.On call
//   - capabilityId string
//   - donId uint32
func (_e *Dispatcher_Expecter) RemoveReceiver(capabilityId interface{}, donId interface{}) *Dispatcher_RemoveReceiver_Call {
	return &Dispatcher_RemoveReceiver_Call{Call: _e.mock.On("RemoveReceiver", capabilityId, donId)}
}

func (_c *Dispatcher_RemoveReceiver_Call) Run(run func(capabilityId string, donId uint32)) *Dispatcher_RemoveReceiver_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(uint32))
	})
	return _c
}

func (_c *Dispatcher_RemoveReceiver_Call) Return() *Dispatcher_RemoveReceiver_Call {
	_c.Call.Return()
	return _c
}

func (_c *Dispatcher_RemoveReceiver_Call) RunAndReturn(run func(string, uint32)) *Dispatcher_RemoveReceiver_Call {
	_c.Call.Return(run)
	return _c
}

// Send provides a mock function with given fields: peerID, msgBody
func (_m *Dispatcher) Send(peerID p2ptypes.PeerID, msgBody *types.MessageBody) error {
	ret := _m.Called(peerID, msgBody)

	if len(ret) == 0 {
		panic("no return value specified for Send")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(p2ptypes.PeerID, *types.MessageBody) error); ok {
		r0 = rf(peerID, msgBody)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Dispatcher_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type Dispatcher_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - peerID p2ptypes.PeerID
//   - msgBody *types.MessageBody
func (_e *Dispatcher_Expecter) Send(peerID interface{}, msgBody interface{}) *Dispatcher_Send_Call {
	return &Dispatcher_Send_Call{Call: _e.mock.On("Send", peerID, msgBody)}
}

func (_c *Dispatcher_Send_Call) Run(run func(peerID p2ptypes.PeerID, msgBody *types.MessageBody)) *Dispatcher_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(p2ptypes.PeerID), args[1].(*types.MessageBody))
	})
	return _c
}

func (_c *Dispatcher_Send_Call) Return(_a0 error) *Dispatcher_Send_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Dispatcher_Send_Call) RunAndReturn(run func(p2ptypes.PeerID, *types.MessageBody) error) *Dispatcher_Send_Call {
	_c.Call.Return(run)
	return _c
}

// SetReceiver provides a mock function with given fields: capabilityId, donId, receiver
func (_m *Dispatcher) SetReceiver(capabilityId string, donId uint32, receiver types.Receiver) error {
	ret := _m.Called(capabilityId, donId, receiver)

	if len(ret) == 0 {
		panic("no return value specified for SetReceiver")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint32, types.Receiver) error); ok {
		r0 = rf(capabilityId, donId, receiver)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Dispatcher_SetReceiver_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetReceiver'
type Dispatcher_SetReceiver_Call struct {
	*mock.Call
}

// SetReceiver is a helper method to define mock.On call
//   - capabilityId string
//   - donId uint32
//   - receiver types.Receiver
func (_e *Dispatcher_Expecter) SetReceiver(capabilityId interface{}, donId interface{}, receiver interface{}) *Dispatcher_SetReceiver_Call {
	return &Dispatcher_SetReceiver_Call{Call: _e.mock.On("SetReceiver", capabilityId, donId, receiver)}
}

func (_c *Dispatcher_SetReceiver_Call) Run(run func(capabilityId string, donId uint32, receiver types.Receiver)) *Dispatcher_SetReceiver_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(uint32), args[2].(types.Receiver))
	})
	return _c
}

func (_c *Dispatcher_SetReceiver_Call) Return(_a0 error) *Dispatcher_SetReceiver_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Dispatcher_SetReceiver_Call) RunAndReturn(run func(string, uint32, types.Receiver) error) *Dispatcher_SetReceiver_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: _a0
func (_m *Dispatcher) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Dispatcher_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type Dispatcher_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *Dispatcher_Expecter) Start(_a0 interface{}) *Dispatcher_Start_Call {
	return &Dispatcher_Start_Call{Call: _e.mock.On("Start", _a0)}
}

func (_c *Dispatcher_Start_Call) Run(run func(_a0 context.Context)) *Dispatcher_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Dispatcher_Start_Call) Return(_a0 error) *Dispatcher_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Dispatcher_Start_Call) RunAndReturn(run func(context.Context) error) *Dispatcher_Start_Call {
	_c.Call.Return(run)
	return _c
}

// NewDispatcher creates a new instance of Dispatcher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDispatcher(t interface {
	mock.TestingT
	Cleanup(func())
}) *Dispatcher {
	mock := &Dispatcher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
