// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.24.4
// source: file_create.proto

package proto

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

// *
// Create a new file, containing the given contents.
// After the file is created, the FileID for it can be found in the receipt, or record, or retrieved
// with a GetByKey query.
//
// The file contains the specified contents (possibly empty). The file will automatically disappear
// at the expirationTime, unless its expiration is extended by another transaction before that time.
// If the file is deleted, then its contents will become empty and it will be marked as deleted
// until it expires, and then it will cease to exist.
//
// The keys field is a list of keys. All keys within the top-level key list must sign (M-M) to
// create or modify a file. However, to delete the file, only one key (1-M) is required to sign from
// the top-level key list.  Each of those "keys" may itself be threshold key containing other keys
// (including other threshold keys). In other words, the behavior is an AND for create/modify, OR
// for delete. This is useful for acting as a revocation server. If it is desired to have the
// behavior be AND for all 3 operations (or OR for all 3), then the list should have only a single
// Key, which is a threshold key, with N=1 for OR, N=M for AND. If the auto_renew_account field
// is set, the key of the referenced account must sign.
//
// If a file is created without ANY keys in the keys field, the file is immutable and ONLY the
// expirationTime of the file can be changed with a FileUpdate transaction. The file contents or its
// keys cannot be changed.
//
// An entity (account, file, or smart contract instance) must be created in a particular realm. If
// the realmID is left null, then a new realm will be created with the given admin key. If a new
// realm has a null adminKey, then anyone can create/modify/delete entities in that realm. But if an
// admin key is given, then any transaction to create/modify/delete an entity in that realm must be
// signed by that key, though anyone can still call functions on smart contract instances that exist
// in that realm. A realm ceases to exist when everything within it has expired and no longer
// exists.
//
// The current API ignores shardID, realmID, and newRealmAdminKey, and creates everything in shard 0
// and realm 0, with a null key. Future versions of the API will support multiple realms and
// multiple shards.
type FileCreateTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// The time at which this file should expire (unless FileUpdateTransaction is used before then
	// to extend its life)
	ExpirationTime *Timestamp `protobuf:"bytes,2,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
	// *
	// All keys at the top level of a key list must sign to create or modify the file. Any one of
	// the keys at the top level key list can sign to delete the file.
	Keys *KeyList `protobuf:"bytes,3,opt,name=keys,proto3" json:"keys,omitempty"`
	// *
	// The bytes that are the contents of the file
	Contents []byte `protobuf:"bytes,4,opt,name=contents,proto3" json:"contents,omitempty"`
	// *
	// Shard in which this file is created
	ShardID *ShardID `protobuf:"bytes,5,opt,name=shardID,proto3" json:"shardID,omitempty"`
	// *
	// The Realm in which to the file is created (leave this null to create a new realm)
	RealmID *RealmID `protobuf:"bytes,6,opt,name=realmID,proto3" json:"realmID,omitempty"`
	// *
	// If realmID is null, then this the admin key for the new realm that will be created
	NewRealmAdminKey *Key `protobuf:"bytes,7,opt,name=newRealmAdminKey,proto3" json:"newRealmAdminKey,omitempty"`
	// *
	// The memo associated with the file (UTF-8 encoding max 100 bytes)
	Memo string `protobuf:"bytes,8,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (x *FileCreateTransactionBody) Reset() {
	*x = FileCreateTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileCreateTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileCreateTransactionBody) ProtoMessage() {}

func (x *FileCreateTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_file_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileCreateTransactionBody.ProtoReflect.Descriptor instead.
func (*FileCreateTransactionBody) Descriptor() ([]byte, []int) {
	return file_file_create_proto_rawDescGZIP(), []int{0}
}

func (x *FileCreateTransactionBody) GetExpirationTime() *Timestamp {
	if x != nil {
		return x.ExpirationTime
	}
	return nil
}

func (x *FileCreateTransactionBody) GetKeys() *KeyList {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *FileCreateTransactionBody) GetContents() []byte {
	if x != nil {
		return x.Contents
	}
	return nil
}

func (x *FileCreateTransactionBody) GetShardID() *ShardID {
	if x != nil {
		return x.ShardID
	}
	return nil
}

func (x *FileCreateTransactionBody) GetRealmID() *RealmID {
	if x != nil {
		return x.RealmID
	}
	return nil
}

func (x *FileCreateTransactionBody) GetNewRealmAdminKey() *Key {
	if x != nil {
		return x.NewRealmAdminKey
	}
	return nil
}

func (x *FileCreateTransactionBody) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

var File_file_create_proto protoreflect.FileDescriptor

var file_file_create_proto_rawDesc = []byte{
	0x0a, 0x11, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5,
	0x02, 0x0a, 0x19, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x38, 0x0a, 0x0e,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x07, 0x73, 0x68, 0x61, 0x72, 0x64, 0x49,
	0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x68, 0x61, 0x72, 0x64, 0x49, 0x44, 0x52, 0x07, 0x73, 0x68, 0x61, 0x72, 0x64, 0x49, 0x44,
	0x12, 0x28, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x49,
	0x44, 0x52, 0x07, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x49, 0x44, 0x12, 0x36, 0x0a, 0x10, 0x6e, 0x65,
	0x77, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79,
	0x52, 0x10, 0x6e, 0x65, 0x77, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4b,
	0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_file_create_proto_rawDescOnce sync.Once
	file_file_create_proto_rawDescData = file_file_create_proto_rawDesc
)

func file_file_create_proto_rawDescGZIP() []byte {
	file_file_create_proto_rawDescOnce.Do(func() {
		file_file_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_create_proto_rawDescData)
	})
	return file_file_create_proto_rawDescData
}

var file_file_create_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_file_create_proto_goTypes = []interface{}{
	(*FileCreateTransactionBody)(nil), // 0: proto.FileCreateTransactionBody
	(*Timestamp)(nil),                 // 1: proto.Timestamp
	(*KeyList)(nil),                   // 2: proto.KeyList
	(*ShardID)(nil),                   // 3: proto.ShardID
	(*RealmID)(nil),                   // 4: proto.RealmID
	(*Key)(nil),                       // 5: proto.Key
}
var file_file_create_proto_depIdxs = []int32{
	1, // 0: proto.FileCreateTransactionBody.expirationTime:type_name -> proto.Timestamp
	2, // 1: proto.FileCreateTransactionBody.keys:type_name -> proto.KeyList
	3, // 2: proto.FileCreateTransactionBody.shardID:type_name -> proto.ShardID
	4, // 3: proto.FileCreateTransactionBody.realmID:type_name -> proto.RealmID
	5, // 4: proto.FileCreateTransactionBody.newRealmAdminKey:type_name -> proto.Key
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_file_create_proto_init() }
func file_file_create_proto_init() {
	if File_file_create_proto != nil {
		return
	}
	file_basic_types_proto_init()
	file_timestamp_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_file_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileCreateTransactionBody); i {
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
			RawDescriptor: file_file_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_file_create_proto_goTypes,
		DependencyIndexes: file_file_create_proto_depIdxs,
		MessageInfos:      file_file_create_proto_msgTypes,
	}.Build()
	File_file_create_proto = out.File
	file_file_create_proto_rawDesc = nil
	file_file_create_proto_goTypes = nil
	file_file_create_proto_depIdxs = nil
}
