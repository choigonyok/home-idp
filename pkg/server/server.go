package server

import (
	"fmt"

	"github.com/choigonyok/home-idp/pkg/config"
)

type Server struct {
	Config config.Config
}

func New(cfg config.Config) *Server {
	return &Server{
		Config: cfg,
	}
}

func (svr *Server) Run() error {
	fmt.Println("SERVER IS RUNING NOW")
	fmt.Println()
	fmt.Println("LOOK FOR port")
	result1, found1, err := svr.Config.Get("port")
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	if !found1 {
		fmt.Println("KEY_VALUE NOT FOUND")
	} else {
		fmt.Println("RESULT: ", result1)
	}

	fmt.Println()
	fmt.Println("LOOK FOR replicas")
	result2, found2, err := svr.Config.Get("replicas")
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	if !found2 {
		fmt.Println("KEY_VALUE NOT FOUND")
	} else {
		fmt.Println("RESULT: ", result2)
	}

	fmt.Println()
	fmt.Println("LOOK FOR Port")
	result3, found3, err := svr.Config.Get("Port")
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	if !found3 {
		fmt.Println("KEY_VALUE NOT FOUND")
	} else {
		fmt.Println("RESULT: ", result3)
	}

	fmt.Println()
	fmt.Println("LOOK FOR Replicas")
	result4, found4, err := svr.Config.Get("Replicas")
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	if !found4 {
		fmt.Println("KEY_VALUE NOT FOUND")
	} else {
		fmt.Println("RESULT: ", result4)
	}

	return nil
}
