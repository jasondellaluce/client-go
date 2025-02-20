//
//Copyright (C) 2020 The Falco Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: outputs.proto

package outputs

import (
	schema "github.com/falcosecurity/client-go/pkg/api/schema"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The `request` message is the logical representation of the request model.
// It is the input of the `output.service` service.
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outputs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_outputs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_outputs_proto_rawDescGZIP(), []int{0}
}

// The `response` message is the representation of the output model.
// It contains all the elements that Falco emits in an output along with the
// definitions for priorities and source.
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time         *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Priority     schema.Priority      `protobuf:"varint,2,opt,name=priority,proto3,enum=falco.schema.Priority" json:"priority,omitempty"`
	Source       schema.Source        `protobuf:"varint,3,opt,name=source,proto3,enum=falco.schema.Source" json:"source,omitempty"`
	Rule         string               `protobuf:"bytes,4,opt,name=rule,proto3" json:"rule,omitempty"`
	Output       string               `protobuf:"bytes,5,opt,name=output,proto3" json:"output,omitempty"`
	OutputFields map[string]string    `protobuf:"bytes,6,rep,name=output_fields,json=outputFields,proto3" json:"output_fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Hostname     string               `protobuf:"bytes,7,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Tags         []string             `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outputs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_outputs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_outputs_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Response) GetPriority() schema.Priority {
	if x != nil {
		return x.Priority
	}
	return schema.Priority_EMERGENCY
}

func (x *Response) GetSource() schema.Source {
	if x != nil {
		return x.Source
	}
	return schema.Source_SYSCALL
}

func (x *Response) GetRule() string {
	if x != nil {
		return x.Rule
	}
	return ""
}

func (x *Response) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *Response) GetOutputFields() map[string]string {
	if x != nil {
		return x.OutputFields
	}
	return nil
}

func (x *Response) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Response) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_outputs_proto protoreflect.FileDescriptor

var file_outputs_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x66, 0x61, 0x6c, 0x63, 0x6f, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0c, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x09, 0x0a,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x89, 0x03, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x66, 0x61, 0x6c, 0x63, 0x6f, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x52,
	0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x66, 0x61, 0x6c, 0x63,
	0x6f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x12, 0x4e, 0x0a, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x66, 0x61, 0x6c,
	0x63, 0x6f, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x1a, 0x3f, 0x0a, 0x11, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x32, 0x7f, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3a, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x12, 0x16, 0x2e, 0x66, 0x61, 0x6c, 0x63, 0x6f, 0x2e, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x66, 0x61, 0x6c, 0x63, 0x6f, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x2e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x38, 0x0a, 0x03, 0x67,
	0x65, 0x74, 0x12, 0x16, 0x2e, 0x66, 0x61, 0x6c, 0x63, 0x6f, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x66, 0x61, 0x6c,
	0x63, 0x6f, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x61, 0x6c, 0x63, 0x6f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_outputs_proto_rawDescOnce sync.Once
	file_outputs_proto_rawDescData = file_outputs_proto_rawDesc
)

func file_outputs_proto_rawDescGZIP() []byte {
	file_outputs_proto_rawDescOnce.Do(func() {
		file_outputs_proto_rawDescData = protoimpl.X.CompressGZIP(file_outputs_proto_rawDescData)
	})
	return file_outputs_proto_rawDescData
}

var file_outputs_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_outputs_proto_goTypes = []interface{}{
	(*Request)(nil),             // 0: falco.outputs.request
	(*Response)(nil),            // 1: falco.outputs.response
	nil,                         // 2: falco.outputs.response.OutputFieldsEntry
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(schema.Priority)(0),        // 4: falco.schema.priority
	(schema.Source)(0),          // 5: falco.schema.source
}
var file_outputs_proto_depIdxs = []int32{
	3, // 0: falco.outputs.response.time:type_name -> google.protobuf.Timestamp
	4, // 1: falco.outputs.response.priority:type_name -> falco.schema.priority
	5, // 2: falco.outputs.response.source:type_name -> falco.schema.source
	2, // 3: falco.outputs.response.output_fields:type_name -> falco.outputs.response.OutputFieldsEntry
	0, // 4: falco.outputs.service.sub:input_type -> falco.outputs.request
	0, // 5: falco.outputs.service.get:input_type -> falco.outputs.request
	1, // 6: falco.outputs.service.sub:output_type -> falco.outputs.response
	1, // 7: falco.outputs.service.get:output_type -> falco.outputs.response
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_outputs_proto_init() }
func file_outputs_proto_init() {
	if File_outputs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_outputs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_outputs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_outputs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_outputs_proto_goTypes,
		DependencyIndexes: file_outputs_proto_depIdxs,
		MessageInfos:      file_outputs_proto_msgTypes,
	}.Build()
	File_outputs_proto = out.File
	file_outputs_proto_rawDesc = nil
	file_outputs_proto_goTypes = nil
	file_outputs_proto_depIdxs = nil
}
