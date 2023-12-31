// Code generated by MockGen. DO NOT EDIT.
// Source: baseapp/abci_utils.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	math "cosmossdk.io/math"
	crypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
	baseapp "github.com/cosmos/cosmos-sdk/baseapp"
	types "github.com/cosmos/cosmos-sdk/crypto/types"
	types0 "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
)

// MockValidator is a mock of Validator interface.
type MockValidator struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorMockRecorder
}

// MockValidatorMockRecorder is the mock recorder for MockValidator.
type MockValidatorMockRecorder struct {
	mock *MockValidator
}

// NewMockValidator creates a new mock instance.
func NewMockValidator(ctrl *gomock.Controller) *MockValidator {
	mock := &MockValidator{ctrl: ctrl}
	mock.recorder = &MockValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidator) EXPECT() *MockValidatorMockRecorder {
	return m.recorder
}

// BondedTokens mocks base method.
func (m *MockValidator) BondedTokens() math.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BondedTokens")
	ret0, _ := ret[0].(math.Int)
	return ret0
}

// BondedTokens indicates an expected call of BondedTokens.
func (mr *MockValidatorMockRecorder) BondedTokens() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BondedTokens", reflect.TypeOf((*MockValidator)(nil).BondedTokens))
}

// CmtConsPublicKey mocks base method.
func (m *MockValidator) CmtConsPublicKey() (crypto.PublicKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CmtConsPublicKey")
	ret0, _ := ret[0].(crypto.PublicKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CmtConsPublicKey indicates an expected call of CmtConsPublicKey.
func (mr *MockValidatorMockRecorder) CmtConsPublicKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CmtConsPublicKey", reflect.TypeOf((*MockValidator)(nil).CmtConsPublicKey))
}

// MockValidatorStore is a mock of ValidatorStore interface.
type MockValidatorStore struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorStoreMockRecorder
}

// MockValidatorStoreMockRecorder is the mock recorder for MockValidatorStore.
type MockValidatorStoreMockRecorder struct {
	mock *MockValidatorStore
}

// NewMockValidatorStore creates a new mock instance.
func NewMockValidatorStore(ctrl *gomock.Controller) *MockValidatorStore {
	mock := &MockValidatorStore{ctrl: ctrl}
	mock.recorder = &MockValidatorStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidatorStore) EXPECT() *MockValidatorStoreMockRecorder {
	return m.recorder
}

// GetValidatorByConsAddr mocks base method.
func (m *MockValidatorStore) GetValidatorByConsAddr(arg0 types0.Context, arg1 types.Address) (baseapp.Validator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorByConsAddr", arg0, arg1)
	ret0, _ := ret[0].(baseapp.Validator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidatorByConsAddr indicates an expected call of GetValidatorByConsAddr.
func (mr *MockValidatorStoreMockRecorder) GetValidatorByConsAddr(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorByConsAddr", reflect.TypeOf((*MockValidatorStore)(nil).GetValidatorByConsAddr), arg0, arg1)
}

// TotalBondedTokens mocks base method.
func (m *MockValidatorStore) TotalBondedTokens(ctx types0.Context) math.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TotalBondedTokens", ctx)
	ret0, _ := ret[0].(math.Int)
	return ret0
}

// TotalBondedTokens indicates an expected call of TotalBondedTokens.
func (mr *MockValidatorStoreMockRecorder) TotalBondedTokens(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TotalBondedTokens", reflect.TypeOf((*MockValidatorStore)(nil).TotalBondedTokens), ctx)
}

// MockGasTx is a mock of GasTx interface.
type MockGasTx struct {
	ctrl     *gomock.Controller
	recorder *MockGasTxMockRecorder
}

// MockGasTxMockRecorder is the mock recorder for MockGasTx.
type MockGasTxMockRecorder struct {
	mock *MockGasTx
}

// NewMockGasTx creates a new mock instance.
func NewMockGasTx(ctrl *gomock.Controller) *MockGasTx {
	mock := &MockGasTx{ctrl: ctrl}
	mock.recorder = &MockGasTxMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGasTx) EXPECT() *MockGasTxMockRecorder {
	return m.recorder
}

// GetGas mocks base method.
func (m *MockGasTx) GetGas() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGas")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetGas indicates an expected call of GetGas.
func (mr *MockGasTxMockRecorder) GetGas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGas", reflect.TypeOf((*MockGasTx)(nil).GetGas))
}

// MockProposalTxVerifier is a mock of ProposalTxVerifier interface.
type MockProposalTxVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockProposalTxVerifierMockRecorder
}

// MockProposalTxVerifierMockRecorder is the mock recorder for MockProposalTxVerifier.
type MockProposalTxVerifierMockRecorder struct {
	mock *MockProposalTxVerifier
}

// NewMockProposalTxVerifier creates a new mock instance.
func NewMockProposalTxVerifier(ctrl *gomock.Controller) *MockProposalTxVerifier {
	mock := &MockProposalTxVerifier{ctrl: ctrl}
	mock.recorder = &MockProposalTxVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProposalTxVerifier) EXPECT() *MockProposalTxVerifierMockRecorder {
	return m.recorder
}

// PrepareProposalVerifyTx mocks base method.
func (m *MockProposalTxVerifier) PrepareProposalVerifyTx(tx types0.Tx) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareProposalVerifyTx", tx)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareProposalVerifyTx indicates an expected call of PrepareProposalVerifyTx.
func (mr *MockProposalTxVerifierMockRecorder) PrepareProposalVerifyTx(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareProposalVerifyTx", reflect.TypeOf((*MockProposalTxVerifier)(nil).PrepareProposalVerifyTx), tx)
}

// ProcessProposalVerifyTx mocks base method.
func (m *MockProposalTxVerifier) ProcessProposalVerifyTx(txBz []byte) (types0.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessProposalVerifyTx", txBz)
	ret0, _ := ret[0].(types0.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessProposalVerifyTx indicates an expected call of ProcessProposalVerifyTx.
func (mr *MockProposalTxVerifierMockRecorder) ProcessProposalVerifyTx(txBz interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessProposalVerifyTx", reflect.TypeOf((*MockProposalTxVerifier)(nil).ProcessProposalVerifyTx), txBz)
}
