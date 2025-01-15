package main

import (
	config "hf/config"
	db "hf/database"
	web "hf/web"
	device "hf/device"
)

func main() {
	config.ParseConfig()

	db.init()

	web.openServer()
	device.openServer()
}
