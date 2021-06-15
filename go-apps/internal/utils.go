package internal

import (
	"log"

	"github.com/nats-io/nats.go"
)

func NatsConnect() (*nats.Conn, error) {
	log.Printf("Connecting to %v\n", nats.DefaultURL)
	return nats.Connect(
		"massimo-mbp.fwx.one",
		nats.ClientCert("/usr/local/filewave/certs/server.crt", "/usr/local/filewave/certs/server.key"),
	)
}

func NatsClose(conn *nats.Conn) {
	log.Printf("Closing connection from %s", conn.ConnectedUrl())
	conn.Close()
}
