package internal

import (
	"log"

	"github.com/nats-io/nats.go"
)

func NatsConnect(options ...nats.Option) (*nats.Conn, error) {
	log.Printf("Connecting to NATS server...\n")

	opts := []nats.Option{
		nats.ClientCert("/usr/local/filewave/certs/server.crt", "/usr/local/filewave/certs/server.key"),
	}
	opts = append(opts, options...)

	return nats.Connect(
		"massimo-mbp.fwx.one",
		// "demo.nats.io",
		opts...,
	)
}

func NatsClose(conn *nats.Conn) {
	log.Printf("Closing connection from %s", conn.ConnectedUrl())
	conn.Close()
}
