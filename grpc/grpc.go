package grpc

import (
	"Projects/Car24/car24_user_service/config"
	"Projects/Car24/car24_user_service/genproto/client_service"
	"Projects/Car24/car24_user_service/grpc/client"
	"Projects/Car24/car24_user_service/grpc/service"
	"Projects/Car24/car24_user_service/pkg/logger"
	"Projects/Car24/car24_user_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {

	grpcServer = grpc.NewServer()

	client_service.RegisterClientServiceServer(grpcServer, service.NewClientService(cfg, log, strg, srvc))

	reflection.Register(grpcServer)
	return
}
