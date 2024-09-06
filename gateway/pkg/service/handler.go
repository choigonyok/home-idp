package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	deployPb "github.com/choigonyok/home-idp/deploy-manager/pkg/proto"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/git"
	"github.com/choigonyok/home-idp/pkg/util"
	"github.com/google/go-github/v63/github"
)

// func (svc *Gateway) InstallArgoCDHandler() http.HandlerFunc {
// 	return func(resp http.ResponseWriter, req *http.Request) {
// 		data := &helm.ArgoCDData{}

// 		if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
// 			http.Error(resp, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		fmt.Println(data)

// 		ok, err := grpc.InstallArgoCDChart(data, svc.ClientSet.GrpcClient[util.InstallManager].GetConnection())
// 		fmt.Println(ok)
// 		fmt.Println(err)
// 	}
// }

func (svc *Gateway) UninstallArgoCDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Body)
		fmt.Println("TESTHANDLER1")
	}
}

func (svc *Gateway) TestHandler2() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Body)
		fmt.Println("TESTHANDLER2")
	}
}

// /api/webhook/harbor?username="choigonyok"&email="choigonyok@naver.com"
func (svc *Gateway) HarborWebhookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Body)
		b, err := io.ReadAll(r.Body)
		fmt.Println("TESTHANDLER3:", err)
		m := make(map[string]interface{})
		err = json.Unmarshal(b, &m)

		img := util.ParseInterfaceMap(m, []string{"event_data", "resources", "resource_url"}).(string)
		s := strings.Split(img, "/")
		name, version, _ := strings.Cut(s[2], ":")

		fmt.Println(err)
		fmt.Println("MAPPED MAP:", name, version)

		if err := svc.ClientSet.GitClient.UpdateManifest(name + ":" + version); err != nil {
			fmt.Println("TEST UPDATE IMAGE FROM MANIFEST ERR:", err)
			return
		}
		return
	}
}

func (svc *Gateway) GithubWebhookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := github.ValidatePayload(r, nil)
		if err != nil {
			fmt.Println("TEST VALIDATE PAYLOAD ERR:", err)
		}
		event, err := github.ParseWebHook(github.WebHookType(r), payload)
		if err != nil {
			fmt.Println("TEST PARSE WEBHOOK ERR:", err)
		}
		switch event := event.(type) {
		case *github.PushEvent:
			t := getFileType(event)
			fmt.Println("TEST WEBHOOK TYPE:", t)
			switch t {
			case "docker":
				name, version := getImageNameAndVersionFromCommit(event)
				username := getUserFromCommit(event)
				fmt.Println("TEST RELATED USERNAME: ", username)
				fmt.Println("TEST WEBHOOK IMGNAME:", name)
				fmt.Println("TEST WEBHOOK IMGVERSION:", version)
				svc.requestBuildDockerfile(name, version, username)
			case "cd":
				path := getFilepathFromCommit(event)
				svc.requestDeploy(path)
				// case "manifest":
				// 	path := forwardToArgoCD(event)
				// 	svc.requestDeploy(path)
			case "manifest":
				svc.requestArgoCDWebhook(r, payload)
			}
		case *github.RepositoryEvent:
			fmt.Println("TEST PING WEBHOOK RECEIVED")
		}
	}
}

func getFileType(e *github.PushEvent) string {
	pushPath := e.Commits[0].Added[0]
	t, _, _ := strings.Cut(pushPath, "/")
	return t
}

func getImageNameAndVersionFromCommit(e *github.PushEvent) (string, string) {
	pushPath := e.Commits[0].Added[0]
	fmt.Println("TEST PUSH PATH:", pushPath)
	re := regexp.MustCompile(`^docker/[^/]+/Dockerfile.`)
	img := re.ReplaceAllString(pushPath, "")
	fmt.Println("TEST IMG:", img)
	name, version, _ := strings.Cut(img, ":")
	return name, version
}

func getUserFromCommit(e *github.PushEvent) string {
	pushPath := e.Commits[0].Added[0]
	fmt.Println("TEST PUSH PATH:", pushPath)

	_, pathWithoutType, _ := strings.Cut(pushPath, "/")

	username, _, _ := strings.Cut(pathWithoutType, "/")
	return username
}

func getFilepathFromCommit(e *github.PushEvent) string {
	return e.Commits[0].Added[0]
}

func (svc *Gateway) requestDeploy(filepath string) {
	c := deployPb.NewDeployClient(svc.ClientSet.GrpcClient[util.Components(util.DeployManager)].GetConnection())
	reply, err := c.Deploy(context.TODO(), &deployPb.DeployRequest{Filepath: filepath})
	if err != nil {
		fmt.Println("TEST DEPLOY REQUEST ERR:", err)
		return
	}

	if reply.Succeed {
		fmt.Println("TEST DEPLOY REQUEST FAILED")
		return
	}
}

