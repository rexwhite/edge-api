// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/services/devices.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	inventory "github.com/redhatinsights/edge-api/pkg/clients/inventory"
	models "github.com/redhatinsights/edge-api/pkg/models"
	gorm "gorm.io/gorm"
)

// MockDeviceServiceInterface is a mock of DeviceServiceInterface interface.
type MockDeviceServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceServiceInterfaceMockRecorder
}

// MockDeviceServiceInterfaceMockRecorder is the mock recorder for MockDeviceServiceInterface.
type MockDeviceServiceInterfaceMockRecorder struct {
	mock *MockDeviceServiceInterface
}

// NewMockDeviceServiceInterface creates a new mock instance.
func NewMockDeviceServiceInterface(ctrl *gomock.Controller) *MockDeviceServiceInterface {
	mock := &MockDeviceServiceInterface{ctrl: ctrl}
	mock.recorder = &MockDeviceServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeviceServiceInterface) EXPECT() *MockDeviceServiceInterfaceMockRecorder {
	return m.recorder
}

// GetDeviceByID mocks base method.
func (m *MockDeviceServiceInterface) GetDeviceByID(deviceID uint) (*models.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceByID", deviceID)
	ret0, _ := ret[0].(*models.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceByID indicates an expected call of GetDeviceByID.
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDeviceByID(deviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceByID", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDeviceByID), deviceID)
}

// GetDeviceByUUID mocks base method.
func (m *MockDeviceServiceInterface) GetDeviceByUUID(deviceUUID string) (*models.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceByUUID", deviceUUID)
	ret0, _ := ret[0].(*models.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceByUUID indicates an expected call of GetDeviceByUUID.
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDeviceByUUID(deviceUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceByUUID", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDeviceByUUID), deviceUUID)
}

// GetDeviceDetails mocks base method.
func (m *MockDeviceServiceInterface) GetDeviceDetails(device inventory.Device) (*models.DeviceDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceDetails", device)
	ret0, _ := ret[0].(*models.DeviceDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceDetails indicates an expected call of GetDeviceDetails.
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDeviceDetails(device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceDetails", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDeviceDetails), device)
}

// GetDeviceDetailsByUUID mocks base method.
func (m *MockDeviceServiceInterface) GetDeviceDetailsByUUID(deviceUUID string) (*models.DeviceDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceDetailsByUUID", deviceUUID)
	ret0, _ := ret[0].(*models.DeviceDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceDetailsByUUID indicates an expected call of GetDeviceDetailsByUUID.
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDeviceDetailsByUUID(deviceUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceDetailsByUUID", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDeviceDetailsByUUID), deviceUUID)
}

// GetDeviceImageInfo mocks base method.
func (m *MockDeviceServiceInterface) GetDeviceImageInfo(device inventory.Device) (*models.ImageInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceImageInfo", device)
	ret0, _ := ret[0].(*models.ImageInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceImageInfo indicates an expected call of GetDeviceImageInfo.
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDeviceImageInfo(device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceImageInfo", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDeviceImageInfo), device)
}

// GetDeviceImageInfoByUUID mocks base method.
func (m *MockDeviceServiceInterface) GetDeviceImageInfoByUUID(deviceUUID string) (*models.ImageInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceImageInfoByUUID", deviceUUID)
	ret0, _ := ret[0].(*models.ImageInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceImageInfoByUUID indicates an expected call of GetDeviceImageInfoByUUID.
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDeviceImageInfoByUUID(deviceUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceImageInfoByUUID", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDeviceImageInfoByUUID), deviceUUID)
}

// GetDeviceLastBootedDeployment mocks base method.
func (m *MockDeviceServiceInterface) GetDeviceLastBootedDeployment(device inventory.Device) *inventory.OSTree {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceLastBootedDeployment", device)
	ret0, _ := ret[0].(*inventory.OSTree)
	return ret0
}

// GetDeviceLastBootedDeployment indicates an expected call of GetDeviceLastBootedDeployment.
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDeviceLastBootedDeployment(device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceLastBootedDeployment", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDeviceLastBootedDeployment), device)
}

// GetDeviceLastDeployment mocks base method.
func (m *MockDeviceServiceInterface) GetDeviceLastDeployment(device inventory.Device) *inventory.OSTree {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceLastDeployment", device)
	ret0, _ := ret[0].(*inventory.OSTree)
	return ret0
}

// GetDeviceLastDeployment indicates an expected call of GetDeviceLastDeployment.
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDeviceLastDeployment(device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceLastDeployment", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDeviceLastDeployment), device)
}

// GetDevices mocks base method.
func (m *MockDeviceServiceInterface) GetDevices(params *inventory.Params) (*models.DeviceDetailsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevices", params)
	ret0, _ := ret[0].(*models.DeviceDetailsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevices indicates an expected call of GetDevices.
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDevices(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevices", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDevices), params)
}

// GetDevicesCount mocks base method.
func (m *MockDeviceServiceInterface) GetDevicesCount(tx *gorm.DB) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevicesCount", tx)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevicesCount indicates an expected call of GetDevicesCount.
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDevicesCount(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevicesCount", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDevicesCount), tx)
}

// GetDevicesView mocks base method.
func (m *MockDeviceServiceInterface) GetDevicesView(limit, offset int, tx *gorm.DB) (*models.DeviceViewList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevicesView", limit, offset, tx)
	ret0, _ := ret[0].(*models.DeviceViewList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevicesView indicates an expected call of GetDevicesView.
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDevicesView(limit, offset, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevicesView", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDevicesView), limit, offset, tx)
}

// GetLatestCommitFromDevices mocks base method.
func (m *MockDeviceServiceInterface) GetLatestCommitFromDevices(account, orgID string, devicesUUID []string) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestCommitFromDevices", account, orgID, devicesUUID)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestCommitFromDevices indicates an expected call of GetLatestCommitFromDevices.
func (mr *MockDeviceServiceInterfaceMockRecorder) GetLatestCommitFromDevices(account, orgID, devicesUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestCommitFromDevices", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetLatestCommitFromDevices), account, orgID, devicesUUID)
}

// GetUpdateAvailableForDevice mocks base method.
func (m *MockDeviceServiceInterface) GetUpdateAvailableForDevice(device inventory.Device, latest bool) ([]models.ImageUpdateAvailable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpdateAvailableForDevice", device, latest)
	ret0, _ := ret[0].([]models.ImageUpdateAvailable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUpdateAvailableForDevice indicates an expected call of GetUpdateAvailableForDevice.
func (mr *MockDeviceServiceInterfaceMockRecorder) GetUpdateAvailableForDevice(device, latest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpdateAvailableForDevice", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetUpdateAvailableForDevice), device, latest)
}

// GetUpdateAvailableForDeviceByUUID mocks base method.
func (m *MockDeviceServiceInterface) GetUpdateAvailableForDeviceByUUID(deviceUUID string, latest bool) ([]models.ImageUpdateAvailable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpdateAvailableForDeviceByUUID", deviceUUID, latest)
	ret0, _ := ret[0].([]models.ImageUpdateAvailable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUpdateAvailableForDeviceByUUID indicates an expected call of GetUpdateAvailableForDeviceByUUID.
func (mr *MockDeviceServiceInterfaceMockRecorder) GetUpdateAvailableForDeviceByUUID(deviceUUID, latest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpdateAvailableForDeviceByUUID", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetUpdateAvailableForDeviceByUUID), deviceUUID, latest)
}

// ProcessPlatformInventoryCreateEvent mocks base method.
func (m *MockDeviceServiceInterface) ProcessPlatformInventoryCreateEvent(message []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessPlatformInventoryCreateEvent", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessPlatformInventoryCreateEvent indicates an expected call of ProcessPlatformInventoryCreateEvent.
func (mr *MockDeviceServiceInterfaceMockRecorder) ProcessPlatformInventoryCreateEvent(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessPlatformInventoryCreateEvent", reflect.TypeOf((*MockDeviceServiceInterface)(nil).ProcessPlatformInventoryCreateEvent), message)
}

// ProcessPlatformInventoryDeleteEvent mocks base method.
func (m *MockDeviceServiceInterface) ProcessPlatformInventoryDeleteEvent(message []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessPlatformInventoryDeleteEvent", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessPlatformInventoryDeleteEvent indicates an expected call of ProcessPlatformInventoryDeleteEvent.
func (mr *MockDeviceServiceInterfaceMockRecorder) ProcessPlatformInventoryDeleteEvent(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessPlatformInventoryDeleteEvent", reflect.TypeOf((*MockDeviceServiceInterface)(nil).ProcessPlatformInventoryDeleteEvent), message)
}

// ProcessPlatformInventoryUpdatedEvent mocks base method.
func (m *MockDeviceServiceInterface) ProcessPlatformInventoryUpdatedEvent(message []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessPlatformInventoryUpdatedEvent", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessPlatformInventoryUpdatedEvent indicates an expected call of ProcessPlatformInventoryUpdatedEvent.
func (mr *MockDeviceServiceInterfaceMockRecorder) ProcessPlatformInventoryUpdatedEvent(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessPlatformInventoryUpdatedEvent", reflect.TypeOf((*MockDeviceServiceInterface)(nil).ProcessPlatformInventoryUpdatedEvent), message)
}
