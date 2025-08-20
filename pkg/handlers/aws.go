package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ListAWSClusters(c *gin.Context) {
	clusters, err := h.aws.ListEKSClusters(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"clusters": clusters})
}

func (h *Handler) ListK8sNamespaces(c *gin.Context) {
	ns, err := h.k8s.ListNamespaces()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"namespaces": ns})
}

func (h *Handler) CreateK8sDeployment(c *gin.Context) {
	var req struct {
		Namespace string `json:"namespace" binding:"required"`
		Name      string `json:"name" binding:"required"`
		Image     string `json:"image" binding:"required"`
		Replicas  int32  `json:"replicas"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Replicas == 0 {
		req.Replicas = 1
	}
	if err := h.k8s.CreateDeployment(req.Namespace, req.Name, req.Image, req.Replicas); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "created"})
}
