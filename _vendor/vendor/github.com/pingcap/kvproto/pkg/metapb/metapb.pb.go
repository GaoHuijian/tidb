// Code generated by protoc-gen-gogo.
// source: metapb.proto
// DO NOT EDIT!

/*
	Package metapb is a generated protocol buffer package.

	It is generated from these files:
		metapb.proto

	It has these top-level messages:
		Cluster
		Store
		RegionEpoch
		Region
		Peer
*/
package metapb

import (
	"fmt"
	"io"
	"math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StoreState int32

const (
	StoreState_Up        StoreState = 0
	StoreState_Offline   StoreState = 1
	StoreState_Tombstone StoreState = 2
)

var StoreState_name = map[int32]string{
	0: "Up",
	1: "Offline",
	2: "Tombstone",
}
var StoreState_value = map[string]int32{
	"Up":        0,
	"Offline":   1,
	"Tombstone": 2,
}

func (x StoreState) Enum() *StoreState {
	p := new(StoreState)
	*p = x
	return p
}
func (x StoreState) String() string {
	return proto.EnumName(StoreState_name, int32(x))
}
func (x *StoreState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(StoreState_value, data, "StoreState")
	if err != nil {
		return err
	}
	*x = StoreState(value)
	return nil
}
func (StoreState) EnumDescriptor() ([]byte, []int) { return fileDescriptorMetapb, []int{0} }

type Cluster struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id"`
	// max peer count for a region.
	// pd will do the auto-balance if region peer count mismatches.
	MaxPeerCount     uint32 `protobuf:"varint,2,opt,name=max_peer_count" json:"max_peer_count"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Cluster) Reset()                    { *m = Cluster{} }
func (m *Cluster) String() string            { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()               {}
func (*Cluster) Descriptor() ([]byte, []int) { return fileDescriptorMetapb, []int{0} }

func (m *Cluster) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Cluster) GetMaxPeerCount() uint32 {
	if m != nil {
		return m.MaxPeerCount
	}
	return 0
}

type Store struct {
	Id               uint64     `protobuf:"varint,1,opt,name=id" json:"id"`
	Address          string     `protobuf:"bytes,2,opt,name=address" json:"address"`
	State            StoreState `protobuf:"varint,3,opt,name=state,enum=metapb.StoreState" json:"state"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Store) Reset()                    { *m = Store{} }
func (m *Store) String() string            { return proto.CompactTextString(m) }
func (*Store) ProtoMessage()               {}
func (*Store) Descriptor() ([]byte, []int) { return fileDescriptorMetapb, []int{1} }

func (m *Store) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Store) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Store) GetState() StoreState {
	if m != nil {
		return m.State
	}
	return StoreState_Up
}

type RegionEpoch struct {
	// Conf change version, auto increment when add or remove peer
	ConfVer uint64 `protobuf:"varint,1,opt,name=conf_ver" json:"conf_ver"`
	// Region version, auto increment when split or merge
	Version          uint64 `protobuf:"varint,2,opt,name=version" json:"version"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RegionEpoch) Reset()                    { *m = RegionEpoch{} }
func (m *RegionEpoch) String() string            { return proto.CompactTextString(m) }
func (*RegionEpoch) ProtoMessage()               {}
func (*RegionEpoch) Descriptor() ([]byte, []int) { return fileDescriptorMetapb, []int{2} }

func (m *RegionEpoch) GetConfVer() uint64 {
	if m != nil {
		return m.ConfVer
	}
	return 0
}

func (m *RegionEpoch) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

type Region struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id"`
	// Region key range [start_key, end_key).
	StartKey         []byte       `protobuf:"bytes,2,opt,name=start_key" json:"start_key,omitempty"`
	EndKey           []byte       `protobuf:"bytes,3,opt,name=end_key" json:"end_key,omitempty"`
	RegionEpoch      *RegionEpoch `protobuf:"bytes,4,opt,name=region_epoch" json:"region_epoch,omitempty"`
	Peers            []*Peer      `protobuf:"bytes,5,rep,name=peers" json:"peers,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Region) Reset()                    { *m = Region{} }
func (m *Region) String() string            { return proto.CompactTextString(m) }
func (*Region) ProtoMessage()               {}
func (*Region) Descriptor() ([]byte, []int) { return fileDescriptorMetapb, []int{3} }

func (m *Region) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Region) GetStartKey() []byte {
	if m != nil {
		return m.StartKey
	}
	return nil
}

func (m *Region) GetEndKey() []byte {
	if m != nil {
		return m.EndKey
	}
	return nil
}

func (m *Region) GetRegionEpoch() *RegionEpoch {
	if m != nil {
		return m.RegionEpoch
	}
	return nil
}

func (m *Region) GetPeers() []*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

type Peer struct {
	Id               uint64 `protobuf:"varint,1,opt,name=id" json:"id"`
	StoreId          uint64 `protobuf:"varint,2,opt,name=store_id" json:"store_id"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Peer) Reset()                    { *m = Peer{} }
