// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service/verify.proto

package service

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/nspcc-dev/neofs-api-go/refs"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Verb is an enumeration of session request types
type Token_Info_Verb int32

const (
	// Put refers to object.Put RPC call
	Token_Info_Put Token_Info_Verb = 0
	// Get refers to object.Get RPC call
	Token_Info_Get Token_Info_Verb = 1
	// Head refers to object.Head RPC call
	Token_Info_Head Token_Info_Verb = 2
	// Search refers to object.Search RPC call
	Token_Info_Search Token_Info_Verb = 3
	// Delete refers to object.Delete RPC call
	Token_Info_Delete Token_Info_Verb = 4
	// Range refers to object.GetRange RPC call
	Token_Info_Range Token_Info_Verb = 5
	// RangeHash refers to object.GetRangeHash RPC call
	Token_Info_RangeHash Token_Info_Verb = 6
)

var Token_Info_Verb_name = map[int32]string{
	0: "Put",
	1: "Get",
	2: "Head",
	3: "Search",
	4: "Delete",
	5: "Range",
	6: "RangeHash",
}

var Token_Info_Verb_value = map[string]int32{
	"Put":       0,
	"Get":       1,
	"Head":      2,
	"Search":    3,
	"Delete":    4,
	"Range":     5,
	"RangeHash": 6,
}

func (x Token_Info_Verb) String() string {
	return proto.EnumName(Token_Info_Verb_name, int32(x))
}

func (Token_Info_Verb) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4bdd5bc50ec96238, []int{1, 0, 0}
}

// RequestVerificationHeader is a set of signatures of every NeoFS Node that processed request
// (should be embedded into message).
type RequestVerificationHeader struct {
	// Signatures is a set of signatures of every passed NeoFS Node
	Signatures []*RequestVerificationHeader_Signature `protobuf:"bytes,1,rep,name=Signatures,proto3" json:"Signatures,omitempty"`
	// Token is a token of the session within which the request is sent
	Token                *Token   `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestVerificationHeader) Reset()         { *m = RequestVerificationHeader{} }
func (m *RequestVerificationHeader) String() string { return proto.CompactTextString(m) }
func (*RequestVerificationHeader) ProtoMessage()    {}
func (*RequestVerificationHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bdd5bc50ec96238, []int{0}
}
func (m *RequestVerificationHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestVerificationHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *RequestVerificationHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestVerificationHeader.Merge(m, src)
}
func (m *RequestVerificationHeader) XXX_Size() int {
	return m.Size()
}
func (m *RequestVerificationHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestVerificationHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RequestVerificationHeader proto.InternalMessageInfo

func (m *RequestVerificationHeader) GetSignatures() []*RequestVerificationHeader_Signature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *RequestVerificationHeader) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type RequestVerificationHeader_Signature struct {
	// Sign is signature of the request or session key.
	Sign []byte `protobuf:"bytes,1,opt,name=Sign,proto3" json:"Sign,omitempty"`
	// Peer is compressed public key used for signature.
	Peer                 []byte   `protobuf:"bytes,2,opt,name=Peer,proto3" json:"Peer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestVerificationHeader_Signature) Reset()         { *m = RequestVerificationHeader_Signature{} }
func (m *RequestVerificationHeader_Signature) String() string { return proto.CompactTextString(m) }
func (*RequestVerificationHeader_Signature) ProtoMessage()    {}
func (*RequestVerificationHeader_Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bdd5bc50ec96238, []int{0, 0}
}
func (m *RequestVerificationHeader_Signature) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestVerificationHeader_Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *RequestVerificationHeader_Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestVerificationHeader_Signature.Merge(m, src)
}
func (m *RequestVerificationHeader_Signature) XXX_Size() int {
	return m.Size()
}
func (m *RequestVerificationHeader_Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestVerificationHeader_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_RequestVerificationHeader_Signature proto.InternalMessageInfo

func (m *RequestVerificationHeader_Signature) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *RequestVerificationHeader_Signature) GetPeer() []byte {
	if m != nil {
		return m.Peer
	}
	return nil
}

