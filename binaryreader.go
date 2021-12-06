package binarycursor

import (
	"encoding/binary"
	"io"
)

type BinaryReader struct {
	r     io.Reader
	Order binary.ByteOrder
}

func NewBinaryReader(r io.Reader) BinaryReader {
	return BinaryReader{
		r:     r,
		Order: binary.LittleEndian,
	}
}

func NewBinaryReaderAt(r io.ReaderAt, pos int64) BinaryReader {
	pr := NewPositionReaderAt(r, pos)
	return NewBinaryReader(&pr)
}

func (br *BinaryReader) Read(p []byte) (n int, err error) {
	return br.r.Read(p)
}

func (br *BinaryReader) FlipOrder() {
	switch br.Order {
	case binary.BigEndian:
		br.Order = binary.LittleEndian
	case binary.LittleEndian:
		br.Order = binary.BigEndian
	}
}

func (br *BinaryReader) ReadUint8() (uint8, error) {
	buf := []byte{0x0}

	n, err := br.r.Read(buf)

	if err != nil {
		return 0, err
	}

	if n != len(buf) {

		return 0, ErrReadWrongSize
	}

	return uint8(buf[0]), err
}

func (br *BinaryReader) ReadUint16() (uint16, error) {
	buf := []byte{0x0, 0x0}

	n, err := br.Read(buf)
	if err != nil {
		return 0, err
	}

	if n != len(buf) {
		return 0, ErrReadWrongSize
	}

	return br.Order.Uint16(buf), err
}

func (br *BinaryReader) ReadUint32() (uint32, error) {
	buf := []byte{0x0, 0x0, 0x0, 0x0}

	n, err := br.Read(buf)
	if err != nil {
		return 0, err
	}

	if n != len(buf) {
		return 0, ErrReadWrongSize
	}

	return br.Order.Uint32(buf), err
}

func (br *BinaryReader) ReadUint64() (uint64, error) {
	buf := []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}

	n, err := br.Read(buf)
	if err != nil {
		return 0, err
	}

	if n != len(buf) {
		return 0, ErrReadWrongSize
	}

	return br.Order.Uint64(buf), err
}

func (br *BinaryReader) ReadNullTerminatedUTF8String() (string, error) {
	data := []byte{}

	for {
		n, err := br.ReadUint8()
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