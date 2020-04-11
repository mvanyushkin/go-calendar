// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package apicontracts

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type EventDto struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Time                 int64    `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventDto) Reset()         { *m = EventDto{} }
func (m *EventDto) String() string { return proto.CompactTextString(m) }
func (*EventDto) ProtoMessage()    {}
func (*EventDto) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *EventDto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventDto.Unmarshal(m, b)
}
func (m *EventDto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventDto.Marshal(b, m, deterministic)
}
func (m *EventDto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventDto.Merge(m, src)
}
func (m *EventDto) XXX_Size() int {
	return xxx_messageInfo_EventDto.Size(m)
}
func (m *EventDto) XXX_DiscardUnknown() {
	xxx_messageInfo_EventDto.DiscardUnknown(m)
}

var xxx_messageInfo_EventDto proto.InternalMessageInfo

func (m *EventDto) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EventDto) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *EventDto) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EventDto) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type CreateResponseDto struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponseDto) Reset()         { *m = CreateResponseDto{} }
func (m *CreateResponseDto) String() string { return proto.CompactTextString(m) }
func (*CreateResponseDto) ProtoMessage()    {}
func (*CreateResponseDto) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *CreateResponseDto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponseDto.Unmarshal(m, b)
}
func (m *CreateResponseDto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponseDto.Marshal(b, m, deterministic)
}
func (m *CreateResponseDto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponseDto.Merge(m, src)
}
func (m *CreateResponseDto) XXX_Size() int {
	return xxx_messageInfo_CreateResponseDto.Size(m)
}
func (m *CreateResponseDto) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponseDto.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponseDto proto.InternalMessageInfo

func (m *CreateResponseDto) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type EventsResponse struct {
	Events               []*EventDto `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *EventsResponse) Reset()         { *m = EventsResponse{} }
