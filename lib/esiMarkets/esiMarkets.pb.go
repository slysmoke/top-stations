// Code generated by protoc-gen-go. DO NOT EDIT.
// source: esiMarkets.proto

/*
Package esiMarkets is a generated protocol buffer package.

It is generated from these files:
	esiMarkets.proto

It has these top-level messages:
	GetOrderRequest
	GetRegionRequest
	GetTypeRequest
	GetRegionTypeRequest
	GetOrdersResponse
	GetRegionTypeUpdateStreamResponse
	RegionType
	Order
*/
package esiMarkets

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import google_protobuf2 "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetOrderRequest struct {
	// Fetch recorded history of an order
	OrderId uint64 `protobuf:"varint,1,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
}

func (m *GetOrderRequest) Reset()                    { *m = GetOrderRequest{} }
func (m *GetOrderRequest) String() string            { return proto.CompactTextString(m) }
func (*GetOrderRequest) ProtoMessage()               {}
func (*GetOrderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetOrderRequest) GetOrderId() uint64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type GetRegionRequest struct {
	// Defines which region the data is fetched for
	RegionId uint64 `protobuf:"varint,1,opt,name=region_id,json=regionId" json:"region_id,omitempty"`
}

func (m *GetRegionRequest) Reset()                    { *m = GetRegionRequest{} }
func (m *GetRegionRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRegionRequest) ProtoMessage()               {}
func (*GetRegionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetRegionRequest) GetRegionId() uint64 {
	if m != nil {
		return m.RegionId
	}
	return 0
}

type GetTypeRequest struct {
	// Defines which type the data is fetched for
	TypeId uint64 `protobuf:"varint,1,opt,name=type_id,json=typeId" json:"type_id,omitempty"`
}

func (m *GetTypeRequest) Reset()                    { *m = GetTypeRequest{} }
func (m *GetTypeRequest) String() string            { return proto.CompactTextString(m) }
func (*GetTypeRequest) ProtoMessage()               {}
func (*GetTypeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetTypeRequest) GetTypeId() uint64 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

type GetRegionTypeRequest struct {
	// Defines which region the data is fetched for
	RegionId uint64 `protobuf:"varint,1,opt,name=region_id,json=regionId" json:"region_id,omitempty"`
	// Defines which type the data is fetched for
	TypeId uint64 `protobuf:"varint,2,opt,name=type_id,json=typeId" json:"type_id,omitempty"`
}

func (m *GetRegionTypeRequest) Reset()                    { *m = GetRegionTypeRequest{} }
func (m *GetRegionTypeRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRegionTypeRequest) ProtoMessage()               {}
func (*GetRegionTypeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetRegionTypeRequest) GetRegionId() uint64 {
	if m != nil {
		return m.RegionId
	}
	return 0
}

func (m *GetRegionTypeRequest) GetTypeId() uint64 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

type GetOrdersResponse struct {
	// Matching orders
	Orders []*Order `protobuf:"bytes,1,rep,name=orders" json:"orders,omitempty"`
}

func (m *GetOrdersResponse) Reset()                    { *m = GetOrdersResponse{} }
func (m *GetOrdersResponse) String() string            { return proto.CompactTextString(m) }
func (*GetOrdersResponse) ProtoMessage()               {}
func (*GetOrdersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetOrdersResponse) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

type GetRegionTypeUpdateStreamResponse struct {
	// Region/tye pairs affected by update
	RegionTypes []*RegionType `protobuf:"bytes,1,rep,name=region_types,json=regionTypes" json:"region_types,omitempty"`
}

func (m *GetRegionTypeUpdateStreamResponse) Reset()         { *m = GetRegionTypeUpdateStreamResponse{} }
func (m *GetRegionTypeUpdateStreamResponse) String() string { return proto.CompactTextString(m) }
func (*GetRegionTypeUpdateStreamResponse) ProtoMessage()    {}
func (*GetRegionTypeUpdateStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5}
}

func (m *GetRegionTypeUpdateStreamResponse) GetRegionTypes() []*RegionType {
	if m != nil {
		return m.RegionTypes
	}
	return nil
}

type RegionType struct {
	// The update's region's ID
	RegionId uint64 `protobuf:"varint,1,opt,name=region_id,json=regionId" json:"region_id,omitempty"`
	// The update's type's ID
	TypeId uint64 `protobuf:"varint,2,opt,name=type_id,json=typeId" json:"type_id,omitempty"`
}

func (m *RegionType) Reset()                    { *m = RegionType{} }
func (m *RegionType) String() string            { return proto.CompactTextString(m) }
func (*RegionType) ProtoMessage()               {}
func (*RegionType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RegionType) GetRegionId() uint64 {
	if m != nil {
		return m.RegionId
	}
	return 0
}

func (m *RegionType) GetTypeId() uint64 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

type Order struct {
	// The order's ID
	OrderId uint64 `protobuf:"varint,1,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
	// The order's type's ID
	TypeId uint64 `protobuf:"varint,2,opt,name=type_id,json=typeId" json:"type_id,omitempty"`
	// The order's region's ID
	RegionId uint64 `protobuf:"varint,3,opt,name=region_id,json=regionId" json:"region_id,omitempty"`
	// ID of the order's location/station
	LocationId uint64 `protobuf:"varint,4,opt,name=location_id,json=locationId" json:"location_id,omitempty"`
	// Initital number of items of the order
	VolumeTotal uint64 `protobuf:"varint,5,opt,name=volume_total,json=volumeTotal" json:"volume_total,omitempty"`
	// Number of items remaining
	VolumeRemain uint64 `protobuf:"varint,6,opt,name=volume_remain,json=volumeRemain" json:"volume_remain,omitempty"`
	// Minimum volume to be traded for this order
	MinVolume uint64 `protobuf:"varint,7,opt,name=min_volume,json=minVolume" json:"min_volume,omitempty"`
	// The price the type is bought/sold for
	Price float64 `protobuf:"fixed64,8,opt,name=price" json:"price,omitempty"`
	// True: Bid/buy order | False: ask/sell order
	IsBuyOrder bool `protobuf:"varint,9,opt,name=is_buy_order,json=isBuyOrder" json:"is_buy_order,omitempty"`
	// Defines how long the order exists after creation
	Duration int32 `protobuf:"varint,10,opt,name=duration" json:"duration,omitempty"`
	// Order's range
	Range string `protobuf:"bytes,11,opt,name=range" json:"range,omitempty"`
	// Date the order was issued
	Issued *google_protobuf2.Timestamp `protobuf:"bytes,12,opt,name=issued" json:"issued,omitempty"`
	// When the order was last seen in this state: when using market-streamer the last-modified date from EMDRToNSQService
	SeenAt *google_protobuf2.Timestamp `protobuf:"bytes,13,opt,name=seen_at,json=seenAt" json:"seen_at,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Order) GetOrderId() uint64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *Order) GetTypeId() uint64 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

func (m *Order) GetRegionId() uint64 {
	if m != nil {
		return m.RegionId
	}
	return 0
}

func (m *Order) GetLocationId() uint64 {
	if m != nil {
		return m.LocationId
	}
	return 0
}

func (m *Order) GetVolumeTotal() uint64 {
	if m != nil {
		return m.VolumeTotal
	}
	return 0
}

func (m *Order) GetVolumeRemain() uint64 {
	if m != nil {
		return m.VolumeRemain
	}
	return 0
}

func (m *Order) GetMinVolume() uint64 {
	if m != nil {
		return m.MinVolume
	}
	return 0
}

func (m *Order) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Order) GetIsBuyOrder() bool {
	if m != nil {
		return m.IsBuyOrder
	}
	return false
}

func (m *Order) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Order) GetRange() string {
	if m != nil {
		return m.Range
	}
	return ""
}

func (m *Order) GetIssued() *google_protobuf2.Timestamp {
	if m != nil {
		return m.Issued
	}
	return nil
}

func (m *Order) GetSeenAt() *google_protobuf2.Timestamp {
	if m != nil {
		return m.SeenAt
	}
	return nil
}

func init() {
	proto.RegisterType((*GetOrderRequest)(nil), "esiMarkets.GetOrderRequest")
	proto.RegisterType((*GetRegionRequest)(nil), "esiMarkets.GetRegionRequest")
	proto.RegisterType((*GetTypeRequest)(nil), "esiMarkets.GetTypeRequest")
	proto.RegisterType((*GetRegionTypeRequest)(nil), "esiMarkets.GetRegionTypeRequest")
	proto.RegisterType((*GetOrdersResponse)(nil), "esiMarkets.GetOrdersResponse")
	proto.RegisterType((*GetRegionTypeUpdateStreamResponse)(nil), "esiMarkets.GetRegionTypeUpdateStreamResponse")
	proto.RegisterType((*RegionType)(nil), "esiMarkets.RegionType")
	proto.RegisterType((*Order)(nil), "esiMarkets.Order")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ESIMarkets service

type ESIMarketsClient interface {
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error)
	GetRegion(ctx context.Context, in *GetRegionRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error)
	GetType(ctx context.Context, in *GetTypeRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error)
	GetRegionType(ctx context.Context, in *GetRegionTypeRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error)
	GetRegionTypeUpdateStream(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (ESIMarkets_GetRegionTypeUpdateStreamClient, error)
}

type eSIMarketsClient struct {
	cc *grpc.ClientConn
}

func NewESIMarketsClient(cc *grpc.ClientConn) ESIMarketsClient {
	return &eSIMarketsClient{cc}
}

func (c *eSIMarketsClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error) {
	out := new(GetOrdersResponse)
	err := grpc.Invoke(ctx, "/esiMarkets.ESIMarkets/GetOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eSIMarketsClient) GetRegion(ctx context.Context, in *GetRegionRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error) {
	out := new(GetOrdersResponse)
	err := grpc.Invoke(ctx, "/esiMarkets.ESIMarkets/GetRegion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eSIMarketsClient) GetType(ctx context.Context, in *GetTypeRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error) {
	out := new(GetOrdersResponse)
	err := grpc.Invoke(ctx, "/esiMarkets.ESIMarkets/GetType", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eSIMarketsClient) GetRegionType(ctx context.Context, in *GetRegionTypeRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error) {
	out := new(GetOrdersResponse)
	err := grpc.Invoke(ctx, "/esiMarkets.ESIMarkets/GetRegionType", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eSIMarketsClient) GetRegionTypeUpdateStream(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (ESIMarkets_GetRegionTypeUpdateStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ESIMarkets_serviceDesc.Streams[0], c.cc, "/esiMarkets.ESIMarkets/GetRegionTypeUpdateStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &eSIMarketsGetRegionTypeUpdateStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ESIMarkets_GetRegionTypeUpdateStreamClient interface {
	Recv() (*GetRegionTypeUpdateStreamResponse, error)
	grpc.ClientStream
}

type eSIMarketsGetRegionTypeUpdateStreamClient struct {
	grpc.ClientStream
}

func (x *eSIMarketsGetRegionTypeUpdateStreamClient) Recv() (*GetRegionTypeUpdateStreamResponse, error) {
	m := new(GetRegionTypeUpdateStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ESIMarkets service

type ESIMarketsServer interface {
	GetOrder(context.Context, *GetOrderRequest) (*GetOrdersResponse, error)
	GetRegion(context.Context, *GetRegionRequest) (*GetOrdersResponse, error)
	GetType(context.Context, *GetTypeRequest) (*GetOrdersResponse, error)
	GetRegionType(context.Context, *GetRegionTypeRequest) (*GetOrdersResponse, error)
	GetRegionTypeUpdateStream(*google_protobuf1.Empty, ESIMarkets_GetRegionTypeUpdateStreamServer) error
}

func RegisterESIMarketsServer(s *grpc.Server, srv ESIMarketsServer) {
	s.RegisterService(&_ESIMarkets_serviceDesc, srv)
}

func _ESIMarkets_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ESIMarketsServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/esiMarkets.ESIMarkets/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ESIMarketsServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ESIMarkets_GetRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ESIMarketsServer).GetRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/esiMarkets.ESIMarkets/GetRegion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ESIMarketsServer).GetRegion(ctx, req.(*GetRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ESIMarkets_GetType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ESIMarketsServer).GetType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/esiMarkets.ESIMarkets/GetType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ESIMarketsServer).GetType(ctx, req.(*GetTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ESIMarkets_GetRegionType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegionTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ESIMarketsServer).GetRegionType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/esiMarkets.ESIMarkets/GetRegionType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ESIMarketsServer).GetRegionType(ctx, req.(*GetRegionTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ESIMarkets_GetRegionTypeUpdateStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(google_protobuf1.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ESIMarketsServer).GetRegionTypeUpdateStream(m, &eSIMarketsGetRegionTypeUpdateStreamServer{stream})
}

type ESIMarkets_GetRegionTypeUpdateStreamServer interface {
	Send(*GetRegionTypeUpdateStreamResponse) error
	grpc.ServerStream
}

type eSIMarketsGetRegionTypeUpdateStreamServer struct {
	grpc.ServerStream
}

func (x *eSIMarketsGetRegionTypeUpdateStreamServer) Send(m *GetRegionTypeUpdateStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _ESIMarkets_serviceDesc = grpc.ServiceDesc{
	ServiceName: "esiMarkets.ESIMarkets",
	HandlerType: (*ESIMarketsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrder",
			Handler:    _ESIMarkets_GetOrder_Handler,
		},
		{
			MethodName: "GetRegion",
			Handler:    _ESIMarkets_GetRegion_Handler,
		},
		{
			MethodName: "GetType",
			Handler:    _ESIMarkets_GetType_Handler,
		},
		{
			MethodName: "GetRegionType",
			Handler:    _ESIMarkets_GetRegionType_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetRegionTypeUpdateStream",
			Handler:       _ESIMarkets_GetRegionTypeUpdateStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "esiMarkets.proto",
}

func init() { proto.RegisterFile("esiMarkets.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 670 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcb, 0x6e, 0xd3, 0x5c,
	0x10, 0x96, 0xdb, 0xe6, 0x36, 0x49, 0xff, 0xbf, 0x1d, 0x55, 0xc5, 0x75, 0x5a, 0xd5, 0x35, 0x15,
	0x72, 0x51, 0x89, 0x21, 0x15, 0x0b, 0x84, 0x84, 0x44, 0xa5, 0xaa, 0x8a, 0x04, 0x42, 0x72, 0x0b,
	0x4b, 0x22, 0xb7, 0x1e, 0xa2, 0x23, 0xe2, 0x0b, 0x3e, 0xc7, 0x95, 0xa2, 0xaa, 0x9b, 0x6e, 0x58,
	0xb1, 0xe2, 0x51, 0x78, 0x14, 0x5e, 0x81, 0x07, 0x41, 0x3e, 0xbe, 0xc4, 0x71, 0x7a, 0x01, 0x56,
	0xc9, 0xcc, 0x7c, 0xe7, 0xfb, 0x66, 0x8e, 0xbf, 0x39, 0xb0, 0x42, 0x9c, 0xbd, 0x75, 0xa2, 0xcf,
	0x24, 0x78, 0x2f, 0x8c, 0x02, 0x11, 0x20, 0x4c, 0x33, 0xda, 0xe6, 0x28, 0x08, 0x46, 0x63, 0xb2,
	0x9c, 0x90, 0x59, 0x8e, 0xef, 0x07, 0xc2, 0x11, 0x2c, 0xf0, 0x33, 0xa4, 0xd6, 0xcd, 0xaa, 0x32,
	0x3a, 0x8b, 0x3f, 0x59, 0xe4, 0x85, 0x62, 0x92, 0x15, 0xb7, 0xab, 0x45, 0xc1, 0x3c, 0xe2, 0xc2,
	0xf1, 0xc2, 0x14, 0x60, 0xec, 0xc3, 0xff, 0xc7, 0x24, 0xde, 0x45, 0x2e, 0x45, 0x36, 0x7d, 0x89,
	0x89, 0x0b, 0xdc, 0x80, 0x66, 0x90, 0xc4, 0x43, 0xe6, 0xaa, 0x8a, 0xae, 0x98, 0x4b, 0x76, 0x43,
	0xc6, 0x03, 0xd7, 0xb0, 0x60, 0xe5, 0x98, 0x84, 0x4d, 0x23, 0x16, 0xf8, 0x39, 0xbc, 0x0b, 0xad,
	0x48, 0x26, 0xa6, 0xf8, 0x66, 0x9a, 0x18, 0xb8, 0xc6, 0x1e, 0xfc, 0x77, 0x4c, 0xe2, 0x74, 0x12,
	0x52, 0x0e, 0x7f, 0x00, 0x0d, 0x31, 0x09, 0x69, 0x0a, 0xae, 0x27, 0xe1, 0xc0, 0x35, 0xde, 0xc0,
	0x5a, 0xc1, 0x5d, 0x3e, 0x70, 0x17, 0x7f, 0x99, 0x6d, 0x61, 0x86, 0xed, 0x15, 0xac, 0xe6, 0x73,
	0x71, 0x9b, 0x78, 0x18, 0xf8, 0x9c, 0x70, 0x0f, 0xea, 0x72, 0x12, 0xae, 0x2a, 0xfa, 0xa2, 0xd9,
	0xee, 0xaf, 0xf6, 0x4a, 0xf7, 0x9e, 0xde, 0x41, 0x06, 0x30, 0x3e, 0xc2, 0xce, 0x4c, 0x37, 0xef,
	0x43, 0xd7, 0x11, 0x74, 0x22, 0x22, 0x72, 0xbc, 0x82, 0xef, 0x05, 0x74, 0xb2, 0xd6, 0x12, 0xd5,
	0x9c, 0x75, 0xbd, 0xcc, 0x5a, 0x9a, 0xa7, 0x1d, 0x15, 0xff, 0xb9, 0x71, 0x08, 0x30, 0x2d, 0xfd,
	0xe3, 0x8c, 0x3f, 0x16, 0xa1, 0x26, 0xbb, 0xbe, 0xe3, 0x93, 0xdd, 0x7a, 0x7a, 0x56, 0x73, 0xb1,
	0xa2, 0xb9, 0x0d, 0xed, 0x71, 0x70, 0x2e, 0x7d, 0x96, 0x94, 0x97, 0x64, 0x19, 0xf2, 0xd4, 0xc0,
	0xc5, 0x1d, 0xe8, 0x5c, 0x04, 0xe3, 0xd8, 0xa3, 0xa1, 0x08, 0x84, 0x33, 0x56, 0x6b, 0x12, 0xd1,
	0x4e, 0x73, 0xa7, 0x49, 0x0a, 0x1f, 0xc2, 0x72, 0x06, 0x89, 0xc8, 0x73, 0x98, 0xaf, 0xd6, 0x25,
	0x26, 0x3b, 0x67, 0xcb, 0x1c, 0x6e, 0x01, 0x78, 0xcc, 0x1f, 0xa6, 0x39, 0xb5, 0x21, 0x11, 0x2d,
	0x8f, 0xf9, 0x1f, 0x64, 0x02, 0xd7, 0xa0, 0x16, 0x46, 0xec, 0x9c, 0xd4, 0xa6, 0xae, 0x98, 0x8a,
	0x9d, 0x06, 0xa8, 0x43, 0x87, 0xf1, 0xe1, 0x59, 0x3c, 0x19, 0xca, 0x29, 0xd5, 0x96, 0xae, 0x98,
	0x4d, 0x1b, 0x18, 0x3f, 0x8c, 0x27, 0xe9, 0x85, 0x68, 0xd0, 0x74, 0xe3, 0x48, 0x36, 0xab, 0x82,
	0xae, 0x98, 0x35, 0xbb, 0x88, 0x13, 0xce, 0xc8, 0xf1, 0x47, 0xa4, 0xb6, 0x75, 0xc5, 0x6c, 0xd9,
	0x69, 0x80, 0x7d, 0xa8, 0x33, 0xce, 0x63, 0x72, 0xd5, 0x8e, 0xae, 0x98, 0xed, 0xbe, 0xd6, 0x4b,
	0x57, 0xa7, 0x97, 0xaf, 0x4e, 0xef, 0x34, 0x5f, 0x1d, 0x3b, 0x43, 0xe2, 0x01, 0x34, 0x38, 0x91,
	0x3f, 0x74, 0x84, 0xba, 0x7c, 0xff, 0xa1, 0x04, 0xfa, 0x5a, 0xf4, 0xaf, 0x6b, 0x00, 0x47, 0x27,
	0x83, 0xcc, 0x20, 0x18, 0x41, 0x33, 0x37, 0x2a, 0x76, 0xcb, 0xce, 0xa9, 0xac, 0xa5, 0xb6, 0x75,
	0x53, 0xb1, 0xf0, 0xb6, 0xf1, 0xf8, 0xfa, 0xe7, 0xaf, 0xef, 0x0b, 0xbb, 0x68, 0x58, 0x17, 0xcf,
	0x2c, 0x4f, 0xc2, 0xac, 0xd4, 0xcb, 0xe9, 0x8f, 0x75, 0x99, 0x3b, 0xe4, 0x0a, 0x2f, 0xa0, 0x55,
	0x98, 0x1b, 0x37, 0x2b, 0xbc, 0x33, 0xdb, 0x7d, 0x9f, 0xea, 0xbe, 0x54, 0x7d, 0x84, 0xbb, 0xf3,
	0xaa, 0xa9, 0x97, 0xac, 0xcb, 0xc2, 0x64, 0x57, 0xe8, 0x43, 0x23, 0x7b, 0x0d, 0x50, 0xab, 0xf0,
	0x96, 0x36, 0xfe, 0x3e, 0x4d, 0x53, 0x6a, 0x1a, 0xa8, 0xcf, 0x6b, 0x26, 0xd6, 0xb6, 0x2e, 0x33,
	0xbf, 0x5f, 0xe1, 0x37, 0x05, 0x96, 0x67, 0xb6, 0x18, 0xf5, 0x1b, 0x87, 0xfd, 0x0b, 0xf1, 0x97,
	0x52, 0xfc, 0x39, 0x1e, 0xfc, 0xc9, 0xc0, 0xd5, 0x7e, 0xbe, 0x2a, 0xb0, 0x71, 0xeb, 0xab, 0x82,
	0xeb, 0x73, 0xe6, 0x39, 0x4a, 0x5e, 0x72, 0xed, 0xc9, 0xad, 0x3d, 0xdf, 0xf4, 0x28, 0x19, 0x3b,
	0xb2, 0xc3, 0x2e, 0x6e, 0xcc, 0x77, 0x18, 0x4b, 0x3c, 0x7f, 0xaa, 0x9c, 0xd5, 0xa5, 0xc6, 0xc1,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x5d, 0x25, 0x99, 0x79, 0x06, 0x00, 0x00,
}