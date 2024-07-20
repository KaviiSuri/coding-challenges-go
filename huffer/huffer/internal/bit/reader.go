package bit

import (
	"io"
)

type Reader struct {
	r      io.Reader
	buffer byte
	bitPos int
}

func NewBitReader(r io.Reader) Reader {
	return Reader{
		r: r,
	}
}

func (br *Reader) GetNext() (bool, error) {
	if br.bitPos == 0 {
		if err := br.load(); err != nil {
			if err == io.EOF {
				return false, err
			}
			return false, err
		}
	}

	bit := (br.buffer & (1 << (7 - br.bitPos))) != 0
	br.bitPos++

	if br.bitPos == 8 {
		br.bitPos = 0
	}

	return bit, nil
}

func (r *Reader) ReadAllBits() ([]bool, error) {
	var bits []bool

	for {
		bit, err := r.GetNext()
		if err == io.EOF {
			// End of file reached
			break
		}
		if err != nil {
			return nil, err
		}
		bits = append(bits, bit)
	}

	return bits, nil
}

func (br *Reader) load() error {
	var buf [1]byte
	_, err := br.r.Read(buf[:])
	if err != nil {
		if err == io.EOF && br.bitPos == 0 {
			return err // No more data to read
		}
		return err
	}
	br.buffer = buf[0]
	return nil
}
