package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/choigonyok/idp/pkg/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateService(c *gin.Context) {
	svc := model.Service{}
	if err := c.BindJSON(&svc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err,
			"message": "invalid binding json request",
		})
		return
	}

	// to do
	// platform/platformName(namespace)에 해당하는 서비스가 실존하는지 검증
	// git url 존재하는지 검증
	// 서비스명 중복아닌지 검증

	if err := h.service.Create(&svc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "invalid data creation",
		})
		return
	}

	c.JSON(http.StatusOK, "service creation succeeded")
}

func (h *Handler) GetServices(c *gin.Context) {
	result, err := h.service.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	b, _ := json.Marshal(result)

	c.Data(http.StatusOK, "application/json", b)
}

func (h *Handler) GetServiceLogs(c *gin.Context) {
	err := h.o11y.GetServiceLog(c.Param("serviceID"), c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil) // gin.H{"clusters": clusters}
}

func (h *Handler) GetServiceMetrics(c *gin.Context) {
	result, err := h.o11y.GetServiceMetrics(c.Param("serviceID"), c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", []byte(result.String()))
}

func (h *Handler) GetServiceTraces(c *gin.Context) {
	result, err := h.o11y.GetServiceTraces(c.Param("serviceID"), c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	b, _ := json.Marshal(result)
	c.Data(http.StatusOK, "application/json", b)
}
