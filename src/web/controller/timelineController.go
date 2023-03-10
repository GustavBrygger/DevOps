package controller

import (
	"go-minitwit/src/application"
	"go-minitwit/src/persistence"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MapTimelineEndpoints(router *gin.Engine) {
	router.GET("/", renderTimeline)
	router.GET("/public", renderPublicTimeline)
	router.GET("/user-timeline", renderUserTimeline)
}

func renderTimeline(context *gin.Context) {
	db := persistence.GetDbConnection()
	currUser := getCurrentUser(context, db)
	if currUser == nil {
		renderPublicTimeline(context)
		return
	}

	messages := application.GetMessagesByUserID(db, currUser.ID)
	context.HTML(http.StatusOK, "timeline.html", gin.H{
		"Endpoint": "/",
		"User":     currUser,
		"Messages": messages,
	})
}

func renderPublicTimeline(context *gin.Context) {
	db := persistence.GetDbConnection()
	currUser := getCurrentUser(context, db)
	messages := application.GetFirstNMessages(db,30)
	context.HTML(http.StatusOK, "timeline.html", gin.H{
		"Endpoint": "public",
		"User": currUser,
		"Messages": messages,
	})
}

func renderUserTimeline(context *gin.Context) {
	db := persistence.GetDbConnection()
	username := context.Query("username")
	profileUser, err := application.GetUserByUsername(db, username)
	if err != nil {
		context.AbortWithError(http.StatusNotFound, err)
	}

	var endpoint string = "user-timeline"
	currUser := getCurrentUser(context, db)
	messages := application.GetMessagesByUserID(db, profileUser.ID)
	if currUser == nil {
		context.HTML(http.StatusOK, "timeline.html", gin.H{
			"Endpoint":    endpoint,
			"Messages":    messages,
			"ProfileUser": profileUser,
		})
		return
	}

	if currUser.ID == profileUser.ID {
		endpoint = "/"
	}

	followed := application.IsUserFollowing(db, currUser.ID, profileUser.ID)
	context.HTML(http.StatusOK, "timeline.html", gin.H{
		"Endpoint":    endpoint,
		"User":        currUser,
		"Messages":    messages,
		"ProfileUser": profileUser,
		"Followed":    followed,
	})
}

func getCurrentUser(context *gin.Context, db *gorm.DB) *application.User {
	currUserId := getCurrentUserId(context)
	if currUserId != 0 {
		user, _ := application.GetUserByID(db, getCurrentUserId(context))
		return &user
	}

	return nil
}
