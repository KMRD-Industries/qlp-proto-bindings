// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: comm.proto

package _go

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

type StateVariant int32

const (
	StateVariant_NONE         StateVariant = 0
	StateVariant_CONNECTED    StateVariant = 1
	StateVariant_DISCONNECTED StateVariant = 2
	StateVariant_ROOM_CHANGED StateVariant = 3
)

// Enum value maps for StateVariant.
var (
	StateVariant_name = map[int32]string{
		0: "NONE",
		1: "CONNECTED",
		2: "DISCONNECTED",
		3: "ROOM_CHANGED",
	}
	StateVariant_value = map[string]int32{
		"NONE":         0,
		"CONNECTED":    1,
		"DISCONNECTED": 2,
		"ROOM_CHANGED": 3,
	}
)

func (x StateVariant) Enum() *StateVariant {
	p := new(StateVariant)
	*p = x
	return p
}

func (x StateVariant) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StateVariant) Descriptor() protoreflect.EnumDescriptor {
	return file_comm_proto_enumTypes[0].Descriptor()
}

func (StateVariant) Type() protoreflect.EnumType {
	return &file_comm_proto_enumTypes[0]
}

func (x StateVariant) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StateVariant.Descriptor instead.
func (StateVariant) EnumDescriptor() ([]byte, []int) {
	return file_comm_proto_rawDescGZIP(), []int{0}
}

// sent via UDP
type MovementUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId     uint32  `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	PositionX    float32 `protobuf:"fixed32,2,opt,name=position_x,json=positionX,proto3" json:"position_x,omitempty"`
	PositionY    float32 `protobuf:"fixed32,3,opt,name=position_y,json=positionY,proto3" json:"position_y,omitempty"`
	WeaponPivotX float32 `protobuf:"fixed32,4,opt,name=weapon_pivot_x,json=weaponPivotX,proto3" json:"weapon_pivot_x,omitempty"`
	WeaponPivotY float32 `protobuf:"fixed32,5,opt,name=weapon_pivot_y,json=weaponPivotY,proto3" json:"weapon_pivot_y,omitempty"`
	Direction    float32 `protobuf:"fixed32,6,opt,name=direction,proto3" json:"direction,omitempty"`
	Attack       bool    `protobuf:"varint,7,opt,name=attack,proto3" json:"attack,omitempty"`
	CurrRoom     *Room   `protobuf:"bytes,8,opt,name=curr_room,json=currRoom,proto3" json:"curr_room,omitempty"`
}

func (x *MovementUpdate) Reset() {
	*x = MovementUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovementUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovementUpdate) ProtoMessage() {}

func (x *MovementUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_comm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovementUpdate.ProtoReflect.Descriptor instead.
func (*MovementUpdate) Descriptor() ([]byte, []int) {
	return file_comm_proto_rawDescGZIP(), []int{0}
}

func (x *MovementUpdate) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *MovementUpdate) GetPositionX() float32 {
	if x != nil {
		return x.PositionX
	}
	return 0
}

func (x *MovementUpdate) GetPositionY() float32 {
	if x != nil {
		return x.PositionY
	}
	return 0
}

func (x *MovementUpdate) GetWeaponPivotX() float32 {
	if x != nil {
		return x.WeaponPivotX
	}
	return 0
}

func (x *MovementUpdate) GetWeaponPivotY() float32 {
	if x != nil {
		return x.WeaponPivotY
	}
	return 0
}

func (x *MovementUpdate) GetDirection() float32 {
	if x != nil {
		return x.Direction
	}
	return 0
}

func (x *MovementUpdate) GetAttack() bool {
	if x != nil {
		return x.Attack
	}
	return false
}

func (x *MovementUpdate) GetCurrRoom() *Room {
	if x != nil {
		return x.CurrRoom
	}
	return nil
}

// sent via TCP
type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Weapon *Weapon `protobuf:"bytes,2,opt,name=weapon,proto3" json:"weapon,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_comm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_comm_proto_rawDescGZIP(), []int{1}
}

