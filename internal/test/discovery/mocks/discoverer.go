// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	discovery "github.com/jorgebay/soda/internal/discovery"
	mock "github.com/stretchr/testify/mock"

	types "github.com/jorgebay/soda/internal/types"

	uuid "github.com/google/uuid"
)

// Discoverer is an autogenerated mock type for the Discoverer type
type Discoverer struct {
	mock.Mock
}

// Brokers provides a mock function with given fields:
func (_m *Discoverer) Brokers() []types.BrokerInfo {
	ret := _m.Called()

	var r0 []types.BrokerInfo
	if rf, ok := ret.Get(0).(func() []types.BrokerInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.BrokerInfo)
		}
	}

	return r0
}

// Generation provides a mock function with given fields: token
func (_m *Discoverer) Generation(token types.Token) *types.Generation {
	ret := _m.Called(token)

	var r0 *types.Generation
	if rf, ok := ret.Get(0).(func(types.Token) *types.Generation); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Generation)
		}
	}

	return r0
}

// GenerationProposed provides a mock function with given fields: token
func (_m *Discoverer) GenerationProposed(token types.Token) (*types.Generation, *types.Generation) {
	ret := _m.Called(token)

	var r0 *types.Generation
	if rf, ok := ret.Get(0).(func(types.Token) *types.Generation); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Generation)
		}
	}

	var r1 *types.Generation
	if rf, ok := ret.Get(1).(func(types.Token) *types.Generation); ok {
		r1 = rf(token)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*types.Generation)
		}
	}

	return r0, r1
}

// Init provides a mock function with given fields:
func (_m *Discoverer) Init() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Leader provides a mock function with given fields: partitionKey
func (_m *Discoverer) Leader(partitionKey string) types.ReplicationInfo {
	ret := _m.Called(partitionKey)

	var r0 types.ReplicationInfo
	if rf, ok := ret.Get(0).(func(string) types.ReplicationInfo); ok {
		r0 = rf(partitionKey)
	} else {
		r0 = ret.Get(0).(types.ReplicationInfo)
	}

	return r0
}

// LocalInfo provides a mock function with given fields:
func (_m *Discoverer) LocalInfo() *types.BrokerInfo {
	ret := _m.Called()

	var r0 *types.BrokerInfo
	if rf, ok := ret.Get(0).(func() *types.BrokerInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.BrokerInfo)
		}
	}

	return r0
}

// Peers provides a mock function with given fields:
func (_m *Discoverer) Peers() []types.BrokerInfo {
	ret := _m.Called()

	var r0 []types.BrokerInfo
	if rf, ok := ret.Get(0).(func() []types.BrokerInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.BrokerInfo)
		}
	}

	return r0
}

// RegisterListener provides a mock function with given fields: l
func (_m *Discoverer) RegisterListener(l discovery.TopologyChangeHandler) {
	_m.Called(l)
}

// SetAsCommitted provides a mock function with given fields: token, tx
func (_m *Discoverer) SetAsCommitted(token types.Token, tx uuid.UUID) error {
	ret := _m.Called(token, tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Token, uuid.UUID) error); ok {
		r0 = rf(token, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetGenerationProposed provides a mock function with given fields: gen, expectedTx
func (_m *Discoverer) SetGenerationProposed(gen types.Generation, expectedTx *uuid.UUID) error {
	ret := _m.Called(gen, expectedTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Generation, *uuid.UUID) error); ok {
		r0 = rf(gen, expectedTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Shutdown provides a mock function with given fields:
func (_m *Discoverer) Shutdown() {
	_m.Called()
}

// TokenByOrdinal provides a mock function with given fields: ordinal
func (_m *Discoverer) TokenByOrdinal(ordinal int) types.Token {
	ret := _m.Called(ordinal)

	var r0 types.Token
	if rf, ok := ret.Get(0).(func(int) types.Token); ok {
		r0 = rf(ordinal)
	} else {
		r0 = ret.Get(0).(types.Token)
	}

	return r0
}