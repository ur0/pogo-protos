// Code generated by protoc-gen-go.
// source: data_logs.proto
// DO NOT EDIT!

package pogo_protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ItemData from inventory_item.proto

// Ignoring public import of ItemAward from inventory_item.proto

// Ignoring public import of ItemId from inventory_item.proto

// Ignoring public import of ItemType from inventory_item.proto

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

type FortSearchLogEntry_Result int32

const (
	FortSearchLogEntry_UNSET   FortSearchLogEntry_Result = 0
	FortSearchLogEntry_SUCCESS FortSearchLogEntry_Result = 1
)

var FortSearchLogEntry_Result_name = map[int32]string{
	0: "UNSET",
	1: "SUCCESS",
}
var FortSearchLogEntry_Result_value = map[string]int32{
	"UNSET":   0,
	"SUCCESS": 1,
}

func (x FortSearchLogEntry_Result) String() string {
	return proto.EnumName(FortSearchLogEntry_Result_name, int32(x))
}
func (FortSearchLogEntry_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor15, []int{0, 0}
}

type CatchPokemonLogEntry_Result int32

const (
	CatchPokemonLogEntry_UNSET            CatchPokemonLogEntry_Result = 0
	CatchPokemonLogEntry_POKEMON_CAPTURED CatchPokemonLogEntry_Result = 1
	CatchPokemonLogEntry_POKEMON_FLED     CatchPokemonLogEntry_Result = 2
)

var CatchPokemonLogEntry_Result_name = map[int32]string{
	0: "UNSET",
	1: "POKEMON_CAPTURED",
	2: "POKEMON_FLED",
}
var CatchPokemonLogEntry_Result_value = map[string]int32{
	"UNSET":            0,
	"POKEMON_CAPTURED": 1,
	"POKEMON_FLED":     2,
}

func (x CatchPokemonLogEntry_Result) String() string {
	return proto.EnumName(CatchPokemonLogEntry_Result_name, int32(x))
}
func (CatchPokemonLogEntry_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor15, []int{1, 0}
}

type FortSearchLogEntry struct {
	Result FortSearchLogEntry_Result `protobuf:"varint,1,opt,name=result,enum=POGOProtos.Data.Logs.FortSearchLogEntry_Result" json:"result,omitempty"`
	FortId string                    `protobuf:"bytes,2,opt,name=fort_id,json=fortId" json:"fort_id,omitempty"`
	Items  []*ItemData               `protobuf:"bytes,3,rep,name=items" json:"items,omitempty"`
	Eggs   int32                     `protobuf:"varint,4,opt,name=eggs" json:"eggs,omitempty"`
}

func (m *FortSearchLogEntry) Reset()                    { *m = FortSearchLogEntry{} }
func (m *FortSearchLogEntry) String() string            { return proto.CompactTextString(m) }
func (*FortSearchLogEntry) ProtoMessage()               {}
func (*FortSearchLogEntry) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{0} }

func (m *FortSearchLogEntry) GetItems() []*ItemData {
	if m != nil {
		return m.Items
	}
	return nil
}

type CatchPokemonLogEntry struct {
	Result        CatchPokemonLogEntry_Result `protobuf:"varint,1,opt,name=result,enum=POGOProtos.Data.Logs.CatchPokemonLogEntry_Result" json:"result,omitempty"`
	PokemonId     PokemonId                   `protobuf:"varint,2,opt,name=pokemon_id,json=pokemonId,enum=POGOProtos.Enums.PokemonId" json:"pokemon_id,omitempty"`
	CombatPoints  int32                       `protobuf:"varint,3,opt,name=combat_points,json=combatPoints" json:"combat_points,omitempty"`
	PokemonDataId uint64                      `protobuf:"varint,4,opt,name=pokemon_data_id,json=pokemonDataId" json:"pokemon_data_id,omitempty"`
}

func (m *CatchPokemonLogEntry) Reset()                    { *m = CatchPokemonLogEntry{} }
func (m *CatchPokemonLogEntry) String() string            { return proto.CompactTextString(m) }
func (*CatchPokemonLogEntry) ProtoMessage()               {}
func (*CatchPokemonLogEntry) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{1} }

type ActionLogEntry struct {
	TimestampMs int64 `protobuf:"varint,1,opt,name=timestamp_ms,json=timestampMs" json:"timestamp_ms,omitempty"`
	Sfida       bool  `protobuf:"varint,2,opt,name=sfida" json:"sfida,omitempty"`
	// Types that are valid to be assigned to Action:
	//	*ActionLogEntry_CatchPokemon
	//	*ActionLogEntry_FortSearch
	Action isActionLogEntry_Action `protobuf_oneof:"Action"`
}

func (m *ActionLogEntry) Reset()                    { *m = ActionLogEntry{} }
func (m *ActionLogEntry) String() string            { return proto.CompactTextString(m) }
func (*ActionLogEntry) ProtoMessage()               {}
func (*ActionLogEntry) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{2} }

type isActionLogEntry_Action interface {
	isActionLogEntry_Action()
}

