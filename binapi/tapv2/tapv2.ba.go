// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.5.0-dev
//  VPP:              22.02-release
// source: /usr/share/vpp/api/core/tapv2.api.json

// Package tapv2 contains generated bindings for API file tapv2.api.
//
// Contents:
//   1 enum
//   8 messages
//
package tapv2

import (
	"strconv"

	api "go.fd.io/govpp/api"
	ethernet_types "go.fd.io/govpp/binapi/ethernet_types"
	interface_types "go.fd.io/govpp/binapi/interface_types"
	ip_types "go.fd.io/govpp/binapi/ip_types"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "tapv2"
	APIVersion = "4.0.0"
	VersionCrc = 0xc2f80dc7
)

// TapFlags defines enum 'tap_flags'.
type TapFlags uint32

const (
	TAP_API_FLAG_GSO          TapFlags = 1
	TAP_API_FLAG_CSUM_OFFLOAD TapFlags = 2
	TAP_API_FLAG_PERSIST      TapFlags = 4
	TAP_API_FLAG_ATTACH       TapFlags = 8
	TAP_API_FLAG_TUN          TapFlags = 16
	TAP_API_FLAG_GRO_COALESCE TapFlags = 32
	TAP_API_FLAG_PACKED       TapFlags = 64
	TAP_API_FLAG_IN_ORDER     TapFlags = 128
)

var (
	TapFlags_name = map[uint32]string{
		1:   "TAP_API_FLAG_GSO",
		2:   "TAP_API_FLAG_CSUM_OFFLOAD",
		4:   "TAP_API_FLAG_PERSIST",
		8:   "TAP_API_FLAG_ATTACH",
		16:  "TAP_API_FLAG_TUN",
		32:  "TAP_API_FLAG_GRO_COALESCE",
		64:  "TAP_API_FLAG_PACKED",
		128: "TAP_API_FLAG_IN_ORDER",
	}
	TapFlags_value = map[string]uint32{
		"TAP_API_FLAG_GSO":          1,
		"TAP_API_FLAG_CSUM_OFFLOAD": 2,
		"TAP_API_FLAG_PERSIST":      4,
		"TAP_API_FLAG_ATTACH":       8,
		"TAP_API_FLAG_TUN":          16,
		"TAP_API_FLAG_GRO_COALESCE": 32,
		"TAP_API_FLAG_PACKED":       64,
		"TAP_API_FLAG_IN_ORDER":     128,
	}
)

func (x TapFlags) String() string {
	s, ok := TapFlags_name[uint32(x)]
	if ok {
		return s
	}
	str := func(n uint32) string {
		s, ok := TapFlags_name[uint32(n)]
		if ok {
			return s
		}
		return "TapFlags(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint32(0); i <= 32; i++ {
		val := uint32(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint32(x))
	}
	return s
}

// SwInterfaceTapV2Details defines message 'sw_interface_tap_v2_details'.
type SwInterfaceTapV2Details struct {
	SwIfIndex     uint32                        `binapi:"u32,name=sw_if_index" json:"sw_if_index,omitempty"`
	ID            uint32                        `binapi:"u32,name=id" json:"id,omitempty"`
	TxRingSz      uint16                        `binapi:"u16,name=tx_ring_sz" json:"tx_ring_sz,omitempty"`
	RxRingSz      uint16                        `binapi:"u16,name=rx_ring_sz" json:"rx_ring_sz,omitempty"`
	HostMtuSize   uint32                        `binapi:"u32,name=host_mtu_size" json:"host_mtu_size,omitempty"`
	HostMacAddr   ethernet_types.MacAddress     `binapi:"mac_address,name=host_mac_addr" json:"host_mac_addr,omitempty"`
	HostIP4Prefix ip_types.IP4AddressWithPrefix `binapi:"ip4_address_with_prefix,name=host_ip4_prefix" json:"host_ip4_prefix,omitempty"`
	HostIP6Prefix ip_types.IP6AddressWithPrefix `binapi:"ip6_address_with_prefix,name=host_ip6_prefix" json:"host_ip6_prefix,omitempty"`
	TapFlags      TapFlags                      `binapi:"tap_flags,name=tap_flags" json:"tap_flags,omitempty"`
	DevName       string                        `binapi:"string[64],name=dev_name" json:"dev_name,omitempty"`
	HostIfName    string                        `binapi:"string[64],name=host_if_name" json:"host_if_name,omitempty"`
	HostNamespace string                        `binapi:"string[64],name=host_namespace" json:"host_namespace,omitempty"`
	HostBridge    string                        `binapi:"string[64],name=host_bridge" json:"host_bridge,omitempty"`
}

