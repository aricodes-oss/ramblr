package server

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"ramblr/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	// TODO: Actually make a secret
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("ramblr", store))

	users := &controllers.UserController{}

	v1 := router.Group("v1")
	userGroup := v1.Group("users")
	users.Mount(userGroup)

	return router
}
