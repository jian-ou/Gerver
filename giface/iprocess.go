package giface

type IProcess interface {
	Run()
	AddRequest(IRequest)
	GetSize() int
}
