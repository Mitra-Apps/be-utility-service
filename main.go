package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Mitra-Apps/be-utility-service/config/postgre"
	imagePostgreRepo "github.com/Mitra-Apps/be-utility-service/domain/image/repository/postgre"
	pb "github.com/Mitra-Apps/be-utility-service/domain/proto/image"
	grpcRoute "github.com/Mitra-Apps/be-utility-service/handler/grpc"
	"github.com/Mitra-Apps/be-utility-service/service"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedImageServiceServer
}

func main() {
	lis, err := net.Listen("tcp", os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db := postgre.Connection()
	usrRepo := imagePostgreRepo.NewPostgre(db)
	svc := service.New(usrRepo)
	grpcServer := grpc.NewServer()
	route := grpcRoute.New(svc)
	pb.RegisterImageServiceServer(grpcServer, route)

	fmt.Printf("GRPC Server listening at %v ", lis.Addr())
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v \n", err)
	}
}