func (m *SwInterfaceTapV2Details) Reset()               { *m = SwInterfaceTapV2Details{} }
func (*SwInterfaceTapV2Details) GetMessageName() string { return "sw_interface_tap_v2_details" }
func (*SwInterfaceTapV2Details) GetCrcString() string   { return "1e2b2a47" }
func (*SwInterfaceTapV2Details) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SwInterfaceTapV2Details) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.SwIfIndex
	size += 4      // m.ID
	size += 2      // m.TxRingSz
	size += 2      // m.RxRingSz
	size += 4      // m.HostMtuSize
	size += 1 * 6  // m.HostMacAddr
	size += 1 * 4  // m.HostIP4Prefix.Address
	size += 1      // m.HostIP4Prefix.Len
	size += 1 * 16 // m.HostIP6Prefix.Address
	size += 1      // m.HostIP6Prefix.Len
	size += 4      // m.TapFlags
	size += 64     // m.DevName
	size += 64     // m.HostIfName
	size += 64     // m.HostNamespace
	size += 64     // m.HostBridge
	return size
}
func (m *SwInterfaceTapV2Details) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.SwIfIndex)
	buf.EncodeUint32(m.ID)
	buf.EncodeUint16(m.TxRingSz)
	buf.EncodeUint16(m.RxRingSz)
	buf.EncodeUint32(m.HostMtuSize)
	buf.EncodeBytes(m.HostMacAddr[:], 6)
	buf.EncodeBytes(m.HostIP4Prefix.Address[:], 4)
	buf.EncodeUint8(m.HostIP4Prefix.Len)
	buf.EncodeBytes(m.HostIP6Prefix.Address[:], 16)
	buf.EncodeUint8(m.HostIP6Prefix.Len)
	buf.EncodeUint32(uint32(m.TapFlags))
	buf.EncodeString(m.DevName, 64)
	buf.EncodeString(m.HostIfName, 64)
	buf.EncodeString(m.HostNamespace, 64)
	buf.EncodeString(m.HostBridge, 64)
	return buf.Bytes(), nil
}
func (m *SwInterfaceTapV2Details) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = buf.DecodeUint32()
	m.ID = buf.DecodeUint32()
	m.TxRingSz = buf.DecodeUint16()
	m.RxRingSz = buf.DecodeUint16()
	m.HostMtuSize = buf.DecodeUint32()
	copy(m.HostMacAddr[:], buf.DecodeBytes(6))
	copy(m.HostIP4Prefix.Address[:], buf.DecodeBytes(4))
	m.HostIP4Prefix.Len = buf.DecodeUint8()
	copy(m.HostIP6Prefix.Address[:], buf.DecodeBytes(16))
	m.HostIP6Prefix.Len = buf.DecodeUint8()
	m.TapFlags = TapFlags(buf.DecodeUint32())
	m.DevName = buf.DecodeString(64)
	m.HostIfName = buf.DecodeString(64)
	m.HostNamespace = buf.DecodeString(64)
	m.HostBridge = buf.DecodeString(64)
	return nil
}

// SwInterfaceTapV2Dump defines message 'sw_interface_tap_v2_dump'.
type SwInterfaceTapV2Dump struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index,default=4294967295" json:"sw_if_index,omitempty"`
}

