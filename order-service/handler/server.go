package handler

import (
	pb "github.com/dzuura/bookstore-microservice/order-service/proto"
	"database/sql"
)

type Server struct {
	pb.UnimplementedOrderServiceServer
	DB *sql.DB
}