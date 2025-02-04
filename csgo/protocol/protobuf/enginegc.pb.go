// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v3.12.4
// source: engine_gcmessages.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CEngineGotvSyncPacket struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	MatchId          *uint64                `protobuf:"varint,1,opt,name=match_id,json=matchId" json:"match_id,omitempty"`
	InstanceId       *uint32                `protobuf:"varint,2,opt,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
	Signupfragment   *uint32                `protobuf:"varint,3,opt,name=signupfragment" json:"signupfragment,omitempty"`
	Currentfragment  *uint32                `protobuf:"varint,4,opt,name=currentfragment" json:"currentfragment,omitempty"`
	Tickrate         *float32               `protobuf:"fixed32,5,opt,name=tickrate" json:"tickrate,omitempty"`
	Tick             *uint32                `protobuf:"varint,6,opt,name=tick" json:"tick,omitempty"`
	Rtdelay          *float32               `protobuf:"fixed32,8,opt,name=rtdelay" json:"rtdelay,omitempty"`
	Rcvage           *float32               `protobuf:"fixed32,9,opt,name=rcvage" json:"rcvage,omitempty"`
	KeyframeInterval *float32               `protobuf:"fixed32,10,opt,name=keyframe_interval,json=keyframeInterval" json:"keyframe_interval,omitempty"`
	Cdndelay         *uint32                `protobuf:"varint,11,opt,name=cdndelay" json:"cdndelay,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *CEngineGotvSyncPacket) Reset() {
	*x = CEngineGotvSyncPacket{}
	mi := &file_engine_gcmessages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CEngineGotvSyncPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CEngineGotvSyncPacket) ProtoMessage() {}

func (x *CEngineGotvSyncPacket) ProtoReflect() protoreflect.Message {
	mi := &file_engine_gcmessages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CEngineGotvSyncPacket.ProtoReflect.Descriptor instead.
func (*CEngineGotvSyncPacket) Descriptor() ([]byte, []int) {
	return file_engine_gcmessages_proto_rawDescGZIP(), []int{0}
}

func (x *CEngineGotvSyncPacket) GetMatchId() uint64 {
	if x != nil && x.MatchId != nil {
		return *x.MatchId
	}
	return 0
}

func (x *CEngineGotvSyncPacket) GetInstanceId() uint32 {
	if x != nil && x.InstanceId != nil {
		return *x.InstanceId
	}
	return 0
}

func (x *CEngineGotvSyncPacket) GetSignupfragment() uint32 {
	if x != nil && x.Signupfragment != nil {
		return *x.Signupfragment
	}
	return 0
}

func (x *CEngineGotvSyncPacket) GetCurrentfragment() uint32 {
	if x != nil && x.Currentfragment != nil {
		return *x.Currentfragment
	}
	return 0
}

func (x *CEngineGotvSyncPacket) GetTickrate() float32 {
	if x != nil && x.Tickrate != nil {
		return *x.Tickrate
	}
	return 0
}

func (x *CEngineGotvSyncPacket) GetTick() uint32 {
	if x != nil && x.Tick != nil {
		return *x.Tick
	}
	return 0
}

func (x *CEngineGotvSyncPacket) GetRtdelay() float32 {
	if x != nil && x.Rtdelay != nil {
		return *x.Rtdelay
	}
	return 0
}

func (x *CEngineGotvSyncPacket) GetRcvage() float32 {
	if x != nil && x.Rcvage != nil {
		return *x.Rcvage
	}
	return 0
}

func (x *CEngineGotvSyncPacket) GetKeyframeInterval() float32 {
	if x != nil && x.KeyframeInterval != nil {
		return *x.KeyframeInterval
	}
	return 0
}

func (x *CEngineGotvSyncPacket) GetCdndelay() uint32 {
	if x != nil && x.Cdndelay != nil {
		return *x.Cdndelay
	}
	return 0
}

var File_engine_gcmessages_proto protoreflect.FileDescriptor

var file_engine_gcmessages_proto_rawDesc = string([]byte{
	0x0a, 0x17, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x67, 0x63, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x02, 0x0a, 0x15,
	0x43, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x47, 0x6f, 0x74, 0x76, 0x53, 0x79, 0x6e, 0x63, 0x50,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x66, 0x72, 0x61, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x75,
	0x70, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x66, 0x72, 0x61, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x72, 0x61, 0x74, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x69, 0x63, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74,
	0x69, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x74, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x72, 0x74, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x63, 0x76, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72,
	0x63, 0x76, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x6b, 0x65, 0x79, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x10, 0x6b, 0x65, 0x79, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x64, 0x6e, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x64, 0x6e, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x42, 0x03,
	0x80, 0x01, 0x00,
})

var (
	file_engine_gcmessages_proto_rawDescOnce sync.Once
	file_engine_gcmessages_proto_rawDescData []byte
)

func file_engine_gcmessages_proto_rawDescGZIP() []byte {
	file_engine_gcmessages_proto_rawDescOnce.Do(func() {
		file_engine_gcmessages_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_engine_gcmessages_proto_rawDesc), len(file_engine_gcmessages_proto_rawDesc)))
	})
	return file_engine_gcmessages_proto_rawDescData
}

var file_engine_gcmessages_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_engine_gcmessages_proto_goTypes = []any{
	(*CEngineGotvSyncPacket)(nil), // 0: CEngineGotvSyncPacket
}
var file_engine_gcmessages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_engine_gcmessages_proto_init() }
func file_engine_gcmessages_proto_init() {
	if File_engine_gcmessages_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_engine_gcmessages_proto_rawDesc), len(file_engine_gcmessages_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_engine_gcmessages_proto_goTypes,
		DependencyIndexes: file_engine_gcmessages_proto_depIdxs,
		MessageInfos:      file_engine_gcmessages_proto_msgTypes,
	}.Build()
	File_engine_gcmessages_proto = out.File
	file_engine_gcmessages_proto_goTypes = nil
	file_engine_gcmessages_proto_depIdxs = nil
}
