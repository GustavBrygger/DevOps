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

func getConnectionString(dbPassword string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", "minitwit_db", "postgres", dbPassword, "postgres", 5432)
}

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
	print(isProduction)
	if isProduction == "TRUE" {
		dbPassword := os.Getenv("DB_PASSWORD")
		isAzure := os.Getenv("IS_AZURE")
		if isAzure == "FALSE" {
			conn, err := gorm.Open(postgres.Open(getConnectionString(dbPassword)), &gorm.Config{})
			if err != nil {
				log.Fatal("Failed to connect to prod database")
			}

			return conn
		}

		azureConn, err := gorm.Open(sqlserver.Open(getAzureConnString(dbPassword)), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to azure database")
		}

		return azureConn
	}

	localConn, err := gorm.Open(postgres.Open(getConnectionString("postgres")), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to local database")
	}

	return localConn
}

func ConfigurePersistence() {
	db = initDbConnection()

	applyMigrations(db)
	seed(db)
}

func applyMigrations(db *gorm.DB) {
	err1 := db.AutoMigrate(&application.User{})
	err2 := db.AutoMigrate(&application.Message{})
	if err1 != nil || err2 != nil {
		log.Fatal("Migration failed")
	}
}
