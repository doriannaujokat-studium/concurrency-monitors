package main

import (
	"buffer/monbuff"
	"crypto/rand"
	"fmt"
	"sync"
	"time"
)

var buffer = monbuff.New(5)
var waitgroup sync.WaitGroup
var starttime = time.Now()

func timeDiff() int64 {
	return time.Since(starttime).Nanoseconds()
}

func reader() {
	defer waitgroup.Done()
	var b = buffer.Get()
	var timestamp = timeDiff()
	fmt.Printf("[%10d] A reader read %02x\n", timestamp, b)
}

func writer() {
	defer waitgroup.Done()
	var b []byte = make([]byte, 1)
	rand.Read(b)
	buffer.Ins(b[0])
	var timestamp = timeDiff()
	fmt.Printf("[%10d] A writer wrote %02x\n", timestamp, b)
}

func main() {
	for i := 0; i < 20; i++ {
		waitgroup.Add(1)
		go writer()
	}
	for i := 0; i < 20; i++ {
		waitgroup.Add(1)
		go reader()
	}
	waitgroup.Wait()
}
