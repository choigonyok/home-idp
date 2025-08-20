package handlers

import (
	"context"
	"log"
	"net/http"

	jenkinscli "github.com/choigonyok/idp/pkg/client/jenkins"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func initTracer() func() {
	ctx := context.Background()
	exp, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint("localhost:4317"),
	)
	if err != nil {
		log.Fatal(err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("home-idp.test123.querygrafana"),
		)),
	)
	otel.SetTracerProvider(tp)

	return func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}
}

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

// func test1(ctx context.Context) {
// 	tracer := otel.Tracer("test2-tracer")
// 	_, span := tracer.Start(ctx, "test2")
// 	defer span.End()

// 	for i := 0; i < 10; i++ {
// 		_, err := bcrypt.GenerateFromPassword([]byte("password"), 12) // cost=14
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	test2(ctx)
// }

// func test2(ctx context.Context) {
// 	tracer := otel.Tracer("test3-tracer")
// 	_, span := tracer.Start(ctx, "test3")
// 	defer span.End()

// 	a := big.NewInt(1)
// 	for i := 0; i < 200000; i++ {
// 		a.Mul(a, big.NewInt(int64(i+1)))
// 	}
// }
