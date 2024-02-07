// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	entity "github.com/REST-API/entity"
	mock "github.com/stretchr/testify/mock"
)

// PostService is an autogenerated mock type for the PostService type
type PostService struct {
	mock.Mock
}

// Create provides a mock function with given fields: post
func (_m *PostService) Create(post *entity.Post) (*entity.Post, error) {
	ret := _m.Called(post)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *entity.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.Post) (*entity.Post, error)); ok {
		return rf(post)
	}
	if rf, ok := ret.Get(0).(func(*entity.Post) *entity.Post); ok {
		r0 = rf(post)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(*entity.Post) error); ok {
		r1 = rf(post)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields:
func (_m *PostService) FindAll() ([]entity.Post, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []entity.Post
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]entity.Post, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []entity.Post); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Post)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Validate provides a mock function with given fields: post
func (_m *PostService) Validate(post *entity.Post) error {
	ret := _m.Called(post)

	if len(ret) == 0 {
		panic("no return value specified for Validate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Post) error); ok {
		r0 = rf(post)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPostService creates a new instance of PostService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPostService(t interface {
	mock.TestingT
	Cleanup(func())
}) *PostService {
	mock := &PostService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}