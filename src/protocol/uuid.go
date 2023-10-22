package protocol

import (
	"encoding/binary"
	"fmt"
	"io"
)

type UUID struct {
	High uint64
	Low  uint64
}

func (u UUID) String() string {
	return fmt.Sprintf("%X%X", u.High, u.Low)
}

func (u UUID) Encode(w io.Writer) error {
	if err := binary.Write(w, binary.BigEndian, u.High); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, u.Low); err != nil {
		return err
	}

	return nil
}

func (u *UUID) Decode(r io.Reader) error {
	if err := binary.Read(r, binary.BigEndian, &u.High); err != nil {
		return err
	}

	if err := binary.Read(r, binary.BigEndian, &u.Low); err != nil {
		return err
	}

	return nil
}
