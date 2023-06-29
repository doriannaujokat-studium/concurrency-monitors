package main

import (
	. "buffer/buff"
	"fmt"
)

func main() {
	var b = New(5)
	var val byte = 0x0
	var rval byte = 0x0
	for !b.Full() {
		b.Ins(val)
		fmt.Printf("Inserted %02x into buffer\n", val)
		val++
	}
	for !b.Empty() {
		rval = b.Get()
		fmt.Printf("Read %02x from buffer\n", rval)
	}
	for !b.Full() {
		b.Ins(val)
		fmt.Printf("Inserted %02x into buffer\n", val)
		val++
	}
	for !b.Empty() {
		rval = b.Get()
		fmt.Printf("Read %02x from buffer\n", rval)
	}
}
