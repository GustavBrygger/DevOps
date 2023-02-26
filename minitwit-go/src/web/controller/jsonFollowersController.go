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

	context.JSON(204, map[string]interface{}{"Hello": "Gusatv"})

	userID := abortIfNoUserID(context)
	if userID == 0 {
		context.AbortWithStatus(404)
	}

	//Read body and convert form byteArray => string  => JSON
	bodyBites, err := ioutil.ReadAll(context.Request.Body)
	bodyString := string(bodyBites)
	var bodyJson map[string]interface{}
	jsonError := json.Unmarshal([]byte(bodyString), &bodyJson)
	if err != nil || jsonError != nil {
		context.AbortWithStatus(404)
	}

	//Check if we need to follow or unFollow
	followUsername, isFollowInBody := bodyJson["follow"]
	unfollowUsername, _ := bodyJson["unfollow"]

	if isFollowInBody {
		err := application.FollowUser(persistence.GetDbConnection(), userID, followUsername.(string))
		if err != nil {
			context.AbortWithError(http.StatusUnauthorized, err)
		}
	} else {
		err := application.UnfollowUser(persistence.GetDbConnection(), userID, unfollowUsername.(string))
		if err != nil {
			context.AbortWithError(http.StatusUnauthorized, err)
		}
	}

}

func jsonGetFollowersToUser(context *gin.Context) {
	updateLatest(context.Request)

	db := persistence.GetDbConnection()

	userID := abortIfNoUserID(context)
	if userID == 0 {
		context.AbortWithStatus(404)
	}

	limitToQuery := context.Request.URL.Query().Get("no")
	limitToQueryInt, _ := strconv.Atoi(limitToQuery)

	users, err := application.GetFirstNFollowersToUserid(db, userID, uint(limitToQueryInt))
	if err != nil {
		context.AbortWithError(http.StatusUnauthorized, err)
	}

	userNameListToReturn := []string{}
	for _, user := range users {
		userNameListToReturn = append(userNameListToReturn, user.Username)
	}

	usernames, err := json.Marshal(userNameListToReturn)
	context.Writer.Write(usernames)
}
