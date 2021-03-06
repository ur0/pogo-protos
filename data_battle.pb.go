// Code generated by protoc-gen-go.
// source: data_battle.proto
// DO NOT EDIT!

package pogo_protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of AssetDigestEntry from data.proto

// Ignoring public import of PokemonData from data.proto

// Ignoring public import of PokedexEntry from data.proto

// Ignoring public import of DownloadUrlEntry from data.proto

// Ignoring public import of PlayerBadge from data.proto

// Ignoring public import of PlayerData from data.proto

// Ignoring public import of DailyBonus from data_player.proto

// Ignoring public import of ContactSettings from data_player.proto

// Ignoring public import of EquippedBadge from data_player.proto

// Ignoring public import of Currency from data_player.proto

// Ignoring public import of PlayerPublicProfile from data_player.proto

// Ignoring public import of PlayerStats from data_player.proto

// Ignoring public import of PlayerCamera from data_player.proto

// Ignoring public import of PlayerAvatar from data_player.proto

// Ignoring public import of PlayerCurrency from data_player.proto

// Ignoring public import of GymMembership from data_gym.proto

// Ignoring public import of GymState from data_gym.proto

type BattleType int32

const (
	BattleType_BATTLE_TYPE_UNSET BattleType = 0
	BattleType_NORMAL            BattleType = 1
	BattleType_TRAINING          BattleType = 2
)

var BattleType_name = map[int32]string{
	0: "BATTLE_TYPE_UNSET",
	1: "NORMAL",
	2: "TRAINING",
}
var BattleType_value = map[string]int32{
	"BATTLE_TYPE_UNSET": 0,
	"NORMAL":            1,
	"TRAINING":          2,
}

