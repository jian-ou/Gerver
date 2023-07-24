package Gerver

import IGerver "Gerver/Gerver/iface"

type BaseRouter struct{}

func (br *BaseRouter) PreHandle(req IGerver.IRequest)  {}
func (br *BaseRouter) Handle(req IGerver.IRequest)     {}
func (br *BaseRouter) PostHandle(req IGerver.IRequest) {}
