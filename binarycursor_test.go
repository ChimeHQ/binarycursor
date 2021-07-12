package binarycursor

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadUint8(t *testing.T) {
	data := []byte{0x08, 0x09, 0x10, 0x11}
	reader := bytes.NewReader(data)

	c := NewBinaryCursor(reader)

	v, err := c.ReadUint8()
	if assert.Nil(t, err) == false {
		return
	}

	assert.Equal(t, uint8(8), v)
}

func TestReadUint16(t *testing.T) {
	data := []byte{0x08, 0x09, 0x10, 0x11}
	reader := bytes.NewReader(data)

	c := NewBinaryCursor(reader)

	v, err := c.ReadUint16()
	if assert.Nil(t, err) == false {
		return
	}

	assert.Equal(t, uint16(0x0908), v)

	c.FlipOrder()

	v, err = c.ReadUint16()
	if assert.Nil(t, err) == false {
		return
	}

	assert.Equal(t, uint16(0x1011), v)
}