type ActionLogEntry_CatchPokemon struct {
	CatchPokemon *CatchPokemonLogEntry `protobuf:"bytes,3,opt,name=catch_pokemon,json=catchPokemon,oneof"`
}
type ActionLogEntry_FortSearch struct {
	FortSearch *FortSearchLogEntry `protobuf:"bytes,4,opt,name=fort_search,json=fortSearch,oneof"`
}

func (*ActionLogEntry_CatchPokemon) isActionLogEntry_Action() {}
func (*ActionLogEntry_FortSearch) isActionLogEntry_Action()   {}

func (m *ActionLogEntry) GetAction() isActionLogEntry_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *ActionLogEntry) GetCatchPokemon() *CatchPokemonLogEntry {
	if x, ok := m.GetAction().(*ActionLogEntry_CatchPokemon); ok {
		return x.CatchPokemon
	}
	return nil
}

func (m *ActionLogEntry) GetFortSearch() *FortSearchLogEntry {
	if x, ok := m.GetAction().(*ActionLogEntry_FortSearch); ok {
		return x.FortSearch
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ActionLogEntry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ActionLogEntry_OneofMarshaler, _ActionLogEntry_OneofUnmarshaler, _ActionLogEntry_OneofSizer, []interface{}{
		(*ActionLogEntry_CatchPokemon)(nil),
		(*ActionLogEntry_FortSearch)(nil),
	}
}

func _ActionLogEntry_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ActionLogEntry)
	// Action
	switch x := m.Action.(type) {
	case *ActionLogEntry_CatchPokemon:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CatchPokemon); err != nil {
			return err
		}
	case *ActionLogEntry_FortSearch:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FortSearch); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ActionLogEntry.Action has unexpected type %T", x)
	}
	return nil
}

func _ActionLogEntry_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ActionLogEntry)
	switch tag {
	case 3: // Action.catch_pokemon
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CatchPokemonLogEntry)
		err := b.DecodeMessage(msg)
		m.Action = &ActionLogEntry_CatchPokemon{msg}
		return true, err
	case 4: // Action.fort_search
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FortSearchLogEntry)
		err := b.DecodeMessage(msg)
		m.Action = &ActionLogEntry_FortSearch{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ActionLogEntry_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ActionLogEntry)
	// Action
	switch x := m.Action.(type) {
	case *ActionLogEntry_CatchPokemon:
		s := proto.Size(x.CatchPokemon)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ActionLogEntry_FortSearch:
		s := proto.Size(x.FortSearch)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*FortSearchLogEntry)(nil), "POGOProtos.Data.Logs.FortSearchLogEntry")
	proto.RegisterType((*CatchPokemonLogEntry)(nil), "POGOProtos.Data.Logs.CatchPokemonLogEntry")
	proto.RegisterType((*ActionLogEntry)(nil), "POGOProtos.Data.Logs.ActionLogEntry")
	proto.RegisterEnum("POGOProtos.Data.Logs.FortSearchLogEntry_Result", FortSearchLogEntry_Result_name, FortSearchLogEntry_Result_value)
	proto.RegisterEnum("POGOProtos.Data.Logs.CatchPokemonLogEntry_Result", CatchPokemonLogEntry_Result_name, CatchPokemonLogEntry_Result_value)
}

func init() { proto.RegisterFile("data_logs.proto", fileDescriptor15) }

