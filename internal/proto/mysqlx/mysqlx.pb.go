// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mysqlx.proto

/*
	Package mysqlx is a generated protocol buffer package.

	It is generated from these files:
		mysqlx.proto

	It has these top-level messages:
		ClientMessages
		ServerMessages
		Ok
		Error
*/
package mysqlx

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

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

type ClientMessages_Type int32

const (
	ClientMessages_CON_CAPABILITIES_GET       ClientMessages_Type = 1
	ClientMessages_CON_CAPABILITIES_SET       ClientMessages_Type = 2
	ClientMessages_CON_CLOSE                  ClientMessages_Type = 3
	ClientMessages_SESS_AUTHENTICATE_START    ClientMessages_Type = 4
	ClientMessages_SESS_AUTHENTICATE_CONTINUE ClientMessages_Type = 5
	ClientMessages_SESS_RESET                 ClientMessages_Type = 6
	ClientMessages_SESS_CLOSE                 ClientMessages_Type = 7
	ClientMessages_SQL_STMT_EXECUTE           ClientMessages_Type = 12
	ClientMessages_CRUD_FIND                  ClientMessages_Type = 17
	ClientMessages_CRUD_INSERT                ClientMessages_Type = 18
	ClientMessages_CRUD_UPDATE                ClientMessages_Type = 19
	ClientMessages_CRUD_DELETE                ClientMessages_Type = 20
	ClientMessages_EXPECT_OPEN                ClientMessages_Type = 24
	ClientMessages_EXPECT_CLOSE               ClientMessages_Type = 25
	ClientMessages_CRUD_CREATE_VIEW           ClientMessages_Type = 30
	ClientMessages_CRUD_MODIFY_VIEW           ClientMessages_Type = 31
	ClientMessages_CRUD_DROP_VIEW             ClientMessages_Type = 32
)

var ClientMessages_Type_name = map[int32]string{
	1:  "CON_CAPABILITIES_GET",
	2:  "CON_CAPABILITIES_SET",
	3:  "CON_CLOSE",
	4:  "SESS_AUTHENTICATE_START",
	5:  "SESS_AUTHENTICATE_CONTINUE",
	6:  "SESS_RESET",
	7:  "SESS_CLOSE",
	12: "SQL_STMT_EXECUTE",
	17: "CRUD_FIND",
	18: "CRUD_INSERT",
	19: "CRUD_UPDATE",
	20: "CRUD_DELETE",
	24: "EXPECT_OPEN",
	25: "EXPECT_CLOSE",
	30: "CRUD_CREATE_VIEW",
	31: "CRUD_MODIFY_VIEW",
	32: "CRUD_DROP_VIEW",
}
var ClientMessages_Type_value = map[string]int32{
	"CON_CAPABILITIES_GET":       1,
	"CON_CAPABILITIES_SET":       2,
	"CON_CLOSE":                  3,
	"SESS_AUTHENTICATE_START":    4,
	"SESS_AUTHENTICATE_CONTINUE": 5,
	"SESS_RESET":                 6,
	"SESS_CLOSE":                 7,
	"SQL_STMT_EXECUTE":           12,
	"CRUD_FIND":                  17,
	"CRUD_INSERT":                18,
	"CRUD_UPDATE":                19,
	"CRUD_DELETE":                20,
	"EXPECT_OPEN":                24,
	"EXPECT_CLOSE":               25,
	"CRUD_CREATE_VIEW":           30,
	"CRUD_MODIFY_VIEW":           31,
	"CRUD_DROP_VIEW":             32,
}

func (x ClientMessages_Type) Enum() *ClientMessages_Type {
	p := new(ClientMessages_Type)
	*p = x
	return p
}
func (x ClientMessages_Type) String() string {
	return proto.EnumName(ClientMessages_Type_name, int32(x))
}
func (x *ClientMessages_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ClientMessages_Type_value, data, "ClientMessages_Type")
	if err != nil {
		return err
	}
	*x = ClientMessages_Type(value)
	return nil
}
func (ClientMessages_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptorMysqlx, []int{0, 0} }

type ServerMessages_Type int32

