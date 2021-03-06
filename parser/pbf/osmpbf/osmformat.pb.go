// Code generated by protoc-gen-go.
// source: osmformat.proto
// DO NOT EDIT!

package osmpbf

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Relation_MemberType int32

const (
	Relation_NODE     Relation_MemberType = 0
	Relation_WAY      Relation_MemberType = 1
	Relation_RELATION Relation_MemberType = 2
)

var Relation_MemberType_name = map[int32]string{
	0: "NODE",
	1: "WAY",
	2: "RELATION",
}
var Relation_MemberType_value = map[string]int32{
	"NODE":     0,
	"WAY":      1,
	"RELATION": 2,
}

func (x Relation_MemberType) Enum() *Relation_MemberType {
	p := new(Relation_MemberType)
	*p = x
	return p
}
func (x Relation_MemberType) String() string {
	return proto.EnumName(Relation_MemberType_name, int32(x))
}
func (x Relation_MemberType) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Relation_MemberType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Relation_MemberType_value, data, "Relation_MemberType")
	if err != nil {
		return err
	}
	*x = Relation_MemberType(value)
	return nil
}

type HeaderBlock struct {
	Bbox                             *HeaderBBox `protobuf:"bytes,1,opt,name=bbox" json:"bbox,omitempty"`
	RequiredFeatures                 []string    `protobuf:"bytes,4,rep,name=required_features" json:"required_features,omitempty"`
	OptionalFeatures                 []string    `protobuf:"bytes,5,rep,name=optional_features" json:"optional_features,omitempty"`
	Writingprogram                   *string     `protobuf:"bytes,16,opt,name=writingprogram" json:"writingprogram,omitempty"`
	Source                           *string     `protobuf:"bytes,17,opt,name=source" json:"source,omitempty"`
	OsmosisReplicationTimestamp      *int64      `protobuf:"varint,32,opt,name=osmosis_replication_timestamp" json:"osmosis_replication_timestamp,omitempty"`
	OsmosisReplicationSequenceNumber *int64      `protobuf:"varint,33,opt,name=osmosis_replication_sequence_number" json:"osmosis_replication_sequence_number,omitempty"`
	OsmosisReplicationBaseUrl        *string     `protobuf:"bytes,34,opt,name=osmosis_replication_base_url" json:"osmosis_replication_base_url,omitempty"`
	XXX_unrecognized                 []byte      `json:"-"`
}

func (this *HeaderBlock) Reset()         { *this = HeaderBlock{} }
func (this *HeaderBlock) String() string { return proto.CompactTextString(this) }
func (*HeaderBlock) ProtoMessage()       {}

func (this *HeaderBlock) GetBbox() *HeaderBBox {
	if this != nil {
		return this.Bbox
	}
	return nil
}

func (this *HeaderBlock) GetWritingprogram() string {
	if this != nil && this.Writingprogram != nil {
		return *this.Writingprogram
	}
	return ""
}

func (this *HeaderBlock) GetSource() string {
	if this != nil && this.Source != nil {
		return *this.Source
	}
	return ""
}

func (this *HeaderBlock) GetOsmosisReplicationTimestamp() int64 {
	if this != nil && this.OsmosisReplicationTimestamp != nil {
		return *this.OsmosisReplicationTimestamp
	}
	return 0
}

func (this *HeaderBlock) GetOsmosisReplicationSequenceNumber() int64 {
	if this != nil && this.OsmosisReplicationSequenceNumber != nil {
		return *this.OsmosisReplicationSequenceNumber
	}
	return 0
}

func (this *HeaderBlock) GetOsmosisReplicationBaseUrl() string {
	if this != nil && this.OsmosisReplicationBaseUrl != nil {
		return *this.OsmosisReplicationBaseUrl
	}
	return ""
}

