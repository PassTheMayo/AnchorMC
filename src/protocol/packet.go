package protocol

import (
	"bytes"
	"io"
)

type PacketBuffer struct {
	ID   int32
	Data *bytes.Buffer
}

func NewPacketBuffer(id int32) *PacketBuffer {
	return &PacketBuffer{
		ID:   id,
		Data: &bytes.Buffer{},
	}
}

func (p PacketBuffer) Write(data []byte) (int, error) {
	return p.Data.Write(data)
}

var _ io.Writer = &PacketBuffer{}
