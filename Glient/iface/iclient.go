package IGlient

type IClient interface {
	Start()
	Send(b []byte) error
}
