package service

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthHandler for sign in
type AuthHandler struct {
	res Response
}

// SignInRequest contains sign in information
type SignInRequest struct {
	Username string `json:"username" example:"demo"`
	Password string `json:"password" example:"password"`
}

// SignIn with username and password
// @Summary Sign In
// @Schemes
// @Description Sign in into service
// @Tags security
// @Accept json
// @Produce json
// @Param signInRequest body service.SignInRequest true "Sign in"
// @Success 200 {object} service.Response
// @Failure 500 {object} service.Response
// @Router /signin [post]
func (a AuthHandler) SignIn(c *gin.Context) {
	var req SignInRequest

	err := c.BindJSON(&req)

	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, a.res.internal())
		return
	}

	if req.Username == "demo" && req.Password == "123456" {
		session := sessions.Default(c)
		session.Set("demo", true)
		session.Save()
		c.JSON(http.StatusOK, a.res.success())
		return
	}

	c.AbortWithStatusJSON(http.StatusBadRequest, a.res.error("incorrectCredentials"))

}
