package api

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-inspire/message/npool"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// https://github.com/grpc/grpc-go/issues/3794
// require_unimplemented_servers=false
type Server struct {
	npool.UnimplementedCloudHashingInspireServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterCloudHashingInspireServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterCloudHashingInspireHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
