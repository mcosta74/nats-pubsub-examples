package main

import (
	"fmt"
	"log"

	"github.com/mcosta74/play-with-nats/internal"
)

func main() {
	fmt.Println("Producer")

	nc, err := internal.NatsConnect()
	if err != nil {
		log.Fatalf("Error connecting to nats: %v", err)
	}
	defer internal.NatsClose(nc)

	for i := 0; i < 100; i++ {
		if err := nc.Publish("chat.msg", []byte(fmt.Sprintf("Hello #%d", i))); err != nil {
			log.Fatal(err)
		}
	}
	nc.Publish("chat.close", []byte(""))
}
