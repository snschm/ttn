// Automatically generated by MockGen. DO NOT EDIT!
// Source: ./api/networkserver/networkserver.pb.go

package networkserver

import (
	broker "github.com/TheThingsNetwork/ttn/api/broker"
	handler "github.com/TheThingsNetwork/ttn/api/handler"
	lorawan "github.com/TheThingsNetwork/ttn/api/protocol/lorawan"
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Mock of NetworkServerClient interface
type MockNetworkServerClient struct {
	ctrl     *gomock.Controller
	recorder *_MockNetworkServerClientRecorder
}

// Recorder for MockNetworkServerClient (not exported)
type _MockNetworkServerClientRecorder struct {
	mock *MockNetworkServerClient
}

func NewMockNetworkServerClient(ctrl *gomock.Controller) *MockNetworkServerClient {
	mock := &MockNetworkServerClient{ctrl: ctrl}
	mock.recorder = &_MockNetworkServerClientRecorder{mock}
	return mock
}

func (_m *MockNetworkServerClient) EXPECT() *_MockNetworkServerClientRecorder {
	return _m.recorder
}

func (_m *MockNetworkServerClient) GetDevices(ctx context.Context, in *DevicesRequest, opts ...grpc.CallOption) (*DevicesResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetDevices", _s...)
	ret0, _ := ret[0].(*DevicesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkServerClientRecorder) GetDevices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDevices", _s...)
}

func (_m *MockNetworkServerClient) GetDevice(ctx context.Context, in *lorawan.DeviceIdentifier, opts ...grpc.CallOption) (*lorawan.Device, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetDevice", _s...)
	ret0, _ := ret[0].(*lorawan.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkServerClientRecorder) GetDevice(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDevice", _s...)
}

func (_m *MockNetworkServerClient) PrepareActivation(ctx context.Context, in *broker.DeduplicatedDeviceActivationRequest, opts ...grpc.CallOption) (*broker.DeduplicatedDeviceActivationRequest, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "PrepareActivation", _s...)
	ret0, _ := ret[0].(*broker.DeduplicatedDeviceActivationRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkServerClientRecorder) PrepareActivation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PrepareActivation", _s...)
}

func (_m *MockNetworkServerClient) Activate(ctx context.Context, in *handler.DeviceActivationResponse, opts ...grpc.CallOption) (*handler.DeviceActivationResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Activate", _s...)
	ret0, _ := ret[0].(*handler.DeviceActivationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkServerClientRecorder) Activate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Activate", _s...)
}

func (_m *MockNetworkServerClient) Uplink(ctx context.Context, in *broker.DeduplicatedUplinkMessage, opts ...grpc.CallOption) (*broker.DeduplicatedUplinkMessage, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Uplink", _s...)
	ret0, _ := ret[0].(*broker.DeduplicatedUplinkMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkServerClientRecorder) Uplink(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Uplink", _s...)
}

func (_m *MockNetworkServerClient) Downlink(ctx context.Context, in *broker.DownlinkMessage, opts ...grpc.CallOption) (*broker.DownlinkMessage, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Downlink", _s...)
	ret0, _ := ret[0].(*broker.DownlinkMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkServerClientRecorder) Downlink(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Downlink", _s...)
}

// Mock of NetworkServerServer interface
type MockNetworkServerServer struct {
	ctrl     *gomock.Controller
	recorder *_MockNetworkServerServerRecorder
}

// Recorder for MockNetworkServerServer (not exported)
type _MockNetworkServerServerRecorder struct {
	mock *MockNetworkServerServer
}

func NewMockNetworkServerServer(ctrl *gomock.Controller) *MockNetworkServerServer {
	mock := &MockNetworkServerServer{ctrl: ctrl}
	mock.recorder = &_MockNetworkServerServerRecorder{mock}
	return mock
}

func (_m *MockNetworkServerServer) EXPECT() *_MockNetworkServerServerRecorder {
	return _m.recorder
}

func (_m *MockNetworkServerServer) GetDevices(_param0 context.Context, _param1 *DevicesRequest) (*DevicesResponse, error) {
	ret := _m.ctrl.Call(_m, "GetDevices", _param0, _param1)
	ret0, _ := ret[0].(*DevicesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkServerServerRecorder) GetDevices(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDevices", arg0, arg1)
}

func (_m *MockNetworkServerServer) GetDevice(_param0 context.Context, _param1 *lorawan.DeviceIdentifier) (*lorawan.Device, error) {
	ret := _m.ctrl.Call(_m, "GetDevice", _param0, _param1)
	ret0, _ := ret[0].(*lorawan.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkServerServerRecorder) GetDevice(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDevice", arg0, arg1)
}

func (_m *MockNetworkServerServer) PrepareActivation(_param0 context.Context, _param1 *broker.DeduplicatedDeviceActivationRequest) (*broker.DeduplicatedDeviceActivationRequest, error) {
	ret := _m.ctrl.Call(_m, "PrepareActivation", _param0, _param1)
	ret0, _ := ret[0].(*broker.DeduplicatedDeviceActivationRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkServerServerRecorder) PrepareActivation(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PrepareActivation", arg0, arg1)
}

func (_m *MockNetworkServerServer) Activate(_param0 context.Context, _param1 *handler.DeviceActivationResponse) (*handler.DeviceActivationResponse, error) {
	ret := _m.ctrl.Call(_m, "Activate", _param0, _param1)
	ret0, _ := ret[0].(*handler.DeviceActivationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkServerServerRecorder) Activate(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Activate", arg0, arg1)
}

func (_m *MockNetworkServerServer) Uplink(_param0 context.Context, _param1 *broker.DeduplicatedUplinkMessage) (*broker.DeduplicatedUplinkMessage, error) {
	ret := _m.ctrl.Call(_m, "Uplink", _param0, _param1)
	ret0, _ := ret[0].(*broker.DeduplicatedUplinkMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkServerServerRecorder) Uplink(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Uplink", arg0, arg1)
}

func (_m *MockNetworkServerServer) Downlink(_param0 context.Context, _param1 *broker.DownlinkMessage) (*broker.DownlinkMessage, error) {
	ret := _m.ctrl.Call(_m, "Downlink", _param0, _param1)
	ret0, _ := ret[0].(*broker.DownlinkMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkServerServerRecorder) Downlink(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Downlink", arg0, arg1)
}

// Mock of NetworkServerManagerClient interface
type MockNetworkServerManagerClient struct {
	ctrl     *gomock.Controller
	recorder *_MockNetworkServerManagerClientRecorder
}

// Recorder for MockNetworkServerManagerClient (not exported)
type _MockNetworkServerManagerClientRecorder struct {
	mock *MockNetworkServerManagerClient
}

func NewMockNetworkServerManagerClient(ctrl *gomock.Controller) *MockNetworkServerManagerClient {
	mock := &MockNetworkServerManagerClient{ctrl: ctrl}
	mock.recorder = &_MockNetworkServerManagerClientRecorder{mock}
	return mock
}

func (_m *MockNetworkServerManagerClient) EXPECT() *_MockNetworkServerManagerClientRecorder {
	return _m.recorder
}

func (_m *MockNetworkServerManagerClient) GetStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*Status, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetStatus", _s...)
	ret0, _ := ret[0].(*Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkServerManagerClientRecorder) GetStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetStatus", _s...)
}

// Mock of NetworkServerManagerServer interface
type MockNetworkServerManagerServer struct {
	ctrl     *gomock.Controller
	recorder *_MockNetworkServerManagerServerRecorder
}

// Recorder for MockNetworkServerManagerServer (not exported)
type _MockNetworkServerManagerServerRecorder struct {
	mock *MockNetworkServerManagerServer
}

func NewMockNetworkServerManagerServer(ctrl *gomock.Controller) *MockNetworkServerManagerServer {
	mock := &MockNetworkServerManagerServer{ctrl: ctrl}
	mock.recorder = &_MockNetworkServerManagerServerRecorder{mock}
	return mock
}

func (_m *MockNetworkServerManagerServer) EXPECT() *_MockNetworkServerManagerServerRecorder {
	return _m.recorder
}

func (_m *MockNetworkServerManagerServer) GetStatus(_param0 context.Context, _param1 *StatusRequest) (*Status, error) {
	ret := _m.ctrl.Call(_m, "GetStatus", _param0, _param1)
	ret0, _ := ret[0].(*Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkServerManagerServerRecorder) GetStatus(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetStatus", arg0, arg1)
}
