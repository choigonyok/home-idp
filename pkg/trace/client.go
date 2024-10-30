package trace

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/choigonyok/home-idp/pkg/model"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type TraceClient struct {
	Conn *grpc.ClientConn
}

type Trace struct {
	Conn      *grpc.ClientConn
	Trailer   metadata.MD
	Writer    http.ResponseWriter
	Request   *http.Request
	StartTime time.Time
	TraceID   string
	SpanID    string
	Error     error
	Context   context.Context
	Service   string
}

type Span1 struct {
	Context      context.Context
	SpanID       string
	TraceID      string
	ParentSpanID string
	Status       string
	StarTime     time.Time
	EndTime      time.Time
}

type TraceString string

type Trace1 struct {
	ID string
}

func NewTraceClient(port int) *TraceClient {
	grpcOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, _ := grpc.NewClient("localhost:"+strconv.Itoa(port), grpcOptions...)

	return &TraceClient{
		Conn: conn,
	}
}

func (c *TraceClient) Set(i interface{}) {
	c.Conn = parseKubeClientFromInterface(i).Conn
}

func parseKubeClientFromInterface(i interface{}) *TraceClient {
	client := i.(*TraceClient)
	return client
}

func (c *TraceClient) NewTrace(traceId string) *Span1 {
	return &Span1{
		SpanID:       uuid.NewString(),
		TraceID:      traceId,
		ParentSpanID: "",
	}
}

func (c *TraceClient) NewSpanFromOutgoingContext(ctx context.Context) *Span1 {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		fmt.Println("No metadata found in outgoing")
	}
	traceId := md["trace-id"]
	parentSpanId := md["parent-span-id"]

	fmt.Println("OUTTEST1:", traceId)
	fmt.Println("OUTTEST2:", parentSpanId)

	return &Span1{
		SpanID:       uuid.NewString(),
		TraceID:      traceId[0],
		ParentSpanID: parentSpanId[0],
	}
}

func (c *TraceClient) NewSpanFromIncomingContext(ctx context.Context) *Span1 {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("No metadata found in incoming")
	}
	traceId := md["trace-id"]
	parentSpanId := md["parent-span-id"]

	fmt.Println("INTEST1:", traceId)
	fmt.Println("INTEST2:", parentSpanId)

	return &Span1{
		SpanID:       uuid.NewString(),
		TraceID:      traceId[0],
		ParentSpanID: parentSpanId[0],
	}
}

type Key string

func (s *Span1) Start(ctx context.Context) error {
	// nextCtx := context.WithValue(ctx, Key("trace-id"), c.TraceID)
	// nextCtx = context.WithValue(nextCtx, Key("parent-span-id"), c.SpanID)
	md := metadata.Pairs(
		"trace-id", s.TraceID,
		"parent-span-id", s.SpanID,
	)
	c := metadata.NewOutgoingContext(ctx, md)
	s.Context = c

	// fmt.Println("STARTTEST1 ", c.TraceID)
	// fmt.Println("STARTTEST2 ", c.ParentSpanID)
	// fmt.Println("STARTTEST3 ", nextCtx.Value(Key("trace-id")))
	// fmt.Println("STARTTEST4 ", nextCtx.Value(Key("parent-span-id")))
	// fmt.Println("STARTTEST5 ", nextCtx.Value(Key("trace-id")))
	// fmt.Println("STARTTEST6 ", nextCtx.Value(Key("parent-span-id")))

	data := model.Trace{
		SpanID:       s.SpanID,
		TraceID:      s.TraceID,
		ParentSpanID: s.ParentSpanID,
		StartTime:    time.Now().Format("2006-01-02T15:04:05.999Z"),
		EndTime:      "",
		Status:       "START",
	}

	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://home-idp-trace-manager:5103/api/span", bytes.NewBuffer(jsonData))
	_, err := http.DefaultClient.Do(req)
	return err
}

func (c *Span1) Stop() error {
	data := model.Trace{
		SpanID:       c.SpanID,
		TraceID:      c.TraceID,
		ParentSpanID: c.ParentSpanID,
		StartTime:    "",
		EndTime:      time.Now().Format("2006-01-02T15:04:05.999Z"),
		Status:       "END",
	}

	jsonData, _ := json.Marshal(data)
	req, _ := http.NewRequest("PUT", "http://home-idp-trace-manager:5103/api/span", bytes.NewBuffer(jsonData))
	if _, err := http.DefaultClient.Do(req); err != nil {
		return err
	}

	c.Context.Done()
	return nil
}