func (m *SwInterfaceTapV2Dump) Reset()               { *m = SwInterfaceTapV2Dump{} }
func (*SwInterfaceTapV2Dump) GetMessageName() string { return "sw_interface_tap_v2_dump" }
func (*SwInterfaceTapV2Dump) GetCrcString() string   { return "f9e6675e" }
func (*SwInterfaceTapV2Dump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SwInterfaceTapV2Dump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *SwInterfaceTapV2Dump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *SwInterfaceTapV2Dump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// TapCreateV2 defines message 'tap_create_v2'.
type TapCreateV2 struct {
	ID               uint32                        `binapi:"u32,name=id,default=4294967295" json:"id,omitempty"`
	UseRandomMac     bool                          `binapi:"bool,name=use_random_mac,default=true" json:"use_random_mac,omitempty"`
	MacAddress       ethernet_types.MacAddress     `binapi:"mac_address,name=mac_address" json:"mac_address,omitempty"`
	NumRxQueues      uint8                         `binapi:"u8,name=num_rx_queues,default=1" json:"num_rx_queues,omitempty"`
	TxRingSz         uint16                        `binapi:"u16,name=tx_ring_sz,default=256" json:"tx_ring_sz,omitempty"`
	RxRingSz         uint16                        `binapi:"u16,name=rx_ring_sz,default=256" json:"rx_ring_sz,omitempty"`
	HostMtuSet       bool                          `binapi:"bool,name=host_mtu_set" json:"host_mtu_set,omitempty"`
	HostMtuSize      uint32                        `binapi:"u32,name=host_mtu_size" json:"host_mtu_size,omitempty"`
	HostMacAddrSet   bool                          `binapi:"bool,name=host_mac_addr_set" json:"host_mac_addr_set,omitempty"`
	HostMacAddr      ethernet_types.MacAddress     `binapi:"mac_address,name=host_mac_addr" json:"host_mac_addr,omitempty"`
	HostIP4PrefixSet bool                          `binapi:"bool,name=host_ip4_prefix_set" json:"host_ip4_prefix_set,omitempty"`
	HostIP4Prefix    ip_types.IP4AddressWithPrefix `binapi:"ip4_address_with_prefix,name=host_ip4_prefix" json:"host_ip4_prefix,omitempty"`
	HostIP6PrefixSet bool                          `binapi:"bool,name=host_ip6_prefix_set" json:"host_ip6_prefix_set,omitempty"`
	HostIP6Prefix    ip_types.IP6AddressWithPrefix `binapi:"ip6_address_with_prefix,name=host_ip6_prefix" json:"host_ip6_prefix,omitempty"`
	HostIP4GwSet     bool                          `binapi:"bool,name=host_ip4_gw_set" json:"host_ip4_gw_set,omitempty"`
	HostIP4Gw        ip_types.IP4Address           `binapi:"ip4_address,name=host_ip4_gw" json:"host_ip4_gw,omitempty"`
	HostIP6GwSet     bool                          `binapi:"bool,name=host_ip6_gw_set" json:"host_ip6_gw_set,omitempty"`
	HostIP6Gw        ip_types.IP6Address           `binapi:"ip6_address,name=host_ip6_gw" json:"host_ip6_gw,omitempty"`
	TapFlags         TapFlags                      `binapi:"tap_flags,name=tap_flags" json:"tap_flags,omitempty"`
	HostNamespaceSet bool                          `binapi:"bool,name=host_namespace_set" json:"host_namespace_set,omitempty"`
	HostNamespace    string                        `binapi:"string[64],name=host_namespace" json:"host_namespace,omitempty"`
	HostIfNameSet    bool                          `binapi:"bool,name=host_if_name_set" json:"host_if_name_set,omitempty"`
	HostIfName       string                        `binapi:"string[64],name=host_if_name" json:"host_if_name,omitempty"`
	HostBridgeSet    bool                          `binapi:"bool,name=host_bridge_set" json:"host_bridge_set,omitempty"`
	HostBridge       string                        `binapi:"string[64],name=host_bridge" json:"host_bridge,omitempty"`
	Tag              string                        `binapi:"string[],name=tag" json:"tag,omitempty"`
}

func (m *TapCreateV2) Reset()               { *m = TapCreateV2{} }
func (*TapCreateV2) GetMessageName() string { return "tap_create_v2" }
func (*TapCreateV2) GetCrcString() string   { return "2d0d6570" }
func (*TapCreateV2) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TapCreateV2) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4              // m.ID
	size += 1              // m.UseRandomMac
	size += 1 * 6          // m.MacAddress
	size += 1              // m.NumRxQueues
	size += 2              // m.TxRingSz
	size += 2              // m.RxRingSz
	size += 1              // m.HostMtuSet
	size += 4              // m.HostMtuSize
	size += 1              // m.HostMacAddrSet
	size += 1 * 6          // m.HostMacAddr
	size += 1              // m.HostIP4PrefixSet
	size += 1 * 4          // m.HostIP4Prefix.Address
	size += 1              // m.HostIP4Prefix.Len
	size += 1              // m.HostIP6PrefixSet
	size += 1 * 16         // m.HostIP6Prefix.Address
	size += 1              // m.HostIP6Prefix.Len
	size += 1              // m.HostIP4GwSet
	size += 1 * 4          // m.HostIP4Gw
	size += 1              // m.HostIP6GwSet
	size += 1 * 16         // m.HostIP6Gw
	size += 4              // m.TapFlags
	size += 1              // m.HostNamespaceSet
	size += 64             // m.HostNamespace
	size += 1              // m.HostIfNameSet
	size += 64             // m.HostIfName
	size += 1              // m.HostBridgeSet
	size += 64             // m.HostBridge
	size += 4 + len(m.Tag) // m.Tag
	return size
}
func (m *TapCreateV2) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.ID)
	buf.EncodeBool(m.UseRandomMac)
	buf.EncodeBytes(m.MacAddress[:], 6)
	buf.EncodeUint8(m.NumRxQueues)
	buf.EncodeUint16(m.TxRingSz)
	buf.EncodeUint16(m.RxRingSz)
	buf.EncodeBool(m.HostMtuSet)
	buf.EncodeUint32(m.HostMtuSize)
	buf.EncodeBool(m.HostMacAddrSet)
	buf.EncodeBytes(m.HostMacAddr[:], 6)
	buf.EncodeBool(m.HostIP4PrefixSet)
	buf.EncodeBytes(m.HostIP4Prefix.Address[:], 4)
	buf.EncodeUint8(m.HostIP4Prefix.Len)
	buf.EncodeBool(m.HostIP6PrefixSet)
	buf.EncodeBytes(m.HostIP6Prefix.Address[:], 16)
	buf.EncodeUint8(m.HostIP6Prefix.Len)
	buf.EncodeBool(m.HostIP4GwSet)
	buf.EncodeBytes(m.HostIP4Gw[:], 4)
	buf.EncodeBool(m.HostIP6GwSet)
	buf.EncodeBytes(m.HostIP6Gw[:], 16)
	buf.EncodeUint32(uint32(m.TapFlags))
	buf.EncodeBool(m.HostNamespaceSet)
	buf.EncodeString(m.HostNamespace, 64)
	buf.EncodeBool(m.HostIfNameSet)
	buf.EncodeString(m.HostIfName, 64)
	buf.EncodeBool(m.HostBridgeSet)
	buf.EncodeString(m.HostBridge, 64)
	buf.EncodeString(m.Tag, 0)
	return buf.Bytes(), nil
}
func (m *TapCreateV2) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ID = buf.DecodeUint32()
	m.UseRandomMac = buf.DecodeBool()
	copy(m.MacAddress[:], buf.DecodeBytes(6))
	m.NumRxQueues = buf.DecodeUint8()
	m.TxRingSz = buf.DecodeUint16()
	m.RxRingSz = buf.DecodeUint16()
	m.HostMtuSet = buf.DecodeBool()
	m.HostMtuSize = buf.DecodeUint32()
	m.HostMacAddrSet = buf.DecodeBool()
	copy(m.HostMacAddr[:], buf.DecodeBytes(6))
	m.HostIP4PrefixSet = buf.DecodeBool()
	copy(m.HostIP4Prefix.Address[:], buf.DecodeBytes(4))
	m.HostIP4Prefix.Len = buf.DecodeUint8()
	m.HostIP6PrefixSet = buf.DecodeBool()
	copy(m.HostIP6Prefix.Address[:], buf.DecodeBytes(16))
	m.HostIP6Prefix.Len = buf.DecodeUint8()
	m.HostIP4GwSet = buf.DecodeBool()
	copy(m.HostIP4Gw[:], buf.DecodeBytes(4))
	m.HostIP6GwSet = buf.DecodeBool()
	copy(m.HostIP6Gw[:], buf.DecodeBytes(16))
	m.TapFlags = TapFlags(buf.DecodeUint32())
	m.HostNamespaceSet = buf.DecodeBool()
	m.HostNamespace = buf.DecodeString(64)
	m.HostIfNameSet = buf.DecodeBool()
	m.HostIfName = buf.DecodeString(64)
	m.HostBridgeSet = buf.DecodeBool()
	m.HostBridge = buf.DecodeString(64)
	m.Tag = buf.DecodeString(0)
	return nil
}

