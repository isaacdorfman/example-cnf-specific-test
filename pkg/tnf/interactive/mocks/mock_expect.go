// Code generated by MockGen. DO NOT EDIT.
// Source: ./expect.go

// Package mock_interactive is a generated GoMock package.
package mock_interactive

import (
	reflect "reflect"
	regexp "regexp"
	time "time"

	gomock "github.com/golang/mock/gomock"
	goexpect "github.com/google/goexpect"
)

// MockBatcher is a mock of Batcher interface.
type MockBatcher struct {
	ctrl     *gomock.Controller
	recorder *MockBatcherMockRecorder
}

// MockBatcherMockRecorder is the mock recorder for MockBatcher.
type MockBatcherMockRecorder struct {
	mock *MockBatcher
}

// NewMockBatcher creates a new mock instance.
func NewMockBatcher(ctrl *gomock.Controller) *MockBatcher {
	mock := &MockBatcher{ctrl: ctrl}
	mock.recorder = &MockBatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBatcher) EXPECT() *MockBatcherMockRecorder {
	return m.recorder
}

// Arg mocks base method.
func (m *MockBatcher) Arg() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Arg")
	ret0, _ := ret[0].(string)
	return ret0
}

// Arg indicates an expected call of Arg.
func (mr *MockBatcherMockRecorder) Arg() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Arg", reflect.TypeOf((*MockBatcher)(nil).Arg))
}

// Cases mocks base method.
func (m *MockBatcher) Cases() []goexpect.Caser {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cases")
	ret0, _ := ret[0].([]goexpect.Caser)
	return ret0
}

// Cases indicates an expected call of Cases.
func (mr *MockBatcherMockRecorder) Cases() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cases", reflect.TypeOf((*MockBatcher)(nil).Cases))
}

// Cmd mocks base method.
func (m *MockBatcher) Cmd() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cmd")
	ret0, _ := ret[0].(int)
	return ret0
}

// Cmd indicates an expected call of Cmd.
func (mr *MockBatcherMockRecorder) Cmd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cmd", reflect.TypeOf((*MockBatcher)(nil).Cmd))
}

// Timeout mocks base method.
func (m *MockBatcher) Timeout() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Timeout")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// Timeout indicates an expected call of Timeout.
func (mr *MockBatcherMockRecorder) Timeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timeout", reflect.TypeOf((*MockBatcher)(nil).Timeout))
}

// MockCaser is a mock of Caser interface.
type MockCaser struct {
	ctrl     *gomock.Controller
	recorder *MockCaserMockRecorder
}

// MockCaserMockRecorder is the mock recorder for MockCaser.
type MockCaserMockRecorder struct {
	mock *MockCaser
}

// NewMockCaser creates a new mock instance.
func NewMockCaser(ctrl *gomock.Controller) *MockCaser {
	mock := &MockCaser{ctrl: ctrl}
	mock.recorder = &MockCaserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCaser) EXPECT() *MockCaserMockRecorder {
	return m.recorder
}

// RE mocks base method.
func (m *MockCaser) RE() (*regexp.Regexp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RE")
	ret0, _ := ret[0].(*regexp.Regexp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RE indicates an expected call of RE.
func (mr *MockCaserMockRecorder) RE() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RE", reflect.TypeOf((*MockCaser)(nil).RE))
}

// Retry mocks base method.
func (m *MockCaser) Retry() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Retry")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Retry indicates an expected call of Retry.
func (mr *MockCaserMockRecorder) Retry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Retry", reflect.TypeOf((*MockCaser)(nil).Retry))
}

// String mocks base method.
func (m *MockCaser) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockCaserMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockCaser)(nil).String))
}

// Tag mocks base method.
func (m *MockCaser) Tag() (goexpect.Tag, *goexpect.Status) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tag")
	ret0, _ := ret[0].(goexpect.Tag)
	ret1, _ := ret[1].(*goexpect.Status)
	return ret0, ret1
}

// Tag indicates an expected call of Tag.
func (mr *MockCaserMockRecorder) Tag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockCaser)(nil).Tag))
}

// MockExpecter is a mock of Expecter interface.
type MockExpecter struct {
	ctrl     *gomock.Controller
	recorder *MockExpecterMockRecorder
}

// MockExpecterMockRecorder is the mock recorder for MockExpecter.
type MockExpecterMockRecorder struct {
	mock *MockExpecter
}

// NewMockExpecter creates a new mock instance.
func NewMockExpecter(ctrl *gomock.Controller) *MockExpecter {
	mock := &MockExpecter{ctrl: ctrl}
	mock.recorder = &MockExpecterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExpecter) EXPECT() *MockExpecterMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockExpecter) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockExpecterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockExpecter)(nil).Close))
}

// Expect mocks base method.
func (m *MockExpecter) Expect(arg0 *regexp.Regexp, arg1 time.Duration) (string, []string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Expect", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Expect indicates an expected call of Expect.
func (mr *MockExpecterMockRecorder) Expect(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Expect", reflect.TypeOf((*MockExpecter)(nil).Expect), arg0, arg1)
}

// ExpectBatch mocks base method.
func (m *MockExpecter) ExpectBatch(arg0 []goexpect.Batcher, arg1 time.Duration) ([]goexpect.BatchRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpectBatch", arg0, arg1)
	ret0, _ := ret[0].([]goexpect.BatchRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExpectBatch indicates an expected call of ExpectBatch.
func (mr *MockExpecterMockRecorder) ExpectBatch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpectBatch", reflect.TypeOf((*MockExpecter)(nil).ExpectBatch), arg0, arg1)
}

// ExpectSwitchCase mocks base method.
func (m *MockExpecter) ExpectSwitchCase(arg0 []goexpect.Caser, arg1 time.Duration) (string, []string, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpectSwitchCase", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]string)
	ret2, _ := ret[2].(int)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ExpectSwitchCase indicates an expected call of ExpectSwitchCase.
func (mr *MockExpecterMockRecorder) ExpectSwitchCase(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpectSwitchCase", reflect.TypeOf((*MockExpecter)(nil).ExpectSwitchCase), arg0, arg1)
}

// Send mocks base method.
func (m *MockExpecter) Send(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockExpecterMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockExpecter)(nil).Send), arg0)
}