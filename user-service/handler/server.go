package handler

import (
	pb "github.com/dzuura/bookstore-microservice/user-service/proto"
	"database/sql"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	DB *sql.DB
}