package packets

import (
	"errors"
	"io"

	"github.com/anchormc/anchor/src/protocol"
)

var (
	ErrClientboundDecode = errors.New("attempted to decode a clientbound packet")
)

type Packet interface {
	EncodePacket(io.Writer) error
	DecodePacket(*protocol.PacketBuffer) error
	MarshalData(io.Writer) error
	UnmarshalData(io.Reader) error
}
