package binarycursor

import (
	"encoding/binary"
	"errors"
	"io"
)

var ErrReadWrongSize = errors.New("Wrong size")
var ErrReaderInvalid = errors.New("Reader is nil")

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

func NewBinaryCursorAt(r io.ReaderAt, pos int64) BinaryCursor {
	pr := NewPositionReaderAt(r, pos)

	return NewBinaryCursor(&pr)
}

func (c *BinaryCursor) Read(p []byte) (n int, err error) {
	return c.r.Read(p)
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

	n, err := c.Read(buf)
	if err != nil {
		return 0, err
	}

	if n != len(buf) {
		return 0, ErrReadWrongSize
	}

	return uint8(buf[0]), err
}

func (c *BinaryCursor) ReadUint16() (uint16, error) {
	buf := []byte{0x0, 0x0}

	n, err := c.Read(buf)
	if err != nil {
		return 0, err
	}

	if n != len(buf) {
		return 0, ErrReadWrongSize
	}

	return c.order.Uint16(buf), err
}

func (c *BinaryCursor) ReadUint32() (uint32, error) {
	buf := []byte{0x0, 0x0, 0x0, 0x0}

	n, err := c.Read(buf)
	if err != nil {
		return 0, err
	}

	if n != len(buf) {
		return 0, ErrReadWrongSize
	}

	return c.order.Uint32(buf), err
}

func (c *BinaryCursor) ReadUint64() (uint64, error) {
	buf := []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}

	n, err := c.Read(buf)
	if err != nil {
		return 0, err
	}

	if n != len(buf) {
		return 0, ErrReadWrongSize
	}

	return c.order.Uint64(buf), err
}

func (c *BinaryCursor) ReadNullTerminatedUTF8String() (string, error) {
	data := []byte{}

	for {
		n, err := c.ReadUint8()
		if err != nil {
			return "", err
		}

		if n == 0 {
			break
		}

		data = append(data, n)
	}

	return string(data), nil
}
