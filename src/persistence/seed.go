package persistence

import (
	"go-minitwit/src/application"
	"go-minitwit/src/util"

	"gorm.io/gorm"
)

func seed(db *gorm.DB) {
	var users []application.User
	result := db.Find(&users)
	if result.RowsAffected == 0 {
		addUsersAndMessages(db)
	}
}

func addUsersAndMessages(db *gorm.DB) {
	user1 := application.User{
		Username: "Tester",
		Email:    "tester@gmail.com",
		PW_hash:  util.HashPassword("Test"),
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
