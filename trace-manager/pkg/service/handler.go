package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/choigonyok/home-idp/pkg/model"
	"github.com/gorilla/mux"
)

func (svc *TraceManager) apiPostSpanHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		b, _ := io.ReadAll(r.Body)

		data := model.Trace{}
		json.Unmarshal(b, &data)

		svc.ClientSet.StorageClient.DB().Exec(`INSERT INTO spans (span_id, trace_id, parent_span_id, start_time, status) VALUES ('` + data.SpanID + `', '` + data.TraceID + `', '` + data.ParentSpanID + `', '` + data.StartTime + `', '` + data.Status + `')`)
	}
}

func (svc *TraceManager) apiPutSpanHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		b, _ := io.ReadAll(r.Body)
		data := model.Trace{}
		json.Unmarshal(b, &data)

		svc.ClientSet.StorageClient.DB().Exec(`UPDATE spans SET end_time = '` + data.EndTime + `', status = '` + data.Status + `' WHERE span_id = '` + data.SpanID + `'`)
	}
}

func (svc *TraceManager) apiGetTraceHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		traceId := vars["traceId"]

		rows, err := svc.ClientSet.StorageClient.DB().Query(`SELECT span_id, parent_span_id, start_time, end_time, status FROM spans WHERE trace_id = '` + traceId + `' ORDER BY create_time ASC`)
		if err != nil {
			fmt.Println("ERR GETTING TRACE:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		spans := []*model.Trace{}

		for rows.Next() {
			span := model.Trace{}
			rows.Scan(&span.SpanID, &span.ParentSpanID, &span.StartTime, &span.EndTime, &span.Status)
			span.TraceID = traceId
			spans = append(spans, &span)
		}

		b, _ := json.Marshal(spans)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
