package main

import "fmt"

type fifo struct {
}

func (f *fifo) evict() {
	fmt.Println("use fifo strategy")
}
