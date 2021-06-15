package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/mcosta74/play-with-nats/internal"
	"github.com/nats-io/nats.go"
)

func main() {
	fmt.Println("Consumer")
	defer fmt.Println("Consumer Done")

	nc, err := internal.NatsConnect()
	if err != nil {
		log.Fatalf("Error connecting to nats: %v", err)
	}
	defer internal.NatsClose(nc)

	wg := sync.WaitGroup{}
	wg.Add(1)

	if _, err := nc.Subscribe("chat.*", func(msg *nats.Msg) {
		if msg.Subject == "chat.close" {
			fmt.Println("Chat Closed")
			wg.Done()
		} else {
			fmt.Printf("Msg: %s, %s\n", msg.Subject, msg.Data)
		}
	}); err != nil {
		log.Fatal(err)
	}
	wg.Wait()
}

func handleChatMsg(msg *nats.Msg) {
	if msg.Subject == "chat.close" {

	}
	fmt.Printf("Msg: %s, %s\n", msg.Subject, msg.Data)
}
