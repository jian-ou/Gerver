package giface

type ICoder interface {
	Encode(uint32, []byte) []byte
	Decode([]byte) (uint32, uint32, []byte)
}
