package huffer

import (
	"encoding/binary"
	"io"

	"github.com/KaviiSuri/coding-challenges/huffer/huffer/internal"
	"github.com/KaviiSuri/coding-challenges/huffer/huffer/internal/bit"
)

const (
	MagicNumber = 0x48464D4E // "HFMN" in ASCII
)

func Encode(w io.Writer, data []byte) error {
	f := internal.AnalyseFreq(data)
	hh := internal.NewHuffmanHeapFromFreq(f)
	root := hh.BuildTree()
	if root == nil {
		return nil
	}
	code := root.BuildCode()

	if err := writeHeader(w, f); err != nil {
		return err
	}
	bw := bit.NewBitWriter(w)

	for _, ch := range data {
		bits := code[ch]
		for _, bit := range bits {
			if err := bw.WriteBit(bit); err != nil {
				return err
			}
		}
	}

	// Flush any remaining bits
	if err := bw.Flush(); err != nil {
		return err
	}

	return nil
}

func writeHeader(w io.Writer, f internal.Freqs) error {
	if err := binary.Write(w, binary.BigEndian, uint32(MagicNumber)); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, uint32(len(f))); err != nil {
		return err
	}

	// Write character frequencies
	for ch, freq := range f {
		if err := binary.Write(w, binary.BigEndian, ch); err != nil {
			return err
		}
		if err := binary.Write(w, binary.BigEndian, freq); err != nil {
			return err
		}
	}

	return nil
}
