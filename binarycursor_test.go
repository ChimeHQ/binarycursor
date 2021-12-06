package binarycursor

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCursorReadConsecutiveUint8s(t *testing.T) {
	data := []byte{0x08, 0x09, 0x0a, 0x0b}
	reader := bytes.NewReader(data)

	c := NewBinaryCursor(reader)

	v, err := c.ReadUint8()
	if assert.Nil(t, err) == false {
		return
	}

	assert.Equal(t, uint8(8), v)

	v, err = c.ReadUint8()
	if assert.Nil(t, err) == false {
		return
	}

	assert.Equal(t, uint8(9), v)

	v, err = c.ReadUint8()
	if assert.Nil(t, err) == false {
		return
	}

	assert.Equal(t, uint8(10), v)

	v, err = c.ReadUint8()
	if assert.Nil(t, err) == false {
		return
	}

	assert.Equal(t, uint8(11), v)
}

func TestReadUint32(t *testing.T) {
	data := []byte{0x08, 0x09, 0x10, 0x11}
	reader := bytes.NewReader(data)

	c := NewBinaryCursor(reader)

	v, err := c.ReadUint32()
	if assert.Nil(t, err) == false {
		return
	}

	assert.Equal(t, uint32(0x11100908), v)

	reader = bytes.NewReader(data)

	c = NewBinaryCursor(reader)
	c.FlipOrder()

	v, err = c.ReadUint32()
	if assert.Nil(t, err) == false {
		return
	}

	assert.Equal(t, uint32(0x08091011), v)
}

func TestReadUint364(t *testing.T) {
	data := []byte{0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15}
	reader := bytes.NewReader(data)

	c := NewBinaryCursor(reader)

	v, err := c.ReadUint64()
	if assert.Nil(t, err) == false {
		return
	}

	assert.Equal(t, uint64(0x1514131211100908), v)

	reader = bytes.NewReader(data)

	c = NewBinaryCursor(reader)
	c.FlipOrder()

	v, err = c.ReadUint64()
	if assert.Nil(t, err) == false {
		return
	}

	assert.Equal(t, uint64(0x0809101112131415), v)
}

func TestReadString(t *testing.T) {
	data := []byte{0x61, 0x62, 0x63, 0x00}
	reader := bytes.NewReader(data)

	c := NewBinaryCursor(reader)

	s, err := c.ReadNullTerminatedUTF8String()
	if assert.Nil(t, err) == false {
		return
	}

	assert.Equal(t, "abc", s)
}
