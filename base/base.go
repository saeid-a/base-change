package base

import (
	"github.com/tunabay/go-bitarray"
	"math"
)

// Encode encodes src into EncodedLen(len(src)) bytes of dst.
func (enc *Converter) Encode(src, dst []byte) {
	if len(src) == 0 {
		return
	}
	base := uint16(enc.base - 1)
	basebits := int(math.Ceil(math.Log2(float64(base))))

	di := 0
	cb := 0
	offset := 0
	bb := bitarray.NewBufferFromByteSlice(src)
	srcLen := GetLen(src)
	// Encode the current value
	for offset < GetLen(src)-srcLen%basebits {
		b := int(bb.BitArrayAt(offset, basebits).ToUint64())
		offset += basebits
		dst[di] = enc.charset[b]
		di++
		cb++
	}

	// if there are any remaining values that was left
	if srcLen%basebits != 0 {
		b := int(bb.BitArrayAt(offset, srcLen%basebits).ToUint64())
		offset += srcLen % basebits
		dst[di] = enc.charset[b]
	}
}
func (enc *Converter) Decode(src, dst []byte) {
	if len(src) == 0 {
		return
	}
	base := uint16(enc.base - 1)
	basebits := int(math.Ceil(math.Log2(float64(base))))

	dstOffset := 0
	srcOffsetIndex := 0
	dstBits := int(math.Ceil(float64(len(src)) * float64(basebits) / 8.0))
	if dst == nil {
		dst = make([]byte, dstBits)
	}
	tmp := bitarray.NewBuffer(8)

	bb := bitarray.NewBufferFromByteSlice(dst)
	// Encode the current value
	for srcOffsetIndex < len(src) {
		bb.PutByteAt(dstOffset, enc.decodeMap[src[srcOffsetIndex]]<<(8-basebits))
		dstOffset += basebits
		srcOffsetIndex++
	}

	// if there are any remaining values that was left
	if GetLen(dst)%basebits != 0 {
		remain := GetLen(dst) - dstOffset

		tmp.PutByteAt(0, enc.decodeMap[src[srcOffsetIndex]]<<(8-remain))
		bb.PutBitArrayAt(dstOffset, tmp.Slice(0, remain))
	}

	dst = dst[:len(dst)-1]
}

type Converter struct {
	charset   string
	decodeMap []byte
	padding   bool
	padChar   rune
	base      uint8
}

// New init encoder alphabet
func New(charSet string) *Converter {

	return &Converter{
		charset:   charSet,
		decodeMap: generateDecodeMap(charSet),
		padding:   true,
		padChar:   '=',
		base:      64,
	}
}

func (enc *Converter) encodeLength(n int) int {
	return int(math.Ceil(math.Log(2) / (math.Log(float64(enc.base))) * float64(n*8)))
}
func logb(n float64, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

// WithPadding changes the padding character to the specified rune.
func (enc *Converter) WithPadding(padding rune) *Converter {
	if padding == '\n' || padding == '\r' || padding == 0xFF {
		panic("invalid padding")
	}
	enc.padChar = padding
	return enc
}

// WithCharset changes the encoding alphabet to the provided string.
func (enc *Converter) WithCharset(padding rune) *Converter {
	if padding == '\n' || padding == '\r' || padding == 0xFF {
		panic("invalid padding")
	}
	enc.padChar = padding
	return enc
}

func generateDecodeMap(charMap string) []byte {
	m := make([]byte, 255)
	cm := make(map[rune]int, 255)
	for i, i2 := range charMap {
		cm[i2] = i
	}
	for i := 0; i < len(m); i++ {
		if c, ok := cm[rune(i)]; ok {
			m[i] = byte(c)
		}
	}
	return m
}
