package protocol

import (
	"errors"
	"io"
)

const (
	segmentBits byte = 0x7F
	continueBit byte = 0x80
)

type VarInt int32

func ReadVarInt(r io.Reader) (int32, error) {
	var (
		value    int32 = 0
		position int   = 0
	)

	for {
		data := make([]byte, 1)

		if _, err := r.Read(data); err != nil {
			return 0, err
		}

		value |= int32(data[0]&segmentBits) << position

		if (data[0] & continueBit) == 0 {
			break
		}

		position += 7

		if position >= 32 {
			return 0, errors.New("varint is too big")
		}
	}

	return value, nil
}

func WriteVarInt(w io.Writer, value int32) error {
	for {
		if (byte(value) & ^segmentBits) == 0 {
			_, err := w.Write([]byte{byte(value)})

			return err
		}

		w.Write([]byte{(byte(value) & segmentBits) | continueBit})

		value = int32(uint32(value) >> 7)
	}
}
