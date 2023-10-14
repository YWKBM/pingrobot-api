package transport

import (
	"net/http"
	"strconv"

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

func (wh *WebServiceHandler) serviceGetWebServiceByUserID(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
	}

	resp, err := wh.webServiceService.GetWebServiceByUserId(c, int64(id))
	if err != nil {
		c.AbortWithStatusJSON(404, err)
	}

	c.JSON(200, resp)
}

func (wh *WebServiceHandler) initWebServicedRoutes(api *gin.RouterGroup) {
	webServices := api.Group("/web-service")
	{
		webServices.GET("/id:}", wh.serviceGetWebServiceByUserID)
	}
}
