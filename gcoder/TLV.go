package gcoder

import (
	"Gerver/giface"
	"bytes"
	"encoding/binary"
)

const TLV_HEADER_SIZE = 8 //表示TLV空包长度

type TLVcoder struct {
}

func NewTLVDecoder() giface.ICoder {
	return &TLVcoder{}
}

func (c *TLVcoder) Encode(tag uint32, value []byte) []byte {
	b := make([]byte, 0)
	b = binary.BigEndian.AppendUint32(b, tag)
	b = binary.BigEndian.AppendUint32(b, uint32(len(value)))
	b = append(b, value...)
	return b
}

func (c *TLVcoder) Decode(data []byte) (uint32, uint32, []byte) {
	tag := binary.BigEndian.Uint32(data[0:4])
	length := binary.BigEndian.Uint32(data[4:8])
	value := make([]byte, length)

	binary.Read(bytes.NewBuffer(data[8:8+length]), binary.BigEndian, value)

	return tag, 8 + length, value
}
