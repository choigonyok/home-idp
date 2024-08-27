package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/util"
)

type InstallManagerHttpClient struct {
	Client *http.Client
}

func (cli *InstallManagerHttpClient) Set(i interface{}) {
	cli.Client = parseHttpClientFromInterface(i)
}

func parseHttpClientFromInterface(i interface{}) *http.Client {
	client := i.(*http.Client)
	return client
}

func (cli *InstallManagerHttpClient) CreateHarborWebhook() error {
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

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("TEST MARSHAL HARBOR WEBHOOK ERR:", err)
		return err
	}

	b := bytes.NewBuffer(jsonData)

	req, err := http.NewRequest(http.MethodPost, "http://harbor."+env.Get("HOME_IDP_NAMESPACE")+".svc.cluster.local:80/api/v2.0/projects/library/webhook/policies", b)
	if err != nil {
		fmt.Println("TEST CREATE HARBOR WEBHOOK REQUEST ERR:", err)
		return err
	}
	defer req.Body.Close()

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("admin", env.Get("HOME_IDP_ADMIN_PASSWORD"))

	resp, err := cli.Client.Do(req)
	if err != nil {
		fmt.Println("TEST REQUEST HARBOR WEBHOOK ERR:", err)
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

func (cli *InstallManagerHttpClient) IsHarborHealthy() (bool, error) {
	req, err := http.NewRequest(http.MethodGet, "http://harbor."+env.Get("HOME_IDP_NAMESPACE")+".svc.cluster.local:80/api/v2.0/health", nil)
	if err != nil {
		return false, err
	}

	req.SetBasicAuth("admin", env.Get("HOME_IDP_ADMIN_PASSWORD"))

	resp, err := cli.Client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	fmt.Println("TEST HARBOR HEALTH RESPONSE STATUS CODE:", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
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

func (cli *InstallManagerHttpClient) CreateArgoCDRepository(password string) error {
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

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("TEST MARSHAL ARGOCD REPOSITORY ERR:", err)
		return err
	}

	b := bytes.NewBuffer(jsonData)

	req, err := http.NewRequest(http.MethodPost, "http://argocd-server."+env.Get("HOME_IDP_NAMESPACE")+".svc.cluster.local:80/api/v1/repositories", b)
	if err != nil {
		fmt.Println("TEST CREATE ARGOCD REPOSITORY REQUEST ERR:", err)
		return err
	}
	defer req.Body.Close()

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("admin", password)

	resp, err := cli.Client.Do(req)
	if err != nil {
		fmt.Println("TEST REQUEST ARGOCD REPOSITORY ERR:", err)
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
