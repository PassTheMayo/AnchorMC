package packets

import (
	"encoding/json"
	"io"

	"github.com/anchormc/anchor/src/protocol"
	"github.com/anchormc/anchor/src/types"
)

type ClientboundStatusResponse struct {
	Payload types.StatusResponse
}

func (p *ClientboundStatusResponse) EncodePacket(w io.Writer) error {
	if err := protocol.WriteVarInt(w, 0x00); err != nil {
		return err
	}

	return p.MarshalData(w)
}

func (p *ClientboundStatusResponse) DecodePacket(packet *protocol.PacketBuffer) error {
	return ErrClientboundDecode
}

func (p ClientboundStatusResponse) MarshalData(w io.Writer) error {
	data, err := json.Marshal(p.Payload)

	if err != nil {
		return err
	}

	return protocol.Marshal(w, string(data))
}

func (p *ClientboundStatusResponse) UnmarshalData(r io.Reader) error {
	var payload string

	if err := protocol.Unmarshal(r, &payload); err != nil {
		return err
	}

	return json.Unmarshal([]byte(payload), &p.Payload)
}

var _ Packet = &ClientboundStatusResponse{}
