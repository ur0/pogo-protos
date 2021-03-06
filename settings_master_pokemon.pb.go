// Code generated by protoc-gen-go.
// source: settings_master_pokemon.proto
// DO NOT EDIT!

package pogo_protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of PokemonMovementType from enums.proto

// Ignoring public import of HoloIapItemCategory from enums.proto

// Ignoring public import of PokemonMove from enums.proto

// Ignoring public import of ItemCategory from enums.proto

// Ignoring public import of CameraInterpolation from enums.proto

// Ignoring public import of TutorialState from enums.proto

// Ignoring public import of ItemEffect from enums.proto

// Ignoring public import of Platform from enums.proto

// Ignoring public import of BadgeType from enums.proto

// Ignoring public import of PokemonRarity from enums.proto

// Ignoring public import of PokemonFamilyId from enums.proto

// Ignoring public import of PokemonType from enums.proto

// Ignoring public import of TeamColor from enums.proto

// Ignoring public import of ActivityType from enums.proto

// Ignoring public import of CameraTarget from enums.proto

// Ignoring public import of Gender from enums.proto

// Ignoring public import of PokemonId from enums.proto

type StatsAttributes struct {
	BaseStamina      int32 `protobuf:"varint,1,opt,name=base_stamina,json=baseStamina" json:"base_stamina,omitempty"`
	BaseAttack       int32 `protobuf:"varint,2,opt,name=base_attack,json=baseAttack" json:"base_attack,omitempty"`
	BaseDefense      int32 `protobuf:"varint,3,opt,name=base_defense,json=baseDefense" json:"base_defense,omitempty"`
	DodgeEnergyDelta int32 `protobuf:"varint,8,opt,name=dodge_energy_delta,json=dodgeEnergyDelta" json:"dodge_energy_delta,omitempty"`
}

func (m *StatsAttributes) Reset()                    { *m = StatsAttributes{} }
func (m *StatsAttributes) String() string            { return proto.CompactTextString(m) }
func (*StatsAttributes) ProtoMessage()               {}
func (*StatsAttributes) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

type CameraAttributes struct {
	DiskRadiusM       float32 `protobuf:"fixed32,1,opt,name=disk_radius_m,json=diskRadiusM" json:"disk_radius_m,omitempty"`
	CylinderRadiusM   float32 `protobuf:"fixed32,2,opt,name=cylinder_radius_m,json=cylinderRadiusM" json:"cylinder_radius_m,omitempty"`
	CylinderHeightM   float32 `protobuf:"fixed32,3,opt,name=cylinder_height_m,json=cylinderHeightM" json:"cylinder_height_m,omitempty"`
	CylinderGroundM   float32 `protobuf:"fixed32,4,opt,name=cylinder_ground_m,json=cylinderGroundM" json:"cylinder_ground_m,omitempty"`
	ShoulderModeScale float32 `protobuf:"fixed32,5,opt,name=shoulder_mode_scale,json=shoulderModeScale" json:"shoulder_mode_scale,omitempty"`
}

func (m *CameraAttributes) Reset()                    { *m = CameraAttributes{} }
func (m *CameraAttributes) String() string            { return proto.CompactTextString(m) }
func (*CameraAttributes) ProtoMessage()               {}
func (*CameraAttributes) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{1} }

type EncounterAttributes struct {
	BaseCaptureRate      float32             `protobuf:"fixed32,1,opt,name=base_capture_rate,json=baseCaptureRate" json:"base_capture_rate,omitempty"`
	BaseFleeRate         float32             `protobuf:"fixed32,2,opt,name=base_flee_rate,json=baseFleeRate" json:"base_flee_rate,omitempty"`
	CollisionRadiusM     float32             `protobuf:"fixed32,3,opt,name=collision_radius_m,json=collisionRadiusM" json:"collision_radius_m,omitempty"`
	CollisionHeightM     float32             `protobuf:"fixed32,4,opt,name=collision_height_m,json=collisionHeightM" json:"collision_height_m,omitempty"`
	CollisionHeadRadiusM float32             `protobuf:"fixed32,5,opt,name=collision_head_radius_m,json=collisionHeadRadiusM" json:"collision_head_radius_m,omitempty"`
	MovementType         PokemonMovementType `protobuf:"varint,6,opt,name=movement_type,json=movementType,enum=POGOProtos.Enums.PokemonMovementType" json:"movement_type,omitempty"`
	MovementTimerS       float32             `protobuf:"fixed32,7,opt,name=movement_timer_s,json=movementTimerS" json:"movement_timer_s,omitempty"`
	JumpTimeS            float32             `protobuf:"fixed32,8,opt,name=jump_time_s,json=jumpTimeS" json:"jump_time_s,omitempty"`
	AttackTimerS         float32             `protobuf:"fixed32,9,opt,name=attack_timer_s,json=attackTimerS" json:"attack_timer_s,omitempty"`
}

