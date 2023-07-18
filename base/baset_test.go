package base

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestEncodeToBase(t *testing.T) {
	data := []byte("test1234")

	var output []byte
	//data := []byte{240}
	d := uint32(63)
	fmt.Printf("%032b\n%032b\n", d, d<<26)
	x := base64.StdEncoding.EncodeToString(data)
	t.Log(x, len(x))
	tmp := make([]byte, 10)
	base64.StdEncoding.Encode(data, tmp)
	t.Log(x, len(x))
	t.Logf("%b\n", tmp)

	encoder := New(base64Charset)
	fmt.Println(encoder.encodeLength(len(data)))
	dst := make([]byte, encoder.encodeLength(len(data)))
	encoder.Encode(data, dst)
	result := string(dst)
	t.Logf("%s\t%d\n", result, len(result))

	encoder.Decode(dst, output)
	result = string(dst)
	t.Log(result)

}
