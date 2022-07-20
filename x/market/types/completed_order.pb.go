// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: market/completed_order.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type CompletedOrder struct {
	Index       string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	OrderIndex  string `protobuf:"bytes,2,opt,name=orderIndex,proto3" json:"orderIndex,omitempty"`
	CompletedBy string `protobuf:"bytes,3,opt,name=completedBy,proto3" json:"completedBy,omitempty"`
	Height      string `protobuf:"bytes,4,opt,name=height,proto3" json:"height,omitempty"`
	Date        string `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	FuelUsed    string `protobuf:"bytes,6,opt,name=fuelUsed,proto3" json:"fuelUsed,omitempty"`
}

func (m *CompletedOrder) Reset()         { *m = CompletedOrder{} }
func (m *CompletedOrder) String() string { return proto.CompactTextString(m) }
func (*CompletedOrder) ProtoMessage()    {}
func (*CompletedOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9c5747e52889860, []int{0}
}
func (m *CompletedOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CompletedOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CompletedOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CompletedOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompletedOrder.Merge(m, src)
}
func (m *CompletedOrder) XXX_Size() int {
	return m.Size()
}
func (m *CompletedOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_CompletedOrder.DiscardUnknown(m)
}

var xxx_messageInfo_CompletedOrder proto.InternalMessageInfo

func (m *CompletedOrder) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *CompletedOrder) GetOrderIndex() string {
	if m != nil {
		return m.OrderIndex
	}
	return ""
}

func (m *CompletedOrder) GetCompletedBy() string {
	if m != nil {
		return m.CompletedBy
	}
	return ""
}

func (m *CompletedOrder) GetHeight() string {
	if m != nil {
		return m.Height
	}
	return ""
}

func (m *CompletedOrder) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *CompletedOrder) GetFuelUsed() string {
	if m != nil {
		return m.FuelUsed
	}
	return ""
}

func init() {
	proto.RegisterType((*CompletedOrder)(nil), "txlabs.blocklesschain.market.CompletedOrder")
}

func init() { proto.RegisterFile("market/completed_order.proto", fileDescriptor_c9c5747e52889860) }

var fileDescriptor_c9c5747e52889860 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc9, 0x4d, 0x2c, 0xca,
	0x4e, 0x2d, 0xd1, 0x4f, 0xce, 0xcf, 0x2d, 0xc8, 0x49, 0x2d, 0x49, 0x4d, 0x89, 0xcf, 0x2f, 0x4a,
	0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x29, 0xa9, 0xc8, 0x49, 0x4c, 0x2a,
	0xd6, 0x4b, 0xca, 0xc9, 0x4f, 0xce, 0xce, 0x49, 0x2d, 0x2e, 0x4e, 0xce, 0x48, 0xcc, 0xcc, 0xd3,
	0x83, 0xe8, 0x51, 0xda, 0xc0, 0xc8, 0xc5, 0xe7, 0x0c, 0xd3, 0xe7, 0x0f, 0xd2, 0x26, 0x24, 0xc2,
	0xc5, 0x9a, 0x99, 0x97, 0x92, 0x5a, 0x21, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe1, 0x08,
	0xc9, 0x71, 0x71, 0x81, 0x4d, 0xf5, 0x04, 0x4b, 0x31, 0x81, 0xa5, 0x90, 0x44, 0x84, 0x14, 0xb8,
	0xb8, 0xe1, 0xf6, 0x3b, 0x55, 0x4a, 0x30, 0x83, 0x15, 0x20, 0x0b, 0x09, 0x89, 0x71, 0xb1, 0x65,
	0xa4, 0x66, 0xa6, 0x67, 0x94, 0x48, 0xb0, 0x80, 0x25, 0xa1, 0x3c, 0x21, 0x21, 0x2e, 0x96, 0x94,
	0xc4, 0x92, 0x54, 0x09, 0x56, 0xb0, 0x28, 0x98, 0x2d, 0x24, 0xc5, 0xc5, 0x91, 0x56, 0x9a, 0x9a,
	0x13, 0x5a, 0x9c, 0x9a, 0x22, 0xc1, 0x06, 0x16, 0x87, 0xf3, 0x9d, 0xbc, 0x4e, 0x3c, 0x92, 0x63,
	0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96,
	0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x20, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39,
	0x3f, 0x57, 0x1f, 0xe2, 0x6b, 0x7d, 0xb8, 0xaf, 0x75, 0xc1, 0xde, 0xd6, 0xaf, 0xd0, 0x87, 0x06,
	0x56, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0x8c, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xd8, 0xfb, 0x5a, 0xca, 0x43, 0x01, 0x00, 0x00,
}

func (m *CompletedOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CompletedOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CompletedOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FuelUsed) > 0 {
		i -= len(m.FuelUsed)
		copy(dAtA[i:], m.FuelUsed)
		i = encodeVarintCompletedOrder(dAtA, i, uint64(len(m.FuelUsed)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Date) > 0 {
		i -= len(m.Date)
		copy(dAtA[i:], m.Date)
		i = encodeVarintCompletedOrder(dAtA, i, uint64(len(m.Date)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Height) > 0 {
		i -= len(m.Height)
		copy(dAtA[i:], m.Height)
		i = encodeVarintCompletedOrder(dAtA, i, uint64(len(m.Height)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CompletedBy) > 0 {
		i -= len(m.CompletedBy)
		copy(dAtA[i:], m.CompletedBy)
		i = encodeVarintCompletedOrder(dAtA, i, uint64(len(m.CompletedBy)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OrderIndex) > 0 {
		i -= len(m.OrderIndex)
		copy(dAtA[i:], m.OrderIndex)
		i = encodeVarintCompletedOrder(dAtA, i, uint64(len(m.OrderIndex)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintCompletedOrder(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCompletedOrder(dAtA []byte, offset int, v uint64) int {
	offset -= sovCompletedOrder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CompletedOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovCompletedOrder(uint64(l))
	}
	l = len(m.OrderIndex)
	if l > 0 {
		n += 1 + l + sovCompletedOrder(uint64(l))
	}
	l = len(m.CompletedBy)
	if l > 0 {
		n += 1 + l + sovCompletedOrder(uint64(l))
	}
	l = len(m.Height)
	if l > 0 {
		n += 1 + l + sovCompletedOrder(uint64(l))
	}
	l = len(m.Date)
	if l > 0 {
		n += 1 + l + sovCompletedOrder(uint64(l))
	}
	l = len(m.FuelUsed)
	if l > 0 {
		n += 1 + l + sovCompletedOrder(uint64(l))
	}
	return n
}

func sovCompletedOrder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCompletedOrder(x uint64) (n int) {
	return sovCompletedOrder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CompletedOrder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCompletedOrder
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
			return fmt.Errorf("proto: CompletedOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CompletedOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCompletedOrder
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
				return ErrInvalidLengthCompletedOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCompletedOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCompletedOrder
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
				return ErrInvalidLengthCompletedOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCompletedOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderIndex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompletedBy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCompletedOrder
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
				return ErrInvalidLengthCompletedOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCompletedOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CompletedBy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCompletedOrder
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
				return ErrInvalidLengthCompletedOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCompletedOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Height = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Date", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCompletedOrder
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
				return ErrInvalidLengthCompletedOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCompletedOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Date = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FuelUsed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCompletedOrder
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
				return ErrInvalidLengthCompletedOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCompletedOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FuelUsed = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCompletedOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCompletedOrder
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
func skipCompletedOrder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCompletedOrder
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
					return 0, ErrIntOverflowCompletedOrder
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
					return 0, ErrIntOverflowCompletedOrder
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
				return 0, ErrInvalidLengthCompletedOrder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCompletedOrder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCompletedOrder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCompletedOrder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCompletedOrder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCompletedOrder = fmt.Errorf("proto: unexpected end of group")
)