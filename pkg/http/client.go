package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type HttpClient struct {
	Client *http.Client
}

func NewClient() *HttpClient {
	return &HttpClient{}
}

func (c *HttpClient) Set(i interface{}) {
	c.Client = parseHttpClientFromInterface(i).Client
}

func parseHttpClientFromInterface(i interface{}) *HttpClient {
	client := i.(*HttpClient)

	return client
}

func (c *HttpClient) RequestJenkinsCrumb() {
	time.Sleep(time.Minute * 3)
	req, err := http.NewRequest(http.MethodGet, "http://home-idp-ci-jenkins:8080/crumbIssuer/api/json", nil)
	fmt.Println(err)
	fmt.Println(err)

	req.SetBasicAuth("admin", "tester1234")

	resp, err := c.Client.Do(req)
	body2, err := io.ReadAll(resp.Body)
	m := make(map[string]interface{}, 10)
	json.Unmarshal(body2, &m)
	fmt.Println("CRUMB:", m["crumb"])

	data := "newTokenName=testtoken"
	req2, err := http.NewRequest("POST", "http://home-idp-ci-jenkins:8080/user/admin/descriptorByName/jenkins.security.ApiTokenProperty/generateNewToken", strings.NewReader(data))
	fmt.Println(err)
	fmt.Println(err)
	req2.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req2.Header.Set("Jenkins-Crumb", m["crumb"].(string))
	req2.SetBasicAuth("admin", "tester1234")
	resp2, err := c.Client.Do(req)
	fmt.Println(err)
	fmt.Println(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp2.Body)
	fmt.Println(err)
	fmt.Println(err)

	fmt.Println("TOKEN:", string(body))
}

// curl -u admin:tester1234  -cookie-jar -s  http://home-idp-ci-jenkins:8080/crumbIssuer/api/json/xml?xpath=concat\(//crumbRequestField,%22:%22,//crumb\)

// curl -u "admin:tester1234" -X POST "http://home-idp-ci-jenkins:8080/user/test/descriptorByName/jenkins.security.ApiTokenProperty/generateNewToken" -H "Jenkins-Crumb:96c7c82103e4b9ab36addd21ab372863fd075411b77ea492e2c25dcf0b02c77d" -d "newTokenName=test-token" -H "Content-Type: application/x-www-form-urlencoded"
