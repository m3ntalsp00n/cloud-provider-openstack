// Code generated by mockery v1.0.0. DO NOT EDIT.

package keystone

import mock "github.com/stretchr/testify/mock"

// MockIKeystone is an autogenerated mock type for the IKeystone type
type MockIKeystone struct {
	mock.Mock
}

// GetGroups provides a mock function with given fields: _a0, _a1
func (_m *MockIKeystone) GetGroups(_a0 string, _a1 string) ([]string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, string) []string); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTokenInfo provides a mock function with given fields: _a0
func (_m *MockIKeystone) GetTokenInfo(_a0 string) (*tokenInfo, error) {
	ret := _m.Called(_a0)

	var r0 *tokenInfo
	if rf, ok := ret.Get(0).(func(string) *tokenInfo); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tokenInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
