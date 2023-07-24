package gnet

import "Gerver/giface"

type BaseRouter struct{}

func (br *BaseRouter) PreHandle(req giface.IRequest)  {}
func (br *BaseRouter) Handle(req giface.IRequest)     {}
func (br *BaseRouter) PostHandle(req giface.IRequest) {}
