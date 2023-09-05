// Code generated by MockGen. DO NOT EDIT.
// Source: types/module/module.go

// Package mock is a generated GoMock package.
package mock

import (
	json "encoding/json"
	reflect "reflect"

	types "github.com/cometbft/cometbft/abci/types"
	client "github.com/cosmos/cosmos-sdk/client"
	codec "github.com/cosmos/cosmos-sdk/codec"
	types0 "github.com/cosmos/cosmos-sdk/codec/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	module "github.com/cosmos/cosmos-sdk/types/module"
	gomock "github.com/golang/mock/gomock"
	runtime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	cobra "github.com/spf13/cobra"
)

// MockAppModuleBasic is a mock of AppModuleBasic interface.
type MockAppModuleBasic struct {
	ctrl     *gomock.Controller
	recorder *MockAppModuleBasicMockRecorder
}

// MockAppModuleBasicMockRecorder is the mock recorder for MockAppModuleBasic.
type MockAppModuleBasicMockRecorder struct {
	mock *MockAppModuleBasic
}

// NewMockAppModuleBasic creates a new mock instance.
func NewMockAppModuleBasic(ctrl *gomock.Controller) *MockAppModuleBasic {
	mock := &MockAppModuleBasic{ctrl: ctrl}
	mock.recorder = &MockAppModuleBasicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppModuleBasic) EXPECT() *MockAppModuleBasicMockRecorder {
	return m.recorder
}

// GetQueryCmd mocks base method.
func (m *MockAppModuleBasic) GetQueryCmd() *cobra.Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueryCmd")
	ret0, _ := ret[0].(*cobra.Command)
	return ret0
}

// GetQueryCmd indicates an expected call of GetQueryCmd.
func (mr *MockAppModuleBasicMockRecorder) GetQueryCmd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryCmd", reflect.TypeOf((*MockAppModuleBasic)(nil).GetQueryCmd))
}

// GetTxCmd mocks base method.
func (m *MockAppModuleBasic) GetTxCmd() *cobra.Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxCmd")
	ret0, _ := ret[0].(*cobra.Command)
	return ret0
}

// GetTxCmd indicates an expected call of GetTxCmd.
func (mr *MockAppModuleBasicMockRecorder) GetTxCmd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxCmd", reflect.TypeOf((*MockAppModuleBasic)(nil).GetTxCmd))
}

// Name mocks base method.
func (m *MockAppModuleBasic) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockAppModuleBasicMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockAppModuleBasic)(nil).Name))
}

// RegisterGRPCGatewayRoutes mocks base method.
func (m *MockAppModuleBasic) RegisterGRPCGatewayRoutes(arg0 client.Context, arg1 *runtime.ServeMux) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterGRPCGatewayRoutes", arg0, arg1)
}

// RegisterGRPCGatewayRoutes indicates an expected call of RegisterGRPCGatewayRoutes.
func (mr *MockAppModuleBasicMockRecorder) RegisterGRPCGatewayRoutes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterGRPCGatewayRoutes", reflect.TypeOf((*MockAppModuleBasic)(nil).RegisterGRPCGatewayRoutes), arg0, arg1)
}

// RegisterInterfaces mocks base method.
func (m *MockAppModuleBasic) RegisterInterfaces(arg0 types0.InterfaceRegistry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterInterfaces", arg0)
}

// RegisterInterfaces indicates an expected call of RegisterInterfaces.
func (mr *MockAppModuleBasicMockRecorder) RegisterInterfaces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInterfaces", reflect.TypeOf((*MockAppModuleBasic)(nil).RegisterInterfaces), arg0)
}

// RegisterLegacyAminoCodec mocks base method.
func (m *MockAppModuleBasic) RegisterLegacyAminoCodec(arg0 *codec.LegacyAmino) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterLegacyAminoCodec", arg0)
}

// RegisterLegacyAminoCodec indicates an expected call of RegisterLegacyAminoCodec.
func (mr *MockAppModuleBasicMockRecorder) RegisterLegacyAminoCodec(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterLegacyAminoCodec", reflect.TypeOf((*MockAppModuleBasic)(nil).RegisterLegacyAminoCodec), arg0)
}

// MockHasName is a mock of HasName interface.
type MockHasName struct {
	ctrl     *gomock.Controller
	recorder *MockHasNameMockRecorder
}

