package web

import (
	"go-minitwit/src/web/controller"

	"github.com/gin-gonic/gin"
)

func ConfigureWeb(router *gin.Engine) {
	router.LoadHTMLGlob("./web/templates/*")
	router.Static("./web/static", "./web/static/")
	// configureSession must be called before mapAuthEndpoints
	controller.ConfigureSession(router)
	mapEndpoints(router)
}

func mapEndpoints(router *gin.Engine) {
	controller.MapAuthEndpoints(router)
	controller.MapUserEndpoints(router)
	controller.MapTimelineEndpoints(router)
	controller.MapMessageEndpoints(router)
}
