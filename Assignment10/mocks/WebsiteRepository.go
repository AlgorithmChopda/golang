// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	internal "github.com/AlgorithmChopda/Website-Checker.git/internal"
	mock "github.com/stretchr/testify/mock"
)

// WebsiteRepository is an autogenerated mock type for the WebsiteRepository type
type WebsiteRepository struct {
	mock.Mock
}

// addWebsite provides a mock function with given fields: websiteUrl, isActive
func (_m *WebsiteRepository) addWebsite(websiteUrl string, isActive bool) {
	_m.Called(websiteUrl, isActive)
}

// deleteWebsite provides a mock function with given fields: websiteUrl
func (_m *WebsiteRepository) deleteWebsite(websiteUrl string) {
	_m.Called(websiteUrl)
}

// getAllStatus provides a mock function with given fields:
func (_m *WebsiteRepository) getAllStatus() internal.WebsiteDB {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for getAllStatus")
	}

	var r0 internal.WebsiteDB
	if rf, ok := ret.Get(0).(func() internal.WebsiteDB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(internal.WebsiteDB)
		}
	}

	return r0
}

// getStatus provides a mock function with given fields: websiteUrl
func (_m *WebsiteRepository) getStatus(websiteUrl string) internal.Status {
	ret := _m.Called(websiteUrl)

	if len(ret) == 0 {
		panic("no return value specified for getStatus")
	}

	var r0 internal.Status
	if rf, ok := ret.Get(0).(func(string) internal.Status); ok {
		r0 = rf(websiteUrl)
	} else {
		r0 = ret.Get(0).(internal.Status)
	}

	return r0
}

// getWebsites provides a mock function with given fields:
func (_m *WebsiteRepository) getWebsites() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for getWebsites")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// isPresent provides a mock function with given fields: websiteUrl
func (_m *WebsiteRepository) isPresent(websiteUrl string) bool {
	ret := _m.Called(websiteUrl)

	if len(ret) == 0 {
		panic("no return value specified for isPresent")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(websiteUrl)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// updateStatus provides a mock function with given fields: websiteUrl, isActive
func (_m *WebsiteRepository) updateStatus(websiteUrl string, isActive bool) {
	_m.Called(websiteUrl, isActive)
}

// NewWebsiteRepository creates a new instance of WebsiteRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWebsiteRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *WebsiteRepository {
	mock := &WebsiteRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
