package Glient

type IClient interface {
	Start()
	Send(b []byte) error
}
