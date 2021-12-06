package binarycursor

func (c *BinaryCursor) ReadUleb128() (uint64, error) {
	return c.br.ReadUleb128()
}

func (c *BinaryCursor) ReadSleb128() (int64, error) {
	return c.br.ReadSleb128()
}

func (br *BinaryReader) ReadUleb128() (uint64, error) {
	var u uint64
	var n int64

	for n = 0; n < 10; n++ {
		value, err := br.ReadUint8()
		if err != nil {
			return 0, err
		}

		u |= uint64(value&0x7f) << (7 * n)
		if value&0x80 == 0 {
			break
		}
	}

	return u, nil
}

func (br *BinaryReader) ReadSleb128() (int64, error) {
	var s int64
	var n int64

	for n = 0; n < 10; n++ {
		value, err := br.ReadUint8()
		if err != nil {
			return 0, err
		}

		s |= int64(value&0x7f) << (7 * n)
		if value&0x80 == 0 {
			if value&0x40 != 0 {
				s |= ^0 << (7 * (n + 1))
			}

			break
		}
	}

	return s, nil
}
