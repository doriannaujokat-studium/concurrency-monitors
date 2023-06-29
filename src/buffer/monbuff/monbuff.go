package monbuff

import (
	"buffer/buff"
	. "sync"
)

type MBuffer interface {
	Ins(a byte)
	Get() byte
}

func New(n uint) MBuffer { return new_(n) }

type mbuffer struct {
	buff.Buffer
	notFull, notEmpty *Cond
	Mutex
}

func new_(n uint) MBuffer {
	x := new(mbuffer)
	x.Buffer = buff.New(n)
	x.notFull = NewCond(&x.Mutex)
	x.notEmpty = NewCond(&x.Mutex)
	return x
}

func (x *mbuffer) Ins(b byte) {
	x.Mutex.Lock()
	defer x.Mutex.Unlock()
	for x.Buffer.Full() {
		x.notFull.Wait()
	}
	x.Buffer.Ins(b)
	Debugf("Wrote %02x", b)
	x.notEmpty.Signal()
}
func (x *mbuffer) Get() byte {
	x.Mutex.Lock()
	defer x.Mutex.Unlock()
	for x.Buffer.Num() == 0 {
		x.notEmpty.Wait()
	}
	var b = x.Buffer.Get()
	Debugf("Read %02x", b)
	x.notFull.Signal()
	return b
}
