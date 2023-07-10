package controllers

import (
	"net/http"
	"ramblr/query"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u *UserController) Mount(group gin.IRouter) {
	group.POST("/login", u.Login)
}

func (u *UserController) Login(c *gin.Context) {
	session := sessions.Default(c)

	var user = query.User
	var req = struct {
		username string
		password string
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	match, err := user.Where(user.Username.Eq(req.username)).First()
	if match == nil || err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	// Return the same response for both email and password failure
	if err = match.Authenticate(req.password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	session.Set("user", match)
	c.JSON(http.StatusOK, match)
}
