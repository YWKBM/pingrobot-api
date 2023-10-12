package transport

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"pingrobot-api.go/domain"
)

type WebService interface {
	Create(ctx context.Context, webService domain.WebSerice) error
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

func (wh *WebServiceHandler) create(c *gin.Context) error {
	var inp domain.WebSerice
	if err := c.BindJSON(&inp); err != nil {
		return err
	}

	err := wh.webServiceService.Create(c, inp)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
	}
	c.JSON(200, nil)

	return nil
}

func (wh *WebServiceHandler) getWebServiceByUserID(c *gin.Context) error {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		return err
	}

	resp, err := wh.webServiceService.GetWebServiceByUserId(c, int64(id))
	if err != nil {
		c.AbortWithStatusJSON(404, err)
	}
	c.JSON(200, resp)

	return nil
}
