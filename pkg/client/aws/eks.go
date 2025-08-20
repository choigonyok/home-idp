package client

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/eks"
)

func (a *AWS) ListEKSClusters(ctx context.Context) ([]string, error) {
	out, err := a.EKSClient.ListClusters(ctx, &eks.ListClustersInput{})
	if err != nil {
		return nil, err
	}
	return out.Clusters, nil
}

func (a *AWS) DescribeCluster(ctx context.Context, name string) (string, error) {
	out, err := a.EKSClient.DescribeCluster(ctx, &eks.DescribeClusterInput{Name: &name})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("cluster %s status=%s", name, out.Cluster.Status), nil
}
