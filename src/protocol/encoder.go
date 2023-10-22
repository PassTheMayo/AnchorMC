package protocol

import (
	"encoding/binary"
	"io"
)

func Marshal(w io.Writer, args ...interface{}) error {
	for _, arg := range args {
		switch value := arg.(type) {
		case VarInt:
			{
				if err := WriteVarInt(w, int32(value)); err != nil {
					return err
				}

				break
			}
		case string:
			{
				if err := WriteString(w, value); err != nil {
					return err
				}

				break
			}
		case UUID:
			{
				if err := value.Encode(w); err != nil {
					return err
				}

				break
			}
		case []byte:
			{
				if err := WriteVarInt(w, int32(len(value))); err != nil {
					return err
				}

				if _, err := w.Write(value); err != nil {
					return err
				}

				break
			}
		default:
			{
				if err := binary.Write(w, binary.BigEndian, value); err != nil {
					return err
				}

				break
			}
		}
	}

	return nil
}
