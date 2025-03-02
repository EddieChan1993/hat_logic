// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/comm.proto

package pbdb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type Equip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                        // 配置表ID
	Lv         int32   `protobuf:"varint,2,opt,name=lv,proto3" json:"lv,omitempty"`                        //强化等级
	ExtAttrIds []int32 `protobuf:"varint,3,rep,packed,name=extAttrIds,proto3" json:"extAttrIds,omitempty"` //附加属性
}

func (x *Equip) Reset() {
	*x = Equip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Equip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Equip) ProtoMessage() {}

func (x *Equip) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Equip.ProtoReflect.Descriptor instead.
func (*Equip) Descriptor() ([]byte, []int) {
	return file_proto_comm_proto_rawDescGZIP(), []int{0}
}

func (x *Equip) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Equip) GetLv() int32 {
	if x != nil {
		return x.Lv
	}
	return 0
}

func (x *Equip) GetExtAttrIds() []int32 {
	if x != nil {
		return x.ExtAttrIds
	}
	return nil
}

type Setting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChangedName bool `protobuf:"varint,4,opt,name=changedName,proto3" json:"changedName,omitempty"` //是否改过名字 true-是 false-不是
}

func (x *Setting) Reset() {
	*x = Setting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Setting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Setting) ProtoMessage() {}

func (x *Setting) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Setting.ProtoReflect.Descriptor instead.
func (*Setting) Descriptor() ([]byte, []int) {
	return file_proto_comm_proto_rawDescGZIP(), []int{1}
}

func (x *Setting) GetChangedName() bool {
	if x != nil {
		return x.ChangedName
	}
	return false
}

type ExtStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastLoginZeroAt int64 `protobuf:"varint,1,opt,name=lastLoginZeroAt,proto3" json:"lastLoginZeroAt,omitempty"` //最近一次登陆时间
}

func (x *ExtStatus) Reset() {
	*x = ExtStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtStatus) ProtoMessage() {}

func (x *ExtStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtStatus.ProtoReflect.Descriptor instead.
func (*ExtStatus) Descriptor() ([]byte, []int) {
	return file_proto_comm_proto_rawDescGZIP(), []int{2}
}

func (x *ExtStatus) GetLastLoginZeroAt() int64 {
	if x != nil {
		return x.LastLoginZeroAt
	}
	return 0
}

var File_proto_comm_proto protoreflect.FileDescriptor

var file_proto_comm_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x47, 0x0a, 0x05, 0x45, 0x71, 0x75, 0x69, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x6c,
	0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x6c, 0x76, 0x12, 0x1e, 0x0a, 0x0a, 0x65,
	0x78, 0x74, 0x41, 0x74, 0x74, 0x72, 0x49, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x0a, 0x65, 0x78, 0x74, 0x41, 0x74, 0x74, 0x72, 0x49, 0x64, 0x73, 0x22, 0x2b, 0x0a, 0x07, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x35, 0x0a, 0x09, 0x45, 0x78, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x5a, 0x65, 0x72, 0x6f, 0x41, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f,
	0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x5a, 0x65, 0x72, 0x6f, 0x41, 0x74, 0x42,
	0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x70, 0x62, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x64, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_comm_proto_rawDescOnce sync.Once
	file_proto_comm_proto_rawDescData = file_proto_comm_proto_rawDesc
)

func file_proto_comm_proto_rawDescGZIP() []byte {
	file_proto_comm_proto_rawDescOnce.Do(func() {
		file_proto_comm_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_comm_proto_rawDescData)
	})
	return file_proto_comm_proto_rawDescData
}

var file_proto_comm_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_comm_proto_goTypes = []interface{}{
	(*Equip)(nil),     // 0: Equip
	(*Setting)(nil),   // 1: Setting
	(*ExtStatus)(nil), // 2: ExtStatus
}
var file_proto_comm_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_comm_proto_init() }
func file_proto_comm_proto_init() {
	if File_proto_comm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_comm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Equip); i {
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
		file_proto_comm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Setting); i {
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
		file_proto_comm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtStatus); i {
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
			RawDescriptor: file_proto_comm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_comm_proto_goTypes,
		DependencyIndexes: file_proto_comm_proto_depIdxs,
		MessageInfos:      file_proto_comm_proto_msgTypes,
	}.Build()
	File_proto_comm_proto = out.File
	file_proto_comm_proto_rawDesc = nil
	file_proto_comm_proto_goTypes = nil
	file_proto_comm_proto_depIdxs = nil
}
