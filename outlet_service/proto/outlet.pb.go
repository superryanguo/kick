// Code generated by protoc-gen-go. DO NOT EDIT.
// source: outlet.proto

package outlet

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Order struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	SumPrice             int32        `protobuf:"varint,3,opt,name=sum_price,json=sumPrice,proto3" json:"sum_price,omitempty"`
	UserId               string       `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CourierId            string       `protobuf:"bytes,5,opt,name=courier_id,json=courierId,proto3" json:"courier_id,omitempty"`
	Commoditys           []*Commodity `protobuf:"bytes,6,rep,name=commoditys,proto3" json:"commoditys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5d3308bacd7807a, []int{0}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Order) GetSumPrice() int32 {
	if m != nil {
		return m.SumPrice
	}
	return 0
}

func (m *Order) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Order) GetCourierId() string {
	if m != nil {
		return m.CourierId
	}
	return ""
}

func (m *Order) GetCommoditys() []*Commodity {
	if m != nil {
		return m.Commoditys
	}
	return nil
}

type Commodity struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Price                int32    `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	Quantity             int32    `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Commodity) Reset()         { *m = Commodity{} }
func (m *Commodity) String() string { return proto.CompactTextString(m) }
func (*Commodity) ProtoMessage()    {}
func (*Commodity) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5d3308bacd7807a, []int{1}
}

func (m *Commodity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Commodity.Unmarshal(m, b)
}
func (m *Commodity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Commodity.Marshal(b, m, deterministic)
}
func (m *Commodity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commodity.Merge(m, src)
}
func (m *Commodity) XXX_Size() int {
	return xxx_messageInfo_Commodity.Size(m)
}
func (m *Commodity) XXX_DiscardUnknown() {
	xxx_messageInfo_Commodity.DiscardUnknown(m)
}

var xxx_messageInfo_Commodity proto.InternalMessageInfo

