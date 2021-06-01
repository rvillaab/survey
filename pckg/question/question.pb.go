// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: question.proto

package question

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request message containing the user's name.
type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_question_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_question_proto_rawDescGZIP(), []int{0}
}

// The request message containing the user's name.
type RequestWithId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RequestWithId) Reset() {
	*x = RequestWithId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestWithId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestWithId) ProtoMessage() {}

func (x *RequestWithId) ProtoReflect() protoreflect.Message {
	mi := &file_question_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestWithId.ProtoReflect.Descriptor instead.
func (*RequestWithId) Descriptor() ([]byte, []int) {
	return file_question_proto_rawDescGZIP(), []int{1}
}

func (x *RequestWithId) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_question_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_question_proto_rawDescGZIP(), []int{2}
}

func (x *Result) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Content     string                 `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UserCreated string                 `protobuf:"bytes,5,opt,name=UserCreated,proto3" json:"UserCreated,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	UserUpdated string                 `protobuf:"bytes,7,opt,name=UserUpdated,proto3" json:"UserUpdated,omitempty"`
}

func (x *Question) Reset() {
	*x = Question{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_question_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Question.ProtoReflect.Descriptor instead.
func (*Question) Descriptor() ([]byte, []int) {
	return file_question_proto_rawDescGZIP(), []int{3}
}

func (x *Question) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Question) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Question) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Question) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Question) GetUserCreated() string {
	if x != nil {
		return x.UserCreated
	}
	return ""
}

func (x *Question) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Question) GetUserUpdated() string {
	if x != nil {
		return x.UserUpdated
	}
	return ""
}

type AllQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Questions []*Question `protobuf:"bytes,1,rep,name=questions,proto3" json:"questions,omitempty"`
}

func (x *AllQuestionResponse) Reset() {
	*x = AllQuestionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllQuestionResponse) ProtoMessage() {}

func (x *AllQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_question_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllQuestionResponse.ProtoReflect.Descriptor instead.
func (*AllQuestionResponse) Descriptor() ([]byte, []int) {
	return file_question_proto_rawDescGZIP(), []int{4}
}

func (x *AllQuestionResponse) GetQuestions() []*Question {
	if x != nil {
		return x.Questions
	}
	return nil
}

var File_question_proto protoreflect.FileDescriptor

var file_question_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0e, 0x0a, 0x0c, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x23, 0x0a, 0x0d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x22, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x8e, 0x02, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x47, 0x0a, 0x13, 0x41, 0x6c, 0x6c, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x09,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xcd,
	0x02, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x36, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x10, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x48, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x41, 0x6c, 0x6c, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x49, 0x64, 0x1a,
	0x10, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x3b, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3e,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x17, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x16,
	0x5a, 0x14, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2f, 0x70, 0x63, 0x6b, 0x67, 0x2f, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_question_proto_rawDescOnce sync.Once
	file_question_proto_rawDescData = file_question_proto_rawDesc
)

func file_question_proto_rawDescGZIP() []byte {
	file_question_proto_rawDescOnce.Do(func() {
		file_question_proto_rawDescData = protoimpl.X.CompressGZIP(file_question_proto_rawDescData)
	})
	return file_question_proto_rawDescData
}

var file_question_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_question_proto_goTypes = []interface{}{
	(*EmptyRequest)(nil),          // 0: question.EmptyRequest
	(*RequestWithId)(nil),         // 1: question.RequestWithId
	(*Result)(nil),                // 2: question.Result
	(*Question)(nil),              // 3: question.Question
	(*AllQuestionResponse)(nil),   // 4: question.AllQuestionResponse
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_question_proto_depIdxs = []int32{
	5, // 0: question.Question.CreatedAt:type_name -> google.protobuf.Timestamp
	5, // 1: question.Question.UpdatedAt:type_name -> google.protobuf.Timestamp
	3, // 2: question.AllQuestionResponse.questions:type_name -> question.Question
	3, // 3: question.QuestionService.CreateQuestion:input_type -> question.Question
	0, // 4: question.QuestionService.GetAllQuestions:input_type -> question.EmptyRequest
	1, // 5: question.QuestionService.UpdateQuestion:input_type -> question.RequestWithId
	1, // 6: question.QuestionService.DeleteQuestion:input_type -> question.RequestWithId
	1, // 7: question.QuestionService.GetQuestionById:input_type -> question.RequestWithId
	2, // 8: question.QuestionService.CreateQuestion:output_type -> question.Result
	4, // 9: question.QuestionService.GetAllQuestions:output_type -> question.AllQuestionResponse
	2, // 10: question.QuestionService.UpdateQuestion:output_type -> question.Result
	2, // 11: question.QuestionService.DeleteQuestion:output_type -> question.Result
	3, // 12: question.QuestionService.GetQuestionById:output_type -> question.Question
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_question_proto_init() }
func file_question_proto_init() {
	if File_question_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_question_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
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
		file_question_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestWithId); i {
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
		file_question_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_question_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Question); i {
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
		file_question_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllQuestionResponse); i {
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
			RawDescriptor: file_question_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_question_proto_goTypes,
		DependencyIndexes: file_question_proto_depIdxs,
		MessageInfos:      file_question_proto_msgTypes,
	}.Build()
	File_question_proto = out.File
	file_question_proto_rawDesc = nil
	file_question_proto_goTypes = nil
	file_question_proto_depIdxs = nil
}
