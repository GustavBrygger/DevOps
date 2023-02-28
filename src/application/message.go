package application

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	UserID  uint
	Text    string
	flagged bool
}

type MessageDTO struct {
	CreatedAt string `json:"pub_date"`
	Username  string `json:"user"`
	Text      string `json:"content"`
	AvatarURL string `json:"-"`
}

func GetAllMessages(db *gorm.DB) []MessageDTO {
	var messages []Message
	db.Find(&messages)

	return toMessageDTO(db, messages)
}

func GetFirstNMessages(db *gorm.DB, n int) []MessageDTO {
	var messages []Message
	db.Limit(n).Where(&Message{flagged: false}).Find(&messages)

	return toMessageDTO(db, messages)
}

func GetMessagesByUserID(db *gorm.DB, userID uint) []MessageDTO {
	var messages []Message
	db.Where(&Message{UserID: userID}).Find(&messages)

	return toMessageDTO(db, messages)
}

func GetNMessagesByUsername(db *gorm.DB, username string, n int) []MessageDTO {
	user, _ := GetUserByUsername(db, username)
	var messages []Message
	db.Limit(n).Where(&Message{UserID: user.ID, flagged: false}).Find(&messages)

	return toMessageDTO(db, messages)
}

func AddMessage(db *gorm.DB, userID uint, text string) {
	message := Message{UserID: userID, Text: text}
	db.Create(&message)
	db.Save(&message)
}

func toMessageDTO(db *gorm.DB, messages []Message) []MessageDTO {
	var messageDTOs []MessageDTO
	for _, message := range messages {
		user, _ := GetUserByID(db, message.UserID)
		var avatarURL = getAvatarURL(user.Email)
		messageDTO := MessageDTO{
			CreatedAt: message.CreatedAt.Format("2006-01-02"),
			Username:  user.Username,
			Text:      message.Text,
			AvatarURL: avatarURL,
		}
		messageDTOs = append(messageDTOs, messageDTO)
	}

	return messageDTOs
}

func getAvatarURL(email string) string {
	emailMD5 := fmt.Sprintf("%s", md5.Sum([]byte(email)))
	emailHex := hex.EncodeToString([]byte(emailMD5))
	url := fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon&s=48", emailHex)

	return url
}
