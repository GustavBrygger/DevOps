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

var localConnectionString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", "minitwit_db", "postgres", "postgres", "postgres", 5432)

func getAzureConnString(dbPassword string) string {
	return fmt.Sprintf("sqlserver://%s:%s@minitwit-db.database.windows.net:1433?database=minitwit-db", "minitwit", dbPassword)
}

var db *gorm.DB = nil

func GetDbConnection() *gorm.DB {
	if db != nil {
		return db
	}

	return initDbConnection()
}

func initDbConnection() *gorm.DB {
	isProduction := os.Getenv("IS_PRODUCTION")
	if isProduction == "TRUE" {
		dbPassword := os.Getenv("DB_PASSWORD")
		azureConn, err := gorm.Open(sqlserver.Open(getAzureConnString(dbPassword)), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database")
		}

		return azureConn
	}

	localConn, err := gorm.Open(postgres.Open(localConnectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	return localConn
}

func ConfigurePersistence() {
	db = initDbConnection()

	applyMigrations(db)
	seed(db)
}

func applyMigrations(db *gorm.DB) {
	db.AutoMigrate(&application.User{})
	db.AutoMigrate(&application.Message{})
}
