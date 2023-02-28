package controller

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func ConfigureSession(router *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24})
	router.Use(sessions.Sessions("mysession", store))
}

func getCurrentUserId(context *gin.Context) uint {
	session := sessions.Default(context)
	userID := session.Get("userID")

	if userID != nil {
		return session.Get("userID").(uint)
	}

	return 0
}

func abortIfNoUserID(context *gin.Context) uint {
	userID := getCurrentUserId(context)
	if userID == 0 {
		context.AbortWithError(http.StatusUnauthorized, errors.New("Invalid session"))
		return 0
	}

	return userID
}

func clearSession(context *gin.Context) {
	session := sessions.Default(context)
	session.Clear()
	session.Save()
}
