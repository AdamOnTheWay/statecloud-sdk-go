// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.25.1
// source: tag.proto

package tag

import (
	_ "github.com/state-cloud/statecloud-sdk-go/service/eci/types/api"
	containergroup "github.com/state-cloud/statecloud-sdk-go/service/eci/types/containergroup"
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

type BindTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId   string                `protobuf:"bytes,1,opt,name=resourceId,proto3" form:"resourceId" json:"resourceId,omitempty"`
	Tags         []*containergroup.Tag `protobuf:"bytes,2,rep,name=tags,proto3" form:"tags" json:"tags,omitempty"`
	ResourceType string                `protobuf:"bytes,3,opt,name=resourceType,proto3" form:"resourceType" json:"resourceType,omitempty"`
}

func (x *BindTagRequest) Reset() {
	*x = BindTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindTagRequest) ProtoMessage() {}

func (x *BindTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindTagRequest.ProtoReflect.Descriptor instead.
func (*BindTagRequest) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{0}
}

func (x *BindTagRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *BindTagRequest) GetTags() []*containergroup.Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *BindTagRequest) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

type BindTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=requestId,proto3" form:"requestId" json:"requestId,omitempty"`
}

func (x *BindTagResponse) Reset() {
	*x = BindTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindTagResponse) ProtoMessage() {}

func (x *BindTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindTagResponse.ProtoReflect.Descriptor instead.
func (*BindTagResponse) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{1}
}

func (x *BindTagResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type ListTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceIds  []string              `protobuf:"bytes,1,rep,name=resourceIds,proto3" form:"resourceIds" json:"resourceIds,omitempty"`
	Tags         []*containergroup.Tag `protobuf:"bytes,2,rep,name=tags,proto3" form:"tags" json:"tags,omitempty"`
	ResourceType string                `protobuf:"bytes,3,opt,name=resourceType,proto3" form:"resourceType" json:"resourceType,omitempty"`
	Limit        int32                 `protobuf:"varint,4,opt,name=limit,proto3" form:"limit" json:"limit,omitempty"`
	NextToken    string                `protobuf:"bytes,5,opt,name=nextToken,proto3" form:"nextToken" json:"nextToken,omitempty"`
}

func (x *ListTagRequest) Reset() {
	*x = ListTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTagRequest) ProtoMessage() {}

func (x *ListTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTagRequest.ProtoReflect.Descriptor instead.
func (*ListTagRequest) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{2}
}

func (x *ListTagRequest) GetResourceIds() []string {
	if x != nil {
		return x.ResourceIds
	}
	return nil
}

func (x *ListTagRequest) GetTags() []*containergroup.Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *ListTagRequest) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *ListTagRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListTagRequest) GetNextToken() string {
	if x != nil {
		return x.NextToken
	}
	return ""
}

type ListTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId    string          `protobuf:"bytes,1,opt,name=requestId,proto3" form:"requestId" json:"requestId,omitempty"`
	NextToken    string          `protobuf:"bytes,2,opt,name=nextToken,proto3" form:"nextToken" json:"nextToken,omitempty"`
	TagResources []*TagResources `protobuf:"bytes,3,rep,name=TagResources,proto3" form:"tagResources" json:"tagResources,omitempty"`
}

func (x *ListTagResponse) Reset() {
	*x = ListTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTagResponse) ProtoMessage() {}

func (x *ListTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTagResponse.ProtoReflect.Descriptor instead.
func (*ListTagResponse) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{3}
}

func (x *ListTagResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ListTagResponse) GetNextToken() string {
	if x != nil {
		return x.NextToken
	}
	return ""
}

func (x *ListTagResponse) GetTagResources() []*TagResources {
	if x != nil {
		return x.TagResources
	}
	return nil
}

type TagResources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId   string `protobuf:"bytes,1,opt,name=resourceId,proto3" form:"resourceId" json:"resourceId,omitempty"`
	ResourceType string `protobuf:"bytes,2,opt,name=resourceType,proto3" form:"resourceType" json:"resourceType,omitempty"`
	TagKey       string `protobuf:"bytes,3,opt,name=tagKey,proto3" form:"tagKey" json:"tagKey,omitempty"`
	TagValue     string `protobuf:"bytes,4,opt,name=tagValue,proto3" form:"tagValue" json:"tagValue,omitempty"`
}

func (x *TagResources) Reset() {
	*x = TagResources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagResources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagResources) ProtoMessage() {}

func (x *TagResources) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagResources.ProtoReflect.Descriptor instead.
func (*TagResources) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{4}
}

