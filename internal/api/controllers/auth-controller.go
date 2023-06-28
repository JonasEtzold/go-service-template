package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gitlab.com/JonasEtzold/go-service-template/pkg/crypto"
	http_err "gitlab.com/JonasEtzold/go-service-template/pkg/http-err"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AuthToken string `json:"auth_token"`
}

// Login godoc
// @Summary      Logins an example API user
// @Description  provides an authorization token for any user with valid master password. Only dummy, don't use in production.
// @Produce      json
// @Param        username  body      string  true  "Username"
// @Param        password  body      string  true  "Password"
// @Success      200       {object}  string
// @Failure      403       {object}  http_err.HTTPError
// @Failure      500       {object}  http_err.HTTPError
// @Router       /api/v1/login [post]
// @Security     Authorization Token
func Login(c *gin.Context) {
	hashedMasterPassword, err := crypto.HashAndSalt([]byte(viper.GetString("server_default_password")))
	if err != nil {
		http_err.NewError(c, http.StatusInternalServerError, errors.New("authentication not configured"))
		return
	}

	var loginInput LoginInput
	_ = c.BindJSON(&loginInput)

	if !crypto.ComparePasswords(hashedMasterPassword, []byte(loginInput.Password)) {
		http_err.NewError(c, http.StatusForbidden, errors.New("wrong password"))
		return
	}

	token, _ := crypto.CreateToken(loginInput.Username)
	c.JSON(http.StatusOK, LoginResponse{token})
}
