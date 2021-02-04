package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"github.com/twsm000/imersao-fullstack-fullcycle/application/grpc/pb"
	"github.com/twsm000/imersao-fullstack-fullcycle/application/usecase"
	"github.com/twsm000/imersao-fullstack-fullcycle/infrastructure/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// StartGrpcServer ...
func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pixRepo := &repository.PixKeyRepositoryDB{Db: database}
	pixUseCase := &usecase.PixUseCase{PixKeyRepository: pixRepo}
	pixGrpcService := NewPixGrpcService(pixUseCase)
	pb.RegisterPixServiceServer(grpcServer, pixGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("can't start grpc server", err)
	}

	log.Printf("gPRC server has been started on port %d", port)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}
}
