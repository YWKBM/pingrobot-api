package transport

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"pingrobot-api.go/domain"
	"pingrobot-api.go/service"
)

type AuthHandler struct {
	authService service.Authorization
}

func newAuthHadnler(authService service.Authorization) *AuthHandler {
	return &AuthHandler{authService}
}

func (a *AuthHandler) signUp(c *gin.Context) {
	var input domain.User

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := a.authService.CreateUser(input)
	if err != nil {
		fmt.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (a *AuthHandler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := a.authService.GenerateToken(input.Name, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "user not found")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
