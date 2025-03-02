// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/player.proto

package pbdb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type RoleData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BagMgr  *BagMgr  `protobuf:"bytes,3,opt,name=bagMgr,proto3" json:"bagMgr,omitempty"` //背包
	RoleMgr *RoleMgr `protobuf:"bytes,20,opt,name=roleMgr,proto3" json:"roleMgr,omitempty"`
}

func (x *RoleData) Reset() {
	*x = RoleData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleData) ProtoMessage() {}

func (x *RoleData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleData.ProtoReflect.Descriptor instead.
func (*RoleData) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{0}
}

func (x *RoleData) GetBagMgr() *BagMgr {
	if x != nil {
		return x.BagMgr
	}
	return nil
}

func (x *RoleData) GetRoleMgr() *RoleMgr {
	if x != nil {
		return x.RoleMgr
	}
	return nil
}

type RoleMgr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Setting   *Setting   `protobuf:"bytes,1,opt,name=Setting,proto3" json:"Setting,omitempty"`     //设置
	ExtStatus *ExtStatus `protobuf:"bytes,2,opt,name=extStatus,proto3" json:"extStatus,omitempty"` //玩家附加状态
}

func (x *RoleMgr) Reset() {
	*x = RoleMgr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleMgr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleMgr) ProtoMessage() {}

func (x *RoleMgr) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleMgr.ProtoReflect.Descriptor instead.
func (*RoleMgr) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{1}
}

func (x *RoleMgr) GetSetting() *Setting {
	if x != nil {
		return x.Setting
	}
	return nil
}

func (x *RoleMgr) GetExtStatus() *ExtStatus {
	if x != nil {
		return x.ExtStatus
	}
	return nil
}

//背包
type BagMgr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items   map[int32]int64  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`  //道具
	Equips  map[int32]*Equip `protobuf:"bytes,2,rep,name=equips,proto3" json:"equips,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 每个装备单独ID
	MaxEqId int32            `protobuf:"varint,3,opt,name=maxEqId,proto3" json:"maxEqId,omitempty"`                                                                                       //装备自增id
}

func (x *BagMgr) Reset() {
	*x = BagMgr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BagMgr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BagMgr) ProtoMessage() {}

func (x *BagMgr) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BagMgr.ProtoReflect.Descriptor instead.
func (*BagMgr) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{2}
}

func (x *BagMgr) GetItems() map[int32]int64 {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *BagMgr) GetEquips() map[int32]*Equip {
	if x != nil {
		return x.Equips
	}
	return nil
}

func (x *BagMgr) GetMaxEqId() int32 {
	if x != nil {
		return x.MaxEqId
	}
	return 0
}

var File_proto_player_proto protoreflect.FileDescriptor

var file_proto_player_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1f, 0x0a, 0x06, 0x62, 0x61, 0x67, 0x4d, 0x67, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x07, 0x2e, 0x42, 0x61, 0x67, 0x4d, 0x67, 0x72, 0x52, 0x06, 0x62, 0x61, 0x67,
	0x4d, 0x67, 0x72, 0x12, 0x22, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x4d, 0x67, 0x72, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x67, 0x72, 0x52, 0x07,
	0x72, 0x6f, 0x6c, 0x65, 0x4d, 0x67, 0x72, 0x22, 0x57, 0x0a, 0x07, 0x52, 0x6f, 0x6c, 0x65, 0x4d,
	0x67, 0x72, 0x12, 0x22, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x28, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x45, 0x78, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x65, 0x78, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0xf6, 0x01, 0x0a, 0x06, 0x42, 0x61, 0x67, 0x4d, 0x67, 0x72, 0x12, 0x28, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x42, 0x61, 0x67,
	0x4d, 0x67, 0x72, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x65, 0x71, 0x75, 0x69, 0x70, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x42, 0x61, 0x67, 0x4d, 0x67, 0x72, 0x2e, 0x45,
	0x71, 0x75, 0x69, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x65, 0x71, 0x75, 0x69,
	0x70, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x45, 0x71, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x45, 0x71, 0x49, 0x64, 0x1a, 0x38, 0x0a, 0x0a,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x41, 0x0a, 0x0b, 0x45, 0x71, 0x75, 0x69, 0x70, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x70,
	0x62, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_player_proto_rawDescOnce sync.Once
	file_proto_player_proto_rawDescData = file_proto_player_proto_rawDesc
)

func file_proto_player_proto_rawDescGZIP() []byte {
	file_proto_player_proto_rawDescOnce.Do(func() {
		file_proto_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_player_proto_rawDescData)
	})
	return file_proto_player_proto_rawDescData
}

var file_proto_player_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_player_proto_goTypes = []interface{}{
	(*RoleData)(nil),  // 0: RoleData
	(*RoleMgr)(nil),   // 1: RoleMgr
	(*BagMgr)(nil),    // 2: BagMgr
	nil,               // 3: BagMgr.ItemsEntry
	nil,               // 4: BagMgr.EquipsEntry
	(*Setting)(nil),   // 5: Setting
	(*ExtStatus)(nil), // 6: ExtStatus
	(*Equip)(nil),     // 7: Equip
}
var file_proto_player_proto_depIdxs = []int32{
	2, // 0: RoleData.bagMgr:type_name -> BagMgr
	1, // 1: RoleData.roleMgr:type_name -> RoleMgr
	5, // 2: RoleMgr.Setting:type_name -> Setting
	6, // 3: RoleMgr.extStatus:type_name -> ExtStatus
	3, // 4: BagMgr.items:type_name -> BagMgr.ItemsEntry
	4, // 5: BagMgr.equips:type_name -> BagMgr.EquipsEntry
	7, // 6: BagMgr.EquipsEntry.value:type_name -> Equip
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_player_proto_init() }
func file_proto_player_proto_init() {
	if File_proto_player_proto != nil {
		return
	}
	file_proto_comm_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_player_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleData); i {
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
		file_proto_player_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleMgr); i {
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
		file_proto_player_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BagMgr); i {
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
			RawDescriptor: file_proto_player_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_player_proto_goTypes,
		DependencyIndexes: file_proto_player_proto_depIdxs,
		MessageInfos:      file_proto_player_proto_msgTypes,
	}.Build()
	File_proto_player_proto = out.File
	file_proto_player_proto_rawDesc = nil
	file_proto_player_proto_goTypes = nil
	file_proto_player_proto_depIdxs = nil
}
