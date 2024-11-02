package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	deployPb "github.com/choigonyok/home-idp/deploy-manager/pkg/proto"
	"github.com/choigonyok/home-idp/gateway/pkg/progress"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/model"
	"github.com/choigonyok/home-idp/pkg/trace"
	"github.com/choigonyok/home-idp/pkg/util"
	rbacPb "github.com/choigonyok/home-idp/rbac-manager/pkg/proto"
	"github.com/google/go-github/v63/github"
	"github.com/gorilla/mux"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/yaml.v2"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
)

var testUserId = "37e54287-af53-42a1-80a6-ac95361d3005"
var testUserName = "choigonyok"
var testUserEmail = "achoistic98@naver.com"

var Spans = make(map[string]*trace.Span1)
var RootSpans = make(map[string]*trace.Span1)

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
		json.Unmarshal(b, &m)

		img := util.ParseInterfaceMap(m, []string{"event_data", "resources", "resource_url"}).(string)
		s := strings.Split(img, "/")
		name, version, _ := strings.Cut(s[2], ":")
		c := rbacPb.NewRbacServiceClient(svc.ClientSet.RbacGrpcClient.GetConnection())
		resp, err := c.GetTraceId(context.TODO(), &rbacPb.GetTraceIdRequest{
			ImageName:    name,
			ImageVersion: version,
		})
		if err != nil {
			fmt.Println("ERR GETTING TRACE ID :", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ImageSpan := Spans[resp.GetTraceId()]
		err = ImageSpan.Stop()
		if err != nil {
			fmt.Println("BUILD SPAN STOP ERR:", err)
		}
		rootSpan := RootSpans[resp.GetTraceId()]
		deploySpan := svc.ClientSet.TraceClient.NewSpanFromOutgoingContext(rootSpan.Context)
		err = deploySpan.Start(rootSpan.Context)
		if err != nil {
			fmt.Println("DEPLOY SPAN START ERR:", err)
		}
		manifestSpan := svc.ClientSet.TraceClient.NewSpanFromOutgoingContext(deploySpan.Context)
		err = manifestSpan.Start(deploySpan.Context)
		if err != nil {
			fmt.Println("UPDATE MANIFEST SPAN START ERR:", err)
		}

		if err := svc.ClientSet.GitClient.CreatePodManifestFile(testUserName, testUserEmail, name, version, 80); err != nil {
			fmt.Println("TEST UPDATE IMAGE FROM MANIFEST ERR:", err)
			return
		}
		// if err := svc.ClientSet.GitClient.UpdateManifest(name + ":" + version); err != nil {
		// 	fmt.Println("TEST UPDATE IMAGE FROM MANIFEST ERR:", err)
		// 	return
		// }
		err = manifestSpan.Stop()
		if err != nil {
			fmt.Println("UPDATE MANIFEST SPAN STOP ERR:", err)
		}

		Spans[deploySpan.TraceID] = deploySpan

		err = rootSpan.Stop()
		if err != nil {
			fmt.Println("ROOT SPAN STOP ERR:", err)
		}

	}
}

func (svc *Gateway) LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		b, _ := io.ReadAll(r.Body)

		tmp := struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{}

		json.Unmarshal(b, &tmp)

		// svc.requestLogin(tmp.Username, tmp.Password)

		fmt.Println(tmp)
	}
}

