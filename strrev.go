package piscine

import (
	"unicode/utf8"
)

func StrRev(s string) string {
	reversed := make([]byte, len(s))
	i := 0

	for len(s) > 0 {
		r, size := utf8.DecodeLastRuneInString(s)
		s = s[:len(s)-size]
		i += utf8.EncodeRune(reversed[i:], r)
	}

	return string(reversed)
}
