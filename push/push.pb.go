// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: push/push.proto

package push

import (
	sdkws "github.com/SupersStone/serverpro/sdkws"
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

type PushMsgReq struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	MsgData        *sdkws.MsgData         `protobuf:"bytes,1,opt,name=msgData,proto3" json:"msgData"`
	ConversationID string                 `protobuf:"bytes,2,opt,name=conversationID,proto3" json:"conversationID"`
	UserIDs        []string               `protobuf:"bytes,3,rep,name=userIDs,proto3" json:"userIDs"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *PushMsgReq) Reset() {
	*x = PushMsgReq{}
	mi := &file_push_push_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PushMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsgReq) ProtoMessage() {}

func (x *PushMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_push_push_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsgReq.ProtoReflect.Descriptor instead.
func (*PushMsgReq) Descriptor() ([]byte, []int) {
	return file_push_push_proto_rawDescGZIP(), []int{0}
}

func (x *PushMsgReq) GetMsgData() *sdkws.MsgData {
	if x != nil {
		return x.MsgData
	}
	return nil
}

func (x *PushMsgReq) GetConversationID() string {
	if x != nil {
		return x.ConversationID
	}
	return ""
}

func (x *PushMsgReq) GetUserIDs() []string {
	if x != nil {
		return x.UserIDs
	}
	return nil
}

type PushMsgResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PushMsgResp) Reset() {
	*x = PushMsgResp{}
	mi := &file_push_push_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PushMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsgResp) ProtoMessage() {}

func (x *PushMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_push_push_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsgResp.ProtoReflect.Descriptor instead.
func (*PushMsgResp) Descriptor() ([]byte, []int) {
	return file_push_push_proto_rawDescGZIP(), []int{1}
}

type DelUserPushTokenReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserID        string                 `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID"`
	PlatformID    int32                  `protobuf:"varint,2,opt,name=platformID,proto3" json:"platformID"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DelUserPushTokenReq) Reset() {
	*x = DelUserPushTokenReq{}
	mi := &file_push_push_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelUserPushTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelUserPushTokenReq) ProtoMessage() {}

func (x *DelUserPushTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_push_push_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelUserPushTokenReq.ProtoReflect.Descriptor instead.
func (*DelUserPushTokenReq) Descriptor() ([]byte, []int) {
	return file_push_push_proto_rawDescGZIP(), []int{2}
}

func (x *DelUserPushTokenReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *DelUserPushTokenReq) GetPlatformID() int32 {
	if x != nil {
		return x.PlatformID
	}
	return 0
}

type DelUserPushTokenResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DelUserPushTokenResp) Reset() {
	*x = DelUserPushTokenResp{}
	mi := &file_push_push_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelUserPushTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelUserPushTokenResp) ProtoMessage() {}

func (x *DelUserPushTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_push_push_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelUserPushTokenResp.ProtoReflect.Descriptor instead.
func (*DelUserPushTokenResp) Descriptor() ([]byte, []int) {
	return file_push_push_proto_rawDescGZIP(), []int{3}
}

var File_push_push_proto protoreflect.FileDescriptor

var file_push_push_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x75, 0x73, 0x68, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x1a, 0x11,
	0x73, 0x64, 0x6b, 0x77, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x7f, 0x0a, 0x0a, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x12,
	0x2f, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x73, 0x64, 0x6b, 0x77, 0x73, 0x2e,
	0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x73, 0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x4d, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x50, 0x75, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44,
	0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x50, 0x75, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x32, 0xa7, 0x01, 0x0a, 0x0e, 0x50, 0x75, 0x73,
	0x68, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x50,
	0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x12, 0x17, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e,
	0x70, 0x75, 0x73, 0x68, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x1a,
	0x18, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x50, 0x75,
	0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x57, 0x0a, 0x10, 0x44, 0x65, 0x6c,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a,
	0x21, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x44, 0x65,
	0x6c, 0x55, 0x73, 0x65, 0x72, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_push_push_proto_rawDescOnce sync.Once
	file_push_push_proto_rawDescData = file_push_push_proto_rawDesc
)

func file_push_push_proto_rawDescGZIP() []byte {
	file_push_push_proto_rawDescOnce.Do(func() {
		file_push_push_proto_rawDescData = protoimpl.X.CompressGZIP(file_push_push_proto_rawDescData)
	})
	return file_push_push_proto_rawDescData
}

var file_push_push_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_push_push_proto_goTypes = []any{
	(*PushMsgReq)(nil),           // 0: openim.push.PushMsgReq
	(*PushMsgResp)(nil),          // 1: openim.push.PushMsgResp
	(*DelUserPushTokenReq)(nil),  // 2: openim.push.DelUserPushTokenReq
	(*DelUserPushTokenResp)(nil), // 3: openim.push.DelUserPushTokenResp
	(*sdkws.MsgData)(nil),        // 4: openim.sdkws.MsgData
}
var file_push_push_proto_depIdxs = []int32{
	4, // 0: openim.push.PushMsgReq.msgData:type_name -> openim.sdkws.MsgData
	0, // 1: openim.push.PushMsgService.PushMsg:input_type -> openim.push.PushMsgReq
	2, // 2: openim.push.PushMsgService.DelUserPushToken:input_type -> openim.push.DelUserPushTokenReq
	1, // 3: openim.push.PushMsgService.PushMsg:output_type -> openim.push.PushMsgResp
	3, // 4: openim.push.PushMsgService.DelUserPushToken:output_type -> openim.push.DelUserPushTokenResp
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_push_push_proto_init() }
func file_push_push_proto_init() {
	if File_push_push_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_push_push_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_push_push_proto_goTypes,
		DependencyIndexes: file_push_push_proto_depIdxs,
		MessageInfos:      file_push_push_proto_msgTypes,
	}.Build()
	File_push_push_proto = out.File
	file_push_push_proto_rawDesc = nil
	file_push_push_proto_goTypes = nil
	file_push_push_proto_depIdxs = nil
}
