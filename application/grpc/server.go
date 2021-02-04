package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

// StartGrpcServer ...
func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()

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
