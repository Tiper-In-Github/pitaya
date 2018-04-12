// Code generated by MockGen. DO NOT EDIT.
// Source: cluster/rpc_interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	cluster "github.com/topfreegames/pitaya/cluster"
	message "github.com/topfreegames/pitaya/internal/message"
	protos "github.com/topfreegames/pitaya/protos"
	route "github.com/topfreegames/pitaya/route"
	session "github.com/topfreegames/pitaya/session"
	reflect "reflect"
)

// MockRPCServer is a mock of RPCServer interface
type MockRPCServer struct {
	ctrl     *gomock.Controller
	recorder *MockRPCServerMockRecorder
}

// MockRPCServerMockRecorder is the mock recorder for MockRPCServer
type MockRPCServerMockRecorder struct {
	mock *MockRPCServer
}

// NewMockRPCServer creates a new mock instance
func NewMockRPCServer(ctrl *gomock.Controller) *MockRPCServer {
	mock := &MockRPCServer{ctrl: ctrl}
	mock.recorder = &MockRPCServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRPCServer) EXPECT() *MockRPCServerMockRecorder {
	return m.recorder
}

// GetUnhandledRequestsChannel mocks base method
func (m *MockRPCServer) GetUnhandledRequestsChannel() chan *protos.Request {
	ret := m.ctrl.Call(m, "GetUnhandledRequestsChannel")
	ret0, _ := ret[0].(chan *protos.Request)
	return ret0
}

// GetUnhandledRequestsChannel indicates an expected call of GetUnhandledRequestsChannel
func (mr *MockRPCServerMockRecorder) GetUnhandledRequestsChannel() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnhandledRequestsChannel", reflect.TypeOf((*MockRPCServer)(nil).GetUnhandledRequestsChannel))
}

// GetUserPushChannel mocks base method
func (m *MockRPCServer) GetUserPushChannel() chan *protos.Push {
	ret := m.ctrl.Call(m, "GetUserPushChannel")
	ret0, _ := ret[0].(chan *protos.Push)
	return ret0
}

// GetUserPushChannel indicates an expected call of GetUserPushChannel
func (mr *MockRPCServerMockRecorder) GetUserPushChannel() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPushChannel", reflect.TypeOf((*MockRPCServer)(nil).GetUserPushChannel))
}

// Init mocks base method
func (m *MockRPCServer) Init() error {
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockRPCServerMockRecorder) Init() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockRPCServer)(nil).Init))
}

// AfterInit mocks base method
func (m *MockRPCServer) AfterInit() {
	m.ctrl.Call(m, "AfterInit")
}

// AfterInit indicates an expected call of AfterInit
func (mr *MockRPCServerMockRecorder) AfterInit() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterInit", reflect.TypeOf((*MockRPCServer)(nil).AfterInit))
}

// BeforeShutdown mocks base method
func (m *MockRPCServer) BeforeShutdown() {
	m.ctrl.Call(m, "BeforeShutdown")
}

// BeforeShutdown indicates an expected call of BeforeShutdown
func (mr *MockRPCServerMockRecorder) BeforeShutdown() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeShutdown", reflect.TypeOf((*MockRPCServer)(nil).BeforeShutdown))
}

// Shutdown mocks base method
func (m *MockRPCServer) Shutdown() error {
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown
func (mr *MockRPCServerMockRecorder) Shutdown() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockRPCServer)(nil).Shutdown))
}

// MockRPCClient is a mock of RPCClient interface
type MockRPCClient struct {
	ctrl     *gomock.Controller
	recorder *MockRPCClientMockRecorder
}

// MockRPCClientMockRecorder is the mock recorder for MockRPCClient
type MockRPCClientMockRecorder struct {
	mock *MockRPCClient
}

// NewMockRPCClient creates a new mock instance
func NewMockRPCClient(ctrl *gomock.Controller) *MockRPCClient {
	mock := &MockRPCClient{ctrl: ctrl}
	mock.recorder = &MockRPCClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRPCClient) EXPECT() *MockRPCClientMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockRPCClient) Send(route string, data []byte) error {
	ret := m.ctrl.Call(m, "Send", route, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockRPCClientMockRecorder) Send(route, data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockRPCClient)(nil).Send), route, data)
}

// Call mocks base method
func (m *MockRPCClient) Call(rpcType protos.RPCType, route *route.Route, session *session.Session, msg *message.Message, server *cluster.Server) (*protos.Response, error) {
	ret := m.ctrl.Call(m, "Call", rpcType, route, session, msg, server)
	ret0, _ := ret[0].(*protos.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call
func (mr *MockRPCClientMockRecorder) Call(rpcType, route, session, msg, server interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockRPCClient)(nil).Call), rpcType, route, session, msg, server)
}

// Init mocks base method
func (m *MockRPCClient) Init() error {
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockRPCClientMockRecorder) Init() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockRPCClient)(nil).Init))
}

// AfterInit mocks base method
func (m *MockRPCClient) AfterInit() {
	m.ctrl.Call(m, "AfterInit")
}

// AfterInit indicates an expected call of AfterInit
func (mr *MockRPCClientMockRecorder) AfterInit() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterInit", reflect.TypeOf((*MockRPCClient)(nil).AfterInit))
}

// BeforeShutdown mocks base method
func (m *MockRPCClient) BeforeShutdown() {
	m.ctrl.Call(m, "BeforeShutdown")
}

// BeforeShutdown indicates an expected call of BeforeShutdown
func (mr *MockRPCClientMockRecorder) BeforeShutdown() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeShutdown", reflect.TypeOf((*MockRPCClient)(nil).BeforeShutdown))
}

// Shutdown mocks base method
func (m *MockRPCClient) Shutdown() error {
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown
func (mr *MockRPCClientMockRecorder) Shutdown() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockRPCClient)(nil).Shutdown))
}
