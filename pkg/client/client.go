package client

import "github.com/choigonyok/home-idp/pkg/util"

type ClientOption interface {
	Apply(ClientSet) error
}

type ClientSet interface {
	Set(util.Clients, interface{})
}
