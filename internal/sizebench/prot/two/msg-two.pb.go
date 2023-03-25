// Copyright (C) 2021 Storx Labs, Inc.
// See LICENSE for copying information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: msg-two.proto

package two

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Language int32

const (
	Language_UNKNOWN Language = 0
	Language_ENGLISH Language = 1
	Language_SPANISH Language = 3
	Language_FRENCH  Language = 4
	Language_GERMAN  Language = 5
)

// Enum value maps for Language.
var (
	Language_name = map[int32]string{
		0: "UNKNOWN",
		1: "ENGLISH",
		3: "SPANISH",
		4: "FRENCH",
		5: "GERMAN",
	}
	Language_value = map[string]int32{
		"UNKNOWN": 0,
		"ENGLISH": 1,
		"SPANISH": 3,
		"FRENCH":  4,
		"GERMAN":  5,
	}
)

func (x Language) Enum() *Language {
	p := new(Language)
	*p = x
	return p
}

func (x Language) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Language) Descriptor() protoreflect.EnumDescriptor {
	return file_msg_two_proto_enumTypes[0].Descriptor()
}

func (Language) Type() protoreflect.EnumType {
	return &file_msg_two_proto_enumTypes[0]
}

func (x Language) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Language.Descriptor instead.
func (Language) EnumDescriptor() ([]byte, []int) {
	return file_msg_two_proto_rawDescGZIP(), []int{0}
}

type Language2 int32

const (
	Language2_UNKNOWN2 Language2 = 0
	Language2_ENGLISH2 Language2 = 1
	Language2_SPANISH2 Language2 = 3
	Language2_FRENCH2  Language2 = 4
	Language2_GERMAN2  Language2 = 5
)

// Enum value maps for Language2.
var (
	Language2_name = map[int32]string{
		0: "UNKNOWN2",
		1: "ENGLISH2",
		3: "SPANISH2",
		4: "FRENCH2",
		5: "GERMAN2",
	}
	Language2_value = map[string]int32{
		"UNKNOWN2": 0,
		"ENGLISH2": 1,
		"SPANISH2": 3,
		"FRENCH2":  4,
		"GERMAN2":  5,
	}
)

func (x Language2) Enum() *Language2 {
	p := new(Language2)
	*p = x
	return p
}

func (x Language2) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Language2) Descriptor() protoreflect.EnumDescriptor {
	return file_msg_two_proto_enumTypes[1].Descriptor()
}

func (Language2) Type() protoreflect.EnumType {
	return &file_msg_two_proto_enumTypes[1]
}

func (x Language2) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Language2.Descriptor instead.
func (Language2) EnumDescriptor() ([]byte, []int) {
	return file_msg_two_proto_rawDescGZIP(), []int{1}
}

type Nop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Nop) Reset() {
	*x = Nop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_two_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nop) ProtoMessage() {}

