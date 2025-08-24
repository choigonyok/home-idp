package handlers

import (
	"net/http"

	jenkinscli "github.com/choigonyok/idp/pkg/client/jenkins"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ListJenkinsJobs(c *gin.Context) {
	resp, err := h.jenkins.List(jenkinscli.Job)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, *resp)
}

func (h *Handler) BuildJenkinsJobs(c *gin.Context) {
	jobName := c.Param("jobName")
	m := make(map[string][]string)
	for k, v := range c.Request.URL.Query() {
		m[k] = v
	}
	resp, err := h.jenkins.Run(jenkinscli.Job, jobName, m)
	if err != nil {
		// fmt.Println("test3")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// fmt.Println("test2")
		return
	}
	// fmt.Println("test1")
	c.JSON(http.StatusOK, resp.Body)
}
