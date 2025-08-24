package server

import (
	"github.com/choigonyok/idp/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func setupRouter(engine *gin.Engine, h *handlers.Handler) {
	engine.Use(gin.Recovery())

	engine.Use(func(c *gin.Context) {
		h.Logger.Infof("Request: %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	{
		engine.OPTIONS("/")
	}
	{
		jenkins := engine.Group("/jenkins")
		// jenkins.Use(handlers.JWTMiddleware())
		{
			jobs := jenkins.Group("/jobs")
			jobs.GET("", h.ListJenkinsJobs)
			jobs.POST("/:jobName/build", h.BuildJenkinsJobs)
		}
	}
	{
		log := engine.Group("/log")
		// k8s.Use(handlers.JWTMiddleware())
		{
			k8s := log.Group("/k8s")
			k8s.GET("/pods/:name", h.LogsPod)
		}
		// log.GET("/aws", h.xxxx)
		// log.GET("/jenkins", h.xxxx)
	}
	{
		k8s := engine.Group("/k8s")
		// k8s.Use(handlers.JWTMiddleware())
		k8s.POST("/deploy", h.CreateK8sDeployment)
		k8s.GET("/namespace", h.ListK8sNamespaces)
		k8s.GET("/pod")
	}
	{
		traces := engine.Group("/traces")
		// k8s.Use(handlers.JWTMiddleware())

		traces.GET("/services/:serviceName", h.GetServiceTraces)
		traces.GET("/:traceId", h.GetTrace)
	}
	{
		grafana := engine.Group("/grafana")
		// grafana.Use(handlers.JWTMiddleware())
		grafana.POST("", h.QueryGrafana)
	}
}

// curl http://localhost:3200/api/search\?start\=$(date -v -15M +%s)\&end=$(date +%s)\&limit=20\&service=home-idp.test123.querygrafana

// curl http://localhost:3200/api/traces/90955ecb5225135cee22150ebbee1d97

// curl http://localhost:3200/api/spans/90955ecb5225135cee22150ebbee1d97
