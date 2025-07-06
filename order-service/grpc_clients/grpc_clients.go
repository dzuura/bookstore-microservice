package grpc_clients

import (
    "context"
    "log"
    "time"

    userpb "github.com/dzuura/bookstore-microservice/user-service/proto"
    bookpb "github.com/dzuura/bookstore-microservice/book-service/proto"

    "google.golang.org/grpc"
)

func GetUserInfo(userID int32) (*userpb.UserResponse, error) {
    conn, err := grpc.Dial("user-service:50051", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
    if err != nil {
        log.Println("Failed to connect to user service:", err)
        return nil, err
    }
    defer conn.Close()

    client := userpb.NewUserServiceClient(conn)
    resp, err := client.GetUser(context.Background(), &userpb.UserRequest{Id: userID})
    if err != nil {
        log.Println("Error calling GetUser:", err)
        return nil, err
    }

    return resp, nil
}

func GetBookInfo(bookID int32) (*bookpb.BookResponse, error) {
    conn, err := grpc.Dial("book-service:50052", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
    if err != nil {
        log.Println("Failed to connect to book service:", err)
        return nil, err
    }
    defer conn.Close()

    client := bookpb.NewBookServiceClient(conn)
    resp, err := client.GetBook(context.Background(), &bookpb.BookRequest{Id: bookID})
    if err != nil {
        log.Println("Error calling GetBook:", err)
        return nil, err
    }

    return resp, nil
}
