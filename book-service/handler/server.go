package handler

import (
	pb "github.com/dzuura/bookstore-microservice/book-service/proto"
	"database/sql"
)

type Server struct {
	pb.UnimplementedBookServiceServer
	DB *sql.DB
}