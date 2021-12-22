// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: repo_owners.proto

package protocol

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

type Branch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Platform string `protobuf:"bytes,1,opt,name=platform,proto3" json:"platform,omitempty"`
	Org      string `protobuf:"bytes,2,opt,name=org,proto3" json:"org,omitempty"`
	Repo     string `protobuf:"bytes,3,opt,name=repo,proto3" json:"repo,omitempty"`
	Branch   string `protobuf:"bytes,4,opt,name=branch,proto3" json:"branch,omitempty"`
}

func (x *Branch) Reset() {
	*x = Branch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repo_owners_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Branch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Branch) ProtoMessage() {}

func (x *Branch) ProtoReflect() protoreflect.Message {
	mi := &file_repo_owners_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Branch.ProtoReflect.Descriptor instead.
func (*Branch) Descriptor() ([]byte, []int) {
	return file_repo_owners_proto_rawDescGZIP(), []int{0}
}

func (x *Branch) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *Branch) GetOrg() string {
	if x != nil {
		return x.Org
	}
	return ""
}

func (x *Branch) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *Branch) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

type RepoFilePath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Branch *Branch `protobuf:"bytes,1,opt,name=branch,proto3" json:"branch,omitempty"`
	File   string  `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *RepoFilePath) Reset() {
	*x = RepoFilePath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repo_owners_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepoFilePath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepoFilePath) ProtoMessage() {}

func (x *RepoFilePath) ProtoReflect() protoreflect.Message {
	mi := &file_repo_owners_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepoFilePath.ProtoReflect.Descriptor instead.
func (*RepoFilePath) Descriptor() ([]byte, []int) {
	return file_repo_owners_proto_rawDescGZIP(), []int{1}
}

func (x *RepoFilePath) GetBranch() *Branch {
	if x != nil {
		return x.Branch
	}
	return nil
}

func (x *RepoFilePath) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

type Path struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *Path) Reset() {
	*x = Path{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repo_owners_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Path) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Path) ProtoMessage() {}

func (x *Path) ProtoReflect() protoreflect.Message {
	mi := &file_repo_owners_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Path.ProtoReflect.Descriptor instead.
func (*Path) Descriptor() ([]byte, []int) {
	return file_repo_owners_proto_rawDescGZIP(), []int{2}
}

func (x *Path) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type Owners struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owners []string `protobuf:"bytes,1,rep,name=owners,proto3" json:"owners,omitempty"`
}

func (x *Owners) Reset() {
	*x = Owners{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repo_owners_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Owners) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Owners) ProtoMessage() {}

func (x *Owners) ProtoReflect() protoreflect.Message {
	mi := &file_repo_owners_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Owners.ProtoReflect.Descriptor instead.
func (*Owners) Descriptor() ([]byte, []int) {
	return file_repo_owners_proto_rawDescGZIP(), []int{3}
}

func (x *Owners) GetOwners() []string {
	if x != nil {
		return x.Owners
	}
	return nil
}

var File_repo_owners_proto protoreflect.FileDescriptor

var file_repo_owners_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x22,
	0x62, 0x0a, 0x06, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x72, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6f, 0x72, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x22, 0x4e, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x2a, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73,
	0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66,
	0x69, 0x6c, 0x65, 0x22, 0x1a, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22,
	0x20, 0x0a, 0x06, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x73, 0x32, 0x98, 0x04, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73,
	0x12, 0x49, 0x0a, 0x19, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x2e,
	0x72, 0x65, 0x70, 0x6f, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x46,
	0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x1a, 0x10, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x1a, 0x46,
	0x69, 0x6e, 0x64, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x73, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x73, 0x46, 0x6f, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x72, 0x65, 0x70, 0x6f,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x1a, 0x10, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73,
	0x2e, 0x50, 0x61, 0x74, 0x68, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0d, 0x4c, 0x65, 0x61, 0x66, 0x41,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x12, 0x18, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x1a, 0x12, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x2e,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x72, 0x73, 0x12, 0x18, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x1a,
	0x12, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x73, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0d, 0x4c, 0x65, 0x61, 0x66, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x65, 0x72, 0x73, 0x12, 0x18, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x1a, 0x12, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x73, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x09, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x65, 0x72, 0x73, 0x12, 0x18, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73,
	0x2e, 0x52, 0x65, 0x70, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x1a, 0x12, 0x2e,
	0x72, 0x65, 0x70, 0x6f, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x73, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0c, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x65, 0x72, 0x73, 0x12, 0x12, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73,
	0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x1a, 0x12, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x22, 0x00, 0x12, 0x3d, 0x0a,
	0x11, 0x54, 0x6f, 0x70, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x72, 0x73, 0x12, 0x12, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x2e,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x1a, 0x12, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x73, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b,
	0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_repo_owners_proto_rawDescOnce sync.Once
	file_repo_owners_proto_rawDescData = file_repo_owners_proto_rawDesc
)

func file_repo_owners_proto_rawDescGZIP() []byte {
	file_repo_owners_proto_rawDescOnce.Do(func() {
		file_repo_owners_proto_rawDescData = protoimpl.X.CompressGZIP(file_repo_owners_proto_rawDescData)
	})
	return file_repo_owners_proto_rawDescData
}

var file_repo_owners_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_repo_owners_proto_goTypes = []interface{}{
	(*Branch)(nil),       // 0: repoOwners.Branch
	(*RepoFilePath)(nil), // 1: repoOwners.RepoFilePath
	(*Path)(nil),         // 2: repoOwners.Path
	(*Owners)(nil),       // 3: repoOwners.Owners
}
var file_repo_owners_proto_depIdxs = []int32{
	0, // 0: repoOwners.RepoFilePath.branch:type_name -> repoOwners.Branch
	1, // 1: repoOwners.RepoOwners.FindApproverOwnersForFile:input_type -> repoOwners.RepoFilePath
	1, // 2: repoOwners.RepoOwners.FindReviewersOwnersForFile:input_type -> repoOwners.RepoFilePath
	1, // 3: repoOwners.RepoOwners.LeafApprovers:input_type -> repoOwners.RepoFilePath
	1, // 4: repoOwners.RepoOwners.Approvers:input_type -> repoOwners.RepoFilePath
	1, // 5: repoOwners.RepoOwners.LeafReviewers:input_type -> repoOwners.RepoFilePath
	1, // 6: repoOwners.RepoOwners.Reviewers:input_type -> repoOwners.RepoFilePath
	0, // 7: repoOwners.RepoOwners.AllReviewers:input_type -> repoOwners.Branch
	0, // 8: repoOwners.RepoOwners.TopLevelApprovers:input_type -> repoOwners.Branch
	2, // 9: repoOwners.RepoOwners.FindApproverOwnersForFile:output_type -> repoOwners.Path
	2, // 10: repoOwners.RepoOwners.FindReviewersOwnersForFile:output_type -> repoOwners.Path
	3, // 11: repoOwners.RepoOwners.LeafApprovers:output_type -> repoOwners.Owners
	3, // 12: repoOwners.RepoOwners.Approvers:output_type -> repoOwners.Owners
	3, // 13: repoOwners.RepoOwners.LeafReviewers:output_type -> repoOwners.Owners
	3, // 14: repoOwners.RepoOwners.Reviewers:output_type -> repoOwners.Owners
	3, // 15: repoOwners.RepoOwners.AllReviewers:output_type -> repoOwners.Owners
	3, // 16: repoOwners.RepoOwners.TopLevelApprovers:output_type -> repoOwners.Owners
	9, // [9:17] is the sub-list for method output_type
	1, // [1:9] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_repo_owners_proto_init() }
func file_repo_owners_proto_init() {
	if File_repo_owners_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_repo_owners_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Branch); i {
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
		file_repo_owners_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepoFilePath); i {
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
		file_repo_owners_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Path); i {
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
		file_repo_owners_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Owners); i {
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
			RawDescriptor: file_repo_owners_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_repo_owners_proto_goTypes,
		DependencyIndexes: file_repo_owners_proto_depIdxs,
		MessageInfos:      file_repo_owners_proto_msgTypes,
	}.Build()
	File_repo_owners_proto = out.File
	file_repo_owners_proto_rawDesc = nil
	file_repo_owners_proto_goTypes = nil
	file_repo_owners_proto_depIdxs = nil
}