func (x *Nop) ProtoReflect() protoreflect.Message {
	mi := &file_msg_two_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nop.ProtoReflect.Descriptor instead.
func (*Nop) Descriptor() ([]byte, []int) {
	return file_msg_two_proto_rawDescGZIP(), []int{0}
}

type Types struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Int32     int32             `protobuf:"varint,1,opt,name=int32,proto3" json:"int32,omitempty"`
	Int64     int64             `protobuf:"varint,2,opt,name=int64,proto3" json:"int64,omitempty"`
	Uint32    uint32            `protobuf:"varint,3,opt,name=uint32,proto3" json:"uint32,omitempty"`
	Uint64    uint64            `protobuf:"varint,4,opt,name=uint64,proto3" json:"uint64,omitempty"`
	Sint32    int32             `protobuf:"zigzag32,5,opt,name=sint32,proto3" json:"sint32,omitempty"`
	Sint64    int64             `protobuf:"zigzag64,6,opt,name=sint64,proto3" json:"sint64,omitempty"`
	Fixed32   uint32            `protobuf:"fixed32,7,opt,name=fixed32,proto3" json:"fixed32,omitempty"`
	Fixed64   uint64            `protobuf:"fixed64,8,opt,name=fixed64,proto3" json:"fixed64,omitempty"`
	Sfixed32  int32             `protobuf:"fixed32,9,opt,name=sfixed32,proto3" json:"sfixed32,omitempty"`
	Sfixed64  int64             `protobuf:"fixed64,10,opt,name=sfixed64,proto3" json:"sfixed64,omitempty"`
	Float     float32           `protobuf:"fixed32,11,opt,name=float,proto3" json:"float,omitempty"`
	Double    float64           `protobuf:"fixed64,12,opt,name=double,proto3" json:"double,omitempty"`
	Bool      bool              `protobuf:"varint,13,opt,name=bool,proto3" json:"bool,omitempty"`
	String_   string            `protobuf:"bytes,14,opt,name=string,proto3" json:"string,omitempty"`
	Bytes     []byte            `protobuf:"bytes,15,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Message   *Message          `protobuf:"bytes,16,opt,name=message,proto3" json:"message,omitempty"`
	Language  Language          `protobuf:"varint,17,opt,name=language,proto3,enum=two.Language" json:"language,omitempty"`
	Int32S    []int32           `protobuf:"varint,18,rep,packed,name=int32s,proto3" json:"int32s,omitempty"`
	Int64S    []int64           `protobuf:"varint,19,rep,packed,name=int64s,proto3" json:"int64s,omitempty"`
	Uint32S   []uint32          `protobuf:"varint,20,rep,packed,name=uint32s,proto3" json:"uint32s,omitempty"`
	Uint64S   []uint64          `protobuf:"varint,21,rep,packed,name=uint64s,proto3" json:"uint64s,omitempty"`
	Sint32S   []int32           `protobuf:"zigzag32,22,rep,packed,name=sint32s,proto3" json:"sint32s,omitempty"`
	Sint64S   []int64           `protobuf:"zigzag64,23,rep,packed,name=sint64s,proto3" json:"sint64s,omitempty"`
	Fixed32S  []uint32          `protobuf:"fixed32,24,rep,packed,name=fixed32s,proto3" json:"fixed32s,omitempty"`
	Fixed64S  []uint64          `protobuf:"fixed64,25,rep,packed,name=fixed64s,proto3" json:"fixed64s,omitempty"`
	Sfixed32S []int32           `protobuf:"fixed32,26,rep,packed,name=sfixed32s,proto3" json:"sfixed32s,omitempty"`
	Sfixed64S []int64           `protobuf:"fixed64,27,rep,packed,name=sfixed64s,proto3" json:"sfixed64s,omitempty"`
	Floats    []float32         `protobuf:"fixed32,28,rep,packed,name=floats,proto3" json:"floats,omitempty"`
	Doubles   []float64         `protobuf:"fixed64,29,rep,packed,name=doubles,proto3" json:"doubles,omitempty"`
	Bools     []bool            `protobuf:"varint,30,rep,packed,name=bools,proto3" json:"bools,omitempty"`
	Strings   []string          `protobuf:"bytes,31,rep,name=strings,proto3" json:"strings,omitempty"`
	Bytess    [][]byte          `protobuf:"bytes,32,rep,name=bytess,proto3" json:"bytess,omitempty"`
	Messages  []*Message        `protobuf:"bytes,33,rep,name=messages,proto3" json:"messages,omitempty"`
	Languages []Language        `protobuf:"varint,34,rep,packed,name=languages,proto3,enum=two.Language" json:"languages,omitempty"`
	Values    map[string]string `protobuf:"bytes,35,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Types) Reset() {
	*x = Types{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_two_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Types) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Types) ProtoMessage() {}

