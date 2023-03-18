// Copyright 2020 Google LLC
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/actions/sdk/v2/interactionmodel/prompt/content/static_collection_prompt.proto

package prompt

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A card for presenting a collection of options to select from.
type StaticCollectionPrompt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. Title of the collection.
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// Optional. Subtitle of the collection.
	Subtitle string `protobuf:"bytes,2,opt,name=subtitle,proto3" json:"subtitle,omitempty"`
	// Required. Collection items.
	Items []*StaticCollectionPrompt_CollectionItem `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	// Optional. Type of image display option.
	ImageFill StaticImagePrompt_ImageFill `protobuf:"varint,4,opt,name=image_fill,json=imageFill,proto3,enum=google.actions.sdk.v2.interactionmodel.prompt.StaticImagePrompt_ImageFill" json:"image_fill,omitempty"`
}

func (x *StaticCollectionPrompt) Reset() {
	*x = StaticCollectionPrompt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticCollectionPrompt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticCollectionPrompt) ProtoMessage() {}

func (x *StaticCollectionPrompt) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticCollectionPrompt.ProtoReflect.Descriptor instead.
func (*StaticCollectionPrompt) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_rawDescGZIP(), []int{0}
}

func (x *StaticCollectionPrompt) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *StaticCollectionPrompt) GetSubtitle() string {
	if x != nil {
		return x.Subtitle
	}
	return ""
}

func (x *StaticCollectionPrompt) GetItems() []*StaticCollectionPrompt_CollectionItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *StaticCollectionPrompt) GetImageFill() StaticImagePrompt_ImageFill {
	if x != nil {
		return x.ImageFill
	}
	return StaticImagePrompt_UNSPECIFIED
}

// An item in the collection.
type StaticCollectionPrompt_CollectionItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The NLU key that matches the entry key name in the associated
	// Type. When item tapped, this key will be posted back as a select option
	// parameter.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Required. Title of the item. When tapped, this text will be
	// posted back to the conversation verbatim as if the user had typed it.
	// Each title must be unique among the set of items.
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// Optional. Body text of the item.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Optional. Item image.
	Image *StaticImagePrompt `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *StaticCollectionPrompt_CollectionItem) Reset() {
	*x = StaticCollectionPrompt_CollectionItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticCollectionPrompt_CollectionItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticCollectionPrompt_CollectionItem) ProtoMessage() {}

func (x *StaticCollectionPrompt_CollectionItem) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticCollectionPrompt_CollectionItem.ProtoReflect.Descriptor instead.
func (*StaticCollectionPrompt_CollectionItem) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_rawDescGZIP(), []int{0, 0}
}

func (x *StaticCollectionPrompt_CollectionItem) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *StaticCollectionPrompt_CollectionItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *StaticCollectionPrompt_CollectionItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *StaticCollectionPrompt_CollectionItem) GetImage() *StaticImagePrompt {
	if x != nil {
		return x.Image
	}
	return nil
}

var File_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto protoreflect.FileDescriptor

var file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_rawDesc = []byte{
	0x0a, 0x54, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x1a, 0x4f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x03, 0x0a, 0x16, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6d,
	0x70, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1f, 0x0a,
	0x08, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x6f,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x54, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x74, 0x65, 0x6d, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x6e, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x4a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x69, 0x6c, 0x6c, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x69, 0x6c, 0x6c, 0x1a,
	0xc6, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x15, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5b, 0x0a, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e,
	0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x42, 0x03, 0xe0, 0x41,
	0x01, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0xa7, 0x01, 0x0a, 0x31, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x42, 0x1b,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b,
	0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x3b, 0x70, 0x72, 0x6f, 0x6d,
	0x70, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_rawDescOnce sync.Once
	file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_rawDescData = file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_rawDesc
)

func file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_rawDescGZIP() []byte {
	file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_rawDescOnce.Do(func() {
		file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_rawDescData)
	})
	return file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_rawDescData
}

var file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_goTypes = []interface{}{
	(*StaticCollectionPrompt)(nil),                // 0: google.actions.sdk.v2.interactionmodel.prompt.StaticCollectionPrompt
	(*StaticCollectionPrompt_CollectionItem)(nil), // 1: google.actions.sdk.v2.interactionmodel.prompt.StaticCollectionPrompt.CollectionItem
	(StaticImagePrompt_ImageFill)(0),              // 2: google.actions.sdk.v2.interactionmodel.prompt.StaticImagePrompt.ImageFill
	(*StaticImagePrompt)(nil),                     // 3: google.actions.sdk.v2.interactionmodel.prompt.StaticImagePrompt
}
var file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_depIdxs = []int32{
	1, // 0: google.actions.sdk.v2.interactionmodel.prompt.StaticCollectionPrompt.items:type_name -> google.actions.sdk.v2.interactionmodel.prompt.StaticCollectionPrompt.CollectionItem
	2, // 1: google.actions.sdk.v2.interactionmodel.prompt.StaticCollectionPrompt.image_fill:type_name -> google.actions.sdk.v2.interactionmodel.prompt.StaticImagePrompt.ImageFill
	3, // 2: google.actions.sdk.v2.interactionmodel.prompt.StaticCollectionPrompt.CollectionItem.image:type_name -> google.actions.sdk.v2.interactionmodel.prompt.StaticImagePrompt
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_init()
}
func file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_init() {
	if File_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto != nil {
		return
	}
	file_google_actions_sdk_v2_interactionmodel_prompt_content_static_image_prompt_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticCollectionPrompt); i {
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
		file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticCollectionPrompt_CollectionItem); i {
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
			RawDescriptor: file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_goTypes,
		DependencyIndexes: file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_depIdxs,
		MessageInfos:      file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_msgTypes,
	}.Build()
	File_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto = out.File
	file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_rawDesc = nil
	file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_goTypes = nil
	file_google_actions_sdk_v2_interactionmodel_prompt_content_static_collection_prompt_proto_depIdxs = nil
}
