package trace

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/choigonyok/idp/pkg/observability/trace/model"
	"github.com/go-resty/resty/v2"
)

type Trace struct {
	resty         *resty.Client
	host          string
	pyroscopeHost string
}

func NewTrace(resty *resty.Client, host string) *Trace {
	return &Trace{
		resty: resty,
		host:  host,
	}
}

func (t *Trace) GetTempoTraces(service string) (*model.Traces, error) {
	now := time.Now().Unix()
	now15mAgo := time.Now().Add(-15 * time.Minute).Unix()

	resp, err := t.resty.R().
		SetHeader("Content-Type", "application/json").
		Get(t.host + "/api/search?start=" + strconv.Itoa(int(now15mAgo)) + "&end=" + strconv.Itoa(int(now)) + "&limit=5&sort=startTime&direction=desc&service=" + service)

	if err != nil {
		return nil, err
	}

	traces := model.Traces{}

	json.Unmarshal(resp.Body(), &traces)

	for _, trace := range traces.Traces {
		spans := model.Spans{}
		resp, err := t.resty.R().
			SetHeader("Content-Type", "application/json").
			Get(t.host + "/api/traces/" + trace.TraceID)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(resp.Body(), &spans)
		trace.Spans = spans

		for _, b := range trace.Spans.Batches {
			for _, s := range b.ScopeSpans {
				for _, ss := range s.Spans {
					s1, err := base64.StdEncoding.DecodeString(ss.SpanID)
					if err != nil {
						return nil, err
					}
					s2, err := base64.StdEncoding.DecodeString(ss.TraceID)
					if err != nil {
						return nil, err
					}

					b1 := hex.EncodeToString(s1)
					b2 := hex.EncodeToString(s2)
					ss.SpanID = b1
					ss.TraceID = b2

					resp, err := t.resty.R().
						SetHeader("Content-Type", "application/json").
						SetQueryParam("query", `process_cpu:cpu:nanoseconds:cpu:nanoseconds{service_name="tempo.trace.spans", trace_id="`+ss.TraceID+`"}`).
						SetQueryParam("from", "now-1h").
						Get(t.pyroscopeHost + "/pyroscope/render")
					if err != nil {
						return nil, err
					}

					treeNode := getTree(resp.Body())

					trace.Profile = treeNode
				}
			}
		}
	}
	return &traces, nil
}

func getTree(b []byte) *model.TreeNode {
	var data model.Pyroscope
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

func BuildFlameTree(levels [][]float64, names []string) *model.TreeNode {
	if len(levels) == 0 {
		return nil
	}

	// 1) 각 레벨을 절대 좌표(start/end)로 디코딩
	decoded := make([][]model.DecodedBlock, 0, len(levels))
	for _, lvl := range levels {
		cursor := int64(0)
		row := make([]model.DecodedBlock, 0, len(lvl)/4)

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

			node := &model.TreeNode{
				Name:     name,
				Value:    total,
				SelfTime: self,
				Start:    start,
				End:      end,
			}
			row = append(row, model.DecodedBlock{
				Start: start,
				End:   end,
				Node:  node,
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
			for j < len(children) && children[j].End <= p.Start {
				j++
			}
			// 부모 범위에 들어오는 애들만 붙임
			for j < len(children) && children[j].Start >= p.Start && children[j].End <= p.End {
				p.Node.Children = append(p.Node.Children, children[j].Node)
				j++
			}
		}
	}

	// 3) 루트 반환 (일반적으로 level 0 첫 블록이 total)
	if len(decoded[0]) == 0 {
		return nil
	}
	return decoded[0][0].Node
}