// User token granting rights for object manipulation
type Token struct {
	// TokenInfo is a grouped information about token
	Token_Info `protobuf:"bytes,1,opt,name=TokenInfo,proto3,embedded=TokenInfo" json:"TokenInfo"`
	// Signature is a signature of session token information
	Signature            []byte   `protobuf:"bytes,8,opt,name=Signature,proto3" json:"Signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bdd5bc50ec96238, []int{1}
}
func (m *Token) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return m.Size()
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Token_Info struct {
	// ID is a token identifier. valid UUIDv4 represented in bytes
	ID TokenID `protobuf:"bytes,1,opt,name=ID,proto3,customtype=TokenID" json:"ID"`
	// OwnerID is an owner of manipulation object
	OwnerID OwnerID `protobuf:"bytes,2,opt,name=OwnerID,proto3,customtype=OwnerID" json:"OwnerID"`
	// Verb is a type of request for which the token is issued
	Verb Token_Info_Verb `protobuf:"varint,3,opt,name=verb,proto3,enum=service.Token_Info_Verb" json:"verb,omitempty"`
	// Address is an object address for which token is issued
	Address Address `protobuf:"bytes,4,opt,name=Address,proto3,customtype=Address" json:"Address"`
	// Created is an initial epoch of token lifetime
	Created uint64 `protobuf:"varint,5,opt,name=Created,proto3" json:"Created,omitempty"`
	// ValidUntil is a last epoch of token lifetime
	ValidUntil uint64 `protobuf:"varint,6,opt,name=ValidUntil,proto3" json:"ValidUntil,omitempty"`
	// SessionKey is a public key of session key
	SessionKey           []byte   `protobuf:"bytes,7,opt,name=SessionKey,proto3" json:"SessionKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token_Info) Reset()         { *m = Token_Info{} }
func (m *Token_Info) String() string { return proto.CompactTextString(m) }
func (*Token_Info) ProtoMessage()    {}
func (*Token_Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bdd5bc50ec96238, []int{1, 0}
}
func (m *Token_Info) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Token_Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Token_Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token_Info.Merge(m, src)
}
func (m *Token_Info) XXX_Size() int {
	return m.Size()
}
func (m *Token_Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Token_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Token_Info proto.InternalMessageInfo

func (m *Token_Info) GetVerb() Token_Info_Verb {
	if m != nil {
		return m.Verb
	}
	return Token_Info_Put
}

func (m *Token_Info) GetCreated() uint64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Token_Info) GetValidUntil() uint64 {
	if m != nil {
		return m.ValidUntil
	}
	return 0
}

func (m *Token_Info) GetSessionKey() []byte {
	if m != nil {
		return m.SessionKey
	}
	return nil
}

func init() {
	proto.RegisterEnum("service.Token_Info_Verb", Token_Info_Verb_name, Token_Info_Verb_value)
	proto.RegisterType((*RequestVerificationHeader)(nil), "service.RequestVerificationHeader")
	proto.RegisterType((*RequestVerificationHeader_Signature)(nil), "service.RequestVerificationHeader.Signature")
	proto.RegisterType((*Token)(nil), "service.Token")
	proto.RegisterType((*Token_Info)(nil), "service.Token.Info")
}

func init() { proto.RegisterFile("service/verify.proto", fileDescriptor_4bdd5bc50ec96238) }

