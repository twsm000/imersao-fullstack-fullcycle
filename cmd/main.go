package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/twsm000/imersao-fullstack-fullcycle/application/grpc"
	"github.com/twsm000/imersao-fullstack-fullcycle/infrastructure/db"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
