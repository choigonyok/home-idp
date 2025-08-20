package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/grafana/pyroscope-go"
	"go.opentelemetry.io/otel"
)

func (h *Handler) QueryGrafana(c *gin.Context) {
	shutdown := initTracer()
	defer shutdown()
	ctx := c.Request.Context()
	tracer := otel.Tracer("example-test-tracer")

	// 최상위 Span
	ctx, span := tracer.Start(ctx, "HTTP Request")
	defer span.End()
	traceID := span.SpanContext().TraceID().String()

	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: "test123",
		ServerAddress:   "http://localhost:4040",
		Tags: map[string]string{
			"trace_id": traceID,
		},
	})
	if err != nil {
		panic(err)
	}

	// test1(ctx)

	var body struct {
		Query string `json:"query"`
		Range string `json:"range"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+h.grafana.Token).
		SetBody(map[string]interface{}{
			"queries": []map[string]interface{}{
				{
					"refId": "A",
					"datasource": map[string]string{
						"type": "prometheus",
						"uid":  (*h.grafana.DataSourcesUID)["prometheus"],
					},
					"expr": body.Query,
				},
			},
			"from": "now-" + body.Range,
			"to":   "now",
		}).
		Post(h.grafana.Host + "/api/ds/query")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Grafana 요청 실패"})
		return
	}
	// fmt.Println(resp.Body())
	c.Data(resp.StatusCode(), "application/json", resp.Body())
}