func (x BattleType) String() string {
	return proto.EnumName(BattleType_name, int32(x))
}
func (BattleType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type BattleActionType int32

const (
	BattleActionType_ACTION_UNSET          BattleActionType = 0
	BattleActionType_ACTION_ATTACK         BattleActionType = 1
	BattleActionType_ACTION_DODGE          BattleActionType = 2
	BattleActionType_ACTION_SPECIAL_ATTACK BattleActionType = 3
	BattleActionType_ACTION_SWAP_POKEMON   BattleActionType = 4
	BattleActionType_ACTION_FAINT          BattleActionType = 5
	BattleActionType_ACTION_PLAYER_JOIN    BattleActionType = 6
	BattleActionType_ACTION_PLAYER_QUIT    BattleActionType = 7
	BattleActionType_ACTION_VICTORY        BattleActionType = 8
	BattleActionType_ACTION_DEFEAT         BattleActionType = 9
	BattleActionType_ACTION_TIMED_OUT      BattleActionType = 10
)

var BattleActionType_name = map[int32]string{
	0:  "ACTION_UNSET",
	1:  "ACTION_ATTACK",
	2:  "ACTION_DODGE",
	3:  "ACTION_SPECIAL_ATTACK",
	4:  "ACTION_SWAP_POKEMON",
	5:  "ACTION_FAINT",
	6:  "ACTION_PLAYER_JOIN",
	7:  "ACTION_PLAYER_QUIT",
	8:  "ACTION_VICTORY",
	9:  "ACTION_DEFEAT",
	10: "ACTION_TIMED_OUT",
}
var BattleActionType_value = map[string]int32{
	"ACTION_UNSET":          0,
	"ACTION_ATTACK":         1,
	"ACTION_DODGE":          2,
	"ACTION_SPECIAL_ATTACK": 3,
	"ACTION_SWAP_POKEMON":   4,
	"ACTION_FAINT":          5,
	"ACTION_PLAYER_JOIN":    6,
	"ACTION_PLAYER_QUIT":    7,
	"ACTION_VICTORY":        8,
	"ACTION_DEFEAT":         9,
	"ACTION_TIMED_OUT":      10,
}

func (x BattleActionType) String() string {
	return proto.EnumName(BattleActionType_name, int32(x))
}
func (BattleActionType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type BattleState int32

const (
	BattleState_STATE_UNSET BattleState = 0
	BattleState_ACTIVE      BattleState = 1
	BattleState_VICTORY     BattleState = 2
	BattleState_DEFEATED    BattleState = 3
	BattleState_TIMED_OUT   BattleState = 4
)

var BattleState_name = map[int32]string{
	0: "STATE_UNSET",
	1: "ACTIVE",
	2: "VICTORY",
	3: "DEFEATED",
	4: "TIMED_OUT",
}
var BattleState_value = map[string]int32{
	"STATE_UNSET": 0,
	"ACTIVE":      1,
	"VICTORY":     2,
	"DEFEATED":    3,
	"TIMED_OUT":   4,
}

func (x BattleState) String() string {
	return proto.EnumName(BattleState_name, int32(x))
}
func (BattleState) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

type BattleLog struct {
	State                  BattleState     `protobuf:"varint,1,opt,name=state,enum=POGOProtos.Data.Battle.BattleState" json:"state,omitempty"`
	BattleType             BattleType      `protobuf:"varint,2,opt,name=battle_type,json=battleType,enum=POGOProtos.Data.Battle.BattleType" json:"battle_type,omitempty"`
	ServerMs               int64           `protobuf:"varint,3,opt,name=server_ms,json=serverMs" json:"server_ms,omitempty"`
	BattleActions          []*BattleAction `protobuf:"bytes,4,rep,name=battle_actions,json=battleActions" json:"battle_actions,omitempty"`
	BattleStartTimestampMs int64           `protobuf:"varint,5,opt,name=battle_start_timestamp_ms,json=battleStartTimestampMs" json:"battle_start_timestamp_ms,omitempty"`
	BattleEndTimestampMs   int64           `protobuf:"varint,6,opt,name=battle_end_timestamp_ms,json=battleEndTimestampMs" json:"battle_end_timestamp_ms,omitempty"`
}

func (m *BattleLog) Reset()                    { *m = BattleLog{} }
func (m *BattleLog) String() string            { return proto.CompactTextString(m) }
func (*BattleLog) ProtoMessage()               {}
func (*BattleLog) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *BattleLog) GetBattleActions() []*BattleAction {
	if m != nil {
		return m.BattleActions
	}
	return nil
}

type BattlePokemonInfo struct {
	PokemonData   *PokemonData `protobuf:"bytes,1,opt,name=pokemon_data,json=pokemonData" json:"pokemon_data,omitempty"`
	CurrentHealth int32        `protobuf:"varint,2,opt,name=current_health,json=currentHealth" json:"current_health,omitempty"`
	CurrentEnergy int32        `protobuf:"varint,3,opt,name=current_energy,json=currentEnergy" json:"current_energy,omitempty"`
}

func (m *BattlePokemonInfo) Reset()                    { *m = BattlePokemonInfo{} }
func (m *BattlePokemonInfo) String() string            { return proto.CompactTextString(m) }
func (*BattlePokemonInfo) ProtoMessage()               {}
func (*BattlePokemonInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *BattlePokemonInfo) GetPokemonData() *PokemonData {
	if m != nil {
		return m.PokemonData
	}
	return nil
}

type BattleParticipant struct {
	ActivePokemon        *BattlePokemonInfo   `protobuf:"bytes,1,opt,name=active_pokemon,json=activePokemon" json:"active_pokemon,omitempty"`
	TrainerPublicProfile *PlayerPublicProfile `protobuf:"bytes,2,opt,name=trainer_public_profile,json=trainerPublicProfile" json:"trainer_public_profile,omitempty"`
	ReversePokemon       []*BattlePokemonInfo `protobuf:"bytes,3,rep,name=reverse_pokemon,json=reversePokemon" json:"reverse_pokemon,omitempty"`
	DefeatedPokemon      []*BattlePokemonInfo `protobuf:"bytes,4,rep,name=defeated_pokemon,json=defeatedPokemon" json:"defeated_pokemon,omitempty"`
}

func (m *BattleParticipant) Reset()                    { *m = BattleParticipant{} }
func (m *BattleParticipant) String() string            { return proto.CompactTextString(m) }
func (*BattleParticipant) ProtoMessage()               {}
func (*BattleParticipant) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *BattleParticipant) GetActivePokemon() *BattlePokemonInfo {
	if m != nil {
		return m.ActivePokemon
	}
	return nil
}

