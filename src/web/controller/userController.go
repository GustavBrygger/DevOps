package controller

import (
	"go-minitwit/src/application"
	"go-minitwit/src/persistence"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MapUserEndpoints(router *gin.Engine) {
	router.GET("/follow", followUser)
	router.GET("/unfollow", unfollowUser)
}

func followUser(context *gin.Context) {
	userID := abortIfNoUserID(context)
	username := context.Query("username")
	err := application.FollowUser(persistence.GetDbConnection(), userID, username)
	if err != nil {
		context.AbortWithError(http.StatusUnauthorized, err)
	}

	context.Redirect(http.StatusFound, "/user-timeline?username="+username)
}

func unfollowUser(context *gin.Context) {
	userID := abortIfNoUserID(context)
	username := context.Query("username")
	err := application.UnfollowUser(persistence.GetDbConnection(), userID, username)
	if err != nil {
		context.AbortWithError(http.StatusUnauthorized, err)
	}

	context.Redirect(http.StatusFound, "/user-timeline?username="+username)
}
