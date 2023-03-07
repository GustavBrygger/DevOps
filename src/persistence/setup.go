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

func getAzureConnString(dbPassword string) string {
	return fmt.Sprintf("sqlserver://%s:%s@minitwit-db.database.windows.net:1433?database=minitwit-db", "minitwit", dbPassword)
}

func GetDbConnection() *gorm.DB {
	isProduction := os.Getenv("IS_PRODUCTION")
	if isProduction == "TRUE" {
		dbPassword := os.Getenv("DB_PASSWORD")
		db, err := gorm.Open(sqlserver.Open(getAzureConnString(dbPassword)), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database")
		}

		return db
	}

	db, err := gorm.Open(postgres.Open(localConnectionString), &gorm.Config{})
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
