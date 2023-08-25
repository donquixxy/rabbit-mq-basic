package broker

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func InitBroker() *amqp.Channel {
	con, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672")

	if err != nil {
		log.Fatalf("failed to dial amqp %v", err)
	}

	conn, err := con.Channel()

	if err != nil {
		log.Fatalf("failed to connect to channel %v", err)
	}

	log.Println("Success connected to amqp")
	return conn
}

func InitBrokerSend() *amqp.Channel {
	con, err := amqp.Dial("amqp://guest:guest@localhost:37888")

	if err != nil {
		log.Fatalf("failed to dial amqp %v", err)
	}

	conn, err := con.Channel()

	if err != nil {
		log.Fatalf("failed to connect to channel %v", err)
	}

	log.Println("Success connected to amqp")
	return conn
}
