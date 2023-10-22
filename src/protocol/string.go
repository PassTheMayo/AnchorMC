package protocol

import "io"

func ReadString(r io.Reader) (string, error) {
	length, err := ReadVarInt(r)

	if err != nil {
		return "", err
	}

	data := make([]byte, length)

	if _, err := r.Read(data); err != nil {
		return "", err
	}

	return string(data), nil
}

func WriteString(w io.Writer, value string) error {
	if err := WriteVarInt(w, int32(len(value))); err != nil {
		return err
	}

	_, err := w.Write([]byte(value))

	return err
}
