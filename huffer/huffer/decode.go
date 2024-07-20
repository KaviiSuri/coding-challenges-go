package huffer

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/KaviiSuri/coding-challenges/huffer/huffer/internal"
	"github.com/KaviiSuri/coding-challenges/huffer/huffer/internal/bit"
)

func Decode(r io.Reader) ([]byte, error) {
	var data []byte
	if err := ensureMagicNumber(r); err != nil {
		return data, err
	}
	f, err := decodeHeader(r)
	if err != nil {
		return nil, err
	}

	hh := internal.NewHuffmanHeapFromFreq(f)
	root := hh.BuildTree()
	if root.Left == nil && root.Right == nil {
		return decodeSingleUniqCharCase(r, root, f)
	}
	cur := root

	br := bit.NewBitReader(r)

	for {
		bit, err := br.GetNext()
		if cur.IsLeaf() {
			data = append(data, cur.Ch)
			cur = root
		}
		if err == io.EOF {
			// End of file reached
			break
		}
		if err != nil {
			return nil, err
		}
		if bit {
			if cur.Right == nil {
				return nil, errors.New("invalid code")
			}
			cur = cur.Right
		} else {
			if cur.Left == nil {
				return nil, errors.New("invalid code")
			}
			cur = cur.Left
		}
	}

	return data, nil
}

func ensureMagicNumber(r io.Reader) error {
	var magicNum uint32
	if err := binary.Read(r, binary.BigEndian, &magicNum); err != nil {
		if err == io.EOF {
			return errors.New("file is too short to contain a valid header")
		}
		return err
	}
	if magicNum != MagicNumber {
		return errors.New("file format invalid: magic number mismatch")
	}
	return nil
}

func decodeHeader(r io.Reader) (internal.Freqs, error) {
	var n uint32
	if err := binary.Read(r, binary.BigEndian, &n); err != nil {
		if err == io.EOF {
			return nil, errors.New("file is too short to contain a valid header")
		}
		return nil, fmt.Errorf("failed to read size: %w", err)
	}

	f := make(internal.Freqs)

	for i := uint32(0); i < n; i++ {
		var ch byte
		var freq uint32

		if err := binary.Read(r, binary.BigEndian, &ch); err != nil {
			if err == io.EOF {
				return nil, errors.New("file is too short to contain a valid header")
			}
			return nil, fmt.Errorf("failed to read ch: %w", err)
		}
		if err := binary.Read(r, binary.BigEndian, &freq); err != nil {
			if err == io.EOF {
				return nil, errors.New("file is too short to contain a valid header")
			}
			return nil, fmt.Errorf("failed to read freq: %w", err)
		}
		f[ch] = freq
	}

	return f, nil
}

func decodeSingleUniqCharCase(r io.Reader, root *internal.Node, f internal.Freqs) ([]byte, error) {
	var singleChar byte
	for b := range f {
		singleChar = b
		break
	}
	// Read the entire stream as a repetition of singleChar
	data := make([]byte, 0)
	br := bit.NewBitReader(r)
	for {
		_, err := br.GetNext()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		data = append(data, singleChar)
	}
	return data, nil
}
