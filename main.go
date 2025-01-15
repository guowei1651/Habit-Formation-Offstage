package main

import (
	"hf/config"
	"hf/database"
	"hf/web"
	"hf/device"
)

func main() {
	config.ParseConfig()

	database.OpenConnectPool()

	web.OpenServer()
	device.OpenServer()
}
