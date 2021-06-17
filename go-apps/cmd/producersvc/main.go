package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mcosta74/play-with-nats/internal"
	"github.com/nats-io/nats.go"
)

func main() {
	fmt.Println("Producer")

	nc, err := internal.NatsConnect(nats.Name("Go Producer"))
	if err != nil {
		log.Fatalf("Error connecting to nats: %v", err)
	}
	defer internal.NatsClose(nc)

	ticker := time.NewTicker(time.Millisecond * 500)
	errs := make(chan error)

	go func(amount int) {
		count := 0

		for t := range ticker.C {
			count += 1
			if err := nc.Publish("chat.msg", []byte(fmt.Sprintf("%s Hello", t.Format(time.StampMilli)))); err != nil {
				errs <- fmt.Errorf("error sending: %v", err)
			}

			if count >= amount {
				close(errs)
			}
		}
	}(30)

	// Wait for done
	err = <-errs
	fmt.Println(err)

	nc.Publish("chat.close", []byte(""))

	fmt.Println("Bye")
}
