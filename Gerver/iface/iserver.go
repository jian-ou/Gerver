package IGerver

type IServer interface {
	Start()
	AddRouter(router IRouter)
	GetRouter() IRouter
}