func (svc *Gateway) LoginCallbackHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		schema := "http"
		if env.Get("HOME_IDP_API_TLS_ENABLED") == "true" {
			schema = "https"
		}
		host := env.Get("HOME_IDP_API_HOST")
		port := env.Get("HOME_IDP_API_PORT")

		tokenURL := "https://github.com/login/oauth/access_token"
		data := url.Values{
			"client_id":     {env.Get("HOME_IDP_GIT_OAUTH_CLIENT_ID")},
			"client_secret": {env.Get("HOME_IDP_GIT_OAUTH_CLIENT_SECRET")},
			"code":          {code},
			"redirect_uri":  {fmt.Sprintf("%s://%s:%s/login/callback", schema, host, port)},
		}

		req, err := http.NewRequest("POST", tokenURL, nil)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("Accept", "application/json")
		req.URL.RawQuery = data.Encode()
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)

		var tokenResponse map[string]interface{}
		if err := json.Unmarshal(body, &tokenResponse); err != nil {
			log.Fatal(err)
		}

		accessToken := tokenResponse["access_token"].(string)

		userReq, _ := http.NewRequest("GET", "https://api.github.com/user", nil)
		userReq.Header.Set("Authorization", fmt.Sprintf("token %s", accessToken))

		userResp, err := client.Do(userReq)
		if err != nil {
			log.Fatal(err)
		}
		defer userResp.Body.Close()

		userBody, _ := io.ReadAll(userResp.Body)

		var userInfo map[string]interface{}
		json.Unmarshal(userBody, &userInfo)

		c := rbacPb.NewRbacServiceClient(svc.ClientSet.RbacGrpcClient.GetConnection())
		reply, err := c.Check(context.TODO(), &rbacPb.RbacRequest{
			Username: userInfo["login"].(string),
			Target:   "everything",
			Action:   rbacPb.Action_CREATE,
		})
		if err != nil {
			fmt.Println("TEST RBAC CHECK ERR: ", err)
		}

		switch reply.Result.String() {
		case "ASK":
			fmt.Println("ASK RETURN")
		case "ALLOW":
			fmt.Println("ALLOW RETURN")
		case "DENY":
			fmt.Println("DENY RETURN")
		case "ERROR":
			fmt.Println("ERROR RETURN")
		default:
			fmt.Println("DEFAULT RETURN")
		}
	}
}

