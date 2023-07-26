package gnet

import "Gerver/giface"

type Dispatch struct {
	server    giface.IServer
	maxSize   int
	processes []giface.IProcess
}

func NewDispatch(server giface.IServer, maxSize int) giface.IDispatch {
	d := &Dispatch{
		server:    server,
		maxSize:   maxSize,
		processes: make([]giface.IProcess, maxSize),
	}
	for i := 0; i < d.maxSize; i++ {
		d.processes[i] = NewProcess(d, 100, i)
	}
	return d
}

func (d *Dispatch) AddRequest(r giface.IRequest) {
	index := 0
	minstacksize := 2000
	for i := 0; i < d.maxSize; i++ {
		temp := d.processes[i].GetSize()
		if minstacksize > temp {
			minstacksize = temp
			index = i
		}
	}
	d.processes[index].AddRequest(r)
}
