package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpClient struct {
	*http.Client
}

const (
	Post   string = http.MethodPost
	Get    string = http.MethodGet
	Delete string = http.MethodDelete
	Put    string = http.MethodPut

	StatusOK int = http.StatusOK
)

type Header struct {
	Key   string
	Value string
}

type Request struct {
	username string
	password string
	url      string
	method   string
	body     map[string]interface{}
	headers  map[string]string
}

func NewClient() *HttpClient {
	return &HttpClient{
		http.DefaultClient,
	}
}

func NewRequest(method, url string, body map[string]interface{}) *Request {
	return &Request{
		method:  method,
		url:     url,
		body:    body,
		headers: make(map[string]string),
	}
}

func (c *HttpClient) Request(r *Request) (*http.Response, error) {
	jsonData, err := json.Marshal(r.body)
	if err != nil {
		return nil, err
	}
	b := bytes.NewBuffer(jsonData)

	req := &http.Request{}
	if r.method == Get {
		req, err = http.NewRequest(r.method, r.url, nil)
		fmt.Println(err)
	} else {
		req, err = http.NewRequest(r.method, r.url, b)
		fmt.Println(err)
	}

	if r.username != "" || r.password != "" {
		req.SetBasicAuth(r.username, r.password)
	}

	for k, v := range r.headers {
		req.Header.Set(k, v)
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
