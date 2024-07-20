package bit

import "io"

type Writer struct {
	w        io.Writer
	buffer   byte
	bitCount int
}

func NewBitWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

func (bw *Writer) WriteBit(bit bool) error {
	if bit {
		bw.buffer |= (1 << (7 - bw.bitCount))
	}
	bw.bitCount++
	if bw.bitCount == 8 {
		if err := bw.Flush(); err != nil {
			return err
		}
	}

	return nil
}

func (bw *Writer) Flush() error {
	if bw.bitCount > 0 {
		if _, err := bw.w.Write([]byte{bw.buffer}); err != nil {
			return err
		}
		bw.buffer = 0
		bw.bitCount = 0
	}
	return nil
}