func (x *Player) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Player) GetWeapon() *Weapon {
	if x != nil {
		return x.Weapon
	}
	return nil
}

type Weapon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Weapon) Reset() {
	*x = Weapon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Weapon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Weapon) ProtoMessage() {}

func (x *Weapon) ProtoReflect() protoreflect.Message {
	mi := &file_comm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Weapon.ProtoReflect.Descriptor instead.
func (*Weapon) Descriptor() ([]byte, []int) {
	return file_comm_proto_rawDescGZIP(), []int{2}
}

func (x *Weapon) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type InitialInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId uint32 `protobuf:"varint,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	WeaponId uint32 `protobuf:"varint,2,opt,name=weapon_id,json=weaponId,proto3" json:"weapon_id,omitempty"`
	Seed     int64  `protobuf:"varint,3,opt,name=seed,proto3" json:"seed,omitempty"`
}

func (x *InitialInfo) Reset() {
	*x = InitialInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitialInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitialInfo) ProtoMessage() {}

func (x *InitialInfo) ProtoReflect() protoreflect.Message {
	mi := &file_comm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitialInfo.ProtoReflect.Descriptor instead.
func (*InitialInfo) Descriptor() ([]byte, []int) {
	return file_comm_proto_rawDescGZIP(), []int{3}
}

func (x *InitialInfo) GetPlayerId() uint32 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *InitialInfo) GetWeaponId() uint32 {
	if x != nil {
		return x.WeaponId
	}
	return 0
}

func (x *InitialInfo) GetSeed() int64 {
	if x != nil {
		return x.Seed
	}
	return 0
}

type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Room) Reset() {
	*x = Room{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_comm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_comm_proto_rawDescGZIP(), []int{4}
}

func (x *Room) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Room) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type BytePrefix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bytes uint32 `protobuf:"varint,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (x *BytePrefix) Reset() {
	*x = BytePrefix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comm_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BytePrefix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytePrefix) ProtoMessage() {}

func (x *BytePrefix) ProtoReflect() protoreflect.Message {
	mi := &file_comm_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BytePrefix.ProtoReflect.Descriptor instead.
func (*BytePrefix) Descriptor() ([]byte, []int) {
	return file_comm_proto_rawDescGZIP(), []int{5}
}

func (x *BytePrefix) GetBytes() uint32 {
	if x != nil {
		return x.Bytes
	}
	return 0
}

type StateUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player  *Player      `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	Weapon  *Weapon      `protobuf:"bytes,2,opt,name=weapon,proto3" json:"weapon,omitempty"`
	Variant StateVariant `protobuf:"varint,3,opt,name=variant,proto3,enum=comm.StateVariant" json:"variant,omitempty"`
	Room    *Room        `protobuf:"bytes,4,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *StateUpdate) Reset() {
	*x = StateUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comm_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateUpdate) ProtoMessage() {}

