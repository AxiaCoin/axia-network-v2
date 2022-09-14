// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	ids "github.com/axiacoin/axia-network-v2/ids"
	common "github.com/axiacoin/axia-network-v2/snow/engine/common"

	manager "github.com/axiacoin/axia-network-v2/database/manager"

	mock "github.com/stretchr/testify/mock"

	snow "github.com/axiacoin/axia-network-v2/snow"

	snowstorm "github.com/axiacoin/axia-network-v2/snow/consensus/snowstorm"

	time "time"

	version "github.com/axiacoin/axia-network-v2/version"
)

// DAGVM is an autogenerated mock type for the DAGVM type
type DAGVM struct {
	mock.Mock
}

// AppGossip provides a mock function with given fields: nodeID, msg
func (_m *DAGVM) AppGossip(nodeID ids.ShortID, msg []byte) error {
	ret := _m.Called(nodeID, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, []byte) error); ok {
		r0 = rf(nodeID, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppRequest provides a mock function with given fields: nodeID, requestID, deadline, request
func (_m *DAGVM) AppRequest(nodeID ids.ShortID, requestID uint32, deadline time.Time, request []byte) error {
	ret := _m.Called(nodeID, requestID, deadline, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, uint32, time.Time, []byte) error); ok {
		r0 = rf(nodeID, requestID, deadline, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppRequestFailed provides a mock function with given fields: nodeID, requestID
func (_m *DAGVM) AppRequestFailed(nodeID ids.ShortID, requestID uint32) error {
	ret := _m.Called(nodeID, requestID)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, uint32) error); ok {
		r0 = rf(nodeID, requestID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppResponse provides a mock function with given fields: nodeID, requestID, response
func (_m *DAGVM) AppResponse(nodeID ids.ShortID, requestID uint32, response []byte) error {
	ret := _m.Called(nodeID, requestID, response)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, uint32, []byte) error); ok {
		r0 = rf(nodeID, requestID, response)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Connected provides a mock function with given fields: id, nodeVersion
func (_m *DAGVM) Connected(id ids.ShortID, nodeVersion version.Application) error {
	ret := _m.Called(id, nodeVersion)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID, version.Application) error); ok {
		r0 = rf(id, nodeVersion)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateHandlers provides a mock function with given fields:
func (_m *DAGVM) CreateHandlers() (map[string]*common.HTTPHandler, error) {
	ret := _m.Called()

	var r0 map[string]*common.HTTPHandler
	if rf, ok := ret.Get(0).(func() map[string]*common.HTTPHandler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*common.HTTPHandler)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateStaticHandlers provides a mock function with given fields:
func (_m *DAGVM) CreateStaticHandlers() (map[string]*common.HTTPHandler, error) {
	ret := _m.Called()

	var r0 map[string]*common.HTTPHandler
	if rf, ok := ret.Get(0).(func() map[string]*common.HTTPHandler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*common.HTTPHandler)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Disconnected provides a mock function with given fields: id
func (_m *DAGVM) Disconnected(id ids.ShortID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(ids.ShortID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTx provides a mock function with given fields: _a0
func (_m *DAGVM) GetTx(_a0 ids.ID) (snowstorm.Tx, error) {
	ret := _m.Called(_a0)

	var r0 snowstorm.Tx
	if rf, ok := ret.Get(0).(func(ids.ID) snowstorm.Tx); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(snowstorm.Tx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ids.ID) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HealthCheck provides a mock function with given fields:
func (_m *DAGVM) HealthCheck() (interface{}, error) {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Initialize provides a mock function with given fields: ctx, dbManager, genesisBytes, upgradeBytes, configBytes, toEngine, fxs, appSender
func (_m *DAGVM) Initialize(ctx *snow.Context, dbManager manager.Manager, genesisBytes []byte, upgradeBytes []byte, configBytes []byte, toEngine chan<- common.Message, fxs []*common.Fx, appSender common.AppSender) error {
	ret := _m.Called(ctx, dbManager, genesisBytes, upgradeBytes, configBytes, toEngine, fxs, appSender)

	var r0 error
	if rf, ok := ret.Get(0).(func(*snow.Context, manager.Manager, []byte, []byte, []byte, chan<- common.Message, []*common.Fx, common.AppSender) error); ok {
		r0 = rf(ctx, dbManager, genesisBytes, upgradeBytes, configBytes, toEngine, fxs, appSender)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ParseTx provides a mock function with given fields: tx
func (_m *DAGVM) ParseTx(tx []byte) (snowstorm.Tx, error) {
	ret := _m.Called(tx)

	var r0 snowstorm.Tx
	if rf, ok := ret.Get(0).(func([]byte) snowstorm.Tx); ok {
		r0 = rf(tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(snowstorm.Tx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PendingTxs provides a mock function with given fields:
func (_m *DAGVM) PendingTxs() []snowstorm.Tx {
	ret := _m.Called()

	var r0 []snowstorm.Tx
	if rf, ok := ret.Get(0).(func() []snowstorm.Tx); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]snowstorm.Tx)
		}
	}

	return r0
}

// SetState provides a mock function with given fields: state
func (_m *DAGVM) SetState(state snow.State) error {
	ret := _m.Called(state)

	var r0 error
	if rf, ok := ret.Get(0).(func(snow.State) error); ok {
		r0 = rf(state)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Shutdown provides a mock function with given fields:
func (_m *DAGVM) Shutdown() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Version provides a mock function with given fields:
func (_m *DAGVM) Version() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
