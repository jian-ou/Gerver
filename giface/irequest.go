package giface

type IRequest interface {
	GetData() []byte
	GetConnection() IConnection
	Run()
}
