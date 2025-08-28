package handlers

import (
	"time"

	"github.com/go-resty/resty/v2"
	otelpyroscope "github.com/grafana/otel-profiling-go"
	"github.com/grafana/pyroscope-go"

	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type Profile struct {
	Version     int         `json:"version"`
	FlameBearer FlameBearer `json:"flamebearer"`
	Metadata    Metadata    `json:"metadata"`
	Timeline    Timeline    `json:"timeline"`
	Groups      bool        `json:"groups"`
	Heatmap     bool        `json:"heatmap"`
}

type FlameBearer struct {
	Names    []string `json:"names"`
	Levels   [][]int  `json:"levels"`
	NumTicks int      `json:"numTicks"`
	MaxSelf  int      `json:"maxSelf"`
}

type Metadata struct {
	Format     string `json:"format"`
	SpyName    string `json:"spyName"`
	SampleRate int    `json:"sampleRate"`
	Units      string `json:"units"`
	Name       string `json:"name"`
}

type Timeline struct {
	StartTime     int   `json:"startTime"`
	Samples       []int `json:"samples"`
	DurationDelta int   `json:"durationDelta"`
	Watermarks    bool  `json:"watermarks"`
}

type Traces struct {
	Traces []*Trace `json:"traces"`
}

type Trace struct {
	TraceID           string    `json:"traceID"`
	RootServiceName   string    `json:"rootServiceName"`
	RootTraceName     string    `json:"rootTraceName"`
	StartTimeUnixNano string    `json:"startTimeUnixNano"`
	DurationMs        string    `json:"durationMs"`
	Spans             Spans     `json:"spans"`
	Profile           *TreeNode `json:"profile"`
}

type Spans struct {
	Batches []struct {
		Resource struct {
			Attributes []struct {
				Key   string            `json:"key"`
				Value map[string]string `json:"value"`
			} `json:"attributes"`
		} `json:"resource"`
		ScopeSpans []struct {
			Scope struct {
				Name string `json:"name"`
			} `json:"scope"`
			Spans []*struct {
				TraceID           string `json:"traceId"`
				SpanID            string `json:"spanId"`
				Name              string `json:"name"`
				Kind              string `json:"kind"`
				StartTimeUnixNano string `json:"startTimeUnixNano"`
				EndTimeUnixNano   string `json:"endTimeUnixNano"`
				Status            string `json:"status"`
			}
		} `json:"scopeSpans"`
	} `json:"batches"`
}

type TreeNode struct {
	Name     string      `json:"name"`
	Value    int64       `json:"value"`    // totalTime
	SelfTime int64       `json:"selfTime"` // selfTime
	Children []*TreeNode `json:"children,omitempty"`
	// 디버깅용(원하지 않으면 지워도 됨)
	Start int64 `json:"start,omitempty"` // 절대 시작 위치(ns)
	End   int64 `json:"end,omitempty"`   // 절대 끝 위치(ns)
}

type decodedBlock struct {
	start int64
	end   int64
	node  *TreeNode
}

func initTracer() *sdktrace.TracerProvider {
	ctx := context.Background()
	exp, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint("localhost:4317"),
	)
	if err != nil {
		log.Fatal(err)
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("home-idp.tester.query"),
		)),
	)
}

func test1(ctx context.Context) {
	ctx, span := otel.Tracer("function/test1").Start(ctx, "Internal Caculation")
	defer span.End()

	p, _ := pyroscope.Start(pyroscope.Config{
		ApplicationName: "tempo.trace.spans",
		ServerAddress:   "http://localhost:4040",
		Tags: map[string]string{
			"trace_id": span.SpanContext().TraceID().String(),
			"span_id":  span.SpanContext().SpanID().String(),
		},
	})
	defer p.Stop()

	end := time.Now().Add(3 * time.Second)

	for {
		if time.Now().After(end) {
			break
		}
	}

	test2(ctx)
}

func test2(ctx context.Context) {
	ctx, span := otel.Tracer("function/test2").Start(ctx, "Internal Sleep")
	defer span.End()

	p, _ := pyroscope.Start(pyroscope.Config{
		ApplicationName: "tempo.trace.spans",
		ServerAddress:   "http://localhost:4040",
		Tags: map[string]string{
			"trace_id": span.SpanContext().TraceID().String(),
			"span_id":  span.SpanContext().SpanID().String(),
		},
	})
	defer p.Stop()

	end := time.Now().Add(2 * time.Second)

	for {
		if time.Now().After(end) {
			break
		}
	}
}

