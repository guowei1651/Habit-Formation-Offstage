package main

import (
	"log"

	"hf/config"
	"hf/database"
	"hf/web"
	"hf/device"
)

func main() {
	config.ParseConfig()

	database.OpenConnectPool()

	ch := make(chan string)
	go web.OpenServer(ch)
	go device.OpenServer(ch)
	err := <- ch
	log.Fatal("错误退出。%s", err)
}