func (m *BattleParticipant) GetTrainerPublicProfile() *PlayerPublicProfile {
	if m != nil {
		return m.TrainerPublicProfile
	}
	return nil
}

func (m *BattleParticipant) GetReversePokemon() []*BattlePokemonInfo {
	if m != nil {
		return m.ReversePokemon
	}
	return nil
}

func (m *BattleParticipant) GetDefeatedPokemon() []*BattlePokemonInfo {
	if m != nil {
		return m.DefeatedPokemon
	}
	return nil
}

type BattleResults struct {
	GymState                *GymState            `protobuf:"bytes,1,opt,name=gym_state,json=gymState" json:"gym_state,omitempty"`
	Attackers               []*BattleParticipant `protobuf:"bytes,2,rep,name=attackers" json:"attackers,omitempty"`
	PlayerExperienceAwarded []int32              `protobuf:"varint,3,rep,name=player_experience_awarded,json=playerExperienceAwarded" json:"player_experience_awarded,omitempty"`
	NextDefenderPokemonId   int64                `protobuf:"varint,4,opt,name=next_defender_pokemon_id,json=nextDefenderPokemonId" json:"next_defender_pokemon_id,omitempty"`
	GymPointsDelta          int32                `protobuf:"varint,5,opt,name=gym_points_delta,json=gymPointsDelta" json:"gym_points_delta,omitempty"`
}

func (m *BattleResults) Reset()                    { *m = BattleResults{} }
func (m *BattleResults) String() string            { return proto.CompactTextString(m) }
func (*BattleResults) ProtoMessage()               {}
func (*BattleResults) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *BattleResults) GetGymState() *GymState {
	if m != nil {
		return m.GymState
	}
	return nil
}

func (m *BattleResults) GetAttackers() []*BattleParticipant {
	if m != nil {
		return m.Attackers
	}
	return nil
}

type BattleAction struct {
	Type                           BattleActionType   `protobuf:"varint,1,opt,name=Type,json=type,enum=POGOProtos.Data.Battle.BattleActionType" json:"Type,omitempty"`
	ActionStartMs                  int64              `protobuf:"varint,2,opt,name=action_start_ms,json=actionStartMs" json:"action_start_ms,omitempty"`
	DurationMs                     int32              `protobuf:"varint,3,opt,name=duration_ms,json=durationMs" json:"duration_ms,omitempty"`
	EnergyDelta                    int32              `protobuf:"varint,5,opt,name=energy_delta,json=energyDelta" json:"energy_delta,omitempty"`
	AttackerIndex                  int32              `protobuf:"varint,6,opt,name=attacker_index,json=attackerIndex" json:"attacker_index,omitempty"`
	TargetIndex                    int32              `protobuf:"varint,7,opt,name=target_index,json=targetIndex" json:"target_index,omitempty"`
	ActivePokemonId                uint64             `protobuf:"fixed64,8,opt,name=active_pokemon_id,json=activePokemonId" json:"active_pokemon_id,omitempty"`
	PlayerJoined                   *BattleParticipant `protobuf:"bytes,9,opt,name=player_joined,json=playerJoined" json:"player_joined,omitempty"`
	BattleResults                  *BattleResults     `protobuf:"bytes,10,opt,name=battle_results,json=battleResults" json:"battle_results,omitempty"`
	DamageWindowsStartTimestampMss int64              `protobuf:"varint,11,opt,name=damage_windows_start_timestamp_mss,json=damageWindowsStartTimestampMss" json:"damage_windows_start_timestamp_mss,omitempty"`
	DamageWindowsEndTimestampMss   int64              `protobuf:"varint,12,opt,name=damage_windows_end_timestamp_mss,json=damageWindowsEndTimestampMss" json:"damage_windows_end_timestamp_mss,omitempty"`
	PlayerLeft                     *BattleParticipant `protobuf:"bytes,13,opt,name=player_left,json=playerLeft" json:"player_left,omitempty"`
	TargetPokemonId                uint64             `protobuf:"fixed64,14,opt,name=target_pokemon_id,json=targetPokemonId" json:"target_pokemon_id,omitempty"`
}

