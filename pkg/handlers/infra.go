package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/choigonyok/idp/pkg/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateInfra(c *gin.Context) {
	infra := model.Infra{}
	if err := c.BindJSON(&infra); err != nil {
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

	if err := h.infra.Create(&infra); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err,
			"message": "invalid data creation",
		})
		return
	}

	c.JSON(http.StatusOK, "service creation succeeded")
}

func (h *Handler) ListInfra(c *gin.Context) {
	result, err := h.infra.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	b, _ := json.Marshal(*result)

	c.Data(http.StatusOK, "application/json", b)
}
