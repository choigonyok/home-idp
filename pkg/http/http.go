package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type HttpClient struct {
	Client *http.Client
}

type Method string

const (
	Post   Method = http.MethodPost
	Get    Method = http.MethodGet
	Delete Method = http.MethodDelete
	Put    Method = http.MethodPut

	StatusOK int = http.StatusOK
)

type Header struct {
	Key   string
	Value string
}

type Request struct {
	username string
	password string
	URL      string
	Method   Method
	Body     map[string]interface{}
	headers  map[string]string
}

func NewClient() *HttpClient {
	return &HttpClient{}
}

func NewRequest(method Method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(string(method), url, body)
}

func (c *HttpClient) Request(r *Request, headers ...Header) (*http.Response, error) {
	var b *bytes.Buffer = nil

	if r.Body != nil {
		jsonData, err := json.Marshal(r.Body)
		if err != nil {
			return nil, err
		}

		b = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(string(r.Method), r.URL, b)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	if r.username != "" || r.password != "" {
		req.SetBasicAuth(r.username, r.password)
	}

	for _, h := range headers {
		req.Header.Set(h.Key, h.Value)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *Request) SetHeader(key, value string) {
	r.headers[key] = value
}

func (r *Request) SetBasicAuth(username, password string) {
	r.username = username
	r.password = password
}
