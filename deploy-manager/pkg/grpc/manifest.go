package grpc

import (
	"context"

	pb "github.com/choigonyok/home-idp/deploy-manager/pkg/proto"
)

type ManifestServiceServer struct {
	pb.UnimplementedManifestServiceServer
}

func (s *ManifestServiceServer) ApplyManifest(ctx context.Context, in *pb.ApplyManifestRequest) (*pb.SuccessReply, error) {
	return nil, nil
	// in.Kind.GetKind()
}

func (s *ManifestServiceServer) DeleteManifest(ctx context.Context, in *pb.DeleteManifestRequest) (*pb.SuccessReply, error) {
	return nil, nil
}
