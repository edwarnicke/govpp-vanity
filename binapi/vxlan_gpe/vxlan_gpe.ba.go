// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.5.0-dev
//  VPP:              22.02-release
// source: /usr/share/vpp/api/core/vxlan_gpe.api.json

// Package vxlan_gpe contains generated bindings for API file vxlan_gpe.api.
//
// Contents:
//  10 messages
//
package vxlan_gpe

import (
	api "go.fd.io/govpp/api"
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
	APIFile    = "vxlan_gpe"
	APIVersion = "2.1.0"
	VersionCrc = 0x3bc06278
)

// SwInterfaceSetVxlanGpeBypass defines message 'sw_interface_set_vxlan_gpe_bypass'.
type SwInterfaceSetVxlanGpeBypass struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	IsIPv6    bool                           `binapi:"bool,name=is_ipv6" json:"is_ipv6,omitempty"`
	Enable    bool                           `binapi:"bool,name=enable,default=true" json:"enable,omitempty"`
}

func (m *SwInterfaceSetVxlanGpeBypass) Reset() { *m = SwInterfaceSetVxlanGpeBypass{} }
func (*SwInterfaceSetVxlanGpeBypass) GetMessageName() string {
	return "sw_interface_set_vxlan_gpe_bypass"
}
func (*SwInterfaceSetVxlanGpeBypass) GetCrcString() string { return "65247409" }
func (*SwInterfaceSetVxlanGpeBypass) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SwInterfaceSetVxlanGpeBypass) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 1 // m.IsIPv6
	size += 1 // m.Enable
	return size
}
func (m *SwInterfaceSetVxlanGpeBypass) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBool(m.IsIPv6)
	buf.EncodeBool(m.Enable)
	return buf.Bytes(), nil
}
func (m *SwInterfaceSetVxlanGpeBypass) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.IsIPv6 = buf.DecodeBool()
	m.Enable = buf.DecodeBool()
	return nil
}