// MockHasNameMockRecorder is the mock recorder for MockHasName.
type MockHasNameMockRecorder struct {
	mock *MockHasName
}

// NewMockHasName creates a new mock instance.
func NewMockHasName(ctrl *gomock.Controller) *MockHasName {
	mock := &MockHasName{ctrl: ctrl}
	mock.recorder = &MockHasNameMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHasName) EXPECT() *MockHasNameMockRecorder {
	return m.recorder
}

// Name mocks base method.
func (m *MockHasName) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockHasNameMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockHasName)(nil).Name))
}

// MockHasGenesisBasics is a mock of HasGenesisBasics interface.
type MockHasGenesisBasics struct {
	ctrl     *gomock.Controller
	recorder *MockHasGenesisBasicsMockRecorder
}

// MockHasGenesisBasicsMockRecorder is the mock recorder for MockHasGenesisBasics.
type MockHasGenesisBasicsMockRecorder struct {
	mock *MockHasGenesisBasics
}

// NewMockHasGenesisBasics creates a new mock instance.
func NewMockHasGenesisBasics(ctrl *gomock.Controller) *MockHasGenesisBasics {
	mock := &MockHasGenesisBasics{ctrl: ctrl}
	mock.recorder = &MockHasGenesisBasicsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHasGenesisBasics) EXPECT() *MockHasGenesisBasicsMockRecorder {
	return m.recorder
}

// DefaultGenesis mocks base method.
func (m *MockHasGenesisBasics) DefaultGenesis(arg0 codec.JSONCodec) json.RawMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultGenesis", arg0)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}

// DefaultGenesis indicates an expected call of DefaultGenesis.
func (mr *MockHasGenesisBasicsMockRecorder) DefaultGenesis(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultGenesis", reflect.TypeOf((*MockHasGenesisBasics)(nil).DefaultGenesis), arg0)
}

// ValidateGenesis mocks base method.
func (m *MockHasGenesisBasics) ValidateGenesis(arg0 codec.JSONCodec, arg1 client.TxEncodingConfig, arg2 json.RawMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateGenesis", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateGenesis indicates an expected call of ValidateGenesis.
func (mr *MockHasGenesisBasicsMockRecorder) ValidateGenesis(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateGenesis", reflect.TypeOf((*MockHasGenesisBasics)(nil).ValidateGenesis), arg0, arg1, arg2)
}

// MockAppModuleGenesis is a mock of AppModuleGenesis interface.
type MockAppModuleGenesis struct {
	ctrl     *gomock.Controller
	recorder *MockAppModuleGenesisMockRecorder
}

// MockAppModuleGenesisMockRecorder is the mock recorder for MockAppModuleGenesis.
type MockAppModuleGenesisMockRecorder struct {
	mock *MockAppModuleGenesis
}

// NewMockAppModuleGenesis creates a new mock instance.
func NewMockAppModuleGenesis(ctrl *gomock.Controller) *MockAppModuleGenesis {
	mock := &MockAppModuleGenesis{ctrl: ctrl}
	mock.recorder = &MockAppModuleGenesisMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppModuleGenesis) EXPECT() *MockAppModuleGenesisMockRecorder {
	return m.recorder
}

// DefaultGenesis mocks base method.
func (m *MockAppModuleGenesis) DefaultGenesis(arg0 codec.JSONCodec) json.RawMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultGenesis", arg0)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}

// DefaultGenesis indicates an expected call of DefaultGenesis.
func (mr *MockAppModuleGenesisMockRecorder) DefaultGenesis(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultGenesis", reflect.TypeOf((*MockAppModuleGenesis)(nil).DefaultGenesis), arg0)
}

// ExportGenesis mocks base method.
func (m *MockAppModuleGenesis) ExportGenesis(arg0 types1.Context, arg1 codec.JSONCodec) json.RawMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportGenesis", arg0, arg1)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}

// ExportGenesis indicates an expected call of ExportGenesis.
func (mr *MockAppModuleGenesisMockRecorder) ExportGenesis(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportGenesis", reflect.TypeOf((*MockAppModuleGenesis)(nil).ExportGenesis), arg0, arg1)
}

// GetQueryCmd mocks base method.
func (m *MockAppModuleGenesis) GetQueryCmd() *cobra.Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueryCmd")
	ret0, _ := ret[0].(*cobra.Command)
	return ret0
}

