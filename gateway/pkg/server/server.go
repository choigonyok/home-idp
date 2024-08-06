package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	gwconfig "github.com/choigonyok/home-idp/gateway/pkg/config"
	"github.com/choigonyok/home-idp/gateway/pkg/grpc"
	"github.com/choigonyok/home-idp/install-manager/pkg/helm"
	"github.com/choigonyok/home-idp/pkg/config"
	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/choigonyok/home-idp/pkg/server"
	"github.com/gorilla/mux"
)

type Gateway struct {
	Http   server.Server
	Grpc   *grpc.GrpcClient
	Config config.Config
}

func (gw *Gateway) Close() error {
	if err := gw.Http.Close(); err != nil {
		return err
	}
	if err := gw.Grpc.Close(); err != nil {
		return err
	}
	return nil
}

func (gw *Gateway) Run() {
	gw.Http.Run()
}

func New(cfg *gwconfig.GatewayConfig) *Gateway {
	r := mux.NewRouter()
	svr := &Gateway{
		Http: &GatewayServer{
			Server: &http.Server{
				Addr:    ":" + env.Get("GATEWAY_SERVICE_PORT"),
				Handler: r,
			},
		},
		Grpc:   grpc.NewClient(),
		Config: cfg,
	}

	r.Handle("/test", http.HandlerFunc(svr.Test)).Methods("GET")
	r.Handle("/deploy", http.HandlerFunc(svr.Test2)).Methods("POST")
	r.Handle("/charts/argocd", http.HandlerFunc(svr.InstallArgoCDHandler)).Methods("POST")
	r.Handle("/charts/upgrade", http.HandlerFunc(svr.UpgradeArgoCDHandler)).Methods("POST")

	return svr
}

func (s *Gateway) Test(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("HELLO WORLD"))
	name := req.URL.Query().Get("name")
	email := req.URL.Query().Get("email")
	password := req.URL.Query().Get("password")
	projectId := req.URL.Query().Get("project_id")
	pid, _ := strconv.Atoi(projectId)

	ok, _ := s.Grpc.PutUser(email, name, password, int32(pid))
	fmt.Println("TEST REQUEST RESULT: ", ok.Succeed)
}

func (s *Gateway) Test2(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("HELLO2 WORLD2"))
	name := req.URL.Query().Get("name")
	email := req.URL.Query().Get("email")
	password := req.URL.Query().Get("password")
	projectId := req.URL.Query().Get("project_id")
	pid, _ := strconv.Atoi(projectId)

	ok, _ := s.Grpc.PutUser(email, name, password, int32(pid))
	fmt.Println("TEST REQUEST RESULT: ", ok.Succeed)
}

func (s *Gateway) InstallArgoCDHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("START")
	data := &helm.ArgoCDData{}

	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(data)

	ok, err := s.Grpc.InstallArgoCD(data)
	fmt.Println(err)
	fmt.Println(err)
	fmt.Println(err)
	fmt.Println("TEST REQUEST RESULT: ", ok.Succeed)
}

func (s *Gateway) UpgradeArgoCDHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("START")
	data := &helm.ArgoCDData{}

	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(data)

	ok, err := s.Grpc.InstallArgoCD(data)
	fmt.Println(err)
	fmt.Println(err)
	fmt.Println(err)
	fmt.Println("TEST REQUEST RESULT: ", ok.Succeed)
}
