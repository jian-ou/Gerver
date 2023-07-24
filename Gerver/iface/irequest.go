package IGerver

type IRequest interface {
	GetData() []byte
	GetConnection() IConnection
}
