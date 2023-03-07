package persistence

import (
	"fmt"
	"go-minitwit/src/application"
	"os"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var localConnectionString = fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s", "localhost", "postgres", "postgres", 5432, "postgres")
var azureConnectionString = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
	"minitwit-db.database.windows.net", "minitwit", "dbpassworD1!", 1433, "minitwit-db")

func GetDbConnection() *gorm.DB {
	connString := localConnectionString
	isProduction := os.Getenv("IS_PRODUCTION")
	if isProduction == "TRUE" {
		connString = azureConnectionString
		dsn := "sqlserver://minitwit:dbpassworD1!@minitwit-db.database.windows.net:1433?database=minitwit-db"
		db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database")
		}

		return db
	}

	println(connString)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	return db
}

func ConfigurePersistence() {
	db := GetDbConnection()

	applyMigrations(db)
	seed(db)
}

func applyMigrations(db *gorm.DB) {
	db.AutoMigrate(&application.User{})
	db.AutoMigrate(&application.Message{})
}
