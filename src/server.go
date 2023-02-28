package main

import (
	"go-minitwit/src/persistence"
	"go-minitwit/src/web"
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	config := persistence.Config{}
	err = env.Parse(&config)
	if err != nil {
		log.Fatal("unable to parse ennvironment variables: %e", err)
	}

	persistence.ConfigurePersistence(config)
	web.ConfigureWeb(router)
	router.Run()
}
