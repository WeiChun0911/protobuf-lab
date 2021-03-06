// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: proto/syncData/gameRecord.proto

package syncData

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

type GameRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId        uint32 `protobuf:"varint,1,opt,name=ChannelId,proto3" json:"ChannelId,omitempty"`
	Accounts         string `protobuf:"bytes,2,opt,name=Accounts,proto3" json:"Accounts,omitempty"`
	KindId           uint32 `protobuf:"varint,3,opt,name=KindId,proto3" json:"KindId,omitempty"`
	ServerId         uint32 `protobuf:"varint,4,opt,name=ServerId,proto3" json:"ServerId,omitempty"`
	CellScore        int64  `protobuf:"zigzag64,5,opt,name=CellScore,proto3" json:"CellScore,omitempty"`
	Profit           int64  `protobuf:"zigzag64,6,opt,name=Profit,proto3" json:"Profit,omitempty"`
	CreateTime       string `protobuf:"bytes,7,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	IsNew            bool   `protobuf:"varint,8,opt,name=IsNew,proto3" json:"IsNew,omitempty"`
	LastLoginTime    string `protobuf:"bytes,9,opt,name=LastLoginTime,proto3" json:"LastLoginTime,omitempty"`
	RegisterTime     string `protobuf:"bytes,10,opt,name=RegisterTime,proto3" json:"RegisterTime,omitempty"`
	HistoryCellScore int64  `protobuf:"zigzag64,11,opt,name=HistoryCellScore,proto3" json:"HistoryCellScore,omitempty"`
	HistoryProfit    int64  `protobuf:"zigzag64,12,opt,name=HistoryProfit,proto3" json:"HistoryProfit,omitempty"`
	HistoryGameNum   int64  `protobuf:"zigzag64,13,opt,name=HistoryGameNum,proto3" json:"HistoryGameNum,omitempty"`
}

func (x *GameRecord) Reset() {
	*x = GameRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_syncData_gameRecord_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameRecord) ProtoMessage() {}

func (x *GameRecord) ProtoReflect() protoreflect.Message {
	mi := &file_proto_syncData_gameRecord_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameRecord.ProtoReflect.Descriptor instead.
func (*GameRecord) Descriptor() ([]byte, []int) {
	return file_proto_syncData_gameRecord_proto_rawDescGZIP(), []int{0}
}

func (x *GameRecord) GetChannelId() uint32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *GameRecord) GetAccounts() string {
	if x != nil {
		return x.Accounts
	}
	return ""
}

func (x *GameRecord) GetKindId() uint32 {
	if x != nil {
		return x.KindId
	}
	return 0
}

func (x *GameRecord) GetServerId() uint32 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *GameRecord) GetCellScore() int64 {
	if x != nil {
		return x.CellScore
	}
	return 0
}

func (x *GameRecord) GetProfit() int64 {
	if x != nil {
		return x.Profit
	}
	return 0
}

func (x *GameRecord) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *GameRecord) GetIsNew() bool {
	if x != nil {
		return x.IsNew
	}
	return false
}

func (x *GameRecord) GetLastLoginTime() string {
	if x != nil {
		return x.LastLoginTime
	}
	return ""
}

func (x *GameRecord) GetRegisterTime() string {
	if x != nil {
		return x.RegisterTime
	}
	return ""
}

func (x *GameRecord) GetHistoryCellScore() int64 {
	if x != nil {
		return x.HistoryCellScore
	}
	return 0
}

func (x *GameRecord) GetHistoryProfit() int64 {
	if x != nil {
		return x.HistoryProfit
	}
	return 0
}

func (x *GameRecord) GetHistoryGameNum() int64 {
	if x != nil {
		return x.HistoryGameNum
	}
	return 0
}

var File_proto_syncData_gameRecord_proto protoreflect.FileDescriptor

var file_proto_syncData_gameRecord_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x44, 0x61, 0x74, 0x61,
	0x2f, 0x67, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x03, 0x0a, 0x0a, 0x47, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x4b, 0x69, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x4b, 0x69, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x65, 0x6c, 0x6c, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x12, 0x52, 0x09, 0x43, 0x65, 0x6c, 0x6c, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x12, 0x52, 0x06, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x49,
	0x73, 0x4e, 0x65, 0x77, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x49, 0x73, 0x4e, 0x65,
	0x77, 0x12, 0x24, 0x0a, 0x0d, 0x4c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x4c, 0x61, 0x73, 0x74, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x65, 0x6c, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x12, 0x52, 0x10, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x65,
	0x6c, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x12, 0x52, 0x0d,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x12, 0x26, 0x0a,
	0x0e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x12, 0x52, 0x0e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x47, 0x61,
	0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_syncData_gameRecord_proto_rawDescOnce sync.Once
	file_proto_syncData_gameRecord_proto_rawDescData = file_proto_syncData_gameRecord_proto_rawDesc
)

func file_proto_syncData_gameRecord_proto_rawDescGZIP() []byte {
	file_proto_syncData_gameRecord_proto_rawDescOnce.Do(func() {
		file_proto_syncData_gameRecord_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_syncData_gameRecord_proto_rawDescData)
	})
	return file_proto_syncData_gameRecord_proto_rawDescData
}

var file_proto_syncData_gameRecord_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_syncData_gameRecord_proto_goTypes = []interface{}{
	(*GameRecord)(nil), // 0: proto.GameRecord
}
var file_proto_syncData_gameRecord_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_syncData_gameRecord_proto_init() }
func file_proto_syncData_gameRecord_proto_init() {
	if File_proto_syncData_gameRecord_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_syncData_gameRecord_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameRecord); i {
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
			RawDescriptor: file_proto_syncData_gameRecord_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_syncData_gameRecord_proto_goTypes,
		DependencyIndexes: file_proto_syncData_gameRecord_proto_depIdxs,
		MessageInfos:      file_proto_syncData_gameRecord_proto_msgTypes,
	}.Build()
	File_proto_syncData_gameRecord_proto = out.File
	file_proto_syncData_gameRecord_proto_rawDesc = nil
	file_proto_syncData_gameRecord_proto_goTypes = nil
	file_proto_syncData_gameRecord_proto_depIdxs = nil
}
