// Code generated by protoc-gen-go. DO NOT EDIT.
// source: retrieve.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	retrieve.proto

It has these top-level messages:
	RetrievePara
	Retrieve
	RetrieveAction
	BackupRetrieve
	PrepareRetrieve
	PerformRetrieve
	CancelRetrieve
	ReqRetrieveInfo
	RetrieveQuery
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import types2 "github.com/33cn/chain33/types"

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

// message for retrieve start
type RetrievePara struct {
	DefaultAddress string `protobuf:"bytes,1,opt,name=defaultAddress" json:"defaultAddress,omitempty"`
	Status         int32  `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	CreateTime     int64  `protobuf:"varint,3,opt,name=createTime" json:"createTime,omitempty"`
	PrepareTime    int64  `protobuf:"varint,4,opt,name=prepareTime" json:"prepareTime,omitempty"`
	DelayPeriod    int64  `protobuf:"varint,5,opt,name=delayPeriod" json:"delayPeriod,omitempty"`
}

func (m *RetrievePara) Reset()                    { *m = RetrievePara{} }
func (m *RetrievePara) String() string            { return proto.CompactTextString(m) }
func (*RetrievePara) ProtoMessage()               {}
func (*RetrievePara) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RetrievePara) GetDefaultAddress() string {
	if m != nil {
		return m.DefaultAddress
	}
	return ""
}

func (m *RetrievePara) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *RetrievePara) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *RetrievePara) GetPrepareTime() int64 {
	if m != nil {
		return m.PrepareTime
	}
	return 0
}

func (m *RetrievePara) GetDelayPeriod() int64 {
	if m != nil {
		return m.DelayPeriod
	}
	return 0
}

type Retrieve struct {
	// used as key
	BackupAddress string          `protobuf:"bytes,1,opt,name=backupAddress" json:"backupAddress,omitempty"`
	RetPara       []*RetrievePara `protobuf:"bytes,2,rep,name=retPara" json:"retPara,omitempty"`
}

func (m *Retrieve) Reset()                    { *m = Retrieve{} }
func (m *Retrieve) String() string            { return proto.CompactTextString(m) }
func (*Retrieve) ProtoMessage()               {}
func (*Retrieve) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Retrieve) GetBackupAddress() string {
	if m != nil {
		return m.BackupAddress
	}
	return ""
}

func (m *Retrieve) GetRetPara() []*RetrievePara {
	if m != nil {
		return m.RetPara
	}
	return nil
}

type RetrieveAction struct {
	// Types that are valid to be assigned to Value:
	//	*RetrieveAction_Prepare
	//	*RetrieveAction_Perform
	//	*RetrieveAction_Backup
	//	*RetrieveAction_Cancel
	Value isRetrieveAction_Value `protobuf_oneof:"value"`
	Ty    int32                  `protobuf:"varint,5,opt,name=ty" json:"ty,omitempty"`
}

func (m *RetrieveAction) Reset()                    { *m = RetrieveAction{} }
func (m *RetrieveAction) String() string            { return proto.CompactTextString(m) }
func (*RetrieveAction) ProtoMessage()               {}
func (*RetrieveAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isRetrieveAction_Value interface {
	isRetrieveAction_Value()
}

type RetrieveAction_Prepare struct {
	Prepare *PrepareRetrieve `protobuf:"bytes,1,opt,name=prepare,oneof"`
}
type RetrieveAction_Perform struct {
	Perform *PerformRetrieve `protobuf:"bytes,2,opt,name=perform,oneof"`
}
type RetrieveAction_Backup struct {
	Backup *BackupRetrieve `protobuf:"bytes,3,opt,name=backup,oneof"`
}
type RetrieveAction_Cancel struct {
	Cancel *CancelRetrieve `protobuf:"bytes,4,opt,name=cancel,oneof"`
}

func (*RetrieveAction_Prepare) isRetrieveAction_Value() {}
func (*RetrieveAction_Perform) isRetrieveAction_Value() {}
func (*RetrieveAction_Backup) isRetrieveAction_Value()  {}
func (*RetrieveAction_Cancel) isRetrieveAction_Value()  {}

func (m *RetrieveAction) GetValue() isRetrieveAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *RetrieveAction) GetPrepare() *PrepareRetrieve {
	if x, ok := m.GetValue().(*RetrieveAction_Prepare); ok {
		return x.Prepare
	}
	return nil
}

func (m *RetrieveAction) GetPerform() *PerformRetrieve {
	if x, ok := m.GetValue().(*RetrieveAction_Perform); ok {
		return x.Perform
	}
	return nil
}

func (m *RetrieveAction) GetBackup() *BackupRetrieve {
	if x, ok := m.GetValue().(*RetrieveAction_Backup); ok {
		return x.Backup
	}
	return nil
}

func (m *RetrieveAction) GetCancel() *CancelRetrieve {
	if x, ok := m.GetValue().(*RetrieveAction_Cancel); ok {
		return x.Cancel
	}
	return nil
}

func (m *RetrieveAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RetrieveAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RetrieveAction_OneofMarshaler, _RetrieveAction_OneofUnmarshaler, _RetrieveAction_OneofSizer, []interface{}{
		(*RetrieveAction_Prepare)(nil),
		(*RetrieveAction_Perform)(nil),
		(*RetrieveAction_Backup)(nil),
		(*RetrieveAction_Cancel)(nil),
	}
}

func _RetrieveAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RetrieveAction)
	// value
	switch x := m.Value.(type) {
	case *RetrieveAction_Prepare:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Prepare); err != nil {
			return err
		}
	case *RetrieveAction_Perform:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Perform); err != nil {
			return err
		}
	case *RetrieveAction_Backup:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Backup); err != nil {
			return err
		}
	case *RetrieveAction_Cancel:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Cancel); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RetrieveAction.Value has unexpected type %T", x)
	}
	return nil
}

func _RetrieveAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RetrieveAction)
	switch tag {
	case 1: // value.prepare
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PrepareRetrieve)
		err := b.DecodeMessage(msg)
		m.Value = &RetrieveAction_Prepare{msg}
		return true, err
	case 2: // value.perform
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PerformRetrieve)
		err := b.DecodeMessage(msg)
		m.Value = &RetrieveAction_Perform{msg}
		return true, err
	case 3: // value.backup
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BackupRetrieve)
		err := b.DecodeMessage(msg)
		m.Value = &RetrieveAction_Backup{msg}
		return true, err
	case 4: // value.cancel
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CancelRetrieve)
		err := b.DecodeMessage(msg)
		m.Value = &RetrieveAction_Cancel{msg}
		return true, err
	default:
		return false, nil
	}
}

func _RetrieveAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RetrieveAction)
	// value
	switch x := m.Value.(type) {
	case *RetrieveAction_Prepare:
		s := proto.Size(x.Prepare)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RetrieveAction_Perform:
		s := proto.Size(x.Perform)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RetrieveAction_Backup:
		s := proto.Size(x.Backup)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RetrieveAction_Cancel:
		s := proto.Size(x.Cancel)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type BackupRetrieve struct {
	BackupAddress  string `protobuf:"bytes,1,opt,name=backupAddress" json:"backupAddress,omitempty"`
	DefaultAddress string `protobuf:"bytes,2,opt,name=defaultAddress" json:"defaultAddress,omitempty"`
	DelayPeriod    int64  `protobuf:"varint,3,opt,name=delayPeriod" json:"delayPeriod,omitempty"`
}

func (m *BackupRetrieve) Reset()                    { *m = BackupRetrieve{} }
func (m *BackupRetrieve) String() string            { return proto.CompactTextString(m) }
func (*BackupRetrieve) ProtoMessage()               {}
func (*BackupRetrieve) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BackupRetrieve) GetBackupAddress() string {
	if m != nil {
		return m.BackupAddress
	}
	return ""
}

func (m *BackupRetrieve) GetDefaultAddress() string {
	if m != nil {
		return m.DefaultAddress
	}
	return ""
}

func (m *BackupRetrieve) GetDelayPeriod() int64 {
	if m != nil {
		return m.DelayPeriod
	}
	return 0
}

type PrepareRetrieve struct {
	BackupAddress  string `protobuf:"bytes,1,opt,name=backupAddress" json:"backupAddress,omitempty"`
	DefaultAddress string `protobuf:"bytes,2,opt,name=defaultAddress" json:"defaultAddress,omitempty"`
}

func (m *PrepareRetrieve) Reset()                    { *m = PrepareRetrieve{} }
func (m *PrepareRetrieve) String() string            { return proto.CompactTextString(m) }
func (*PrepareRetrieve) ProtoMessage()               {}
func (*PrepareRetrieve) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PrepareRetrieve) GetBackupAddress() string {
	if m != nil {
		return m.BackupAddress
	}
	return ""
}

func (m *PrepareRetrieve) GetDefaultAddress() string {
	if m != nil {
		return m.DefaultAddress
	}
	return ""
}

type PerformRetrieve struct {
	BackupAddress  string `protobuf:"bytes,1,opt,name=backupAddress" json:"backupAddress,omitempty"`
	DefaultAddress string `protobuf:"bytes,2,opt,name=defaultAddress" json:"defaultAddress,omitempty"`
}

func (m *PerformRetrieve) Reset()                    { *m = PerformRetrieve{} }
func (m *PerformRetrieve) String() string            { return proto.CompactTextString(m) }
func (*PerformRetrieve) ProtoMessage()               {}
func (*PerformRetrieve) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PerformRetrieve) GetBackupAddress() string {
	if m != nil {
		return m.BackupAddress
	}
	return ""
}

func (m *PerformRetrieve) GetDefaultAddress() string {
	if m != nil {
		return m.DefaultAddress
	}
	return ""
}

type CancelRetrieve struct {
	BackupAddress  string `protobuf:"bytes,1,opt,name=backupAddress" json:"backupAddress,omitempty"`
	DefaultAddress string `protobuf:"bytes,2,opt,name=defaultAddress" json:"defaultAddress,omitempty"`
}

func (m *CancelRetrieve) Reset()                    { *m = CancelRetrieve{} }
func (m *CancelRetrieve) String() string            { return proto.CompactTextString(m) }
func (*CancelRetrieve) ProtoMessage()               {}
func (*CancelRetrieve) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CancelRetrieve) GetBackupAddress() string {
	if m != nil {
		return m.BackupAddress
	}
	return ""
}

func (m *CancelRetrieve) GetDefaultAddress() string {
	if m != nil {
		return m.DefaultAddress
	}
	return ""
}

type ReqRetrieveInfo struct {
	BackupAddress  string `protobuf:"bytes,1,opt,name=backupAddress" json:"backupAddress,omitempty"`
	DefaultAddress string `protobuf:"bytes,2,opt,name=defaultAddress" json:"defaultAddress,omitempty"`
}

func (m *ReqRetrieveInfo) Reset()                    { *m = ReqRetrieveInfo{} }
func (m *ReqRetrieveInfo) String() string            { return proto.CompactTextString(m) }
func (*ReqRetrieveInfo) ProtoMessage()               {}
func (*ReqRetrieveInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ReqRetrieveInfo) GetBackupAddress() string {
	if m != nil {
		return m.BackupAddress
	}
	return ""
}

func (m *ReqRetrieveInfo) GetDefaultAddress() string {
	if m != nil {
		return m.DefaultAddress
	}
	return ""
}

type RetrieveQuery struct {
	BackupAddress  string `protobuf:"bytes,1,opt,name=backupAddress" json:"backupAddress,omitempty"`
	DefaultAddress string `protobuf:"bytes,2,opt,name=defaultAddress" json:"defaultAddress,omitempty"`
	DelayPeriod    int64  `protobuf:"varint,3,opt,name=delayPeriod" json:"delayPeriod,omitempty"`
	PrepareTime    int64  `protobuf:"varint,4,opt,name=prepareTime" json:"prepareTime,omitempty"`
	RemainTime     int64  `protobuf:"varint,5,opt,name=remainTime" json:"remainTime,omitempty"`
	Status         int32  `protobuf:"varint,6,opt,name=status" json:"status,omitempty"`
}

func (m *RetrieveQuery) Reset()                    { *m = RetrieveQuery{} }
func (m *RetrieveQuery) String() string            { return proto.CompactTextString(m) }
func (*RetrieveQuery) ProtoMessage()               {}
func (*RetrieveQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RetrieveQuery) GetBackupAddress() string {
	if m != nil {
		return m.BackupAddress
	}
	return ""
}

func (m *RetrieveQuery) GetDefaultAddress() string {
	if m != nil {
		return m.DefaultAddress
	}
	return ""
}

func (m *RetrieveQuery) GetDelayPeriod() int64 {
	if m != nil {
		return m.DelayPeriod
	}
	return 0
}

func (m *RetrieveQuery) GetPrepareTime() int64 {
	if m != nil {
		return m.PrepareTime
	}
	return 0
}

func (m *RetrieveQuery) GetRemainTime() int64 {
	if m != nil {
		return m.RemainTime
	}
	return 0
}

func (m *RetrieveQuery) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*RetrievePara)(nil), "types.RetrievePara")
	proto.RegisterType((*Retrieve)(nil), "types.Retrieve")
	proto.RegisterType((*RetrieveAction)(nil), "types.RetrieveAction")
	proto.RegisterType((*BackupRetrieve)(nil), "types.BackupRetrieve")
	proto.RegisterType((*PrepareRetrieve)(nil), "types.PrepareRetrieve")
	proto.RegisterType((*PerformRetrieve)(nil), "types.PerformRetrieve")
	proto.RegisterType((*CancelRetrieve)(nil), "types.CancelRetrieve")
	proto.RegisterType((*ReqRetrieveInfo)(nil), "types.ReqRetrieveInfo")
	proto.RegisterType((*RetrieveQuery)(nil), "types.RetrieveQuery")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Retrieve service

type RetrieveClient interface {
	Prepare(ctx context.Context, in *PrepareRetrieve, opts ...grpc.CallOption) (*types2.UnsignTx, error)
	Perform(ctx context.Context, in *PerformRetrieve, opts ...grpc.CallOption) (*types2.UnsignTx, error)
	Backup(ctx context.Context, in *BackupRetrieve, opts ...grpc.CallOption) (*types2.UnsignTx, error)
	Cancel(ctx context.Context, in *CancelRetrieve, opts ...grpc.CallOption) (*types2.UnsignTx, error)
}

type retrieveClient struct {
	cc *grpc.ClientConn
}

func NewRetrieveClient(cc *grpc.ClientConn) RetrieveClient {
	return &retrieveClient{cc}
}

func (c *retrieveClient) Prepare(ctx context.Context, in *PrepareRetrieve, opts ...grpc.CallOption) (*types2.UnsignTx, error) {
	out := new(types2.UnsignTx)
	err := grpc.Invoke(ctx, "/types.retrieve/Prepare", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *retrieveClient) Perform(ctx context.Context, in *PerformRetrieve, opts ...grpc.CallOption) (*types2.UnsignTx, error) {
	out := new(types2.UnsignTx)
	err := grpc.Invoke(ctx, "/types.retrieve/Perform", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *retrieveClient) Backup(ctx context.Context, in *BackupRetrieve, opts ...grpc.CallOption) (*types2.UnsignTx, error) {
	out := new(types2.UnsignTx)
	err := grpc.Invoke(ctx, "/types.retrieve/Backup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *retrieveClient) Cancel(ctx context.Context, in *CancelRetrieve, opts ...grpc.CallOption) (*types2.UnsignTx, error) {
	out := new(types2.UnsignTx)
	err := grpc.Invoke(ctx, "/types.retrieve/Cancel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Retrieve service

type RetrieveServer interface {
	Prepare(context.Context, *PrepareRetrieve) (*types2.UnsignTx, error)
	Perform(context.Context, *PerformRetrieve) (*types2.UnsignTx, error)
	Backup(context.Context, *BackupRetrieve) (*types2.UnsignTx, error)
	Cancel(context.Context, *CancelRetrieve) (*types2.UnsignTx, error)
}

func RegisterRetrieveServer(s *grpc.Server, srv RetrieveServer) {
	s.RegisterService(&_Retrieve_serviceDesc, srv)
}

func _Retrieve_Prepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareRetrieve)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrieveServer).Prepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.retrieve/Prepare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrieveServer).Prepare(ctx, req.(*PrepareRetrieve))
	}
	return interceptor(ctx, in, info, handler)
}

func _Retrieve_Perform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PerformRetrieve)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrieveServer).Perform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.retrieve/Perform",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrieveServer).Perform(ctx, req.(*PerformRetrieve))
	}
	return interceptor(ctx, in, info, handler)
}

func _Retrieve_Backup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupRetrieve)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrieveServer).Backup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.retrieve/Backup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrieveServer).Backup(ctx, req.(*BackupRetrieve))
	}
	return interceptor(ctx, in, info, handler)
}

func _Retrieve_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelRetrieve)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrieveServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.retrieve/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrieveServer).Cancel(ctx, req.(*CancelRetrieve))
	}
	return interceptor(ctx, in, info, handler)
}

var _Retrieve_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.retrieve",
	HandlerType: (*RetrieveServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Prepare",
			Handler:    _Retrieve_Prepare_Handler,
		},
		{
			MethodName: "Perform",
			Handler:    _Retrieve_Perform_Handler,
		},
		{
			MethodName: "Backup",
			Handler:    _Retrieve_Backup_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _Retrieve_Cancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "retrieve.proto",
}

func init() { proto.RegisterFile("retrieve.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xee, 0x3a, 0xd8, 0x29, 0x13, 0xea, 0x8a, 0x45, 0x54, 0x56, 0x0f, 0x95, 0x65, 0x21, 0x94,
	0x0b, 0x41, 0x32, 0xbc, 0x40, 0xcb, 0x05, 0x6e, 0x61, 0x55, 0xae, 0xa0, 0xad, 0x3d, 0x41, 0x16,
	0x8e, 0x6d, 0xd6, 0xeb, 0x0a, 0xdf, 0x78, 0x26, 0xde, 0x86, 0x0b, 0x2f, 0xc1, 0x0b, 0xa0, 0xec,
	0x8f, 0x58, 0x07, 0x47, 0xca, 0xc1, 0xe2, 0xe8, 0xcf, 0xdf, 0xb7, 0x33, 0xfb, 0xed, 0xcc, 0x07,
	0xa1, 0x40, 0x29, 0x0a, 0xbc, 0xc7, 0x55, 0x23, 0x6a, 0x59, 0x53, 0x5f, 0xf6, 0x0d, 0xb6, 0x97,
	0x8f, 0xa5, 0xe0, 0x55, 0xcb, 0x33, 0x59, 0xd4, 0x95, 0xfe, 0x93, 0xfc, 0x20, 0xf0, 0x88, 0x19,
	0xf2, 0x9a, 0x0b, 0x4e, 0x9f, 0x43, 0x98, 0xe3, 0x86, 0x77, 0xa5, 0xbc, 0xce, 0x73, 0x81, 0x6d,
	0x1b, 0x91, 0x98, 0x2c, 0x1f, 0xb2, 0x3d, 0x94, 0x5e, 0x40, 0xd0, 0x4a, 0x2e, 0xbb, 0x36, 0xf2,
	0x62, 0xb2, 0xf4, 0x99, 0xf9, 0xa2, 0x57, 0x00, 0x99, 0x40, 0x2e, 0xf1, 0xb6, 0xd8, 0x62, 0x34,
	0x8b, 0xc9, 0x72, 0xc6, 0x1c, 0x84, 0xc6, 0xb0, 0x68, 0x04, 0x36, 0x5c, 0x68, 0xc2, 0x03, 0x45,
	0x70, 0xa1, 0x1d, 0x23, 0xc7, 0x92, 0xf7, 0x6b, 0x14, 0x45, 0x9d, 0x47, 0xbe, 0x66, 0x38, 0x50,
	0xf2, 0x09, 0x4e, 0x6d, 0xcf, 0xf4, 0x19, 0x9c, 0xdd, 0xf1, 0xec, 0x4b, 0xd7, 0x0c, 0xdb, 0x1d,
	0x82, 0xf4, 0x05, 0xcc, 0x05, 0xca, 0xdd, 0x05, 0x23, 0x2f, 0x9e, 0x2d, 0x17, 0xe9, 0x93, 0x95,
	0xb2, 0x64, 0xe5, 0xde, 0x9d, 0x59, 0x4e, 0xf2, 0x9b, 0x40, 0x68, 0xff, 0x5c, 0x2b, 0xbb, 0x68,
	0x0a, 0x73, 0xd3, 0xa4, 0xaa, 0xb0, 0x48, 0x2f, 0xcc, 0x09, 0x6b, 0x8d, 0x5a, 0xfa, 0xdb, 0x13,
	0x66, 0x89, 0x4a, 0x83, 0x62, 0x53, 0x8b, 0xad, 0x32, 0xc9, 0xd1, 0x68, 0x74, 0xa0, 0xd1, 0x10,
	0x7d, 0x09, 0x81, 0x6e, 0x5d, 0x79, 0xb7, 0x48, 0x9f, 0x1a, 0xc9, 0x8d, 0x02, 0x1d, 0x85, 0xa1,
	0xed, 0x04, 0x19, 0xaf, 0x32, 0x2c, 0x95, 0x97, 0x7f, 0x05, 0x6f, 0x14, 0xe8, 0x0a, 0x34, 0x8d,
	0x86, 0xe0, 0xc9, 0x5e, 0xd9, 0xea, 0x33, 0x4f, 0xf6, 0x37, 0x73, 0xf0, 0xef, 0x79, 0xd9, 0x61,
	0xf2, 0x9d, 0x40, 0x38, 0x2c, 0x73, 0xa4, 0xbb, 0xff, 0xce, 0x8c, 0x37, 0x3a, 0x33, 0x7b, 0x2f,
	0x3b, 0x1b, 0x7b, 0xd9, 0xf3, 0x3d, 0x3f, 0xa7, 0x6d, 0x41, 0x15, 0x18, 0x9a, 0x3f, 0x71, 0x81,
	0x8f, 0x10, 0x0e, 0x9d, 0x9f, 0xfe, 0x02, 0x0c, 0xbf, 0xda, 0xc3, 0xdf, 0x55, 0x9b, 0x7a, 0xe2,
	0x02, 0x3f, 0x09, 0x9c, 0xd9, 0xe3, 0xdf, 0x77, 0x28, 0xfa, 0xff, 0x3d, 0x04, 0x47, 0x44, 0xc4,
	0x15, 0x80, 0xc0, 0x2d, 0x2f, 0x2a, 0x45, 0xd0, 0x09, 0xe1, 0x20, 0x4e, 0x38, 0x05, 0x6e, 0x38,
	0xa5, 0xbf, 0x08, 0x9c, 0xda, 0x68, 0xa4, 0xaf, 0x61, 0x6e, 0x66, 0x8d, 0x1e, 0xd8, 0xe5, 0xcb,
	0x73, 0x83, 0x7f, 0xa8, 0xda, 0xe2, 0x73, 0x75, 0xfb, 0x2d, 0x39, 0x51, 0x2a, 0xb3, 0xaa, 0x07,
	0xb6, 0x79, 0x4c, 0x95, 0x42, 0xa0, 0x37, 0x8b, 0x8e, 0xef, 0xf3, 0x01, 0x8d, 0x9e, 0x24, 0x3a,
	0xbe, 0xd2, 0x23, 0x9a, 0xbb, 0x40, 0xa5, 0xfa, 0xab, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x01,
	0x17, 0x39, 0x99, 0x01, 0x06, 0x00, 0x00,
}