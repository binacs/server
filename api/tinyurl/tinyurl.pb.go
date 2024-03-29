// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: tinyurl.proto

package api_tinyurl

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// TinyURLEncode
type TinyURLEncodeResObj struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Turl string `protobuf:"bytes,1,opt,name=turl,proto3" json:"turl,omitempty"`
}

func (x *TinyURLEncodeResObj) Reset() {
	*x = TinyURLEncodeResObj{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tinyurl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TinyURLEncodeResObj) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TinyURLEncodeResObj) ProtoMessage() {}

func (x *TinyURLEncodeResObj) ProtoReflect() protoreflect.Message {
	mi := &file_tinyurl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TinyURLEncodeResObj.ProtoReflect.Descriptor instead.
func (*TinyURLEncodeResObj) Descriptor() ([]byte, []int) {
	return file_tinyurl_proto_rawDescGZIP(), []int{0}
}

func (x *TinyURLEncodeResObj) GetTurl() string {
	if x != nil {
		return x.Turl
	}
	return ""
}

type TinyURLEncodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *TinyURLEncodeReq) Reset() {
	*x = TinyURLEncodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tinyurl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TinyURLEncodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TinyURLEncodeReq) ProtoMessage() {}

func (x *TinyURLEncodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_tinyurl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TinyURLEncodeReq.ProtoReflect.Descriptor instead.
func (*TinyURLEncodeReq) Descriptor() ([]byte, []int) {
	return file_tinyurl_proto_rawDescGZIP(), []int{1}
}

func (x *TinyURLEncodeReq) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type TinyURLEncodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64                `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string               `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *TinyURLEncodeResObj `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TinyURLEncodeResp) Reset() {
	*x = TinyURLEncodeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tinyurl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TinyURLEncodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TinyURLEncodeResp) ProtoMessage() {}

func (x *TinyURLEncodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_tinyurl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TinyURLEncodeResp.ProtoReflect.Descriptor instead.
func (*TinyURLEncodeResp) Descriptor() ([]byte, []int) {
	return file_tinyurl_proto_rawDescGZIP(), []int{2}
}

func (x *TinyURLEncodeResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *TinyURLEncodeResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *TinyURLEncodeResp) GetData() *TinyURLEncodeResObj {
	if x != nil {
		return x.Data
	}
	return nil
}

// TinyURLDecode
type TinyURLDecodeResObj struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *TinyURLDecodeResObj) Reset() {
	*x = TinyURLDecodeResObj{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tinyurl_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TinyURLDecodeResObj) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TinyURLDecodeResObj) ProtoMessage() {}

func (x *TinyURLDecodeResObj) ProtoReflect() protoreflect.Message {
	mi := &file_tinyurl_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TinyURLDecodeResObj.ProtoReflect.Descriptor instead.
func (*TinyURLDecodeResObj) Descriptor() ([]byte, []int) {
	return file_tinyurl_proto_rawDescGZIP(), []int{3}
}

func (x *TinyURLDecodeResObj) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type TinyURLDecodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Turl string `protobuf:"bytes,1,opt,name=turl,proto3" json:"turl,omitempty"`
}

func (x *TinyURLDecodeReq) Reset() {
	*x = TinyURLDecodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tinyurl_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TinyURLDecodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TinyURLDecodeReq) ProtoMessage() {}

func (x *TinyURLDecodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_tinyurl_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TinyURLDecodeReq.ProtoReflect.Descriptor instead.
func (*TinyURLDecodeReq) Descriptor() ([]byte, []int) {
	return file_tinyurl_proto_rawDescGZIP(), []int{4}
}

func (x *TinyURLDecodeReq) GetTurl() string {
	if x != nil {
		return x.Turl
	}
	return ""
}

type TinyURLDecodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64                `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string               `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *TinyURLDecodeResObj `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TinyURLDecodeResp) Reset() {
	*x = TinyURLDecodeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tinyurl_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TinyURLDecodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TinyURLDecodeResp) ProtoMessage() {}

