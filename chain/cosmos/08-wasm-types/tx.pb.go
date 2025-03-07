// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/lightclients/wasm/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgStoreCode defines the request type for the StoreCode rpc.
type MsgStoreCode struct {
	Signer string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	// wasm byte code of light client contract. It can be raw or gzip compressed
	WasmByteCode []byte `protobuf:"bytes,2,opt,name=wasm_byte_code,json=wasmByteCode,proto3" json:"wasm_byte_code,omitempty"`
}

func (m *MsgStoreCode) Reset()         { *m = MsgStoreCode{} }
func (m *MsgStoreCode) String() string { return proto.CompactTextString(m) }
func (*MsgStoreCode) ProtoMessage()    {}
func (*MsgStoreCode) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d9737363bf1e38d, []int{0}
}
func (m *MsgStoreCode) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStoreCode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStoreCode.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStoreCode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStoreCode.Merge(m, src)
}
func (m *MsgStoreCode) XXX_Size() int {
	return m.Size()
}
func (m *MsgStoreCode) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStoreCode.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStoreCode proto.InternalMessageInfo

func (m *MsgStoreCode) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgStoreCode) GetWasmByteCode() []byte {
	if m != nil {
		return m.WasmByteCode
	}
	return nil
}

// MsgStoreCodeResponse defines the response type for the StoreCode rpc
type MsgStoreCodeResponse struct {
	// the sha256 hash of the stored code
	Checksum []byte `protobuf:"bytes,1,opt,name=checksum,proto3" json:"checksum,omitempty"`
}

func (m *MsgStoreCodeResponse) Reset()         { *m = MsgStoreCodeResponse{} }
func (m *MsgStoreCodeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgStoreCodeResponse) ProtoMessage()    {}
func (*MsgStoreCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d9737363bf1e38d, []int{1}
}
func (m *MsgStoreCodeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStoreCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStoreCodeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStoreCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStoreCodeResponse.Merge(m, src)
}
func (m *MsgStoreCodeResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgStoreCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStoreCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStoreCodeResponse proto.InternalMessageInfo

func (m *MsgStoreCodeResponse) GetChecksum() []byte {
	if m != nil {
		return m.Checksum
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgStoreCode)(nil), "ibc.lightclients.wasm.v1.MsgStoreCode")
	proto.RegisterType((*MsgStoreCodeResponse)(nil), "ibc.lightclients.wasm.v1.MsgStoreCodeResponse")
}

func init() { proto.RegisterFile("ibc/lightclients/wasm/v1/tx.proto", fileDescriptor_1d9737363bf1e38d) }

