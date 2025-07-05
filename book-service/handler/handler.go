package handler

import (
	pb "github.com/dzuura/bookstore-microservice/book-service/proto"
	"context"
	"log"
)

func (s *Server) GetBook(ctx context.Context, req *pb.BookRequest) (*pb.BookResponse, error) {
	row := s.DB.QueryRow("SELECT id, title, author, price FROM books WHERE id=$1", req.Id)
	var id int32
	var title, author string
	var price float64
	err := row.Scan(&id, &title, &author, &price)
	if err != nil {
		log.Println("Error getting book:", err)
		return nil, err
	}
	return &pb.BookResponse{Id: id, Title: title, Author: author, Price: price}, nil
}

func (s *Server) ListBooks(ctx context.Context, _ *pb.Empty) (*pb.BookList, error) {
	rows, err := s.DB.Query("SELECT id, title, author, price FROM books")
	if err != nil {
		return nil, err
	}
	var books []*pb.BookResponse
	for rows.Next() {
		var id int32
		var title, author string
		var price float64
		err := rows.Scan(&id, &title, &author, &price)
		if err != nil {
			continue
		}
		books = append(books, &pb.BookResponse{Id: id, Title: title, Author: author, Price: price})
	}
	return &pb.BookList{Books: books}, nil
}