var fileDescriptor_4bdd5bc50ec96238 = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4d, 0x8e, 0xd3, 0x4c,
	0x10, 0x86, 0xa7, 0x1d, 0x27, 0x4e, 0x6a, 0x7e, 0x3e, 0x7f, 0x0d, 0x0b, 0x13, 0x21, 0x27, 0x8a,
	0x58, 0x64, 0x24, 0x62, 0x4b, 0x19, 0x09, 0x21, 0xb1, 0x9a, 0x10, 0xc1, 0x44, 0x20, 0x88, 0xda,
	0x43, 0x16, 0xec, 0x1c, 0xbb, 0xe2, 0x58, 0x04, 0x77, 0xe8, 0x76, 0x82, 0x72, 0x13, 0xce, 0xc0,
	0x39, 0x58, 0xcc, 0x72, 0x96, 0xc0, 0x22, 0x42, 0xe1, 0x0a, 0x1c, 0x00, 0xb9, 0xed, 0xfc, 0x20,
	0xc1, 0xee, 0xad, 0xa7, 0xaa, 0xde, 0xaa, 0x6e, 0x15, 0xdc, 0x95, 0x28, 0x96, 0x71, 0x80, 0xee,
	0x12, 0x45, 0x3c, 0x59, 0x39, 0x73, 0xc1, 0x53, 0x4e, 0x8d, 0x82, 0xd6, 0x4d, 0x81, 0x13, 0xe9,
	0xa6, 0xab, 0x39, 0xca, 0x3c, 0x55, 0xef, 0x44, 0x71, 0x3a, 0x5d, 0x8c, 0x9d, 0x80, 0xbf, 0x77,
	0x23, 0x1e, 0x71, 0x57, 0xe1, 0xf1, 0x62, 0xa2, 0x22, 0x15, 0x28, 0x95, 0x97, 0xb7, 0xbe, 0x10,
	0xb8, 0xc7, 0xf0, 0xc3, 0x02, 0x65, 0x3a, 0xca, 0x26, 0xc4, 0x81, 0x9f, 0xc6, 0x3c, 0xb9, 0x42,
	0x3f, 0x44, 0x41, 0x5f, 0x02, 0x78, 0x71, 0x94, 0xf8, 0xe9, 0x42, 0xa0, 0xb4, 0x48, 0xb3, 0xd4,
	0x3e, 0xee, 0x3e, 0x74, 0x8a, 0xe1, 0xce, 0x3f, 0xfb, 0x9c, 0x5d, 0x13, 0x3b, 0xe8, 0xa7, 0x0f,
	0xa0, 0x7c, 0xcd, 0xdf, 0x61, 0x62, 0x69, 0x4d, 0xd2, 0x3e, 0xee, 0x9e, 0xed, 0x8c, 0x14, 0x65,
	0x79, 0xb2, 0x7e, 0x01, 0xb5, 0x5d, 0x0f, 0xa5, 0xa0, 0x67, 0x81, 0x45, 0x9a, 0xa4, 0x7d, 0xc2,
	0x94, 0xce, 0xd8, 0x10, 0x51, 0x28, 0x97, 0x13, 0xa6, 0x74, 0xeb, 0x5b, 0xa9, 0xf0, 0xa6, 0x4f,
	0xa0, 0xa6, 0xc4, 0x20, 0x99, 0x70, 0xd5, 0x76, 0xdc, 0xbd, 0xf3, 0xe7, 0x20, 0x27, 0x4b, 0xf5,
	0xaa, 0x37, 0xeb, 0xc6, 0xd1, 0xed, 0xba, 0x41, 0xd8, 0xbe, 0x9e, 0xde, 0x3f, 0x98, 0x6d, 0x55,
	0x95, 0xff, 0x1e, 0xd4, 0x7f, 0x69, 0xa0, 0xab, 0xb2, 0x06, 0x68, 0x83, 0x7e, 0xbe, 0x53, 0xef,
	0xbf, 0xcc, 0xe7, 0xfb, 0xba, 0x61, 0xe4, 0x2e, 0x7d, 0xa6, 0x0d, 0xfa, 0xf4, 0x1c, 0x8c, 0xd7,
	0x1f, 0x13, 0x14, 0x83, 0x7e, 0xbe, 0xe5, 0xbe, 0xaa, 0xc0, 0x6c, 0x2b, 0xe8, 0x23, 0xd0, 0x97,
	0x28, 0xc6, 0x56, 0xa9, 0x49, 0xda, 0x67, 0x5d, 0xeb, 0x2f, 0xab, 0x3a, 0x23, 0x14, 0xe3, 0x5e,
	0x75, 0xb3, 0x6e, 0xe8, 0x99, 0x62, 0xaa, 0x9e, 0x3e, 0x06, 0xe3, 0x32, 0x0c, 0x05, 0x4a, 0x69,
	0xe9, 0xea, 0x95, 0xa7, 0x4e, 0x76, 0x0b, 0x4e, 0x01, 0xf7, 0x13, 0x0b, 0xc0, 0xb6, 0x82, 0x5a,
	0x60, 0x3c, 0x15, 0xe8, 0xa7, 0x18, 0x5a, 0xe5, 0x26, 0x69, 0xeb, 0x6c, 0x1b, 0x52, 0x1b, 0x60,
	0xe4, 0xcf, 0xe2, 0xf0, 0x4d, 0x92, 0xc6, 0x33, 0xab, 0xa2, 0x92, 0x07, 0x24, 0xcb, 0x7b, 0x28,
	0x65, 0xcc, 0x93, 0x17, 0xb8, 0xb2, 0x0c, 0xf5, 0x3f, 0x07, 0xa4, 0x75, 0x0d, 0x6a, 0x43, 0x6a,
	0x40, 0x69, 0xb8, 0x48, 0xcd, 0xa3, 0x4c, 0x3c, 0xc7, 0xd4, 0x24, 0xb4, 0x0a, 0x7a, 0x76, 0x1a,
	0xa6, 0x46, 0x01, 0x2a, 0x1e, 0xfa, 0x22, 0x98, 0x9a, 0xa5, 0x4c, 0xf7, 0x71, 0x86, 0x29, 0x9a,
	0x3a, 0xad, 0x41, 0x99, 0xf9, 0x49, 0x84, 0x66, 0x99, 0x9e, 0x42, 0x4d, 0xc9, 0x2b, 0x5f, 0x4e,
	0xcd, 0x4a, 0xcf, 0xbb, 0xd9, 0xd8, 0xe4, 0x76, 0x63, 0x93, 0xaf, 0x1b, 0x9b, 0xfc, 0xd8, 0xd8,
	0xe4, 0xd3, 0x4f, 0xfb, 0xe8, 0xed, 0xf9, 0xc1, 0x9d, 0x27, 0x72, 0x1e, 0x04, 0x9d, 0x10, 0x97,
	0x6e, 0x82, 0x7c, 0x22, 0x3b, 0xfe, 0x3c, 0xee, 0x44, 0xdc, 0x2d, 0xbe, 0xf2, 0xb3, 0xf6, 0xff,
	0x2b, 0xe4, 0xcf, 0x3c, 0xe7, 0x72, 0x38, 0x70, 0xbc, 0x9c, 0x8d, 0x2b, 0xea, 0xfc, 0x2f, 0x7e,
	0x07, 0x00, 0x00, 0xff, 0xff, 0xea, 0xcf, 0xa5, 0xdd, 0x60, 0x03, 0x00, 0x00,
}

