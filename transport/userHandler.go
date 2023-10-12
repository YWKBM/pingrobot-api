package transport

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"pingrobot-api.go/domain"
)

type User interface {
	Create(ctx context.Context, user domain.User) error
	GetUserById(ctx context.Context, id int64) (domain.User, error)
}

type UserHandler struct {
	userService User
}

func NewUserHandler(userService User) *UserHandler {
	return &UserHandler{userService: userService}
}

func (uh *UserHandler) Create(c *gin.Context) error {
	var inp domain.User
	if err := c.BindJSON(&inp); err != nil {
		return err
	}

	err := uh.userService.Create(c, inp)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
	}
	c.JSON(200, nil)

	return nil
}

func (uh *UserHandler) GetUserById(c *gin.Context) error {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		return err
	}

	resp, err := uh.userService.GetUserById(c, int64(id))
	if err != nil {
		c.AbortWithStatusJSON(404, err)
	}

	c.JSON(200, resp)

	return nil
}
