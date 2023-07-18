package base

import (
	"fmt"
	"math"
	"reflect"
)

type BitMan struct {
	buffer []byte
	pos    int
	subPos int
}

func (b *BitMan) GetBits(num int, data uint64) error {
	if num > GetTypeLen(data) {
		return fmt.Errorf("bit length bigger than bumber of bits requested")
	}
	//b.subPos + num
	for i := 0; i < int(math.Ceil(float64(num/8)))-1; i++ {
		//data += b.buffer[b.pos] << i * 8
		b.pos += num
	}

	//remainder := num % 8
	//data += b.buffer[b.pos+1] << int(math.Ceil(float64(num/8))-1) * 8
	b.pos++

	return nil
}
func GetTypeLen(v any) int {
	return reflect.TypeOf(v).Bits()
}
func GetLen(v any) int {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice, reflect.Array:
		return reflect.TypeOf(v).Elem().Bits() * reflect.ValueOf(v).Len()
	default:
		return reflect.TypeOf(v).Bits()
	}
}
