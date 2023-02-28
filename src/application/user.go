package application

import (
	"errors"
	"strconv"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string
	Email     string
	PW_hash   string
	Messages  []*Message
	Followers []*User `gorm:"many2many:user_followers"`
}

func GetUserByID(db *gorm.DB, userID uint) (User, error) {
	var user User
	result := db.Find(&user, userID)
	if result.Error != nil {
		return user, errors.New("User not found")
	}

	return user, nil
}

func GetUserByUsername(db *gorm.DB, username string) (User, error) {
	user := User{Username: username}
	result := db.Where("Username = ?", username).Find(&user)
	if result.Error != nil {
		return user, errors.New("User not found")
	}

	return user, nil
}

func FollowUser(db *gorm.DB, currUserID uint, usernameToFollow string) error {
	userToFollow, err := GetUserByUsername(db, usernameToFollow)
	currUser, _ := GetUserByID(db, currUserID)
	if err != nil {
		return err
	}

	userToFollow.Followers = append(userToFollow.Followers, &currUser)
	db.Save(&userToFollow)
	return nil
}

func UnfollowUser(db *gorm.DB, currUserID uint, usernameToUnFollow string) error {
	userToUnFollow, err := GetUserByUsername(db, usernameToUnFollow)
	if err != nil {
		return err
	}

	db.Unscoped().Exec("DELETE from user_followers WHERE user_id =" + strconv.Itoa(int(userToUnFollow.ID)) + " AND follower_id = " + strconv.Itoa(int(currUserID)))
	return nil
}

func IsUserFollowing(db *gorm.DB, currUserID uint, userID uint) bool {
	if currUserID == 0 {
		return false
	}

	user, _ := GetUserByID(db, userID)
	var followers []User
	db.Model(user).Where("follower_id = ?", currUserID).Association("Followers").Find(&followers)
	return len(followers) > 0
}
