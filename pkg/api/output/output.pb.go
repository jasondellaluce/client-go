// Code generated by protoc-gen-go. DO NOT EDIT.
// source: output.proto

package output

import (
	context "context"
	fmt "fmt"
	schema "github.com/falcosecurity/client-go/pkg/api/schema"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

// The `request` message is the logical representation of the request model.
// It is the input of the `subscribe` service.
// It is used to configure the kind of subscription to the gRPC streaming server.
//
// By default the request asks to the server to only receive the accumulated events.
// In case you want to wait indefinitely for new events to come set the keepalive option to true.
type Request struct {
	Keepalive            bool     `protobuf:"varint,1,opt,name=keepalive,proto3" json:"keepalive,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b2b3ae2e703b013, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetKeepalive() bool {
	if m != nil {
		return m.Keepalive
	}
	return false
}

// The `response` message is the logical representation of the output model.
// It contains all the elements that Falco emits in an output along with the
// definitions for priorities and source.
type Response struct {
	Time                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Priority             schema.Priority      `protobuf:"varint,2,opt,name=priority,proto3,enum=falco.schema.Priority" json:"priority,omitempty"`
	Source               schema.Source        `protobuf:"varint,3,opt,name=source,proto3,enum=falco.schema.Source" json:"source,omitempty"`
	Rule                 string               `protobuf:"bytes,4,opt,name=rule,proto3" json:"rule,omitempty"`
	Output               string               `protobuf:"bytes,5,opt,name=output,proto3" json:"output,omitempty"`
	OutputFields         map[string]string    `protobuf:"bytes,6,rep,name=output_fields,json=outputFields,proto3" json:"output_fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Hostname             string               `protobuf:"bytes,7,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b2b3ae2e703b013, []int{1}
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

func (m *Response) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *Response) GetPriority() schema.Priority {
	if m != nil {
		return m.Priority
	}
	return schema.Priority_EMERGENCY
}

func (m *Response) GetSource() schema.Source {
	if m != nil {
		return m.Source
	}
	return schema.Source_SYSCALL
}

func (m *Response) GetRule() string {
	if m != nil {
		return m.Rule
	}
	return ""
}

func (m *Response) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *Response) GetOutputFields() map[string]string {
	if m != nil {
		return m.OutputFields
	}
	return nil
}

func (m *Response) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "falco.output.request")
	proto.RegisterType((*Response)(nil), "falco.output.response")
	proto.RegisterMapType((map[string]string)(nil), "falco.output.response.OutputFieldsEntry")
}

func init() { proto.RegisterFile("output.proto", fileDescriptor_0b2b3ae2e703b013) }

