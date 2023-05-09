package controller

import (
	"encoding/json"
	"fmt"
	"go-minitwit/src/application"
	"go-minitwit/src/persistence"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MapJSONMessageEndpoints(router *gin.Engine) {
	router.GET("/msgs", getNMessagesJSON)
	router.GET("/msgs/:username", getNUserMessagesJSON)
	router.POST("/msgs/:username", postMessageAsUser)
	router.GET("/latest", getLatest)
}

func getNMessagesJSON(context *gin.Context) {
	no_query := context.Request.URL.Query().Get("no")
	n, err := strconv.ParseInt(no_query, 10, 64)
	if err != nil {
		n = 100
	}

	messages := application.GetFirstNMessages(persistence.GetDbConnection(), int(n))
	msgs_json, err := json.Marshal(messages)
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
	}

	updateLatest(context.Request)
	context.Writer.Write(msgs_json)
}

func getNUserMessagesJSON(context *gin.Context) {
	username := context.Param("username")
	no_query := context.Request.URL.Query().Get("no")
	n, err := strconv.ParseInt(no_query, 10, 64)
	if err != nil {
		n = 100
	}
	messages := application.GetNMessagesByUsername(persistence.GetDbConnection(), username, int(n))
	msgs_json, err := json.Marshal(messages)
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
	}
	updateLatest(context.Request)
	context.Writer.Write(msgs_json)
}

func postMessageAsUser(context *gin.Context) {
	username := context.Param("username")
	var messageText map[string]string
	context.BindJSON(&messageText)

	user, err := application.GetUserByUsername(persistence.GetDbConnection(), username)
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
	}

	application.AddMessage(persistence.GetDbConnection(), user.ID, messageText["content"])
	updateLatest(context.Request)
	context.Status(http.StatusNoContent)
}

func getLatest(context *gin.Context) {
	redisConn := persistence.GetRedisConnection()
	latest, err := redisConn.Get("latest").Result()

	if err != nil {
		fmt.Println(err)
	}

	latestInt, _ := strconv.Atoi(latest)
	latestJson, _ := json.Marshal(map[string]int{"latest": latestInt})
	context.Writer.Write(latestJson)
}

func updateLatest(request *http.Request) {
	latest_query, _ := strconv.ParseInt(request.URL.Query().Get("latest"), 10, 64)
	redisConn := persistence.GetRedisConnection()

	if latest_query != -1 {
		err := redisConn.Set("latest", latest_query, 0).Err()
		if err != nil {
			fmt.Println(err)
		}
	}
}
