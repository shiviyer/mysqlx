// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mysqlx_connection.proto

/*
	Package mysqlx_connection is a generated protocol buffer package.

	It is generated from these files:
		mysqlx_connection.proto

	It has these top-level messages:
		Capability
		Capabilities
		CapabilitiesGet
		CapabilitiesSet
		Close
*/
package mysqlx_connection

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import Mysqlx_Datatypes "github.com/AlekSi/mysqlx/internal/proto/mysqlx_datatypes"
import _ "github.com/AlekSi/mysqlx/internal/proto/mysqlx"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// a Capability
//
// a tuple of a ``name`` and a :protobuf:msg:`Mysqlx.Datatypes::Any`
type Capability struct {
	Name             *string               `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value            *Mysqlx_Datatypes.Any `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *Capability) Reset()                    { *m = Capability{} }
func (m *Capability) String() string            { return proto.CompactTextString(m) }
func (*Capability) ProtoMessage()               {}
func (*Capability) Descriptor() ([]byte, []int) { return fileDescriptorMysqlxConnection, []int{0} }

func (m *Capability) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Capability) GetValue() *Mysqlx_Datatypes.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

// Capabilities
type Capabilities struct {
	Capabilities     []*Capability `protobuf:"bytes,1,rep,name=capabilities" json:"capabilities,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Capabilities) Reset()                    { *m = Capabilities{} }
func (m *Capabilities) String() string            { return proto.CompactTextString(m) }
func (*Capabilities) ProtoMessage()               {}
func (*Capabilities) Descriptor() ([]byte, []int) { return fileDescriptorMysqlxConnection, []int{1} }

func (m *Capabilities) GetCapabilities() []*Capability {
	if m != nil {
		return m.Capabilities
	}
	return nil
}

// get supported connection capabilities and their current state
//
//   :returns: :protobuf:msg:`Mysqlx.Connection::Capabilities` or :protobuf:msg:`Mysqlx::Error`
//
type CapabilitiesGet struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CapabilitiesGet) Reset()                    { *m = CapabilitiesGet{} }
func (m *CapabilitiesGet) String() string            { return proto.CompactTextString(m) }
func (*CapabilitiesGet) ProtoMessage()               {}
func (*CapabilitiesGet) Descriptor() ([]byte, []int) { return fileDescriptorMysqlxConnection, []int{2} }

// sets connection capabilities atomically
//
// only provided values are changed, other values are left unchanged.
// If any of the changes fails, all changes are discarded.
//
// :precond: active sessions == 0
// :returns: :protobuf:msg:`Mysqlx::Ok` or :protobuf:msg:`Mysqlx::Error`
type CapabilitiesSet struct {
	Capabilities     *Capabilities `protobuf:"bytes,1,req,name=capabilities" json:"capabilities,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *CapabilitiesSet) Reset()                    { *m = CapabilitiesSet{} }
func (m *CapabilitiesSet) String() string            { return proto.CompactTextString(m) }
func (*CapabilitiesSet) ProtoMessage()               {}
func (*CapabilitiesSet) Descriptor() ([]byte, []int) { return fileDescriptorMysqlxConnection, []int{3} }

func (m *CapabilitiesSet) GetCapabilities() *Capabilities {
	if m != nil {
		return m.Capabilities
	}
	return nil
}

// announce to the server that the client wants to close the connection
//
// it discards any session state of the server
//
// :Returns: :protobuf:msg:`Mysqlx::Ok`
type Close struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Close) Reset()                    { *m = Close{} }
func (m *Close) String() string            { return proto.CompactTextString(m) }
func (*Close) ProtoMessage()               {}
func (*Close) Descriptor() ([]byte, []int) { return fileDescriptorMysqlxConnection, []int{4} }

