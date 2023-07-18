package base

import (
	"github.com/tunabay/go-bitarray"
	"log"
	"testing"
)

func TestGetTypeLen(t *testing.T) {
	a := 20.0
	ab := []int32{20, 10, 33}
	ac := []byte{20, 10, 33}
	log.Println(GetTypeLen(a))
	log.Println(GetLen(a))
	log.Println(GetLen(ab))
	bb := bitarray.NewBufferFromByteSlice(ac)
	log.Println(bb.BitArrayAt(0, 10).ToUint64())

}
