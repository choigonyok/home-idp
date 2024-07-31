package server

import (
	"fmt"
	"net/http"

	"github.com/choigonyok/home-idp/pkg/env"
	"github.com/gorilla/mux"
)

type GatewayServer struct {
	Http *http.Server
}

func (s *GatewayServer) Close() error {
	return s.Http.Close()
}

func (s *GatewayServer) Run() {
	s.Http.ListenAndServe()
}

func New() *GatewayServer {
	r := mux.NewRouter()
	r.Handle("/test", http.HandlerFunc(Test)).Methods("GET")

	return &GatewayServer{
		Http: &http.Server{
			Addr:    ":" + env.Get("GATEWAY_SERVICE_PORT"),
			Handler: r,
		},
	}
}

func Test(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("REQ METHOD:", req.Method)
	resp.Write([]byte("HELLO WORLD"))
}
