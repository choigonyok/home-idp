package client

import (
	"encoding/json"
	"io"
	"net/http"
)

func listJenkinsViews(j *Jenkins, r *JenkinsResp) error {
	req, _ := http.NewRequest(http.MethodGet, j.Host+"/api/json", nil)
	q := req.URL.Query()
	q.Add("tree", "views[name,url]")
	req.URL.RawQuery = q.Encode()
	if j.Token != "" {
		req.Header.Add("Authorization", "Basic "+j.Token)
	}

	resp, err := j.Client.Do(req)
	if err != nil {
		return err
	}

	b, _ := io.ReadAll(resp.Body)
	json.Unmarshal(b, r)

	return nil
}

func listJobsInView(j *Jenkins, viewName string) (*JenkinsResp, error) {
	req, _ := http.NewRequest(http.MethodGet, j.Host+"/view/"+viewName+"/api/json", nil)
	q := req.URL.Query()
	q.Add("tree", "jobs[name,color]")
	req.URL.RawQuery = q.Encode()
	if j.Token != "" {
		req.Header.Add("Authorization", "Basic "+j.Token)
	}

	resp, err := j.Client.Do(req)
	if err != nil {
		return nil, err
	}

	b, _ := io.ReadAll(resp.Body)
	r := JenkinsResp{}
	json.Unmarshal(b, &r)

	return &r, nil
}
