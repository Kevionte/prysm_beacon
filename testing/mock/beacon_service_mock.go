// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Kevionte/prysm_beacon/v2/proto/prysm/v1alpha1 (interfaces: BeaconChainClient)
//
// Generated by this command:
//
//	mockgen -package=mock -destination=testing/mock/beacon_service_mock.go github.com/Kevionte/prysm_beacon/v2/proto/prysm/v1alpha1 BeaconChainClient
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	eth "github.com/Kevionte/prysm_beacon/v2/proto/prysm/v1alpha1"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockBeaconChainClient is a mock of BeaconChainClient interface.
type MockBeaconChainClient struct {
	ctrl     *gomock.Controller
	recorder *MockBeaconChainClientMockRecorder
}

// MockBeaconChainClientMockRecorder is the mock recorder for MockBeaconChainClient.
type MockBeaconChainClientMockRecorder struct {
	mock *MockBeaconChainClient
}

// NewMockBeaconChainClient creates a new mock instance.
func NewMockBeaconChainClient(ctrl *gomock.Controller) *MockBeaconChainClient {
	mock := &MockBeaconChainClient{ctrl: ctrl}
	mock.recorder = &MockBeaconChainClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBeaconChainClient) EXPECT() *MockBeaconChainClientMockRecorder {
	return m.recorder
}

