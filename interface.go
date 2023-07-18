package base_change

import "io"

type BaseConversion interface {
	Encode(src, dst []byte)
	Decode(src, dst []byte)
	EncodeToString(src []byte) string
	DecodeString(src string) ([]byte, error)
	EncodeLength(n int) int
	NewEncoder(charSet string) *io.Writer
	NewDecoder(charSet string) *io.Reader
}
