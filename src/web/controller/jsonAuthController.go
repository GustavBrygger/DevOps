package controller

import (
	"go-minitwit/src/application"
	"go-minitwit/src/persistence"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MapJSONAuthEndpoints(router *gin.Engine) {
	router.POST("/register", registerUser)
}

func registerUser(context *gin.Context){
	if err := application.HandleRegister(context, persistence.GetDbConnection()); err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
	}
	updateLatest(context.Request)
	context.Status(http.StatusNoContent)
}