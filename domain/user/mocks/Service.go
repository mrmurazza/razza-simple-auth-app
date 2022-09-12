// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	user "dealljobs/domain/user"

	mock "github.com/stretchr/testify/mock"

	request "dealljobs/domain/request"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateUserIfNotAny provides a mock function with given fields: _a0
func (_m *Service) CreateUserIfNotAny(_a0 request.CreateUserRequest) {
	_m.Called(_a0)
}

// GetAllUsers provides a mock function with given fields:
func (_m *Service) GetAllUsers() []*user.User {
	ret := _m.Called()

	var r0 []*user.User
	if rf, ok := ret.Get(0).(func() []*user.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*user.User)
		}
	}

	return r0
}

// GetRequestedUserMap provides a mock function with given fields: requestedSkuList
func (_m *Service) GetRequestedUserMap(requestedSkuList []string) map[string]*user.User {
	ret := _m.Called(requestedSkuList)

	var r0 map[string]*user.User
	if rf, ok := ret.Get(0).(func([]string) map[string]*user.User); ok {
		r0 = rf(requestedSkuList)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*user.User)
		}
	}

	return r0
}

// UpdateUserStock provides a mock function with given fields: sku, quantity
func (_m *Service) UpdateUserStock(sku string, quantity int) {
	_m.Called(sku, quantity)
}