func (m *Commodity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Commodity) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Commodity) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Commodity) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5d3308bacd7807a, []int{2}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type Response struct {
	Created              bool     `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Orders               []*Order `protobuf:"bytes,2,rep,name=orders,proto3" json:"orders,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5d3308bacd7807a, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

func init() {
	proto.RegisterType((*Order)(nil), "outlet.Order")
	proto.RegisterType((*Commodity)(nil), "outlet.Commodity")
	proto.RegisterType((*GetRequest)(nil), "outlet.GetRequest")
	proto.RegisterType((*Response)(nil), "outlet.Response")
}

func init() { proto.RegisterFile("outlet.proto", fileDescriptor_e5d3308bacd7807a) }

var fileDescriptor_e5d3308bacd7807a = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcd, 0x6a, 0xf2, 0x40,
	0x14, 0x86, 0xbf, 0xc4, 0x2f, 0x31, 0x39, 0x6a, 0x69, 0x0f, 0x85, 0x0e, 0x96, 0x42, 0x08, 0x14,
	0x5c, 0x49, 0xab, 0x97, 0xe0, 0x42, 0xa4, 0x0b, 0xcb, 0xf4, 0x02, 0xc4, 0x66, 0xce, 0x62, 0xa0,
	0x71, 0xe2, 0xfc, 0x08, 0x5e, 0x5e, 0xef, 0xac, 0x64, 0xe2, 0x58, 0x0b, 0x5d, 0x75, 0xf9, 0x9c,
	0x67, 0x4e, 0x78, 0xdf, 0xc9, 0xc0, 0x50, 0x39, 0xfb, 0x41, 0x76, 0xda, 0x68, 0x65, 0x15, 0xa6,
	0x1d, 0x95, 0x9f, 0x11, 0x24, 0x6b, 0x2d, 0x48, 0xe3, 0x15, 0xc4, 0x52, 0xb0, 0xa8, 0x88, 0x26,
	0x39, 0x8f, 0xa5, 0xc0, 0x02, 0x06, 0x82, 0x4c, 0xa5, 0x65, 0x63, 0xa5, 0xda, 0xb1, 0xd8, 0x8b,
	0xcb, 0x11, 0xde, 0x43, 0x6e, 0x5c, 0xbd, 0x69, 0xb4, 0xac, 0x88, 0xf5, 0x8a, 0x68, 0x92, 0xf0,
	0xcc, 0xb8, 0xfa, 0xb5, 0x65, 0xbc, 0x83, 0xbe, 0x33, 0xa4, 0x37, 0x52, 0xb0, 0xff, 0x7e, 0x35,
	0x6d, 0x71, 0x25, 0xf0, 0x01, 0xa0, 0x52, 0x4e, 0xcb, 0xce, 0x25, 0xde, 0xe5, 0xa7, 0xc9, 0x4a,
	0xe0, 0x73, 0xab, 0xeb, 0x5a, 0x09, 0x69, 0x8f, 0x86, 0xa5, 0x45, 0x6f, 0x32, 0x98, 0xdd, 0x4c,
	0x4f, 0xd9, 0x17, 0xc1, 0xf0, 0x8b, 0x43, 0xa5, 0x82, 0xfc, 0x2c, 0xfe, 0x50, 0xe3, 0x16, 0x92,
	0xcb, 0x0a, 0x1d, 0xe0, 0x18, 0xb2, 0xbd, 0xdb, 0xee, 0xac, 0xb4, 0x47, 0x5f, 0x20, 0xe1, 0x67,
	0x2e, 0x87, 0x00, 0x4b, 0xb2, 0x9c, 0xf6, 0x8e, 0x8c, 0x2d, 0x5f, 0x20, 0xe3, 0x64, 0x1a, 0xb5,
	0x33, 0x84, 0x0c, 0xfa, 0x95, 0xa6, 0xad, 0xa5, 0x2e, 0x42, 0xc6, 0x03, 0xe2, 0x23, 0xa4, 0xaa,
	0xbd, 0x67, 0xc3, 0x62, 0xdf, 0x69, 0x14, 0x3a, 0xf9, 0xdb, 0xe7, 0x27, 0x39, 0x3b, 0xc0, 0x68,
	0xed, 0xe7, 0x6f, 0xa4, 0x0f, 0x6d, 0x8e, 0x27, 0x18, 0x2c, 0xfc, 0x27, 0xba, 0xbf, 0xf4, 0x73,
	0x6d, 0x7c, 0x1d, 0x30, 0x24, 0x28, 0xff, 0xe1, 0x1c, 0xf2, 0x25, 0x59, 0xef, 0x0d, 0x62, 0x38,
	0xf0, 0x1d, 0xf8, 0xb7, 0xa5, 0xf7, 0xd4, 0x3f, 0x8b, 0xf9, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xd9, 0xdf, 0xa2, 0x8c, 0x26, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for OutletService service

type OutletServiceClient interface {
	CreateOrder(ctx context.Context, in *Order, opts ...client.CallOption) (*Response, error)
	GetOrders(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error)
}

type outletServiceClient struct {
	c           client.Client
	serviceName string
}

func NewOutletServiceClient(serviceName string, c client.Client) OutletServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "outlet"
	}
	return &outletServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *outletServiceClient) CreateOrder(ctx context.Context, in *Order, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "OutletService.CreateOrder", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outletServiceClient) GetOrders(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "OutletService.GetOrders", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OutletService service

type OutletServiceHandler interface {
	CreateOrder(context.Context, *Order, *Response) error
	GetOrders(context.Context, *GetRequest, *Response) error
}

func RegisterOutletServiceHandler(s server.Server, hdlr OutletServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&OutletService{hdlr}, opts...))
}

type OutletService struct {
	OutletServiceHandler
}

func (h *OutletService) CreateOrder(ctx context.Context, in *Order, out *Response) error {
	return h.OutletServiceHandler.CreateOrder(ctx, in, out)
}

func (h *OutletService) GetOrders(ctx context.Context, in *GetRequest, out *Response) error {
	return h.OutletServiceHandler.GetOrders(ctx, in, out)
}
