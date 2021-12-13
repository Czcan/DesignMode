package main

import "fmt"

//客户端代码

type client struct{}

func (c *client) insertLightingConnectorIntoComputer(com computer) {
	fmt.Println("Client insert lighting connector into computer")
	com.insertIntoLightingPort()
}