func (x *StateUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_comm_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateUpdate.ProtoReflect.Descriptor instead.
func (*StateUpdate) Descriptor() ([]byte, []int) {
	return file_comm_proto_rawDescGZIP(), []int{6}
}

func (x *StateUpdate) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

func (x *StateUpdate) GetWeapon() *Weapon {
	if x != nil {
		return x.Weapon
	}
	return nil
}

func (x *StateUpdate) GetVariant() StateVariant {
	if x != nil {
		return x.Variant
	}
	return StateVariant_NONE
}

func (x *StateUpdate) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

var File_comm_proto protoreflect.FileDescriptor

var file_comm_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f,
	0x6d, 0x6d, 0x22, 0x96, 0x02, 0x0a, 0x0e, 0x4d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x58, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x59,
	0x12, 0x24, 0x0a, 0x0e, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x5f, 0x70, 0x69, 0x76, 0x6f, 0x74,
	0x5f, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e,
	0x50, 0x69, 0x76, 0x6f, 0x74, 0x58, 0x12, 0x24, 0x0a, 0x0e, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e,
	0x5f, 0x70, 0x69, 0x76, 0x6f, 0x74, 0x5f, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c,
	0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x50, 0x69, 0x76, 0x6f, 0x74, 0x59, 0x12, 0x1c, 0x0a, 0x09,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x74,
	0x74, 0x61, 0x63, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x74, 0x74, 0x61,
	0x63, 0x6b, 0x12, 0x27, 0x0a, 0x09, 0x63, 0x75, 0x72, 0x72, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x52, 0x6f, 0x6f, 0x6d, 0x22, 0x3e, 0x0a, 0x06, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x06, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x57, 0x65, 0x61,
	0x70, 0x6f, 0x6e, 0x52, 0x06, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x22, 0x18, 0x0a, 0x06, 0x57,
	0x65, 0x61, 0x70, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5b, 0x0a, 0x0b, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x65, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x65,
	0x65, 0x64, 0x22, 0x22, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x22, 0x22, 0x0a, 0x0a, 0x42, 0x79, 0x74, 0x65, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x22, 0xa7, 0x01, 0x0a, 0x0b, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x24, 0x0a, 0x06, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x52, 0x06,
	0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x07, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04,
	0x72, 0x6f, 0x6f, 0x6d, 0x2a, 0x4b, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x56, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x10, 0x0a,
	0x0c, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x10, 0x0a, 0x0c, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10,
	0x03, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6b, 0x6d, 0x72, 0x64, 0x2d, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f,
	0x71, 0x6c, 0x70, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_comm_proto_rawDescOnce sync.Once
	file_comm_proto_rawDescData = file_comm_proto_rawDesc
)

func file_comm_proto_rawDescGZIP() []byte {
	file_comm_proto_rawDescOnce.Do(func() {
		file_comm_proto_rawDescData = protoimpl.X.CompressGZIP(file_comm_proto_rawDescData)
	})
	return file_comm_proto_rawDescData
}

var file_comm_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_comm_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_comm_proto_goTypes = []any{
	(StateVariant)(0),      // 0: comm.StateVariant
	(*MovementUpdate)(nil), // 1: comm.MovementUpdate
	(*Player)(nil),         // 2: comm.Player
	(*Weapon)(nil),         // 3: comm.Weapon
	(*InitialInfo)(nil),    // 4: comm.InitialInfo
	(*Room)(nil),           // 5: comm.Room
	(*BytePrefix)(nil),     // 6: comm.BytePrefix
	(*StateUpdate)(nil),    // 7: comm.StateUpdate
}
var file_comm_proto_depIdxs = []int32{
	5, // 0: comm.MovementUpdate.curr_room:type_name -> comm.Room
	3, // 1: comm.Player.weapon:type_name -> comm.Weapon
	2, // 2: comm.StateUpdate.player:type_name -> comm.Player
	3, // 3: comm.StateUpdate.weapon:type_name -> comm.Weapon
	0, // 4: comm.StateUpdate.variant:type_name -> comm.StateVariant
	5, // 5: comm.StateUpdate.room:type_name -> comm.Room
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_comm_proto_init() }
func file_comm_proto_init() {
	if File_comm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comm_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MovementUpdate); i {
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
		file_comm_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Player); i {
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
		file_comm_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Weapon); i {
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
		file_comm_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*InitialInfo); i {
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
		file_comm_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Room); i {
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
		file_comm_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*BytePrefix); i {
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
		file_comm_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*StateUpdate); i {
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
			RawDescriptor: file_comm_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_comm_proto_goTypes,
		DependencyIndexes: file_comm_proto_depIdxs,
		EnumInfos:         file_comm_proto_enumTypes,
		MessageInfos:      file_comm_proto_msgTypes,
	}.Build()
	File_comm_proto = out.File
	file_comm_proto_rawDesc = nil
	file_comm_proto_goTypes = nil
	file_comm_proto_depIdxs = nil
}
