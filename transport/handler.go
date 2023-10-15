package transport

import (
	"github.com/gin-gonic/gin"
	"pingrobot-api.go/service"
)

type Handler struct {
	userHadnler       *UserHandler
	webServiceHandler *WebServiceHandler
}

func NewHadnler(userService service.Users, webService service.WebServices) *Handler {
	return &Handler{
		userHadnler:       newUserHandler(userService),
		webServiceHandler: newWebServiceHandler(webService),
	}
}

func (h *Handler) Init() {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
	)

	h.initApi(router)
}

func (h *Handler) initApi(router *gin.Engine) {
	api := router.Group("/api")
	{
		h.userHadnler.initUserRoutes(api)
		h.webServiceHandler.initWebServicedRoutes(api)
	}

	router.Run()
}
