package packets

import (
	"fmt"
	"io"

	"github.com/anchormc/anchor/src/protocol"
)

type ServerboundStatusRequest struct{}

func (p *ServerboundStatusRequest) EncodePacket(w io.Writer) error {
	if err := protocol.WriteVarInt(w, 0x00); err != nil {
		return err
	}

	return p.MarshalData(w)
}

func (p *ServerboundStatusRequest) DecodePacket(packet *protocol.PacketBuffer) error {
	if packet.ID != 0x00 {
		return fmt.Errorf("packet ID mismatch (type=ServerboundStatusRequest, expected=0x00, received=0x%02X)", packet.ID)
	}

	return p.UnmarshalData(packet.Data)
}

func (p ServerboundStatusRequest) MarshalData(w io.Writer) error {
	return nil
}

func (p *ServerboundStatusRequest) UnmarshalData(r io.Reader) error {
	return nil
}

var _ Packet = &ServerboundStatusRequest{}