func (m *RequestVerificationHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestVerificationHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestVerificationHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Token != nil {
		{
			size, err := m.Token.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVerify(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signatures) > 0 {
		for iNdEx := len(m.Signatures) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Signatures[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVerify(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *RequestVerificationHeader_Signature) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestVerificationHeader_Signature) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestVerificationHeader_Signature) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Peer) > 0 {
		i -= len(m.Peer)
		copy(dAtA[i:], m.Peer)
		i = encodeVarintVerify(dAtA, i, uint64(len(m.Peer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sign) > 0 {
		i -= len(m.Sign)
		copy(dAtA[i:], m.Sign)
		i = encodeVarintVerify(dAtA, i, uint64(len(m.Sign)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Token) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Token) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Token) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintVerify(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x42
	}
	{
		size, err := m.Token_Info.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintVerify(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Token_Info) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Token_Info) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Token_Info) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.SessionKey) > 0 {
		i -= len(m.SessionKey)
		copy(dAtA[i:], m.SessionKey)
		i = encodeVarintVerify(dAtA, i, uint64(len(m.SessionKey)))
		i--
		dAtA[i] = 0x3a
	}
	if m.ValidUntil != 0 {
		i = encodeVarintVerify(dAtA, i, uint64(m.ValidUntil))
		i--
		dAtA[i] = 0x30
	}
	if m.Created != 0 {
		i = encodeVarintVerify(dAtA, i, uint64(m.Created))
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.Address.Size()
		i -= size
		if _, err := m.Address.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintVerify(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Verb != 0 {
		i = encodeVarintVerify(dAtA, i, uint64(m.Verb))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.OwnerID.Size()
		i -= size
		if _, err := m.OwnerID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintVerify(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.ID.Size()
		i -= size
		if _, err := m.ID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintVerify(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintVerify(dAtA []byte, offset int, v uint64) int {
	offset -= sovVerify(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RequestVerificationHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Signatures) > 0 {
		for _, e := range m.Signatures {
			l = e.Size()
			n += 1 + l + sovVerify(uint64(l))
		}
	}
	if m.Token != nil {
		l = m.Token.Size()
		n += 1 + l + sovVerify(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RequestVerificationHeader_Signature) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sign)
	if l > 0 {
		n += 1 + l + sovVerify(uint64(l))
	}
	l = len(m.Peer)
	if l > 0 {
		n += 1 + l + sovVerify(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Token) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Token_Info.Size()
	n += 1 + l + sovVerify(uint64(l))
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovVerify(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Token_Info) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovVerify(uint64(l))
	l = m.OwnerID.Size()
	n += 1 + l + sovVerify(uint64(l))
	if m.Verb != 0 {
		n += 1 + sovVerify(uint64(m.Verb))
	}
	l = m.Address.Size()
	n += 1 + l + sovVerify(uint64(l))
	if m.Created != 0 {
		n += 1 + sovVerify(uint64(m.Created))
	}
	if m.ValidUntil != 0 {
		n += 1 + sovVerify(uint64(m.ValidUntil))
	}
	l = len(m.SessionKey)
	if l > 0 {
		n += 1 + l + sovVerify(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovVerify(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVerify(x uint64) (n int) {
	return sovVerify(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RequestVerificationHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVerify
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
			return fmt.Errorf("proto: RequestVerificationHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestVerificationHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signatures", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerify
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthVerify
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVerify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signatures = append(m.Signatures, &RequestVerificationHeader_Signature{})
			if err := m.Signatures[len(m.Signatures)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerify
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthVerify
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVerify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Token == nil {
				m.Token = &Token{}
			}
			if err := m.Token.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVerify(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVerify
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthVerify
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
func (m *RequestVerificationHeader_Signature) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVerify
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
			return fmt.Errorf("proto: Signature: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Signature: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sign", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerify
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
				return ErrInvalidLengthVerify
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthVerify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sign = append(m.Sign[:0], dAtA[iNdEx:postIndex]...)
			if m.Sign == nil {
				m.Sign = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Peer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerify
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
				return ErrInvalidLengthVerify
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthVerify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Peer = append(m.Peer[:0], dAtA[iNdEx:postIndex]...)
			if m.Peer == nil {
				m.Peer = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVerify(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVerify
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthVerify
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
func (m *Token) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVerify
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
			return fmt.Errorf("proto: Token: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Token: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token_Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerify
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthVerify
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVerify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Token_Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerify
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
				return ErrInvalidLengthVerify
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthVerify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVerify(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVerify
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthVerify
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
func (m *Token_Info) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVerify
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
			return fmt.Errorf("proto: Info: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Info: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerify
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
				return ErrInvalidLengthVerify
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthVerify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerify
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
				return ErrInvalidLengthVerify
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthVerify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OwnerID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Verb", wireType)
			}
			m.Verb = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerify
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Verb |= Token_Info_Verb(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerify
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthVerify
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVerify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Address.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			m.Created = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerify
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Created |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidUntil", wireType)
			}
			m.ValidUntil = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerify
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidUntil |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerify
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
				return ErrInvalidLengthVerify
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthVerify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionKey = append(m.SessionKey[:0], dAtA[iNdEx:postIndex]...)
			if m.SessionKey == nil {
				m.SessionKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVerify(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVerify
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthVerify
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
func skipVerify(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVerify
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
					return 0, ErrIntOverflowVerify
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
					return 0, ErrIntOverflowVerify
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
				return 0, ErrInvalidLengthVerify
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVerify
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVerify
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVerify        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVerify          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVerify = fmt.Errorf("proto: unexpected end of group")
)
