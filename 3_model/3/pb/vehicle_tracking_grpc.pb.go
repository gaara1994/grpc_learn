// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: vehicle_tracking.proto

package pb

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

// VehicleTrackingServiceClient is the client API for VehicleTrackingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VehicleTrackingServiceClient interface {
	RecordVehicleRoute(ctx context.Context, opts ...grpc.CallOption) (VehicleTrackingService_RecordVehicleRouteClient, error)
}

type vehicleTrackingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVehicleTrackingServiceClient(cc grpc.ClientConnInterface) VehicleTrackingServiceClient {
	return &vehicleTrackingServiceClient{cc}
}

func (c *vehicleTrackingServiceClient) RecordVehicleRoute(ctx context.Context, opts ...grpc.CallOption) (VehicleTrackingService_RecordVehicleRouteClient, error) {
	stream, err := c.cc.NewStream(ctx, &VehicleTrackingService_ServiceDesc.Streams[0], "/vehicletracking.VehicleTrackingService/RecordVehicleRoute", opts...)
	if err != nil {
		return nil, err
	}
	x := &vehicleTrackingServiceRecordVehicleRouteClient{stream}
	return x, nil
}

type VehicleTrackingService_RecordVehicleRouteClient interface {
	Send(*VehicleLocation) error
	CloseAndRecv() (*RouteSummary, error)
	grpc.ClientStream
}

type vehicleTrackingServiceRecordVehicleRouteClient struct {
	grpc.ClientStream
}

func (x *vehicleTrackingServiceRecordVehicleRouteClient) Send(m *VehicleLocation) error {
	return x.ClientStream.SendMsg(m)
}

func (x *vehicleTrackingServiceRecordVehicleRouteClient) CloseAndRecv() (*RouteSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RouteSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VehicleTrackingServiceServer is the server API for VehicleTrackingService service.
// All implementations must embed UnimplementedVehicleTrackingServiceServer
// for forward compatibility
type VehicleTrackingServiceServer interface {
	RecordVehicleRoute(VehicleTrackingService_RecordVehicleRouteServer) error
	mustEmbedUnimplementedVehicleTrackingServiceServer()
}

// UnimplementedVehicleTrackingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVehicleTrackingServiceServer struct {
}

func (UnimplementedVehicleTrackingServiceServer) RecordVehicleRoute(VehicleTrackingService_RecordVehicleRouteServer) error {
	return status.Errorf(codes.Unimplemented, "method RecordVehicleRoute not implemented")
}
func (UnimplementedVehicleTrackingServiceServer) mustEmbedUnimplementedVehicleTrackingServiceServer() {
}

// UnsafeVehicleTrackingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VehicleTrackingServiceServer will
// result in compilation errors.
type UnsafeVehicleTrackingServiceServer interface {
	mustEmbedUnimplementedVehicleTrackingServiceServer()
}

func RegisterVehicleTrackingServiceServer(s grpc.ServiceRegistrar, srv VehicleTrackingServiceServer) {
	s.RegisterService(&VehicleTrackingService_ServiceDesc, srv)
}

func _VehicleTrackingService_RecordVehicleRoute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VehicleTrackingServiceServer).RecordVehicleRoute(&vehicleTrackingServiceRecordVehicleRouteServer{stream})
}

type VehicleTrackingService_RecordVehicleRouteServer interface {
	SendAndClose(*RouteSummary) error
	Recv() (*VehicleLocation, error)
	grpc.ServerStream
}

type vehicleTrackingServiceRecordVehicleRouteServer struct {
	grpc.ServerStream
}

func (x *vehicleTrackingServiceRecordVehicleRouteServer) SendAndClose(m *RouteSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *vehicleTrackingServiceRecordVehicleRouteServer) Recv() (*VehicleLocation, error) {
	m := new(VehicleLocation)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VehicleTrackingService_ServiceDesc is the grpc.ServiceDesc for VehicleTrackingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VehicleTrackingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vehicletracking.VehicleTrackingService",
	HandlerType: (*VehicleTrackingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RecordVehicleRoute",
			Handler:       _VehicleTrackingService_RecordVehicleRoute_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "vehicle_tracking.proto",
}
