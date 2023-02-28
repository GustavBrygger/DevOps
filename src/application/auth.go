package application

import (
	"errors"
	"go-minitwit/src/util"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Register struct {
	Username  string `json:"username" form:"username" binding:"required,min=1"`
	Email     string `json:"email" form:"email" binding:"required,min=1"`
	Password  string `json:"pwd" form:"pwd" binding:"required,min=1"`
	Password2 string `form:"password2"`
}

type Login struct {
	Username string `form:"username" binding:"required,min=1"`
	Password string `form:"password" binding:"required,min=1"`
}

func HandleRegister(context *gin.Context, db *gorm.DB) error {
	var register Register
	if err := context.ShouldBind(&register); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			return errors.New(ve[0].Field() + " is required")
		}
	}

	var existingUser = []User{}
	db.Where(&User{Username: register.Username}).Find(&existingUser)

	if len(existingUser) != 0 {
		return errors.New("The username is already taken")
	}

	new_user := User{
		Username: register.Username,
		Email:    register.Email,
		PW_hash:  util.HashPassword(register.Password),
	}

	db.Create(&new_user)

	return nil
}

func HandleLogin(context *gin.Context, db *gorm.DB, session sessions.Session) error {
	var login Login
	if err := context.ShouldBind(&login); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			return errors.New(ve[0].Field() + " is required")
		}
	}

	var user = []User{}
	db.Where(&User{Username: login.Username}).Find(&user)

	if len(user) == 0 {
		return errors.New("Invalid username")
	}

	if !util.PasswordMatch(login.Password, user[0].PW_hash) {
		return errors.New("Invalid password")
	}

	session.Set("userID", user[0].ID)
	session.Save()

	return nil
}
