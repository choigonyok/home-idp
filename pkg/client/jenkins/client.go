package client

import (
	"encoding/base64"
	"net/http"
)

type Jenkins struct {
	Client *http.Client
	Host   string
	Token  string
}

func NewJenkinsFromAPIToken(jenkinsHost string, userID string, token string) *Jenkins {
	t := base64.RawStdEncoding.EncodeToString([]byte(userID + ":" + token))

	cli := http.DefaultClient
	return &Jenkins{
		Host:   jenkinsHost,
		Token:  t,
		Client: cli,
	}
}

func (j *Jenkins) List(kind string) (*Pipeline, error) {
	switch kind {
	case Job:
		return listJobs(j)
	}
	return nil, nil
}

func (j *Jenkins) Run(kind string, resource string, query map[string][]string) (*http.Response, error) {
	switch kind {
	case "Job":
		return runJobs(j, resource, query)
	}
	return nil, nil
}
