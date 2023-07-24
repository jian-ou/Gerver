package giface

type IClient interface {
	Start()
	Send(b []byte) error
}