// GetQueryCmd indicates an expected call of GetQueryCmd.
func (mr *MockAppModuleGenesisMockRecorder) GetQueryCmd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryCmd", reflect.TypeOf((*MockAppModuleGenesis)(nil).GetQueryCmd))
}

// GetTxCmd mocks base method.
func (m *MockAppModuleGenesis) GetTxCmd() *cobra.Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxCmd")
	ret0, _ := ret[0].(*cobra.Command)
	return ret0
}

// GetTxCmd indicates an expected call of GetTxCmd.
func (mr *MockAppModuleGenesisMockRecorder) GetTxCmd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxCmd", reflect.TypeOf((*MockAppModuleGenesis)(nil).GetTxCmd))
}

// InitGenesis mocks base method.
func (m *MockAppModuleGenesis) InitGenesis(arg0 types1.Context, arg1 codec.JSONCodec, arg2 json.RawMessage) []types.ValidatorUpdate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitGenesis", arg0, arg1, arg2)
	ret0, _ := ret[0].([]types.ValidatorUpdate)
	return ret0
}

// InitGenesis indicates an expected call of InitGenesis.
func (mr *MockAppModuleGenesisMockRecorder) InitGenesis(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitGenesis", reflect.TypeOf((*MockAppModuleGenesis)(nil).InitGenesis), arg0, arg1, arg2)
}

// Name mocks base method.
func (m *MockAppModuleGenesis) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockAppModuleGenesisMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockAppModuleGenesis)(nil).Name))
}

// RegisterGRPCGatewayRoutes mocks base method.
func (m *MockAppModuleGenesis) RegisterGRPCGatewayRoutes(arg0 client.Context, arg1 *runtime.ServeMux) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterGRPCGatewayRoutes", arg0, arg1)
}

// RegisterGRPCGatewayRoutes indicates an expected call of RegisterGRPCGatewayRoutes.
func (mr *MockAppModuleGenesisMockRecorder) RegisterGRPCGatewayRoutes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterGRPCGatewayRoutes", reflect.TypeOf((*MockAppModuleGenesis)(nil).RegisterGRPCGatewayRoutes), arg0, arg1)
}

// RegisterInterfaces mocks base method.
func (m *MockAppModuleGenesis) RegisterInterfaces(arg0 types0.InterfaceRegistry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterInterfaces", arg0)
}

// RegisterInterfaces indicates an expected call of RegisterInterfaces.
func (mr *MockAppModuleGenesisMockRecorder) RegisterInterfaces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInterfaces", reflect.TypeOf((*MockAppModuleGenesis)(nil).RegisterInterfaces), arg0)
}

// RegisterLegacyAminoCodec mocks base method.
func (m *MockAppModuleGenesis) RegisterLegacyAminoCodec(arg0 *codec.LegacyAmino) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterLegacyAminoCodec", arg0)
}

// RegisterLegacyAminoCodec indicates an expected call of RegisterLegacyAminoCodec.
func (mr *MockAppModuleGenesisMockRecorder) RegisterLegacyAminoCodec(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterLegacyAminoCodec", reflect.TypeOf((*MockAppModuleGenesis)(nil).RegisterLegacyAminoCodec), arg0)
}

