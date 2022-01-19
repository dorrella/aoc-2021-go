package main

type PacketType int

const (
	PKT_LIT PacketType = iota
	PKT_OP
)

type Header struct {
	Version uint
	Id      uint
}

type Packet struct {
	Header  Header
	Type    PacketType
	Literal uint
	Subs    []*Packet
}

func NewPacket(h Header, t PacketType) Packet {
	return Packet{
		Header:  h,
		Type:    t,
		Literal: 0,
		Subs:    make([]*Packet, 0),
	}
}
