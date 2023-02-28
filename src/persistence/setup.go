package persistence

import (
	"fmt"
	"go-minitwit/src/application"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DbServer   string `env:"DBSERVER,required"`
	DbPort     int    `env:"DBPORT,required"`
	DbName     string `env:"DBNAME,required"`
	DbUser     string `env:"DBUSER,required"`
	DbPassword string `env:"DBPASSWORD,required"`
}

var connString = ""

func GetDbConnection() *gorm.DB {
	println(connString)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	return db
}

func ConfigurePersistence(config Config) {
	connString = fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s",
		config.DbServer, config.DbUser, config.DbPassword, config.DbPort, config.DbName)
	println(connString)
	db := GetDbConnection()

	applyMigrations(db)
	seed(db)
}

func applyMigrations(db *gorm.DB) {
	db.AutoMigrate(&application.User{})
	db.AutoMigrate(&application.Message{})
}
