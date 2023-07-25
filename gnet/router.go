package gnet

import (
	"Gerver/giface"
	"fmt"
)

type BaseRouter struct {
}

func (br *BaseRouter) PreHandle(req giface.IRequest)  {}
func (br *BaseRouter) Handle(req giface.IRequest)     {}
func (br *BaseRouter) PostHandle(req giface.IRequest) {}

type NoneRouter struct {
	BaseRouter
}

func (br *NoneRouter) Handle(req giface.IRequest) {
	fmt.Println("None Router")
}
