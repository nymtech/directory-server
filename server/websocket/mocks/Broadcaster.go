// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Broadcaster is an autogenerated mock type for the Broadcaster type
type Broadcaster struct {
	mock.Mock
}

// Notify provides a mock function with given fields: msg
func (_m *Broadcaster) Notify(msg []byte) {
	_m.Called(msg)
}
