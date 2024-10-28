// Code generated by mockery v2.45.1. DO NOT EDIT.

package repository

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	model "github.com/thinhhuy997/go-windows/internal/model"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

type MockRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepository) EXPECT() *MockRepository_Expecter {
	return &MockRepository_Expecter{mock: &_m.Mock}
}

// All provides a mock function with given fields: ctx
func (_m *MockRepository) All(ctx context.Context) ([]model.Todo, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for All")
	}

	var r0 []model.Todo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]model.Todo, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []model.Todo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Todo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_All_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'All'
type MockRepository_All_Call struct {
	*mock.Call
}

// All is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockRepository_Expecter) All(ctx interface{}) *MockRepository_All_Call {
	return &MockRepository_All_Call{Call: _e.mock.On("All", ctx)}
}

func (_c *MockRepository_All_Call) Run(run func(ctx context.Context)) *MockRepository_All_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockRepository_All_Call) Return(_a0 []model.Todo, _a1 error) *MockRepository_All_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_All_Call) RunAndReturn(run func(context.Context) ([]model.Todo, error)) *MockRepository_All_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, todo
func (_m *MockRepository) Create(ctx context.Context, todo model.Todo) error {
	ret := _m.Called(ctx, todo)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Todo) error); ok {
		r0 = rf(ctx, todo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - todo model.Todo
func (_e *MockRepository_Expecter) Create(ctx interface{}, todo interface{}) *MockRepository_Create_Call {
	return &MockRepository_Create_Call{Call: _e.mock.On("Create", ctx, todo)}
}

func (_c *MockRepository_Create_Call) Run(run func(ctx context.Context, todo model.Todo)) *MockRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.Todo))
	})
	return _c
}

func (_c *MockRepository_Create_Call) Return(_a0 error) *MockRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_Create_Call) RunAndReturn(run func(context.Context, model.Todo) error) *MockRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *MockRepository) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id int
func (_e *MockRepository_Expecter) Delete(ctx interface{}, id interface{}) *MockRepository_Delete_Call {
	return &MockRepository_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *MockRepository_Delete_Call) Run(run func(ctx context.Context, id int)) *MockRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockRepository_Delete_Call) Return(_a0 error) *MockRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_Delete_Call) RunAndReturn(run func(context.Context, int) error) *MockRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
