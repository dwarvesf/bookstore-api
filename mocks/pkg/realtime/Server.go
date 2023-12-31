// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"

	realtime "github.com/dwarvesf/bookstore-api/pkg/realtime"
)

// Server is an autogenerated mock type for the Server type
type Server struct {
	mock.Mock
}

type Server_Expecter struct {
	mock *mock.Mock
}

func (_m *Server) EXPECT() *Server_Expecter {
	return &Server_Expecter{mock: &_m.Mock}
}

// BroadcastData provides a mock function with given fields: data
func (_m *Server) BroadcastData(data interface{}) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Server_BroadcastData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BroadcastData'
type Server_BroadcastData_Call struct {
	*mock.Call
}

// BroadcastData is a helper method to define mock.On call
//   - data interface{}
func (_e *Server_Expecter) BroadcastData(data interface{}) *Server_BroadcastData_Call {
	return &Server_BroadcastData_Call{Call: _e.mock.On("BroadcastData", data)}
}

func (_c *Server_BroadcastData_Call) Run(run func(data interface{})) *Server_BroadcastData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *Server_BroadcastData_Call) Return(_a0 error) *Server_BroadcastData_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Server_BroadcastData_Call) RunAndReturn(run func(interface{}) error) *Server_BroadcastData_Call {
	_c.Call.Return(run)
	return _c
}

// BroadcastMessage provides a mock function with given fields: message
func (_m *Server) BroadcastMessage(message string) error {
	ret := _m.Called(message)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Server_BroadcastMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BroadcastMessage'
type Server_BroadcastMessage_Call struct {
	*mock.Call
}

// BroadcastMessage is a helper method to define mock.On call
//   - message string
func (_e *Server_Expecter) BroadcastMessage(message interface{}) *Server_BroadcastMessage_Call {
	return &Server_BroadcastMessage_Call{Call: _e.mock.On("BroadcastMessage", message)}
}

func (_c *Server_BroadcastMessage_Call) Run(run func(message string)) *Server_BroadcastMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Server_BroadcastMessage_Call) Return(_a0 error) *Server_BroadcastMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Server_BroadcastMessage_Call) RunAndReturn(run func(string) error) *Server_BroadcastMessage_Call {
	_c.Call.Return(run)
	return _c
}

// DisconnectUser provides a mock function with given fields: u
func (_m *Server) DisconnectUser(u realtime.User) error {
	ret := _m.Called(u)

	var r0 error
	if rf, ok := ret.Get(0).(func(realtime.User) error); ok {
		r0 = rf(u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Server_DisconnectUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DisconnectUser'
type Server_DisconnectUser_Call struct {
	*mock.Call
}

// DisconnectUser is a helper method to define mock.On call
//   - u realtime.User
func (_e *Server_Expecter) DisconnectUser(u interface{}) *Server_DisconnectUser_Call {
	return &Server_DisconnectUser_Call{Call: _e.mock.On("DisconnectUser", u)}
}

func (_c *Server_DisconnectUser_Call) Run(run func(u realtime.User)) *Server_DisconnectUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(realtime.User))
	})
	return _c
}

func (_c *Server_DisconnectUser_Call) Return(_a0 error) *Server_DisconnectUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Server_DisconnectUser_Call) RunAndReturn(run func(realtime.User) error) *Server_DisconnectUser_Call {
	_c.Call.Return(run)
	return _c
}

// HandleConnection provides a mock function with given fields: c
func (_m *Server) HandleConnection(c *gin.Context) (*realtime.User, error) {
	ret := _m.Called(c)

	var r0 *realtime.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*gin.Context) (*realtime.User, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(*gin.Context) *realtime.User); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*realtime.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*gin.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Server_HandleConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleConnection'
type Server_HandleConnection_Call struct {
	*mock.Call
}

