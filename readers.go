package binarycursor

import (
	"io"
)

type CountingReader struct {
	r     io.Reader
	count int64
}

func NewCountingReader(r io.Reader) CountingReader {
	return CountingReader{
		r:     r,
		count: 0,
	}
}

func (c *CountingReader) Read(p []byte) (n int, err error) {
	n, err = c.Read(p)

	c.count += int64(n)

	return n, err
}

type PositionReaderAt struct {
	r   io.ReaderAt
	pos int64
}

func NewPositionReaderAt(r io.ReaderAt, pos int64) PositionReaderAt {
	return PositionReaderAt{
		r:   r,
		pos: pos,
	}
}

func (c *PositionReaderAt) Read(p []byte) (n int, err error) {
	n, err = c.r.ReadAt(p, c.pos)

	c.pos += int64(n)

	return n, err
}
