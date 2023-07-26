package gcoder_test

import (
	"Gerver/gcoder"
	"fmt"
	"testing"
)

func TestTLVcoder(t *testing.T) {
	c := gcoder.NewTLVCoder()
	d := c.Encode(1000, []byte("helloworld"))
	d = append(d, c.Encode(1000, []byte("helloworld"))...)
	fmt.Println(d)
	tag, length, value, _ := c.Decode(d)
	if len(d) > int(length) {
		d = d[length:]
	}
	fmt.Println(d)
	fmt.Println("tag : ", tag, "value : ", string(value))
}