func (h *Handler) QueryGrafana(c *gin.Context) {
	tp := initTracer()
	otel.SetTracerProvider(otelpyroscope.NewTracerProvider(tp))
	defer tp.Shutdown(c.Request.Context())

	ctx := c.Request.Context()
	ctx, span := otel.Tracer("yunsuk-test-tracer").Start(ctx, "HTTP Request")
	defer span.End()

	p, _ := pyroscope.Start(pyroscope.Config{
		ApplicationName: "tempo.trace.spans",
		ServerAddress:   "http://localhost:4040",
		Tags: map[string]string{
			"trace_id": span.SpanContext().TraceID().String(),
			"span_id":  span.SpanContext().SpanID().String(),
		},
	})
	defer p.Stop()

	test1(ctx)

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

// func (h *Handler) GetServiceTrace(c *gin.Context) {
// 	v := viper.GetViper()
// 	if !v.GetBool("grafana.dataSources.tempo.enabled") {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "tempo is not enabled"})
// 		return
// 	}
// 	serviceName := c.Param("serviceName")
// 	now := time.Now().Unix()
// 	now15mAgo := time.Now().Add(-15 * time.Minute).Unix()
// 	client := resty.New()
// 	resp, err := client.R().
// 		SetHeader("Content-Type", "application/json").
// 		SetHeader("Authorization", "Bearer "+h.grafana.Token).
// 		Get(h.grafana.Host + "/api/datasources/proxy/4/api/search?start=" + strconv.Itoa(int(now15mAgo)) + "&end=" + strconv.Itoa(int(now)) + "&limit=5&sort=startTime&direction=desc&service=" + serviceName)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Grafana 요청 실패"})
// 		fmt.Println(err)
// 		return
// 	}

// 	traces := Traces{}

// 	json.Unmarshal(resp.Body(), &traces)

// 	for _, t := range traces.Traces {
// 		spans := Spans{}
// 		resp, err := client.R().
// 			SetHeader("Content-Type", "application/json").
// 			SetHeader("Authorization", "Bearer "+h.grafana.Token).
// 			Get(h.grafana.Host + "/api/datasources/proxy/4/api/traces/" + t.TraceID)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Grafana 요청 실패"})
// 			return
// 		}
// 		json.Unmarshal(resp.Body(), &spans)
// 		t.Spans = spans

// 		for _, b := range t.Spans.Batches {
// 			for _, s := range b.ScopeSpans {
// 				for _, ss := range s.Spans {
// 					s1, err := base64.StdEncoding.DecodeString(ss.SpanID)
// 					if err != nil {
// 						fmt.Println(err)
// 						return
// 					}
// 					s2, err := base64.StdEncoding.DecodeString(ss.TraceID)
// 					if err != nil {
// 						fmt.Println(err)
// 						return
// 					}

// 					b1 := hex.EncodeToString(s1)
// 					b2 := hex.EncodeToString(s2)
// 					ss.SpanID = b1
// 					ss.TraceID = b2

// 					resp, err := client.R().
// 						SetHeader("Content-Type", "application/json").
// 						SetHeader("Authorization", "Bearer "+h.grafana.Token).
// 						SetQueryParam("query", `process_cpu:cpu:nanoseconds:cpu:nanoseconds{service_name="tempo.trace.spans", trace_id="`+ss.TraceID+`"}`).
// 						SetQueryParam("from", "now-1h").
// 						Get(h.grafana.Host + "/api/datasources/proxy/3/pyroscope/render")
// 					if err != nil {
// 						c.JSON(http.StatusInternalServerError, gin.H{"error": "Grafana 요청 실패"})
// 						return
// 					}

// 					treeNode := getTree(resp.Body())

// 					t.Profile = treeNode
// 				}
// 			}
// 		}
// 	}
// 	b, _ := json.Marshal(traces)

// 	c.Data(resp.StatusCode(), "application/json", b)
// }

// func (h *Handler) GetTrace(c *gin.Context) {}