func (x *TagResources) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *TagResources) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *TagResources) GetTagKey() string {
	if x != nil {
		return x.TagKey
	}
	return ""
}

func (x *TagResources) GetTagValue() string {
	if x != nil {
		return x.TagValue
	}
	return ""
}

type UnbindTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId   string   `protobuf:"bytes,1,opt,name=resourceId,proto3" form:"resourceId" json:"resourceId,omitempty"`
	TagKeys      []string `protobuf:"bytes,2,rep,name=tagKeys,proto3" form:"tagKeys" json:"tagKeys,omitempty"`
	ResourceType string   `protobuf:"bytes,3,opt,name=resourceType,proto3" form:"resourceType" json:"resourceType,omitempty"`
	All          bool     `protobuf:"varint,4,opt,name=all,proto3" form:"all" json:"all,omitempty"`
}

func (x *UnbindTagRequest) Reset() {
	*x = UnbindTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnbindTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnbindTagRequest) ProtoMessage() {}

func (x *UnbindTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnbindTagRequest.ProtoReflect.Descriptor instead.
func (*UnbindTagRequest) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{5}
}

func (x *UnbindTagRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *UnbindTagRequest) GetTagKeys() []string {
	if x != nil {
		return x.TagKeys
	}
	return nil
}

func (x *UnbindTagRequest) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *UnbindTagRequest) GetAll() bool {
	if x != nil {
		return x.All
	}
	return false
}

type UnbindTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=requestId,proto3" form:"requestId" json:"requestId,omitempty"`
}

func (x *UnbindTagResponse) Reset() {
	*x = UnbindTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnbindTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnbindTagResponse) ProtoMessage() {}

func (x *UnbindTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnbindTagResponse.ProtoReflect.Descriptor instead.
func (*UnbindTagResponse) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{6}
}

func (x *UnbindTagResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

var File_tag_proto protoreflect.FileDescriptor

var file_tag_proto_rawDesc = []byte{
	0x0a, 0x09, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa0, 0x01, 0x0a, 0x0e, 0x42, 0x69, 0x6e, 0x64, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xca, 0xbb, 0x18, 0x0a, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x67, 0x42, 0x08, 0xca, 0xbb,
	0x18, 0x04, 0x74, 0x61, 0x67, 0x73, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x34, 0x0a, 0x0c,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x10, 0xca, 0xbb, 0x18, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x3e, 0x0a, 0x0f, 0x42, 0x69, 0x6e, 0x64, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xca, 0xbb, 0x18, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x22, 0xf1, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0f, 0xca, 0xbb, 0x18, 0x0b,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x73, 0x52, 0x0b, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x73, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x61, 0x67, 0x42, 0x08, 0xca, 0xbb, 0x18, 0x04, 0x74, 0x61, 0x67, 0x73, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x12, 0x34, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xca, 0xbb, 0x18, 0x0c, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x09, 0xca, 0xbb, 0x18, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x2b, 0x0a, 0x09, 0x6e, 0x65, 0x78,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xca, 0xbb,
	0x18, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x09, 0x6e, 0x65, 0x78,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xb6, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xca,
	0xbb, 0x18, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x52, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xca, 0xbb, 0x18, 0x09,
	0x6e, 0x65, 0x78, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x49, 0x0a, 0x0c, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42,
	0x10, 0xca, 0xbb, 0x18, 0x0c, 0x74, 0x61, 0x67, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x52, 0x0c, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22,
	0xc2, 0x01, 0x0a, 0x0c, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x12, 0x2e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xca, 0xbb, 0x18, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x64, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x34, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xca, 0xbb, 0x18, 0x0c, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x74, 0x61, 0x67, 0x4b, 0x65, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xca, 0xbb, 0x18, 0x06, 0x74, 0x61, 0x67, 0x4b,
	0x65, 0x79, 0x52, 0x06, 0x74, 0x61, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x08, 0x74, 0x61,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xca, 0xbb,
	0x18, 0x08, 0x74, 0x61, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x74, 0x61, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0xba, 0x01, 0x0a, 0x10, 0x55, 0x6e, 0x62, 0x69, 0x6e, 0x64, 0x54,
	0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x0a, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xca,
	0xbb, 0x18, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x52, 0x0a, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x07, 0x74, 0x61, 0x67,
	0x4b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0b, 0xca, 0xbb, 0x18, 0x07,
	0x74, 0x61, 0x67, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x07, 0x74, 0x61, 0x67, 0x4b, 0x65, 0x79, 0x73,
	0x12, 0x34, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xca, 0xbb, 0x18, 0x0c, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x03, 0x61, 0x6c, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x42, 0x07, 0xca, 0xbb, 0x18, 0x03, 0x61, 0x6c, 0x6c, 0x52, 0x03, 0x61, 0x6c,
	0x6c, 0x22, 0x40, 0x0a, 0x11, 0x55, 0x6e, 0x62, 0x69, 0x6e, 0x64, 0x54, 0x61, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xca, 0xbb, 0x18, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x32, 0xc0, 0x02, 0x0a, 0x0a, 0x54, 0x61, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x55, 0x0a, 0x07, 0x42, 0x69, 0x6e, 0x64, 0x54, 0x61, 0x67, 0x12, 0x15, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x69, 0x6e,
	0x64, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0xd2, 0xc1,
	0x18, 0x17, 0x2f, 0x65, 0x63, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61,
	0x67, 0x2f, 0x62, 0x69, 0x6e, 0x64, 0x54, 0x61, 0x67, 0x12, 0x55, 0x0a, 0x07, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x61, 0x67, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1b, 0xd2, 0xc1, 0x18, 0x17, 0x2f, 0x65, 0x63, 0x69, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67,
	0x12, 0x5d, 0x0a, 0x09, 0x55, 0x6e, 0x62, 0x69, 0x6e, 0x64, 0x54, 0x61, 0x67, 0x12, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x6e, 0x62, 0x69, 0x6e, 0x64, 0x54, 0x61, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x6e, 0x62, 0x69, 0x6e, 0x64, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1d, 0xd2, 0xc1, 0x18, 0x19, 0x2f, 0x65, 0x63, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x75, 0x6e, 0x62, 0x69, 0x6e, 0x64, 0x54, 0x61, 0x67, 0x1a,
	0x25, 0x92, 0xce, 0x18, 0x21, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x65, 0x63, 0x69,
	0x2d, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x63, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x74,
	0x79, 0x75, 0x6e, 0x2e, 0x63, 0x6e, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x73, 0x64, 0x6b, 0x2d,
	0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x65, 0x63, 0x69, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x61, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tag_proto_rawDescOnce sync.Once
	file_tag_proto_rawDescData = file_tag_proto_rawDesc
)

func file_tag_proto_rawDescGZIP() []byte {
	file_tag_proto_rawDescOnce.Do(func() {
		file_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_tag_proto_rawDescData)
	})
	return file_tag_proto_rawDescData
}

