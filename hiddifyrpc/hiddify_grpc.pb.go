// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: hiddifyrpc/hiddify.proto

package hiddifyrpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	BarVPN_SayHello_FullMethodName              = "/hiddifyrpc.BarVPN/SayHello"
	BarVPN_SayHelloStream_FullMethodName        = "/hiddifyrpc.BarVPN/SayHelloStream"
	BarVPN_Start_FullMethodName                 = "/hiddifyrpc.BarVPN/Start"
	BarVPN_CoreInfoListener_FullMethodName      = "/hiddifyrpc.BarVPN/CoreInfoListener"
	BarVPN_OutboundsInfo_FullMethodName         = "/hiddifyrpc.BarVPN/OutboundsInfo"
	BarVPN_MainOutboundsInfo_FullMethodName     = "/hiddifyrpc.BarVPN/MainOutboundsInfo"
	BarVPN_GetSystemInfo_FullMethodName         = "/hiddifyrpc.BarVPN/GetSystemInfo"
	BarVPN_Setup_FullMethodName                 = "/hiddifyrpc.BarVPN/Setup"
	BarVPN_Parse_FullMethodName                 = "/hiddifyrpc.BarVPN/Parse"
	BarVPN_StartService_FullMethodName          = "/hiddifyrpc.BarVPN/StartService"
	BarVPN_Stop_FullMethodName                  = "/hiddifyrpc.BarVPN/Stop"
	BarVPN_Restart_FullMethodName               = "/hiddifyrpc.BarVPN/Restart"
	BarVPN_SelectOutbound_FullMethodName        = "/hiddifyrpc.BarVPN/SelectOutbound"
	BarVPN_UrlTest_FullMethodName               = "/hiddifyrpc.BarVPN/UrlTest"
	BarVPN_GenerateWarpConfig_FullMethodName    = "/hiddifyrpc.BarVPN/GenerateWarpConfig"
	BarVPN_GetSystemProxyStatus_FullMethodName  = "/hiddifyrpc.BarVPN/GetSystemProxyStatus"
	BarVPN_SetSystemProxyEnabled_FullMethodName = "/hiddifyrpc.BarVPN/SetSystemProxyEnabled"
	BarVPN_LogListener_FullMethodName           = "/hiddifyrpc.BarVPN/LogListener"
)

// BarVPNClient is the client API for BarVPN service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BarVPNClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	SayHelloStream(ctx context.Context, opts ...grpc.CallOption) (BarVPN_SayHelloStreamClient, error)
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*CoreInfoResponse, error)
	CoreInfoListener(ctx context.Context, opts ...grpc.CallOption) (BarVPN_CoreInfoListenerClient, error)
	OutboundsInfo(ctx context.Context, opts ...grpc.CallOption) (BarVPN_OutboundsInfoClient, error)
	MainOutboundsInfo(ctx context.Context, opts ...grpc.CallOption) (BarVPN_MainOutboundsInfoClient, error)
	GetSystemInfo(ctx context.Context, opts ...grpc.CallOption) (BarVPN_GetSystemInfoClient, error)
	Setup(ctx context.Context, in *SetupRequest, opts ...grpc.CallOption) (*Response, error)
	Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResponse, error)
	//rpc ChangeConfigOptions (ChangeConfigOptionsRequest) returns (CoreInfoResponse);
	//rpc GenerateConfig (GenerateConfigRequest) returns (GenerateConfigResponse);
	StartService(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*CoreInfoResponse, error)
	Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CoreInfoResponse, error)
	Restart(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*CoreInfoResponse, error)
	SelectOutbound(ctx context.Context, in *SelectOutboundRequest, opts ...grpc.CallOption) (*Response, error)
	UrlTest(ctx context.Context, in *UrlTestRequest, opts ...grpc.CallOption) (*Response, error)
	GenerateWarpConfig(ctx context.Context, in *GenerateWarpConfigRequest, opts ...grpc.CallOption) (*WarpGenerationResponse, error)
	GetSystemProxyStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SystemProxyStatus, error)
	SetSystemProxyEnabled(ctx context.Context, in *SetSystemProxyEnabledRequest, opts ...grpc.CallOption) (*Response, error)
	LogListener(ctx context.Context, opts ...grpc.CallOption) (BarVPN_LogListenerClient, error)
}

type hiddifyClient struct {
	cc grpc.ClientConnInterface
}

func NewBarVPNClient(cc grpc.ClientConnInterface) BarVPNClient {
	return &hiddifyClient{cc}
}

