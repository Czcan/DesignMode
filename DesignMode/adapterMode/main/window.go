package main

import "fmt"

//未知服务

type windows struct {
	version string
}

func (w *windows) InsertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine")
}
