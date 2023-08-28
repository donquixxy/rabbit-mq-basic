package main

import (
	"context"
	"encoding/json"
	"log"
	"micro-company/database/broker"
	"micro-company/internal/domain"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func main() {

	mq := broker.InitBrokerSend()

	q, err := mq.QueueDeclare("bss", false, false, false, false, nil)

	if err != nil {
		log.Fatalf("failed to declare queue %v", err)
	}

	company := &domain.Company{
		ID:        "8891231333",
		Name:      "PT Sakatan",
		Phone:     "08991212",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	table := make(amqp091.Table)

	table["message_type"] = "company"

	v, err := json.Marshal(company)

	if err != nil {
		log.Fatalf("failed to marshal company %v", err)
	}

	err = mq.PublishWithContext(context.Background(), "", q.Name, false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        v,
		Headers:     table,
	})

	if err != nil {
		log.Fatalf("failed to publish %v", err)
	}

}