// HandleConnection is a helper method to define mock.On call
//   - c *gin.Context
func (_e *Server_Expecter) HandleConnection(c interface{}) *Server_HandleConnection_Call {
	return &Server_HandleConnection_Call{Call: _e.mock.On("HandleConnection", c)}
}

func (_c *Server_HandleConnection_Call) Run(run func(c *gin.Context)) *Server_HandleConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *Server_HandleConnection_Call) Return(_a0 *realtime.User, _a1 error) *Server_HandleConnection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Server_HandleConnection_Call) RunAndReturn(run func(*gin.Context) (*realtime.User, error)) *Server_HandleConnection_Call {
	_c.Call.Return(run)
	return _c
}

// HandleEvent provides a mock function with given fields: c, u, callback
func (_m *Server) HandleEvent(c *gin.Context, u realtime.User, callback func(*gin.Context, interface{}) error) {
	_m.Called(c, u, callback)
}

// Server_HandleEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleEvent'
type Server_HandleEvent_Call struct {
	*mock.Call
}

// HandleEvent is a helper method to define mock.On call
//   - c *gin.Context
//   - u realtime.User
//   - callback func(*gin.Context , interface{}) error
func (_e *Server_Expecter) HandleEvent(c interface{}, u interface{}, callback interface{}) *Server_HandleEvent_Call {
	return &Server_HandleEvent_Call{Call: _e.mock.On("HandleEvent", c, u, callback)}
}

func (_c *Server_HandleEvent_Call) Run(run func(c *gin.Context, u realtime.User, callback func(*gin.Context, interface{}) error)) *Server_HandleEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context), args[1].(realtime.User), args[2].(func(*gin.Context, interface{}) error))
	})
	return _c
}

func (_c *Server_HandleEvent_Call) Return() *Server_HandleEvent_Call {
	_c.Call.Return()
	return _c
}

func (_c *Server_HandleEvent_Call) RunAndReturn(run func(*gin.Context, realtime.User, func(*gin.Context, interface{}) error)) *Server_HandleEvent_Call {
	_c.Call.Return(run)
	return _c
}

// SendData provides a mock function with given fields: userID, data
func (_m *Server) SendData(userID string, data interface{}) error {
	ret := _m.Called(userID, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(userID, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Server_SendData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendData'
type Server_SendData_Call struct {
	*mock.Call
}

// SendData is a helper method to define mock.On call
//   - userID string
//   - data interface{}
func (_e *Server_Expecter) SendData(userID interface{}, data interface{}) *Server_SendData_Call {
	return &Server_SendData_Call{Call: _e.mock.On("SendData", userID, data)}
}

func (_c *Server_SendData_Call) Run(run func(userID string, data interface{})) *Server_SendData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(interface{}))
	})
	return _c
}

func (_c *Server_SendData_Call) Return(_a0 error) *Server_SendData_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Server_SendData_Call) RunAndReturn(run func(string, interface{}) error) *Server_SendData_Call {
	_c.Call.Return(run)
	return _c
}

// SendMessage provides a mock function with given fields: userID, message
func (_m *Server) SendMessage(userID string, message string) error {
	ret := _m.Called(userID, message)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userID, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Server_SendMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMessage'
type Server_SendMessage_Call struct {
	*mock.Call
}

// SendMessage is a helper method to define mock.On call
//   - userID string
//   - message string
func (_e *Server_Expecter) SendMessage(userID interface{}, message interface{}) *Server_SendMessage_Call {
	return &Server_SendMessage_Call{Call: _e.mock.On("SendMessage", userID, message)}
}

func (_c *Server_SendMessage_Call) Run(run func(userID string, message string)) *Server_SendMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *Server_SendMessage_Call) Return(_a0 error) *Server_SendMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Server_SendMessage_Call) RunAndReturn(run func(string, string) error) *Server_SendMessage_Call {
	_c.Call.Return(run)
	return _c
}

// NewServer creates a new instance of Server. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Server {
	mock := &Server{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