func (m *EncounterAttributes) Reset()                    { *m = EncounterAttributes{} }
func (m *EncounterAttributes) String() string            { return proto.CompactTextString(m) }
func (*EncounterAttributes) ProtoMessage()               {}
func (*EncounterAttributes) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{2} }

func init() {
	proto.RegisterType((*StatsAttributes)(nil), "POGOProtos.Settings.Master.Pokemon.StatsAttributes")
	proto.RegisterType((*CameraAttributes)(nil), "POGOProtos.Settings.Master.Pokemon.CameraAttributes")
	proto.RegisterType((*EncounterAttributes)(nil), "POGOProtos.Settings.Master.Pokemon.EncounterAttributes")
}

func init() { proto.RegisterFile("settings_master_pokemon.proto", fileDescriptor16) }

var fileDescriptor16 = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x93, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x69, 0xbb, 0x8d, 0xcd, 0xdd, 0xda, 0x2e, 0x43, 0x22, 0x42, 0xe2, 0xcf, 0x0a, 0x48,
	0x13, 0x82, 0x4c, 0x02, 0xf1, 0x00, 0x63, 0x2b, 0x43, 0x48, 0xd5, 0xaa, 0x86, 0x2b, 0x6e, 0x22,
	0x37, 0x3e, 0x4b, 0x4d, 0x63, 0x3b, 0x8a, 0x1d, 0xa4, 0xbe, 0x10, 0x8f, 0xc0, 0x43, 0xf1, 0x14,
	0xd8, 0xc7, 0x4d, 0x9a, 0x22, 0xee, 0x9c, 0xef, 0xfb, 0xf9, 0xd8, 0xe7, 0xf3, 0x09, 0x79, 0xaa,
	0xc1, 0x18, 0x2e, 0x33, 0x9d, 0x08, 0xaa, 0x0d, 0x94, 0x49, 0xa1, 0x56, 0x20, 0x94, 0x8c, 0x8a,
	0x52, 0x19, 0x15, 0x8c, 0x67, 0x77, 0xb7, 0x77, 0x33, 0xb7, 0xd4, 0x51, 0xbc, 0x21, 0xa3, 0x29,
	0x92, 0xd1, 0xcc, 0x93, 0x4f, 0xfa, 0x20, 0x2b, 0xa1, 0xfd, 0x86, 0xf1, 0xaf, 0x0e, 0x19, 0xc6,
	0x86, 0x1a, 0x7d, 0x65, 0x4c, 0xc9, 0x17, 0x95, 0x01, 0x1d, 0x9c, 0x93, 0xe3, 0x05, 0xd5, 0x90,
	0x68, 0x43, 0x05, 0x97, 0x34, 0xec, 0xbc, 0xe8, 0x5c, 0xec, 0xcf, 0xfb, 0x4e, 0x8b, 0xbd, 0x14,
	0x3c, 0x27, 0xf8, 0x99, 0x50, 0x63, 0x68, 0xba, 0x0a, 0xbb, 0x48, 0x10, 0x27, 0x5d, 0xa1, 0xd2,
	0xd4, 0x60, 0x70, 0x0f, 0x52, 0x43, 0xd8, 0xdb, 0xd6, 0xb8, 0xf1, 0x52, 0xf0, 0x96, 0x04, 0x4c,
	0xb1, 0x0c, 0x12, 0x90, 0x50, 0x66, 0x6b, 0x8b, 0xe6, 0x86, 0x86, 0x87, 0x08, 0x8e, 0xd0, 0x99,
	0xa0, 0x71, 0xe3, 0xf4, 0xf1, 0x9f, 0x0e, 0x19, 0x5d, 0x53, 0x01, 0x25, 0x6d, 0xdd, 0x74, 0x4c,
	0x4e, 0x18, 0xd7, 0xab, 0xa4, 0xa4, 0x8c, 0x57, 0x36, 0x12, 0xbc, 0x6a, 0x77, 0xde, 0x77, 0xe2,
	0x1c, 0xb5, 0x69, 0xf0, 0x86, 0x9c, 0xa6, 0xeb, 0x9c, 0x4b, 0x66, 0xc3, 0x6a, 0xb8, 0x2e, 0x72,
	0xc3, 0xda, 0xf8, 0x1f, 0xbb, 0x04, 0x9e, 0x2d, 0x8d, 0x65, 0x7b, 0xbb, 0xec, 0x17, 0xd4, 0x77,
	0xd9, 0xac, 0x54, 0x95, 0x64, 0x96, 0xdd, 0xdb, 0x65, 0x6f, 0x51, 0x9f, 0x06, 0x11, 0x39, 0xd3,
	0x4b, 0x55, 0xe5, 0x8e, 0x15, 0x8a, 0xd9, 0x68, 0x53, 0x9a, 0x43, 0xb8, 0x8f, 0xf4, 0x69, 0x6d,
	0x4d, 0xad, 0x13, 0x3b, 0x63, 0xfc, 0xbb, 0x47, 0xce, 0x26, 0x32, 0xb5, 0x9b, 0xed, 0xc3, 0xb5,
	0xfa, 0xb5, 0x67, 0x62, 0xaa, 0x29, 0x2d, 0x4c, 0x55, 0x82, 0xed, 0xc7, 0xc0, 0xa6, 0xe7, 0xa1,
	0x33, 0xae, 0xbd, 0x3e, 0xb7, 0x72, 0xf0, 0x8a, 0x0c, 0x90, 0xbd, 0xcf, 0x61, 0x03, 0xfa, 0xa6,
	0xf1, 0x5d, 0x3e, 0x5b, 0x11, 0x29, 0xfb, 0x08, 0xa9, 0xca, 0x73, 0xae, 0xb9, 0x92, 0xdb, 0x78,
	0x7c, 0xcb, 0xa3, 0xc6, 0xa9, 0xf3, 0xd9, 0xa1, 0x9b, 0x80, 0xf6, 0xfe, 0xa1, 0xeb, 0x84, 0x3e,
	0x92, 0xc7, 0x6d, 0x9a, 0xb2, 0xed, 0x01, 0xbe, 0xf3, 0x47, 0xad, 0x2d, 0x94, 0xd5, 0x87, 0x7c,
	0x25, 0x27, 0x42, 0xfd, 0x04, 0x01, 0xd2, 0x24, 0x66, 0x5d, 0x40, 0x78, 0x60, 0xe1, 0xc1, 0xfb,
	0xd7, 0x51, 0x6b, 0xb6, 0x27, 0x38, 0xc2, 0x9b, 0x89, 0x9e, 0x6e, 0xe8, 0x6f, 0x16, 0x9e, 0x1f,
	0x8b, 0xd6, 0x57, 0x70, 0x41, 0x46, 0xdb, 0x5a, 0xdc, 0x4e, 0x4f, 0xa2, 0xc3, 0x87, 0x78, 0xf6,
	0xa0, 0xe1, 0x9c, 0x1c, 0x07, 0xcf, 0x48, 0xff, 0x47, 0x25, 0x0a, 0xa4, 0x2c, 0x74, 0x88, 0xd0,
	0x91, 0x93, 0x1c, 0x10, 0xbb, 0x38, 0xfd, 0xb0, 0x37, 0x75, 0x8e, 0x7c, 0x9c, 0x5e, 0xf5, 0x55,
	0x3e, 0xbd, 0xfc, 0x7e, 0x9e, 0x71, 0xb3, 0xac, 0x16, 0x51, 0xaa, 0xc4, 0x65, 0xb1, 0x12, 0x32,
	0x53, 0xef, 0x14, 0xe3, 0x97, 0x85, 0xb2, 0x0b, 0xfc, 0xe5, 0xf4, 0xec, 0xc1, 0xe2, 0x00, 0x57,
	0x1f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xbd, 0x08, 0x34, 0xce, 0x03, 0x00, 0x00,
}
