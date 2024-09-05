package http

import (
	"encoding/json"
	"fmt"

	"github.com/choigonyok/home-idp/pkg/http"
)

type GatewayHttpClient struct {
	Client *http.HttpClient
}

func (c *GatewayHttpClient) Set(i interface{}) {
	c.Client = parseHttpClientFromInterface(i)
}

func parseHttpClientFromInterface(i interface{}) *http.HttpClient {
	client := i.(*http.HttpClient)
	return client
}

func (c *GatewayHttpClient) SendArgoCDWebhook(payload []byte, headers map[string]string) error {
	tmp := make(map[string]interface{})
	json.Unmarshal(payload, &tmp)

	fmt.Println("TEST TMP:", tmp)
	req := http.NewRequest(http.Post, "http://home-idp-cd-argocd-server:80/api/webhook", tmp)

	for k, v := range headers {
		req.SetHeader(k, v)
	}

	resp, err := c.Client.Request(req)
	fmt.Println("TEST RESPONSE STATUS CODE OF MANIFEST ARGOCD WEBHOOK REQUEST:", resp.StatusCode)
	return err
}
