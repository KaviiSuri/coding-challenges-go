package huffer

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/KaviiSuri/coding-challenges/huffer/huffer/internal"
)

func TestTemporary(t *testing.T) {
	str := "aabbcc"
	f := internal.AnalyseFreq([]byte(str))
	_ = getCodeFor(f)

	w := new(bytes.Buffer)
	if err := writeHeader(w, f); err != nil {
		t.Fatal(err)
	}
	if err := ensureMagicNumber(w); err != nil {
		t.Fatal(err)
	}
	decodedF, err := decodeHeader(w)
	if err != nil {
		t.Fatal(err)
	}
	_ = getCodeFor(decodedF)
}

func getCodeFor(f internal.Freqs) internal.Code {
	hh := internal.NewHuffmanHeapFromFreq(f)
	root := hh.BuildTree()
	code := root.BuildCode()
	return code
}

func TestCodeAndFreqWork(t *testing.T) {
	str := "aabbcc"
	f := internal.AnalyseFreq([]byte(str))
	code := getCodeFor(f)

	w := new(bytes.Buffer)
	if err := writeHeader(w, f); err != nil {
		t.Fatal(err)
	}
	if err := ensureMagicNumber(w); err != nil {
		t.Fatal(err)
	}
	decodedF, err := decodeHeader(w)
	if err != nil {
		t.Fatal(err)
	}
	decodedCode := getCodeFor(decodedF)

	if !reflect.DeepEqual(f, decodedF) {
		t.Fatal(
			"Expected: ",
			internal.PrettyPrentFreqs(f),
			"Got: ",
			internal.PrettyPrentFreqs(decodedF),
		)

	}

	if !reflect.DeepEqual(code, decodedCode) {
		t.Fatal(
			"Expected: ",
			internal.PrettyPrentCode(code),
			"Got: ",
			internal.PrettyPrentCode(decodedCode),
		)
	}
}
