package packets

import (
	"io"

	"github.com/anchormc/anchor/src/protocol"
)

type ClientboundEncryptionRequest struct {
	ServerID    string
	PublicKey   []byte
	VerifyToken []byte
}

func (p *ClientboundEncryptionRequest) EncodePacket(w io.Writer) error {
	if err := protocol.WriteVarInt(w, 0x01); err != nil {
		return err
	}

	return p.MarshalData(w)
}

func (p *ClientboundEncryptionRequest) DecodePacket(packet *protocol.PacketBuffer) error {
	return ErrClientboundDecode
}

func (p ClientboundEncryptionRequest) MarshalData(w io.Writer) error {
	return protocol.Marshal(
		w,
		p.ServerID,
		p.PublicKey,
		p.VerifyToken,
	)
}

func (p *ClientboundEncryptionRequest) UnmarshalData(r io.Reader) error {
	return protocol.Unmarshal(
		r,
		&p.ServerID,
		&p.PublicKey,
		&p.VerifyToken,
	)
}

var _ Packet = &ClientboundEncryptionRequest{}
