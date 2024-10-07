// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
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
	StateVariant_NONE                   StateVariant = 0
	StateVariant_CONNECTED              StateVariant = 1
	StateVariant_DISCONNECTED           StateVariant = 2
	StateVariant_ROOM_CHANGED           StateVariant = 3
	StateVariant_MAP_UPDATE             StateVariant = 4
	StateVariant_PLAYER_POSITION_UPDATE StateVariant = 5
)

// Enum value maps for StateVariant.
var (
	StateVariant_name = map[int32]string{
		0: "NONE",
		1: "CONNECTED",
		2: "DISCONNECTED",
		3: "ROOM_CHANGED",
		4: "MAP_UPDATE",
		5: "PLAYER_POSITION_UPDATE",
	}
	StateVariant_value = map[string]int32{
		"NONE":                   0,
		"CONNECTED":              1,
		"DISCONNECTED":           2,
		"ROOM_CHANGED":           3,
		"MAP_UPDATE":             4,
		"PLAYER_POSITION_UPDATE": 5,
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
type PositionUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId  uint32  `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	X         float32 `protobuf:"fixed32,2,opt,name=x,proto3" json:"x,omitempty"`
	Y         float32 `protobuf:"fixed32,3,opt,name=y,proto3" json:"y,omitempty"`
	Direction float32 `protobuf:"fixed32,4,opt,name=direction,proto3" json:"direction,omitempty"`
	CurrRoom  *Room   `protobuf:"bytes,5,opt,name=curr_room,json=currRoom,proto3" json:"curr_room,omitempty"`
}

func (x *PositionUpdate) Reset() {
	*x = PositionUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionUpdate) ProtoMessage() {}

func (x *PositionUpdate) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PositionUpdate.ProtoReflect.Descriptor instead.
func (*PositionUpdate) Descriptor() ([]byte, []int) {
	return file_comm_proto_rawDescGZIP(), []int{0}
}

func (x *PositionUpdate) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *PositionUpdate) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *PositionUpdate) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *PositionUpdate) GetDirection() float32 {
	if x != nil {
		return x.Direction
	}
	return 0
}

func (x *PositionUpdate) GetCurrRoom() *Room {
	if x != nil {
		return x.CurrRoom
	}
	return nil
}

// sent via TCP
type GameState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId         uint32   `protobuf:"varint,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Seed             int64    `protobuf:"varint,2,opt,name=seed,proto3" json:"seed,omitempty"`
	ConnectedPlayers []uint32 `protobuf:"varint,3,rep,packed,name=connected_players,json=connectedPlayers,proto3" json:"connected_players,omitempty"`
}

func (x *GameState) Reset() {
	*x = GameState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameState) ProtoMessage() {}

func (x *GameState) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GameState.ProtoReflect.Descriptor instead.
func (*GameState) Descriptor() ([]byte, []int) {
	return file_comm_proto_rawDescGZIP(), []int{1}
}

func (x *GameState) GetPlayerId() uint32 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *GameState) GetSeed() int64 {
	if x != nil {
		return x.Seed
	}
	return 0
}

func (x *GameState) GetConnectedPlayers() []uint32 {
	if x != nil {
		return x.ConnectedPlayers
	}
	return nil
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
		mi := &file_comm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_comm_proto_rawDescGZIP(), []int{2}
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

type Obstacle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height int32 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Width  int32 `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Left   int32 `protobuf:"varint,3,opt,name=left,proto3" json:"left,omitempty"`
	Top    int32 `protobuf:"varint,4,opt,name=top,proto3" json:"top,omitempty"`
}

func (x *Obstacle) Reset() {
	*x = Obstacle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Obstacle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Obstacle) ProtoMessage() {}

func (x *Obstacle) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Obstacle.ProtoReflect.Descriptor instead.
func (*Obstacle) Descriptor() ([]byte, []int) {
	return file_comm_proto_rawDescGZIP(), []int{3}
}

func (x *Obstacle) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Obstacle) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Obstacle) GetLeft() int32 {
	if x != nil {
		return x.Left
	}
	return 0
}

func (x *Obstacle) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

type Enemy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	X  float32 `protobuf:"fixed32,2,opt,name=x,proto3" json:"x,omitempty"`
	Y  float32 `protobuf:"fixed32,3,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Enemy) Reset() {
	*x = Enemy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Enemy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Enemy) ProtoMessage() {}

func (x *Enemy) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Enemy.ProtoReflect.Descriptor instead.
func (*Enemy) Descriptor() ([]byte, []int) {
	return file_comm_proto_rawDescGZIP(), []int{4}
}

func (x *Enemy) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Enemy) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Enemy) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	X  int32  `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	Y  int32  `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comm_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_comm_proto_rawDescGZIP(), []int{5}
}

func (x *Player) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Player) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Player) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type EnemyPositionsUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnemyPositions []*Enemy `protobuf:"bytes,1,rep,name=enemyPositions,proto3" json:"enemyPositions,omitempty"`
}

