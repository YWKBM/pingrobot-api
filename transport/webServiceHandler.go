package transport

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"pingrobot-api.go/domain"
	"pingrobot-api.go/service"
)

type WebServiceHandler struct {
	webServiceService service.WebServices
}

func newWebServiceHandler(webServiceService service.WebServices) *WebServiceHandler {
	return &WebServiceHandler{webServiceService}
}

func (wh *WebServiceHandler) createWebService(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(500, http.StatusInternalServerError)
	}

	var input domain.WebService
	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(400, http.StatusBadRequest)
		return
	}

	id, err := wh.webServiceService.Create(userId, input)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(400, http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (wh *WebServiceHandler) getAllWebServices(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(500, http.StatusInternalServerError)
		return
	}

	webServices, err := wh.webServiceService.GetAll(userId)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(500, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, webServices)
}

func (wh *WebServiceHandler) getWebServiceById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(500, http.StatusInternalServerError)
		return
	}

	webServiceId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(400, http.StatusBadRequest)
		return
	}

	webService, err := wh.webServiceService.GetById(userId, webServiceId)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(500, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, webService)
}

func (wh *WebServiceHandler) updateWebService(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(500, http.StatusInternalServerError)
		return
	}

	webServiceId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(400, http.StatusBadRequest)
		return
	}

	var input domain.UpdateWebServiceInput
	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(400, http.StatusBadRequest)
		return
	}

	if err := wh.webServiceService.Update(userId, webServiceId, input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(400, http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, "OK")
}

func (wh *WebServiceHandler) deleteWebService(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(500, http.StatusInternalServerError)
		return
	}

	webServiceId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(400, http.StatusBadRequest)
		return
	}

	if err := wh.webServiceService.Delete(userId, webServiceId); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(500, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, "OK")
}

func (wh *WebServiceHandler) initWebServicedRoutes(api *gin.RouterGroup) {
	webServices := api.Group("/web-service")
	{
		webServices.POST("/", wh.createWebService)
		webServices.GET("/", wh.getAllWebServices)
		webServices.GET("/:id", wh.getWebServiceById)
		webServices.PUT("/:id", wh.updateWebService)
		webServices.DELETE("/:id", wh.deleteWebService)
	}
}
