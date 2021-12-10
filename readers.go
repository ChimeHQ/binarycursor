package binarycursor

import (
	"io"
)

type CountingReader struct {
	r      io.Reader
	Offset int64
}

func NewCountingReader(r io.Reader) CountingReader {
	return CountingReader{
		r:      r,
		Offset: 0,
	}
}

func (c *CountingReader) Read(p []byte) (n int, err error) {
	n, err = c.r.Read(p)

	c.Offset += int64(n)

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

func (pr *PositionReaderAt) Read(p []byte) (n int, err error) {
	n, err = pr.r.ReadAt(p, pr.pos)

	pr.pos += int64(n)

	return n, err
}
