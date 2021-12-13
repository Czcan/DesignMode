package main

import "fmt"

type windowsAdapter struct {
	windowsMachine *windows //对象结构适配器
}

func (w *windowsAdapter) insertIntoLightingPort() { //相当于实现了computer接口
	fmt.Println("Adapter converts Lighting signal to USB")
	w.windowsMachine.InsertIntoUSBPort() //接口里面执行windows的方法
}
