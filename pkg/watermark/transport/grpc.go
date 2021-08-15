package transport

import (
	"microgo/pkg/watermark"
	"microgo/pkg/watermark/endpoints"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	get           grpctransport.Handler
	status        grpctransport.Handler
	addDocument   grpctransport.Handler
	watermark     grpctransport.Handler
	serviceStatus grpctransport.Handler
}

func NewGrpcServer(ep endpoints.Set) watermark.WatermarkServer {

}
