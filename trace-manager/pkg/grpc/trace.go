package grpc

import (
	"github.com/choigonyok/home-idp/pkg/storage"
	pb "github.com/choigonyok/home-idp/trace-manager/pkg/proto"
)

type TraceServiceServer struct {
	pb.UnimplementedTraceServiceServer
	StorageClient storage.StorageClient
}

// METHODS IN HERE
