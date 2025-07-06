package main

import (
	"log"
	"net"

	"github.com/dzuura/bookstore-microservice/order-service/db"
	"github.com/dzuura/bookstore-microservice/order-service/handler"
	pb "github.com/dzuura/bookstore-microservice/order-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	database := db.InitDB()
	grpcServer := grpc.NewServer()
	
	pb.RegisterOrderServiceServer(grpcServer, &handler.Server{DB: database})

	reflection.Register(grpcServer)

	log.Println("Order service running on port 50053")
	grpcServer.Serve(lis)
}