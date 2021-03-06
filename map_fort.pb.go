// Code generated by protoc-gen-go.
// source: map_fort.proto
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

// Ignoring public import of ItemData from inventory_item.proto

// Ignoring public import of ItemAward from inventory_item.proto

// Ignoring public import of ItemId from inventory_item.proto

// Ignoring public import of ItemType from inventory_item.proto

type FortType int32

const (
	FortType_GYM        FortType = 0
	FortType_CHECKPOINT FortType = 1
)

var FortType_name = map[int32]string{
	0: "GYM",
	1: "CHECKPOINT",
}
var FortType_value = map[string]int32{
	"GYM":        0,
	"CHECKPOINT": 1,
}

func (x FortType) String() string {
	return proto.EnumName(FortType_name, int32(x))
}
func (FortType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type FortRenderingType int32

const (
	FortRenderingType_DEFAULT       FortRenderingType = 0
	FortRenderingType_INTERNAL_TEST FortRenderingType = 1
)

var FortRenderingType_name = map[int32]string{
	0: "DEFAULT",
	1: "INTERNAL_TEST",
}
var FortRenderingType_value = map[string]int32{
	"DEFAULT":       0,
	"INTERNAL_TEST": 1,
}

func (x FortRenderingType) String() string {
	return proto.EnumName(FortRenderingType_name, int32(x))
}
func (FortRenderingType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type FortSponsor int32

const (
	FortSponsor_UNSET_SPONSOR FortSponsor = 0
	FortSponsor_MCDONALDS     FortSponsor = 1
	FortSponsor_POKEMON_STORE FortSponsor = 2
)

var FortSponsor_name = map[int32]string{
	0: "UNSET_SPONSOR",
	1: "MCDONALDS",
	2: "POKEMON_STORE",
}
var FortSponsor_value = map[string]int32{
	"UNSET_SPONSOR": 0,
	"MCDONALDS":     1,
	"POKEMON_STORE": 2,
}

func (x FortSponsor) String() string {
	return proto.EnumName(FortSponsor_name, int32(x))
}
func (FortSponsor) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type FortData struct {
	Id                      string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	LastModifiedTimestampMs int64    `protobuf:"varint,2,opt,name=last_modified_timestamp_ms,json=lastModifiedTimestampMs" json:"last_modified_timestamp_ms,omitempty"`
	Latitude                float64  `protobuf:"fixed64,3,opt,name=latitude" json:"latitude,omitempty"`
	Longitude               float64  `protobuf:"fixed64,4,opt,name=longitude" json:"longitude,omitempty"`
	Enabled                 bool     `protobuf:"varint,8,opt,name=enabled" json:"enabled,omitempty"`
	Type                    FortType `protobuf:"varint,9,opt,name=type,enum=POGOProtos.Map.Fort.FortType" json:"type,omitempty"`
	// Team that owns the gym
	OwnedByTeam TeamColor `protobuf:"varint,5,opt,name=owned_by_team,json=ownedByTeam,enum=POGOProtos.Enums.TeamColor" json:"owned_by_team,omitempty"`
	// Highest CP Pokemon at the gym
	GuardPokemonId PokemonId `protobuf:"varint,6,opt,name=guard_pokemon_id,json=guardPokemonId,enum=POGOProtos.Enums.PokemonId" json:"guard_pokemon_id,omitempty"`
	GuardPokemonCp int32     `protobuf:"varint,7,opt,name=guard_pokemon_cp,json=guardPokemonCp" json:"guard_pokemon_cp,omitempty"`
	// Prestigate / experience of the gym
	GymPoints int64 `protobuf:"varint,10,opt,name=gym_points,json=gymPoints" json:"gym_points,omitempty"`
	// Whether someone is battling at the gym currently
	IsInBattle bool `protobuf:"varint,11,opt,name=is_in_battle,json=isInBattle" json:"is_in_battle,omitempty"`
	// Timestamp when the pokestop can be activated again to get items / xp
	CooldownCompleteTimestampMs int64             `protobuf:"varint,14,opt,name=cooldown_complete_timestamp_ms,json=cooldownCompleteTimestampMs" json:"cooldown_complete_timestamp_ms,omitempty"`
	Sponsor                     FortSponsor       `protobuf:"varint,15,opt,name=sponsor,enum=POGOProtos.Map.Fort.FortSponsor" json:"sponsor,omitempty"`
	RenderingType               FortRenderingType `protobuf:"varint,16,opt,name=rendering_type,json=renderingType,enum=POGOProtos.Map.Fort.FortRenderingType" json:"rendering_type,omitempty"`
	// Might represent the type of item applied to the pokestop, right now only lures can be applied
	ActiveFortModifier []byte        `protobuf:"bytes,12,opt,name=active_fort_modifier,json=activeFortModifier,proto3" json:"active_fort_modifier,omitempty"`
	LureInfo           *FortLureInfo `protobuf:"bytes,13,opt,name=lure_info,json=lureInfo" json:"lure_info,omitempty"`
}

func (m *FortData) Reset()                    { *m = FortData{} }
func (m *FortData) String() string            { return proto.CompactTextString(m) }
func (*FortData) ProtoMessage()               {}
func (*FortData) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *FortData) GetLureInfo() *FortLureInfo {
	if m != nil {
		return m.LureInfo
	}
	return nil
}

type FortLureInfo struct {
	FortId                 string    `protobuf:"bytes,1,opt,name=fort_id,json=fortId" json:"fort_id,omitempty"`
	EncounterId            uint64    `protobuf:"fixed64,2,opt,name=encounter_id,json=encounterId" json:"encounter_id,omitempty"`
	ActivePokemonId        PokemonId `protobuf:"varint,3,opt,name=active_pokemon_id,json=activePokemonId,enum=POGOProtos.Enums.PokemonId" json:"active_pokemon_id,omitempty"`
	LureExpiresTimestampMs int64     `protobuf:"varint,4,opt,name=lure_expires_timestamp_ms,json=lureExpiresTimestampMs" json:"lure_expires_timestamp_ms,omitempty"`
}

func (m *FortLureInfo) Reset()                    { *m = FortLureInfo{} }
func (m *FortLureInfo) String() string            { return proto.CompactTextString(m) }
func (*FortLureInfo) ProtoMessage()               {}
func (*FortLureInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type FortModifier struct {
	ItemId                 ItemId `protobuf:"varint,1,opt,name=item_id,json=itemId,enum=POGOProtos.Inventory.Item.ItemId" json:"item_id,omitempty"`
	ExpirationTimestampMs  int64  `protobuf:"varint,2,opt,name=expiration_timestamp_ms,json=expirationTimestampMs" json:"expiration_timestamp_ms,omitempty"`
	DeployerPlayerCodename string `protobuf:"bytes,3,opt,name=deployer_player_codename,json=deployerPlayerCodename" json:"deployer_player_codename,omitempty"`
}

func (m *FortModifier) Reset()                    { *m = FortModifier{} }
func (m *FortModifier) String() string            { return proto.CompactTextString(m) }
func (*FortModifier) ProtoMessage()               {}
func (*FortModifier) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type FortSummary struct {
	FortSummaryId           string  `protobuf:"bytes,1,opt,name=fort_summary_id,json=fortSummaryId" json:"fort_summary_id,omitempty"`
	LastModifiedTimestampMs int64   `protobuf:"varint,2,opt,name=last_modified_timestamp_ms,json=lastModifiedTimestampMs" json:"last_modified_timestamp_ms,omitempty"`
	Latitude                float64 `protobuf:"fixed64,3,opt,name=latitude" json:"latitude,omitempty"`
	Longitude               float64 `protobuf:"fixed64,4,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *FortSummary) Reset()                    { *m = FortSummary{} }
func (m *FortSummary) String() string            { return proto.CompactTextString(m) }
func (*FortSummary) ProtoMessage()               {}
func (*FortSummary) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func init() {
	proto.RegisterType((*FortData)(nil), "POGOProtos.Map.Fort.FortData")
	proto.RegisterType((*FortLureInfo)(nil), "POGOProtos.Map.Fort.FortLureInfo")
	proto.RegisterType((*FortModifier)(nil), "POGOProtos.Map.Fort.FortModifier")
	proto.RegisterType((*FortSummary)(nil), "POGOProtos.Map.Fort.FortSummary")
	proto.RegisterEnum("POGOProtos.Map.Fort.FortType", FortType_name, FortType_value)
	proto.RegisterEnum("POGOProtos.Map.Fort.FortRenderingType", FortRenderingType_name, FortRenderingType_value)
	proto.RegisterEnum("POGOProtos.Map.Fort.FortSponsor", FortSponsor_name, FortSponsor_value)
}

func init() { proto.RegisterFile("map_fort.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 820 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x55, 0x6d, 0x6f, 0xe3, 0x44,
	0x10, 0x3e, 0x37, 0xbd, 0xbc, 0x4c, 0x5e, 0x9a, 0x2e, 0xc7, 0xd5, 0xe4, 0x38, 0xd4, 0xe6, 0xa4,
	0xd3, 0xa9, 0x12, 0x29, 0xb4, 0x12, 0xe2, 0x45, 0x02, 0x35, 0x8e, 0x5b, 0xa2, 0x36, 0xb1, 0xe5,
	0xb8, 0x1f, 0xe0, 0x8b, 0xe5, 0xc4, 0x9b, 0x60, 0xd5, 0xde, 0xb5, 0xec, 0x4d, 0x21, 0xbf, 0x83,
	0xbf, 0xc1, 0x4f, 0xe0, 0x37, 0xf0, 0x9b, 0x18, 0xaf, 0xed, 0xbc, 0x40, 0x23, 0x3e, 0xf2, 0xc5,
	0xde, 0x9d, 0xe7, 0x99, 0xf1, 0xce, 0xcc, 0xb3, 0x63, 0x68, 0x85, 0x6e, 0xe4, 0xcc, 0x79, 0x2c,
	0x7a, 0x51, 0xcc, 0x05, 0x27, 0x1f, 0x99, 0xc6, 0xad, 0x61, 0xa6, 0xcb, 0xa4, 0x37, 0x72, 0xa3,
	0xde, 0x0d, 0x42, 0x9d, 0x3a, 0x65, 0xcb, 0x30, 0xc9, 0x18, 0x9d, 0x57, 0x3e, 0x7b, 0xa2, 0x4c,
	0xf0, 0x78, 0xe5, 0xf8, 0x82, 0x86, 0x99, 0xb5, 0xfb, 0x7b, 0x19, 0xaa, 0x29, 0x77, 0xe0, 0x0a,
	0x97, 0xb4, 0xe0, 0xc0, 0xf7, 0x54, 0xe5, 0x54, 0xf9, 0x50, 0xb3, 0x70, 0x45, 0xbe, 0x83, 0x4e,
	0xe0, 0x26, 0xc2, 0x09, 0xb9, 0xe7, 0xcf, 0x7d, 0xea, 0x39, 0xc2, 0x0f, 0x69, 0x22, 0xdc, 0x30,
	0x72, 0xc2, 0x44, 0x3d, 0x40, 0x5e, 0xc9, 0x3a, 0x49, 0x19, 0xa3, 0x9c, 0x60, 0x17, 0xf8, 0x28,
	0x21, 0x1d, 0xa8, 0x06, 0xae, 0xf0, 0xc5, 0xd2, 0xa3, 0x6a, 0x09, 0xa9, 0x8a, 0xb5, 0xde, 0x93,
	0x4f, 0xa1, 0x16, 0x70, 0xb6, 0xc8, 0xc0, 0x43, 0x09, 0x6e, 0x0c, 0x44, 0x85, 0x0a, 0x65, 0xee,
	0x34, 0xa0, 0x9e, 0x5a, 0x45, 0xac, 0x6a, 0x15, 0x5b, 0xf2, 0x25, 0x1c, 0x8a, 0x55, 0x44, 0xd5,
	0x1a, 0x9a, 0x5b, 0x97, 0x6f, 0x7b, 0xcf, 0x24, 0x2d, 0x1f, 0x36, 0x92, 0x2c, 0x49, 0x25, 0x3f,
	0x40, 0x93, 0xff, 0xca, 0xf0, 0xec, 0xd3, 0x95, 0x23, 0xa8, 0x1b, 0xaa, 0x2f, 0xa5, 0xef, 0x9b,
	0x6d, 0x5f, 0x5d, 0x96, 0xc9, 0x46, 0x54, 0xe3, 0x01, 0x8f, 0xad, 0xba, 0xf4, 0xe8, 0xaf, 0x52,
	0x0b, 0xd1, 0xa1, 0xbd, 0x58, 0xba, 0xb1, 0xe7, 0x44, 0xfc, 0x91, 0x86, 0x9c, 0x39, 0x58, 0xa2,
	0xf2, 0xbe, 0x18, 0x66, 0xc6, 0x19, 0x7a, 0x56, 0x4b, 0x3a, 0xad, 0xf7, 0xe4, 0xc3, 0x3f, 0xc3,
	0xcc, 0x22, 0xb5, 0x82, 0x61, 0x5e, 0xee, 0x32, 0xb5, 0x88, 0xbc, 0x05, 0x58, 0xac, 0x42, 0xe4,
	0xf9, 0x4c, 0x24, 0x2a, 0xc8, 0x2a, 0xd7, 0xd0, 0x62, 0x4a, 0x03, 0x39, 0x85, 0x86, 0x9f, 0x38,
	0x3e, 0x73, 0xa6, 0xae, 0x10, 0x01, 0x55, 0xeb, 0xb2, 0x44, 0xe0, 0x27, 0x43, 0xd6, 0x97, 0x16,
	0xa2, 0xc1, 0x67, 0x33, 0xce, 0x03, 0x0f, 0xb3, 0x70, 0x66, 0x3c, 0x8c, 0x02, 0x2a, 0xe8, 0x6e,
	0xeb, 0x5a, 0x32, 0xe8, 0x9b, 0x82, 0xa5, 0xe5, 0xa4, 0xed, 0xf6, 0x7d, 0x0b, 0x95, 0x24, 0xe2,
	0x2c, 0xe1, 0xb1, 0x7a, 0x24, 0xb3, 0x3d, 0xdd, 0x5b, 0xed, 0x49, 0xc6, 0xb3, 0x0a, 0x07, 0x32,
	0x82, 0x56, 0x4c, 0x99, 0x47, 0x63, 0x9f, 0x2d, 0x1c, 0xd9, 0xb0, 0xb6, 0x0c, 0xf1, 0x7e, 0x6f,
	0x08, 0xab, 0xa0, 0xcb, 0xce, 0x35, 0xe3, 0xed, 0x2d, 0xf9, 0x02, 0x5e, 0xb9, 0x33, 0xe1, 0x3f,
	0x51, 0x29, 0xf8, 0x42, 0x8d, 0xb1, 0xda, 0xc0, 0xa0, 0x0d, 0x8b, 0x64, 0x58, 0x1a, 0x25, 0x97,
	0x61, 0x4c, 0xbe, 0x47, 0x7d, 0x2d, 0x63, 0x8a, 0x55, 0x9a, 0x73, 0xb5, 0x89, 0xb4, 0xfa, 0xe5,
	0xd9, 0xde, 0x6f, 0xdf, 0x23, 0x73, 0x88, 0x44, 0xd4, 0x67, 0xbe, 0xea, 0xfe, 0xa5, 0x40, 0x63,
	0x1b, 0x22, 0x27, 0x50, 0x91, 0xdf, 0x5e, 0x5f, 0x8f, 0x72, 0xba, 0xc5, 0xb6, 0x9e, 0x41, 0x83,
	0xb2, 0x19, 0x5f, 0x32, 0x41, 0xe3, 0x14, 0x4d, 0x2f, 0x45, 0xd9, 0xaa, 0xaf, 0x6d, 0x48, 0xb9,
	0x85, 0xe3, 0xfc, 0xf8, 0x5b, 0x0a, 0x2a, 0xfd, 0xb7, 0x82, 0x8e, 0x32, 0xaf, 0x8d, 0x84, 0xbe,
	0x81, 0x4f, 0x64, 0x56, 0xf4, 0xb7, 0xc8, 0x8f, 0x69, 0xb2, 0xdb, 0xd2, 0x43, 0xd9, 0xd2, 0xd7,
	0x29, 0x41, 0xcf, 0xf0, 0xad, 0x6e, 0x76, 0xff, 0xcc, 0x13, 0x5a, 0x57, 0x08, 0xdb, 0x9b, 0x4e,
	0x81, 0x22, 0xa1, 0xd6, 0x6e, 0x7d, 0x86, 0xc5, 0xa8, 0xe8, 0x0d, 0xd3, 0x51, 0x91, 0x3e, 0xf0,
	0x40, 0x65, 0x5f, 0xbe, 0xc9, 0x57, 0x70, 0x22, 0x8f, 0x80, 0xb7, 0x19, 0x93, 0x79, 0x66, 0x26,
	0x7c, 0xbc, 0x81, 0xb7, 0x25, 0xf5, 0x35, 0xa8, 0x1e, 0x8d, 0x02, 0xbe, 0xc2, 0x52, 0x45, 0x81,
	0x9b, 0xbe, 0x66, 0xdc, 0xc3, 0xbb, 0x1d, 0x66, 0x13, 0xa2, 0x66, 0xbd, 0x2e, 0x70, 0x53, 0xc2,
	0x5a, 0x8e, 0x76, 0xff, 0x50, 0xa0, 0x2e, 0x95, 0xb6, 0x0c, 0x43, 0x37, 0x5e, 0x91, 0xf7, 0x70,
	0x24, 0xdb, 0x91, 0x64, 0xfb, 0x4d, 0x5b, 0x9a, 0xf3, 0x0d, 0x6b, 0xf8, 0x7f, 0x0d, 0xb0, 0xf3,
	0x77, 0xd9, 0x4c, 0x95, 0xe2, 0xad, 0x40, 0xe9, 0xf6, 0xa7, 0x51, 0xfb, 0x05, 0x0e, 0x57, 0xd0,
	0x7e, 0xd4, 0xb5, 0x3b, 0xd3, 0x18, 0x8e, 0xed, 0xb6, 0x72, 0x7e, 0x05, 0xc7, 0xff, 0x52, 0x3e,
	0xa9, 0x43, 0x65, 0xa0, 0xdf, 0x5c, 0x3f, 0xdc, 0xdb, 0xe8, 0x71, 0x0c, 0x4d, 0xa4, 0xea, 0xd6,
	0xf8, 0xfa, 0xde, 0xb1, 0xf5, 0x49, 0xea, 0xd4, 0xcf, 0xeb, 0x90, 0x5f, 0x34, 0x64, 0x3c, 0x8c,
	0x27, 0xba, 0xed, 0x4c, 0x4c, 0x63, 0x3c, 0x31, 0x2c, 0x74, 0x6a, 0x42, 0x6d, 0xa4, 0x0d, 0x0c,
	0xf4, 0x19, 0x4c, 0xda, 0x4a, 0xca, 0x30, 0x8d, 0x3b, 0x7d, 0x64, 0x8c, 0x9d, 0x89, 0x6d, 0x58,
	0x7a, 0xfb, 0xa0, 0xff, 0xee, 0xe7, 0x33, 0x3c, 0xe7, 0x2f, 0xcb, 0x69, 0x0f, 0x87, 0xc3, 0x45,
	0xf4, 0x18, 0xb2, 0x05, 0xff, 0x1c, 0xf3, 0xbf, 0x88, 0x38, 0x2e, 0xe4, 0x6f, 0x21, 0x31, 0x5f,
	0x98, 0xca, 0xb4, 0x2c, 0xd7, 0x57, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x59, 0x48, 0x5e,
	0x6c, 0x06, 0x00, 0x00,
}
