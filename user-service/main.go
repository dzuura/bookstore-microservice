package main

import (
	"log"
	"net"

	"github.com/dzuura/bookstore-microservice/user-service/db"
	"github.com/dzuura/bookstore-microservice/user-service/handler"
	pb "github.com/dzuura/bookstore-microservice/user-service/proto"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	database := db.InitDB()
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &handler.Server{DB: database})
	log.Println("User service running on port 50051")
	grpcServer.Serve(lis)
}