func (x *TinyURLDecodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_tinyurl_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TinyURLDecodeResp.ProtoReflect.Descriptor instead.
func (*TinyURLDecodeResp) Descriptor() ([]byte, []int) {
	return file_tinyurl_proto_rawDescGZIP(), []int{5}
}

func (x *TinyURLDecodeResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *TinyURLDecodeResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *TinyURLDecodeResp) GetData() *TinyURLDecodeResObj {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_tinyurl_proto protoreflect.FileDescriptor

var file_tinyurl_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x69, 0x6e, 0x79, 0x75, 0x72, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a,
	0x13, 0x54, 0x69, 0x6e, 0x79, 0x55, 0x52, 0x4c, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x4f, 0x62, 0x6a, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x75, 0x72, 0x6c, 0x22, 0x24, 0x0a, 0x10, 0x54, 0x69, 0x6e, 0x79,
	0x55, 0x52, 0x4c, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x63,
	0x0a, 0x11, 0x54, 0x69, 0x6e, 0x79, 0x55, 0x52, 0x4c, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x54, 0x69, 0x6e, 0x79, 0x55, 0x52,
	0x4c, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x4f, 0x62, 0x6a, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x27, 0x0a, 0x13, 0x54, 0x69, 0x6e, 0x79, 0x55, 0x52, 0x4c, 0x44, 0x65,
	0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x4f, 0x62, 0x6a, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x26, 0x0a, 0x10,
	0x54, 0x69, 0x6e, 0x79, 0x55, 0x52, 0x4c, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x75, 0x72, 0x6c, 0x22, 0x63, 0x0a, 0x11, 0x54, 0x69, 0x6e, 0x79, 0x55, 0x52, 0x4c, 0x44,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x54, 0x69, 0x6e, 0x79, 0x55, 0x52, 0x4c, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x4f, 0x62, 0x6a, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xb1, 0x01, 0x0a, 0x07, 0x54, 0x69,
	0x6e, 0x79, 0x55, 0x52, 0x4c, 0x12, 0x52, 0x0a, 0x0d, 0x54, 0x69, 0x6e, 0x79, 0x55, 0x52, 0x4c,
	0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x11, 0x2e, 0x54, 0x69, 0x6e, 0x79, 0x55, 0x52, 0x4c,
	0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x54, 0x69, 0x6e, 0x79,
	0x55, 0x52, 0x4c, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x74, 0x69, 0x6e, 0x79, 0x75,
	0x72, 0x6c, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x54, 0x69, 0x6e,
	0x79, 0x55, 0x52, 0x4c, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x11, 0x2e, 0x54, 0x69, 0x6e,
	0x79, 0x55, 0x52, 0x4c, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e,
	0x54, 0x69, 0x6e, 0x79, 0x55, 0x52, 0x4c, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x74,
	0x69, 0x6e, 0x79, 0x75, 0x72, 0x6c, 0x2f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x0f, 0x5a,
	0x0d, 0x2e, 0x3b, 0x61, 0x70, 0x69, 0x5f, 0x74, 0x69, 0x6e, 0x79, 0x75, 0x72, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tinyurl_proto_rawDescOnce sync.Once
	file_tinyurl_proto_rawDescData = file_tinyurl_proto_rawDesc
)

func file_tinyurl_proto_rawDescGZIP() []byte {
	file_tinyurl_proto_rawDescOnce.Do(func() {
		file_tinyurl_proto_rawDescData = protoimpl.X.CompressGZIP(file_tinyurl_proto_rawDescData)
	})
	return file_tinyurl_proto_rawDescData
}

var file_tinyurl_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tinyurl_proto_goTypes = []interface{}{
	(*TinyURLEncodeResObj)(nil), // 0: TinyURLEncodeResObj
	(*TinyURLEncodeReq)(nil),    // 1: TinyURLEncodeReq
	(*TinyURLEncodeResp)(nil),   // 2: TinyURLEncodeResp
	(*TinyURLDecodeResObj)(nil), // 3: TinyURLDecodeResObj
	(*TinyURLDecodeReq)(nil),    // 4: TinyURLDecodeReq
	(*TinyURLDecodeResp)(nil),   // 5: TinyURLDecodeResp
}
var file_tinyurl_proto_depIdxs = []int32{
	0, // 0: TinyURLEncodeResp.data:type_name -> TinyURLEncodeResObj
	3, // 1: TinyURLDecodeResp.data:type_name -> TinyURLDecodeResObj
	1, // 2: TinyURL.TinyURLEncode:input_type -> TinyURLEncodeReq
	4, // 3: TinyURL.TinyURLDecode:input_type -> TinyURLDecodeReq
	2, // 4: TinyURL.TinyURLEncode:output_type -> TinyURLEncodeResp
	5, // 5: TinyURL.TinyURLDecode:output_type -> TinyURLDecodeResp
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tinyurl_proto_init() }
func file_tinyurl_proto_init() {
	if File_tinyurl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tinyurl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TinyURLEncodeResObj); i {
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
		file_tinyurl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TinyURLEncodeReq); i {
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
		file_tinyurl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TinyURLEncodeResp); i {
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
		file_tinyurl_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TinyURLDecodeResObj); i {
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
		file_tinyurl_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TinyURLDecodeReq); i {
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
		file_tinyurl_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TinyURLDecodeResp); i {
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
			RawDescriptor: file_tinyurl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tinyurl_proto_goTypes,
		DependencyIndexes: file_tinyurl_proto_depIdxs,
		MessageInfos:      file_tinyurl_proto_msgTypes,
	}.Build()
	File_tinyurl_proto = out.File
	file_tinyurl_proto_rawDesc = nil
	file_tinyurl_proto_goTypes = nil
	file_tinyurl_proto_depIdxs = nil
}
