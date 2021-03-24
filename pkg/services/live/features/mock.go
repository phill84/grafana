// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/grafana/grafana/pkg/services/live/features (interfaces: ChannelPublisher,PresenceGetter,PluginContextGetter,StreamRunner)

// Package features is a generated GoMock package.
package features

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	backend "github.com/grafana/grafana-plugin-sdk-go/backend"
)

// MockChannelPublisher is a mock of ChannelPublisher interface.
type MockChannelPublisher struct {
	ctrl     *gomock.Controller
	recorder *MockChannelPublisherMockRecorder
}

// MockChannelPublisherMockRecorder is the mock recorder for MockChannelPublisher.
type MockChannelPublisherMockRecorder struct {
	mock *MockChannelPublisher
}

// NewMockChannelPublisher creates a new mock instance.
func NewMockChannelPublisher(ctrl *gomock.Controller) *MockChannelPublisher {
	mock := &MockChannelPublisher{ctrl: ctrl}
	mock.recorder = &MockChannelPublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChannelPublisher) EXPECT() *MockChannelPublisherMockRecorder {
	return m.recorder
}

// Publish mocks base method.
func (m *MockChannelPublisher) Publish(arg0 string, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockChannelPublisherMockRecorder) Publish(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockChannelPublisher)(nil).Publish), arg0, arg1)
}

// MockPresenceGetter is a mock of PresenceGetter interface.
type MockPresenceGetter struct {
	ctrl     *gomock.Controller
	recorder *MockPresenceGetterMockRecorder
}

// MockPresenceGetterMockRecorder is the mock recorder for MockPresenceGetter.
type MockPresenceGetterMockRecorder struct {
	mock *MockPresenceGetter
}

// NewMockPresenceGetter creates a new mock instance.
func NewMockPresenceGetter(ctrl *gomock.Controller) *MockPresenceGetter {
	mock := &MockPresenceGetter{ctrl: ctrl}
	mock.recorder = &MockPresenceGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPresenceGetter) EXPECT() *MockPresenceGetterMockRecorder {
	return m.recorder
}

// GetNumSubscribers mocks base method.
func (m *MockPresenceGetter) GetNumSubscribers(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNumSubscribers", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNumSubscribers indicates an expected call of GetNumSubscribers.
func (mr *MockPresenceGetterMockRecorder) GetNumSubscribers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNumSubscribers", reflect.TypeOf((*MockPresenceGetter)(nil).GetNumSubscribers), arg0)
}

// MockPluginContextGetter is a mock of PluginContextGetter interface.
type MockPluginContextGetter struct {
	ctrl     *gomock.Controller
	recorder *MockPluginContextGetterMockRecorder
}

// MockPluginContextGetterMockRecorder is the mock recorder for MockPluginContextGetter.
type MockPluginContextGetterMockRecorder struct {
	mock *MockPluginContextGetter
}

// NewMockPluginContextGetter creates a new mock instance.
func NewMockPluginContextGetter(ctrl *gomock.Controller) *MockPluginContextGetter {
	mock := &MockPluginContextGetter{ctrl: ctrl}
	mock.recorder = &MockPluginContextGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPluginContextGetter) EXPECT() *MockPluginContextGetterMockRecorder {
	return m.recorder
}

// GetPluginContext mocks base method.
func (m *MockPluginContextGetter) GetPluginContext(arg0 context.Context, arg1, arg2 string) (backend.PluginContext, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPluginContext", arg0, arg1, arg2)
	ret0, _ := ret[0].(backend.PluginContext)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPluginContext indicates an expected call of GetPluginContext.
func (mr *MockPluginContextGetterMockRecorder) GetPluginContext(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPluginContext", reflect.TypeOf((*MockPluginContextGetter)(nil).GetPluginContext), arg0, arg1, arg2)
}

// MockStreamRunner is a mock of StreamRunner interface.
type MockStreamRunner struct {
	ctrl     *gomock.Controller
	recorder *MockStreamRunnerMockRecorder
}

// MockStreamRunnerMockRecorder is the mock recorder for MockStreamRunner.
type MockStreamRunnerMockRecorder struct {
	mock *MockStreamRunner
}

// NewMockStreamRunner creates a new mock instance.
func NewMockStreamRunner(ctrl *gomock.Controller) *MockStreamRunner {
	mock := &MockStreamRunner{ctrl: ctrl}
	mock.recorder = &MockStreamRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamRunner) EXPECT() *MockStreamRunnerMockRecorder {
	return m.recorder
}

// RunStream mocks base method.
func (m *MockStreamRunner) RunStream(arg0 context.Context, arg1 *backend.RunStreamRequest, arg2 backend.StreamPacketSender) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunStream", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunStream indicates an expected call of RunStream.
func (mr *MockStreamRunnerMockRecorder) RunStream(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunStream", reflect.TypeOf((*MockStreamRunner)(nil).RunStream), arg0, arg1, arg2)
}
