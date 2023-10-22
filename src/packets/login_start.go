package packets

import (
	"fmt"
	"io"

	"github.com/anchormc/anchor/src/protocol"
)

type ServerboundLoginStart struct {
	Name string
	UUID protocol.UUID
}

func (p *ServerboundLoginStart) EncodePacket(w io.Writer) error {
	if err := protocol.WriteVarInt(w, 0x00); err != nil {
		return err
	}

	return p.MarshalData(w)
}

func (p *ServerboundLoginStart) DecodePacket(packet *protocol.PacketBuffer) error {
	if packet.ID != 0x00 {
		return fmt.Errorf("packet ID mismatch (type=ServerboundLoginStart, expected=0x00, received=0x%02X)", packet.ID)
	}

	return p.UnmarshalData(packet.Data)
}

func (p ServerboundLoginStart) MarshalData(w io.Writer) error {
	return protocol.Marshal(
		w,
		p.Name,
		p.UUID,
	)
}

func (p *ServerboundLoginStart) UnmarshalData(r io.Reader) error {
	return protocol.Unmarshal(
		r,
		&p.Name,
		&p.UUID,
	)
}

var _ Packet = &ServerboundLoginStart{}
