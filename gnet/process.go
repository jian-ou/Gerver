package gnet

import (
	"Gerver/giface"
	"fmt"
)

type Process struct {
	dispatch    giface.IDispatch
	ID          int
	RequestChan chan giface.IRequest
	size        int
	maxSize     int
}

func NewProcess(dispatch giface.IDispatch, maxSize int, ID int) giface.IProcess {
	p := &Process{
		dispatch:    dispatch,
		ID:          ID,
		maxSize:     maxSize,
		size:        0,
		RequestChan: make(chan giface.IRequest),
	}
	fmt.Printf("Process [%d] start\n", ID)
	go p.Run()
	return p
}

func (p *Process) Run() {
	for {
		select {
		case r := <-p.RequestChan:
			r.Run()
			p.size--
		}
	}
}

func (p *Process) AddRequest(r giface.IRequest) {
	p.RequestChan <- r
	p.size++
}

func (p *Process) GetSize() int {
	return p.size
}