func (m *BattleAction) Reset()                    { *m = BattleAction{} }
func (m *BattleAction) String() string            { return proto.CompactTextString(m) }
func (*BattleAction) ProtoMessage()               {}
func (*BattleAction) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *BattleAction) GetPlayerJoined() *BattleParticipant {
	if m != nil {
		return m.PlayerJoined
	}
	return nil
}

func (m *BattleAction) GetBattleResults() *BattleResults {
	if m != nil {
		return m.BattleResults
	}
	return nil
}

func (m *BattleAction) GetPlayerLeft() *BattleParticipant {
	if m != nil {
		return m.PlayerLeft
	}
	return nil
}

func init() {
	proto.RegisterType((*BattleLog)(nil), "POGOProtos.Data.Battle.BattleLog")
	proto.RegisterType((*BattlePokemonInfo)(nil), "POGOProtos.Data.Battle.BattlePokemonInfo")
	proto.RegisterType((*BattleParticipant)(nil), "POGOProtos.Data.Battle.BattleParticipant")
	proto.RegisterType((*BattleResults)(nil), "POGOProtos.Data.Battle.BattleResults")
	proto.RegisterType((*BattleAction)(nil), "POGOProtos.Data.Battle.BattleAction")
	proto.RegisterEnum("POGOProtos.Data.Battle.BattleType", BattleType_name, BattleType_value)
	proto.RegisterEnum("POGOProtos.Data.Battle.BattleActionType", BattleActionType_name, BattleActionType_value)
	proto.RegisterEnum("POGOProtos.Data.Battle.BattleState", BattleState_name, BattleState_value)
}

