package giface

type IServer interface {
	Start()
	AddRouter(uint32, IRouter)
	GetRouter(uint32) IRouter
	AddPreHandle(func(IConnection))
	AddPostHandle(func(IConnection))
	GetPreHandle() func(IConnection)
	GetPostHandle() func(IConnection)
	GetDispatch() IDispatch
}