func (m *EventsResponse) String() string { return proto.CompactTextString(m) }
func (*EventsResponse) ProtoMessage()    {}
func (*EventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *EventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventsResponse.Unmarshal(m, b)
}
func (m *EventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventsResponse.Marshal(b, m, deterministic)
}
func (m *EventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventsResponse.Merge(m, src)
}
func (m *EventsResponse) XXX_Size() int {
	return xxx_messageInfo_EventsResponse.Size(m)
}
func (m *EventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventsResponse proto.InternalMessageInfo

func (m *EventsResponse) GetEvents() []*EventDto {
	if m != nil {
		return m.Events
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type DateRequest struct {
	Day                  int64    `protobuf:"varint,1,opt,name=day,proto3" json:"day,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DateRequest) Reset()         { *m = DateRequest{} }
func (m *DateRequest) String() string { return proto.CompactTextString(m) }
func (*DateRequest) ProtoMessage()    {}
func (*DateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *DateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateRequest.Unmarshal(m, b)
}
func (m *DateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateRequest.Marshal(b, m, deterministic)
}
func (m *DateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateRequest.Merge(m, src)
}
func (m *DateRequest) XXX_Size() int {
	return xxx_messageInfo_DateRequest.Size(m)
}
func (m *DateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DateRequest proto.InternalMessageInfo

func (m *DateRequest) GetDay() int64 {
	if m != nil {
		return m.Day
	}
	return 0
}

func init() {
	proto.RegisterType((*EventDto)(nil), "apicontracts.EventDto")
	proto.RegisterType((*CreateResponseDto)(nil), "apicontracts.CreateResponseDto")
	proto.RegisterType((*EventsResponse)(nil), "apicontracts.EventsResponse")
	proto.RegisterType((*Empty)(nil), "apicontracts.Empty")
	proto.RegisterType((*DateRequest)(nil), "apicontracts.DateRequest")
}

func init() {
	proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626)
}

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0x4b, 0xf3, 0x40,
	0x10, 0x80, 0x49, 0xd2, 0xe6, 0xed, 0x3b, 0xc1, 0x62, 0x87, 0x82, 0xa1, 0x17, 0x43, 0x44, 0xc9,
	0xa9, 0x6a, 0x45, 0xf0, 0xde, 0x56, 0x4f, 0x5e, 0x16, 0xc4, 0x73, 0x6c, 0x46, 0x5d, 0x6c, 0xb3,
	0xeb, 0xee, 0x58, 0xe8, 0x0f, 0xf2, 0x7f, 0x4a, 0x36, 0x0d, 0x06, 0xad, 0xe8, 0x6d, 0xbe, 0x9e,
	0x99, 0x27, 0x61, 0x61, 0xcf, 0x92, 0x59, 0xcb, 0x05, 0x8d, 0xb5, 0x51, 0xac, 0xb0, 0xf3, 0x64,
	0xf4, 0x22, 0x7d, 0x84, 0xde, 0x7c, 0x4d, 0x25, 0xcf, 0x58, 0x61, 0x1f, 0x7c, 0x59, 0xc4, 0x5e,
	0xe2, 0x65, 0x5d, 0xe1, 0xcb, 0x02, 0x87, 0xd0, 0x65, 0xc9, 0x4b, 0x8a, 0xfd, 0xc4, 0xcb, 0xfe,
	0x8b, 0x3a, 0xc1, 0x04, 0xa2, 0x82, 0xec, 0xc2, 0x48, 0xcd, 0x52, 0x95, 0x71, 0xe0, 0x7a, 0xed,
	0x12, 0x22, 0x74, 0x58, 0xae, 0x28, 0xee, 0x24, 0x5e, 0x16, 0x08, 0x17, 0xa7, 0x47, 0x30, 0x98,
	0x1a, 0xca, 0x99, 0x04, 0x59, 0xad, 0x4a, 0x4b, 0x3b, 0x0e, 0xa6, 0x57, 0xd0, 0x77, 0x32, 0xb6,
	0x19, 0xc2, 0x13, 0x08, 0xc9, 0x55, 0x62, 0x2f, 0x09, 0xb2, 0x68, 0xd2, 0x1f, 0x57, 0xd6, 0xe3,
	0x46, 0x59, 0x6c, 0xbb, 0xe9, 0x3f, 0xe8, 0xce, 0x57, 0x9a, 0x37, 0xe9, 0x21, 0x44, 0x33, 0x77,
	0xe5, 0xf5, 0x8d, 0x2c, 0xe3, 0x3e, 0x04, 0x45, 0xbe, 0x71, 0x27, 0x02, 0x51, 0x85, 0x93, 0x77,
	0x1f, 0x7a, 0xd3, 0x7c, 0x49, 0x65, 0x91, 0x1b, 0x3c, 0x87, 0xb0, 0xb6, 0xc2, 0x2f, 0x8b, 0x47,
	0x07, 0x75, 0xfe, 0xdd, 0xf9, 0x18, 0xc2, 0x3b, 0x5d, 0xec, 0x42, 0xa2, 0x6d, 0x5e, 0x79, 0x54,
	0x63, 0x82, 0x56, 0x6a, 0xfd, 0xcb, 0xd8, 0x25, 0xc0, 0x0d, 0xf1, 0xb5, 0x32, 0x95, 0x34, 0x0e,
	0xea, 0x56, 0xeb, 0x03, 0x46, 0xc3, 0x16, 0xfd, 0xf9, 0x5b, 0x4e, 0x1b, 0xec, 0x9e, 0xe8, 0x05,
	0xdb, 0x1b, 0x7f, 0x00, 0xce, 0x20, 0xaa, 0x81, 0x5b, 0x55, 0xf2, 0xf3, 0x1f, 0x88, 0x87, 0xd0,
	0xbd, 0x92, 0x8b, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x29, 0x26, 0xc3, 0xe9, 0x36, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the apicontracts package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalendarClient is the client API for Calendar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalendarClient interface {
	Create(ctx context.Context, in *EventDto, opts ...grpc.CallOption) (*CreateResponseDto, error)
	Update(ctx context.Context, in *EventDto, opts ...grpc.CallOption) (*Empty, error)
	Remove(ctx context.Context, in *EventDto, opts ...grpc.CallOption) (*Empty, error)
	GetForDate(ctx context.Context, in *DateRequest, opts ...grpc.CallOption) (*EventsResponse, error)
	GetForWeek(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*EventsResponse, error)
	GetForMonth(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*EventsResponse, error)
}

type calendarClient struct {
	cc grpc.ClientConnInterface
}

func NewCalendarClient(cc grpc.ClientConnInterface) CalendarClient {
	return &calendarClient{cc}
}

func (c *calendarClient) Create(ctx context.Context, in *EventDto, opts ...grpc.CallOption) (*CreateResponseDto, error) {
	out := new(CreateResponseDto)
	err := c.cc.Invoke(ctx, "/apicontracts.Calendar/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) Update(ctx context.Context, in *EventDto, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/apicontracts.Calendar/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) Remove(ctx context.Context, in *EventDto, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/apicontracts.Calendar/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) GetForDate(ctx context.Context, in *DateRequest, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/apicontracts.Calendar/GetForDate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) GetForWeek(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/apicontracts.Calendar/GetForWeek", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) GetForMonth(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/apicontracts.Calendar/GetForMonth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServer is the server API for Calendar service.
type CalendarServer interface {
	Create(context.Context, *EventDto) (*CreateResponseDto, error)
	Update(context.Context, *EventDto) (*Empty, error)
	Remove(context.Context, *EventDto) (*Empty, error)
	GetForDate(context.Context, *DateRequest) (*EventsResponse, error)
	GetForWeek(context.Context, *Empty) (*EventsResponse, error)
	GetForMonth(context.Context, *Empty) (*EventsResponse, error)
}

// UnimplementedCalendarServer can be embedded to have forward compatible implementations.
type UnimplementedCalendarServer struct {
}

func (*UnimplementedCalendarServer) Create(ctx context.Context, req *EventDto) (*CreateResponseDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedCalendarServer) Update(ctx context.Context, req *EventDto) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedCalendarServer) Remove(ctx context.Context, req *EventDto) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (*UnimplementedCalendarServer) GetForDate(ctx context.Context, req *DateRequest) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetForDate not implemented")
}
func (*UnimplementedCalendarServer) GetForWeek(ctx context.Context, req *Empty) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetForWeek not implemented")
}
func (*UnimplementedCalendarServer) GetForMonth(ctx context.Context, req *Empty) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetForMonth not implemented")
}

func RegisterCalendarServer(s *grpc.Server, srv CalendarServer) {
	s.RegisterService(&_Calendar_serviceDesc, srv)
}

func _Calendar_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apicontracts.Calendar/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).Create(ctx, req.(*EventDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apicontracts.Calendar/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).Update(ctx, req.(*EventDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apicontracts.Calendar/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).Remove(ctx, req.(*EventDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_GetForDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).GetForDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apicontracts.Calendar/GetForDate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).GetForDate(ctx, req.(*DateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_GetForWeek_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).GetForWeek(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apicontracts.Calendar/GetForWeek",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).GetForWeek(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_GetForMonth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).GetForMonth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apicontracts.Calendar/GetForMonth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).GetForMonth(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calendar_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apicontracts.Calendar",
	HandlerType: (*CalendarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Calendar_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Calendar_Update_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Calendar_Remove_Handler,
		},
		{
			MethodName: "GetForDate",
			Handler:    _Calendar_GetForDate_Handler,
		},
		{
			MethodName: "GetForWeek",
			Handler:    _Calendar_GetForWeek_Handler,
		},
		{
			MethodName: "GetForMonth",
			Handler:    _Calendar_GetForMonth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}