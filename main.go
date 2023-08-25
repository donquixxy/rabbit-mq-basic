package main

import (
	"encoding/json"
	"log"
	"micro-company/config"
	"micro-company/database/broker"
	"micro-company/database/mysql"
	"micro-company/database/postgre"
	"micro-company/internal/domain"
	"micro-company/internal/repository"
	"os"
	"os/signal"
	"syscall"
)

// func main() {

// 	config := config.InitConfig()

// 	db, _ := postgre.InitPostgreDbLocal(*config.Database)

// 	mq := broker.InitBrokerSend()

// 	q, err := mq.QueueDeclare("bss", false, false, false, false, nil)

// 	if err != nil {
// 		log.Fatalf("failed to declare queue %v", err)
// 	}

// 	postgre.AutoMigrate(db)
// 	// item := &domain.Company{
// 	// 	ID:        "91314134",
// 	// 	Name:      "name132",
// 	// 	Phone:     "9983",
// 	// 	CreatedAt: time.Now(),
// 	// 	UpdatedAt: time.Now(),
// 	// }

// 	user := &domain.User{
// 		ID:        "1231312",
// 		Name:      "user",
// 		Age:       "231",
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}

// 	marshall, err := json.Marshal(user)

// 	if err != nil {
// 		log.Fatalf("failed to marshall %v", err)
// 	}

// 	var u *domain.User
// 	json.Unmarshal(marshall, &u)

// 	log.Println(u)

// 	headers := make(amqp091.Table)
// 	headers["message_type"] = "user"

// 	err = mq.PublishWithContext(context.Background(), "", q.Name, false, false, amqp091.Publishing{
// 		ContentType: "application/json",
// 		Body:        marshall,
// 		Timestamp:   time.Now(),
// 		Headers:     headers,
// 	})

// 	if err != nil {
// 		log.Fatalf("failed send request %v", err)
// 	}

// 	log.Println("Successfully sent request")
// }

// =================================================================

func main() {

	config := config.InitConfig()

	userRepository := repository.NewUserRepository()
	companyRepository := repository.NewCommpanyRepository()

	mysql.InitMySqlServer(config.Mysql)

	db, _ := postgre.InitPostgreDb(*config.Database)

	postgre.AutoMigrate(db)

	conn := broker.InitBroker()

	q, err := conn.QueueDeclare("bss", false, false, false, false, nil)

	if err != nil {
		log.Fatalf("failed to declare queue %v", err)
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

					// Insert to db
					v, err := userRepository.Create(db, user)

					if err != nil {
						log.Println("Failed create user: ", err)
					}

					log.Println("Success creating user :", v)

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

					// Insert to db
					v, err := companyRepository.Create(db, company)

					if err != nil {
						log.Println("Failed create company: ", err)
					}

					log.Println("Success creating company :", v)
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
