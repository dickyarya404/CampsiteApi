// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	dto "campsite_reservation/dto"
	model "campsite_reservation/model"

	mock "github.com/stretchr/testify/mock"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: payloads
func (_m *UserUsecase) Create(payloads dto.CreateUserRequest) error {
	ret := _m.Called(payloads)

	var r0 error
	if rf, ok := ret.Get(0).(func(dto.CreateUserRequest) error); ok {
		r0 = rf(payloads)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Login provides a mock function with given fields: email, password
func (_m *UserUsecase) Login(email string, password string) (model.User, string, error) {
	ret := _m.Called(email, password)

	var r0 model.User
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) (model.User, string, error)); ok {
		return rf(email, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) model.User); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(email, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SelectAll provides a mock function with given fields:
func (_m *UserUsecase) SelectAll() ([]model.User, error) {
	ret := _m.Called()

	var r0 []model.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserUsecase {
	mock := &UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
