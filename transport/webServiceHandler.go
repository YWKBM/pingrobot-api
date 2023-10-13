package transport

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"pingrobot-api.go/domain"
)

type WebService interface {
	GetWebServiceByUserId(ctx context.Context, id int64) (domain.WebSerice, error)
}

type WebServiceHandler struct {
	webServiceService WebService
}

func NewWebServiceHandler(webService WebService) *WebServiceHandler {
	return &WebServiceHandler{
		webServiceService: webService,
	}
}

func (wh *WebServiceHandler) getWebServiceByUserID(c *gin.Context) {
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
