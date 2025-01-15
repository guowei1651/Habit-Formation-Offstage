package main

import (
	"hf/config"
	"hf/database"
	"hf/web"
	"hf/device"
)

func main() {
	config.ParseConfig()

	database.openConnectPool()

	web.openServer()
	device.openServer()
}
