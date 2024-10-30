package model

type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	RoleID     string `json:"role_id"`
	CreateTime string `json:"create_time"`
}

type Project struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Creator string `json:"creator"`
}

type Policy struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Json string `json:"json"`
}

type Trace struct {
	SpanID       string `json:"span_id"`
	TraceID      string `json:"trace_id"`
	ParentSpanID string `json:"parent_span_id"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	Status       string `json:"status"`
}