const (
	ServerMessages_OK                         ServerMessages_Type = 0
	ServerMessages_ERROR                      ServerMessages_Type = 1
	ServerMessages_CONN_CAPABILITIES          ServerMessages_Type = 2
	ServerMessages_SESS_AUTHENTICATE_CONTINUE ServerMessages_Type = 3
	ServerMessages_SESS_AUTHENTICATE_OK       ServerMessages_Type = 4
	// NOTICE has to stay at 11 forever
	ServerMessages_NOTICE                               ServerMessages_Type = 11
	ServerMessages_RESULTSET_COLUMN_META_DATA           ServerMessages_Type = 12
	ServerMessages_RESULTSET_ROW                        ServerMessages_Type = 13
	ServerMessages_RESULTSET_FETCH_DONE                 ServerMessages_Type = 14
	ServerMessages_RESULTSET_FETCH_SUSPENDED            ServerMessages_Type = 15
	ServerMessages_RESULTSET_FETCH_DONE_MORE_RESULTSETS ServerMessages_Type = 16
	ServerMessages_SQL_STMT_EXECUTE_OK                  ServerMessages_Type = 17
	ServerMessages_RESULTSET_FETCH_DONE_MORE_OUT_PARAMS ServerMessages_Type = 18
)

var ServerMessages_Type_name = map[int32]string{
	0:  "OK",
	1:  "ERROR",
	2:  "CONN_CAPABILITIES",
	3:  "SESS_AUTHENTICATE_CONTINUE",
	4:  "SESS_AUTHENTICATE_OK",
	11: "NOTICE",
	12: "RESULTSET_COLUMN_META_DATA",
	13: "RESULTSET_ROW",
	14: "RESULTSET_FETCH_DONE",
	15: "RESULTSET_FETCH_SUSPENDED",
	16: "RESULTSET_FETCH_DONE_MORE_RESULTSETS",
	17: "SQL_STMT_EXECUTE_OK",
	18: "RESULTSET_FETCH_DONE_MORE_OUT_PARAMS",
}
var ServerMessages_Type_value = map[string]int32{
	"OK":                                   0,
	"ERROR":                                1,
	"CONN_CAPABILITIES":                    2,
	"SESS_AUTHENTICATE_CONTINUE":           3,
	"SESS_AUTHENTICATE_OK":                 4,
	"NOTICE":                               11,
	"RESULTSET_COLUMN_META_DATA":           12,
	"RESULTSET_ROW":                        13,
	"RESULTSET_FETCH_DONE":                 14,
	"RESULTSET_FETCH_SUSPENDED":            15,
	"RESULTSET_FETCH_DONE_MORE_RESULTSETS": 16,
	"SQL_STMT_EXECUTE_OK":                  17,
	"RESULTSET_FETCH_DONE_MORE_OUT_PARAMS": 18,
}

func (x ServerMessages_Type) Enum() *ServerMessages_Type {
	p := new(ServerMessages_Type)
	*p = x
	return p
}
func (x ServerMessages_Type) String() string {
	return proto.EnumName(ServerMessages_Type_name, int32(x))
}
func (x *ServerMessages_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ServerMessages_Type_value, data, "ServerMessages_Type")
	if err != nil {
		return err
	}
	*x = ServerMessages_Type(value)
	return nil
}
func (ServerMessages_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptorMysqlx, []int{1, 0} }

type Error_Severity int32

const (
	Error_ERROR Error_Severity = 0
	Error_FATAL Error_Severity = 1
)

var Error_Severity_name = map[int32]string{
	0: "ERROR",
	1: "FATAL",
}
var Error_Severity_value = map[string]int32{
	"ERROR": 0,
	"FATAL": 1,
}

func (x Error_Severity) Enum() *Error_Severity {
	p := new(Error_Severity)
	*p = x
	return p
}
func (x Error_Severity) String() string {
	return proto.EnumName(Error_Severity_name, int32(x))
}
func (x *Error_Severity) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Error_Severity_value, data, "Error_Severity")
	if err != nil {
		return err
	}
	*x = Error_Severity(value)
	return nil
}
func (Error_Severity) EnumDescriptor() ([]byte, []int) { return fileDescriptorMysqlx, []int{3, 0} }