func init() {
	proto.RegisterType((*Capability)(nil), "Mysqlx.Connection.Capability")
	proto.RegisterType((*Capabilities)(nil), "Mysqlx.Connection.Capabilities")
	proto.RegisterType((*CapabilitiesGet)(nil), "Mysqlx.Connection.CapabilitiesGet")
	proto.RegisterType((*CapabilitiesSet)(nil), "Mysqlx.Connection.CapabilitiesSet")
	proto.RegisterType((*Close)(nil), "Mysqlx.Connection.Close")
}
func (m *Capability) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Capability) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Name == nil {
		return 0, new(proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMysqlxConnection(dAtA, i, uint64(len(*m.Name)))
		i += copy(dAtA[i:], *m.Name)
	}
	if m.Value == nil {
		return 0, new(proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMysqlxConnection(dAtA, i, uint64(m.Value.Size()))
		n1, err := m.Value.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Capabilities) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Capabilities) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Capabilities) > 0 {
		for _, msg := range m.Capabilities {
			dAtA[i] = 0xa
			i++
			i = encodeVarintMysqlxConnection(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *CapabilitiesGet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CapabilitiesGet) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *CapabilitiesSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CapabilitiesSet) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Capabilities == nil {
		return 0, new(proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMysqlxConnection(dAtA, i, uint64(m.Capabilities.Size()))
		n2, err := m.Capabilities.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Close) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Close) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintMysqlxConnection(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Capability) Size() (n int) {
	var l int
	_ = l
	if m.Name != nil {
		l = len(*m.Name)
		n += 1 + l + sovMysqlxConnection(uint64(l))
	}
	if m.Value != nil {
		l = m.Value.Size()
		n += 1 + l + sovMysqlxConnection(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Capabilities) Size() (n int) {
	var l int
	_ = l
	if len(m.Capabilities) > 0 {
		for _, e := range m.Capabilities {
			l = e.Size()
			n += 1 + l + sovMysqlxConnection(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CapabilitiesGet) Size() (n int) {
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CapabilitiesSet) Size() (n int) {
	var l int
	_ = l
	if m.Capabilities != nil {
		l = m.Capabilities.Size()
		n += 1 + l + sovMysqlxConnection(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Close) Size() (n int) {
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMysqlxConnection(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMysqlxConnection(x uint64) (n int) {
	return sovMysqlxConnection(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Capability) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMysqlxConnection
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Capability: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Capability: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlxConnection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMysqlxConnection
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Name = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlxConnection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMysqlxConnection
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Value == nil {
				m.Value = &Mysqlx_Datatypes.Any{}
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		default:
			iNdEx = preIndex
			skippy, err := skipMysqlxConnection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMysqlxConnection
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return new(proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Capabilities) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMysqlxConnection
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Capabilities: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Capabilities: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Capabilities", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlxConnection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMysqlxConnection
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Capabilities = append(m.Capabilities, &Capability{})
			if err := m.Capabilities[len(m.Capabilities)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMysqlxConnection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMysqlxConnection
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CapabilitiesGet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMysqlxConnection
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CapabilitiesGet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CapabilitiesGet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMysqlxConnection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMysqlxConnection
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CapabilitiesSet) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMysqlxConnection
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CapabilitiesSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CapabilitiesSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Capabilities", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlxConnection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMysqlxConnection
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Capabilities == nil {
				m.Capabilities = &Capabilities{}
			}
			if err := m.Capabilities.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		default:
			iNdEx = preIndex
			skippy, err := skipMysqlxConnection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMysqlxConnection
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Close) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMysqlxConnection
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Close: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Close: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMysqlxConnection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMysqlxConnection
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMysqlxConnection(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMysqlxConnection
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
					return 0, ErrIntOverflowMysqlxConnection
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMysqlxConnection
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMysqlxConnection
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMysqlxConnection
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMysqlxConnection(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMysqlxConnection = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMysqlxConnection   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("mysqlx_connection.proto", fileDescriptorMysqlxConnection) }

var fileDescriptorMysqlxConnection = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcf, 0xad, 0x2c, 0x2e,
	0xcc, 0xa9, 0x88, 0x4f, 0xce, 0xcf, 0xcb, 0x4b, 0x4d, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x12, 0xf4, 0x05, 0x4b, 0xe8, 0x39, 0xc3, 0x25, 0xa4, 0xc4, 0xa0, 0x6a,
	0x53, 0x12, 0x4b, 0x12, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x21, 0x4a, 0xa5, 0x78, 0x20, 0xe2, 0x10,
	0x9e, 0x92, 0x2f, 0x17, 0x97, 0x73, 0x62, 0x41, 0x62, 0x52, 0x66, 0x4e, 0x66, 0x49, 0xa5, 0x90,
	0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0x93, 0x06, 0x67, 0x10, 0x98, 0x2d,
	0xa4, 0xcd, 0xc5, 0x5a, 0x96, 0x98, 0x53, 0x9a, 0x2a, 0xc1, 0xa4, 0xc0, 0xa4, 0xc1, 0x6d, 0x24,
	0xaa, 0x07, 0xb5, 0xca, 0x05, 0x6e, 0xae, 0x63, 0x5e, 0x65, 0x10, 0x44, 0x8d, 0x52, 0x38, 0x17,
	0x0f, 0xdc, 0xb8, 0xcc, 0xd4, 0x62, 0x21, 0x47, 0x2e, 0x9e, 0x64, 0x24, 0xbe, 0x04, 0xa3, 0x02,
	0xb3, 0x06, 0xb7, 0x91, 0xac, 0x1e, 0x86, 0x73, 0xf5, 0x10, 0xae, 0x08, 0x42, 0xd1, 0x62, 0xc5,
	0x32, 0xe1, 0x95, 0x01, 0x93, 0x92, 0x38, 0x17, 0x3f, 0xb2, 0xc1, 0xee, 0xa9, 0x25, 0x56, 0x2c,
	0x1d, 0xaf, 0x0c, 0x18, 0x95, 0x62, 0x50, 0x25, 0x82, 0x53, 0x4b, 0x84, 0x9c, 0x31, 0x2c, 0x05,
	0x39, 0x5c, 0x1e, 0x9f, 0xa5, 0x99, 0xa9, 0xc5, 0xe8, 0xd6, 0x76, 0x80, 0xac, 0xe5, 0xe5, 0x62,
	0x75, 0xce, 0xc9, 0x2f, 0x4e, 0x05, 0x73, 0x99, 0x9d, 0x34, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0,
	0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x19, 0x8f, 0xe5, 0x18, 0xb8, 0xc4, 0x93, 0xf3, 0x73,
	0xf5, 0xc0, 0x21, 0xaa, 0x97, 0x9c, 0xa5, 0x07, 0x0d, 0xd5, 0xa4, 0xd2, 0x34, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd7, 0x1d, 0x15, 0x2a, 0xab, 0x01, 0x00, 0x00,
}
