// Code generated by mockery v1.1.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	time "time"

	wfv1 "github.com/nholuongut/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
)

// Throttler is an autogenerated mock type for the Throttler type
type Throttler struct {
	mock.Mock
}

// Init provides a mock function with given running workflows
func (_m *Throttler) Init(wfs []wfv1.Workflow) error {
	_m.Called(wfs)
	return nil
}

// Add provides a mock function with given fields: key, priority, creationTime
func (_m *Throttler) Add(key string, priority int32, creationTime time.Time) {
	_m.Called(key, priority, creationTime)
}

// Admit provides a mock function with given fields: key
func (_m *Throttler) Admit(key string) bool {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Remove provides a mock function with given fields: key
func (_m *Throttler) Remove(key string) {
	_m.Called(key)
}
