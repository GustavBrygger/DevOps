package controller

import (
	"go-minitwit/src/application"
	"go-minitwit/src/persistence"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MapMessageEndpoints(router *gin.Engine) {
	router.POST("/add_message", addMessage)
}

func addMessage(context *gin.Context) {
	userID := abortIfNoUserID(context)
	messageText := context.Request.FormValue("text")
	application.AddMessage(persistence.GetDbConnection(), userID, messageText)

	context.Redirect(http.StatusFound, "/")
}
