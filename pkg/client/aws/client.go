package client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

type AWS struct {
	EKSClient *eks.Client
}

func NewFromEnv() *AWS {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}
	e := eks.NewFromConfig(cfg)
	return &AWS{EKSClient: e}
}