// TapCreateV2Reply defines message 'tap_create_v2_reply'.
type TapCreateV2Reply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *TapCreateV2Reply) Reset()               { *m = TapCreateV2Reply{} }
func (*TapCreateV2Reply) GetMessageName() string { return "tap_create_v2_reply" }
func (*TapCreateV2Reply) GetCrcString() string   { return "5383d31f" }
func (*TapCreateV2Reply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *TapCreateV2Reply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *TapCreateV2Reply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *TapCreateV2Reply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// TapCreateV3 defines message 'tap_create_v3'.
type TapCreateV3 struct {
	ID               uint32                        `binapi:"u32,name=id,default=4294967295" json:"id,omitempty"`
	UseRandomMac     bool                          `binapi:"bool,name=use_random_mac,default=true" json:"use_random_mac,omitempty"`
	MacAddress       ethernet_types.MacAddress     `binapi:"mac_address,name=mac_address" json:"mac_address,omitempty"`
	NumRxQueues      uint16                        `binapi:"u16,name=num_rx_queues,default=1" json:"num_rx_queues,omitempty"`
	NumTxQueues      uint16                        `binapi:"u16,name=num_tx_queues,default=1" json:"num_tx_queues,omitempty"`
	TxRingSz         uint16                        `binapi:"u16,name=tx_ring_sz,default=256" json:"tx_ring_sz,omitempty"`
	RxRingSz         uint16                        `binapi:"u16,name=rx_ring_sz,default=256" json:"rx_ring_sz,omitempty"`
	HostMtuSet       bool                          `binapi:"bool,name=host_mtu_set" json:"host_mtu_set,omitempty"`
	HostMtuSize      uint32                        `binapi:"u32,name=host_mtu_size" json:"host_mtu_size,omitempty"`
	HostMacAddrSet   bool                          `binapi:"bool,name=host_mac_addr_set" json:"host_mac_addr_set,omitempty"`
	HostMacAddr      ethernet_types.MacAddress     `binapi:"mac_address,name=host_mac_addr" json:"host_mac_addr,omitempty"`
	HostIP4PrefixSet bool                          `binapi:"bool,name=host_ip4_prefix_set" json:"host_ip4_prefix_set,omitempty"`
	HostIP4Prefix    ip_types.IP4AddressWithPrefix `binapi:"ip4_address_with_prefix,name=host_ip4_prefix" json:"host_ip4_prefix,omitempty"`
	HostIP6PrefixSet bool                          `binapi:"bool,name=host_ip6_prefix_set" json:"host_ip6_prefix_set,omitempty"`
	HostIP6Prefix    ip_types.IP6AddressWithPrefix `binapi:"ip6_address_with_prefix,name=host_ip6_prefix" json:"host_ip6_prefix,omitempty"`
	HostIP4GwSet     bool                          `binapi:"bool,name=host_ip4_gw_set" json:"host_ip4_gw_set,omitempty"`
	HostIP4Gw        ip_types.IP4Address           `binapi:"ip4_address,name=host_ip4_gw" json:"host_ip4_gw,omitempty"`
	HostIP6GwSet     bool                          `binapi:"bool,name=host_ip6_gw_set" json:"host_ip6_gw_set,omitempty"`
	HostIP6Gw        ip_types.IP6Address           `binapi:"ip6_address,name=host_ip6_gw" json:"host_ip6_gw,omitempty"`
	TapFlags         TapFlags                      `binapi:"tap_flags,name=tap_flags" json:"tap_flags,omitempty"`
	HostNamespaceSet bool                          `binapi:"bool,name=host_namespace_set" json:"host_namespace_set,omitempty"`
	HostNamespace    string                        `binapi:"string[64],name=host_namespace" json:"host_namespace,omitempty"`
	HostIfNameSet    bool                          `binapi:"bool,name=host_if_name_set" json:"host_if_name_set,omitempty"`
	HostIfName       string                        `binapi:"string[64],name=host_if_name" json:"host_if_name,omitempty"`
	HostBridgeSet    bool                          `binapi:"bool,name=host_bridge_set" json:"host_bridge_set,omitempty"`
	HostBridge       string                        `binapi:"string[64],name=host_bridge" json:"host_bridge,omitempty"`
	Tag              string                        `binapi:"string[],name=tag" json:"tag,omitempty"`
}

func (m *TapCreateV3) Reset()               { *m = TapCreateV3{} }
func (*TapCreateV3) GetMessageName() string { return "tap_create_v3" }
func (*TapCreateV3) GetCrcString() string   { return "3f3fd1df" }
func (*TapCreateV3) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TapCreateV3) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4              // m.ID
	size += 1              // m.UseRandomMac
	size += 1 * 6          // m.MacAddress
	size += 2              // m.NumRxQueues
	size += 2              // m.NumTxQueues
	size += 2              // m.TxRingSz
	size += 2              // m.RxRingSz
	size += 1              // m.HostMtuSet
	size += 4              // m.HostMtuSize
	size += 1              // m.HostMacAddrSet
	size += 1 * 6          // m.HostMacAddr
	size += 1              // m.HostIP4PrefixSet
	size += 1 * 4          // m.HostIP4Prefix.Address
	size += 1              // m.HostIP4Prefix.Len
	size += 1              // m.HostIP6PrefixSet
	size += 1 * 16         // m.HostIP6Prefix.Address
	size += 1              // m.HostIP6Prefix.Len
	size += 1              // m.HostIP4GwSet
	size += 1 * 4          // m.HostIP4Gw
	size += 1              // m.HostIP6GwSet
	size += 1 * 16         // m.HostIP6Gw
	size += 4              // m.TapFlags
	size += 1              // m.HostNamespaceSet
	size += 64             // m.HostNamespace
	size += 1              // m.HostIfNameSet
	size += 64             // m.HostIfName
	size += 1              // m.HostBridgeSet
	size += 64             // m.HostBridge
	size += 4 + len(m.Tag) // m.Tag
	return size
}
func (m *TapCreateV3) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.ID)
	buf.EncodeBool(m.UseRandomMac)
	buf.EncodeBytes(m.MacAddress[:], 6)
	buf.EncodeUint16(m.NumRxQueues)
	buf.EncodeUint16(m.NumTxQueues)
	buf.EncodeUint16(m.TxRingSz)
	buf.EncodeUint16(m.RxRingSz)
	buf.EncodeBool(m.HostMtuSet)
	buf.EncodeUint32(m.HostMtuSize)
	buf.EncodeBool(m.HostMacAddrSet)
	buf.EncodeBytes(m.HostMacAddr[:], 6)
	buf.EncodeBool(m.HostIP4PrefixSet)
	buf.EncodeBytes(m.HostIP4Prefix.Address[:], 4)
	buf.EncodeUint8(m.HostIP4Prefix.Len)
	buf.EncodeBool(m.HostIP6PrefixSet)
	buf.EncodeBytes(m.HostIP6Prefix.Address[:], 16)
	buf.EncodeUint8(m.HostIP6Prefix.Len)
	buf.EncodeBool(m.HostIP4GwSet)
	buf.EncodeBytes(m.HostIP4Gw[:], 4)
	buf.EncodeBool(m.HostIP6GwSet)
	buf.EncodeBytes(m.HostIP6Gw[:], 16)
	buf.EncodeUint32(uint32(m.TapFlags))
	buf.EncodeBool(m.HostNamespaceSet)
	buf.EncodeString(m.HostNamespace, 64)
	buf.EncodeBool(m.HostIfNameSet)
	buf.EncodeString(m.HostIfName, 64)
	buf.EncodeBool(m.HostBridgeSet)
	buf.EncodeString(m.HostBridge, 64)
	buf.EncodeString(m.Tag, 0)
	return buf.Bytes(), nil
}
func (m *TapCreateV3) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ID = buf.DecodeUint32()
	m.UseRandomMac = buf.DecodeBool()
	copy(m.MacAddress[:], buf.DecodeBytes(6))
	m.NumRxQueues = buf.DecodeUint16()
	m.NumTxQueues = buf.DecodeUint16()
	m.TxRingSz = buf.DecodeUint16()
	m.RxRingSz = buf.DecodeUint16()
	m.HostMtuSet = buf.DecodeBool()
	m.HostMtuSize = buf.DecodeUint32()
	m.HostMacAddrSet = buf.DecodeBool()
	copy(m.HostMacAddr[:], buf.DecodeBytes(6))
	m.HostIP4PrefixSet = buf.DecodeBool()
	copy(m.HostIP4Prefix.Address[:], buf.DecodeBytes(4))
	m.HostIP4Prefix.Len = buf.DecodeUint8()
	m.HostIP6PrefixSet = buf.DecodeBool()
	copy(m.HostIP6Prefix.Address[:], buf.DecodeBytes(16))
	m.HostIP6Prefix.Len = buf.DecodeUint8()
	m.HostIP4GwSet = buf.DecodeBool()
	copy(m.HostIP4Gw[:], buf.DecodeBytes(4))
	m.HostIP6GwSet = buf.DecodeBool()
	copy(m.HostIP6Gw[:], buf.DecodeBytes(16))
	m.TapFlags = TapFlags(buf.DecodeUint32())
	m.HostNamespaceSet = buf.DecodeBool()
	m.HostNamespace = buf.DecodeString(64)
	m.HostIfNameSet = buf.DecodeBool()
	m.HostIfName = buf.DecodeString(64)
	m.HostBridgeSet = buf.DecodeBool()
	m.HostBridge = buf.DecodeString(64)
	m.Tag = buf.DecodeString(0)
	return nil
}

