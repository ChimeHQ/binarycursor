package binarycursor

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadUint8(t *testing.T) {
	data := []byte{0x08, 0x09, 0x10, 0x11}
	reader := bytes.NewReader(data)

	br := NewBinaryReader(reader)

	v, err := br.ReadUint8()
	if assert.Nil(t, err) == false {
		return
	}

	assert.Equal(t, uint8(8), v)
	assert.Equal(t, int64(1), br.Offset())
}

func TestReadUint16(t *testing.T) {
	data := []byte{0x08, 0x09}
	reader := bytes.NewReader(data)

	br := NewBinaryReader(reader)

	v, err := br.ReadUint16()
	if assert.Nil(t, err) == false {
		return
	}

	assert.Equal(t, uint16(0x0908), v)
	assert.Equal(t, int64(2), br.Offset())

	reader = bytes.NewReader(data)

	br = NewBinaryReader(reader)
	br.FlipOrder()

	v, err = br.ReadUint16()
	if assert.Nil(t, err) == false {
		return
	}

	assert.Equal(t, uint16(0x0809), v)
	assert.Equal(t, int64(2), br.Offset())
}