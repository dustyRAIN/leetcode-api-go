package leetcodeapi

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// IUtilMock is an autogenerated mock type for the IUtilMock type
type IUtilMock struct {
	mock.Mock
}

// MakeGraphQLRequest provides a mock function with given fields: payload, resultRef
func (_m *IUtilMock) MakeGraphQLRequest(payload string, resultRef interface{}) error {
	ret := _m.Called(payload, resultRef)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(payload, resultRef)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// convertListToString provides a mock function with given fields: list
func (_m *IUtilMock) convertListToString(list []string) string {
	ret := _m.Called(list)

	var r0 string
	if rf, ok := ret.Get(0).(func([]string) string); ok {
		r0 = rf(list)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// makeHttpRequest provides a mock function with given fields: method, url, contentType, body, resultRef
func (_m *IUtilMock) makeHttpRequest(method string, url string, contentType string, body string, resultRef interface{}) error {
	ret := _m.Called(method, url, contentType, body, resultRef)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, interface{}) error); ok {
		r0 = rf(method, url, contentType, body, resultRef)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// parseCookie provides a mock function with given fields: cookies, cookieName
func (_m *IUtilMock) parseCookie(cookies []*http.Cookie, cookieName string) string {
	ret := _m.Called(cookies, cookieName)

	var r0 string
	if rf, ok := ret.Get(0).(func([]*http.Cookie, string) string); ok {
		r0 = rf(cookies, cookieName)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewIUtilMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUtilMock creates a new instance of IUtilMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUtilMock(t mockConstructorTestingTNewIUtilMock) *IUtilMock {
	mock := &IUtilMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}