// SwInterfaceSetVxlanGpeBypassReply defines message 'sw_interface_set_vxlan_gpe_bypass_reply'.
type SwInterfaceSetVxlanGpeBypassReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SwInterfaceSetVxlanGpeBypassReply) Reset() { *m = SwInterfaceSetVxlanGpeBypassReply{} }
func (*SwInterfaceSetVxlanGpeBypassReply) GetMessageName() string {
	return "sw_interface_set_vxlan_gpe_bypass_reply"
}
func (*SwInterfaceSetVxlanGpeBypassReply) GetCrcString() string { return "e8d4e804" }
func (*SwInterfaceSetVxlanGpeBypassReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SwInterfaceSetVxlanGpeBypassReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SwInterfaceSetVxlanGpeBypassReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SwInterfaceSetVxlanGpeBypassReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// VxlanGpeAddDelTunnel defines message 'vxlan_gpe_add_del_tunnel'.
type VxlanGpeAddDelTunnel struct {
	Local          ip_types.Address               `binapi:"address,name=local" json:"local,omitempty"`
	Remote         ip_types.Address               `binapi:"address,name=remote" json:"remote,omitempty"`
	McastSwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=mcast_sw_if_index" json:"mcast_sw_if_index,omitempty"`
	EncapVrfID     uint32                         `binapi:"u32,name=encap_vrf_id" json:"encap_vrf_id,omitempty"`
	DecapVrfID     uint32                         `binapi:"u32,name=decap_vrf_id" json:"decap_vrf_id,omitempty"`
	Protocol       ip_types.IPProto               `binapi:"ip_proto,name=protocol" json:"protocol,omitempty"`
	Vni            uint32                         `binapi:"u32,name=vni" json:"vni,omitempty"`
	IsAdd          bool                           `binapi:"bool,name=is_add,default=true" json:"is_add,omitempty"`
}

func (m *VxlanGpeAddDelTunnel) Reset()               { *m = VxlanGpeAddDelTunnel{} }
func (*VxlanGpeAddDelTunnel) GetMessageName() string { return "vxlan_gpe_add_del_tunnel" }
func (*VxlanGpeAddDelTunnel) GetCrcString() string   { return "a645b2b0" }
func (*VxlanGpeAddDelTunnel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *VxlanGpeAddDelTunnel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.Local.Af
	size += 1 * 16 // m.Local.Un
	size += 1      // m.Remote.Af
	size += 1 * 16 // m.Remote.Un
	size += 4      // m.McastSwIfIndex
	size += 4      // m.EncapVrfID
	size += 4      // m.DecapVrfID
	size += 1      // m.Protocol
	size += 4      // m.Vni
	size += 1      // m.IsAdd
	return size
}
func (m *VxlanGpeAddDelTunnel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.Local.Af))
	buf.EncodeBytes(m.Local.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(uint8(m.Remote.Af))
	buf.EncodeBytes(m.Remote.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(uint32(m.McastSwIfIndex))
	buf.EncodeUint32(m.EncapVrfID)
	buf.EncodeUint32(m.DecapVrfID)
	buf.EncodeUint8(uint8(m.Protocol))
	buf.EncodeUint32(m.Vni)
	buf.EncodeBool(m.IsAdd)
	return buf.Bytes(), nil
}
func (m *VxlanGpeAddDelTunnel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Local.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Local.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Remote.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Remote.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.McastSwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.EncapVrfID = buf.DecodeUint32()
	m.DecapVrfID = buf.DecodeUint32()
	m.Protocol = ip_types.IPProto(buf.DecodeUint8())
	m.Vni = buf.DecodeUint32()
	m.IsAdd = buf.DecodeBool()
	return nil
}

// VxlanGpeAddDelTunnelReply defines message 'vxlan_gpe_add_del_tunnel_reply'.
type VxlanGpeAddDelTunnelReply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *VxlanGpeAddDelTunnelReply) Reset()               { *m = VxlanGpeAddDelTunnelReply{} }
func (*VxlanGpeAddDelTunnelReply) GetMessageName() string { return "vxlan_gpe_add_del_tunnel_reply" }
func (*VxlanGpeAddDelTunnelReply) GetCrcString() string   { return "5383d31f" }
func (*VxlanGpeAddDelTunnelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *VxlanGpeAddDelTunnelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *VxlanGpeAddDelTunnelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *VxlanGpeAddDelTunnelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// VxlanGpeAddDelTunnelV2 defines message 'vxlan_gpe_add_del_tunnel_v2'.
type VxlanGpeAddDelTunnelV2 struct {
	Local          ip_types.Address               `binapi:"address,name=local" json:"local,omitempty"`
	Remote         ip_types.Address               `binapi:"address,name=remote" json:"remote,omitempty"`
	LocalPort      uint16                         `binapi:"u16,name=local_port" json:"local_port,omitempty"`
	RemotePort     uint16                         `binapi:"u16,name=remote_port" json:"remote_port,omitempty"`
	McastSwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=mcast_sw_if_index" json:"mcast_sw_if_index,omitempty"`
	EncapVrfID     uint32                         `binapi:"u32,name=encap_vrf_id" json:"encap_vrf_id,omitempty"`
	DecapVrfID     uint32                         `binapi:"u32,name=decap_vrf_id" json:"decap_vrf_id,omitempty"`
	Protocol       ip_types.IPProto               `binapi:"ip_proto,name=protocol" json:"protocol,omitempty"`
	Vni            uint32                         `binapi:"u32,name=vni" json:"vni,omitempty"`
	IsAdd          bool                           `binapi:"bool,name=is_add,default=true" json:"is_add,omitempty"`
}

func (m *VxlanGpeAddDelTunnelV2) Reset()               { *m = VxlanGpeAddDelTunnelV2{} }
func (*VxlanGpeAddDelTunnelV2) GetMessageName() string { return "vxlan_gpe_add_del_tunnel_v2" }
func (*VxlanGpeAddDelTunnelV2) GetCrcString() string   { return "d62fdb35" }
func (*VxlanGpeAddDelTunnelV2) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *VxlanGpeAddDelTunnelV2) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.Local.Af
	size += 1 * 16 // m.Local.Un
	size += 1      // m.Remote.Af
	size += 1 * 16 // m.Remote.Un
	size += 2      // m.LocalPort
	size += 2      // m.RemotePort
	size += 4      // m.McastSwIfIndex
	size += 4      // m.EncapVrfID
	size += 4      // m.DecapVrfID
	size += 1      // m.Protocol
	size += 4      // m.Vni
	size += 1      // m.IsAdd
	return size
}
func (m *VxlanGpeAddDelTunnelV2) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.Local.Af))
	buf.EncodeBytes(m.Local.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(uint8(m.Remote.Af))
	buf.EncodeBytes(m.Remote.Un.XXX_UnionData[:], 16)
	buf.EncodeUint16(m.LocalPort)
	buf.EncodeUint16(m.RemotePort)
	buf.EncodeUint32(uint32(m.McastSwIfIndex))
	buf.EncodeUint32(m.EncapVrfID)
	buf.EncodeUint32(m.DecapVrfID)
	buf.EncodeUint8(uint8(m.Protocol))
	buf.EncodeUint32(m.Vni)
	buf.EncodeBool(m.IsAdd)
	return buf.Bytes(), nil
}
func (m *VxlanGpeAddDelTunnelV2) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Local.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Local.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Remote.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Remote.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.LocalPort = buf.DecodeUint16()
	m.RemotePort = buf.DecodeUint16()
	m.McastSwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.EncapVrfID = buf.DecodeUint32()
	m.DecapVrfID = buf.DecodeUint32()
	m.Protocol = ip_types.IPProto(buf.DecodeUint8())
	m.Vni = buf.DecodeUint32()
	m.IsAdd = buf.DecodeBool()
	return nil
}

