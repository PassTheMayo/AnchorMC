package packets

import (
	"fmt"
	"io"

	"github.com/anchormc/anchor/src/protocol"
)

type ServerboundHandshake struct {
	ProtocolVersion protocol.VarInt
	ServerAddress   string
	ServerPort      uint16
	NextState       protocol.VarInt
}

func (p *ServerboundHandshake) EncodePacket(w io.Writer) error {
	if err := protocol.WriteVarInt(w, 0x00); err != nil {
		return err
	}

	return p.MarshalData(w)
}

func (p *ServerboundHandshake) DecodePacket(packet *protocol.PacketBuffer) error {
	if packet.ID != 0x00 {
		return fmt.Errorf("packet ID mismatch (type=ServerboundHandshake, expected=0x00, received=0x%02X)", packet.ID)
	}

	return p.UnmarshalData(packet.Data)
}

func (p ServerboundHandshake) MarshalData(w io.Writer) error {
	return protocol.Marshal(
		w,
		p.ProtocolVersion,
		p.ServerAddress,
		p.ServerPort,
		p.NextState,
	)
}

func (p *ServerboundHandshake) UnmarshalData(r io.Reader) error {
	return protocol.Unmarshal(
		r,
		&p.ProtocolVersion,
		&p.ServerAddress,
		&p.ServerPort,
		&p.NextState,
	)
}

var _ Packet = &ServerboundHandshake{}
