package main

import (
	"fmt"
	"unsafe"
)

type T struct {
	x int
	y *[1 << 23]byte
}

func bar() {
	t := T{y: new([1 << 23]byte)}

	p := uintptr(unsafe.Pointer(&t.y[0]))
	*(*byte)(unsafe.Pointer(p)) = 1
	fmt.Println(t.x)
}

func main() {
	for {
		bar()
	}
	//bar()
}