// VxlanGpeAddDelTunnelV2Reply defines message 'vxlan_gpe_add_del_tunnel_v2_reply'.
type VxlanGpeAddDelTunnelV2Reply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *VxlanGpeAddDelTunnelV2Reply) Reset() { *m = VxlanGpeAddDelTunnelV2Reply{} }
func (*VxlanGpeAddDelTunnelV2Reply) GetMessageName() string {
	return "vxlan_gpe_add_del_tunnel_v2_reply"
}
func (*VxlanGpeAddDelTunnelV2Reply) GetCrcString() string { return "5383d31f" }
func (*VxlanGpeAddDelTunnelV2Reply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *VxlanGpeAddDelTunnelV2Reply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *VxlanGpeAddDelTunnelV2Reply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *VxlanGpeAddDelTunnelV2Reply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// VxlanGpeTunnelDetails defines message 'vxlan_gpe_tunnel_details'.
type VxlanGpeTunnelDetails struct {
	SwIfIndex      interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Local          ip_types.Address               `binapi:"address,name=local" json:"local,omitempty"`
	Remote         ip_types.Address               `binapi:"address,name=remote" json:"remote,omitempty"`
	Vni            uint32                         `binapi:"u32,name=vni" json:"vni,omitempty"`
	Protocol       ip_types.IPProto               `binapi:"ip_proto,name=protocol" json:"protocol,omitempty"`
	McastSwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=mcast_sw_if_index" json:"mcast_sw_if_index,omitempty"`
	EncapVrfID     uint32                         `binapi:"u32,name=encap_vrf_id" json:"encap_vrf_id,omitempty"`
	DecapVrfID     uint32                         `binapi:"u32,name=decap_vrf_id" json:"decap_vrf_id,omitempty"`
	IsIPv6         bool                           `binapi:"bool,name=is_ipv6" json:"is_ipv6,omitempty"`
}

func (m *VxlanGpeTunnelDetails) Reset()               { *m = VxlanGpeTunnelDetails{} }
func (*VxlanGpeTunnelDetails) GetMessageName() string { return "vxlan_gpe_tunnel_details" }
func (*VxlanGpeTunnelDetails) GetCrcString() string   { return "0968fc8b" }
func (*VxlanGpeTunnelDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *VxlanGpeTunnelDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.SwIfIndex
	size += 1      // m.Local.Af
	size += 1 * 16 // m.Local.Un
	size += 1      // m.Remote.Af
	size += 1 * 16 // m.Remote.Un
	size += 4      // m.Vni
	size += 1      // m.Protocol
	size += 4      // m.McastSwIfIndex
	size += 4      // m.EncapVrfID
	size += 4      // m.DecapVrfID
	size += 1      // m.IsIPv6
	return size
}
func (m *VxlanGpeTunnelDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint8(uint8(m.Local.Af))
	buf.EncodeBytes(m.Local.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(uint8(m.Remote.Af))
	buf.EncodeBytes(m.Remote.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(m.Vni)
	buf.EncodeUint8(uint8(m.Protocol))
	buf.EncodeUint32(uint32(m.McastSwIfIndex))
	buf.EncodeUint32(m.EncapVrfID)
	buf.EncodeUint32(m.DecapVrfID)
	buf.EncodeBool(m.IsIPv6)
	return buf.Bytes(), nil
}
func (m *VxlanGpeTunnelDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Local.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Local.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Remote.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Remote.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Vni = buf.DecodeUint32()
	m.Protocol = ip_types.IPProto(buf.DecodeUint8())
	m.McastSwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.EncapVrfID = buf.DecodeUint32()
	m.DecapVrfID = buf.DecodeUint32()
	m.IsIPv6 = buf.DecodeBool()
	return nil
}

// VxlanGpeTunnelDump defines message 'vxlan_gpe_tunnel_dump'.
type VxlanGpeTunnelDump struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *VxlanGpeTunnelDump) Reset()               { *m = VxlanGpeTunnelDump{} }
func (*VxlanGpeTunnelDump) GetMessageName() string { return "vxlan_gpe_tunnel_dump" }
func (*VxlanGpeTunnelDump) GetCrcString() string   { return "f9e6675e" }
func (*VxlanGpeTunnelDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *VxlanGpeTunnelDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *VxlanGpeTunnelDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *VxlanGpeTunnelDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// VxlanGpeTunnelV2Details defines message 'vxlan_gpe_tunnel_v2_details'.
type VxlanGpeTunnelV2Details struct {
	SwIfIndex      interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Local          ip_types.Address               `binapi:"address,name=local" json:"local,omitempty"`
	Remote         ip_types.Address               `binapi:"address,name=remote" json:"remote,omitempty"`
	LocalPort      uint16                         `binapi:"u16,name=local_port" json:"local_port,omitempty"`
	RemotePort     uint16                         `binapi:"u16,name=remote_port" json:"remote_port,omitempty"`
	Vni            uint32                         `binapi:"u32,name=vni" json:"vni,omitempty"`
	Protocol       ip_types.IPProto               `binapi:"ip_proto,name=protocol" json:"protocol,omitempty"`
	McastSwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=mcast_sw_if_index" json:"mcast_sw_if_index,omitempty"`
	EncapVrfID     uint32                         `binapi:"u32,name=encap_vrf_id" json:"encap_vrf_id,omitempty"`
	DecapVrfID     uint32                         `binapi:"u32,name=decap_vrf_id" json:"decap_vrf_id,omitempty"`
	IsIPv6         bool                           `binapi:"bool,name=is_ipv6" json:"is_ipv6,omitempty"`
}

func (m *VxlanGpeTunnelV2Details) Reset()               { *m = VxlanGpeTunnelV2Details{} }
func (*VxlanGpeTunnelV2Details) GetMessageName() string { return "vxlan_gpe_tunnel_v2_details" }
func (*VxlanGpeTunnelV2Details) GetCrcString() string   { return "06be4870" }
func (*VxlanGpeTunnelV2Details) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *VxlanGpeTunnelV2Details) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.SwIfIndex
	size += 1      // m.Local.Af
	size += 1 * 16 // m.Local.Un
	size += 1      // m.Remote.Af
	size += 1 * 16 // m.Remote.Un
	size += 2      // m.LocalPort
	size += 2      // m.RemotePort
	size += 4      // m.Vni
	size += 1      // m.Protocol
	size += 4      // m.McastSwIfIndex
	size += 4      // m.EncapVrfID
	size += 4      // m.DecapVrfID
	size += 1      // m.IsIPv6
	return size
}
func (m *VxlanGpeTunnelV2Details) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint8(uint8(m.Local.Af))
	buf.EncodeBytes(m.Local.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(uint8(m.Remote.Af))
	buf.EncodeBytes(m.Remote.Un.XXX_UnionData[:], 16)
	buf.EncodeUint16(m.LocalPort)
	buf.EncodeUint16(m.RemotePort)
	buf.EncodeUint32(m.Vni)
	buf.EncodeUint8(uint8(m.Protocol))
	buf.EncodeUint32(uint32(m.McastSwIfIndex))
	buf.EncodeUint32(m.EncapVrfID)
	buf.EncodeUint32(m.DecapVrfID)
	buf.EncodeBool(m.IsIPv6)
	return buf.Bytes(), nil
}
func (m *VxlanGpeTunnelV2Details) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Local.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Local.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Remote.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Remote.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.LocalPort = buf.DecodeUint16()
	m.RemotePort = buf.DecodeUint16()
	m.Vni = buf.DecodeUint32()
	m.Protocol = ip_types.IPProto(buf.DecodeUint8())
	m.McastSwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.EncapVrfID = buf.DecodeUint32()
	m.DecapVrfID = buf.DecodeUint32()
	m.IsIPv6 = buf.DecodeBool()
	return nil
}

