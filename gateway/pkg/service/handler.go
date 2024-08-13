package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/choigonyok/home-idp/gateway/pkg/grpc"
	"github.com/choigonyok/home-idp/install-manager/pkg/helm"
	"github.com/choigonyok/home-idp/pkg/util"
)

func (svc *Gateway) InstallArgoCDHandler() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		data := &helm.ArgoCDData{}

		if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
			http.Error(resp, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println(data)

		ok, err := grpc.InstallArgoCDChart(data, svc.ClientSet.GrpcClient[util.InstallManager].GetConnection())
		fmt.Println(ok)
		fmt.Println(err)
	}
}

func (svc *Gateway) TestHandler1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (svc *Gateway) TestHandler2() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (svc *Gateway) TestHandler3() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