func (svc *Gateway) SignUpHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		b, _ := io.ReadAll(r.Body)

		tmp := struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{}

		json.Unmarshal(b, &tmp)

		// svc.requestLogin(tmp.Username, tmp.Password)

		fmt.Println(tmp)
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
				c := rbacPb.NewRbacServiceClient(svc.ClientSet.RbacGrpcClient.GetConnection())
				resp, err := c.GetTraceId(context.TODO(), &rbacPb.GetTraceIdRequest{
					ImageName:    name,
					ImageVersion: version,
				})
				if err != nil {
					fmt.Println("ERR GETTING TRACE ID :", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				RootSpan := RootSpans[resp.GetTraceId()]
				repo := resp.GetRepository()

				ImageSpan := svc.ClientSet.TraceClient.NewSpanFromOutgoingContext(RootSpan.Context)

				err = ImageSpan.Start(RootSpan.Context)
				if err != nil {
					fmt.Println("IMAGE SPAN START ERR:", err)
				}
				username := getUserFromCommit(event)
				svc.requestBuildDockerfile(ImageSpan, name, version, username, repo)
				Spans[ImageSpan.TraceID] = ImageSpan

			case "cd":
				path := getFilepathFromCommit(event)

				svc.requestDeploy(path)
			// case "manifest":
			// 	path := forwardToArgoCD(event)
			// 	svc.requestDeploy(path)
			case "manifest":
				path := getFilepathFromCommit(event)
				file := svc.ClientSet.GitClient.Client.GetFilesByPath(path)
				tmp := struct {
					Spec struct {
						Containers []struct {
							Image string `yaml:"image"`
						} `yaml:"containers"`
					} `yaml:"spec"`
				}{}

				err := yaml.Unmarshal([]byte(file[0]), &tmp)
				if err != nil {
					fmt.Println("ERR UNMARSHALING MANIFEST TO POD : ", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				imageName := ""
				imageVersion := ""
				parts := strings.Split(tmp.Spec.Containers[0].Image, ":")
				if len(parts) == 2 {
					imageName = parts[0]
					imageVersion = parts[1]
					fmt.Println("imageName: ", imageName)
					fmt.Println("imageVersion: ", imageVersion)
				}

				c := rbacPb.NewRbacServiceClient(svc.ClientSet.RbacGrpcClient.GetConnection())
				resp, err := c.GetTraceId(context.TODO(), &rbacPb.GetTraceIdRequest{
					ImageName:    imageName,
					ImageVersion: imageVersion,
				})
				if err != nil {
					fmt.Println("ERR GETTING TRACE ID :", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				fmt.Println("TRACE ID MANIFEST:", resp.GetTraceId())
				deploySpan := Spans[resp.GetTraceId()]
				argocdWebhookSpan := svc.ClientSet.TraceClient.NewSpanFromOutgoingContext(deploySpan.Context)
				err = argocdWebhookSpan.Start(deploySpan.Context)
				if err != nil {
					fmt.Println("ARGOCD WEBHOOK SPAN START ERR:", err)
				}

				if fn := getFilename(event); fn != ".gitkeep" {
					svc.requestArgoCDWebhook(r, payload)
				}

				err = argocdWebhookSpan.Stop()
				if err != nil {
					fmt.Println("ARGOCD WEBHOOK SPAN STOP ERR:", err)
				}
				err = deploySpan.Stop()
				if err != nil {
					fmt.Println("DEPLOY SPAN STOP ERR:", err)
				}
			default:
				fmt.Println("TEST PING WEBHOOK RECEIVED")
			}
		}
	}

}

func getFileType(e *github.PushEvent) string {
	if len(e.Commits[0].Added) != 0 {
		pushPath := e.Commits[0].Added[0]
		t, _, _ := strings.Cut(pushPath, "/")
		return t
	} else if len(e.Commits[0].Removed) != 0 {
		pushPath := e.Commits[0].Removed[0]
		t, _, _ := strings.Cut(pushPath, "/")
		return t
	} else if len(e.Commits[0].Modified) != 0 {
		pushPath := e.Commits[0].Modified[0]
		t, _, _ := strings.Cut(pushPath, "/")
		return t
	} else {
		return ""
	}
}

func getFilename(e *github.PushEvent) string {
	if len(e.Commits[0].Added) != 0 {
		pushPath := e.Commits[0].Added[0]
		i := strings.LastIndex(pushPath, "/")
		return pushPath[i+1:]
	} else if len(e.Commits[0].Removed) != 0 {
		pushPath := e.Commits[0].Removed[0]
		i := strings.LastIndex(pushPath, "/")
		return pushPath[i+1:]
	} else if len(e.Commits[0].Modified) != 0 {
		pushPath := e.Commits[0].Modified[0]
		i := strings.LastIndex(pushPath, "/")
		return pushPath[i+1:]
	} else {
		return ""
	}
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
	c := deployPb.NewDeployClient(svc.ClientSet.DeployGrpcClient.GetConnection())
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

// func (svc *Gateway) requestLogin(username, password string) {
// 	c := rbacPb.NewLoginServiceClient(svc.ClientSet.GrpcClient.GetConnection())
// 	user := rbacPb.User{Username: username, Password: password}
// 	resp, err := c.Login(context.TODO(), &rbacPb.LoginRequest{User: &user})
// 	if err != nil {
// 		fmt.Println("TEST LOGIN GRPC REQUEST ERR:", err)
// 		return
// 	}

// 	if resp.Success {
// 		fmt.Println("TEST LOGIN GRPC REQUEST FAILED")
// 		return
// 	}
// }

func (svc *Gateway) requestBuildDockerfile(span *trace.Span1, name, version, username, repo string) {
	c := deployPb.NewBuildClient(svc.ClientSet.DeployGrpcClient.GetConnection())
	reply, err := c.BuildDockerfile(span.Context, &deployPb.BuildDockerfileRequest{
		Img: &deployPb.Image{
			Pusher:     username,
			Name:       name,
			Version:    version,
			Repository: repo,
		},
	})
	if err != nil {
		fmt.Println("TEST BUILD DOCKERFILE REQUEST ERR:", err)
		return
	}

	if !reply.Succeed {
		fmt.Println("TEST BUILD DOCKERFILE REQUEST FAILED")
		return
	}
}

func (svc *Gateway) requestArgoCDWebhook(r *http.Request, payload []byte) error {
	// step := progress.NewStep(progress.DeployResource, progress.Continue, nil)
	// step.Add(name + ":" + version)

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

func (svc *Gateway) GetProgressHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		image := vars["image"]
		w.Header().Set("Access-Control-Allow-Origin", "*")

		b, err := json.Marshal(progress.Map[image])
		if err != nil {
			fmt.Println("TEST MARSHALING ERR:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(b)
	}
}
func (svc *Gateway) ApiOptionsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uiSchema := "http"
		if env.Get("HOME_IDP_UI_TLS_ENABLED") == "true" {
			uiSchema = "https"
		}
		uiPort := env.Get("HOME_IDP_UI_PORT")
		uiHost := env.Get("HOME_IDP_UI_HOST")

		fmt.Println("TEST ORIGIN: " + uiSchema + "://" + uiHost + ":" + uiPort)

		w.Header().Set("Access-Control-Allow-Origin", "*")

		// 클라이언트가 사용할 수 있는 HTTP 메서드 목록
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		w.WriteHeader(http.StatusOK)
	}
}

func (svc *Gateway) apiPutUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userName := vars["userName"]
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		b, _ := io.ReadAll(r.Body)

		usr := model.User{}
		json.Unmarshal(b, &usr)

		c := rbacPb.NewRbacServiceClient(svc.ClientSet.RbacGrpcClient.GetConnection())
		_, err := c.PutUser(context.TODO(), &rbacPb.PutUserRequest{
			UserId: userName,
			User: &rbacPb.User{
				Name:   userName,
				RoleId: usr.RoleID,
			},
		})
		if err != nil {
			fmt.Println("PUT USER GRPC ERR:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (svc *Gateway) apiPostProjectHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		b, _ := io.ReadAll(r.Body)

		p := rbacPb.Project{}

		json.Unmarshal(b, &p)

		c := rbacPb.NewRbacServiceClient(svc.ClientSet.RbacGrpcClient.GetConnection())
		_, err := c.PostProject(context.TODO(), &rbacPb.PostProjectRequest{
			ProjectName: p.GetName(),
			CreatorId:   testUserId,
		})
		if err != nil {
			fmt.Println("ERR POSTING PROJECT :", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (svc *Gateway) apiPostRoleHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		b, _ := io.ReadAll(r.Body)

		role := rbacPb.Role{}

		json.Unmarshal(b, &role)

		c := rbacPb.NewRbacServiceClient(svc.ClientSet.RbacGrpcClient.GetConnection())
		_, err := c.PostRole(context.TODO(), &rbacPb.PostRoleRequest{
			RoleName: role.Name,
		})
		if err != nil {
			fmt.Println("ERR POSTING NEW ROLE :", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (svc *Gateway) apiPostUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		projectName := vars["projectName"]

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		b, _ := io.ReadAll(r.Body)

		usr := rbacPb.User{}

		json.Unmarshal(b, &usr)

		c := rbacPb.NewRbacServiceClient(svc.ClientSet.RbacGrpcClient.GetConnection())
		_, err := c.PostUser(context.TODO(), &rbacPb.PostUserRequest{
			User: &rbacPb.User{
				RoleId: usr.RoleId,
				Name:   usr.Name,
			},
			ProjectName: projectName,
		})
		if err != nil {
			fmt.Println("POST USER GRPC ERR:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (svc *Gateway) apiPostPodHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		b, _ := io.ReadAll(r.Body)
		pod := deployPb.Pod{}
		json.Unmarshal(b, &pod)

		c := deployPb.NewDeployClient(svc.ClientSet.DeployGrpcClient.GetConnection())
		_, err := c.DeployPod(context.TODO(), &deployPb.DeployPodRequest{
			Pod: &deployPb.Pod{
				Name:          pod.GetName(),
				Namespace:     pod.GetNamespace(),
				Image:         pod.GetImage(),
				ContainerPort: pod.GetContainerPort(),
			},
		})
		if err != nil {
			fmt.Println("ERR POSTTING NEW POD :", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (svc *Gateway) apiGetTraceHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		traceId := vars["traceId"]

		req, _ := http.NewRequest("GET", "http://home-idp-trace-manager:5103/api/traces/"+traceId, nil)
		resp, _ := http.DefaultClient.Do(req)
		bytes, _ := io.ReadAll(resp.Body)
		spans := []*model.Trace{}
		json.Unmarshal(bytes, &spans)

		datas := []struct {
			TraceID      string `json:"trace_id"`
			SpanID       string `json:"span_id"`
			Duration     string `json:"duration"`
			Status       string `json:"status"`
			StartTime    string `json:"start_time"`
			EndTime      string `json:"end_time"`
			ParentSpanID string `json:"parent_span_id"`
		}{}

		for _, span := range spans {
			st, _ := time.Parse("2006-01-02T15:04:05.999Z", span.StartTime)
			et, _ := time.Parse("2006-01-02T15:04:05.999Z", span.EndTime)
			d := et.Sub(st)
			datas = append(datas, struct {
				TraceID      string `json:"trace_id"`
				SpanID       string `json:"span_id"`
				Duration     string `json:"duration"`
				Status       string `json:"status"`
				StartTime    string `json:"start_time"`
				EndTime      string `json:"end_time"`
				ParentSpanID string `json:"parent_span_id"`
			}{
				TraceID:      span.TraceID,
				Status:       span.Status,
				SpanID:       span.SpanID,
				Duration:     d.String(),
				StartTime:    span.StartTime,
				EndTime:      span.EndTime,
				ParentSpanID: span.ParentSpanID,
			})
		}

		b, _ := json.Marshal(datas)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func (svc *Gateway) apiPostDockerfileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		conn := svc.ClientSet.RbacGrpcClient.GetConnection()

		b, _ := io.ReadAll(r.Body)
		// t := rbacPb.Dockerfile{}
		d := struct {
			Id           string `json:"id"`
			ImageName    string `json:"image_name"`
			ImageVersion string `json:"image_version"`
			CreatorId    string `json:"creator_id"`
			Repository   string `json:"repository"`
			Content      string `json:"content"`
			TraceId      string `json:"trace_id"`
		}{}

		// d := rbacPb.Dockerfile{}
		json.Unmarshal(b, &d)
		fmt.Println("PARSE TRACEID:", d.TraceId)
		RootSpan := svc.ClientSet.TraceClient.NewTrace(d.TraceId)

		fmt.Println("SPANID:", RootSpan.SpanID)
		fmt.Println("TRACEID:", RootSpan.TraceID)

		err := RootSpan.Start(context.Background())
		if err != nil {
			fmt.Println("ROOT SPAN START ERR:", err)
		}

		postDockerfileSpan := svc.ClientSet.TraceClient.NewSpanFromOutgoingContext(RootSpan.Context)
		err = postDockerfileSpan.Start(RootSpan.Context)
		if err != nil {
			fmt.Println("POST DOCKERFILE SPAN START ERR:", err)
		}

		// trace.Context, , grpc.Trailer(&trace.Trailer)
		c := rbacPb.NewRbacServiceClient(conn)
		_, err = c.PostDockerfile(postDockerfileSpan.Context, &rbacPb.PostDockerfileRequest{
			Dockerfile: &rbacPb.Dockerfile{
				ImageName:    d.ImageName,
				ImageVersion: d.ImageVersion,
				CreatorId:    testUserId,
				Repository:   d.Repository,
				Content:      d.Content,
				TraceId:      d.TraceId,
			},
		})
		if err != nil {
			fmt.Println("POST DOCKERFILE GRPC ERR:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		gitPushSpan := svc.ClientSet.TraceClient.NewSpanFromOutgoingContext(postDockerfileSpan.Context)
		err = gitPushSpan.Start(postDockerfileSpan.Context)
		if err != nil {
			fmt.Println("ERR GIT CONTEXT:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		RootSpans[RootSpan.TraceID] = RootSpan

		if svc.ClientSet.GitClient.IsDockerfileExist(d.ImageName) {
			svc.ClientSet.GitClient.UpdateDockerFile(testUserName, d.ImageName+":"+d.ImageVersion, d.Content)
			return
		} else {
			svc.ClientSet.GitClient.CreateDockerFile(testUserName, d.ImageName+":"+d.ImageVersion, d.Content)
		}

		err = gitPushSpan.Stop()
		if err != nil {
			fmt.Println("GIT PUSH SPAN STOP ERR:", err)
		}
		err = postDockerfileSpan.Stop()
		if err != nil {
			fmt.Println("POST DOCKERFILE SPAN STOP ERR:", err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func (svc *Gateway) apiGetDockerfilesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userName := vars["userName"]
		c := rbacPb.NewRbacServiceClient(svc.ClientSet.RbacGrpcClient.GetConnection())
		reply, err := c.GetDockerfiles(context.TODO(), &rbacPb.GetDockerfilesRequest{
			UserName: userName,
		})
		if err != nil {
			fmt.Println("ERR GETTING DOCKERFILES ERR:", err)
		}

		if len(reply.GetDockerfiles()) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// dockerfiles := svc.ClientSet.GitClient.GetDockerFiles(projectName)

		// if len(dockerfiles) == 0 {
		// 	w.WriteHeader(http.StatusNoContent)
		// 	return
		// }

		b, _ := json.Marshal(reply.GetDockerfiles())
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func (svc *Gateway) apiGetRoleListHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		conn := svc.ClientSet.RbacGrpcClient.GetConnection()

		c := rbacPb.NewRbacServiceClient(conn)
		reply, _ := c.GetRoles(context.TODO(), nil)

		roles := reply.GetRoles()
		b, _ := json.Marshal(roles)

		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func (svc *Gateway) apiGetRoleHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userName := vars["userName"]

		c := rbacPb.NewRbacServiceClient(svc.ClientSet.RbacGrpcClient.GetConnection())
		reply, err := c.GetRole(context.TODO(), &rbacPb.GetRoleRequest{UserName: userName})
		if err != nil {
			fmt.Println("GET USER ROLE GRPC ERR:", err)
		}

		b, _ := json.Marshal(reply.GetRole())

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func (svc *Gateway) apiGetUsersInProjectHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		proj := vars["projectName"]

		c := rbacPb.NewRbacServiceClient(svc.ClientSet.RbacGrpcClient.GetConnection())
		reply, err := c.GetUsers(context.TODO(), &rbacPb.GetUsersRequest{
			ProjectName: proj,
		})

		if err != nil {
			fmt.Println("GET USERS GRPC ERR:", err)
		}

		b, _ := json.Marshal(reply.GetUsers())
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func (svc *Gateway) apiGetProjectListHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := rbacPb.NewRbacServiceClient(svc.ClientSet.RbacGrpcClient.GetConnection())
		reply, err := c.GetProjects(context.TODO(), &emptypb.Empty{})
		if err != nil {
			fmt.Println("ERR GET PROJECT LIST :", err)
		}

		b, _ := json.Marshal(reply.GetProjects())

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func (svc *Gateway) apiDeleteResourcesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Query().Get("names")
		vars := mux.Vars(r)
		project := vars["projectName"]
		rsc := vars["resourceName"]

		if err := svc.ClientSet.KubeClient.DeleteResources(rsc, p, project); err != nil {
			fmt.Println("TEST DELETE SOME "+rsc+" FOR NAMESPACE "+project+" ERR:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
	}
}

func (svc *Gateway) apiGetConfigmapHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		proj := vars["projectName"]
		cm := vars["configmapName"]

		cms := svc.ClientSet.KubeClient.GetConfigmap(cm, proj)

		datas := []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}{}

		for k, v := range *cms {
			data := struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			}{
				Key:   k,
				Value: v,
			}

			datas = append(datas, data)
		}

		b, _ := json.Marshal(datas)
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func (svc *Gateway) apiGetConfigmapsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		proj := vars["projectName"]

		data := []struct {
			Name     string   `json:"name"`
			Services []string `json:"services"`
		}{}

		configmaps := svc.ClientSet.KubeClient.GetConfigmaps(proj)
		for _, cm := range *configmaps {
			fmt.Println("CM NAME:", cm.Name)
			service := svc.ClientSet.KubeClient.GetConfigmapMountedService(cm.Name, proj)
			services := []string{}
			for _, s := range service {
				services = append(services, s)

				data = append(data, struct {
					Name     string   `json:"name"`
					Services []string `json:"services"`
				}{
					Name:     cm.Name,
					Services: services,
				})
			}
		}

		b, _ := json.Marshal(data)
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func (svc *Gateway) apiGetSecretsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		proj := vars["projectName"]

		data := []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}{}

		secrets := *svc.ClientSet.KubeClient.GetSecret(proj)
		for k, v := range secrets {
			data = append(data, struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			}{
				Key:   k,
				Value: v,
			})
		}

		b, _ := json.Marshal(data)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func (svc *Gateway) apiGetResourcesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		proj := vars["projectName"]

		data := struct {
			Pods       *[]corev1.Pod           `json:"pod"`
			Services   *[]corev1.Service       `json:"service"`
			Ingresses  *[]networkingv1.Ingress `json:"ingress"`
			Configmaps *[]corev1.ConfigMap     `json:"configmap"`
			Secrets    *[]corev1.Secret        `json:"secret"`
		}{
			Pods:       svc.ClientSet.KubeClient.GetPods(proj),
			Services:   svc.ClientSet.KubeClient.GetServices(proj),
			Ingresses:  svc.ClientSet.KubeClient.GetIngresses(proj),
			Configmaps: svc.ClientSet.KubeClient.GetConfigmaps(proj),
			Secrets:    svc.ClientSet.KubeClient.GetSecrets(proj),
		}

		b, _ := json.Marshal(data)
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func (svc *Gateway) apiGetPoliciesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println()
		fmt.Println("GET GET POLICIES REQUEST")

		vars := mux.Vars(r)
		roleId := vars["roleId"]
		userId := r.URL.Query().Get("user_id")

		fmt.Println("USER ", userId, " tried to get policies")
		c := rbacPb.NewRbacServiceClient(svc.ClientSet.RbacGrpcClient.GetConnection())
		reply, err := c.GetPolicies(context.TODO(), &rbacPb.GetPoliciesRequest{
			RoleId: roleId,
		})
		if err != nil {
			fmt.Println("GET POLICIES GRPC ERR:", err)
		}

		policies := reply.GetPolicies()

		datas := []model.Policy{}
		data := model.Policy{}

		for _, p := range policies {
			pId, _ := strconv.Atoi(p.GetId())
			data.ID = pId
			data.Name = p.GetName()
			data.Json = p.GetJson()
			datas = append(datas, data)
		}

		b, _ := json.Marshal(datas)
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

// /api/manifest?username="choigonyok"
func (svc *Gateway) apiPostManifestHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		err := svc.ClientSet.GitClient.CreatePodManifestFile(username, env.Get("HOME_IDP_GIT_EMAIL"), "test", "v1.0", 8080)
		fmt.Println(err)
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}
}