func init() { proto.RegisterFile("data_battle.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 1100 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x56, 0xdb, 0x4e, 0xe3, 0x56,
	0x17, 0x9e, 0x9c, 0x80, 0x2c, 0x27, 0xc1, 0xec, 0x9f, 0x43, 0x98, 0x7f, 0xda, 0x32, 0x99, 0x4e,
	0x45, 0xa9, 0xca, 0x48, 0x54, 0x55, 0xd5, 0xaa, 0x55, 0xe5, 0x21, 0x86, 0x06, 0x42, 0xe2, 0x1a,
	0x33, 0x23, 0x7a, 0x63, 0x39, 0xf1, 0x26, 0xb8, 0x24, 0xb6, 0x65, 0xef, 0x0c, 0xf0, 0x3a, 0x7d,
	0x85, 0xbe, 0x4c, 0xdf, 0xa0, 0xf7, 0xbd, 0xee, 0x45, 0xd7, 0x3e, 0x38, 0x07, 0x18, 0x21, 0xb8,
	0x40, 0xd9, 0xfb, 0xdb, 0x6b, 0x7d, 0xeb, 0xbc, 0x30, 0xac, 0xf8, 0x1e, 0xf3, 0xdc, 0x9e, 0xc7,
	0xd8, 0x90, 0xee, 0xc6, 0x49, 0xc4, 0x22, 0xb2, 0x6e, 0x75, 0x0f, 0xbb, 0x16, 0x3f, 0xa6, 0xbb,
	0x4d, 0x7c, 0xdd, 0x7d, 0x2b, 0x5e, 0x9f, 0x03, 0x17, 0x95, 0x32, 0xcf, 0xa5, 0x5a, 0x3c, 0xf4,
	0x6e, 0x69, 0xa2, 0xa0, 0x9a, 0x80, 0x06, 0xb7, 0x23, 0x79, 0x6f, 0xfc, 0x9d, 0x87, 0xb2, 0xd4,
	0x6c, 0x47, 0x03, 0xf2, 0x3d, 0x94, 0x52, 0xe6, 0x31, 0x5a, 0xcf, 0x6d, 0xe5, 0xb6, 0x6b, 0x7b,
	0xaf, 0x76, 0x3f, 0x6e, 0x44, 0xfd, 0x9c, 0x72, 0x51, 0x5b, 0x6a, 0x90, 0x7d, 0xd0, 0xa4, 0x7f,
	0x2e, 0xbb, 0x8d, 0x69, 0x3d, 0x2f, 0x08, 0x1a, 0x0f, 0x13, 0x38, 0x28, 0x69, 0x43, 0x6f, 0x72,
	0x26, 0xff, 0x87, 0x72, 0x4a, 0x93, 0x0f, 0x34, 0x71, 0x47, 0x69, 0xbd, 0x80, 0x14, 0x05, 0x7b,
	0x49, 0x02, 0x27, 0x29, 0x39, 0x86, 0x9a, 0xb2, 0xe0, 0xf5, 0x59, 0x10, 0x85, 0x69, 0xbd, 0xb8,
	0x55, 0xd8, 0xd6, 0xf6, 0x3e, 0x7f, 0xd8, 0x88, 0x21, 0x84, 0xed, 0x6a, 0x6f, 0xe6, 0x96, 0x62,
	0xa4, 0x9b, 0x8a, 0x0c, 0xdd, 0x4f, 0x98, 0xcb, 0x82, 0x11, 0xc5, 0xd3, 0x28, 0xe6, 0x96, 0x4b,
	0xc2, 0xf2, 0x7a, 0x2f, 0x8b, 0x32, 0x61, 0x4e, 0xf6, 0x8c, 0x7e, 0x7c, 0x0b, 0x1b, 0x4a, 0x95,
	0x86, 0xfe, 0xbc, 0xe2, 0x82, 0x50, 0x5c, 0x95, 0xcf, 0x66, 0xe8, 0xcf, 0xa8, 0x35, 0xfe, 0xc8,
	0xc1, 0x8a, 0xf4, 0xc8, 0x8a, 0xae, 0xe8, 0x28, 0x0a, 0x5b, 0xe1, 0x45, 0x44, 0x7e, 0x86, 0x4a,
	0x2c, 0xaf, 0x2e, 0xaf, 0x8c, 0x48, 0xbc, 0xb6, 0xf7, 0xe2, 0x5e, 0x48, 0x4a, 0x87, 0x9f, 0x6d,
	0x2d, 0x9e, 0x5e, 0xc8, 0x6b, 0xa8, 0xf5, 0xc7, 0x49, 0x42, 0x43, 0xe6, 0x5e, 0x52, 0x6f, 0xc8,
	0x2e, 0x45, 0xea, 0x4b, 0x76, 0x55, 0xa1, 0xbf, 0x08, 0x70, 0x56, 0x8c, 0x86, 0x34, 0x19, 0xdc,
	0x8a, 0xf4, 0x4e, 0xc5, 0x4c, 0x01, 0x36, 0xfe, 0xc9, 0x4f, 0x9c, 0xc4, 0xa8, 0x83, 0x7e, 0x10,
	0x7b, 0x21, 0x23, 0x16, 0xd4, 0x78, 0xca, 0x3f, 0x50, 0x57, 0x59, 0x56, 0x6e, 0x7e, 0xf9, 0x70,
	0xe6, 0x67, 0xe2, 0xb4, 0xab, 0x92, 0x40, 0x41, 0xc4, 0x83, 0x75, 0x96, 0x78, 0x01, 0x1a, 0x75,
	0xe3, 0x71, 0x6f, 0x18, 0xf4, 0x5d, 0x6c, 0xc7, 0x8b, 0x60, 0x28, 0x1b, 0x47, 0xdb, 0xfb, 0xea,
	0x7e, 0x02, 0x64, 0x17, 0xcb, 0x1f, 0x4b, 0xe8, 0x58, 0x52, 0xc5, 0x5e, 0x55, 0x54, 0x73, 0x28,
	0xb1, 0x61, 0x39, 0xa1, 0xd8, 0x39, 0xe9, 0xd4, 0xeb, 0x82, 0xe8, 0x97, 0x27, 0x78, 0x5d, 0x53,
	0x0c, 0x99, 0xdb, 0x0e, 0xe8, 0x3e, 0xbd, 0xa0, 0xd8, 0xef, 0xfe, 0x84, 0xb4, 0xf8, 0x54, 0xd2,
	0xe5, 0x8c, 0x42, 0x81, 0x8d, 0x3f, 0xf3, 0x50, 0x95, 0x62, 0x36, 0x4d, 0xc7, 0x43, 0x96, 0x92,
	0x1f, 0xa0, 0x8c, 0x23, 0xea, 0x4e, 0x67, 0x51, 0xdb, 0xfb, 0xe4, 0x9e, 0x81, 0x43, 0x1c, 0x62,
	0xfc, 0x93, 0x53, 0xb8, 0x34, 0x50, 0x27, 0x72, 0x08, 0x65, 0xe4, 0xf2, 0xfa, 0x57, 0xe8, 0x38,
	0x66, 0xf3, 0x31, 0xce, 0x4d, 0x4b, 0x6d, 0x4f, 0x75, 0xd1, 0x89, 0x4d, 0xb9, 0x3a, 0x5c, 0x7a,
	0x13, 0xd3, 0x24, 0xa0, 0x61, 0x1f, 0x47, 0xef, 0xda, 0x4b, 0x7c, 0xea, 0x8b, 0x54, 0x96, 0xec,
	0x0d, 0x29, 0x60, 0x4e, 0xde, 0x0d, 0xf9, 0x4c, 0xbe, 0x83, 0x7a, 0x48, 0x6f, 0x98, 0xcb, 0x43,
	0x0d, 0x7d, 0x5e, 0x65, 0xd5, 0xe4, 0x81, 0x8f, 0x09, 0xe3, 0x43, 0xb2, 0xc6, 0xdf, 0x9b, 0xea,
	0x39, 0x4b, 0x8f, 0x4f, 0xb6, 0x41, 0xe7, 0x91, 0xc7, 0x51, 0x10, 0xb2, 0x14, 0xd5, 0x87, 0x38,
	0x13, 0x25, 0xd1, 0xa9, 0x35, 0xc4, 0x2d, 0x01, 0x37, 0x39, 0xda, 0xf8, 0xab, 0x04, 0x95, 0xd9,
	0x09, 0x27, 0x3f, 0x42, 0x91, 0x2f, 0x11, 0xb5, 0xbb, 0xb6, 0x1f, 0xb3, 0x15, 0xc4, 0x02, 0x2a,
	0xf2, 0x85, 0x45, 0xbe, 0x80, 0x65, 0xb9, 0x56, 0xd4, 0x42, 0x18, 0xa5, 0xa2, 0x15, 0x0b, 0xb2,
	0x73, 0xa3, 0x50, 0xac, 0x01, 0x9c, 0xfe, 0xcf, 0x40, 0xf3, 0xc7, 0x89, 0x27, 0x24, 0xd5, 0x92,
	0x2a, 0xd9, 0x90, 0x41, 0x28, 0xf0, 0x12, 0x2a, 0x72, 0xc2, 0xe6, 0xbc, 0xd7, 0x24, 0x26, 0x5c,
	0xe7, 0xc3, 0x98, 0xa5, 0xd9, 0x0d, 0x30, 0xfe, 0x1b, 0xb1, 0x38, 0x70, 0x18, 0x33, 0xb4, 0xc5,
	0x41, 0xce, 0x84, 0x46, 0x07, 0x94, 0x29, 0xa1, 0x45, 0xc9, 0x24, 0x31, 0x29, 0xb2, 0x03, 0x2b,
	0xf3, 0x93, 0xc9, 0x13, 0xbc, 0x84, 0x72, 0x0b, 0xf6, 0xf2, 0xdc, 0xc4, 0x61, 0x6a, 0x3b, 0x50,
	0x55, 0xf5, 0xfc, 0x1d, 0xd3, 0x88, 0x35, 0x2c, 0x3f, 0x6a, 0x88, 0x67, 0x9a, 0xa3, 0x22, 0xf5,
	0x8f, 0x84, 0x3a, 0x69, 0x4f, 0xf6, 0x71, 0x22, 0xdb, 0xb6, 0x0e, 0x82, 0xf0, 0xf5, 0xc3, 0x84,
	0xaa, 0xc7, 0xb3, 0x85, 0x9c, 0xb5, 0xfc, 0x11, 0x34, 0x7c, 0x6f, 0xe4, 0x0d, 0xa8, 0x7b, 0x8d,
	0xd1, 0x46, 0xd7, 0xe9, 0x47, 0x16, 0x73, 0x5a, 0xd7, 0x44, 0x49, 0x3e, 0x95, 0x92, 0xef, 0xa5,
	0xe0, 0xdd, 0x05, 0x9d, 0x92, 0x03, 0xd8, 0xba, 0xc3, 0x75, 0x77, 0x53, 0xa7, 0xf5, 0x8a, 0x60,
	0x7a, 0x31, 0xc7, 0x34, 0xbf, 0xb1, 0xb9, 0x4f, 0x9a, 0xca, 0xd8, 0x90, 0x5e, 0xb0, 0x7a, 0xf5,
	0xa9, 0xf9, 0x02, 0xa9, 0xdd, 0x46, 0x65, 0x5e, 0x29, 0x55, 0xcc, 0x99, 0x4a, 0xd5, 0x64, 0xa5,
	0xe4, 0xc3, 0xa4, 0x52, 0x3b, 0x3f, 0x01, 0x4c, 0xff, 0x41, 0x92, 0x35, 0x5c, 0xc9, 0x86, 0xe3,
	0xb4, 0x4d, 0xd7, 0x39, 0xb7, 0x4c, 0xf7, 0xac, 0x73, 0x6a, 0x3a, 0xfa, 0x33, 0x02, 0xb0, 0xd0,
	0xe9, 0xda, 0x27, 0x46, 0x5b, 0xcf, 0x91, 0x0a, 0x2c, 0x39, 0xb6, 0xd1, 0xea, 0xb4, 0x3a, 0x87,
	0x7a, 0x7e, 0xe7, 0xdf, 0x1c, 0xe8, 0x77, 0xbb, 0x9c, 0xe8, 0x50, 0x31, 0xf6, 0x9d, 0x56, 0xb7,
	0x33, 0x21, 0x58, 0x81, 0xaa, 0x42, 0x90, 0xdd, 0xd8, 0x3f, 0x46, 0x9e, 0xa9, 0x50, 0xb3, 0xdb,
	0x3c, 0x34, 0xf5, 0x3c, 0xd9, 0x84, 0x35, 0x85, 0x9c, 0x5a, 0xe6, 0x7e, 0xcb, 0x68, 0x67, 0xc2,
	0x05, 0xb2, 0x01, 0xff, 0xcb, 0x9e, 0xde, 0x1b, 0x96, 0x6b, 0x75, 0x8f, 0xcd, 0x93, 0x6e, 0x47,
	0x2f, 0xce, 0xb0, 0x1c, 0xa0, 0x4f, 0x8e, 0x5e, 0x22, 0xeb, 0x40, 0x14, 0x62, 0xb5, 0x8d, 0x73,
	0xd3, 0x76, 0x8f, 0xba, 0xad, 0x8e, 0xbe, 0x70, 0x1f, 0xff, 0xf5, 0xac, 0xe5, 0xe8, 0x8b, 0x84,
	0x40, 0x4d, 0xe1, 0xef, 0x5a, 0xfb, 0x4e, 0xd7, 0x3e, 0xd7, 0x97, 0x66, 0xdc, 0x6d, 0x9a, 0x07,
	0xa6, 0xe1, 0xe8, 0x65, 0xb2, 0x0a, 0xba, 0x82, 0x9c, 0xd6, 0x89, 0xd9, 0x74, 0xbb, 0x67, 0x8e,
	0x0e, 0x3b, 0x0e, 0x68, 0x33, 0xdf, 0x27, 0x64, 0x19, 0xb4, 0x53, 0xc7, 0x70, 0xe6, 0x12, 0xc7,
	0xb5, 0xde, 0x99, 0x18, 0xb0, 0x06, 0x8b, 0x99, 0x85, 0x3c, 0xcf, 0xa2, 0xa4, 0x36, 0x9b, 0x18,
	0x5e, 0x15, 0xca, 0x53, 0xd6, 0xe2, 0xdb, 0x57, 0xbf, 0xbd, 0x1c, 0x04, 0xec, 0x72, 0xdc, 0xdb,
	0xed, 0x47, 0xa3, 0x37, 0xf1, 0xd5, 0x28, 0x1c, 0x44, 0x5f, 0x47, 0x7e, 0xf0, 0x26, 0x8e, 0xf0,
	0x20, 0x3e, 0xa6, 0x52, 0xeb, 0x99, 0x95, 0xb3, 0xf2, 0xbd, 0x05, 0x71, 0xfb, 0xe6, 0xbf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xe9, 0x8b, 0x7e, 0x00, 0xb6, 0x09, 0x00, 0x00,
}
