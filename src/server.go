package main

import (
	"go-minitwit/src/persistence"
	"go-minitwit/src/web"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	persistence.ConfigurePersistence()
	web.ConfigureWeb(router)
	router.Run()
}
