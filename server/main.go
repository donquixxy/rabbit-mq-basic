package main

import (
	"encoding/json"
	"log"
	"micro-company/config"
	"micro-company/database/broker"
	"micro-company/database/postgre"
	"micro-company/internal/domain"
	"micro-company/internal/repository"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// =================================================================

func main() {

	config := config.InitConfig()

	userRepository := repository.NewUserRepository()
	companyRepository := repository.NewCommpanyRepository()

	db, _ := postgre.InitPostgreDbLocal(*config.Database)

	postgre.AutoMigrate(db)

	conn := broker.InitBrokerSend()

	q, err := conn.QueueDeclare("bss", false, false, false, false, nil)

	if err != nil {
		log.Fatalf("failed to declare queue %v", err)
	}

	err = conn.ExchangeDeclare(
		"logs", "fanout", true, false, false, false, nil,
	)

	if err != nil {
		log.Fatalf("failed to get exchange declaration %v", err)
	}

	msgg, err := conn.Consume(q.Name, "", false, false, false, false, nil)

	if err != nil {
		log.Fatalf("failed to consume %v", err)
	}

	go func() {
		for msg := range msgg {
			log.Println("received message : ", string(msg.Body))

			messageType, ok := msg.Headers["message_type"].(string)
			if !ok {
				// No message_type field, handle accordingly
				log.Println("not ok")
				continue
			}

			switch messageType {
			case "user":
				{

					log.Println("User message type received !")
					// Unmarshall user
					var user *domain.User

					err := json.Unmarshal(msg.Body, &user)

					if err != nil {
						log.Fatalf("failed to unmarshal user message :%v", err)
					}

					time.Sleep(time.Second * 2)

					// Insert to db
					v, err := userRepository.Create(db, user)

					if err != nil {
						log.Println("Failed create user: ", err)
					}

					log.Println("Success creating user :", v)

					msg.Ack(false)

				}
			case "company":
				{
					log.Println("Company message type received !")
					// Unmarshall user
					var company *domain.Company

					err := json.Unmarshal(msg.Body, &company)

					if err != nil {
						log.Fatalf("failed to unmarshal user company :%v", err)
					}

					time.Sleep(time.Second * 2)

					// Insert to db
					v, err := companyRepository.Create(db, company)

					if err != nil {
						log.Println("Failed create company: ", err)
					}

					log.Println("Success creating company :", v)

					msg.Ack(false)
				}
			default:
				{
					log.Println("invalid message type received :", messageType)
				}
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)

	<-quit
}
