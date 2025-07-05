package main

import (
	"log"
	"net"

	"github.com/dzuura/bookstore-microservice/book-service/db"
	"github.com/dzuura/bookstore-microservice/book-service/handler"
	pb "github.com/dzuura/bookstore-microservice/book-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	database := db.InitDB()
	grpcServer := grpc.NewServer()

	pb.RegisterBookServiceServer(grpcServer, &handler.Server{DB: database})

	reflection.Register(grpcServer)
	
	log.Println("Book service running on port 50052")
	grpcServer.Serve(lis)
}