var file_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_tag_proto_goTypes = []interface{}{
	(*BindTagRequest)(nil),     // 0: proto.BindTagRequest
	(*BindTagResponse)(nil),    // 1: proto.BindTagResponse
	(*ListTagRequest)(nil),     // 2: proto.ListTagRequest
	(*ListTagResponse)(nil),    // 3: proto.ListTagResponse
	(*TagResources)(nil),       // 4: proto.TagResources
	(*UnbindTagRequest)(nil),   // 5: proto.UnbindTagRequest
	(*UnbindTagResponse)(nil),  // 6: proto.UnbindTagResponse
	(*containergroup.Tag)(nil), // 7: proto.Tag
}
var file_tag_proto_depIdxs = []int32{
	7, // 0: proto.BindTagRequest.tags:type_name -> proto.Tag
	7, // 1: proto.ListTagRequest.tags:type_name -> proto.Tag
	4, // 2: proto.ListTagResponse.TagResources:type_name -> proto.TagResources
	0, // 3: proto.TagService.BindTag:input_type -> proto.BindTagRequest
	2, // 4: proto.TagService.ListTag:input_type -> proto.ListTagRequest
	5, // 5: proto.TagService.UnbindTag:input_type -> proto.UnbindTagRequest
	1, // 6: proto.TagService.BindTag:output_type -> proto.BindTagResponse
	3, // 7: proto.TagService.ListTag:output_type -> proto.ListTagResponse
	6, // 8: proto.TagService.UnbindTag:output_type -> proto.UnbindTagResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tag_proto_init() }
func file_tag_proto_init() {
	if File_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindTagRequest); i {
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
		file_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindTagResponse); i {
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
		file_tag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTagRequest); i {
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
		file_tag_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTagResponse); i {
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
		file_tag_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagResources); i {
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
		file_tag_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnbindTagRequest); i {
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
		file_tag_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnbindTagResponse); i {
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
			RawDescriptor: file_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tag_proto_goTypes,
		DependencyIndexes: file_tag_proto_depIdxs,
		MessageInfos:      file_tag_proto_msgTypes,
	}.Build()
	File_tag_proto = out.File
	file_tag_proto_rawDesc = nil
	file_tag_proto_goTypes = nil
	file_tag_proto_depIdxs = nil
}