// ValidateGenesis mocks base method.
func (m *MockAppModuleGenesis) ValidateGenesis(arg0 codec.JSONCodec, arg1 client.TxEncodingConfig, arg2 json.RawMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateGenesis", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateGenesis indicates an expected call of ValidateGenesis.
func (mr *MockAppModuleGenesisMockRecorder) ValidateGenesis(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateGenesis", reflect.TypeOf((*MockAppModuleGenesis)(nil).ValidateGenesis), arg0, arg1, arg2)
}

// MockHasGenesis is a mock of HasGenesis interface.
type MockHasGenesis struct {
	ctrl     *gomock.Controller
	recorder *MockHasGenesisMockRecorder
}

// MockHasGenesisMockRecorder is the mock recorder for MockHasGenesis.
type MockHasGenesisMockRecorder struct {
	mock *MockHasGenesis
}

// NewMockHasGenesis creates a new mock instance.
func NewMockHasGenesis(ctrl *gomock.Controller) *MockHasGenesis {
	mock := &MockHasGenesis{ctrl: ctrl}
	mock.recorder = &MockHasGenesisMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHasGenesis) EXPECT() *MockHasGenesisMockRecorder {
	return m.recorder
}

// DefaultGenesis mocks base method.
func (m *MockHasGenesis) DefaultGenesis(arg0 codec.JSONCodec) json.RawMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultGenesis", arg0)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}

// DefaultGenesis indicates an expected call of DefaultGenesis.
func (mr *MockHasGenesisMockRecorder) DefaultGenesis(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultGenesis", reflect.TypeOf((*MockHasGenesis)(nil).DefaultGenesis), arg0)
}

// ExportGenesis mocks base method.
func (m *MockHasGenesis) ExportGenesis(arg0 types1.Context, arg1 codec.JSONCodec) json.RawMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportGenesis", arg0, arg1)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}

// ExportGenesis indicates an expected call of ExportGenesis.
func (mr *MockHasGenesisMockRecorder) ExportGenesis(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportGenesis", reflect.TypeOf((*MockHasGenesis)(nil).ExportGenesis), arg0, arg1)
}

// InitGenesis mocks base method.
func (m *MockHasGenesis) InitGenesis(arg0 types1.Context, arg1 codec.JSONCodec, arg2 json.RawMessage) []types.ValidatorUpdate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitGenesis", arg0, arg1, arg2)
	ret0, _ := ret[0].([]types.ValidatorUpdate)
	return ret0
}

// InitGenesis indicates an expected call of InitGenesis.
func (mr *MockHasGenesisMockRecorder) InitGenesis(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitGenesis", reflect.TypeOf((*MockHasGenesis)(nil).InitGenesis), arg0, arg1, arg2)
}

// ValidateGenesis mocks base method.
func (m *MockHasGenesis) ValidateGenesis(arg0 codec.JSONCodec, arg1 client.TxEncodingConfig, arg2 json.RawMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateGenesis", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateGenesis indicates an expected call of ValidateGenesis.
func (mr *MockHasGenesisMockRecorder) ValidateGenesis(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateGenesis", reflect.TypeOf((*MockHasGenesis)(nil).ValidateGenesis), arg0, arg1, arg2)
}

// MockAppModule is a mock of AppModule interface.
type MockAppModule struct {
	ctrl     *gomock.Controller
	recorder *MockAppModuleMockRecorder
}

// MockAppModuleMockRecorder is the mock recorder for MockAppModule.
type MockAppModuleMockRecorder struct {
	mock *MockAppModule
}

// NewMockAppModule creates a new mock instance.
func NewMockAppModule(ctrl *gomock.Controller) *MockAppModule {
	mock := &MockAppModule{ctrl: ctrl}
	mock.recorder = &MockAppModuleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppModule) EXPECT() *MockAppModuleMockRecorder {
	return m.recorder
}

// GetQueryCmd mocks base method.
func (m *MockAppModule) GetQueryCmd() *cobra.Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueryCmd")
	ret0, _ := ret[0].(*cobra.Command)
	return ret0
}

// GetQueryCmd indicates an expected call of GetQueryCmd.
func (mr *MockAppModuleMockRecorder) GetQueryCmd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryCmd", reflect.TypeOf((*MockAppModule)(nil).GetQueryCmd))
}

// GetTxCmd mocks base method.
func (m *MockAppModule) GetTxCmd() *cobra.Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxCmd")
	ret0, _ := ret[0].(*cobra.Command)
	return ret0
}

// GetTxCmd indicates an expected call of GetTxCmd.
func (mr *MockAppModuleMockRecorder) GetTxCmd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxCmd", reflect.TypeOf((*MockAppModule)(nil).GetTxCmd))
}

// Name mocks base method.
func (m *MockAppModule) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockAppModuleMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockAppModule)(nil).Name))
}

// RegisterGRPCGatewayRoutes mocks base method.
func (m *MockAppModule) RegisterGRPCGatewayRoutes(arg0 client.Context, arg1 *runtime.ServeMux) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterGRPCGatewayRoutes", arg0, arg1)
}