func (m *Peer) String() string            { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()               {}
func (*Peer) Descriptor() ([]byte, []int) { return fileDescriptorMetapb, []int{4} }

func (m *Peer) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Peer) GetStoreId() uint64 {
	if m != nil {
		return m.StoreId
	}
	return 0
}

func init() {
	proto.RegisterType((*Cluster)(nil), "metapb.Cluster")
	proto.RegisterType((*Store)(nil), "metapb.Store")
	proto.RegisterType((*RegionEpoch)(nil), "metapb.RegionEpoch")
	proto.RegisterType((*Region)(nil), "metapb.Region")
	proto.RegisterType((*Peer)(nil), "metapb.Peer")
	proto.RegisterEnum("metapb.StoreState", StoreState_name, StoreState_value)
}
func (m *Cluster) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Cluster) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintMetapb(data, i, uint64(m.Id))
	data[i] = 0x10
	i++
	i = encodeVarintMetapb(data, i, uint64(m.MaxPeerCount))
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Store) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Store) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintMetapb(data, i, uint64(m.Id))
	data[i] = 0x12
	i++
	i = encodeVarintMetapb(data, i, uint64(len(m.Address)))
	i += copy(data[i:], m.Address)
	data[i] = 0x18
	i++
	i = encodeVarintMetapb(data, i, uint64(m.State))
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *RegionEpoch) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RegionEpoch) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintMetapb(data, i, uint64(m.ConfVer))
	data[i] = 0x10
	i++
	i = encodeVarintMetapb(data, i, uint64(m.Version))
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Region) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Region) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintMetapb(data, i, uint64(m.Id))
	if m.StartKey != nil {
		data[i] = 0x12
		i++
		i = encodeVarintMetapb(data, i, uint64(len(m.StartKey)))
		i += copy(data[i:], m.StartKey)
	}
	if m.EndKey != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintMetapb(data, i, uint64(len(m.EndKey)))
		i += copy(data[i:], m.EndKey)
	}
	if m.RegionEpoch != nil {
		data[i] = 0x22
		i++
		i = encodeVarintMetapb(data, i, uint64(m.RegionEpoch.Size()))
		n1, err := m.RegionEpoch.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Peers) > 0 {
		for _, msg := range m.Peers {
			data[i] = 0x2a
			i++
			i = encodeVarintMetapb(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Peer) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Peer) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintMetapb(data, i, uint64(m.Id))
	data[i] = 0x10
	i++
	i = encodeVarintMetapb(data, i, uint64(m.StoreId))
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Metapb(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Metapb(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMetapb(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Cluster) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovMetapb(uint64(m.Id))
	n += 1 + sovMetapb(uint64(m.MaxPeerCount))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Store) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovMetapb(uint64(m.Id))
	l = len(m.Address)
	n += 1 + l + sovMetapb(uint64(l))
	n += 1 + sovMetapb(uint64(m.State))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RegionEpoch) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovMetapb(uint64(m.ConfVer))
	n += 1 + sovMetapb(uint64(m.Version))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Region) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovMetapb(uint64(m.Id))
	if m.StartKey != nil {
		l = len(m.StartKey)
		n += 1 + l + sovMetapb(uint64(l))
	}
	if m.EndKey != nil {
		l = len(m.EndKey)
		n += 1 + l + sovMetapb(uint64(l))
	}
	if m.RegionEpoch != nil {
		l = m.RegionEpoch.Size()
		n += 1 + l + sovMetapb(uint64(l))
	}
	if len(m.Peers) > 0 {
		for _, e := range m.Peers {
			l = e.Size()
			n += 1 + l + sovMetapb(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Peer) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovMetapb(uint64(m.Id))
	n += 1 + sovMetapb(uint64(m.StoreId))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMetapb(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMetapb(x uint64) (n int) {
	return sovMetapb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Cluster) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetapb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Cluster: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Cluster: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Id |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPeerCount", wireType)
			}
			m.MaxPeerCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.MaxPeerCount |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMetapb(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetapb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Store) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetapb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Store: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Store: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Id |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetapb
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.State |= (StoreState(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMetapb(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetapb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RegionEpoch) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetapb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegionEpoch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegionEpoch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfVer", wireType)
			}
			m.ConfVer = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ConfVer |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Version |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMetapb(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetapb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Region) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetapb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Region: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Region: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Id |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMetapb
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StartKey = append(m.StartKey[:0], data[iNdEx:postIndex]...)
			if m.StartKey == nil {
				m.StartKey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMetapb
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EndKey = append(m.EndKey[:0], data[iNdEx:postIndex]...)
			if m.EndKey == nil {
				m.EndKey = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegionEpoch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetapb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RegionEpoch == nil {
				m.RegionEpoch = &RegionEpoch{}
			}
			if err := m.RegionEpoch.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Peers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetapb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Peers = append(m.Peers, &Peer{})
			if err := m.Peers[len(m.Peers)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetapb(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetapb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Peer) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetapb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Peer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Peer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Id |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreId", wireType)
			}
			m.StoreId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.StoreId |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMetapb(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetapb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMetapb(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetapb
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMetapb
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMetapb
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMetapb
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMetapb
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMetapb(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMetapb = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetapb   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("metapb.proto", fileDescriptorMetapb) }

var fileDescriptorMetapb = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0xcd, 0x4e, 0xfa, 0x40,
	0x14, 0xc5, 0x69, 0x69, 0xe9, 0x9f, 0x4b, 0xe1, 0x5f, 0xc7, 0x8f, 0x34, 0x6a, 0x90, 0x74, 0xe3,
	0xc7, 0x02, 0x09, 0x3b, 0x13, 0x57, 0x18, 0xd7, 0x1a, 0xd1, 0x85, 0xab, 0xa6, 0xd0, 0x0b, 0x12,
	0xa1, 0xd3, 0xcc, 0x0c, 0x46, 0x5f, 0xc0, 0x67, 0xf0, 0x91, 0x58, 0xfa, 0x04, 0xc6, 0xe8, 0x8b,
	0x78, 0x3b, 0x2d, 0x91, 0x18, 0x16, 0x93, 0x76, 0x7e, 0xe7, 0x9e, 0xb9, 0xe7, 0x5e, 0x70, 0x67,
	0xa8, 0xa2, 0x74, 0xd0, 0x4e, 0x05, 0x57, 0x9c, 0x55, 0xf2, 0xdb, 0xee, 0xd6, 0x98, 0x8f, 0xb9,
	0x46, 0xa7, 0xd9, 0x5f, 0xae, 0x06, 0x67, 0xe0, 0x5c, 0x4c, 0xe7, 0x52, 0xa1, 0x60, 0x1e, 0x98,
	0x93, 0xd8, 0x37, 0x5a, 0xc6, 0x91, 0xd5, 0xb3, 0x16, 0x1f, 0x07, 0x25, 0xb6, 0x0f, 0x8d, 0x59,
	0xf4, 0x1c, 0xa6, 0x88, 0x22, 0x1c, 0xf2, 0x79, 0xa2, 0x7c, 0x93, 0xd4, 0x7a, 0xae, 0x06, 0xf7,
	0x60, 0xf7, 0x15, 0x17, 0xb8, 0xc6, 0xb8, 0x0d, 0x4e, 0x14, 0xc7, 0x02, 0xa5, 0xd4, 0x8e, 0x6a,
	0x81, 0x0f, 0xc1, 0x96, 0x2a, 0x52, 0xe8, 0x97, 0x09, 0x36, 0xba, 0xac, 0x5d, 0x04, 0xd5, 0xcf,
	0xf4, 0x33, 0xa5, 0x78, 0xfa, 0x1c, 0x6a, 0x37, 0x38, 0x9e, 0xf0, 0xe4, 0x32, 0xe5, 0xc3, 0x07,
	0xb6, 0x03, 0xff, 0x86, 0x3c, 0x19, 0x85, 0x4f, 0x28, 0xfe, 0xb6, 0x21, 0x24, 0xa9, 0x4e, 0xb7,
	0x29, 0x70, 0xf0, 0x6a, 0x40, 0x25, 0xb7, 0xaf, 0x89, 0xb6, 0x01, 0x55, 0xca, 0x20, 0x54, 0xf8,
	0x88, 0x2f, 0xda, 0xe5, 0xb2, 0xff, 0xe0, 0x60, 0x12, 0x6b, 0x50, 0xd6, 0xe0, 0x18, 0x5c, 0xa1,
	0xfd, 0x21, 0x66, 0xfd, 0x7d, 0x8b, 0x68, 0xad, 0xbb, 0xb9, 0x8c, 0xbb, 0x1a, 0x6d, 0x0f, 0xec,
	0x6c, 0x3d, 0xd2, 0xb7, 0x5b, 0x65, 0xaa, 0x71, 0x97, 0x35, 0xd7, 0x04, 0x83, 0x0e, 0x58, 0xd9,
	0x77, 0x4d, 0x0a, 0x9a, 0x48, 0x66, 0x43, 0x87, 0xc4, 0x57, 0xa2, 0x9f, 0x74, 0x00, 0x7e, 0x97,
	0xc1, 0x2a, 0x60, 0xde, 0xa5, 0x5e, 0x89, 0xd5, 0xc0, 0xb9, 0x1a, 0x8d, 0xa6, 0x93, 0x04, 0x3d,
	0x83, 0xd5, 0xa1, 0x7a, 0xcb, 0x67, 0x03, 0xb2, 0xd3, 0xd5, 0xec, 0x79, 0x8b, 0xaf, 0xa6, 0xf1,
	0x4e, 0xe7, 0x93, 0xce, 0xdb, 0x77, 0xb3, 0xf4, 0x13, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x03, 0xd0,
	0xd5, 0xff, 0x01, 0x00, 0x00,
}
