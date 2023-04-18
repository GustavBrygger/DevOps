package controller

import (
	"encoding/json"
	"go-minitwit/src/application"
	"go-minitwit/src/persistence"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MapJSONFollowersEndpoints(router *gin.Engine) {
	router.GET("/fllws/:username", jsonGetFollowersToUser)
	router.POST("/fllws/:username", jsonfollowUser)
}

func jsonfollowUser(context *gin.Context) {
	updateLatest(context.Request)

	userNameToFollow := context.Param("username")

	// Read body and convert from byteArray => string  => JSON
	bodyBites, err := ioutil.ReadAll(context.Request.Body)
	bodyString := string(bodyBites)
	var bodyJson map[string]interface{}
	jsonError := json.Unmarshal([]byte(bodyString), &bodyJson)
	if err != nil || jsonError != nil {
		context.AbortWithStatus(404)
	}

	// Check if we need to follow or unFollow
	followUsername, isFollowInBody := bodyJson["follow"]
	unfollowUsername := bodyJson["unfollow"]

	if isFollowInBody {
		user, err := application.GetUserByUsername(persistence.GetDbConnection(), followUsername.(string))
		if err != nil {
			context.AbortWithError(http.StatusInternalServerError, err)
		}
		userID := user.ID

		errs := application.FollowUser(persistence.GetDbConnection(), userID, userNameToFollow)
		if errs != nil {
			context.AbortWithError(http.StatusUnauthorized, errs)
		}
	} else {
		user, err := application.GetUserByUsername(persistence.GetDbConnection(), unfollowUsername.(string))
		if err != nil {
			context.AbortWithError(http.StatusInternalServerError, err)
		}
		userID := user.ID

		errs := application.UnfollowUser(persistence.GetDbConnection(), userID, userNameToFollow)
		if errs != nil {
			context.AbortWithError(http.StatusUnauthorized, errs)
		}
	}
	context.Status(http.StatusNoContent)
}

func jsonGetFollowersToUser(context *gin.Context) {
	updateLatest(context.Request)

	db := persistence.GetDbConnection()

	userName := context.Param("username")

	user, err := application.GetUserByUsername(persistence.GetDbConnection(), userName)
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
	}

	userID := user.ID
	limitToQuery := context.Request.URL.Query().Get("no")
	limitToQueryInt, err := strconv.Atoi(limitToQuery)
	if err != nil {
		limitToQueryInt = 100
	}

	users, err := application.GetFirstNFollowersToUserid(db, userID, uint(limitToQueryInt))
	if err != nil {
		context.AbortWithError(http.StatusUnauthorized, err)
	}

	userNameListToReturn := []string{}
	for _, user := range users {
		userNameListToReturn = append(userNameListToReturn, user.Username)
	}

	usernames, _ := json.Marshal(map[string]interface{}{"follows": userNameListToReturn})
	context.Writer.Write(usernames)
}