func (x *Types) ProtoReflect() protoreflect.Message {
	mi := &file_msg_two_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Types.ProtoReflect.Descriptor instead.
func (*Types) Descriptor() ([]byte, []int) {
	return file_msg_two_proto_rawDescGZIP(), []int{1}
}

func (x *Types) GetInt32() int32 {
	if x != nil {
		return x.Int32
	}
	return 0
}

func (x *Types) GetInt64() int64 {
	if x != nil {
		return x.Int64
	}
	return 0
}

func (x *Types) GetUint32() uint32 {
	if x != nil {
		return x.Uint32
	}
	return 0
}

func (x *Types) GetUint64() uint64 {
	if x != nil {
		return x.Uint64
	}
	return 0
}

func (x *Types) GetSint32() int32 {
	if x != nil {
		return x.Sint32
	}
	return 0
}

func (x *Types) GetSint64() int64 {
	if x != nil {
		return x.Sint64
	}
	return 0
}

func (x *Types) GetFixed32() uint32 {
	if x != nil {
		return x.Fixed32
	}
	return 0
}

func (x *Types) GetFixed64() uint64 {
	if x != nil {
		return x.Fixed64
	}
	return 0
}

func (x *Types) GetSfixed32() int32 {
	if x != nil {
		return x.Sfixed32
	}
	return 0
}

func (x *Types) GetSfixed64() int64 {
	if x != nil {
		return x.Sfixed64
	}
	return 0
}

func (x *Types) GetFloat() float32 {
	if x != nil {
		return x.Float
	}
	return 0
}

func (x *Types) GetDouble() float64 {
	if x != nil {
		return x.Double
	}
	return 0
}

func (x *Types) GetBool() bool {
	if x != nil {
		return x.Bool
	}
	return false
}

func (x *Types) GetString_() string {
	if x != nil {
		return x.String_
	}
	return ""
}

func (x *Types) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *Types) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Types) GetLanguage() Language {
	if x != nil {
		return x.Language
	}
	return Language_UNKNOWN
}

func (x *Types) GetInt32S() []int32 {
	if x != nil {
		return x.Int32S
	}
	return nil
}

func (x *Types) GetInt64S() []int64 {
	if x != nil {
		return x.Int64S
	}
	return nil
}

func (x *Types) GetUint32S() []uint32 {
	if x != nil {
		return x.Uint32S
	}
	return nil
}

func (x *Types) GetUint64S() []uint64 {
	if x != nil {
		return x.Uint64S
	}
	return nil
}

func (x *Types) GetSint32S() []int32 {
	if x != nil {
		return x.Sint32S
	}
	return nil
}

func (x *Types) GetSint64S() []int64 {
	if x != nil {
		return x.Sint64S
	}
	return nil
}

func (x *Types) GetFixed32S() []uint32 {
	if x != nil {
		return x.Fixed32S
	}
	return nil
}

func (x *Types) GetFixed64S() []uint64 {
	if x != nil {
		return x.Fixed64S
	}
	return nil
}

func (x *Types) GetSfixed32S() []int32 {
	if x != nil {
		return x.Sfixed32S
	}
	return nil
}

func (x *Types) GetSfixed64S() []int64 {
	if x != nil {
		return x.Sfixed64S
	}
	return nil
}

func (x *Types) GetFloats() []float32 {
	if x != nil {
		return x.Floats
	}
	return nil
}

func (x *Types) GetDoubles() []float64 {
	if x != nil {
		return x.Doubles
	}
	return nil
}

func (x *Types) GetBools() []bool {
	if x != nil {
		return x.Bools
	}
	return nil
}

func (x *Types) GetStrings() []string {
	if x != nil {
		return x.Strings
	}
	return nil
}

func (x *Types) GetBytess() [][]byte {
	if x != nil {
		return x.Bytess
	}
	return nil
}

func (x *Types) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *Types) GetLanguages() []Language {
	if x != nil {
		return x.Languages
	}
	return nil
}

