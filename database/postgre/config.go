package postgre

import (
	"fmt"
	"log"
	"micro-company/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreConfig struct {
	Driver       string
	Username     string
	Password     string
	Port         uint
	Address      string
	DatabaseName string
	Schemas      string
	SSLMode      string
}

func InitPostgreDb(config PostgreConfig) (*gorm.DB, error) {
	// dsn := "host=host.docker.internal user=agusari password=12345678 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	// dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
	// 	"agusari",
	// 	"12345678",
	// 	"host.docker.internal",
	// 	"5432",
	// 	"postgres")

	// dsn := "host=postgre-micro user=agusari password=12345678 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		"postgre", "agusari", "12345678", "postgres", 5432,
	)
	database, er := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if er != nil {
		log.Println("error opening database postgre", er.Error())
		return nil, er
	}

	db, er := database.DB()

	if er != nil {
		return nil, er
	}

	er = db.Ping()

	if er != nil {
		return nil, er
	}

	log.Println("Connected to database Postgre!")

	return database, nil
}

func ClosePostgreConnection(db *gorm.DB) {
	database, errDb := db.DB()

	if errDb != nil {
		panic("error connecting to database in order to close connection")
	}

	errClose := database.Close()

	if errClose != nil {
		panic(errClose.Error())
	}

	fmt.Println("Success to close connection Postgres database")
}

func InitPostgreDbLocal(config PostgreConfig) (*gorm.DB, error) {
	// dsn := "host=host.docker.internal user=agusari password=12345678 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	// dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
	// 	"agusari",
	// 	"12345678",
	// 	"host.docker.internal",
	// 	"5432",
	// 	"postgres")

	// dsn := "host=postgre-micro user=agusari password=12345678 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		"localhost", "agusari", "12345678", "postgres", 37892,
	)
	database, er := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if er != nil {
		log.Println("error opening database postgre", er.Error())
		return nil, er
	}

	db, er := database.DB()

	if er != nil {
		return nil, er
	}

	er = db.Ping()

	if er != nil {
		return nil, er
	}

	log.Println("Connected to database Postgre!")

	return database, nil
}

func AutoMigrate(db *gorm.DB) {

	db.AutoMigrate(&domain.Company{}, &domain.User{})

}
