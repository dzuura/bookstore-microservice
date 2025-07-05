package main

import (
	"log"
	"net"

	"github.com/dzuura/bookstore-microservice/user-service/db"
	"github.com/dzuura/bookstore-microservice/user-service/handler"
	pb "github.com/dzuura/bookstore-microservice/user-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    db := db.InitDB()
    grpcServer := grpc.NewServer()

    pb.RegisterUserServiceServer(grpcServer, &handler.Server{DB: db})

    reflection.Register(grpcServer)

    log.Println("User service running on port 50051")
    grpcServer.Serve(lis)
}