var fileDescriptor_0b2b3ae2e703b013 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xcf, 0xcb, 0xd3, 0x40,
	0x10, 0x25, 0x6d, 0xbf, 0x34, 0x99, 0xaf, 0x8a, 0x2e, 0xb5, 0x84, 0x20, 0x18, 0x7a, 0x31, 0x07,
	0xdd, 0x68, 0x7a, 0x11, 0x11, 0x04, 0x41, 0x3d, 0x89, 0xb0, 0x78, 0xf2, 0x22, 0xc9, 0x3a, 0x4d,
	0x97, 0x26, 0xd9, 0x75, 0x7f, 0x14, 0xfa, 0x7f, 0xfb, 0x07, 0x48, 0x77, 0xdb, 0xfa, 0xeb, 0xbb,
	0xcd, 0x9b, 0x79, 0x6f, 0xdf, 0xec, 0x1b, 0x58, 0x48, 0x67, 0x95, 0xb3, 0x54, 0x69, 0x69, 0x25,
	0x59, 0x6c, 0x9b, 0x9e, 0x4b, 0x1a, 0x7a, 0xf9, 0x93, 0x4e, 0xca, 0xae, 0xc7, 0xca, 0xcf, 0x5a,
	0xb7, 0xad, 0xac, 0x18, 0xd0, 0xd8, 0x66, 0x50, 0x81, 0x9e, 0x2f, 0x0c, 0xdf, 0xe1, 0xd0, 0x04,
	0xb4, 0x7e, 0x0a, 0x73, 0x8d, 0x3f, 0x1c, 0x1a, 0x4b, 0x1e, 0x43, 0xba, 0x47, 0x54, 0x4d, 0x2f,
	0x0e, 0x98, 0x45, 0x45, 0x54, 0x26, 0xec, 0x77, 0x63, 0xfd, 0x73, 0x02, 0x89, 0x46, 0xa3, 0xe4,
	0x68, 0x90, 0x50, 0x98, 0x9d, 0x9e, 0xf5, 0xac, 0xdb, 0x3a, 0xa7, 0xc1, 0x93, 0x5e, 0x3c, 0xe9,
	0x97, 0x8b, 0x27, 0xf3, 0x3c, 0x52, 0x43, 0xa2, 0xb4, 0x90, 0x5a, 0xd8, 0x63, 0x36, 0x29, 0xa2,
	0xf2, 0x7e, 0xbd, 0xa2, 0x61, 0xeb, 0xeb, 0x32, 0x61, 0xca, 0xae, 0x3c, 0xf2, 0x0c, 0x62, 0x23,
	0x9d, 0xe6, 0x98, 0x4d, 0xbd, 0x62, 0xf9, 0xb7, 0x22, 0xcc, 0xd8, 0x99, 0x43, 0x08, 0xcc, 0xb4,
	0xeb, 0x31, 0x9b, 0x15, 0x51, 0x99, 0x32, 0x5f, 0x93, 0x15, 0xc4, 0x21, 0x94, 0xec, 0xc6, 0x77,
	0xcf, 0x88, 0x7c, 0x82, 0x7b, 0xa1, 0xfa, 0xb6, 0x15, 0xd8, 0x7f, 0x37, 0x59, 0x5c, 0x4c, 0xcb,
	0xdb, 0xba, 0xa4, 0x7f, 0x06, 0x49, 0x2f, 0x9f, 0xa5, 0x9f, 0x3d, 0xfe, 0xe0, 0xa9, 0xef, 0x47,
	0xab, 0x8f, 0xec, 0x9c, 0x7f, 0x68, 0x91, 0x1c, 0x92, 0x9d, 0x34, 0x76, 0x6c, 0x06, 0xcc, 0xe6,
	0xde, 0xe8, 0x8a, 0xf3, 0xb7, 0xf0, 0xf0, 0x3f, 0x39, 0x79, 0x00, 0xd3, 0x3d, 0x1e, 0x7d, 0x78,
	0x29, 0x3b, 0x95, 0x64, 0x09, 0x37, 0x87, 0xa6, 0x77, 0xe8, 0xc3, 0x49, 0x59, 0x00, 0xaf, 0x27,
	0xaf, 0xa2, 0xfa, 0x23, 0xcc, 0x0d, 0xea, 0x83, 0xe0, 0x48, 0xde, 0x40, 0x6a, 0x5c, 0x6b, 0xb8,
	0x16, 0x2d, 0x92, 0x47, 0xff, 0x2e, 0xeb, 0x6f, 0x98, 0xaf, 0xee, 0xfe, 0xc3, 0x8b, 0xe8, 0xdd,
	0xe6, 0xeb, 0xcb, 0x4e, 0xd8, 0x9d, 0x6b, 0x29, 0x97, 0x43, 0xe5, 0x59, 0x06, 0xb9, 0x3b, 0x45,
	0x5d, 0xf1, 0x5e, 0xe0, 0x68, 0x9f, 0x77, 0xb2, 0x52, 0xfb, 0xae, 0x6a, 0x94, 0xa8, 0x82, 0xbe,
	0x8d, 0xfd, 0x45, 0x37, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x74, 0x41, 0x5b, 0x76, 0x71, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	Subscribe(ctx context.Context, in *Request, opts ...grpc.CallOption) (Service_SubscribeClient, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Subscribe(ctx context.Context, in *Request, opts ...grpc.CallOption) (Service_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Service_serviceDesc.Streams[0], "/falco.output.service/subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Service_SubscribeClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type serviceSubscribeClient struct {
	grpc.ClientStream
}

func (x *serviceSubscribeClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	Subscribe(*Request, Service_SubscribeServer) error
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Subscribe(req *Request, srv Service_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServer).Subscribe(m, &serviceSubscribeServer{stream})
}

type Service_SubscribeServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type serviceSubscribeServer struct {
	grpc.ServerStream
}

func (x *serviceSubscribeServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "falco.output.service",
	HandlerType: (*ServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "subscribe",
			Handler:       _Service_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "output.proto",
}
