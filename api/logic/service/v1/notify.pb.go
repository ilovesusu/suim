// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: v1/notify.proto

package v1

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

//消息通知
type NotifyHandleType int32

const (
	NotifyHandleType_NOTIFY_UNKNOWN_HANDLE       NotifyHandleType = 0  //未知
	NotifyHandleType_NOTIFY_USER_INFORMATION     NotifyHandleType = 1  //个人消息通知
	NotifyHandleType_NOTIFY_GROUP_INFORMATION    NotifyHandleType = 2  //群消息通知
	NotifyHandleType_NOTIFY_ONLINE               NotifyHandleType = 3  //上线通知
	NotifyHandleType_NOTIFY_OFFLINE              NotifyHandleType = 4  // 下线通知
	NotifyHandleType_NOTIFY_CHANNEL_SUBSCRIPTION NotifyHandleType = 5  // 订阅频道通知
	NotifyHandleType_NOTIFY_ADD_GROUP            NotifyHandleType = 6  // 加入群通知
	NotifyHandleType_NOTIFY_ROOM_SUBSCRIPTION    NotifyHandleType = 8  // 订阅房间通知
	NotifyHandleType_NOTIFY_SYSTEM               NotifyHandleType = 9  // 系统通知
	NotifyHandleType_NOTIFY_WITHDRAW_INFORMATION NotifyHandleType = 10 // 撤回消息通知
	NotifyHandleType_NOTIFY_SPECIAL_CARE         NotifyHandleType = 11 // 特别关注通知
)

// Enum value maps for NotifyHandleType.
var (
	NotifyHandleType_name = map[int32]string{
		0:  "NOTIFY_UNKNOWN_HANDLE",
		1:  "NOTIFY_USER_INFORMATION",
		2:  "NOTIFY_GROUP_INFORMATION",
		3:  "NOTIFY_ONLINE",
		4:  "NOTIFY_OFFLINE",
		5:  "NOTIFY_CHANNEL_SUBSCRIPTION",
		6:  "NOTIFY_ADD_GROUP",
		8:  "NOTIFY_ROOM_SUBSCRIPTION",
		9:  "NOTIFY_SYSTEM",
		10: "NOTIFY_WITHDRAW_INFORMATION",
		11: "NOTIFY_SPECIAL_CARE",
	}
	NotifyHandleType_value = map[string]int32{
		"NOTIFY_UNKNOWN_HANDLE":       0,
		"NOTIFY_USER_INFORMATION":     1,
		"NOTIFY_GROUP_INFORMATION":    2,
		"NOTIFY_ONLINE":               3,
		"NOTIFY_OFFLINE":              4,
		"NOTIFY_CHANNEL_SUBSCRIPTION": 5,
		"NOTIFY_ADD_GROUP":            6,
		"NOTIFY_ROOM_SUBSCRIPTION":    8,
		"NOTIFY_SYSTEM":               9,
		"NOTIFY_WITHDRAW_INFORMATION": 10,
		"NOTIFY_SPECIAL_CARE":         11,
	}
)

func (x NotifyHandleType) Enum() *NotifyHandleType {
	p := new(NotifyHandleType)
	*p = x
	return p
}

func (x NotifyHandleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotifyHandleType) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_notify_proto_enumTypes[0].Descriptor()
}

func (NotifyHandleType) Type() protoreflect.EnumType {
	return &file_v1_notify_proto_enumTypes[0]
}

func (x NotifyHandleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotifyHandleType.Descriptor instead.
func (NotifyHandleType) EnumDescriptor() ([]byte, []int) {
	return file_v1_notify_proto_rawDescGZIP(), []int{0}
}

type NotifyPackageData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type NotifyHandleType `protobuf:"varint,1,opt,name=type,proto3,enum=NotifyHandleType" json:"type,omitempty"`
	Data []byte           `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *NotifyPackageData) Reset() {
	*x = NotifyPackageData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_notify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyPackageData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyPackageData) ProtoMessage() {}

func (x *NotifyPackageData) ProtoReflect() protoreflect.Message {
	mi := &file_v1_notify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyPackageData.ProtoReflect.Descriptor instead.
func (*NotifyPackageData) Descriptor() ([]byte, []int) {
	return file_v1_notify_proto_rawDescGZIP(), []int{0}
}

func (x *NotifyPackageData) GetType() NotifyHandleType {
	if x != nil {
		return x.Type
	}
	return NotifyHandleType_NOTIFY_UNKNOWN_HANDLE
}

func (x *NotifyPackageData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_v1_notify_proto protoreflect.FileDescriptor

var file_v1_notify_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4e, 0x0a, 0x11, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x48, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x2a, 0xb1, 0x02, 0x0a, 0x10, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59,
	0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x48, 0x41, 0x4e, 0x44, 0x4c, 0x45, 0x10,
	0x00, 0x12, 0x1b, 0x0a, 0x17, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x55, 0x53, 0x45, 0x52,
	0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x1c,
	0x0a, 0x18, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x49,
	0x4e, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d,
	0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x4f, 0x4e, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x03, 0x12,
	0x12, 0x0a, 0x0e, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x4f, 0x46, 0x46, 0x4c, 0x49, 0x4e,
	0x45, 0x10, 0x04, 0x12, 0x1f, 0x0a, 0x1b, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x43, 0x48,
	0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x41,
	0x44, 0x44, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0x06, 0x12, 0x1c, 0x0a, 0x18, 0x4e, 0x4f,
	0x54, 0x49, 0x46, 0x59, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52,
	0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x08, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x54, 0x49,
	0x46, 0x59, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x10, 0x09, 0x12, 0x1f, 0x0a, 0x1b, 0x4e,
	0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x44, 0x52, 0x41, 0x57, 0x5f, 0x49,
	0x4e, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0a, 0x12, 0x17, 0x0a, 0x13,
	0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x41, 0x4c, 0x5f, 0x43,
	0x41, 0x52, 0x45, 0x10, 0x0b, 0x42, 0x1e, 0x5a, 0x1c, 0x73, 0x75, 0x69, 0x6d, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_notify_proto_rawDescOnce sync.Once
	file_v1_notify_proto_rawDescData = file_v1_notify_proto_rawDesc
)

func file_v1_notify_proto_rawDescGZIP() []byte {
	file_v1_notify_proto_rawDescOnce.Do(func() {
		file_v1_notify_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_notify_proto_rawDescData)
	})
	return file_v1_notify_proto_rawDescData
}

var file_v1_notify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_notify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_v1_notify_proto_goTypes = []interface{}{
	(NotifyHandleType)(0),     // 0: NotifyHandleType
	(*NotifyPackageData)(nil), // 1: NotifyPackageData
}
var file_v1_notify_proto_depIdxs = []int32{
	0, // 0: NotifyPackageData.type:type_name -> NotifyHandleType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_notify_proto_init() }
func file_v1_notify_proto_init() {
	if File_v1_notify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_notify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyPackageData); i {
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
			RawDescriptor: file_v1_notify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_notify_proto_goTypes,
		DependencyIndexes: file_v1_notify_proto_depIdxs,
		EnumInfos:         file_v1_notify_proto_enumTypes,
		MessageInfos:      file_v1_notify_proto_msgTypes,
	}.Build()
	File_v1_notify_proto = out.File
	file_v1_notify_proto_rawDesc = nil
	file_v1_notify_proto_goTypes = nil
	file_v1_notify_proto_depIdxs = nil
}
