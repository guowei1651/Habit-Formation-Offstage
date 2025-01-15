package device

import (
	"log"
	"time"
)

func OpenServer() {
	log.Printf("device open server")
	for true {
		time.Sleep(100 * time.Millisecond)
	}
}