func (x *EnemyPositionsUpdate) Reset() {
	*x = EnemyPositionsUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comm_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnemyPositionsUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnemyPositionsUpdate) ProtoMessage() {}

func (x *EnemyPositionsUpdate) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use EnemyPositionsUpdate.ProtoReflect.Descriptor instead.
func (*EnemyPositionsUpdate) Descriptor() ([]byte, []int) {
	return file_comm_proto_rawDescGZIP(), []int{6}
}

func (x *EnemyPositionsUpdate) GetEnemyPositions() []*Enemy {
	if x != nil {
		return x.EnemyPositions
	}
	return nil
}

type MapPositionsUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players []*Player `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
	Enemies []*Enemy  `protobuf:"bytes,2,rep,name=enemies,proto3" json:"enemies,omitempty"`
}

func (x *MapPositionsUpdate) Reset() {
	*x = MapPositionsUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comm_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapPositionsUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapPositionsUpdate) ProtoMessage() {}

func (x *MapPositionsUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_comm_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapPositionsUpdate.ProtoReflect.Descriptor instead.
func (*MapPositionsUpdate) Descriptor() ([]byte, []int) {
	return file_comm_proto_rawDescGZIP(), []int{7}
}

func (x *MapPositionsUpdate) GetPlayers() []*Player {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *MapPositionsUpdate) GetEnemies() []*Enemy {
	if x != nil {
		return x.Enemies
	}
	return nil
}

type MapDimensionsUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Obstacles []*Obstacle `protobuf:"bytes,1,rep,name=obstacles,proto3" json:"obstacles,omitempty"`
}

func (x *MapDimensionsUpdate) Reset() {
	*x = MapDimensionsUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comm_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapDimensionsUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapDimensionsUpdate) ProtoMessage() {}

func (x *MapDimensionsUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_comm_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapDimensionsUpdate.ProtoReflect.Descriptor instead.
func (*MapDimensionsUpdate) Descriptor() ([]byte, []int) {
	return file_comm_proto_rawDescGZIP(), []int{8}
}

func (x *MapDimensionsUpdate) GetObstacles() []*Obstacle {
	if x != nil {
		return x.Obstacles
	}
	return nil
}

type StateUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   uint32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Variant              StateVariant          `protobuf:"varint,2,opt,name=variant,proto3,enum=comm.StateVariant" json:"variant,omitempty"`
	Room                 *Room                 `protobuf:"bytes,3,opt,name=room,proto3" json:"room,omitempty"`
	MapPositionsUpdate   *MapPositionsUpdate   `protobuf:"bytes,4,opt,name=mapPositionsUpdate,proto3" json:"mapPositionsUpdate,omitempty"`
	PositionUpdate       *PositionUpdate       `protobuf:"bytes,5,opt,name=positionUpdate,proto3" json:"positionUpdate,omitempty"`
	EnemyPositionsUpdate *EnemyPositionsUpdate `protobuf:"bytes,6,opt,name=enemyPositionsUpdate,proto3" json:"enemyPositionsUpdate,omitempty"`
	MapDimensionsUpdate  *MapDimensionsUpdate  `protobuf:"bytes,7,opt,name=mapDimensionsUpdate,proto3" json:"mapDimensionsUpdate,omitempty"`
}

func (x *StateUpdate) Reset() {
	*x = StateUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comm_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateUpdate) ProtoMessage() {}

func (x *StateUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_comm_proto_msgTypes[9]
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
	return file_comm_proto_rawDescGZIP(), []int{9}
}

func (x *StateUpdate) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
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

func (x *StateUpdate) GetMapPositionsUpdate() *MapPositionsUpdate {
	if x != nil {
		return x.MapPositionsUpdate
	}
	return nil
}

func (x *StateUpdate) GetPositionUpdate() *PositionUpdate {
	if x != nil {
		return x.PositionUpdate
	}
	return nil
}

func (x *StateUpdate) GetEnemyPositionsUpdate() *EnemyPositionsUpdate {
	if x != nil {
		return x.EnemyPositionsUpdate
	}
	return nil
}

func (x *StateUpdate) GetMapDimensionsUpdate() *MapDimensionsUpdate {
	if x != nil {
		return x.MapDimensionsUpdate
	}
	return nil
}

var File_comm_proto protoreflect.FileDescriptor

var file_comm_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f,
	0x6d, 0x6d, 0x22, 0x90, 0x01, 0x0a, 0x0e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78,
	0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x09,
	0x63, 0x75, 0x72, 0x72, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x52, 0x6f, 0x6f, 0x6d, 0x22, 0x69, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x65, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73,
	0x65, 0x65, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x10,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x22, 0x22, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x01, 0x79, 0x22, 0x5e, 0x0a, 0x08, 0x4f, 0x62, 0x73, 0x74, 0x61, 0x63, 0x6c, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x65,
	0x66, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x74, 0x6f, 0x70, 0x22, 0x33, 0x0a, 0x05, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0c, 0x0a,
	0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x22, 0x34, 0x0a, 0x06, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01,
	0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x22,
	0x4b, 0x0a, 0x14, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x33, 0x0a, 0x0e, 0x65, 0x6e, 0x65, 0x6d, 0x79,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x52, 0x0e, 0x65, 0x6e,
	0x65, 0x6d, 0x79, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x63, 0x0a, 0x12,
	0x4d, 0x61, 0x70, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x26, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x07, 0x65, 0x6e,
	0x65, 0x6d, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x2e, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x65, 0x6d, 0x69, 0x65,
	0x73, 0x22, 0x43, 0x0a, 0x13, 0x4d, 0x61, 0x70, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x09, 0x6f, 0x62, 0x73, 0x74,
	0x61, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x2e, 0x4f, 0x62, 0x73, 0x74, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x09, 0x6f, 0x62, 0x73,
	0x74, 0x61, 0x63, 0x6c, 0x65, 0x73, 0x22, 0x90, 0x03, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x07, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04,
	0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x48, 0x0a, 0x12, 0x6d, 0x61, 0x70, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x4d, 0x61, 0x70, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x12, 0x6d, 0x61, 0x70, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3c,
	0x0a, 0x0e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x0e, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x4e, 0x0a, 0x14,
	0x65, 0x6e, 0x65, 0x6d, 0x79, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x2e, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x14, 0x65, 0x6e, 0x65, 0x6d, 0x79, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x4b, 0x0a, 0x13,
	0x6d, 0x61, 0x70, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x2e, 0x4d, 0x61, 0x70, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x13, 0x6d, 0x61, 0x70, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2a, 0x77, 0x0a, 0x0c, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44,
	0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54,
	0x45, 0x44, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x43, 0x48, 0x41,
	0x4e, 0x47, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x41, 0x50, 0x5f, 0x55, 0x50,
	0x44, 0x41, 0x54, 0x45, 0x10, 0x04, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52,
	0x5f, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45,
	0x10, 0x05, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x6d, 0x72, 0x64, 0x2d, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x2f, 0x71, 0x6c, 0x70, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x62, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
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
var file_comm_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_comm_proto_goTypes = []any{
	(StateVariant)(0),            // 0: comm.StateVariant
	(*PositionUpdate)(nil),       // 1: comm.PositionUpdate
	(*GameState)(nil),            // 2: comm.GameState
	(*Room)(nil),                 // 3: comm.Room
	(*Obstacle)(nil),             // 4: comm.Obstacle
	(*Enemy)(nil),                // 5: comm.Enemy
	(*Player)(nil),               // 6: comm.Player
	(*EnemyPositionsUpdate)(nil), // 7: comm.EnemyPositionsUpdate
	(*MapPositionsUpdate)(nil),   // 8: comm.MapPositionsUpdate
	(*MapDimensionsUpdate)(nil),  // 9: comm.MapDimensionsUpdate
	(*StateUpdate)(nil),          // 10: comm.StateUpdate
}
var file_comm_proto_depIdxs = []int32{
	3,  // 0: comm.PositionUpdate.curr_room:type_name -> comm.Room
	5,  // 1: comm.EnemyPositionsUpdate.enemyPositions:type_name -> comm.Enemy
	6,  // 2: comm.MapPositionsUpdate.players:type_name -> comm.Player
	5,  // 3: comm.MapPositionsUpdate.enemies:type_name -> comm.Enemy
	4,  // 4: comm.MapDimensionsUpdate.obstacles:type_name -> comm.Obstacle
	0,  // 5: comm.StateUpdate.variant:type_name -> comm.StateVariant
	3,  // 6: comm.StateUpdate.room:type_name -> comm.Room
	8,  // 7: comm.StateUpdate.mapPositionsUpdate:type_name -> comm.MapPositionsUpdate
	1,  // 8: comm.StateUpdate.positionUpdate:type_name -> comm.PositionUpdate
	7,  // 9: comm.StateUpdate.enemyPositionsUpdate:type_name -> comm.EnemyPositionsUpdate
	9,  // 10: comm.StateUpdate.mapDimensionsUpdate:type_name -> comm.MapDimensionsUpdate
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_comm_proto_init() }
func file_comm_proto_init() {
	if File_comm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comm_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PositionUpdate); i {
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
			switch v := v.(*GameState); i {
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
		file_comm_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Obstacle); i {
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
			switch v := v.(*Enemy); i {
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
		file_comm_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*EnemyPositionsUpdate); i {
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
		file_comm_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*MapPositionsUpdate); i {
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
		file_comm_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*MapDimensionsUpdate); i {
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
		file_comm_proto_msgTypes[9].Exporter = func(v any, i int) any {
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
			NumMessages:   10,
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
