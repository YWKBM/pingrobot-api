package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
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
	m := ginmetrics.GetMonitor()

	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	// set middleware for gin
	m.Use(router)

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
