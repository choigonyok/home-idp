package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/choigonyok/home-idp/gateway/pkg/client"
	"github.com/choigonyok/home-idp/install-manager/pkg/helm"
	"github.com/choigonyok/home-idp/pkg/util"
)

func (svr *Gateway) RegisterRoutes() {
	svr.Router.registerRoute(http.MethodGet, "/test0", svr.InstallArgoCDHandler())
	svr.Router.registerRoute(http.MethodGet, "/test1", svr.TestHandler1())
	svr.Router.registerRoute(http.MethodGet, "/test2", svr.TestHandler2())
	svr.Router.registerRoute(http.MethodGet, "/test3", svr.TestHandler3())
}

func (svr *Gateway) InstallArgoCDHandler() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		data := &helm.ArgoCDData{}

		if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
			http.Error(resp, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println(data)

		ok, err := client.InstallArgoCDChart(data, svr.ClientSet.GrpcClient[util.InstallManager].GetConnection())
		fmt.Println(ok)
		fmt.Println(err)
	}
}

func (svr *Gateway) TestHandler1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (svr *Gateway) TestHandler2() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (svr *Gateway) TestHandler3() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
