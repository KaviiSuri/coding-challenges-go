package internal

import "fmt"

type Freqs map[byte]uint32

func AnalyseFreq(data []byte) Freqs {
	f := make(Freqs)

	for _, c := range data {
		f[c] += 1
	}
	return f
}

func PrettyPrentFreqs(f Freqs) string {
	s := "\n"
	for k, v := range f {
		s += fmt.Sprintf("%c: %d\n", k, v)
	}
	return s
}