type HeaderBBox struct {
	Left             *int64 `protobuf:"zigzag64,1,req,name=left" json:"left,omitempty"`
	Right            *int64 `protobuf:"zigzag64,2,req,name=right" json:"right,omitempty"`
	Top              *int64 `protobuf:"zigzag64,3,req,name=top" json:"top,omitempty"`
	Bottom           *int64 `protobuf:"zigzag64,4,req,name=bottom" json:"bottom,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *HeaderBBox) Reset()         { *this = HeaderBBox{} }
func (this *HeaderBBox) String() string { return proto.CompactTextString(this) }
func (*HeaderBBox) ProtoMessage()       {}

func (this *HeaderBBox) GetLeft() int64 {
	if this != nil && this.Left != nil {
		return *this.Left
	}
	return 0
}

func (this *HeaderBBox) GetRight() int64 {
	if this != nil && this.Right != nil {
		return *this.Right
	}
	return 0
}

func (this *HeaderBBox) GetTop() int64 {
	if this != nil && this.Top != nil {
		return *this.Top
	}
	return 0
}

func (this *HeaderBBox) GetBottom() int64 {
	if this != nil && this.Bottom != nil {
		return *this.Bottom
	}
	return 0
}

type PrimitiveBlock struct {
	Stringtable      *StringTable      `protobuf:"bytes,1,req,name=stringtable" json:"stringtable,omitempty"`
	Primitivegroup   []*PrimitiveGroup `protobuf:"bytes,2,rep,name=primitivegroup" json:"primitivegroup,omitempty"`
	Granularity      *int32            `protobuf:"varint,17,opt,name=granularity,def=100" json:"granularity,omitempty"`
	LatOffset        *int64            `protobuf:"varint,19,opt,name=lat_offset,def=0" json:"lat_offset,omitempty"`
	LonOffset        *int64            `protobuf:"varint,20,opt,name=lon_offset,def=0" json:"lon_offset,omitempty"`
	DateGranularity  *int32            `protobuf:"varint,18,opt,name=date_granularity,def=1000" json:"date_granularity,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (this *PrimitiveBlock) Reset()         { *this = PrimitiveBlock{} }
func (this *PrimitiveBlock) String() string { return proto.CompactTextString(this) }
func (*PrimitiveBlock) ProtoMessage()       {}

const Default_PrimitiveBlock_Granularity int32 = 100
const Default_PrimitiveBlock_LatOffset int64 = 0
const Default_PrimitiveBlock_LonOffset int64 = 0
const Default_PrimitiveBlock_DateGranularity int32 = 1000

func (this *PrimitiveBlock) GetStringtable() *StringTable {
	if this != nil {
		return this.Stringtable
	}
	return nil
}

func (this *PrimitiveBlock) GetGranularity() int32 {
	if this != nil && this.Granularity != nil {
		return *this.Granularity
	}
	return Default_PrimitiveBlock_Granularity
}

func (this *PrimitiveBlock) GetLatOffset() int64 {
	if this != nil && this.LatOffset != nil {
		return *this.LatOffset
	}
	return Default_PrimitiveBlock_LatOffset
}

func (this *PrimitiveBlock) GetLonOffset() int64 {
	if this != nil && this.LonOffset != nil {
		return *this.LonOffset
	}
	return Default_PrimitiveBlock_LonOffset
}

func (this *PrimitiveBlock) GetDateGranularity() int32 {
	if this != nil && this.DateGranularity != nil {
		return *this.DateGranularity
	}
	return Default_PrimitiveBlock_DateGranularity
}

type PrimitiveGroup struct {
	Nodes            []*Node      `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
	Dense            *DenseNodes  `protobuf:"bytes,2,opt,name=dense" json:"dense,omitempty"`
	Ways             []*Way       `protobuf:"bytes,3,rep,name=ways" json:"ways,omitempty"`
	Relations        []*Relation  `protobuf:"bytes,4,rep,name=relations" json:"relations,omitempty"`
	Changesets       []*ChangeSet `protobuf:"bytes,5,rep,name=changesets" json:"changesets,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (this *PrimitiveGroup) Reset()         { *this = PrimitiveGroup{} }
func (this *PrimitiveGroup) String() string { return proto.CompactTextString(this) }
func (*PrimitiveGroup) ProtoMessage()       {}

func (this *PrimitiveGroup) GetDense() *DenseNodes {
	if this != nil {
		return this.Dense
	}
	return nil
}

type StringTable struct {
	S                [][]byte `protobuf:"bytes,1,rep,name=s" json:"s,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *StringTable) Reset()         { *this = StringTable{} }
func (this *StringTable) String() string { return proto.CompactTextString(this) }
func (*StringTable) ProtoMessage()       {}

type Info struct {
	Version          *int32  `protobuf:"varint,1,opt,name=version,def=-1" json:"version,omitempty"`
	Timestamp        *int64  `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Changeset        *int64  `protobuf:"varint,3,opt,name=changeset" json:"changeset,omitempty"`
	Uid              *int32  `protobuf:"varint,4,opt,name=uid" json:"uid,omitempty"`
	UserSid          *uint32 `protobuf:"varint,5,opt,name=user_sid" json:"user_sid,omitempty"`
	Visible          *bool   `protobuf:"varint,6,opt,name=visible" json:"visible,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *Info) Reset()         { *this = Info{} }
func (this *Info) String() string { return proto.CompactTextString(this) }
func (*Info) ProtoMessage()       {}

const Default_Info_Version int32 = -1

func (this *Info) GetVersion() int32 {
	if this != nil && this.Version != nil {
		return *this.Version
	}
	return Default_Info_Version
}

func (this *Info) GetTimestamp() int64 {
	if this != nil && this.Timestamp != nil {
		return *this.Timestamp
	}
	return 0
}

func (this *Info) GetChangeset() int64 {
	if this != nil && this.Changeset != nil {
		return *this.Changeset
	}
	return 0
}

func (this *Info) GetUid() int32 {
	if this != nil && this.Uid != nil {
		return *this.Uid
	}
	return 0
}

func (this *Info) GetUserSid() uint32 {
	if this != nil && this.UserSid != nil {
		return *this.UserSid
	}
	return 0
}

func (this *Info) GetVisible() bool {
	if this != nil && this.Visible != nil {
		return *this.Visible
	}
	return false
}

type DenseInfo struct {
	Version          []int32 `protobuf:"varint,1,rep,packed,name=version" json:"version,omitempty"`
	Timestamp        []int64 `protobuf:"zigzag64,2,rep,packed,name=timestamp" json:"timestamp,omitempty"`
	Changeset        []int64 `protobuf:"zigzag64,3,rep,packed,name=changeset" json:"changeset,omitempty"`
	Uid              []int32 `protobuf:"zigzag32,4,rep,packed,name=uid" json:"uid,omitempty"`
	UserSid          []int32 `protobuf:"zigzag32,5,rep,packed,name=user_sid" json:"user_sid,omitempty"`
	Visible          []bool  `protobuf:"varint,6,rep,packed,name=visible" json:"visible,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *DenseInfo) Reset()         { *this = DenseInfo{} }
func (this *DenseInfo) String() string { return proto.CompactTextString(this) }
func (*DenseInfo) ProtoMessage()       {}

type ChangeSet struct {
	Id               *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *ChangeSet) Reset()         { *this = ChangeSet{} }
func (this *ChangeSet) String() string { return proto.CompactTextString(this) }
func (*ChangeSet) ProtoMessage()       {}

func (this *ChangeSet) GetId() int64 {
	if this != nil && this.Id != nil {
		return *this.Id
	}
	return 0
}

type Node struct {
	Id               *int64   `protobuf:"zigzag64,1,req,name=id" json:"id,omitempty"`
	Keys             []uint32 `protobuf:"varint,2,rep,packed,name=keys" json:"keys,omitempty"`
	Vals             []uint32 `protobuf:"varint,3,rep,packed,name=vals" json:"vals,omitempty"`
	Info             *Info    `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
	Lat              *int64   `protobuf:"zigzag64,8,req,name=lat" json:"lat,omitempty"`
	Lon              *int64   `protobuf:"zigzag64,9,req,name=lon" json:"lon,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *Node) Reset()         { *this = Node{} }
func (this *Node) String() string { return proto.CompactTextString(this) }
func (*Node) ProtoMessage()       {}

func (this *Node) GetId() int64 {
	if this != nil && this.Id != nil {
		return *this.Id
	}
	return 0
}

func (this *Node) GetInfo() *Info {
	if this != nil {
		return this.Info
	}
	return nil
}

func (this *Node) GetLat() int64 {
	if this != nil && this.Lat != nil {
		return *this.Lat
	}
	return 0
}

func (this *Node) GetLon() int64 {
	if this != nil && this.Lon != nil {
		return *this.Lon
	}
	return 0
}

type DenseNodes struct {
	Id               []int64    `protobuf:"zigzag64,1,rep,packed,name=id" json:"id,omitempty"`
	Denseinfo        *DenseInfo `protobuf:"bytes,5,opt,name=denseinfo" json:"denseinfo,omitempty"`
	Lat              []int64    `protobuf:"zigzag64,8,rep,packed,name=lat" json:"lat,omitempty"`
	Lon              []int64    `protobuf:"zigzag64,9,rep,packed,name=lon" json:"lon,omitempty"`
	KeysVals         []int32    `protobuf:"varint,10,rep,packed,name=keys_vals" json:"keys_vals,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (this *DenseNodes) Reset()         { *this = DenseNodes{} }
func (this *DenseNodes) String() string { return proto.CompactTextString(this) }
func (*DenseNodes) ProtoMessage()       {}

func (this *DenseNodes) GetDenseinfo() *DenseInfo {
	if this != nil {
		return this.Denseinfo
	}
	return nil
}

type Way struct {
	Id               *int64   `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Keys             []uint32 `protobuf:"varint,2,rep,packed,name=keys" json:"keys,omitempty"`
	Vals             []uint32 `protobuf:"varint,3,rep,packed,name=vals" json:"vals,omitempty"`
	Info             *Info    `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
	Refs             []int64  `protobuf:"zigzag64,8,rep,packed,name=refs" json:"refs,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *Way) Reset()         { *this = Way{} }
func (this *Way) String() string { return proto.CompactTextString(this) }
func (*Way) ProtoMessage()       {}

func (this *Way) GetId() int64 {
	if this != nil && this.Id != nil {
		return *this.Id
	}
	return 0
}

func (this *Way) GetInfo() *Info {
	if this != nil {
		return this.Info
	}
	return nil
}

type Relation struct {
	Id               *int64                `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Keys             []uint32              `protobuf:"varint,2,rep,packed,name=keys" json:"keys,omitempty"`
	Vals             []uint32              `protobuf:"varint,3,rep,packed,name=vals" json:"vals,omitempty"`
	Info             *Info                 `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
	RolesSid         []int32               `protobuf:"varint,8,rep,packed,name=roles_sid" json:"roles_sid,omitempty"`
	Memids           []int64               `protobuf:"zigzag64,9,rep,packed,name=memids" json:"memids,omitempty"`
	Types            []Relation_MemberType `protobuf:"varint,10,rep,packed,name=types,enum=imposm3.Relation_MemberType" json:"types,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (this *Relation) Reset()         { *this = Relation{} }
func (this *Relation) String() string { return proto.CompactTextString(this) }
func (*Relation) ProtoMessage()       {}

func (this *Relation) GetId() int64 {
	if this != nil && this.Id != nil {
		return *this.Id
	}
	return 0
}

func (this *Relation) GetInfo() *Info {
	if this != nil {
		return this.Info
	}
	return nil
}

func init() {
	proto.RegisterEnum("imposm3.Relation_MemberType", Relation_MemberType_name, Relation_MemberType_value)
}
