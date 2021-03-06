// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mysqlx_session.proto

/*
	Package mysqlx_session is a generated protocol buffer package.

	Messages to manage Sessions

	.. uml::

	  == session start ==
	  Client -> Server: AuthenticateStart
	  opt
	  Server --> Client: AuthenticateContinue
	  Client --> Server: AuthenticateContinue
	  end
	  alt
	  Server --> Client: AuthenticateOk
	  else
	  Server --> Client: Error
	  end
	  ...
	  == session reset ==
	  Client -> Server: Reset
	  Server --> Client: Ok
	  == session end ==
	  Client -> Server: Close
	  Server --> Client: Ok


	It is generated from these files:
		mysqlx_session.proto

	It has these top-level messages:
		AuthenticateStart
		AuthenticateContinue
		AuthenticateOk
		Reset
		Close
*/
package mysqlx_session

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

// the initial message send from the client to the server to start the
// authentication proccess
//
// :param mech_name: authentication mechanism name
// :param auth_data: authentication data
// :param initial_response: initial response
// :Returns: :protobuf:msg:`Mysqlx.Session::AuthenticateContinue`
type AuthenticateStart struct {
	MechName         *string `protobuf:"bytes,1,req,name=mech_name,json=mechName" json:"mech_name,omitempty"`
	AuthData         []byte  `protobuf:"bytes,2,opt,name=auth_data,json=authData" json:"auth_data,omitempty"`
	InitialResponse  []byte  `protobuf:"bytes,3,opt,name=initial_response,json=initialResponse" json:"initial_response,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AuthenticateStart) Reset()                    { *m = AuthenticateStart{} }
func (m *AuthenticateStart) String() string            { return proto.CompactTextString(m) }
func (*AuthenticateStart) ProtoMessage()               {}
func (*AuthenticateStart) Descriptor() ([]byte, []int) { return fileDescriptorMysqlxSession, []int{0} }

func (m *AuthenticateStart) GetMechName() string {
	if m != nil && m.MechName != nil {
		return *m.MechName
	}
	return ""
}

func (m *AuthenticateStart) GetAuthData() []byte {
	if m != nil {
		return m.AuthData
	}
	return nil
}

func (m *AuthenticateStart) GetInitialResponse() []byte {
	if m != nil {
		return m.InitialResponse
	}
	return nil
}

// send by client or server after a :protobuf:msg:`Mysqlx.Session::AuthenticateStart` to
// exchange more auth data
//
// :param auth_data: authentication data
// :Returns: :protobuf:msg:`Mysqlx.Session::AuthenticateContinue`
type AuthenticateContinue struct {
	AuthData         []byte `protobuf:"bytes,1,req,name=auth_data,json=authData" json:"auth_data,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *AuthenticateContinue) Reset()         { *m = AuthenticateContinue{} }
func (m *AuthenticateContinue) String() string { return proto.CompactTextString(m) }
func (*AuthenticateContinue) ProtoMessage()    {}
func (*AuthenticateContinue) Descriptor() ([]byte, []int) {
	return fileDescriptorMysqlxSession, []int{1}
}

func (m *AuthenticateContinue) GetAuthData() []byte {
	if m != nil {
		return m.AuthData
	}
	return nil
}

// sent by the server after successful authentication
//
// :param auth_data: authentication data
type AuthenticateOk struct {
	AuthData         []byte `protobuf:"bytes,1,opt,name=auth_data,json=authData" json:"auth_data,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *AuthenticateOk) Reset()                    { *m = AuthenticateOk{} }
func (m *AuthenticateOk) String() string            { return proto.CompactTextString(m) }
func (*AuthenticateOk) ProtoMessage()               {}
func (*AuthenticateOk) Descriptor() ([]byte, []int) { return fileDescriptorMysqlxSession, []int{2} }

func (m *AuthenticateOk) GetAuthData() []byte {
	if m != nil {
		return m.AuthData
	}
	return nil
}

// reset the current session
//
// :Returns: :protobuf:msg:`Mysqlx::Ok`
type Reset struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Reset) Reset()                    { *m = Reset{} }
func (m *Reset) String() string            { return proto.CompactTextString(m) }
func (*Reset) ProtoMessage()               {}
func (*Reset) Descriptor() ([]byte, []int) { return fileDescriptorMysqlxSession, []int{3} }

// close the current session
//
// :Returns: :protobuf:msg:`Mysqlx::Ok`
type Close struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Close) Reset()                    { *m = Close{} }
func (m *Close) String() string            { return proto.CompactTextString(m) }
func (*Close) ProtoMessage()               {}
func (*Close) Descriptor() ([]byte, []int) { return fileDescriptorMysqlxSession, []int{4} }

func init() {
	proto.RegisterType((*AuthenticateStart)(nil), "Mysqlx.Session.AuthenticateStart")
	proto.RegisterType((*AuthenticateContinue)(nil), "Mysqlx.Session.AuthenticateContinue")
	proto.RegisterType((*AuthenticateOk)(nil), "Mysqlx.Session.AuthenticateOk")
	proto.RegisterType((*Reset)(nil), "Mysqlx.Session.Reset")
	proto.RegisterType((*Close)(nil), "Mysqlx.Session.Close")
}
func (m *AuthenticateStart) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthenticateStart) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MechName == nil {
		return 0, new(proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMysqlxSession(dAtA, i, uint64(len(*m.MechName)))
		i += copy(dAtA[i:], *m.MechName)
	}
	if m.AuthData != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMysqlxSession(dAtA, i, uint64(len(m.AuthData)))
		i += copy(dAtA[i:], m.AuthData)
	}
	if m.InitialResponse != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMysqlxSession(dAtA, i, uint64(len(m.InitialResponse)))
		i += copy(dAtA[i:], m.InitialResponse)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *AuthenticateContinue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthenticateContinue) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.AuthData == nil {
		return 0, new(proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMysqlxSession(dAtA, i, uint64(len(m.AuthData)))
		i += copy(dAtA[i:], m.AuthData)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *AuthenticateOk) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthenticateOk) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.AuthData != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMysqlxSession(dAtA, i, uint64(len(m.AuthData)))
		i += copy(dAtA[i:], m.AuthData)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Reset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Reset) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
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

func encodeVarintMysqlxSession(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AuthenticateStart) Size() (n int) {
	var l int
	_ = l
	if m.MechName != nil {
		l = len(*m.MechName)
		n += 1 + l + sovMysqlxSession(uint64(l))
	}
	if m.AuthData != nil {
		l = len(m.AuthData)
		n += 1 + l + sovMysqlxSession(uint64(l))
	}
	if m.InitialResponse != nil {
		l = len(m.InitialResponse)
		n += 1 + l + sovMysqlxSession(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthenticateContinue) Size() (n int) {
	var l int
	_ = l
	if m.AuthData != nil {
		l = len(m.AuthData)
		n += 1 + l + sovMysqlxSession(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthenticateOk) Size() (n int) {
	var l int
	_ = l
	if m.AuthData != nil {
		l = len(m.AuthData)
		n += 1 + l + sovMysqlxSession(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Reset) Size() (n int) {
	var l int
	_ = l
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

func sovMysqlxSession(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMysqlxSession(x uint64) (n int) {
	return sovMysqlxSession(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AuthenticateStart) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMysqlxSession
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
			return fmt.Errorf("proto: AuthenticateStart: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthenticateStart: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MechName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlxSession
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
				return ErrInvalidLengthMysqlxSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.MechName = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlxSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMysqlxSession
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthData = append(m.AuthData[:0], dAtA[iNdEx:postIndex]...)
			if m.AuthData == nil {
				m.AuthData = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialResponse", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlxSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMysqlxSession
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialResponse = append(m.InitialResponse[:0], dAtA[iNdEx:postIndex]...)
			if m.InitialResponse == nil {
				m.InitialResponse = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMysqlxSession(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMysqlxSession
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
func (m *AuthenticateContinue) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMysqlxSession
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
			return fmt.Errorf("proto: AuthenticateContinue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthenticateContinue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlxSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMysqlxSession
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthData = append(m.AuthData[:0], dAtA[iNdEx:postIndex]...)
			if m.AuthData == nil {
				m.AuthData = []byte{}
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		default:
			iNdEx = preIndex
			skippy, err := skipMysqlxSession(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMysqlxSession
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
func (m *AuthenticateOk) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMysqlxSession
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
			return fmt.Errorf("proto: AuthenticateOk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthenticateOk: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlxSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMysqlxSession
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthData = append(m.AuthData[:0], dAtA[iNdEx:postIndex]...)
			if m.AuthData == nil {
				m.AuthData = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMysqlxSession(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMysqlxSession
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
func (m *Reset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMysqlxSession
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
			return fmt.Errorf("proto: Reset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Reset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMysqlxSession(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMysqlxSession
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
func (m *Close) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMysqlxSession
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
			skippy, err := skipMysqlxSession(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMysqlxSession
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
func skipMysqlxSession(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMysqlxSession
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
					return 0, ErrIntOverflowMysqlxSession
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
					return 0, ErrIntOverflowMysqlxSession
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
				return 0, ErrInvalidLengthMysqlxSession
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMysqlxSession
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
				next, err := skipMysqlxSession(dAtA[start:])
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
	ErrInvalidLengthMysqlxSession = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMysqlxSession   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("mysqlx_session.proto", fileDescriptorMysqlxSession) }

var fileDescriptorMysqlxSession = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0xad, 0x2c, 0x2e,
	0xcc, 0xa9, 0x88, 0x2f, 0x4e, 0x2d, 0x2e, 0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0xf3, 0x05, 0x8b, 0xea, 0x05, 0x43, 0x44, 0xa5, 0x78, 0x20, 0xaa, 0x20, 0xb2, 0x4a,
	0x75, 0x5c, 0x82, 0x8e, 0xa5, 0x25, 0x19, 0xa9, 0x79, 0x25, 0x99, 0xc9, 0x89, 0x25, 0xa9, 0xc1,
	0x25, 0x89, 0x45, 0x25, 0x42, 0xd2, 0x5c, 0x9c, 0xb9, 0xa9, 0xc9, 0x19, 0xf1, 0x79, 0x89, 0xb9,
	0xa9, 0x12, 0x8c, 0x0a, 0x4c, 0x1a, 0x9c, 0x41, 0x1c, 0x20, 0x01, 0xbf, 0xc4, 0xdc, 0x54, 0x90,
	0x64, 0x62, 0x69, 0x49, 0x46, 0x7c, 0x4a, 0x62, 0x49, 0xa2, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x4f,
	0x10, 0x07, 0x48, 0xc0, 0x25, 0xb1, 0x24, 0x51, 0x48, 0x93, 0x4b, 0x20, 0x33, 0x2f, 0xb3, 0x24,
	0x33, 0x31, 0x27, 0xbe, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x82, 0x19, 0xac, 0x86,
	0x1f, 0x2a, 0x1e, 0x04, 0x15, 0xb6, 0x62, 0xe9, 0x78, 0x65, 0xc0, 0xa2, 0x64, 0xcb, 0x25, 0x82,
	0x6c, 0xbf, 0x73, 0x7e, 0x5e, 0x49, 0x66, 0x5e, 0x29, 0x9a, 0x2d, 0x20, 0x27, 0x20, 0xd9, 0x62,
	0xc5, 0xd1, 0xf1, 0xca, 0x80, 0x75, 0xc2, 0x2b, 0x03, 0x66, 0x25, 0x63, 0x2e, 0x3e, 0x64, 0xed,
	0xfe, 0xd9, 0xe8, 0x1a, 0x51, 0x9c, 0x67, 0xc5, 0x32, 0x01, 0x64, 0x27, 0x2f, 0x17, 0x6b, 0x50,
	0x6a, 0x71, 0x6a, 0x09, 0xd8, 0x09, 0x6c, 0x20, 0xae, 0x73, 0x4e, 0x3e, 0xd4, 0x45, 0xec, 0x4e,
	0x9a, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8c, 0xc7,
	0x72, 0x0c, 0x5c, 0xe2, 0xc9, 0xf9, 0xb9, 0x7a, 0xe0, 0x50, 0xd3, 0x4b, 0xce, 0xd2, 0x83, 0x86,
	0x5c, 0x52, 0x69, 0x1a, 0x20, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x4e, 0x7b, 0xfb, 0x71, 0x01, 0x00,
	0x00,
}
