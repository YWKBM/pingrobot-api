package transport

import (
	"github.com/gin-gonic/gin"
	"pingrobot-api.go/service"
)

type Handler struct {
	userHadnler       *UserHandler
	webServiceHandler *WebServiceHandler
}

func NewHadnler(userService service.UserService, webService service.WebSericeService) *Handler {
	return &Handler{
		userHadnler:       newUserHandler(userService),
		webServiceHandler: newWebServiceHandler(webService),
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.userHadnler.initUserRoutes(v1)
		h.webServiceHandler.initWebServicedRoutes(v1)
	}
}
