package packets

import (
	"fmt"
	"io"

	"github.com/anchormc/anchor/src/protocol"
)

type BidirectionalPingPong struct {
	Payload int64
}

func (p *BidirectionalPingPong) EncodePacket(w io.Writer) error {
	if err := protocol.WriteVarInt(w, 0x01); err != nil {
		return err
	}

	return p.MarshalData(w)
}

func (p *BidirectionalPingPong) DecodePacket(packet *protocol.PacketBuffer) error {
	if packet.ID != 0x01 {
		return fmt.Errorf("packet ID mismatch (type=BidirectionalPingPong, expected=0x01, received=0x%02X)", packet.ID)
	}

	return p.UnmarshalData(packet.Data)
}

func (p BidirectionalPingPong) MarshalData(w io.Writer) error {
	return protocol.Marshal(
		w,
		p.Payload,
	)
}

func (p *BidirectionalPingPong) UnmarshalData(r io.Reader) error {
	return protocol.Unmarshal(
		r,
		&p.Payload,
	)
}

var _ Packet = &BidirectionalPingPong{}