// RegisterGRPCGatewayRoutes indicates an expected call of RegisterGRPCGatewayRoutes.
func (mr *MockAppModuleMockRecorder) RegisterGRPCGatewayRoutes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterGRPCGatewayRoutes", reflect.TypeOf((*MockAppModule)(nil).RegisterGRPCGatewayRoutes), arg0, arg1)
}

// RegisterInterfaces mocks base method.
func (m *MockAppModule) RegisterInterfaces(arg0 types0.InterfaceRegistry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterInterfaces", arg0)
}

// RegisterInterfaces indicates an expected call of RegisterInterfaces.
func (mr *MockAppModuleMockRecorder) RegisterInterfaces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInterfaces", reflect.TypeOf((*MockAppModule)(nil).RegisterInterfaces), arg0)
}

// RegisterLegacyAminoCodec mocks base method.
func (m *MockAppModule) RegisterLegacyAminoCodec(arg0 *codec.LegacyAmino) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterLegacyAminoCodec", arg0)
}

// RegisterLegacyAminoCodec indicates an expected call of RegisterLegacyAminoCodec.
func (mr *MockAppModuleMockRecorder) RegisterLegacyAminoCodec(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterLegacyAminoCodec", reflect.TypeOf((*MockAppModule)(nil).RegisterLegacyAminoCodec), arg0)
}

// MockHasInvariants is a mock of HasInvariants interface.
type MockHasInvariants struct {
	ctrl     *gomock.Controller
	recorder *MockHasInvariantsMockRecorder
}

// MockHasInvariantsMockRecorder is the mock recorder for MockHasInvariants.
type MockHasInvariantsMockRecorder struct {
	mock *MockHasInvariants
}

// NewMockHasInvariants creates a new mock instance.
func NewMockHasInvariants(ctrl *gomock.Controller) *MockHasInvariants {
	mock := &MockHasInvariants{ctrl: ctrl}
	mock.recorder = &MockHasInvariantsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHasInvariants) EXPECT() *MockHasInvariantsMockRecorder {
	return m.recorder
}

// RegisterInvariants mocks base method.
func (m *MockHasInvariants) RegisterInvariants(arg0 types1.InvariantRegistry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterInvariants", arg0)
}

// RegisterInvariants indicates an expected call of RegisterInvariants.
func (mr *MockHasInvariantsMockRecorder) RegisterInvariants(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInvariants", reflect.TypeOf((*MockHasInvariants)(nil).RegisterInvariants), arg0)
}

// MockHasServices is a mock of HasServices interface.
type MockHasServices struct {
	ctrl     *gomock.Controller
	recorder *MockHasServicesMockRecorder
}

// MockHasServicesMockRecorder is the mock recorder for MockHasServices.
type MockHasServicesMockRecorder struct {
	mock *MockHasServices
}

// NewMockHasServices creates a new mock instance.
func NewMockHasServices(ctrl *gomock.Controller) *MockHasServices {
	mock := &MockHasServices{ctrl: ctrl}
	mock.recorder = &MockHasServicesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHasServices) EXPECT() *MockHasServicesMockRecorder {
	return m.recorder
}

// RegisterServices mocks base method.
func (m *MockHasServices) RegisterServices(arg0 module.Configurator) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterServices", arg0)
}

// RegisterServices indicates an expected call of RegisterServices.
func (mr *MockHasServicesMockRecorder) RegisterServices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterServices", reflect.TypeOf((*MockHasServices)(nil).RegisterServices), arg0)
}

// MockHasConsensusVersion is a mock of HasConsensusVersion interface.
type MockHasConsensusVersion struct {
	ctrl     *gomock.Controller
	recorder *MockHasConsensusVersionMockRecorder
}

// MockHasConsensusVersionMockRecorder is the mock recorder for MockHasConsensusVersion.
type MockHasConsensusVersionMockRecorder struct {
	mock *MockHasConsensusVersion
}

// NewMockHasConsensusVersion creates a new mock instance.
func NewMockHasConsensusVersion(ctrl *gomock.Controller) *MockHasConsensusVersion {
	mock := &MockHasConsensusVersion{ctrl: ctrl}
	mock.recorder = &MockHasConsensusVersionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHasConsensusVersion) EXPECT() *MockHasConsensusVersionMockRecorder {
	return m.recorder
}