// IDs of messages that can be sent from client to the server
//
// .. note::
//   this message is never sent on the wire. It is only used to let ``protoc``
//
//   * generate constants
//   * check for uniqueness
type ClientMessages struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ClientMessages) Reset()                    { *m = ClientMessages{} }
func (m *ClientMessages) String() string            { return proto.CompactTextString(m) }
func (*ClientMessages) ProtoMessage()               {}
func (*ClientMessages) Descriptor() ([]byte, []int) { return fileDescriptorMysqlx, []int{0} }

// IDs of messages that can be sent from server to client
//
// .. note::
//   this message is never sent on the wire. It is only used to let ``protoc``
//
//   * generate constants
//   * check for uniqueness
type ServerMessages struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ServerMessages) Reset()                    { *m = ServerMessages{} }
func (m *ServerMessages) String() string            { return proto.CompactTextString(m) }
func (*ServerMessages) ProtoMessage()               {}
func (*ServerMessages) Descriptor() ([]byte, []int) { return fileDescriptorMysqlx, []int{1} }

// generic Ok message
type Ok struct {
	Msg              *string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Ok) Reset()                    { *m = Ok{} }
func (m *Ok) String() string            { return proto.CompactTextString(m) }
func (*Ok) ProtoMessage()               {}
func (*Ok) Descriptor() ([]byte, []int) { return fileDescriptorMysqlx, []int{2} }

func (m *Ok) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

