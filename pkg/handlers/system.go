package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/choigonyok/idp/pkg/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateSystem(c *gin.Context) {
	sys := model.System{
		Type: c.Param("type"),
	}
	if err := c.BindJSON(&sys); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// json 형식 검증
	if sys.Type == "permission" {
		f := model.JSONFormat{}
		if err := json.Unmarshal([]byte(sys.JSON), &f); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json format"})
			return
		}
		if f.Act == "" || f.Obj == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json format"})
			return
		}
	}

	if err := h.system.Create(&sys); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "success"})
}

func (h *Handler) ListSystem(c *gin.Context) {
	result, err := h.system.List(c.Param("type"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	b, _ := json.Marshal(result)

	c.Data(http.StatusOK, "application/json", b)
}

func (h *Handler) UpdateSystem(c *gin.Context) {
	sys := model.System{
		Type: c.Param("type"),
	}
	if err := c.BindJSON(&sys); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.system.Update(&sys); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "success"})
}

func (h *Handler) ListRolePermissions(c *gin.Context) {
	result, err := h.system.ListRolePermissions("permissions")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	b, _ := json.Marshal(result)

	c.Data(http.StatusOK, "application/json", b)
}
