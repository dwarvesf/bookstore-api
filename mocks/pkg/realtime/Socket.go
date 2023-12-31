// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Socket is an autogenerated mock type for the Socket type
type Socket struct {
	mock.Mock
}

type Socket_Expecter struct {
	mock *mock.Mock
}

func (_m *Socket) EXPECT() *Socket_Expecter {
	return &Socket_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *Socket) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Socket_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type Socket_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *Socket_Expecter) Close() *Socket_Close_Call {
	return &Socket_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *Socket_Close_Call) Run(run func()) *Socket_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Socket_Close_Call) Return(_a0 error) *Socket_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Socket_Close_Call) RunAndReturn(run func() error) *Socket_Close_Call {
	_c.Call.Return(run)
	return _c
}

// ReadMessage provides a mock function with given fields:
func (_m *Socket) ReadMessage() (int, []byte, error) {
	ret := _m.Called()

	var r0 int
	var r1 []byte
	var r2 error
	if rf, ok := ret.Get(0).(func() (int, []byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func() []byte); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Socket_ReadMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadMessage'
type Socket_ReadMessage_Call struct {
	*mock.Call
}

// ReadMessage is a helper method to define mock.On call
func (_e *Socket_Expecter) ReadMessage() *Socket_ReadMessage_Call {
	return &Socket_ReadMessage_Call{Call: _e.mock.On("ReadMessage")}
}

func (_c *Socket_ReadMessage_Call) Run(run func()) *Socket_ReadMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Socket_ReadMessage_Call) Return(messageType int, p []byte, err error) *Socket_ReadMessage_Call {
	_c.Call.Return(messageType, p, err)
	return _c
}

func (_c *Socket_ReadMessage_Call) RunAndReturn(run func() (int, []byte, error)) *Socket_ReadMessage_Call {
	_c.Call.Return(run)
	return _c
}

// WriteJSON provides a mock function with given fields: v
func (_m *Socket) WriteJSON(v interface{}) error {
	ret := _m.Called(v)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Socket_WriteJSON_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteJSON'
type Socket_WriteJSON_Call struct {
	*mock.Call
}

// WriteJSON is a helper method to define mock.On call
//   - v interface{}
func (_e *Socket_Expecter) WriteJSON(v interface{}) *Socket_WriteJSON_Call {
	return &Socket_WriteJSON_Call{Call: _e.mock.On("WriteJSON", v)}
}

func (_c *Socket_WriteJSON_Call) Run(run func(v interface{})) *Socket_WriteJSON_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *Socket_WriteJSON_Call) Return(_a0 error) *Socket_WriteJSON_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Socket_WriteJSON_Call) RunAndReturn(run func(interface{}) error) *Socket_WriteJSON_Call {
	_c.Call.Return(run)
	return _c
}

// WriteMessage provides a mock function with given fields: messageType, data
func (_m *Socket) WriteMessage(messageType int, data []byte) error {
	ret := _m.Called(messageType, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, []byte) error); ok {
		r0 = rf(messageType, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Socket_WriteMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteMessage'
type Socket_WriteMessage_Call struct {
	*mock.Call
}

// WriteMessage is a helper method to define mock.On call
//   - messageType int
//   - data []byte
func (_e *Socket_Expecter) WriteMessage(messageType interface{}, data interface{}) *Socket_WriteMessage_Call {
	return &Socket_WriteMessage_Call{Call: _e.mock.On("WriteMessage", messageType, data)}
}

func (_c *Socket_WriteMessage_Call) Run(run func(messageType int, data []byte)) *Socket_WriteMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].([]byte))
	})
	return _c
}

func (_c *Socket_WriteMessage_Call) Return(_a0 error) *Socket_WriteMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Socket_WriteMessage_Call) RunAndReturn(run func(int, []byte) error) *Socket_WriteMessage_Call {
	_c.Call.Return(run)
	return _c
}

// NewSocket creates a new instance of Socket. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSocket(t interface {
	mock.TestingT
	Cleanup(func())
}) *Socket {
	mock := &Socket{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