// AttestationPool mocks base method.
func (m *MockBeaconChainClient) AttestationPool(arg0 context.Context, arg1 *eth.AttestationPoolRequest, arg2 ...grpc.CallOption) (*eth.AttestationPoolResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AttestationPool", varargs...)
	ret0, _ := ret[0].(*eth.AttestationPoolResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttestationPool indicates an expected call of AttestationPool.
func (mr *MockBeaconChainClientMockRecorder) AttestationPool(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttestationPool", reflect.TypeOf((*MockBeaconChainClient)(nil).AttestationPool), varargs...)
}

// GetBeaconConfig mocks base method.
func (m *MockBeaconChainClient) GetBeaconConfig(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*eth.BeaconConfig, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBeaconConfig", varargs...)
	ret0, _ := ret[0].(*eth.BeaconConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBeaconConfig indicates an expected call of GetBeaconConfig.
func (mr *MockBeaconChainClientMockRecorder) GetBeaconConfig(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBeaconConfig", reflect.TypeOf((*MockBeaconChainClient)(nil).GetBeaconConfig), varargs...)
}

// GetChainHead mocks base method.
func (m *MockBeaconChainClient) GetChainHead(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*eth.ChainHead, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetChainHead", varargs...)
	ret0, _ := ret[0].(*eth.ChainHead)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChainHead indicates an expected call of GetChainHead.
func (mr *MockBeaconChainClientMockRecorder) GetChainHead(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainHead", reflect.TypeOf((*MockBeaconChainClient)(nil).GetChainHead), varargs...)
}

// GetIndividualVotes mocks base method.
func (m *MockBeaconChainClient) GetIndividualVotes(arg0 context.Context, arg1 *eth.IndividualVotesRequest, arg2 ...grpc.CallOption) (*eth.IndividualVotesRespond, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIndividualVotes", varargs...)
	ret0, _ := ret[0].(*eth.IndividualVotesRespond)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIndividualVotes indicates an expected call of GetIndividualVotes.
func (mr *MockBeaconChainClientMockRecorder) GetIndividualVotes(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIndividualVotes", reflect.TypeOf((*MockBeaconChainClient)(nil).GetIndividualVotes), varargs...)
}

// GetValidator mocks base method.
func (m *MockBeaconChainClient) GetValidator(arg0 context.Context, arg1 *eth.GetValidatorRequest, arg2 ...grpc.CallOption) (*eth.Validator, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetValidator", varargs...)
	ret0, _ := ret[0].(*eth.Validator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidator indicates an expected call of GetValidator.
func (mr *MockBeaconChainClientMockRecorder) GetValidator(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidator", reflect.TypeOf((*MockBeaconChainClient)(nil).GetValidator), varargs...)
}

// GetValidatorActiveSetChanges mocks base method.
func (m *MockBeaconChainClient) GetValidatorActiveSetChanges(arg0 context.Context, arg1 *eth.GetValidatorActiveSetChangesRequest, arg2 ...grpc.CallOption) (*eth.ActiveSetChanges, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetValidatorActiveSetChanges", varargs...)
	ret0, _ := ret[0].(*eth.ActiveSetChanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidatorActiveSetChanges indicates an expected call of GetValidatorActiveSetChanges.
func (mr *MockBeaconChainClientMockRecorder) GetValidatorActiveSetChanges(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorActiveSetChanges", reflect.TypeOf((*MockBeaconChainClient)(nil).GetValidatorActiveSetChanges), varargs...)
}

// GetValidatorParticipation mocks base method.
func (m *MockBeaconChainClient) GetValidatorParticipation(arg0 context.Context, arg1 *eth.GetValidatorParticipationRequest, arg2 ...grpc.CallOption) (*eth.ValidatorParticipationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetValidatorParticipation", varargs...)
	ret0, _ := ret[0].(*eth.ValidatorParticipationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidatorParticipation indicates an expected call of GetValidatorParticipation.
func (mr *MockBeaconChainClientMockRecorder) GetValidatorParticipation(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorParticipation", reflect.TypeOf((*MockBeaconChainClient)(nil).GetValidatorParticipation), varargs...)
}

// GetValidatorPerformance mocks base method.
func (m *MockBeaconChainClient) GetValidatorPerformance(arg0 context.Context, arg1 *eth.ValidatorPerformanceRequest, arg2 ...grpc.CallOption) (*eth.ValidatorPerformanceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetValidatorPerformance", varargs...)
	ret0, _ := ret[0].(*eth.ValidatorPerformanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidatorPerformance indicates an expected call of GetValidatorPerformance.
func (mr *MockBeaconChainClientMockRecorder) GetValidatorPerformance(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorPerformance", reflect.TypeOf((*MockBeaconChainClient)(nil).GetValidatorPerformance), varargs...)
}

// GetValidatorQueue mocks base method.
func (m *MockBeaconChainClient) GetValidatorQueue(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*eth.ValidatorQueue, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetValidatorQueue", varargs...)
	ret0, _ := ret[0].(*eth.ValidatorQueue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidatorQueue indicates an expected call of GetValidatorQueue.
func (mr *MockBeaconChainClientMockRecorder) GetValidatorQueue(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorQueue", reflect.TypeOf((*MockBeaconChainClient)(nil).GetValidatorQueue), varargs...)
}

// ListAttestations mocks base method.
func (m *MockBeaconChainClient) ListAttestations(arg0 context.Context, arg1 *eth.ListAttestationsRequest, arg2 ...grpc.CallOption) (*eth.ListAttestationsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAttestations", varargs...)
	ret0, _ := ret[0].(*eth.ListAttestationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAttestations indicates an expected call of ListAttestations.
func (mr *MockBeaconChainClientMockRecorder) ListAttestations(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAttestations", reflect.TypeOf((*MockBeaconChainClient)(nil).ListAttestations), varargs...)
}

// ListBeaconBlocks mocks base method.
func (m *MockBeaconChainClient) ListBeaconBlocks(arg0 context.Context, arg1 *eth.ListBlocksRequest, arg2 ...grpc.CallOption) (*eth.ListBeaconBlocksResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBeaconBlocks", varargs...)
	ret0, _ := ret[0].(*eth.ListBeaconBlocksResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBeaconBlocks indicates an expected call of ListBeaconBlocks.
func (mr *MockBeaconChainClientMockRecorder) ListBeaconBlocks(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBeaconBlocks", reflect.TypeOf((*MockBeaconChainClient)(nil).ListBeaconBlocks), varargs...)
}

// ListBeaconCommittees mocks base method.
func (m *MockBeaconChainClient) ListBeaconCommittees(arg0 context.Context, arg1 *eth.ListCommitteesRequest, arg2 ...grpc.CallOption) (*eth.BeaconCommittees, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBeaconCommittees", varargs...)
	ret0, _ := ret[0].(*eth.BeaconCommittees)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBeaconCommittees indicates an expected call of ListBeaconCommittees.
func (mr *MockBeaconChainClientMockRecorder) ListBeaconCommittees(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBeaconCommittees", reflect.TypeOf((*MockBeaconChainClient)(nil).ListBeaconCommittees), varargs...)
}

// ListIndexedAttestations mocks base method.
func (m *MockBeaconChainClient) ListIndexedAttestations(arg0 context.Context, arg1 *eth.ListIndexedAttestationsRequest, arg2 ...grpc.CallOption) (*eth.ListIndexedAttestationsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIndexedAttestations", varargs...)
	ret0, _ := ret[0].(*eth.ListIndexedAttestationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIndexedAttestations indicates an expected call of ListIndexedAttestations.
func (mr *MockBeaconChainClientMockRecorder) ListIndexedAttestations(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIndexedAttestations", reflect.TypeOf((*MockBeaconChainClient)(nil).ListIndexedAttestations), varargs...)
}

// ListValidatorAssignments mocks base method.
func (m *MockBeaconChainClient) ListValidatorAssignments(arg0 context.Context, arg1 *eth.ListValidatorAssignmentsRequest, arg2 ...grpc.CallOption) (*eth.ValidatorAssignments, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListValidatorAssignments", varargs...)
	ret0, _ := ret[0].(*eth.ValidatorAssignments)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListValidatorAssignments indicates an expected call of ListValidatorAssignments.
func (mr *MockBeaconChainClientMockRecorder) ListValidatorAssignments(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListValidatorAssignments", reflect.TypeOf((*MockBeaconChainClient)(nil).ListValidatorAssignments), varargs...)
}

// ListValidatorBalances mocks base method.
func (m *MockBeaconChainClient) ListValidatorBalances(arg0 context.Context, arg1 *eth.ListValidatorBalancesRequest, arg2 ...grpc.CallOption) (*eth.ValidatorBalances, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListValidatorBalances", varargs...)
	ret0, _ := ret[0].(*eth.ValidatorBalances)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListValidatorBalances indicates an expected call of ListValidatorBalances.
func (mr *MockBeaconChainClientMockRecorder) ListValidatorBalances(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListValidatorBalances", reflect.TypeOf((*MockBeaconChainClient)(nil).ListValidatorBalances), varargs...)
}

// ListValidators mocks base method.
func (m *MockBeaconChainClient) ListValidators(arg0 context.Context, arg1 *eth.ListValidatorsRequest, arg2 ...grpc.CallOption) (*eth.Validators, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListValidators", varargs...)
	ret0, _ := ret[0].(*eth.Validators)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListValidators indicates an expected call of ListValidators.
func (mr *MockBeaconChainClientMockRecorder) ListValidators(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListValidators", reflect.TypeOf((*MockBeaconChainClient)(nil).ListValidators), varargs...)
}

// SubmitAttesterSlashing mocks base method.
func (m *MockBeaconChainClient) SubmitAttesterSlashing(arg0 context.Context, arg1 *eth.AttesterSlashing, arg2 ...grpc.CallOption) (*eth.SubmitSlashingResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubmitAttesterSlashing", varargs...)
	ret0, _ := ret[0].(*eth.SubmitSlashingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitAttesterSlashing indicates an expected call of SubmitAttesterSlashing.
func (mr *MockBeaconChainClientMockRecorder) SubmitAttesterSlashing(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitAttesterSlashing", reflect.TypeOf((*MockBeaconChainClient)(nil).SubmitAttesterSlashing), varargs...)
}

// SubmitProposerSlashing mocks base method.
func (m *MockBeaconChainClient) SubmitProposerSlashing(arg0 context.Context, arg1 *eth.ProposerSlashing, arg2 ...grpc.CallOption) (*eth.SubmitSlashingResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubmitProposerSlashing", varargs...)
	ret0, _ := ret[0].(*eth.SubmitSlashingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitProposerSlashing indicates an expected call of SubmitProposerSlashing.
func (mr *MockBeaconChainClientMockRecorder) SubmitProposerSlashing(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitProposerSlashing", reflect.TypeOf((*MockBeaconChainClient)(nil).SubmitProposerSlashing), varargs...)
}