var fileDescriptor15 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0x9b, 0x4c,
	0x14, 0x35, 0xfe, 0x21, 0xf1, 0xc5, 0x71, 0xd0, 0xc8, 0xd2, 0x67, 0xe5, 0xdb, 0x38, 0x58, 0xaa,
	0xac, 0x4a, 0xc5, 0x2a, 0x5d, 0xb5, 0x5d, 0x25, 0x0e, 0x49, 0x51, 0x7e, 0x4c, 0x87, 0x78, 0xd3,
	0x0d, 0xc2, 0x06, 0x63, 0x94, 0xc0, 0x20, 0x66, 0x5c, 0x29, 0x4f, 0xd4, 0x47, 0xeb, 0x03, 0xf4,
	0x05, 0x3a, 0x33, 0x80, 0x65, 0xa5, 0x5e, 0x64, 0x63, 0xdd, 0x39, 0x33, 0xe7, 0x9e, 0x7b, 0x8e,
	0x2f, 0x70, 0x1a, 0x06, 0x2c, 0xf0, 0x9f, 0x49, 0x4c, 0xcd, 0xbc, 0x20, 0x8c, 0xa0, 0x81, 0x3b,
	0xbf, 0x99, 0xbb, 0xa2, 0xa4, 0xe6, 0x15, 0xbf, 0x33, 0xef, 0xf8, 0xdd, 0xd9, 0x20, 0xc9, 0x7e,
	0x46, 0x19, 0x23, 0xc5, 0x8b, 0x9f, 0xb0, 0x28, 0x2d, 0xdf, 0x9e, 0x69, 0x51, 0xb6, 0x4d, 0x2b,
	0xa2, 0xf1, 0x5b, 0x01, 0x74, 0x4d, 0x0a, 0xe6, 0x45, 0x41, 0xb1, 0xda, 0x70, 0x96, 0x9d, 0xb1,
	0xe2, 0x05, 0xdd, 0x80, 0x5a, 0x44, 0x74, 0xfb, 0xcc, 0x86, 0xca, 0x48, 0x99, 0xf4, 0xad, 0xa9,
	0x79, 0x48, 0xc0, 0xfc, 0x97, 0x69, 0x62, 0x49, 0xc3, 0x15, 0x1d, 0xfd, 0x07, 0x47, 0x6b, 0xfe,
	0xc8, 0x4f, 0xc2, 0x61, 0x93, 0x77, 0xea, 0x62, 0x55, 0x1c, 0x9d, 0x10, 0x7d, 0x86, 0x8e, 0x98,
	0x89, 0x0e, 0x5b, 0xa3, 0xd6, 0x44, 0xb3, 0xc6, 0xfb, 0x02, 0x4e, 0x3d, 0xb6, 0xe9, 0x88, 0xb1,
	0xc5, 0x8f, 0xd0, 0xc4, 0x25, 0x03, 0x21, 0x68, 0x47, 0x71, 0x4c, 0x87, 0x6d, 0xde, 0xb0, 0x83,
	0x65, 0x6d, 0x8c, 0x40, 0x2d, 0x95, 0x51, 0x17, 0x3a, 0x8b, 0x07, 0xcf, 0x7e, 0xd4, 0x1b, 0x48,
	0x83, 0x23, 0x6f, 0x31, 0x9b, 0xd9, 0x9e, 0xa7, 0x2b, 0xc6, 0xaf, 0x26, 0x0c, 0x66, 0x01, 0x5b,
	0x6d, 0x5c, 0xf2, 0x14, 0xa5, 0x24, 0xdb, 0x79, 0x75, 0x5e, 0x79, 0xfd, 0x78, 0xd8, 0xeb, 0x21,
	0xee, 0x6b, 0xb7, 0x5f, 0x00, 0xf2, 0xf2, 0x45, 0x6d, 0xb8, 0x6f, 0xfd, 0xbf, 0xdf, 0xce, 0x96,
	0xd1, 0x57, 0x5d, 0x9c, 0x10, 0x77, 0xf3, 0xba, 0x44, 0x63, 0x38, 0x59, 0x91, 0x74, 0x19, 0x30,
	0x3f, 0x27, 0x49, 0xc6, 0x44, 0x30, 0xc2, 0x5e, 0xaf, 0x04, 0x5d, 0x89, 0xa1, 0x77, 0x70, 0x5a,
	0x0b, 0xc8, 0x15, 0xe0, 0x2a, 0x22, 0x85, 0x36, 0x3e, 0xa9, 0x60, 0x31, 0xaf, 0x13, 0x1a, 0x5f,
	0x0f, 0xc5, 0x31, 0x00, 0xdd, 0x9d, 0xdf, 0xda, 0xf7, 0xf3, 0x07, 0x7f, 0x76, 0xe1, 0x3e, 0x2e,
	0xb0, 0x7d, 0xa5, 0x2b, 0x48, 0x87, 0x5e, 0x8d, 0x5e, 0xdf, 0x71, 0xa4, 0x69, 0xfc, 0x51, 0xa0,
	0x7f, 0xb1, 0x62, 0xc9, 0x5e, 0x46, 0xe7, 0xd0, 0x63, 0x49, 0x1a, 0x51, 0x16, 0xa4, 0xb9, 0xcf,
	0xff, 0x34, 0x91, 0x54, 0x0b, 0x6b, 0x3b, 0xec, 0x9e, 0xf2, 0xee, 0x1d, 0xba, 0x4e, 0xc2, 0x40,
	0xda, 0x3e, 0xc6, 0xe5, 0x01, 0x7d, 0xe7, 0xae, 0x44, 0x70, 0x7e, 0x35, 0x9f, 0x74, 0xa5, 0x59,
	0xef, 0xdf, 0x9e, 0xf1, 0xb7, 0x06, 0xcf, 0x60, 0x0f, 0x47, 0xb7, 0xa0, 0xc9, 0x95, 0xa2, 0x72,
	0xf1, 0xa4, 0x7f, 0xcd, 0x9a, 0xbc, 0x75, 0x41, 0x79, 0x3b, 0x58, 0xef, 0xd0, 0xcb, 0x63, 0x50,
	0x4b, 0xab, 0x97, 0xe3, 0x1f, 0xe7, 0x71, 0xc2, 0x36, 0xdb, 0xa5, 0xc9, 0x13, 0x9f, 0xe6, 0x4f,
	0x69, 0x16, 0x93, 0x0f, 0x24, 0x4c, 0xa6, 0x39, 0xe1, 0x85, 0xfc, 0x5a, 0xa8, 0xdb, 0x70, 0x95,
	0xa5, 0x2a, 0xeb, 0x4f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xe0, 0xf9, 0xcd, 0x85, 0x03,
	0x00, 0x00,
}
