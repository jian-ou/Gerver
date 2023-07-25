package giface

type IServer interface {
	Start()
	AddRouter(uint32, IRouter)
	GetRouter(uint32) IRouter
}
