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

	deployPb "github.com/choigonyok/home-idp/deploy-manager/pkg/proto"
	"github.com/choigonyok/home-idp/gateway/pkg/progress"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/git"
	"github.com/choigonyok/home-idp/pkg/model"
	"github.com/choigonyok/home-idp/pkg/util"
	rbacPb "github.com/choigonyok/home-idp/rbac-manager/pkg/proto"
	"github.com/google/go-github/v63/github"
	"github.com/gorilla/mux"
)

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
		step := progress.NewStep(progress.PushManifest, progress.Continue, nil)
		fmt.Println(r.Body)
		b, err := io.ReadAll(r.Body)
		fmt.Println("TESTHANDLER3:", err)
		m := make(map[string]interface{})
		err = json.Unmarshal(b, &m)

		img := util.ParseInterfaceMap(m, []string{"event_data", "resources", "resource_url"}).(string)
		s := strings.Split(img, "/")
		name, version, _ := strings.Cut(s[2], ":")
		step.Add(name + ":" + version)

		fmt.Println(err)
		fmt.Println("MAPPED MAP:", name, version)

		step.UpdateLog("START TO UPDATE MANIFEST")
		if err := svc.ClientSet.GitClient.UpdateManifest(name + ":" + version); err != nil {
			step.UpdateState(progress.Fail, err.Error())
			fmt.Println("TEST UPDATE IMAGE FROM MANIFEST ERR:", err)
			return
		}
		step.UpdateState(progress.Success, "SUCCESS TO MANIFEST!")
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

		c := rbacPb.NewRbacServiceClient(svc.ClientSet.GrpcClient[util.Components(util.RbacManager)].GetConnection())
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
				if fn := getFilename(event); fn != ".gitkeep" {
					svc.requestArgoCDWebhook(r, payload)
				}
			default:
				return
			}
		case *github.RepositoryEvent:
			fmt.Println("TEST PING WEBHOOK RECEIVED")
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

// func (svc *Gateway) requestLogin(username, password string) {
// 	c := rbacPb.NewLoginServiceClient(svc.ClientSet.GrpcClient[util.Components(util.RbacManager)].GetConnection())
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

func (svc *Gateway) requestBuildDockerfile(name, version, username string) {
	step := progress.NewStep(progress.DeployKaniko, progress.Continue, nil)
	step.Add(name + ":" + version)
	step.UpdateLog("Deploy-manager grpc client creating...")
	c := deployPb.NewBuildClient(svc.ClientSet.GrpcClient[util.Components(util.DeployManager)].GetConnection())

	step.UpdateLog("Deploy Kaniko to build image...")
	reply, err := c.BuildDockerfile(context.TODO(), &deployPb.BuildDockerfileRequest{
		Img: &deployPb.Image{
			Pusher:  username,
			Name:    name,
			Version: version,
		},
	})
	if err != nil {
		fmt.Println("TEST BUILD DOCKERFILE REQUEST ERR:", err)
		step.UpdateState(progress.Fail, err.Error())
		return
	}

	if !reply.Succeed {
		fmt.Println("TEST BUILD DOCKERFILE REQUEST FAILED")
		step.UpdateState(progress.Fail, "FAILED")
		return
	}
	step.UpdateState(progress.Success, "SUCCESS!")
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

		w.WriteHeader(http.StatusOK)
	}
}

// func (svc *Gateway) apiPutUserHandler() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		projectName := vars["projectName"]
// 		userName := vars["userName"]

// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		b, _ := io.ReadAll(r.Body)

// 		usr := model.User{}
// 		json.Unmarshal(b, &usr)
// 	}
// }

func (svc *Gateway) apiPostUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		projectName := vars["projectName"]

		w.Header().Set("Access-Control-Allow-Origin", "*")
		b, _ := io.ReadAll(r.Body)

		usr := model.User{}
		json.Unmarshal(b, &usr)

		c := rbacPb.NewRbacServiceClient(svc.ClientSet.GrpcClient[util.Components(util.RbacManager)].GetConnection())
		reply, err := c.PostUser(context.TODO(), &rbacPb.PostUserRequest{
			User: &rbacPb.User{
				RoleId: usr.RoleID,
				Name:   usr.Name,
			},
			ProjectName: projectName,
		})
		if err != nil {
			fmt.Println("POST USER GRPC ERR:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if reply.Error != nil {
			fmt.Println("POST USER GRPC ERR2:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (svc *Gateway) apiPostDockerfileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET POST DOCKER REQUEST")
		fmt.Println("TEST REQEUST BODY:", r.Body)

		w.Header().Set("Access-Control-Allow-Origin", "*")

		b, _ := io.ReadAll(r.Body)

		f := git.GitDockerFile{}
		json.Unmarshal(b, &f)

		step := progress.NewStep(progress.PushDockerfile, progress.Continue, []string{"GET POST DOCKER REQUEST", fmt.Sprintln("TEST REQEUST BODY:", r.Body)})

		step.Add(f.Image)

		imageName, _, _ := strings.Cut(f.Image, ":")
		if svc.ClientSet.GitClient.IsDockerfileExist(f.Username, imageName) {
			fmt.Println("TEST DOCKERFILE ALREADY EXIST")
			svc.ClientSet.GitClient.UpdateDockerFile(f.Username, f.Image, f.Content)
			step.UpdateState(progress.Fail, "TEST DOCKERFILE ALREADY EXIST")
			return
		}
		step.UpdateLog("DOCKERFILE NOT EXIST! START TO CREATE...")
		if err := svc.ClientSet.GitClient.CreateDockerFile(f.Username, f.Image, f.Content); err != nil {
			step.UpdateState(progress.Fail, err.Error())
		}
		step.UpdateState(progress.Success, "CREATE DOCKERFILE FINISH!")
	}
}

