// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protos.proto

package shared

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Unit struct {
}

func (m *Unit) Reset()      { *m = Unit{} }
func (*Unit) ProtoMessage() {}
func (*Unit) Descriptor() ([]byte, []int) {
	return fileDescriptor_5da3cbeb884d181c, []int{0}
}
func (m *Unit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Unit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Unit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Unit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Unit.Merge(m, src)
}
func (m *Unit) XXX_Size() int {
	return m.Size()
}
func (m *Unit) XXX_DiscardUnknown() {
	xxx_messageInfo_Unit.DiscardUnknown(m)
}

var xxx_messageInfo_Unit proto.InternalMessageInfo

type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()      { *m = HelloRequest{} }
func (*HelloRequest) ProtoMessage() {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5da3cbeb884d181c, []int{1}
}
func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return m.Size()
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *HelloResponse) Reset()      { *m = HelloResponse{} }
func (*HelloResponse) ProtoMessage() {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5da3cbeb884d181c, []int{2}
}
func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return m.Size()
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type AddRequest struct {
	A float64 `protobuf:"fixed64,1,opt,name=a,proto3" json:"a,omitempty"`
	B float64 `protobuf:"fixed64,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (m *AddRequest) Reset()      { *m = AddRequest{} }
func (*AddRequest) ProtoMessage() {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5da3cbeb884d181c, []int{3}
}
func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return m.Size()
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetA() float64 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *AddRequest) GetB() float64 {
	if m != nil {
		return m.B
	}
	return 0
}

type AddResponse struct {
	Result float64 `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *AddResponse) Reset()      { *m = AddResponse{} }
func (*AddResponse) ProtoMessage() {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5da3cbeb884d181c, []int{4}
}
func (m *AddResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResponse.Merge(m, src)
}
func (m *AddResponse) XXX_Size() int {
	return m.Size()
}
func (m *AddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddResponse proto.InternalMessageInfo

func (m *AddResponse) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Unit)(nil), "shared.Unit")
	proto.RegisterType((*HelloRequest)(nil), "shared.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "shared.HelloResponse")
	proto.RegisterType((*AddRequest)(nil), "shared.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "shared.AddResponse")
}

func init() { proto.RegisterFile("protos.proto", fileDescriptor_5da3cbeb884d181c) }

var fileDescriptor_5da3cbeb884d181c = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xbb, 0x4e, 0xf4, 0x30,
	0x10, 0x85, 0x3d, 0xff, 0xbf, 0x84, 0x65, 0x08, 0x8d, 0xb9, 0x28, 0xda, 0x62, 0x84, 0x2c, 0x21,
	0x2d, 0x4d, 0x84, 0x80, 0x86, 0x72, 0x29, 0x10, 0x75, 0x10, 0xf4, 0x0e, 0xb6, 0x60, 0xa5, 0x6c,
	0xb2, 0xc4, 0x49, 0x41, 0xc7, 0x23, 0xf0, 0x10, 0x14, 0x3c, 0x0a, 0x65, 0xca, 0x2d, 0x89, 0xd3,
	0x50, 0xee, 0x23, 0x20, 0x9c, 0x84, 0x9b, 0xa8, 0x3c, 0xe7, 0xf8, 0x9b, 0xe3, 0xb1, 0x8d, 0xfe,
	0x3c, 0xcf, 0x8a, 0xcc, 0x84, 0x6e, 0xe1, 0x9e, 0xb9, 0x95, 0xb9, 0x56, 0xc2, 0xc3, 0xc1, 0x65,
	0x3a, 0x2d, 0x84, 0x40, 0xff, 0x5c, 0x27, 0x49, 0x16, 0xe9, 0xbb, 0x52, 0x9b, 0x82, 0x73, 0x1c,
	0xa4, 0x72, 0xa6, 0x03, 0xd8, 0x85, 0xf1, 0x5a, 0xe4, 0x6a, 0xb1, 0x8f, 0x1b, 0x1d, 0x63, 0xe6,
	0x59, 0x6a, 0x34, 0x0f, 0x70, 0x75, 0xa6, 0x8d, 0x91, 0x37, 0x3d, 0xd7, 0x4b, 0x31, 0x46, 0x9c,
	0x28, 0xd5, 0x87, 0xf9, 0x08, 0xd2, 0x11, 0x10, 0x81, 0xfc, 0x50, 0x71, 0xf0, 0xaf, 0x55, 0xb1,
	0xd8, 0xc3, 0x75, 0x47, 0x76, 0x91, 0x3b, 0xe8, 0xe5, 0xda, 0x94, 0x49, 0xd1, 0xf1, 0x9d, 0x3a,
	0x7c, 0x02, 0x5c, 0x71, 0x87, 0xf3, 0x13, 0x1c, 0x5e, 0xc8, 0xfb, 0xb6, 0xde, 0x0a, 0xdb, 0x6b,
	0x84, 0xdf, 0x67, 0x1f, 0x6d, 0xff, 0x72, 0xdb, 0x68, 0xc1, 0xf8, 0x01, 0xfe, 0x9f, 0x28, 0xc5,
	0x79, 0xbf, 0xff, 0x35, 0xe2, 0x68, 0xf3, 0x87, 0xf7, 0xd9, 0x11, 0xe2, 0xf0, 0x2a, 0x9b, 0xaa,
	0xb3, 0x32, 0xbd, 0xfe, 0xb3, 0xcd, 0xef, 0x3d, 0xf7, 0x88, 0xec, 0xf4, 0xb8, 0xaa, 0x89, 0x2d,
	0x6a, 0x62, 0xcb, 0x9a, 0xe0, 0xc1, 0x12, 0x3c, 0x5b, 0x82, 0x17, 0x4b, 0x50, 0x59, 0x82, 0x57,
	0x4b, 0xf0, 0x66, 0x89, 0x2d, 0x2d, 0xc1, 0x63, 0x43, 0xac, 0x6a, 0x88, 0x2d, 0x1a, 0x62, 0xb1,
	0xe7, 0xfe, 0xe4, 0xe8, 0x3d, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xde, 0xee, 0xcb, 0xa3, 0x01, 0x00,
	0x00,
}

