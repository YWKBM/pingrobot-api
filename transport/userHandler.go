package transport

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"pingrobot-api.go/domain"
	"pingrobot-api.go/service"
)

type UserHandler struct {
	userService service.UserService
}

func newUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (uh *UserHandler) userSingUp(c *gin.Context) {
	var inp service.UserSignUpInput
	if err := c.BindJSON(&inp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)

		return
	}

	if err := uh.userService.SingUp(c, service.UserSignUpInput{
		Name:     inp.Name,
		Email:    inp.Email,
		Password: inp.Password,
	}); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)

		return
	}

	c.Status(http.StatusCreated)
}

func (uh *UserHandler) userSingIn(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	usr, err := uh.userService.SignIn(c, int64(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
	}

	resp, err := json.Marshal(&usr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
	}

	c.JSON(http.StatusFound, resp)
}

// webService should be created by user. Ideas: realization from routig, but there is shoud be a verification, the lower cod is for test functional
// TODO: watch text above
func (uh *UserHandler) userCreateWebService(c *gin.Context) {
	var inp domain.WebSerice
	if err := c.BindJSON(&inp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	err := uh.userService.CreateWebService(c, inp)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
	}

	c.JSON(200, nil)
}

func (uh *UserHandler) initUserRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/sign-up", uh.userSingUp)
		users.GET("/sign-in", uh.userSingIn) //Remember, only test functional
		users.POST("/add-service", uh.userCreateWebService)
	}
}