func (svc *Gateway) apiGetDockerfileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userName := vars["userName"]
		projectName := vars["projectName"]

		fmt.Println("GET GET DOCKER REQUEST")
		dockerfiles := svc.ClientSet.GitClient.GetDockerFiles(projectName, userName)
		if len(dockerfiles) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(dockerfiles)
	}
}

func (svc *Gateway) apiGetRolesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println()
		fmt.Println("GET GET ROLES REQUEST")

		userId := r.URL.Query().Get("user_id")
		fmt.Println("USER ID ", userId, " tried to get every roles")
		c := rbacPb.NewRbacServiceClient(svc.ClientSet.GrpcClient[util.Components(util.RbacManager)].GetConnection())
		reply, err := c.GetRoles(context.TODO(), &rbacPb.GetRolesRequest{})
		if err != nil {
			fmt.Println("GET ROLES GRPC ERR:", err)
		}

		datas := []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}{}
		data := struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}{}

		roles := reply.GetRoles()
		for _, role := range roles {
			fmt.Println(role.Id)
			fmt.Println(role.Name)
			roleId, _ := strconv.Atoi(role.GetId())
			data.ID = roleId
			data.Name = role.GetName()
			datas = append(datas, data)
		}

		b, _ := json.Marshal(datas)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func (svc *Gateway) apiGetRoleHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println()
		fmt.Println("GET GET ROLE REQUEST")

		vars := mux.Vars(r)
		userId := vars["userId"]

		fmt.Println("USER ID ", userId, " tried to get his role")
		c := rbacPb.NewRbacServiceClient(svc.ClientSet.GrpcClient[util.Components(util.RbacManager)].GetConnection())
		reply, err := c.GetRole(context.TODO(), &rbacPb.GetRoleRequest{UserId: userId})
		if err != nil {
			fmt.Println("GET USER ROLE GRPC ERR:", err)
		}

		role := reply.GetRole()
		fmt.Println("---USER " + userId + " ROLE---")
		fmt.Println(role.Id)
		fmt.Println(role.Name)
		fmt.Println("---ROLE END---")
		roleId, _ := strconv.Atoi(role.GetId())

		data := struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}{
			ID:   roleId,
			Name: role.GetName(),
		}

		b, _ := json.Marshal(data)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(b)

	}
}

func (svc *Gateway) apiGetUsersInProjectHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		proj := vars["projectName"]

		c := rbacPb.NewRbacServiceClient(svc.ClientSet.GrpcClient[util.Components(util.RbacManager)].GetConnection())
		reply, err := c.GetUsers(context.TODO(), &rbacPb.GetUsersRequest{
			ProjectName: proj,
		})

		if err != nil {
			fmt.Println("GET USERS GRPC ERR:", err)
		}

		usrs := []model.User{}

		for _, usr := range reply.GetUsers() {
			usrs = append(usrs, model.User{
				ID:         usr.Id,
				Name:       usr.Name,
				RoleID:     usr.RoleId,
				CreateTime: usr.CreateTime,
			})
		}

		b, _ := json.Marshal(usrs)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func (svc *Gateway) apiGetProjectsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := ""
		uid = r.URL.Query().Get("user_id")

		c := rbacPb.NewRbacServiceClient(svc.ClientSet.GrpcClient[util.Components(util.RbacManager)].GetConnection())
		reply, err := c.GetProjects(context.TODO(), &rbacPb.GetProjectsRequest{
			UserId: uid,
		})

		if err != nil {
			fmt.Println("GET PROJECTS GRPC ERR:", err)
		}

		projects := []model.Project{}

		for _, p := range reply.GetProjects() {
			projects = append(projects, model.Project{
				ID:      p.Id,
				Name:    p.Name,
				Creator: p.Creator,
			})
		}

		b, _ := json.Marshal(projects)

		w.Header().Set("Access-Control-Allow-Origin", "*")
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

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func (svc *Gateway) apiGetResourcesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		proj := vars["projectName"]

		data := struct {
			Pods       []model.Pod       `json:"pod"`
			Services   []model.Service   `json:"service"`
			Ingresses  []model.Ingress   `json:"ingress"`
			Configmaps []model.Configmap `json:"configmap"`
			Secrets    []model.Secret    `json:"secret"`
		}{
			Pods:       *svc.ClientSet.KubeClient.GetPods(proj),
			Services:   *svc.ClientSet.KubeClient.GetServices(proj),
			Ingresses:  *svc.ClientSet.KubeClient.GetIngresses(proj),
			Configmaps: *svc.ClientSet.KubeClient.GetConfigmaps(proj),
			Secrets:    *svc.ClientSet.KubeClient.GetSecrets(proj),
		}

		b, _ := json.Marshal(data)

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
		c := rbacPb.NewRbacServiceClient(svc.ClientSet.GrpcClient[util.Components(util.RbacManager)].GetConnection())
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

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

// /api/manifest?username="choigonyok"
func (svc *Gateway) apiPostManifestHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		err := svc.ClientSet.GitClient.CreatePodManifestFile(username, env.Get("HOME_IDP_GIT_EMAIL"), "test:v1.0", 8080)
		fmt.Println(err)

		w.Header().Set("Access-Control-Allow-Origin", "*")
	}
}
