package model

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

type DecodedBlock struct {
	Start int64
	End   int64
	Node  *TreeNode
}

type Flamebearer struct {
	Names  []string    `json:"names"`
	Levels [][]float64 `json:"levels"`
}

type Pyroscope struct {
	Version     int         `json:"version"`
	Flamebearer Flamebearer `json:"flamebearer"`
}
