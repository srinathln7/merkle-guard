// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.3
// source: api/v1/proto/merkle-gaurd.proto

package merkle_gaurd

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

type UploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files [][]byte `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_proto_merkle_gaurd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_proto_merkle_gaurd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_proto_merkle_gaurd_proto_rawDescGZIP(), []int{0}
}

func (x *UploadRequest) GetFiles() [][]byte {
	if x != nil {
		return x.Files
	}
	return nil
}

type UploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MerkleRootHash []byte `protobuf:"bytes,1,opt,name=merkle_root_hash,json=merkleRootHash,proto3" json:"merkle_root_hash,omitempty"`
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_proto_merkle_gaurd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_proto_merkle_gaurd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_proto_merkle_gaurd_proto_rawDescGZIP(), []int{1}
}

func (x *UploadResponse) GetMerkleRootHash() []byte {
	if x != nil {
		return x.MerkleRootHash
	}
	return nil
}

type DownloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileIndex int64 `protobuf:"varint,1,opt,name=file_index,json=fileIndex,proto3" json:"file_index,omitempty"`
}

func (x *DownloadRequest) Reset() {
	*x = DownloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_proto_merkle_gaurd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadRequest) ProtoMessage() {}

func (x *DownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_proto_merkle_gaurd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadRequest.ProtoReflect.Descriptor instead.
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_proto_merkle_gaurd_proto_rawDescGZIP(), []int{2}
}

func (x *DownloadRequest) GetFileIndex() int64 {
	if x != nil {
		return x.FileIndex
	}
	return 0
}

type DownloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileContent []byte `protobuf:"bytes,1,opt,name=file_content,json=fileContent,proto3" json:"file_content,omitempty"`
}

func (x *DownloadResponse) Reset() {
	*x = DownloadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_proto_merkle_gaurd_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadResponse) ProtoMessage() {}

func (x *DownloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_proto_merkle_gaurd_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadResponse.ProtoReflect.Descriptor instead.
func (*DownloadResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_proto_merkle_gaurd_proto_rawDescGZIP(), []int{3}
}

func (x *DownloadResponse) GetFileContent() []byte {
	if x != nil {
		return x.FileContent
	}
	return nil
}

type MerkleProofRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileIndex int64 `protobuf:"varint,1,opt,name=file_index,json=fileIndex,proto3" json:"file_index,omitempty"`
}

func (x *MerkleProofRequest) Reset() {
	*x = MerkleProofRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_proto_merkle_gaurd_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerkleProofRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerkleProofRequest) ProtoMessage() {}

func (x *MerkleProofRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_proto_merkle_gaurd_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerkleProofRequest.ProtoReflect.Descriptor instead.
func (*MerkleProofRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_proto_merkle_gaurd_proto_rawDescGZIP(), []int{4}
}

func (x *MerkleProofRequest) GetFileIndex() int64 {
	if x != nil {
		return x.FileIndex
	}
	return 0
}

type MerkleProofResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proof [][]byte `protobuf:"bytes,1,rep,name=proof,proto3" json:"proof,omitempty"`
}

func (x *MerkleProofResponse) Reset() {
	*x = MerkleProofResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_proto_merkle_gaurd_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerkleProofResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerkleProofResponse) ProtoMessage() {}

func (x *MerkleProofResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_proto_merkle_gaurd_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerkleProofResponse.ProtoReflect.Descriptor instead.
func (*MerkleProofResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_proto_merkle_gaurd_proto_rawDescGZIP(), []int{5}
}

func (x *MerkleProofResponse) GetProof() [][]byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

type VerifyProofRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileIndex   int64    `protobuf:"varint,1,opt,name=file_index,json=fileIndex,proto3" json:"file_index,omitempty"`
	FileContent []byte   `protobuf:"bytes,2,opt,name=file_content,json=fileContent,proto3" json:"file_content,omitempty"`
	Proof       [][]byte `protobuf:"bytes,3,rep,name=proof,proto3" json:"proof,omitempty"`
}

func (x *VerifyProofRequest) Reset() {
	*x = VerifyProofRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_proto_merkle_gaurd_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyProofRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyProofRequest) ProtoMessage() {}

func (x *VerifyProofRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_proto_merkle_gaurd_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyProofRequest.ProtoReflect.Descriptor instead.
func (*VerifyProofRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_proto_merkle_gaurd_proto_rawDescGZIP(), []int{6}
}

func (x *VerifyProofRequest) GetFileIndex() int64 {
	if x != nil {
		return x.FileIndex
	}
	return 0
}

func (x *VerifyProofRequest) GetFileContent() []byte {
	if x != nil {
		return x.FileContent
	}
	return nil
}

func (x *VerifyProofRequest) GetProof() [][]byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

type VerifyProofResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Verified bool `protobuf:"varint,1,opt,name=verified,proto3" json:"verified,omitempty"`
}

func (x *VerifyProofResponse) Reset() {
	*x = VerifyProofResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_proto_merkle_gaurd_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyProofResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyProofResponse) ProtoMessage() {}

func (x *VerifyProofResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_proto_merkle_gaurd_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyProofResponse.ProtoReflect.Descriptor instead.
func (*VerifyProofResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_proto_merkle_gaurd_proto_rawDescGZIP(), []int{7}
}

func (x *VerifyProofResponse) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

var File_api_v1_proto_merkle_gaurd_proto protoreflect.FileDescriptor

var file_api_v1_proto_merkle_gaurd_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d,
	0x65, 0x72, 0x6b, 0x6c, 0x65, 0x2d, 0x67, 0x61, 0x75, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x5f, 0x67, 0x61, 0x75, 0x72, 0x64, 0x22,
	0x25, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x3a, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x65, 0x72, 0x6b,
	0x6c, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0e, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x48, 0x61,
	0x73, 0x68, 0x22, 0x30, 0x0a, 0x0f, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x22, 0x35, 0x0a, 0x10, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b,
	0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x33, 0x0a, 0x12, 0x4d,
	0x65, 0x72, 0x6b, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x22, 0x2b, 0x0a, 0x13, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0x6c, 0x0a,
	0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0x31, 0x0a, 0x13, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x32, 0xcd,
	0x02, 0x0a, 0x0a, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x54, 0x72, 0x65, 0x65, 0x12, 0x43, 0x0a,
	0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1b, 0x2e, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65,
	0x5f, 0x67, 0x61, 0x75, 0x72, 0x64, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x5f, 0x67, 0x61,
	0x75, 0x72, 0x64, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x49, 0x0a, 0x08, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1d,
	0x2e, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x5f, 0x67, 0x61, 0x75, 0x72, 0x64, 0x2e, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x5f, 0x67, 0x61, 0x75, 0x72, 0x64, 0x2e, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12,
	0x20, 0x2e, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x5f, 0x67, 0x61, 0x75, 0x72, 0x64, 0x2e, 0x4d,
	0x65, 0x72, 0x6b, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x5f, 0x67, 0x61, 0x75, 0x72, 0x64,
	0x2e, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x11, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4d, 0x65,
	0x72, 0x6b, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x20, 0x2e, 0x6d, 0x65, 0x72, 0x6b,
	0x6c, 0x65, 0x5f, 0x67, 0x61, 0x75, 0x72, 0x64, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50,
	0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x65,
	0x72, 0x6b, 0x6c, 0x65, 0x5f, 0x67, 0x61, 0x75, 0x72, 0x64, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x28,
	0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x72, 0x69,
	0x6e, 0x61, 0x74, 0x68, 0x4c, 0x4e, 0x37, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x72, 0x6b,
	0x6c, 0x65, 0x5f, 0x67, 0x61, 0x75, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_proto_merkle_gaurd_proto_rawDescOnce sync.Once
	file_api_v1_proto_merkle_gaurd_proto_rawDescData = file_api_v1_proto_merkle_gaurd_proto_rawDesc
)

func file_api_v1_proto_merkle_gaurd_proto_rawDescGZIP() []byte {
	file_api_v1_proto_merkle_gaurd_proto_rawDescOnce.Do(func() {
		file_api_v1_proto_merkle_gaurd_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_proto_merkle_gaurd_proto_rawDescData)
	})
	return file_api_v1_proto_merkle_gaurd_proto_rawDescData
}

var file_api_v1_proto_merkle_gaurd_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_v1_proto_merkle_gaurd_proto_goTypes = []interface{}{
	(*UploadRequest)(nil),       // 0: merkle_gaurd.UploadRequest
	(*UploadResponse)(nil),      // 1: merkle_gaurd.UploadResponse
	(*DownloadRequest)(nil),     // 2: merkle_gaurd.DownloadRequest
	(*DownloadResponse)(nil),    // 3: merkle_gaurd.DownloadResponse
	(*MerkleProofRequest)(nil),  // 4: merkle_gaurd.MerkleProofRequest
	(*MerkleProofResponse)(nil), // 5: merkle_gaurd.MerkleProofResponse
	(*VerifyProofRequest)(nil),  // 6: merkle_gaurd.VerifyProofRequest
	(*VerifyProofResponse)(nil), // 7: merkle_gaurd.VerifyProofResponse
}
var file_api_v1_proto_merkle_gaurd_proto_depIdxs = []int32{
	0, // 0: merkle_gaurd.MerkleTree.Upload:input_type -> merkle_gaurd.UploadRequest
	2, // 1: merkle_gaurd.MerkleTree.Download:input_type -> merkle_gaurd.DownloadRequest
	4, // 2: merkle_gaurd.MerkleTree.GetMerkleProof:input_type -> merkle_gaurd.MerkleProofRequest
	6, // 3: merkle_gaurd.MerkleTree.VerifyMerkleProof:input_type -> merkle_gaurd.VerifyProofRequest
	1, // 4: merkle_gaurd.MerkleTree.Upload:output_type -> merkle_gaurd.UploadResponse
	3, // 5: merkle_gaurd.MerkleTree.Download:output_type -> merkle_gaurd.DownloadResponse
	5, // 6: merkle_gaurd.MerkleTree.GetMerkleProof:output_type -> merkle_gaurd.MerkleProofResponse
	7, // 7: merkle_gaurd.MerkleTree.VerifyMerkleProof:output_type -> merkle_gaurd.VerifyProofResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_proto_merkle_gaurd_proto_init() }
func file_api_v1_proto_merkle_gaurd_proto_init() {
	if File_api_v1_proto_merkle_gaurd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_proto_merkle_gaurd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRequest); i {
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
		file_api_v1_proto_merkle_gaurd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadResponse); i {
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
		file_api_v1_proto_merkle_gaurd_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadRequest); i {
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
		file_api_v1_proto_merkle_gaurd_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadResponse); i {
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
		file_api_v1_proto_merkle_gaurd_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerkleProofRequest); i {
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
		file_api_v1_proto_merkle_gaurd_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerkleProofResponse); i {
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
		file_api_v1_proto_merkle_gaurd_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyProofRequest); i {
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
		file_api_v1_proto_merkle_gaurd_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyProofResponse); i {
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
			RawDescriptor: file_api_v1_proto_merkle_gaurd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_proto_merkle_gaurd_proto_goTypes,
		DependencyIndexes: file_api_v1_proto_merkle_gaurd_proto_depIdxs,
		MessageInfos:      file_api_v1_proto_merkle_gaurd_proto_msgTypes,
	}.Build()
	File_api_v1_proto_merkle_gaurd_proto = out.File
	file_api_v1_proto_merkle_gaurd_proto_rawDesc = nil
	file_api_v1_proto_merkle_gaurd_proto_goTypes = nil
	file_api_v1_proto_merkle_gaurd_proto_depIdxs = nil
}