// VxlanGpeTunnelV2Dump defines message 'vxlan_gpe_tunnel_v2_dump'.
type VxlanGpeTunnelV2Dump struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *VxlanGpeTunnelV2Dump) Reset()               { *m = VxlanGpeTunnelV2Dump{} }
func (*VxlanGpeTunnelV2Dump) GetMessageName() string { return "vxlan_gpe_tunnel_v2_dump" }
func (*VxlanGpeTunnelV2Dump) GetCrcString() string   { return "f9e6675e" }
func (*VxlanGpeTunnelV2Dump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *VxlanGpeTunnelV2Dump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *VxlanGpeTunnelV2Dump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *VxlanGpeTunnelV2Dump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

func init() { file_vxlan_gpe_binapi_init() }
func file_vxlan_gpe_binapi_init() {
	api.RegisterMessage((*SwInterfaceSetVxlanGpeBypass)(nil), "sw_interface_set_vxlan_gpe_bypass_65247409")
	api.RegisterMessage((*SwInterfaceSetVxlanGpeBypassReply)(nil), "sw_interface_set_vxlan_gpe_bypass_reply_e8d4e804")
	api.RegisterMessage((*VxlanGpeAddDelTunnel)(nil), "vxlan_gpe_add_del_tunnel_a645b2b0")
	api.RegisterMessage((*VxlanGpeAddDelTunnelReply)(nil), "vxlan_gpe_add_del_tunnel_reply_5383d31f")
	api.RegisterMessage((*VxlanGpeAddDelTunnelV2)(nil), "vxlan_gpe_add_del_tunnel_v2_d62fdb35")
	api.RegisterMessage((*VxlanGpeAddDelTunnelV2Reply)(nil), "vxlan_gpe_add_del_tunnel_v2_reply_5383d31f")
	api.RegisterMessage((*VxlanGpeTunnelDetails)(nil), "vxlan_gpe_tunnel_details_0968fc8b")
	api.RegisterMessage((*VxlanGpeTunnelDump)(nil), "vxlan_gpe_tunnel_dump_f9e6675e")
	api.RegisterMessage((*VxlanGpeTunnelV2Details)(nil), "vxlan_gpe_tunnel_v2_details_06be4870")
	api.RegisterMessage((*VxlanGpeTunnelV2Dump)(nil), "vxlan_gpe_tunnel_v2_dump_f9e6675e")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*SwInterfaceSetVxlanGpeBypass)(nil),
		(*SwInterfaceSetVxlanGpeBypassReply)(nil),
		(*VxlanGpeAddDelTunnel)(nil),
		(*VxlanGpeAddDelTunnelReply)(nil),
		(*VxlanGpeAddDelTunnelV2)(nil),
		(*VxlanGpeAddDelTunnelV2Reply)(nil),
		(*VxlanGpeTunnelDetails)(nil),
		(*VxlanGpeTunnelDump)(nil),
		(*VxlanGpeTunnelV2Details)(nil),
		(*VxlanGpeTunnelV2Dump)(nil),
	}
}
