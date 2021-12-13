package main

//适配器模式

func main() {
	client := &client{}
	mac := &mac{}

	client.insertLightingConnectorIntoComputer(mac)

	windowsMachine := &windows{version: "AMD"}
	windowsAdapter := &windowsAdapter{
		windowsMachine: windowsMachine,
	}

	client.insertLightingConnectorIntoComputer(windowsAdapter)
}
