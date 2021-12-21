package main

import "fmt"

type lfu struct {
}

func (l *lfu) evict() {
	fmt.Println("use lfu strategy")
}
