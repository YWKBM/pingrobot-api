package transport

import (
	"github.com/gin-gonic/gin"
	"pingrobot-api.go/service"
)

type Handler struct {
	webServiceHandler *WebServiceHandler
	auth              *AuthHandler
}

func NewHadnler(webService service.WebServices, authService service.Authorization) *Handler {
	return &Handler{
		webServiceHandler: newWebServiceHandler(webService),
		auth:              newAuthHadnler(authService),
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

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.auth.signUp)
		auth.POST("/sign-in", h.auth.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		h.webServiceHandler.initWebServicedRoutes(api)
	}

	router.Run()
}
