// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: example/v1/form.proto

package examplev1

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

// FormModel definition
type FormModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age               int32    `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	DateTime          string   `protobuf:"bytes,4,opt,name=dateTime,proto3" json:"dateTime,omitempty"` // RFC3339 formatted string for DateTime
	Skills            []string `protobuf:"bytes,5,rep,name=skills,proto3" json:"skills,omitempty"`
	City              string   `protobuf:"bytes,6,opt,name=city,proto3" json:"city,omitempty"`
	Range             []int32  `protobuf:"varint,7,rep,packed,name=range,proto3" json:"range,omitempty"`
	TermsAccepted     bool     `protobuf:"varint,8,opt,name=termsAccepted,proto3" json:"termsAccepted,omitempty"`
	SatisfactionLevel int32    `protobuf:"varint,9,opt,name=satisfactionLevel,proto3" json:"satisfactionLevel,omitempty"`
	Gender            string   `protobuf:"bytes,10,opt,name=gender,proto3" json:"gender,omitempty"` // Enum "male", "female", "other"
	Todos             []*Todo  `protobuf:"bytes,11,rep,name=todos,proto3" json:"todos,omitempty"`
}

func (x *FormModel) Reset() {
	*x = FormModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_v1_form_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FormModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FormModel) ProtoMessage() {}

func (x *FormModel) ProtoReflect() protoreflect.Message {
	mi := &file_example_v1_form_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FormModel.ProtoReflect.Descriptor instead.
func (*FormModel) Descriptor() ([]byte, []int) {
	return file_example_v1_form_proto_rawDescGZIP(), []int{0}
}

func (x *FormModel) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FormModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FormModel) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *FormModel) GetDateTime() string {
	if x != nil {
		return x.DateTime
	}
	return ""
}

func (x *FormModel) GetSkills() []string {
	if x != nil {
		return x.Skills
	}
	return nil
}

func (x *FormModel) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *FormModel) GetRange() []int32 {
	if x != nil {
		return x.Range
	}
	return nil
}

func (x *FormModel) GetTermsAccepted() bool {
	if x != nil {
		return x.TermsAccepted
	}
	return false
}

func (x *FormModel) GetSatisfactionLevel() int32 {
	if x != nil {
		return x.SatisfactionLevel
	}
	return 0
}

func (x *FormModel) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *FormModel) GetTodos() []*Todo {
	if x != nil {
		return x.Todos
	}
	return nil
}

// Todo item definition
type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Task   string `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	FormId int64  `protobuf:"varint,3,opt,name=formId,proto3" json:"formId,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_v1_form_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_example_v1_form_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_example_v1_form_proto_rawDescGZIP(), []int{1}
}

func (x *Todo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Todo) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

func (x *Todo) GetFormId() int64 {
	if x != nil {
		return x.FormId
	}
	return 0
}

// Request to create a FormModel
type CreateFormModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FormModel *FormModel `protobuf:"bytes,1,opt,name=formModel,proto3" json:"formModel,omitempty"`
}

func (x *CreateFormModelRequest) Reset() {
	*x = CreateFormModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_v1_form_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFormModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFormModelRequest) ProtoMessage() {}

func (x *CreateFormModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_v1_form_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFormModelRequest.ProtoReflect.Descriptor instead.
func (*CreateFormModelRequest) Descriptor() ([]byte, []int) {
	return file_example_v1_form_proto_rawDescGZIP(), []int{2}
}

func (x *CreateFormModelRequest) GetFormModel() *FormModel {
	if x != nil {
		return x.FormModel
	}
	return nil
}

// Response containing a single FormModel
type GetFormModelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FormModel *FormModel `protobuf:"bytes,1,opt,name=formModel,proto3" json:"formModel,omitempty"`
}

func (x *GetFormModelResponse) Reset() {
	*x = GetFormModelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_v1_form_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFormModelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFormModelResponse) ProtoMessage() {}

func (x *GetFormModelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_v1_form_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFormModelResponse.ProtoReflect.Descriptor instead.
func (*GetFormModelResponse) Descriptor() ([]byte, []int) {
	return file_example_v1_form_proto_rawDescGZIP(), []int{3}
}

func (x *GetFormModelResponse) GetFormModel() *FormModel {
	if x != nil {
		return x.FormModel
	}
	return nil
}

// Request to get a single FormModel by ID
type GetFormModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetFormModelRequest) Reset() {
	*x = GetFormModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_v1_form_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFormModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFormModelRequest) ProtoMessage() {}

func (x *GetFormModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_v1_form_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFormModelRequest.ProtoReflect.Descriptor instead.
func (*GetFormModelRequest) Descriptor() ([]byte, []int) {
	return file_example_v1_form_proto_rawDescGZIP(), []int{4}
}

func (x *GetFormModelRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Request to list FormModels with optional pagination
type ListFormModelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *ListFormModelsRequest) Reset() {
	*x = ListFormModelsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_v1_form_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFormModelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFormModelsRequest) ProtoMessage() {}

func (x *ListFormModelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_v1_form_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFormModelsRequest.ProtoReflect.Descriptor instead.
func (*ListFormModelsRequest) Descriptor() ([]byte, []int) {
	return file_example_v1_form_proto_rawDescGZIP(), []int{5}
}

func (x *ListFormModelsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListFormModelsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// Response for listing FormModels
type ListFormModelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FormModels []*FormModel `protobuf:"bytes,1,rep,name=formModels,proto3" json:"formModels,omitempty"`
	TotalCount int32        `protobuf:"varint,2,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
}

