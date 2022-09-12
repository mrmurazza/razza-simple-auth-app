// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	user "dealljobs/domain/user"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetAllUsers provides a mock function with given fields:
func (_m *Repository) GetAllUsers() []*user.User {
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

// GetUser provides a mock function with given fields: sku
func (_m *Repository) GetUser(sku string) user.User {
	ret := _m.Called(sku)

	var r0 user.User
	if rf, ok := ret.Get(0).(func(string) user.User); ok {
		r0 = rf(sku)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	return r0
}

// GetUsers provides a mock function with given fields: skuList
func (_m *Repository) GetUsers(skuList []string) []*user.User {
	ret := _m.Called(skuList)

	var r0 []*user.User
	if rf, ok := ret.Get(0).(func([]string) []*user.User); ok {
		r0 = rf(skuList)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*user.User)
		}
	}

	return r0
}

// IsExist provides a mock function with given fields: i
func (_m *Repository) IsExist(i user.User) bool {
	ret := _m.Called(i)

	var r0 bool
	if rf, ok := ret.Get(0).(func(user.User) bool); ok {
		r0 = rf(i)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Persist provides a mock function with given fields: i
func (_m *Repository) Persist(i user.User) {
	_m.Called(i)
}

// UpdateStock provides a mock function with given fields: i
func (_m *Repository) Update(i user.User) {
	_m.Called(i)
}