func (c *hiddifyClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, BarVPN_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hiddifyClient) SayHelloStream(ctx context.Context, opts ...grpc.CallOption) (BarVPN_SayHelloStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &BarVPN_ServiceDesc.Streams[0], BarVPN_SayHelloStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &hiddifySayHelloStreamClient{stream}
	return x, nil
}

type BarVPN_SayHelloStreamClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type hiddifySayHelloStreamClient struct {
	grpc.ClientStream
}

func (x *hiddifySayHelloStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *hiddifySayHelloStreamClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hiddifyClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*CoreInfoResponse, error) {
	out := new(CoreInfoResponse)
	err := c.cc.Invoke(ctx, BarVPN_Start_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hiddifyClient) CoreInfoListener(ctx context.Context, opts ...grpc.CallOption) (BarVPN_CoreInfoListenerClient, error) {
	stream, err := c.cc.NewStream(ctx, &BarVPN_ServiceDesc.Streams[1], BarVPN_CoreInfoListener_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &hiddifyCoreInfoListenerClient{stream}
	return x, nil
}

type BarVPN_CoreInfoListenerClient interface {
	Send(*StopRequest) error
	Recv() (*CoreInfoResponse, error)
	grpc.ClientStream
}

type hiddifyCoreInfoListenerClient struct {
	grpc.ClientStream
}

func (x *hiddifyCoreInfoListenerClient) Send(m *StopRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *hiddifyCoreInfoListenerClient) Recv() (*CoreInfoResponse, error) {
	m := new(CoreInfoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hiddifyClient) OutboundsInfo(ctx context.Context, opts ...grpc.CallOption) (BarVPN_OutboundsInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &BarVPN_ServiceDesc.Streams[2], BarVPN_OutboundsInfo_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &hiddifyOutboundsInfoClient{stream}
	return x, nil
}

type BarVPN_OutboundsInfoClient interface {
	Send(*StopRequest) error
	Recv() (*OutboundGroupList, error)
	grpc.ClientStream
}

type hiddifyOutboundsInfoClient struct {
	grpc.ClientStream
}

func (x *hiddifyOutboundsInfoClient) Send(m *StopRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *hiddifyOutboundsInfoClient) Recv() (*OutboundGroupList, error) {
	m := new(OutboundGroupList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hiddifyClient) MainOutboundsInfo(ctx context.Context, opts ...grpc.CallOption) (BarVPN_MainOutboundsInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &BarVPN_ServiceDesc.Streams[3], BarVPN_MainOutboundsInfo_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &hiddifyMainOutboundsInfoClient{stream}
	return x, nil
}

type BarVPN_MainOutboundsInfoClient interface {
	Send(*StopRequest) error
	Recv() (*OutboundGroupList, error)
	grpc.ClientStream
}

type hiddifyMainOutboundsInfoClient struct {
	grpc.ClientStream
}

func (x *hiddifyMainOutboundsInfoClient) Send(m *StopRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *hiddifyMainOutboundsInfoClient) Recv() (*OutboundGroupList, error) {
	m := new(OutboundGroupList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hiddifyClient) GetSystemInfo(ctx context.Context, opts ...grpc.CallOption) (BarVPN_GetSystemInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &BarVPN_ServiceDesc.Streams[4], BarVPN_GetSystemInfo_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &hiddifyGetSystemInfoClient{stream}
	return x, nil
}

type BarVPN_GetSystemInfoClient interface {
	Send(*StopRequest) error
	Recv() (*SystemInfo, error)
	grpc.ClientStream
}

type hiddifyGetSystemInfoClient struct {
	grpc.ClientStream
}

func (x *hiddifyGetSystemInfoClient) Send(m *StopRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *hiddifyGetSystemInfoClient) Recv() (*SystemInfo, error) {
	m := new(SystemInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hiddifyClient) Setup(ctx context.Context, in *SetupRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, BarVPN_Setup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hiddifyClient) Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResponse, error) {
	out := new(ParseResponse)
	err := c.cc.Invoke(ctx, BarVPN_Parse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hiddifyClient) StartService(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*CoreInfoResponse, error) {
	out := new(CoreInfoResponse)
	err := c.cc.Invoke(ctx, BarVPN_StartService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hiddifyClient) Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CoreInfoResponse, error) {
	out := new(CoreInfoResponse)
	err := c.cc.Invoke(ctx, BarVPN_Stop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hiddifyClient) Restart(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*CoreInfoResponse, error) {
	out := new(CoreInfoResponse)
	err := c.cc.Invoke(ctx, BarVPN_Restart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hiddifyClient) SelectOutbound(ctx context.Context, in *SelectOutboundRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, BarVPN_SelectOutbound_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hiddifyClient) UrlTest(ctx context.Context, in *UrlTestRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, BarVPN_UrlTest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hiddifyClient) GenerateWarpConfig(ctx context.Context, in *GenerateWarpConfigRequest, opts ...grpc.CallOption) (*WarpGenerationResponse, error) {
	out := new(WarpGenerationResponse)
	err := c.cc.Invoke(ctx, BarVPN_GenerateWarpConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hiddifyClient) GetSystemProxyStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SystemProxyStatus, error) {
	out := new(SystemProxyStatus)
	err := c.cc.Invoke(ctx, BarVPN_GetSystemProxyStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hiddifyClient) SetSystemProxyEnabled(ctx context.Context, in *SetSystemProxyEnabledRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, BarVPN_SetSystemProxyEnabled_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hiddifyClient) LogListener(ctx context.Context, opts ...grpc.CallOption) (BarVPN_LogListenerClient, error) {
	stream, err := c.cc.NewStream(ctx, &BarVPN_ServiceDesc.Streams[5], BarVPN_LogListener_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &hiddifyLogListenerClient{stream}
	return x, nil
}

type BarVPN_LogListenerClient interface {
	Send(*StopRequest) error
	Recv() (*LogMessage, error)
	grpc.ClientStream
}

type hiddifyLogListenerClient struct {
	grpc.ClientStream
}

func (x *hiddifyLogListenerClient) Send(m *StopRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *hiddifyLogListenerClient) Recv() (*LogMessage, error) {
	m := new(LogMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BarVPNServer is the server API for BarVPN service.
// All implementations must embed UnimplementedBarVPNServer
// for forward compatibility
type BarVPNServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	SayHelloStream(BarVPN_SayHelloStreamServer) error
	Start(context.Context, *StartRequest) (*CoreInfoResponse, error)
	CoreInfoListener(BarVPN_CoreInfoListenerServer) error
	OutboundsInfo(BarVPN_OutboundsInfoServer) error
	MainOutboundsInfo(BarVPN_MainOutboundsInfoServer) error
	GetSystemInfo(BarVPN_GetSystemInfoServer) error
	Setup(context.Context, *SetupRequest) (*Response, error)
	Parse(context.Context, *ParseRequest) (*ParseResponse, error)
	//rpc ChangeConfigOptions (ChangeConfigOptionsRequest) returns (CoreInfoResponse);
	//rpc GenerateConfig (GenerateConfigRequest) returns (GenerateConfigResponse);
	StartService(context.Context, *StartRequest) (*CoreInfoResponse, error)
	Stop(context.Context, *Empty) (*CoreInfoResponse, error)
	Restart(context.Context, *StartRequest) (*CoreInfoResponse, error)
	SelectOutbound(context.Context, *SelectOutboundRequest) (*Response, error)
	UrlTest(context.Context, *UrlTestRequest) (*Response, error)
	GenerateWarpConfig(context.Context, *GenerateWarpConfigRequest) (*WarpGenerationResponse, error)
	GetSystemProxyStatus(context.Context, *Empty) (*SystemProxyStatus, error)
	SetSystemProxyEnabled(context.Context, *SetSystemProxyEnabledRequest) (*Response, error)
	LogListener(BarVPN_LogListenerServer) error
	mustEmbedUnimplementedBarVPNServer()
}

// UnimplementedBarVPNServer must be embedded to have forward compatible implementations.
type UnimplementedBarVPNServer struct {
}

func (UnimplementedBarVPNServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedBarVPNServer) SayHelloStream(BarVPN_SayHelloStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloStream not implemented")
}
func (UnimplementedBarVPNServer) Start(context.Context, *StartRequest) (*CoreInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedBarVPNServer) CoreInfoListener(BarVPN_CoreInfoListenerServer) error {
	return status.Errorf(codes.Unimplemented, "method CoreInfoListener not implemented")
}
func (UnimplementedBarVPNServer) OutboundsInfo(BarVPN_OutboundsInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method OutboundsInfo not implemented")
}
func (UnimplementedBarVPNServer) MainOutboundsInfo(BarVPN_MainOutboundsInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method MainOutboundsInfo not implemented")
}
func (UnimplementedBarVPNServer) GetSystemInfo(BarVPN_GetSystemInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method GetSystemInfo not implemented")
}
func (UnimplementedBarVPNServer) Setup(context.Context, *SetupRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Setup not implemented")
}
func (UnimplementedBarVPNServer) Parse(context.Context, *ParseRequest) (*ParseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Parse not implemented")
}
func (UnimplementedBarVPNServer) StartService(context.Context, *StartRequest) (*CoreInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartService not implemented")
}
func (UnimplementedBarVPNServer) Stop(context.Context, *Empty) (*CoreInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedBarVPNServer) Restart(context.Context, *StartRequest) (*CoreInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (UnimplementedBarVPNServer) SelectOutbound(context.Context, *SelectOutboundRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectOutbound not implemented")
}
func (UnimplementedBarVPNServer) UrlTest(context.Context, *UrlTestRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UrlTest not implemented")
}
func (UnimplementedBarVPNServer) GenerateWarpConfig(context.Context, *GenerateWarpConfigRequest) (*WarpGenerationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateWarpConfig not implemented")
}
func (UnimplementedBarVPNServer) GetSystemProxyStatus(context.Context, *Empty) (*SystemProxyStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemProxyStatus not implemented")
}
func (UnimplementedBarVPNServer) SetSystemProxyEnabled(context.Context, *SetSystemProxyEnabledRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSystemProxyEnabled not implemented")
}
func (UnimplementedBarVPNServer) LogListener(BarVPN_LogListenerServer) error {
	return status.Errorf(codes.Unimplemented, "method LogListener not implemented")
}
func (UnimplementedBarVPNServer) mustEmbedUnimplementedBarVPNServer() {}

// UnsafeBarVPNServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BarVPNServer will
// result in compilation errors.
type UnsafeBarVPNServer interface {
	mustEmbedUnimplementedBarVPNServer()
}

func RegisterBarVPNServer(s grpc.ServiceRegistrar, srv BarVPNServer) {
	s.RegisterService(&BarVPN_ServiceDesc, srv)
}

func _BarVPN_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarVPNServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BarVPN_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarVPNServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BarVPN_SayHelloStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BarVPNServer).SayHelloStream(&hiddifySayHelloStreamServer{stream})
}

type BarVPN_SayHelloStreamServer interface {
	Send(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type hiddifySayHelloStreamServer struct {
	grpc.ServerStream
}

func (x *hiddifySayHelloStreamServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *hiddifySayHelloStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BarVPN_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarVPNServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BarVPN_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarVPNServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BarVPN_CoreInfoListener_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BarVPNServer).CoreInfoListener(&hiddifyCoreInfoListenerServer{stream})
}

type BarVPN_CoreInfoListenerServer interface {
	Send(*CoreInfoResponse) error
	Recv() (*StopRequest, error)
	grpc.ServerStream
}

type hiddifyCoreInfoListenerServer struct {
	grpc.ServerStream
}

func (x *hiddifyCoreInfoListenerServer) Send(m *CoreInfoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *hiddifyCoreInfoListenerServer) Recv() (*StopRequest, error) {
	m := new(StopRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BarVPN_OutboundsInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BarVPNServer).OutboundsInfo(&hiddifyOutboundsInfoServer{stream})
}

type BarVPN_OutboundsInfoServer interface {
	Send(*OutboundGroupList) error
	Recv() (*StopRequest, error)
	grpc.ServerStream
}

type hiddifyOutboundsInfoServer struct {
	grpc.ServerStream
}

func (x *hiddifyOutboundsInfoServer) Send(m *OutboundGroupList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *hiddifyOutboundsInfoServer) Recv() (*StopRequest, error) {
	m := new(StopRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BarVPN_MainOutboundsInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BarVPNServer).MainOutboundsInfo(&hiddifyMainOutboundsInfoServer{stream})
}

type BarVPN_MainOutboundsInfoServer interface {
	Send(*OutboundGroupList) error
	Recv() (*StopRequest, error)
	grpc.ServerStream
}

type hiddifyMainOutboundsInfoServer struct {
	grpc.ServerStream
}

func (x *hiddifyMainOutboundsInfoServer) Send(m *OutboundGroupList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *hiddifyMainOutboundsInfoServer) Recv() (*StopRequest, error) {
	m := new(StopRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BarVPN_GetSystemInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BarVPNServer).GetSystemInfo(&hiddifyGetSystemInfoServer{stream})
}

type BarVPN_GetSystemInfoServer interface {
	Send(*SystemInfo) error
	Recv() (*StopRequest, error)
	grpc.ServerStream
}

type hiddifyGetSystemInfoServer struct {
	grpc.ServerStream
}

func (x *hiddifyGetSystemInfoServer) Send(m *SystemInfo) error {
	return x.ServerStream.SendMsg(m)
}

func (x *hiddifyGetSystemInfoServer) Recv() (*StopRequest, error) {
	m := new(StopRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BarVPN_Setup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarVPNServer).Setup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BarVPN_Setup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarVPNServer).Setup(ctx, req.(*SetupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BarVPN_Parse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarVPNServer).Parse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BarVPN_Parse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarVPNServer).Parse(ctx, req.(*ParseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BarVPN_StartService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarVPNServer).StartService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BarVPN_StartService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarVPNServer).StartService(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BarVPN_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarVPNServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BarVPN_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarVPNServer).Stop(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BarVPN_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarVPNServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BarVPN_Restart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarVPNServer).Restart(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BarVPN_SelectOutbound_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectOutboundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarVPNServer).SelectOutbound(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BarVPN_SelectOutbound_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarVPNServer).SelectOutbound(ctx, req.(*SelectOutboundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BarVPN_UrlTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UrlTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarVPNServer).UrlTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BarVPN_UrlTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarVPNServer).UrlTest(ctx, req.(*UrlTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BarVPN_GenerateWarpConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateWarpConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarVPNServer).GenerateWarpConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BarVPN_GenerateWarpConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarVPNServer).GenerateWarpConfig(ctx, req.(*GenerateWarpConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BarVPN_GetSystemProxyStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarVPNServer).GetSystemProxyStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BarVPN_GetSystemProxyStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarVPNServer).GetSystemProxyStatus(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BarVPN_SetSystemProxyEnabled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSystemProxyEnabledRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarVPNServer).SetSystemProxyEnabled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BarVPN_SetSystemProxyEnabled_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarVPNServer).SetSystemProxyEnabled(ctx, req.(*SetSystemProxyEnabledRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BarVPN_LogListener_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BarVPNServer).LogListener(&hiddifyLogListenerServer{stream})
}

type BarVPN_LogListenerServer interface {
	Send(*LogMessage) error
	Recv() (*StopRequest, error)
	grpc.ServerStream
}

type hiddifyLogListenerServer struct {
	grpc.ServerStream
}

func (x *hiddifyLogListenerServer) Send(m *LogMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *hiddifyLogListenerServer) Recv() (*StopRequest, error) {
	m := new(StopRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BarVPN_ServiceDesc is the grpc.ServiceDesc for BarVPN service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BarVPN_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hiddifyrpc.BarVPN",
	HandlerType: (*BarVPNServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _BarVPN_SayHello_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _BarVPN_Start_Handler,
		},
		{
			MethodName: "Setup",
			Handler:    _BarVPN_Setup_Handler,
		},
		{
			MethodName: "Parse",
			Handler:    _BarVPN_Parse_Handler,
		},
		{
			MethodName: "StartService",
			Handler:    _BarVPN_StartService_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _BarVPN_Stop_Handler,
		},
		{
			MethodName: "Restart",
			Handler:    _BarVPN_Restart_Handler,
		},
		{
			MethodName: "SelectOutbound",
			Handler:    _BarVPN_SelectOutbound_Handler,
		},
		{
			MethodName: "UrlTest",
			Handler:    _BarVPN_UrlTest_Handler,
		},
		{
			MethodName: "GenerateWarpConfig",
			Handler:    _BarVPN_GenerateWarpConfig_Handler,
		},
		{
			MethodName: "GetSystemProxyStatus",
			Handler:    _BarVPN_GetSystemProxyStatus_Handler,
		},
		{
			MethodName: "SetSystemProxyEnabled",
			Handler:    _BarVPN_SetSystemProxyEnabled_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloStream",
			Handler:       _BarVPN_SayHelloStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "CoreInfoListener",
			Handler:       _BarVPN_CoreInfoListener_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "OutboundsInfo",
			Handler:       _BarVPN_OutboundsInfo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "MainOutboundsInfo",
			Handler:       _BarVPN_MainOutboundsInfo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetSystemInfo",
			Handler:       _BarVPN_GetSystemInfo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "LogListener",
			Handler:       _BarVPN_LogListener_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hiddifyrpc/hiddify.proto",
}
