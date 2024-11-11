package http

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/choigonyok/home-idp/pkg/env"
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

func (c *GatewayHttpClient) CreateHarborProject(name string) error {
	data := map[string]interface{}{
		"project_name": name,
		"public":       true,
	}

	r := http.NewRequest(http.Post, "http://"+env.Get("HOME_IDP_HARBOR_HOST")+":8080/api/v2.0/projects", data)
	r.SetBasicAuth("admin", env.Get("HOME_IDP_ADMIN_PASSWORD"))
	r.SetHeader("Content-Type", "application/json")

	resp, err := c.Client.Request(r)
	if err != nil {
		fmt.Println("TEST HTTP REQUEST ERR:", err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("TEST READ HARBOR WEBHOOK ERR:", err)
		return err
	}
	fmt.Println("TEST HARBOR WEBHOOK CREATE RESPONSE:", string(body))

	return nil
}

func (c *GatewayHttpClient) CreateHarborWebhook(project string) error {
	host := env.Get("HOME_IDP_API_HOST") + ":" + env.Get("HOME_IDP_API_PORT")
	data := map[string]interface{}{
		"name": "HARBOR_WEBHOOK",
		"targets": []map[string]interface{}{
			{
				"type":             "http",
				"address":          "http://" + host + "/webhooks/harbor",
				"skip_cert_verify": true,
			},
		},
		"event_types": []string{
			"PUSH_ARTIFACT",
		},
		"enabled": true,
	}

	r := http.NewRequest(http.Post, "http://"+env.Get("HOME_IDP_HARBOR_HOST")+":8080/api/v2.0/projects/"+project+"/webhook/policies", data)
	r.SetBasicAuth("admin", env.Get("HOME_IDP_ADMIN_PASSWORD"))
	r.SetHeader("Content-Type", "application/json")

	resp, err := c.Client.Request(r)
	if err != nil {
		fmt.Println("TEST HTTP REQUEST ERR:", err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("TEST READ HARBOR WEBHOOK ERR:", err)
		return err
	}
	fmt.Println("TEST HARBOR WEBHOOK CREATE RESPONSE:", string(body))

	return nil
}
