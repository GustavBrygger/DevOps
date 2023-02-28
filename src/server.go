package main

import (
	"go-minitwit/src/persistence"
	"go-minitwit/src/web"

	"github.com/gin-gonic/gin"
)

type Config struct {
	DbServer   string `env:"ENVIRONMENT,required"`
	DbPort     string `env:"ENVIRONMENT,required"`
	DbName     string `env:"ENVIRONMENT,required"`
	DbUser     string `env:"ENVIRONMENT,required"`
	DbPassword string `env:"ENVIRONMENT,required"`
}

func main() {
	router := gin.Default()

	config := Config{}
	err = env.Parse(&cfg) // 👈 Parse environment variables into `Config`
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}

	println(config.DbServer)
	persistence.ConfigurePersistence()
	web.ConfigureWeb(router)

	router.Run()
}
