package transport

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"pingrobot-api.go/service"
)

type WebServiceHandler struct {
	webServiceService service.WebServices
}

func newWebServiceHandler(webService service.WebServices) *WebServiceHandler {
	return &WebServiceHandler{
		webServiceService: webService,
	}
}

func (wh *WebServiceHandler) serviceGetAllWebServices(c *gin.Context) {
	resp, err := wh.webServiceService.GetAllWebServices(c)

	if err != nil{
		c.AbortWithStatusJSON(http.StatusNotFound, err)
	}

	c.JSON(200, resp)
}

func (wh *WebServiceHandler) initWebServicedRoutes(api *gin.RouterGroup) {
	webServices := api.Group("/web-service")
	{
		webServices.GET("/get-all", wh.serviceGetAllWebServices)
	}
}