func (x *ListFormModelsResponse) Reset() {
	*x = ListFormModelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_v1_form_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFormModelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFormModelsResponse) ProtoMessage() {}

func (x *ListFormModelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_v1_form_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFormModelsResponse.ProtoReflect.Descriptor instead.
func (*ListFormModelsResponse) Descriptor() ([]byte, []int) {
	return file_example_v1_form_proto_rawDescGZIP(), []int{6}
}

func (x *ListFormModelsResponse) GetFormModels() []*FormModel {
	if x != nil {
		return x.FormModels
	}
	return nil
}

func (x *ListFormModelsResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

// Request to delete a single FormModel by ID
type DeleteFormModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteFormModelRequest) Reset() {
	*x = DeleteFormModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_v1_form_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFormModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFormModelRequest) ProtoMessage() {}

func (x *DeleteFormModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_v1_form_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFormModelRequest.ProtoReflect.Descriptor instead.
func (*DeleteFormModelRequest) Descriptor() ([]byte, []int) {
	return file_example_v1_form_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteFormModelRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Response for delete operation
type DeleteFormModelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteFormModelResponse) Reset() {
	*x = DeleteFormModelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_v1_form_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFormModelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFormModelResponse) ProtoMessage() {}

func (x *DeleteFormModelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_v1_form_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFormModelResponse.ProtoReflect.Descriptor instead.
func (*DeleteFormModelResponse) Descriptor() ([]byte, []int) {
	return file_example_v1_form_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteFormModelResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_example_v1_form_proto protoreflect.FileDescriptor

var file_example_v1_form_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x22, 0xb3, 0x02, 0x0a, 0x09, 0x46, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x07, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x41, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x74, 0x65,
	0x72, 0x6d, 0x73, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x73,
	0x61, 0x74, 0x69, 0x73, 0x66, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x73, 0x61, 0x74, 0x69, 0x73, 0x66, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x26, 0x0a, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f,
	0x64, 0x6f, 0x52, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x22, 0x42, 0x0a, 0x04, 0x54, 0x6f, 0x64,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x22, 0x4d, 0x0a,
	0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x09, 0x66, 0x6f, 0x72, 0x6d, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x52, 0x09, 0x66, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x4b, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x66, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x09,
	0x66, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x46, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x47, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x6f, 0x0a, 0x16, 0x4c, 0x69, 0x73,
	0x74, 0x46, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x66, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x0a,
	0x66, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x28, 0x0a, 0x16, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x33, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f,
	0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xf1, 0x02, 0x0a, 0x10, 0x46, 0x6f,
	0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55,
	0x0a, 0x0d, 0x53, 0x61, 0x76, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x22, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x21, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x6d,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46,
	0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5a, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x22, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x6d,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x22, 0x5a,
	0x20, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_v1_form_proto_rawDescOnce sync.Once
	file_example_v1_form_proto_rawDescData = file_example_v1_form_proto_rawDesc
)

func file_example_v1_form_proto_rawDescGZIP() []byte {
	file_example_v1_form_proto_rawDescOnce.Do(func() {
		file_example_v1_form_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_v1_form_proto_rawDescData)
	})
	return file_example_v1_form_proto_rawDescData
}

var file_example_v1_form_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_example_v1_form_proto_goTypes = []any{
	(*FormModel)(nil),               // 0: example.v1.FormModel
	(*Todo)(nil),                    // 1: example.v1.Todo
	(*CreateFormModelRequest)(nil),  // 2: example.v1.CreateFormModelRequest
	(*GetFormModelResponse)(nil),    // 3: example.v1.GetFormModelResponse
	(*GetFormModelRequest)(nil),     // 4: example.v1.GetFormModelRequest
	(*ListFormModelsRequest)(nil),   // 5: example.v1.ListFormModelsRequest
	(*ListFormModelsResponse)(nil),  // 6: example.v1.ListFormModelsResponse
	(*DeleteFormModelRequest)(nil),  // 7: example.v1.DeleteFormModelRequest
	(*DeleteFormModelResponse)(nil), // 8: example.v1.DeleteFormModelResponse
}
var file_example_v1_form_proto_depIdxs = []int32{
	1, // 0: example.v1.FormModel.todos:type_name -> example.v1.Todo
	0, // 1: example.v1.CreateFormModelRequest.formModel:type_name -> example.v1.FormModel
	0, // 2: example.v1.GetFormModelResponse.formModel:type_name -> example.v1.FormModel
	0, // 3: example.v1.ListFormModelsResponse.formModels:type_name -> example.v1.FormModel
	2, // 4: example.v1.FormModelService.SaveFormModel:input_type -> example.v1.CreateFormModelRequest
	4, // 5: example.v1.FormModelService.GetFormModel:input_type -> example.v1.GetFormModelRequest
	5, // 6: example.v1.FormModelService.ListFormModels:input_type -> example.v1.ListFormModelsRequest
	7, // 7: example.v1.FormModelService.DeleteFormModel:input_type -> example.v1.DeleteFormModelRequest
	3, // 8: example.v1.FormModelService.SaveFormModel:output_type -> example.v1.GetFormModelResponse
	3, // 9: example.v1.FormModelService.GetFormModel:output_type -> example.v1.GetFormModelResponse
	6, // 10: example.v1.FormModelService.ListFormModels:output_type -> example.v1.ListFormModelsResponse
	8, // 11: example.v1.FormModelService.DeleteFormModel:output_type -> example.v1.DeleteFormModelResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_example_v1_form_proto_init() }
func file_example_v1_form_proto_init() {
	if File_example_v1_form_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_v1_form_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*FormModel); i {
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
		file_example_v1_form_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Todo); i {
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
		file_example_v1_form_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateFormModelRequest); i {
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
		file_example_v1_form_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetFormModelResponse); i {
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
		file_example_v1_form_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetFormModelRequest); i {
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
		file_example_v1_form_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ListFormModelsRequest); i {
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
		file_example_v1_form_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ListFormModelsResponse); i {
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
		file_example_v1_form_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteFormModelRequest); i {
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
		file_example_v1_form_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteFormModelResponse); i {
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
			RawDescriptor: file_example_v1_form_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_v1_form_proto_goTypes,
		DependencyIndexes: file_example_v1_form_proto_depIdxs,
		MessageInfos:      file_example_v1_form_proto_msgTypes,
	}.Build()
	File_example_v1_form_proto = out.File
	file_example_v1_form_proto_rawDesc = nil
	file_example_v1_form_proto_goTypes = nil
	file_example_v1_form_proto_depIdxs = nil
}
