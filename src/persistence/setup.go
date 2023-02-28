package persistence

import (
	"go-minitwit/src/application"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDbConnection(config *application.Config) *gorm.DB {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		config.DbServer, config.DbUser, config.DbPassword, config.DbPort, config.DbName)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	return db
}

func ConfigurePersistence(config *application.Config) {
	db := GetDbConnection(config)

	applyMigrations(db)
	seed(db)
}

func applyMigrations(db *gorm.DB) {
	db.AutoMigrate(&application.User{})
	db.AutoMigrate(&application.Message{})
}
