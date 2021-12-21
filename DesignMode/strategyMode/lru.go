package main

import "fmt"

type lru struct {
}

func (l *lru) evict() {
	fmt.Println("use lru strategy")
}
