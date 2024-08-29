package http

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/http"
	"github.com/choigonyok/home-idp/pkg/util"
)

type InstallManagerHttpClient struct {
	Client *http.HttpClient
}

func (c *InstallManagerHttpClient) Set(i interface{}) {
	c.Client = parseHttpClientFromInterface(i)
}

func parseHttpClientFromInterface(i interface{}) *http.HttpClient {
	client := i.(*http.HttpClient)
	return client
}

func (c *InstallManagerHttpClient) CreateHarborWebhook() error {
	data := map[string]interface{}{
		"name": "HARBOR_WEBHOOK",
		"targets": []map[string]interface{}{
			{
				"type":             "http",
				"address":          "http://" + env.Get("HOME_IDP_HOST") + "/webhooks/harbor",
				"skip_cert_verify": true,
			},
		},
		"event_types": []string{
			"PUSH_ARTIFACT",
		},
		"enabled": true,
	}

	r := http.NewRequest(http.Post, "http://harbor."+env.Get("HOME_IDP_NAMESPACE")+".svc.cluster.local:80/api/v2.0/projects/library/webhook/policies", data)
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

func (c *InstallManagerHttpClient) IsHarborHealthy() (bool, error) {
	r := http.NewRequest(http.Get, "http://harbor."+env.Get("HOME_IDP_NAMESPACE")+".svc.cluster.local:80/api/v2.0/health", nil)
	r.SetBasicAuth("admin", env.Get("HOME_IDP_ADMIN_PASSWORD"))

	resp, err := c.Client.Request(r)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		fmt.Println(http.StatusOK)
		return false, nil
	}

	body, err := io.ReadAll(resp.Body)

	fmt.Println("TEST HARBOR HEALTH RESPONSE BODY:", string(body))

	m := make(map[string]interface{})

	json.Unmarshal(body, &m)

	fmt.Println("TEST HARBOR STATUS JSON MAP:", m)
	if util.ParseInterfaceMap(m, []string{"status"}).(string) == "healthy" {
		return true, nil
	}

	return false, nil
}

func (c *InstallManagerHttpClient) CreateArgoCDSessionToken(password string) (string, error) {
	data := map[string]interface{}{
		"username": "admin",
		"password": password,
	}

	r := http.NewRequest(http.Post, "http://home-idp-cd-argocd-server."+env.Get("HOME_IDP_NAMESPACE")+".svc.cluster.local:80/api/v1/session", data)

	r.SetHeader("Content-Type", "application/json")
	r.SetBasicAuth("admin", env.Get("HOME_IDP_ADMIN_PASSWORD"))

	resp, err := c.Client.Request(r)
	if err != nil {
		fmt.Println("TEST CREATE ARGOCD SESSION REQUEST ERR:", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	m := make(map[string]interface{})

	json.Unmarshal(body, &m)

	if err != nil {
		fmt.Println("TEST READ ARGOCD SESSION ERR:", err)
		return "", err
	}
	fmt.Println("TEST ARGOCD SESSION CREATE RESPONSE:", string(body))

	fmt.Println("TEST TOKEN RESPONSE MAP:", m)

	return m["token"].(string), nil
}

func (c *InstallManagerHttpClient) CreateArgoCDRepository(password, token string) error {
	data := map[string]interface{}{
		"name":               "home-idp",
		"repo":               "https://github.com/" + env.Get("HOME_IDP_GIT_USERNAME") + "/" + env.Get("HOME_IDP_GIT_REPO") + ".git",
		"username":           env.Get("HOME_IDP_GIT_USERNAME"),
		"password":           env.Get("HOME_IDP_GIT_TOKEN"),
		"type":               "git",
		"project":            "default",
		"forceHttpBasicAuth": true,
		"sshPrivateKey":      "",
		"tlsClientCertData":  "",
		"tlsClientCertKey":   "",
		"insecure":           true,
		"enableLfs":          false,
	}

	r := http.NewRequest(http.Post, "http://home-idp-cd-argocd-server."+env.Get("HOME_IDP_NAMESPACE")+".svc.cluster.local:80/api/v1/repositories", data)
	r.SetHeader("Content-Type", "application/json")
	r.SetHeader("Authorization", "Bearer "+token)
	r.SetBasicAuth("admin", password)

	resp, err := c.Client.Request(r)
	if err != nil {
		fmt.Println("TEST CREATE ARGOCD REPOSITORY REQUEST ERR:", err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("TEST READ ARGOCD REPOSITORY ERR:", err)
		return err
	}
	fmt.Println("TEST ARGOCD REPOSITORY CREATE RESPONSE:", string(body))

	return nil
}
