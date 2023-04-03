// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: examplechain/examplemodule/entity_name.proto

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

type EntityName struct {
	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Field1  string `protobuf:"bytes,2,opt,name=field1,proto3" json:"field1,omitempty"`
	Field2  int32  `protobuf:"varint,3,opt,name=field2,proto3" json:"field2,omitempty"`
	Creator string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *EntityName) Reset()         { *m = EntityName{} }
func (m *EntityName) String() string { return proto.CompactTextString(m) }
func (*EntityName) ProtoMessage()    {}
func (*EntityName) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad3d005a1959f7e9, []int{0}
}
func (m *EntityName) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EntityName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EntityName.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EntityName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityName.Merge(m, src)
}
func (m *EntityName) XXX_Size() int {
	return m.Size()
}
func (m *EntityName) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityName.DiscardUnknown(m)
}

var xxx_messageInfo_EntityName proto.InternalMessageInfo

func (m *EntityName) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EntityName) GetField1() string {
	if m != nil {
		return m.Field1
	}
	return ""
}

func (m *EntityName) GetField2() int32 {
	if m != nil {
		return m.Field2
	}
	return 0
}

func (m *EntityName) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*EntityName)(nil), "examplechain.examplemodule.EntityName")
}

func init() {
	proto.RegisterFile("examplechain/examplemodule/entity_name.proto", fileDescriptor_ad3d005a1959f7e9)
}

var fileDescriptor_ad3d005a1959f7e9 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x49, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x4d, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x87, 0x72, 0x72, 0xf3, 0x53, 0x4a, 0x73,
	0x52, 0xf5, 0x53, 0xf3, 0x4a, 0x32, 0x4b, 0x2a, 0xe3, 0xf3, 0x12, 0x73, 0x53, 0xf5, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0xa4, 0x90, 0x55, 0xeb, 0xa1, 0xa8, 0x56, 0x4a, 0xe3, 0xe2, 0x72, 0x05,
	0x6b, 0xf0, 0x4b, 0xcc, 0x4d, 0x15, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0x60, 0x09, 0x62, 0xca, 0x4c, 0x11, 0x12, 0xe3, 0x62, 0x4b, 0xcb, 0x4c, 0xcd, 0x49, 0x31, 0x94,
	0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2, 0xe0, 0xe2, 0x46, 0x12, 0xcc, 0x0a, 0x8c, 0x1a,
	0xac, 0x50, 0x71, 0x23, 0x21, 0x09, 0x2e, 0xf6, 0xe4, 0xa2, 0xd4, 0xc4, 0x92, 0xfc, 0x22, 0x09,
	0x16, 0xb0, 0x06, 0x18, 0xd7, 0xc9, 0xf6, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f,
	0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18,
	0xa2, 0x94, 0xa1, 0x0e, 0xd2, 0x85, 0x78, 0xa6, 0x02, 0xcd, 0x3b, 0x25, 0x95, 0x05, 0xa9, 0xc5,
	0x49, 0x6c, 0x60, 0x9f, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x78, 0x21, 0xff, 0xa9, 0xf9,
	0x00, 0x00, 0x00,
}

func (m *EntityName) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntityName) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EntityName) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintEntityName(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x22
	}
	if m.Field2 != 0 {
		i = encodeVarintEntityName(dAtA, i, uint64(m.Field2))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Field1) > 0 {
		i -= len(m.Field1)
		copy(dAtA[i:], m.Field1)
		i = encodeVarintEntityName(dAtA, i, uint64(len(m.Field1)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintEntityName(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEntityName(dAtA []byte, offset int, v uint64) int {
	offset -= sovEntityName(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EntityName) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovEntityName(uint64(m.Id))
	}
	l = len(m.Field1)
	if l > 0 {
		n += 1 + l + sovEntityName(uint64(l))
	}
	if m.Field2 != 0 {
		n += 1 + sovEntityName(uint64(m.Field2))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovEntityName(uint64(l))
	}
	return n
}

func sovEntityName(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEntityName(x uint64) (n int) {
	return sovEntityName(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EntityName) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntityName
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
			return fmt.Errorf("proto: EntityName: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EntityName: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntityName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntityName
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
				return ErrInvalidLengthEntityName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEntityName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Field1 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field2", wireType)
			}
			m.Field2 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntityName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Field2 |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntityName
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
				return ErrInvalidLengthEntityName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEntityName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEntityName(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEntityName
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
func skipEntityName(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEntityName
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
					return 0, ErrIntOverflowEntityName
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
					return 0, ErrIntOverflowEntityName
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
				return 0, ErrInvalidLengthEntityName
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEntityName
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEntityName
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEntityName        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEntityName          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEntityName = fmt.Errorf("proto: unexpected end of group")
)