func (this *Unit) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Unit)
	if !ok {
		that2, ok := that.(Unit)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *HelloRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HelloRequest)
	if !ok {
		that2, ok := that.(HelloRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	return true
}
func (this *HelloResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HelloResponse)
	if !ok {
		that2, ok := that.(HelloResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	return true
}
func (this *AddRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AddRequest)
	if !ok {
		that2, ok := that.(AddRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.A != that1.A {
		return false
	}
	if this.B != that1.B {
		return false
	}
	return true
}
func (this *AddResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AddResponse)
	if !ok {
		that2, ok := that.(AddResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Result != that1.Result {
		return false
	}
	return true
}
func (this *Unit) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&shared.Unit{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *HelloRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&shared.HelloRequest{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *HelloResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&shared.HelloResponse{")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *AddRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&shared.AddRequest{")
	s = append(s, "A: "+fmt.Sprintf("%#v", this.A)+",\n")
	s = append(s, "B: "+fmt.Sprintf("%#v", this.B)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *AddResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&shared.AddResponse{")
	s = append(s, "Result: "+fmt.Sprintf("%#v", this.Result)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringProtos(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Unit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Unit) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *HelloRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HelloRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProtos(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func (m *HelloResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HelloResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProtos(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	return i, nil
}

func (m *AddRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.A != 0 {
		dAtA[i] = 0x9
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.A))))
		i += 8
	}
	if m.B != 0 {
		dAtA[i] = 0x11
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.B))))
		i += 8
	}
	return i, nil
}

func (m *AddResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Result != 0 {
		dAtA[i] = 0x9
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Result))))
		i += 8
	}
	return i, nil
}

func encodeVarintProtos(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Unit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *HelloRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProtos(uint64(l))
	}
	return n
}

func (m *HelloResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovProtos(uint64(l))
	}
	return n
}

func (m *AddRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.A != 0 {
		n += 9
	}
	if m.B != 0 {
		n += 9
	}
	return n
}

func (m *AddResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Result != 0 {
		n += 9
	}
	return n
}

func sovProtos(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProtos(x uint64) (n int) {
	return sovProtos(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Unit) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Unit{`,
		`}`,
	}, "")
	return s
}
func (this *HelloRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HelloRequest{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HelloResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HelloResponse{`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AddRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AddRequest{`,
		`A:` + fmt.Sprintf("%v", this.A) + `,`,
		`B:` + fmt.Sprintf("%v", this.B) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AddResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AddResponse{`,
		`Result:` + fmt.Sprintf("%v", this.Result) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringProtos(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Unit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtos
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
			return fmt.Errorf("proto: Unit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Unit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipProtos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtos
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProtos
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
func (m *HelloRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtos
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
			return fmt.Errorf("proto: HelloRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HelloRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtos
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
				return ErrInvalidLengthProtos
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtos
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtos
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProtos
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
func (m *HelloResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtos
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
			return fmt.Errorf("proto: HelloResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HelloResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtos
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
				return ErrInvalidLengthProtos
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtos
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtos
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProtos
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
func (m *AddRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtos
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
			return fmt.Errorf("proto: AddRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.A = float64(math.Float64frombits(v))
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.B = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipProtos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtos
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProtos
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
func (m *AddResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtos
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
			return fmt.Errorf("proto: AddResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Result = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipProtos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtos
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProtos
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
func skipProtos(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProtos
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
					return 0, ErrIntOverflowProtos
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
					return 0, ErrIntOverflowProtos
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
				return 0, ErrInvalidLengthProtos
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthProtos
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProtos
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
				next, err := skipProtos(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthProtos
				}
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
	ErrInvalidLengthProtos = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProtos   = fmt.Errorf("proto: integer overflow")
)