func (x *Types) GetValues() map[string]string {
	if x != nil {
		return x.Values
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Int32 int32 `protobuf:"varint,1,opt,name=int32,proto3" json:"int32,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_two_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_msg_two_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_msg_two_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetInt32() int32 {
	if x != nil {
		return x.Int32
	}
	return 0
}

type Types2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Int32     int32             `protobuf:"varint,1,opt,name=int32,proto3" json:"int32,omitempty"`
	Int64     int64             `protobuf:"varint,2,opt,name=int64,proto3" json:"int64,omitempty"`
	Uint32    uint32            `protobuf:"varint,3,opt,name=uint32,proto3" json:"uint32,omitempty"`
	Uint64    uint64            `protobuf:"varint,4,opt,name=uint64,proto3" json:"uint64,omitempty"`
	Sint32    int32             `protobuf:"zigzag32,5,opt,name=sint32,proto3" json:"sint32,omitempty"`
	Sint64    int64             `protobuf:"zigzag64,6,opt,name=sint64,proto3" json:"sint64,omitempty"`
	Fixed32   uint32            `protobuf:"fixed32,7,opt,name=fixed32,proto3" json:"fixed32,omitempty"`
	Fixed64   uint64            `protobuf:"fixed64,8,opt,name=fixed64,proto3" json:"fixed64,omitempty"`
	Sfixed32  int32             `protobuf:"fixed32,9,opt,name=sfixed32,proto3" json:"sfixed32,omitempty"`
	Sfixed64  int64             `protobuf:"fixed64,10,opt,name=sfixed64,proto3" json:"sfixed64,omitempty"`
	Float     float32           `protobuf:"fixed32,11,opt,name=float,proto3" json:"float,omitempty"`
	Double    float64           `protobuf:"fixed64,12,opt,name=double,proto3" json:"double,omitempty"`
	Bool      bool              `protobuf:"varint,13,opt,name=bool,proto3" json:"bool,omitempty"`
	String_   string            `protobuf:"bytes,14,opt,name=string,proto3" json:"string,omitempty"`
	Bytes     []byte            `protobuf:"bytes,15,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Message   *Message2         `protobuf:"bytes,16,opt,name=message,proto3" json:"message,omitempty"`
	Language  Language2         `protobuf:"varint,17,opt,name=language,proto3,enum=two.Language2" json:"language,omitempty"`
	Int32S    []int32           `protobuf:"varint,18,rep,packed,name=int32s,proto3" json:"int32s,omitempty"`
	Int64S    []int64           `protobuf:"varint,19,rep,packed,name=int64s,proto3" json:"int64s,omitempty"`
	Uint32S   []uint32          `protobuf:"varint,20,rep,packed,name=uint32s,proto3" json:"uint32s,omitempty"`
	Uint64S   []uint64          `protobuf:"varint,21,rep,packed,name=uint64s,proto3" json:"uint64s,omitempty"`
	Sint32S   []int32           `protobuf:"zigzag32,22,rep,packed,name=sint32s,proto3" json:"sint32s,omitempty"`
	Sint64S   []int64           `protobuf:"zigzag64,23,rep,packed,name=sint64s,proto3" json:"sint64s,omitempty"`
	Fixed32S  []uint32          `protobuf:"fixed32,24,rep,packed,name=fixed32s,proto3" json:"fixed32s,omitempty"`
	Fixed64S  []uint64          `protobuf:"fixed64,25,rep,packed,name=fixed64s,proto3" json:"fixed64s,omitempty"`
	Sfixed32S []int32           `protobuf:"fixed32,26,rep,packed,name=sfixed32s,proto3" json:"sfixed32s,omitempty"`
	Sfixed64S []int64           `protobuf:"fixed64,27,rep,packed,name=sfixed64s,proto3" json:"sfixed64s,omitempty"`
	Floats    []float32         `protobuf:"fixed32,28,rep,packed,name=floats,proto3" json:"floats,omitempty"`
	Doubles   []float64         `protobuf:"fixed64,29,rep,packed,name=doubles,proto3" json:"doubles,omitempty"`
	Bools     []bool            `protobuf:"varint,30,rep,packed,name=bools,proto3" json:"bools,omitempty"`
	Strings   []string          `protobuf:"bytes,31,rep,name=strings,proto3" json:"strings,omitempty"`
	Bytess    [][]byte          `protobuf:"bytes,32,rep,name=bytess,proto3" json:"bytess,omitempty"`
	Messages  []*Message2       `protobuf:"bytes,33,rep,name=messages,proto3" json:"messages,omitempty"`
	Languages []Language2       `protobuf:"varint,34,rep,packed,name=languages,proto3,enum=two.Language2" json:"languages,omitempty"`
	Values    map[string]string `protobuf:"bytes,35,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Types2) Reset() {
	*x = Types2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_two_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Types2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Types2) ProtoMessage() {}

func (x *Types2) ProtoReflect() protoreflect.Message {
	mi := &file_msg_two_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Types2.ProtoReflect.Descriptor instead.
func (*Types2) Descriptor() ([]byte, []int) {
	return file_msg_two_proto_rawDescGZIP(), []int{3}
}

func (x *Types2) GetInt32() int32 {
	if x != nil {
		return x.Int32
	}
	return 0
}

func (x *Types2) GetInt64() int64 {
	if x != nil {
		return x.Int64
	}
	return 0
}

func (x *Types2) GetUint32() uint32 {
	if x != nil {
		return x.Uint32
	}
	return 0
}

func (x *Types2) GetUint64() uint64 {
	if x != nil {
		return x.Uint64
	}
	return 0
}

func (x *Types2) GetSint32() int32 {
	if x != nil {
		return x.Sint32
	}
	return 0
}

func (x *Types2) GetSint64() int64 {
	if x != nil {
		return x.Sint64
	}
	return 0
}

func (x *Types2) GetFixed32() uint32 {
	if x != nil {
		return x.Fixed32
	}
	return 0
}

func (x *Types2) GetFixed64() uint64 {
	if x != nil {
		return x.Fixed64
	}
	return 0
}

func (x *Types2) GetSfixed32() int32 {
	if x != nil {
		return x.Sfixed32
	}
	return 0
}

func (x *Types2) GetSfixed64() int64 {
	if x != nil {
		return x.Sfixed64
	}
	return 0
}

func (x *Types2) GetFloat() float32 {
	if x != nil {
		return x.Float
	}
	return 0
}

func (x *Types2) GetDouble() float64 {
	if x != nil {
		return x.Double
	}
	return 0
}

func (x *Types2) GetBool() bool {
	if x != nil {
		return x.Bool
	}
	return false
}

func (x *Types2) GetString_() string {
	if x != nil {
		return x.String_
	}
	return ""
}

func (x *Types2) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *Types2) GetMessage() *Message2 {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Types2) GetLanguage() Language2 {
	if x != nil {
		return x.Language
	}
	return Language2_UNKNOWN2
}

func (x *Types2) GetInt32S() []int32 {
	if x != nil {
		return x.Int32S
	}
	return nil
}

func (x *Types2) GetInt64S() []int64 {
	if x != nil {
		return x.Int64S
	}
	return nil
}

func (x *Types2) GetUint32S() []uint32 {
	if x != nil {
		return x.Uint32S
	}
	return nil
}

func (x *Types2) GetUint64S() []uint64 {
	if x != nil {
		return x.Uint64S
	}
	return nil
}

func (x *Types2) GetSint32S() []int32 {
	if x != nil {
		return x.Sint32S
	}
	return nil
}

func (x *Types2) GetSint64S() []int64 {
	if x != nil {
		return x.Sint64S
	}
	return nil
}

func (x *Types2) GetFixed32S() []uint32 {
	if x != nil {
		return x.Fixed32S
	}
	return nil
}

func (x *Types2) GetFixed64S() []uint64 {
	if x != nil {
		return x.Fixed64S
	}
	return nil
}

func (x *Types2) GetSfixed32S() []int32 {
	if x != nil {
		return x.Sfixed32S
	}
	return nil
}

func (x *Types2) GetSfixed64S() []int64 {
	if x != nil {
		return x.Sfixed64S
	}
	return nil
}

func (x *Types2) GetFloats() []float32 {
	if x != nil {
		return x.Floats
	}
	return nil
}

func (x *Types2) GetDoubles() []float64 {
	if x != nil {
		return x.Doubles
	}
	return nil
}

func (x *Types2) GetBools() []bool {
	if x != nil {
		return x.Bools
	}
	return nil
}

func (x *Types2) GetStrings() []string {
	if x != nil {
		return x.Strings
	}
	return nil
}

func (x *Types2) GetBytess() [][]byte {
	if x != nil {
		return x.Bytess
	}
	return nil
}

func (x *Types2) GetMessages() []*Message2 {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *Types2) GetLanguages() []Language2 {
	if x != nil {
		return x.Languages
	}
	return nil
}

func (x *Types2) GetValues() map[string]string {
	if x != nil {
		return x.Values
	}
	return nil
}

type Message2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Int32 int32 `protobuf:"varint,1,opt,name=int32,proto3" json:"int32,omitempty"`
}

func (x *Message2) Reset() {
	*x = Message2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_two_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message2) ProtoMessage() {}

func (x *Message2) ProtoReflect() protoreflect.Message {
	mi := &file_msg_two_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message2.ProtoReflect.Descriptor instead.
func (*Message2) Descriptor() ([]byte, []int) {
	return file_msg_two_proto_rawDescGZIP(), []int{4}
}

func (x *Message2) GetInt32() int32 {
	if x != nil {
		return x.Int32
	}
	return 0
}

var File_msg_two_proto protoreflect.FileDescriptor

var file_msg_two_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x73, 0x67, 0x2d, 0x74, 0x77, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x74, 0x77, 0x6f, 0x22, 0x05, 0x0a, 0x03, 0x4e, 0x6f, 0x70, 0x22, 0x8a, 0x08, 0x0a, 0x05,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x11, 0x52, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x18, 0x06, 0x20, 0x01, 0x28, 0x12, 0x52, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x07, 0x52, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x08, 0x20, 0x01, 0x28, 0x06, 0x52, 0x07, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33,
	0x32, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33,
	0x32, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x10, 0x52, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x66, 0x6c,
	0x6f, 0x61, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x6f, 0x6f, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x26, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x74, 0x77, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x74, 0x77, 0x6f, 0x2e, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x06, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x73, 0x18, 0x13, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x07, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x04, 0x52, 0x07, 0x75, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x73, 0x18,
	0x16, 0x20, 0x03, 0x28, 0x11, 0x52, 0x07, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x73, 0x18, 0x17, 0x20, 0x03, 0x28, 0x12, 0x52,
	0x07, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x33, 0x32, 0x73, 0x18, 0x18, 0x20, 0x03, 0x28, 0x07, 0x52, 0x08, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x33, 0x32, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x73,
	0x18, 0x19, 0x20, 0x03, 0x28, 0x06, 0x52, 0x08, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x73, 0x18, 0x1a, 0x20,
	0x03, 0x28, 0x0f, 0x52, 0x09, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x73, 0x18, 0x1b, 0x20, 0x03, 0x28,
	0x10, 0x52, 0x09, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x66, 0x6c, 0x6f, 0x61, 0x74, 0x73, 0x18, 0x1c, 0x20, 0x03, 0x28, 0x02, 0x52, 0x06, 0x66, 0x6c,
	0x6f, 0x61, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73, 0x18,
	0x1d, 0x20, 0x03, 0x28, 0x01, 0x52, 0x07, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x1e, 0x20, 0x03, 0x28, 0x08, 0x52, 0x05, 0x62,
	0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x1f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x62, 0x79, 0x74, 0x65, 0x73, 0x73, 0x18, 0x20, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x73, 0x12, 0x28, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x21, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x77, 0x6f, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x2b, 0x0a, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x18, 0x22, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x74, 0x77, 0x6f, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x52, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x12, 0x2e, 0x0a,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x23, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x74, 0x77, 0x6f, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x39, 0x0a,
	0x0b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1f, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x22, 0x90, 0x08, 0x0a, 0x06, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x11,
	0x52, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x18, 0x06, 0x20, 0x01, 0x28, 0x12, 0x52, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x07, 0x52, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x08, 0x20, 0x01, 0x28, 0x06, 0x52, 0x07, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x36, 0x34, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x10, 0x52, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x14, 0x0a, 0x05,
	0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f,
	0x6f, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x74, 0x77, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x74, 0x77, 0x6f, 0x2e, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x32, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x73, 0x18, 0x13, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x73, 0x18, 0x14, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x07, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x75,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x04, 0x52, 0x07, 0x75, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x73,
	0x18, 0x16, 0x20, 0x03, 0x28, 0x11, 0x52, 0x07, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x73, 0x18, 0x17, 0x20, 0x03, 0x28, 0x12,
	0x52, 0x07, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x33, 0x32, 0x73, 0x18, 0x18, 0x20, 0x03, 0x28, 0x07, 0x52, 0x08, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x33, 0x32, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34,
	0x73, 0x18, 0x19, 0x20, 0x03, 0x28, 0x06, 0x52, 0x08, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x73, 0x18, 0x1a,
	0x20, 0x03, 0x28, 0x0f, 0x52, 0x09, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x73, 0x18, 0x1b, 0x20, 0x03,
	0x28, 0x10, 0x52, 0x09, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x73, 0x18, 0x1c, 0x20, 0x03, 0x28, 0x02, 0x52, 0x06, 0x66,
	0x6c, 0x6f, 0x61, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73,
	0x18, 0x1d, 0x20, 0x03, 0x28, 0x01, 0x52, 0x07, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x1e, 0x20, 0x03, 0x28, 0x08, 0x52, 0x05,
	0x62, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x1f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x79, 0x74, 0x65, 0x73, 0x73, 0x18, 0x20, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x06, 0x62, 0x79, 0x74, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x21, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x77, 0x6f, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x2c, 0x0a, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x22, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x74, 0x77, 0x6f, 0x2e, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x32, 0x52, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x2f, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x23, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x74, 0x77, 0x6f, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x32, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x20, 0x0a, 0x08,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x2a, 0x49,
	0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x47, 0x4c, 0x49,
	0x53, 0x48, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x50, 0x41, 0x4e, 0x49, 0x53, 0x48, 0x10,
	0x03, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x52, 0x45, 0x4e, 0x43, 0x48, 0x10, 0x04, 0x12, 0x0a, 0x0a,
	0x06, 0x47, 0x45, 0x52, 0x4d, 0x41, 0x4e, 0x10, 0x05, 0x2a, 0x4f, 0x0a, 0x09, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x32, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x32, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x4e, 0x47, 0x4c, 0x49, 0x53, 0x48, 0x32,
	0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x50, 0x41, 0x4e, 0x49, 0x53, 0x48, 0x32, 0x10, 0x03,
	0x12, 0x0b, 0x0a, 0x07, 0x46, 0x52, 0x45, 0x4e, 0x43, 0x48, 0x32, 0x10, 0x04, 0x12, 0x0b, 0x0a,
	0x07, 0x47, 0x45, 0x52, 0x4d, 0x41, 0x4e, 0x32, 0x10, 0x05, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_msg_two_proto_rawDescOnce sync.Once
	file_msg_two_proto_rawDescData = file_msg_two_proto_rawDesc
)

func file_msg_two_proto_rawDescGZIP() []byte {
	file_msg_two_proto_rawDescOnce.Do(func() {
		file_msg_two_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_two_proto_rawDescData)
	})
	return file_msg_two_proto_rawDescData
}

var file_msg_two_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_msg_two_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_msg_two_proto_goTypes = []interface{}{
	(Language)(0),    // 0: two.Language
	(Language2)(0),   // 1: two.Language2
	(*Nop)(nil),      // 2: two.Nop
	(*Types)(nil),    // 3: two.Types
	(*Message)(nil),  // 4: two.Message
	(*Types2)(nil),   // 5: two.Types2
	(*Message2)(nil), // 6: two.Message2
	nil,              // 7: two.Types.ValuesEntry
	nil,              // 8: two.Types2.ValuesEntry
}
var file_msg_two_proto_depIdxs = []int32{
	4,  // 0: two.Types.message:type_name -> two.Message
	0,  // 1: two.Types.language:type_name -> two.Language
	4,  // 2: two.Types.messages:type_name -> two.Message
	0,  // 3: two.Types.languages:type_name -> two.Language
	7,  // 4: two.Types.values:type_name -> two.Types.ValuesEntry
	6,  // 5: two.Types2.message:type_name -> two.Message2
	1,  // 6: two.Types2.language:type_name -> two.Language2
	6,  // 7: two.Types2.messages:type_name -> two.Message2
	1,  // 8: two.Types2.languages:type_name -> two.Language2
	8,  // 9: two.Types2.values:type_name -> two.Types2.ValuesEntry
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_msg_two_proto_init() }
func file_msg_two_proto_init() {
	if File_msg_two_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_two_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nop); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msg_two_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Types); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msg_two_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msg_two_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Types2); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msg_two_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message2); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msg_two_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_two_proto_goTypes,
		DependencyIndexes: file_msg_two_proto_depIdxs,
		EnumInfos:         file_msg_two_proto_enumTypes,
		MessageInfos:      file_msg_two_proto_msgTypes,
	}.Build()
	File_msg_two_proto = out.File
	file_msg_two_proto_rawDesc = nil
	file_msg_two_proto_goTypes = nil
	file_msg_two_proto_depIdxs = nil
}