var fileDescriptor_1d9737363bf1e38d = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcc, 0x4c, 0x4a, 0xd6,
	0xcf, 0xc9, 0x4c, 0xcf, 0x28, 0x49, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x29, 0xd6, 0x2f, 0x4f, 0x2c,
	0xce, 0xd5, 0x2f, 0x33, 0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0xc8,
	0x4c, 0x4a, 0xd6, 0x43, 0x56, 0xa2, 0x07, 0x52, 0xa2, 0x57, 0x66, 0x28, 0x25, 0x9e, 0x9c, 0x5f,
	0x9c, 0x9b, 0x5f, 0xac, 0x9f, 0x5b, 0x9c, 0x0e, 0xd2, 0x91, 0x5b, 0x9c, 0x0e, 0xd1, 0xa2, 0x14,
	0xc9, 0xc5, 0xe3, 0x5b, 0x9c, 0x1e, 0x5c, 0x92, 0x5f, 0x94, 0xea, 0x9c, 0x9f, 0x92, 0x2a, 0x24,
	0xc6, 0xc5, 0x56, 0x9c, 0x99, 0x9e, 0x97, 0x5a, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04,
	0xe5, 0x09, 0xa9, 0x70, 0xf1, 0x81, 0xcc, 0x8a, 0x4f, 0xaa, 0x2c, 0x49, 0x8d, 0x4f, 0xce, 0x4f,
	0x49, 0x95, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x09, 0xe2, 0x01, 0x89, 0x3a, 0x55, 0x96, 0x80, 0x75,
	0x5b, 0x71, 0x37, 0x3d, 0xdf, 0xa0, 0x05, 0xd5, 0xa2, 0x64, 0xc4, 0x25, 0x82, 0x6c, 0x74, 0x50,
	0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x14, 0x17, 0x47, 0x72, 0x46, 0x6a, 0x72, 0x76,
	0x71, 0x69, 0x2e, 0xd8, 0x12, 0x9e, 0x20, 0x38, 0xdf, 0xa8, 0x90, 0x8b, 0xd9, 0xb7, 0x38, 0x5d,
	0x28, 0x99, 0x8b, 0x13, 0xe1, 0x24, 0x35, 0x3d, 0x5c, 0xde, 0xd2, 0x43, 0x36, 0x5f, 0x4a, 0x8f,
	0x38, 0x75, 0x30, 0x77, 0x48, 0xb1, 0x36, 0x3c, 0xdf, 0xa0, 0xc5, 0xe8, 0x14, 0x76, 0xe2, 0x91,
	0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1,
	0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x36, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a,
	0xc9, 0xf9, 0xb9, 0xfa, 0xd0, 0xf0, 0xcb, 0x4c, 0x4a, 0xd6, 0x4d, 0xcf, 0xd7, 0xcf, 0xcd, 0x4f,
	0x29, 0xcd, 0x49, 0x2d, 0x86, 0x44, 0x87, 0x2e, 0x2c, 0x3e, 0x0c, 0x2c, 0x74, 0xc1, 0x51, 0x52,
	0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x0e, 0x60, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xea, 0x8d, 0x6f, 0x6d, 0xb8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// StoreCode defines a rpc handler method for MsgStoreCode.
	StoreCode(ctx context.Context, in *MsgStoreCode, opts ...grpc.CallOption) (*MsgStoreCodeResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) StoreCode(ctx context.Context, in *MsgStoreCode, opts ...grpc.CallOption) (*MsgStoreCodeResponse, error) {
	out := new(MsgStoreCodeResponse)
	err := c.cc.Invoke(ctx, "/ibc.lightclients.wasm.v1.Msg/StoreCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// StoreCode defines a rpc handler method for MsgStoreCode.
	StoreCode(context.Context, *MsgStoreCode) (*MsgStoreCodeResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) StoreCode(ctx context.Context, req *MsgStoreCode) (*MsgStoreCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreCode not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_StoreCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStoreCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StoreCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.lightclients.wasm.v1.Msg/StoreCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StoreCode(ctx, req.(*MsgStoreCode))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.lightclients.wasm.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreCode",
			Handler:    _Msg_StoreCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/lightclients/wasm/v1/tx.proto",
}

func (m *MsgStoreCode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStoreCode) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStoreCode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WasmByteCode) > 0 {
		i -= len(m.WasmByteCode)
		copy(dAtA[i:], m.WasmByteCode)
		i = encodeVarintTx(dAtA, i, uint64(len(m.WasmByteCode)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgStoreCodeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStoreCodeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStoreCodeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Checksum) > 0 {
		i -= len(m.Checksum)
		copy(dAtA[i:], m.Checksum)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Checksum)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgStoreCode) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.WasmByteCode)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgStoreCodeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Checksum)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgStoreCode) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgStoreCode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStoreCode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WasmByteCode", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WasmByteCode = append(m.WasmByteCode[:0], dAtA[iNdEx:postIndex]...)
			if m.WasmByteCode == nil {
				m.WasmByteCode = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgStoreCodeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgStoreCodeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStoreCodeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checksum", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checksum = append(m.Checksum[:0], dAtA[iNdEx:postIndex]...)
			if m.Checksum == nil {
				m.Checksum = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