// TapCreateV3Reply defines message 'tap_create_v3_reply'.
type TapCreateV3Reply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *TapCreateV3Reply) Reset()               { *m = TapCreateV3Reply{} }
func (*TapCreateV3Reply) GetMessageName() string { return "tap_create_v3_reply" }
func (*TapCreateV3Reply) GetCrcString() string   { return "5383d31f" }
func (*TapCreateV3Reply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *TapCreateV3Reply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *TapCreateV3Reply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *TapCreateV3Reply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// TapDeleteV2 defines message 'tap_delete_v2'.
type TapDeleteV2 struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *TapDeleteV2) Reset()               { *m = TapDeleteV2{} }
func (*TapDeleteV2) GetMessageName() string { return "tap_delete_v2" }
func (*TapDeleteV2) GetCrcString() string   { return "f9e6675e" }
func (*TapDeleteV2) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TapDeleteV2) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *TapDeleteV2) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *TapDeleteV2) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// TapDeleteV2Reply defines message 'tap_delete_v2_reply'.
type TapDeleteV2Reply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *TapDeleteV2Reply) Reset()               { *m = TapDeleteV2Reply{} }
func (*TapDeleteV2Reply) GetMessageName() string { return "tap_delete_v2_reply" }
func (*TapDeleteV2Reply) GetCrcString() string   { return "e8d4e804" }
func (*TapDeleteV2Reply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *TapDeleteV2Reply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *TapDeleteV2Reply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *TapDeleteV2Reply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_tapv2_binapi_init() }
func file_tapv2_binapi_init() {
	api.RegisterMessage((*SwInterfaceTapV2Details)(nil), "sw_interface_tap_v2_details_1e2b2a47")
	api.RegisterMessage((*SwInterfaceTapV2Dump)(nil), "sw_interface_tap_v2_dump_f9e6675e")
	api.RegisterMessage((*TapCreateV2)(nil), "tap_create_v2_2d0d6570")
	api.RegisterMessage((*TapCreateV2Reply)(nil), "tap_create_v2_reply_5383d31f")
	api.RegisterMessage((*TapCreateV3)(nil), "tap_create_v3_3f3fd1df")
	api.RegisterMessage((*TapCreateV3Reply)(nil), "tap_create_v3_reply_5383d31f")
	api.RegisterMessage((*TapDeleteV2)(nil), "tap_delete_v2_f9e6675e")
	api.RegisterMessage((*TapDeleteV2Reply)(nil), "tap_delete_v2_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*SwInterfaceTapV2Details)(nil),
		(*SwInterfaceTapV2Dump)(nil),
		(*TapCreateV2)(nil),
		(*TapCreateV2Reply)(nil),
		(*TapCreateV3)(nil),
		(*TapCreateV3Reply)(nil),
		(*TapDeleteV2)(nil),
		(*TapDeleteV2Reply)(nil),
	}
}