// ConsensusVersion mocks base method.
func (m *MockHasConsensusVersion) ConsensusVersion() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsensusVersion")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// ConsensusVersion indicates an expected call of ConsensusVersion.
func (mr *MockHasConsensusVersionMockRecorder) ConsensusVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsensusVersion", reflect.TypeOf((*MockHasConsensusVersion)(nil).ConsensusVersion))
}

// MockBeginBlockAppModule is a mock of BeginBlockAppModule interface.
type MockBeginBlockAppModule struct {
	ctrl     *gomock.Controller
	recorder *MockBeginBlockAppModuleMockRecorder
}

// MockBeginBlockAppModuleMockRecorder is the mock recorder for MockBeginBlockAppModule.
type MockBeginBlockAppModuleMockRecorder struct {
	mock *MockBeginBlockAppModule
}

// NewMockBeginBlockAppModule creates a new mock instance.
func NewMockBeginBlockAppModule(ctrl *gomock.Controller) *MockBeginBlockAppModule {
	mock := &MockBeginBlockAppModule{ctrl: ctrl}
	mock.recorder = &MockBeginBlockAppModuleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBeginBlockAppModule) EXPECT() *MockBeginBlockAppModuleMockRecorder {
	return m.recorder
}

// BeginBlock mocks base method.
func (m *MockBeginBlockAppModule) BeginBlock(arg0 types1.Context, arg1 types.RequestBeginBlock) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BeginBlock", arg0, arg1)
}

// BeginBlock indicates an expected call of BeginBlock.
func (mr *MockBeginBlockAppModuleMockRecorder) BeginBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginBlock", reflect.TypeOf((*MockBeginBlockAppModule)(nil).BeginBlock), arg0, arg1)
}

// GetQueryCmd mocks base method.
func (m *MockBeginBlockAppModule) GetQueryCmd() *cobra.Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueryCmd")
	ret0, _ := ret[0].(*cobra.Command)
	return ret0
}

// GetQueryCmd indicates an expected call of GetQueryCmd.
func (mr *MockBeginBlockAppModuleMockRecorder) GetQueryCmd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryCmd", reflect.TypeOf((*MockBeginBlockAppModule)(nil).GetQueryCmd))
}

// GetTxCmd mocks base method.
func (m *MockBeginBlockAppModule) GetTxCmd() *cobra.Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxCmd")
	ret0, _ := ret[0].(*cobra.Command)
	return ret0
}

// GetTxCmd indicates an expected call of GetTxCmd.
func (mr *MockBeginBlockAppModuleMockRecorder) GetTxCmd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxCmd", reflect.TypeOf((*MockBeginBlockAppModule)(nil).GetTxCmd))
}

// Name mocks base method.
func (m *MockBeginBlockAppModule) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockBeginBlockAppModuleMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockBeginBlockAppModule)(nil).Name))
}

// RegisterGRPCGatewayRoutes mocks base method.
func (m *MockBeginBlockAppModule) RegisterGRPCGatewayRoutes(arg0 client.Context, arg1 *runtime.ServeMux) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterGRPCGatewayRoutes", arg0, arg1)
}

// RegisterGRPCGatewayRoutes indicates an expected call of RegisterGRPCGatewayRoutes.
func (mr *MockBeginBlockAppModuleMockRecorder) RegisterGRPCGatewayRoutes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterGRPCGatewayRoutes", reflect.TypeOf((*MockBeginBlockAppModule)(nil).RegisterGRPCGatewayRoutes), arg0, arg1)
}

// RegisterInterfaces mocks base method.
func (m *MockBeginBlockAppModule) RegisterInterfaces(arg0 types0.InterfaceRegistry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterInterfaces", arg0)
}

// RegisterInterfaces indicates an expected call of RegisterInterfaces.
func (mr *MockBeginBlockAppModuleMockRecorder) RegisterInterfaces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInterfaces", reflect.TypeOf((*MockBeginBlockAppModule)(nil).RegisterInterfaces), arg0)
}

// RegisterLegacyAminoCodec mocks base method.
func (m *MockBeginBlockAppModule) RegisterLegacyAminoCodec(arg0 *codec.LegacyAmino) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterLegacyAminoCodec", arg0)
}

// RegisterLegacyAminoCodec indicates an expected call of RegisterLegacyAminoCodec.
func (mr *MockBeginBlockAppModuleMockRecorder) RegisterLegacyAminoCodec(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterLegacyAminoCodec", reflect.TypeOf((*MockBeginBlockAppModule)(nil).RegisterLegacyAminoCodec), arg0)
}

