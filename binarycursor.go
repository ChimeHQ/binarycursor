package binarycursor

import (
	"encoding/binary"
	"errors"
	"io"
)

var ErrReadWrongSize = errors.New("Wrong size")

type BinaryCursor struct {
	r     io.Reader
	order binary.ByteOrder
}

func NewBinaryCursor(r io.Reader) BinaryCursor {
	return BinaryCursor{
		r:     r,
		order: binary.LittleEndian,
	}
}

func (c *BinaryCursor) FlipOrder() {
	switch c.order {
	case binary.BigEndian:
		c.order = binary.LittleEndian
	case binary.LittleEndian:
		c.order = binary.BigEndian
	}
}

func (c *BinaryCursor) ReadUint8() (uint8, error) {
	buf := []byte{0x0}

	n, err := c.r.Read(buf)
	if err != nil {
		return 0, err
	}

	if n != 1 {
		return 0, ErrReadWrongSize
	}

	return uint8(buf[0]), err
}

func (c *BinaryCursor) ReadUint16() (uint16, error) {
	buf := []byte{0x0, 0x0}

	n, err := c.r.Read(buf)
	if err != nil {
		return 0, err
	}

	if n != 2 {
		return 0, ErrReadWrongSize
	}

	return c.order.Uint16(buf), err
}