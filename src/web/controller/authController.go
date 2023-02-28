package controller

import (
	"go-minitwit/src/application"
	"go-minitwit/src/persistence"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func MapAuthEndpoints(router *gin.Engine) {
	router.GET("/login", renderLoginPage)
	router.POST("/login", handleLogin)
	router.GET("/register-user", renderRegisterPage)
	router.POST("/register-user", handleRegister)
	router.GET("/logout", handleLogout)
}

func renderLoginPage(context *gin.Context) {
	if getCurrentUserId(context) != 0 {
		context.Redirect(http.StatusFound, "/")
		return
	}

	context.HTML(http.StatusOK, "login.html", gin.H{})
}

func handleLogin(context *gin.Context) {
	if err := application.HandleLogin(context, persistence.GetDbConnection(), sessions.Default(context)); err != nil {
		context.HTML(http.StatusBadRequest, "login.html", gin.H{"Error": err.Error()})
		return
	}
	context.Redirect(http.StatusFound, "/")
}

func renderRegisterPage(context *gin.Context) {
	if getCurrentUserId(context) != 0 {
		context.Redirect(http.StatusFound, "/login")
		return
	}

	context.HTML(http.StatusOK, "register.html", gin.H{})
}

func handleRegister(context *gin.Context) {
	if err := application.HandleRegister(context, persistence.GetDbConnection()); err != nil {
		context.HTML(http.StatusBadRequest, "register.html", gin.H{"Error": err.Error()})
		return
	}

	context.Redirect(http.StatusFound, "/login")
}

func handleLogout(context *gin.Context) {
	clearSession(context)
	context.Redirect(http.StatusFound, "/public")
}
