package transport

import (
	"github.com/gin-gonic/gin"
	"pingrobot-api.go/service"
)

type Handler struct {
	auth              *AuthHandler
	webServiceHandler *WebServiceHandler
}

func NewHadnler(webService service.WebServices, authService service.AuthService) *Handler {
	return &Handler{
		auth:              newAuthHadnler(authService),
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
		h.webServiceHandler.initWebServicedRoutes(api)
	}

	router.Run()
}