func (svc *Gateway) requestBuildDockerfile(name, version, username string) {
	c := deployPb.NewBuildClient(svc.ClientSet.GrpcClient[util.Components(util.DeployManager)].GetConnection())
	reply, err := c.BuildDockerfile(context.TODO(), &deployPb.BuildDockerfileRequest{
		Img: &deployPb.Image{
			Pusher:  username,
			Name:    name,
			Version: version,
		},
	})
	if err != nil {
		fmt.Println("TEST BUILD DOCKERFILE REQUEST ERR:", err)
		return
	}

	if reply.Succeed {
		fmt.Println("TEST BUILD DOCKERFILE REQUEST FAILED")
		return
	}
}

func (svc *Gateway) requestArgoCDWebhook(r *http.Request, payload []byte) error {
	m := make(map[string]string)
	for k, v := range r.Header {
		m[k] = strings.Join(v, ", ")
	}

	fmt.Println("TEST GIT MANIFEST PUSH HEADERS:", m)

	fmt.Println("TEST GIT MANIFEST PUSH PAYLOAD:", string(payload))
	if err := svc.ClientSet.HttpClient.SendArgoCDWebhook(payload, m); err != nil {
		fmt.Println("TEST REQUEST MANIFEST ARGOCD WEBHHOOK ERR:", err)
		return err
	}

	return nil
}

func (svc *Gateway) ApiPostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println()
		fmt.Println("TEST REQUEST PATH:")
		fmt.Println()
		leadPath, _ := strings.CutPrefix(r.URL.Path, "/api/")
		fmt.Println("TEST LEAD PATH:", leadPath)
		dir, _, _ := strings.Cut(leadPath, "/")

		switch dir {
		case "dockerfile":
			svc.apiPostDockerfileHandler(w, r)
		case "manifest":
			svc.apiPostManifestHandler(w, r)
		}
	}
}

func (svc *Gateway) apiPostDockerfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET POST DOCKER REQUEST")
	fmt.Println()
	fmt.Println("TEST REQEUST BODY:", r.Body)
	fmt.Println()

	b, _ := io.ReadAll(r.Body)

	f := git.GitDockerFile{}
	json.Unmarshal(b, &f)

	imageName, _, _ := strings.Cut(f.Image, ":")
	if svc.ClientSet.GitClient.IsDockerfileExist(f.Username, imageName) {
		fmt.Println("TEST DOCKERFILE ALREADY EXIST")
		svc.ClientSet.GitClient.UpdateDockerFile(f.Username, f.Image, f.Content)
		return
	}
	fmt.Println("TEST DOCKERFILE NOT EXIST")
	svc.ClientSet.GitClient.CreateDockerFile(f.Username, f.Image, f.Content)
	return
}

// /api/manifest?username="choigonyok"
func (svc *Gateway) apiPostManifestHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	err := svc.ClientSet.GitClient.CreatePodManifestFile(username, env.Get("HOME_IDP_GIT_EMAIL"), "test:v1.0", 8080)
	fmt.Println(err)
}

// func (svc *Gateway) applyArgoCDApplication(e *github.PushEvent) {
// 	pushPath := e.Commits[0].Added[0]
// 	svc.ClientSet.GitClient.GetArgoCDApplication()
// 	// cd/username/main.yaml
// 	svc.ClientSet.KubeClient.ApplyManifest()
// }

// 다 Running 상태인지,
// webhook이 harbor, github 잘 생성되었는지

// curl -u "admin:tester1234" -X GET http://home-idp-harbor-core.idp-system.svc.cluster.local:80/api/v2.0/projects/library/webhook/policies

// git clone 후에
// 도커파일 /docker/Dockerfile.testimg:v1.123 을 푸시했을 때,
// 정상적으로 kaniko job이 생성되고 harbor에 푸시되는지

// curl -X GET "http://home-idp-harbor-core.idp-system.svc.cluster.local:80/api/v2.0/projects/library/repositories" -H "accept: application/json"

// curl -X GET "http://home-idp-harbor-core.idp-system.svc.cluster.local:80/api/v2.0/projects/library/repositories/testimg77/artifacts" -H "accept: application/json"

// harbor에서 푸시된 이미지에 대한 웹훅이 gateway 로그에 잘 출력되는지

// curl -X POST https://cblicense.front.slexn.com/api/dockerfile \
//      -H "Content-Type: application/json" \
//      -d '{
//            "username": "user123",
//            "tag": "latest",
//            "content": "FROM ubuntu:18.04\nRUN apt-get update && apt-get install -y git"
//          }'

// curl -X POST http://cd.choigonyok.com:8080/api/v1/applications/app-choigonyok/refresh
