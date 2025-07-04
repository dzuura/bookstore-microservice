package handler

import (
	"context"
	"log"
	pb "github.com/dzuura/bookstore-microservice/user-service/proto"
)

func (s *Server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	row := s.DB.QueryRow("SELECT id, name, email FROM users WHERE id=$1", req.Id)
	var id int32
	var name, email string
	err := row.Scan(&id, &name, &email)
	if err != nil {
		log.Println("Error getting user:", err)
		return nil, err
	}
	return &pb.UserResponse{Id: id, Name: name, Email: email}, nil
}