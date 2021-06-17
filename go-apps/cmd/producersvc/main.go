package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mcosta74/play-with-nats/internal"
)

func main() {
	fmt.Println("Producer")

	nc, err := internal.NatsConnect()
	if err != nil {
		log.Fatalf("Error connecting to nats: %v", err)
	}
	defer internal.NatsClose(nc)

	for i := 0; i < 20; i++ {
		if err := nc.Publish("chat.msg", []byte(fmt.Sprintf("%s Hello", time.Now().Format(time.RFC3339)))); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second * 1)
	}
	nc.Publish("chat.close", []byte(""))
}