// MockEndBlockAppModule is a mock of EndBlockAppModule interface.
type MockEndBlockAppModule struct {
	ctrl     *gomock.Controller
	recorder *MockEndBlockAppModuleMockRecorder
}

// MockEndBlockAppModuleMockRecorder is the mock recorder for MockEndBlockAppModule.
type MockEndBlockAppModuleMockRecorder struct {
	mock *MockEndBlockAppModule
}

// NewMockEndBlockAppModule creates a new mock instance.
func NewMockEndBlockAppModule(ctrl *gomock.Controller) *MockEndBlockAppModule {
	mock := &MockEndBlockAppModule{ctrl: ctrl}
	mock.recorder = &MockEndBlockAppModuleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEndBlockAppModule) EXPECT() *MockEndBlockAppModuleMockRecorder {
	return m.recorder
}

// EndBlock mocks base method.
func (m *MockEndBlockAppModule) EndBlock(arg0 types1.Context, arg1 types.RequestEndBlock) []types.ValidatorUpdate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndBlock", arg0, arg1)
	ret0, _ := ret[0].([]types.ValidatorUpdate)
	return ret0
}

// EndBlock indicates an expected call of EndBlock.
func (mr *MockEndBlockAppModuleMockRecorder) EndBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndBlock", reflect.TypeOf((*MockEndBlockAppModule)(nil).EndBlock), arg0, arg1)
}

// GetQueryCmd mocks base method.
func (m *MockEndBlockAppModule) GetQueryCmd() *cobra.Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueryCmd")
	ret0, _ := ret[0].(*cobra.Command)
	return ret0
}

// GetQueryCmd indicates an expected call of GetQueryCmd.
func (mr *MockEndBlockAppModuleMockRecorder) GetQueryCmd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryCmd", reflect.TypeOf((*MockEndBlockAppModule)(nil).GetQueryCmd))
}

// GetTxCmd mocks base method.
func (m *MockEndBlockAppModule) GetTxCmd() *cobra.Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxCmd")
	ret0, _ := ret[0].(*cobra.Command)
	return ret0
}

// GetTxCmd indicates an expected call of GetTxCmd.
func (mr *MockEndBlockAppModuleMockRecorder) GetTxCmd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxCmd", reflect.TypeOf((*MockEndBlockAppModule)(nil).GetTxCmd))
}

// Name mocks base method.
func (m *MockEndBlockAppModule) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockEndBlockAppModuleMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockEndBlockAppModule)(nil).Name))
}

// RegisterGRPCGatewayRoutes mocks base method.
func (m *MockEndBlockAppModule) RegisterGRPCGatewayRoutes(arg0 client.Context, arg1 *runtime.ServeMux) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterGRPCGatewayRoutes", arg0, arg1)
}

// RegisterGRPCGatewayRoutes indicates an expected call of RegisterGRPCGatewayRoutes.
func (mr *MockEndBlockAppModuleMockRecorder) RegisterGRPCGatewayRoutes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterGRPCGatewayRoutes", reflect.TypeOf((*MockEndBlockAppModule)(nil).RegisterGRPCGatewayRoutes), arg0, arg1)
}

// RegisterInterfaces mocks base method.
func (m *MockEndBlockAppModule) RegisterInterfaces(arg0 types0.InterfaceRegistry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterInterfaces", arg0)
}

// RegisterInterfaces indicates an expected call of RegisterInterfaces.
func (mr *MockEndBlockAppModuleMockRecorder) RegisterInterfaces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInterfaces", reflect.TypeOf((*MockEndBlockAppModule)(nil).RegisterInterfaces), arg0)
}

// RegisterLegacyAminoCodec mocks base method.
func (m *MockEndBlockAppModule) RegisterLegacyAminoCodec(arg0 *codec.LegacyAmino) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterLegacyAminoCodec", arg0)
}

// RegisterLegacyAminoCodec indicates an expected call of RegisterLegacyAminoCodec.
func (mr *MockEndBlockAppModuleMockRecorder) RegisterLegacyAminoCodec(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterLegacyAminoCodec", reflect.TypeOf((*MockEndBlockAppModule)(nil).RegisterLegacyAminoCodec), arg0)
}