package mysql

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MysqlConfig struct {
	Driver       string
	Username     string
	Password     string
	Port         uint
	Address      string
	DatabaseName string
	Schemas      string
	SSLMode      string
}

func InitMySqlServer(config *MysqlConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"12345678",
		"mysql",
		"10891",
		"micro-server",
	)

	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if dbErr != nil {
		panic("Error during opening Database !" + dbErr.Error())
	}

	sqlDB, sqlDbErr := db.DB()

	if sqlDbErr != nil {
		panic("Error during Connecting Database ! " + sqlDbErr.Error())
	}

	pingErr := sqlDB.Ping()

	if pingErr != nil {
		panic("Error during connecting database !" + pingErr.Error())
	}

	log.Println("Succes connected to Database !")

	return db
}
