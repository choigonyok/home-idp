package handlers

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	otelpyroscope "github.com/grafana/otel-profiling-go"
	"github.com/grafana/pyroscope-go"
	"github.com/spf13/viper"

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

func (h *Handler) GetServiceTraces(c *gin.Context) {
	v := viper.GetViper()
	if !v.GetBool("grafana.dataSources.tempo.enabled") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tempo is not enabled"})
		return
	}
	serviceName := c.Param("serviceName")
	now := time.Now().Unix()
	now15mAgo := time.Now().Add(-15 * time.Minute).Unix()
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+h.grafana.Token).
		Get(h.grafana.Host + "/api/datasources/proxy/4/api/search?start=" + strconv.Itoa(int(now15mAgo)) + "&end=" + strconv.Itoa(int(now)) + "&limit=5&sort=startTime&direction=desc&service=" + serviceName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Grafana 요청 실패"})
		fmt.Println(err)
		return
	}

	traces := Traces{}

	json.Unmarshal(resp.Body(), &traces)

	for _, t := range traces.Traces {
		spans := Spans{}
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("Authorization", "Bearer "+h.grafana.Token).
			Get(h.grafana.Host + "/api/datasources/proxy/4/api/traces/" + t.TraceID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Grafana 요청 실패"})
			return
		}
		json.Unmarshal(resp.Body(), &spans)
		t.Spans = spans

		for _, b := range t.Spans.Batches {
			for _, s := range b.ScopeSpans {
				for _, ss := range s.Spans {
					s1, err := base64.StdEncoding.DecodeString(ss.SpanID)
					if err != nil {
						fmt.Println(err)
						return
					}
					s2, err := base64.StdEncoding.DecodeString(ss.TraceID)
					if err != nil {
						fmt.Println(err)
						return
					}

					b1 := hex.EncodeToString(s1)
					b2 := hex.EncodeToString(s2)
					ss.SpanID = b1
					ss.TraceID = b2

					resp, err := client.R().
						SetHeader("Content-Type", "application/json").
						SetHeader("Authorization", "Bearer "+h.grafana.Token).
						SetQueryParam("query", `process_cpu:cpu:nanoseconds:cpu:nanoseconds{service_name="tempo.trace.spans", trace_id="`+ss.TraceID+`"}`).
						SetQueryParam("from", "now-1h").
						Get(h.grafana.Host + "/api/datasources/proxy/3/pyroscope/render")
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": "Grafana 요청 실패"})
						return
					}

					treeNode := getTree(resp.Body())

					t.Profile = treeNode
				}
			}
		}
	}
	b, _ := json.Marshal(traces)

	c.Data(resp.StatusCode(), "application/json", b)
}

func (h *Handler) GetTrace(c *gin.Context) {}

type Flamebearer struct {
	Names  []string    `json:"names"`
	Levels [][]float64 `json:"levels"`
}

type PyroscopeData struct {
	Version     int         `json:"version"`
	Flamebearer Flamebearer `json:"flamebearer"`
}

// ---------- 출력(트리) 구조 ----------
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

func getTree(b []byte) *TreeNode {
	var data PyroscopeData
	if err := json.Unmarshal(b, &data); err != nil {
		fmt.Printf("unmarshal err: %v\n", err)
		return nil
	}

	root := BuildFlameTree(data.Flamebearer.Levels, data.Flamebearer.Names)
	if root == nil {
		fmt.Println(`{}`)
		return nil
	}
	return root
	// out, _ := json.MarshalIndent(, "", "  ")
	// fmt.Println(string(out))
}

// BuildFlameTree: Pyroscope v1 levels -> 계층 트리
func BuildFlameTree(levels [][]float64, names []string) *TreeNode {
	if len(levels) == 0 {
		return nil
	}

	// 1) 각 레벨을 절대 좌표(start/end)로 디코딩
	decoded := make([][]decodedBlock, 0, len(levels))
	for _, lvl := range levels {
		cursor := int64(0)
		row := make([]decodedBlock, 0, len(lvl)/4)

		for i := 0; i+3 < len(lvl); i += 4 {
			deltaX := int64(lvl[i])  // "간격" (이전 블록 끝 기준)
			total := int64(lvl[i+1]) // width
			self := int64(lvl[i+2])  // self time
			nameIdx := int(lvl[i+3]) // name index

			name := "unknown"
			if nameIdx >= 0 && nameIdx < len(names) {
				name = names[nameIdx]
			}

			start := cursor + deltaX
			end := start + total
			cursor = end

			node := &TreeNode{
				Name:     name,
				Value:    total,
				SelfTime: self,
				Start:    start,
				End:      end,
			}
			row = append(row, decodedBlock{
				start: start,
				end:   end,
				node:  node,
			})
		}
		decoded = append(decoded, row)
	}

	// 2) 부모-자식 매핑 (다음 레벨의 블록이 부모 구간 안에 있으면 자식으로 연결)
	for l := 0; l < len(decoded)-1; l++ {
		parents := decoded[l]
		children := decoded[l+1]
		j := 0

		for i := 0; i < len(parents); i++ {
			p := parents[i]

			// 자식 포인터를 부모 시작보다 앞인 것들은 스킵
			for j < len(children) && children[j].end <= p.start {
				j++
			}
			// 부모 범위에 들어오는 애들만 붙임
			for j < len(children) && children[j].start >= p.start && children[j].end <= p.end {
				p.node.Children = append(p.node.Children, children[j].node)
				j++
			}
		}
	}

	// 3) 루트 반환 (일반적으로 level 0 첫 블록이 total)
	if len(decoded[0]) == 0 {
		return nil
	}
	return decoded[0][0].node
}
