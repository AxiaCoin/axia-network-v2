// Code generated by mockery v1.0.0. DO NOT EDIT.

package tracker

import (
	ids "github.com/axiacoin/axia-network-v2/ids"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// TimeTracker is an autogenerated mock type for the TimeTracker type
type MockTimeTracker struct {
	mock.Mock
}

// CumulativeUtilization provides a mock function with given fields: _a0
func (_m *MockTimeTracker) CumulativeUtilization(_a0 time.Time) float64 {
	ret := _m.Called(_a0)

	var r0 float64
	if rf, ok := ret.Get(0).(func(time.Time) float64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// Len provides a mock function with given fields:
func (_m *MockTimeTracker) Len() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Utilization provides a mock function with given fields: _a0, _a1
func (_m *MockTimeTracker) Utilization(_a0 ids.ShortID, _a1 time.Time) float64 {
	ret := _m.Called(_a0, _a1)

	var r0 float64
	if rf, ok := ret.Get(0).(func(ids.ShortID, time.Time) float64); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// StartCPU provides a mock function with given fields: _a0, _a1
func (_m *MockTimeTracker) StartCPU(_a0 ids.ShortID, _a1 time.Time) {
	_m.Called(_a0, _a1)
}

// StopCPU provides a mock function with given fields: _a0, _a1
func (_m *MockTimeTracker) StopCPU(_a0 ids.ShortID, _a1 time.Time) {
	_m.Called(_a0, _a1)
}