// generic Error message
//
// A ``severity`` of ``ERROR`` indicates the current message sequence is
// aborted for the given error and the session is ready for more.
//
// In case of a ``FATAL`` error message the client should not expect
// the server to continue handling any further messages and should
// close the connection.
//
// :param severity: severity of the error message
// :param code: error-code
// :param sql_state: SQL state
// :param msg: human readable error message
type Error struct {
	Severity         *Error_Severity `protobuf:"varint,1,opt,name=severity,enum=Mysqlx.Error_Severity,def=0" json:"severity,omitempty"`
	Code             *uint32         `protobuf:"varint,2,req,name=code" json:"code,omitempty"`
	SqlState         *string         `protobuf:"bytes,4,req,name=sql_state,json=sqlState" json:"sql_state,omitempty"`
	Msg              *string         `protobuf:"bytes,3,req,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptorMysqlx, []int{3} }

const Default_Error_Severity Error_Severity = Error_ERROR

func (m *Error) GetSeverity() Error_Severity {
	if m != nil && m.Severity != nil {
		return *m.Severity
	}
	return Default_Error_Severity
}

func (m *Error) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *Error) GetSqlState() string {
	if m != nil && m.SqlState != nil {
		return *m.SqlState
	}
	return ""
}

func (m *Error) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

var E_ClientMessageId = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*ClientMessages_Type)(nil),
	Field:         100001,
	Name:          "Mysqlx.client_message_id",
	Tag:           "varint,100001,opt,name=client_message_id,json=clientMessageId,enum=Mysqlx.ClientMessages_Type",
	Filename:      "mysqlx.proto",
}

var E_ServerMessageId = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*ServerMessages_Type)(nil),
	Field:         100002,
	Name:          "Mysqlx.server_message_id",
	Tag:           "varint,100002,opt,name=server_message_id,json=serverMessageId,enum=Mysqlx.ServerMessages_Type",
	Filename:      "mysqlx.proto",
}

func init() {
	proto.RegisterType((*ClientMessages)(nil), "Mysqlx.ClientMessages")
	proto.RegisterType((*ServerMessages)(nil), "Mysqlx.ServerMessages")
	proto.RegisterType((*Ok)(nil), "Mysqlx.Ok")
	proto.RegisterType((*Error)(nil), "Mysqlx.Error")
	proto.RegisterEnum("Mysqlx.ClientMessages_Type", ClientMessages_Type_name, ClientMessages_Type_value)
	proto.RegisterEnum("Mysqlx.ServerMessages_Type", ServerMessages_Type_name, ServerMessages_Type_value)
	proto.RegisterEnum("Mysqlx.Error_Severity", Error_Severity_name, Error_Severity_value)
	proto.RegisterExtension(E_ClientMessageId)
	proto.RegisterExtension(E_ServerMessageId)
}
func (m *ClientMessages) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientMessages) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ServerMessages) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServerMessages) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Ok) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ok) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Msg != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMysqlx(dAtA, i, uint64(len(*m.Msg)))
		i += copy(dAtA[i:], *m.Msg)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Error) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Error) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Severity != nil {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMysqlx(dAtA, i, uint64(*m.Severity))
	}
	if m.Code == nil {
		return 0, new(proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMysqlx(dAtA, i, uint64(*m.Code))
	}
	if m.Msg == nil {
		return 0, new(proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMysqlx(dAtA, i, uint64(len(*m.Msg)))
		i += copy(dAtA[i:], *m.Msg)
	}
	if m.SqlState == nil {
		return 0, new(proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0x22
		i++
		i = encodeVarintMysqlx(dAtA, i, uint64(len(*m.SqlState)))
		i += copy(dAtA[i:], *m.SqlState)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintMysqlx(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ClientMessages) Size() (n int) {
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ServerMessages) Size() (n int) {
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Ok) Size() (n int) {
	var l int
	_ = l
	if m.Msg != nil {
		l = len(*m.Msg)
		n += 1 + l + sovMysqlx(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Error) Size() (n int) {
	var l int
	_ = l
	if m.Severity != nil {
		n += 1 + sovMysqlx(uint64(*m.Severity))
	}
	if m.Code != nil {
		n += 1 + sovMysqlx(uint64(*m.Code))
	}
	if m.Msg != nil {
		l = len(*m.Msg)
		n += 1 + l + sovMysqlx(uint64(l))
	}
	if m.SqlState != nil {
		l = len(*m.SqlState)
		n += 1 + l + sovMysqlx(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMysqlx(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMysqlx(x uint64) (n int) {
	return sovMysqlx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientMessages) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMysqlx
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
			return fmt.Errorf("proto: ClientMessages: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientMessages: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMysqlx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMysqlx
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
func (m *ServerMessages) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMysqlx
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
			return fmt.Errorf("proto: ServerMessages: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServerMessages: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMysqlx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMysqlx
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
func (m *Ok) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMysqlx
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
			return fmt.Errorf("proto: Ok: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ok: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlx
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
				return ErrInvalidLengthMysqlx
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Msg = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMysqlx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMysqlx
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
func (m *Error) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMysqlx
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
			return fmt.Errorf("proto: Error: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Error: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Severity", wireType)
			}
			var v Error_Severity
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (Error_Severity(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Severity = &v
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Code = &v
			hasFields[0] |= uint64(0x00000001)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlx
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
				return ErrInvalidLengthMysqlx
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Msg = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SqlState", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlx
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
				return ErrInvalidLengthMysqlx
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.SqlState = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000004)
		default:
			iNdEx = preIndex
			skippy, err := skipMysqlx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMysqlx
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
	if hasFields[0]&uint64(0x00000004) == 0 {
		return new(proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMysqlx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMysqlx
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
					return 0, ErrIntOverflowMysqlx
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
					return 0, ErrIntOverflowMysqlx
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
				return 0, ErrInvalidLengthMysqlx
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMysqlx
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
				next, err := skipMysqlx(dAtA[start:])
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
	ErrInvalidLengthMysqlx = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMysqlx   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("mysqlx.proto", fileDescriptorMysqlx) }

var fileDescriptorMysqlx = []byte{
	// 700 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xbd, 0x72, 0xeb, 0x44,
	0x14, 0x8e, 0x64, 0xc7, 0xc4, 0xe7, 0xda, 0xce, 0x7a, 0x6f, 0x20, 0xbe, 0x37, 0xe0, 0x78, 0x3c,
	0x14, 0xa1, 0x51, 0x18, 0x3a, 0xd2, 0x29, 0xd2, 0x31, 0xd1, 0xc4, 0xd2, 0x9a, 0xdd, 0x15, 0x09,
	0xd5, 0x4e, 0xb0, 0x85, 0xc7, 0xc1, 0x8e, 0x1c, 0xc9, 0x64, 0x92, 0x87, 0x60, 0x86, 0x92, 0x16,
	0x7a, 0xde, 0x83, 0x0e, 0x1e, 0x01, 0x4c, 0x47, 0xcd, 0x03, 0x30, 0x5a, 0xd9, 0xca, 0xef, 0xc0,
	0xd0, 0xe9, 0x7c, 0xdf, 0xd1, 0xf7, 0x1d, 0x9d, 0x1f, 0x41, 0x6d, 0x76, 0x97, 0x5e, 0x4f, 0x6f,
	0xad, 0x79, 0x12, 0x2f, 0x62, 0x5a, 0xf1, 0x75, 0xf4, 0xb6, 0x33, 0x8e, 0xe3, 0xf1, 0x34, 0x3a,
	0xd4, 0xe8, 0x57, 0xdf, 0x7e, 0x7d, 0x38, 0x8a, 0xd2, 0x61, 0x32, 0x99, 0x2f, 0xe2, 0x24, 0xcf,
	0xec, 0xfe, 0x6d, 0x42, 0xc3, 0x99, 0x4e, 0xa2, 0xab, 0x85, 0x1f, 0xa5, 0xe9, 0xc5, 0x38, 0x4a,
	0xbb, 0x7f, 0x98, 0x50, 0x96, 0x77, 0xf3, 0x88, 0xb6, 0x60, 0xc7, 0x61, 0x81, 0x72, 0xec, 0x81,
	0x7d, 0xec, 0xf5, 0x3d, 0xe9, 0xa1, 0x50, 0x9f, 0xa1, 0x24, 0xc6, 0x8b, 0x8c, 0x40, 0x49, 0x4c,
	0x5a, 0x87, 0xaa, 0x66, 0xfa, 0x4c, 0x20, 0x29, 0xd1, 0x3d, 0xd8, 0x15, 0x28, 0x84, 0xb2, 0x43,
	0x79, 0x82, 0x81, 0xf4, 0x1c, 0x5b, 0xa2, 0x12, 0xd2, 0xe6, 0x92, 0x94, 0x69, 0x1b, 0xde, 0x3e,
	0x27, 0x1d, 0x16, 0x48, 0x2f, 0x08, 0x91, 0x6c, 0xd2, 0x06, 0x80, 0xe6, 0x39, 0x66, 0xda, 0x95,
	0x22, 0xce, 0xc5, 0xdf, 0xa1, 0x3b, 0x40, 0xc4, 0xe7, 0x7d, 0x25, 0xa4, 0x2f, 0x15, 0x9e, 0xa3,
	0x13, 0x4a, 0x24, 0x35, 0x5d, 0x01, 0x0f, 0x5d, 0xd5, 0xf3, 0x02, 0x97, 0x34, 0xe9, 0x36, 0xbc,
	0xd2, 0xa1, 0x17, 0x08, 0xe4, 0x92, 0xd0, 0x02, 0x08, 0x07, 0xae, 0x2d, 0x91, 0xbc, 0x2e, 0x00,
	0x17, 0xfb, 0x28, 0x91, 0xec, 0x64, 0x00, 0x9e, 0x0f, 0xd0, 0x91, 0x8a, 0x0d, 0x30, 0x20, 0x2d,
	0x4a, 0xa0, 0xb6, 0x02, 0x72, 0xeb, 0x37, 0x99, 0xb5, 0x7e, 0xc7, 0xe1, 0x98, 0x15, 0xfd, 0x85,
	0x87, 0x67, 0xa4, 0x5d, 0xa0, 0x3e, 0x73, 0xbd, 0xde, 0x97, 0x39, 0xba, 0x4f, 0x29, 0x34, 0x72,
	0x7d, 0xce, 0x06, 0x39, 0xd6, 0xe9, 0x2e, 0x4d, 0x68, 0x88, 0x28, 0xb9, 0x89, 0x92, 0xa2, 0xed,
	0xbf, 0xae, 0xdb, 0x5e, 0x01, 0x93, 0x9d, 0x92, 0x0d, 0x5a, 0x85, 0x4d, 0xe4, 0x9c, 0x71, 0x62,
	0xd0, 0x77, 0xa1, 0xe9, 0xb0, 0xe0, 0x71, 0xc3, 0x89, 0xf9, 0x1f, 0x0d, 0x2c, 0x65, 0x63, 0x7a,
	0xce, 0xb3, 0x53, 0x52, 0xa6, 0x00, 0x95, 0x80, 0x49, 0xcf, 0x41, 0xf2, 0x2a, 0x53, 0xe1, 0x28,
	0xc2, 0xbe, 0x14, 0x28, 0x95, 0xc3, 0xfa, 0xa1, 0x1f, 0x28, 0x1f, 0xa5, 0xad, 0x5c, 0x5b, 0xda,
	0xa4, 0x46, 0x9b, 0x50, 0xbf, 0xe7, 0x39, 0x3b, 0x23, 0xf5, 0x4c, 0xf8, 0x1e, 0xea, 0xa1, 0x74,
	0x4e, 0x94, 0xcb, 0x02, 0x24, 0x0d, 0xfa, 0x01, 0xbc, 0x79, 0xca, 0x88, 0x50, 0x0c, 0x30, 0x70,
	0xd1, 0x25, 0xdb, 0xf4, 0x00, 0x3e, 0x7c, 0xe9, 0x45, 0xe5, 0x33, 0x8e, 0xaa, 0x60, 0x04, 0x21,
	0x74, 0x17, 0x5e, 0x3f, 0x1d, 0x6e, 0x56, 0x7a, 0xf3, 0xdf, 0x25, 0x58, 0x28, 0xd5, 0xc0, 0xe6,
	0xb6, 0x2f, 0x08, 0xed, 0xbe, 0x0f, 0x26, 0xfb, 0x86, 0x12, 0x28, 0xcd, 0xd2, 0x71, 0xcb, 0xe8,
	0x18, 0x07, 0x55, 0x9e, 0x3d, 0x1e, 0x95, 0xbf, 0xff, 0xeb, 0xe3, 0x8d, 0xee, 0xcf, 0x06, 0x6c,
	0x62, 0x92, 0xc4, 0x09, 0xfd, 0x14, 0xb6, 0xd2, 0xe8, 0x26, 0x4a, 0x26, 0x8b, 0x3b, 0x9d, 0xd6,
	0xf8, 0xe4, 0x3d, 0x2b, 0x3f, 0x20, 0x4b, 0x27, 0x58, 0x62, 0xc5, 0x1e, 0xe5, 0x03, 0xe1, 0x45,
	0x3a, 0xa5, 0x50, 0x1e, 0xc6, 0xa3, 0xa8, 0x65, 0x76, 0xcc, 0x83, 0x3a, 0xd7, 0xcf, 0x6b, 0xc3,
	0x52, 0xc7, 0x5c, 0x19, 0xd2, 0x3d, 0xa8, 0xa6, 0xd7, 0x53, 0x95, 0x2e, 0x2e, 0x16, 0x51, 0xab,
	0xac, 0xf1, 0xad, 0xf4, 0x7a, 0x2a, 0xb2, 0xb8, 0xdb, 0x81, 0xad, 0xb5, 0xfe, 0xfd, 0xc8, 0xf5,
	0xf4, 0x7b, 0xb6, 0xb4, 0xfb, 0xc4, 0xd0, 0xf5, 0x1a, 0x47, 0x97, 0xd0, 0x1c, 0xea, 0x43, 0x55,
	0xb3, 0x7c, 0x65, 0xd4, 0x64, 0x44, 0xf7, 0xad, 0xfc, 0xc2, 0xad, 0xf5, 0x85, 0x5b, 0xab, 0x7d,
	0x62, 0xf3, 0xc5, 0x24, 0xbe, 0x4a, 0x5b, 0x3f, 0x7e, 0x57, 0xd1, 0x5f, 0xb4, 0xb7, 0xfe, 0xa2,
	0xc7, 0xc7, 0x6e, 0x65, 0x1b, 0xc7, 0xb7, 0x87, 0x0f, 0x41, 0x6f, 0x94, 0x79, 0xa5, 0x7a, 0x3b,
	0xff, 0x97, 0xd7, 0x4f, 0x4f, 0xbd, 0x1e, 0x6f, 0xf8, 0xca, 0x2b, 0x7d, 0x08, 0x7a, 0xa3, 0xe3,
	0x8f, 0x7e, 0x59, 0xb6, 0x8d, 0xdf, 0x96, 0x6d, 0xe3, 0xf7, 0x65, 0xdb, 0xf8, 0xe1, 0xcf, 0xf6,
	0x06, 0xec, 0x0e, 0xe3, 0x99, 0xa5, 0xff, 0x67, 0xd6, 0xf0, 0xd2, 0xba, 0x2d, 0xfc, 0xfe, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x8f, 0x02, 0x12, 0x97, 0xe5, 0x04, 0x00, 0x00,
}
