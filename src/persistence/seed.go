package persistence

import (
	"fmt"
	"go-minitwit/src/application"
	"go-minitwit/src/util"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func seed(db *gorm.DB, redisDb *redis.Client) {
	var users []application.User
	result := db.Find(&users)
	if result.RowsAffected == 0 {
		addUsersAndMessages(db)
	}

	key_count, _ := redisDb.DBSize().Result()
	if key_count == 0 {
		err := redisDb.Set("latest", -1, 0).Err()
		if err != nil {
			fmt.Print(err)
		}
	}
}

func addUsersAndMessages(db *gorm.DB) {
	db.Exec("INSERT INTO users (id, username, email, pw_hash) VALUES (0, 'Roger Histand', 'Roger+Histand@hotmail.com', '$2a$14$o0RLqWDxXvMAvcTrKqYVbuF4JcTT5i8tS2b8nzSN4u.UP5odFsFdG');")

	user1 := application.User{
		Username: "AndenTester",
		Email:    "andentester@gmail.com",
		PW_hash:  util.HashPassword("AndenTest"),
		Messages: []*application.Message{
			{Text: "In Japan"},
		},
	}
	user2 := application.User{
		Username: "Cool",
		Email:    "cool@gmail.com",
		PW_hash:  util.HashPassword("Secret"),
		Messages: []*application.Message{
			{Text: "Hello World"},
		},
		Followers: []*application.User{&user1},
	}

	db.Create([]*application.User{&user1, &user2})
	db.Model(&user1).Association("Followers").Append(&user2)
	db.Model(&user2).Association("Followers").Append(&user1)
	db.Save([]*application.User{&user1, &user2})
}
