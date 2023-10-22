package protocol

import (
	"encoding/binary"
	"io"
	"reflect"
)

func Unmarshal(r io.Reader, args ...interface{}) error {
	for _, arg := range args {
		switch value := arg.(type) {
		case *VarInt:
			{
				val, err := ReadVarInt(r)

				if err != nil {
					return err
				}

				*value = VarInt(val)

				break
			}
		case *string:
			{
				val, err := ReadString(r)

				if err != nil {
					return err
				}

				*value = val

				break
			}
		case *UUID:
			{
				if err := (*value).Decode(r); err != nil {
					return err
				}

				break
			}
		case *[]byte:
			{
				length, err := ReadVarInt(r)

				if err != nil {
					return err
				}

				*value = make([]byte, length)

				if _, err := r.Read(*value); err != nil {
					return err
				}

				break
			}
		default:
			{
				v := reflect.New(reflect.TypeOf(value).Elem())

				if err := binary.Read(r, binary.BigEndian, v.Interface()); err != nil {
					return err
				}

				reflect.ValueOf(value).Elem().Set(v.Elem())

				break
			}
		}
	}

	return nil
}
