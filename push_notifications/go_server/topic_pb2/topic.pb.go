// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: proto/topic.proto

package topic_pb2

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

type FactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`   // how many meow facts will be returned per response
	Second int32 `protobuf:"varint,2,opt,name=second,proto3" json:"second,omitempty"` // time delay in second
	Limit  int32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`   // maximum response, will be stop after reached
}

func (x *FactRequest) Reset() {
	*x = FactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_topic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FactRequest) ProtoMessage() {}

func (x *FactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_topic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FactRequest.ProtoReflect.Descriptor instead.
func (*FactRequest) Descriptor() ([]byte, []int) {
	return file_proto_topic_proto_rawDescGZIP(), []int{0}
}

func (x *FactRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *FactRequest) GetSecond() int32 {
	if x != nil {
		return x.Second
	}
	return 0
}

func (x *FactRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type FactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Facts []string `protobuf:"bytes,1,rep,name=facts,proto3" json:"facts,omitempty"` // list of meow facts
}

func (x *FactResponse) Reset() {
	*x = FactResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_topic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FactResponse) ProtoMessage() {}

func (x *FactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_topic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FactResponse.ProtoReflect.Descriptor instead.
func (*FactResponse) Descriptor() ([]byte, []int) {
	return file_proto_topic_proto_rawDescGZIP(), []int{1}
}

func (x *FactResponse) GetFacts() []string {
	if x != nil {
		return x.Facts
	}
	return nil
}

type ColorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   []string `protobuf:"bytes,1,rep,name=name,proto3" json:"name,omitempty"`      // list of color name
	Second int32    `protobuf:"varint,2,opt,name=second,proto3" json:"second,omitempty"` // time delay in second
}

func (x *ColorRequest) Reset() {
	*x = ColorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_topic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ColorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColorRequest) ProtoMessage() {}

func (x *ColorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_topic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColorRequest.ProtoReflect.Descriptor instead.
func (*ColorRequest) Descriptor() ([]byte, []int) {
	return file_proto_topic_proto_rawDescGZIP(), []int{2}
}

func (x *ColorRequest) GetName() []string {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *ColorRequest) GetSecond() int32 {
	if x != nil {
		return x.Second
	}
	return 0
}

type ColorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hex string `protobuf:"bytes,1,opt,name=hex,proto3" json:"hex,omitempty"` // Hex Value
	Rgb string `protobuf:"bytes,2,opt,name=rgb,proto3" json:"rgb,omitempty"` // RGB Value
	Hsl string `protobuf:"bytes,3,opt,name=hsl,proto3" json:"hsl,omitempty"` // HSL Value
}

func (x *ColorResponse) Reset() {
	*x = ColorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_topic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ColorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColorResponse) ProtoMessage() {}

func (x *ColorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_topic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColorResponse.ProtoReflect.Descriptor instead.
func (*ColorResponse) Descriptor() ([]byte, []int) {
	return file_proto_topic_proto_rawDescGZIP(), []int{3}
}

func (x *ColorResponse) GetHex() string {
	if x != nil {
		return x.Hex
	}
	return ""
}

func (x *ColorResponse) GetRgb() string {
	if x != nil {
		return x.Rgb
	}
	return ""
}

func (x *ColorResponse) GetHsl() string {
	if x != nil {
		return x.Hsl
	}
	return ""
}

var File_proto_topic_proto protoreflect.FileDescriptor

var file_proto_topic_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x51, 0x0a, 0x0b, 0x46, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x24, 0x0a, 0x0c, 0x46, 0x61, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x61, 0x63, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x66, 0x61, 0x63, 0x74, 0x73, 0x22, 0x3a, 0x0a,
	0x0c, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x22, 0x45, 0x0a, 0x0d, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x68, 0x65,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x68, 0x65, 0x78, 0x12, 0x10, 0x0a, 0x03,
	0x72, 0x67, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x67, 0x62, 0x12, 0x10,
	0x0a, 0x03, 0x68, 0x73, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x68, 0x73, 0x6c,
	0x32, 0xb6, 0x01, 0x0a, 0x0e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x50, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x6f,
	0x77, 0x46, 0x61, 0x63, 0x74, 0x73, 0x12, 0x1c, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x52, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43,
	0x6f, 0x6c, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x15, 0x5a, 0x13, 0x67, 0x6f, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x70, 0x62, 0x32,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_topic_proto_rawDescOnce sync.Once
	file_proto_topic_proto_rawDescData = file_proto_topic_proto_rawDesc
)

func file_proto_topic_proto_rawDescGZIP() []byte {
	file_proto_topic_proto_rawDescOnce.Do(func() {
		file_proto_topic_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_topic_proto_rawDescData)
	})
	return file_proto_topic_proto_rawDescData
}

var file_proto_topic_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_topic_proto_goTypes = []interface{}{
	(*FactRequest)(nil),   // 0: topic_selection.FactRequest
	(*FactResponse)(nil),  // 1: topic_selection.FactResponse
	(*ColorRequest)(nil),  // 2: topic_selection.ColorRequest
	(*ColorResponse)(nil), // 3: topic_selection.ColorResponse
}
var file_proto_topic_proto_depIdxs = []int32{
	0, // 0: topic_selection.TopicSelection.StreamMeowFacts:input_type -> topic_selection.FactRequest
	2, // 1: topic_selection.TopicSelection.StreamColorInfo:input_type -> topic_selection.ColorRequest
	1, // 2: topic_selection.TopicSelection.StreamMeowFacts:output_type -> topic_selection.FactResponse
	3, // 3: topic_selection.TopicSelection.StreamColorInfo:output_type -> topic_selection.ColorResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_topic_proto_init() }
func file_proto_topic_proto_init() {
	if File_proto_topic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_topic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FactRequest); i {
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
		file_proto_topic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FactResponse); i {
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
		file_proto_topic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ColorRequest); i {
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
		file_proto_topic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ColorResponse); i {
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
			RawDescriptor: file_proto_topic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_topic_proto_goTypes,
		DependencyIndexes: file_proto_topic_proto_depIdxs,
		MessageInfos:      file_proto_topic_proto_msgTypes,
	}.Build()
	File_proto_topic_proto = out.File
	file_proto_topic_proto_rawDesc = nil
	file_proto_topic_proto_goTypes = nil
	file_proto_topic_proto_depIdxs = nil
}
