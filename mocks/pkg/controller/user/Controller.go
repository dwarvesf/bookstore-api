// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/dwarvesf/bookstore-api/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

type Controller_Expecter struct {
	mock *mock.Mock
}

func (_m *Controller) EXPECT() *Controller_Expecter {
	return &Controller_Expecter{mock: &_m.Mock}
}

// Me provides a mock function with given fields: ctx
func (_m *Controller) Me(ctx context.Context) (*model.User, error) {
	ret := _m.Called(ctx)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*model.User, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *model.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Controller_Me_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Me'
type Controller_Me_Call struct {
	*mock.Call
}

// Me is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Controller_Expecter) Me(ctx interface{}) *Controller_Me_Call {
	return &Controller_Me_Call{Call: _e.mock.On("Me", ctx)}
}

func (_c *Controller_Me_Call) Run(run func(ctx context.Context)) *Controller_Me_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Controller_Me_Call) Return(_a0 *model.User, _a1 error) *Controller_Me_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Controller_Me_Call) RunAndReturn(run func(context.Context) (*model.User, error)) *Controller_Me_Call {
	_c.Call.Return(run)
	return _c
}

// UpdatePassword provides a mock function with given fields: ctx, _a1
func (_m *Controller) UpdatePassword(ctx context.Context, _a1 model.UpdatePasswordRequest) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.UpdatePasswordRequest) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Controller_UpdatePassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdatePassword'
type Controller_UpdatePassword_Call struct {
	*mock.Call
}

// UpdatePassword is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 model.UpdatePasswordRequest
func (_e *Controller_Expecter) UpdatePassword(ctx interface{}, _a1 interface{}) *Controller_UpdatePassword_Call {
	return &Controller_UpdatePassword_Call{Call: _e.mock.On("UpdatePassword", ctx, _a1)}
}

func (_c *Controller_UpdatePassword_Call) Run(run func(ctx context.Context, _a1 model.UpdatePasswordRequest)) *Controller_UpdatePassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UpdatePasswordRequest))
	})
	return _c
}

func (_c *Controller_UpdatePassword_Call) Return(_a0 error) *Controller_UpdatePassword_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Controller_UpdatePassword_Call) RunAndReturn(run func(context.Context, model.UpdatePasswordRequest) error) *Controller_UpdatePassword_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUser provides a mock function with given fields: ctx, _a1
func (_m *Controller) UpdateUser(ctx context.Context, _a1 model.UpdateUserRequest) (*model.User, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.UpdateUserRequest) (*model.User, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.UpdateUserRequest) *model.User); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.UpdateUserRequest) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Controller_UpdateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUser'
type Controller_UpdateUser_Call struct {
	*mock.Call
}

// UpdateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 model.UpdateUserRequest
func (_e *Controller_Expecter) UpdateUser(ctx interface{}, _a1 interface{}) *Controller_UpdateUser_Call {
	return &Controller_UpdateUser_Call{Call: _e.mock.On("UpdateUser", ctx, _a1)}
}

func (_c *Controller_UpdateUser_Call) Run(run func(ctx context.Context, _a1 model.UpdateUserRequest)) *Controller_UpdateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UpdateUserRequest))
	})
	return _c
}

func (_c *Controller_UpdateUser_Call) Return(_a0 *model.User, _a1 error) *Controller_UpdateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Controller_UpdateUser_Call) RunAndReturn(run func(context.Context, model.UpdateUserRequest) (*model.User, error)) *Controller_UpdateUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewController creates a new instance of Controller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewController(t interface {
	mock.TestingT
	Cleanup(func())
}) *Controller {
	mock := &Controller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
