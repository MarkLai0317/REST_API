// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	entity "github.com/REST-API/entity"
	mock "github.com/stretchr/testify/mock"
)

// PostRepository is an autogenerated mock type for the PostRepository type
type PostRepository struct {
	mock.Mock
}

// FindAll provides a mock function with given fields:
func (_m *PostRepository) FindAll() ([]entity.Post, error) {
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

// Save provides a mock function with given fields: post
func (_m *PostRepository) Save(post *entity.Post) (*entity.Post, error) {
	ret := _m.Called(post)

	if len(ret) == 0 {
		panic("no return value specified for Save")
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

// NewPostRepository creates a new instance of PostRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPostRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *PostRepository {
	mock := &PostRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}