package handler

import (
	"context"
	"encoding/json"
	"log"
	"fmt"
	pb "github.com/dzuura/bookstore-microservice/order-service/proto"
	grpcclients "github.com/dzuura/bookstore-microservice/order-service/grpc_clients"
    notificationpb "github.com/dzuura/bookstore-microservice/notification-service/proto"
	"github.com/nats-io/nats.go"
)

func (s *Server) PlaceOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
    user, err := grpcclients.GetUserInfo(req.UserId)
    if err != nil {
        return nil, fmt.Errorf("user not found")
    }

    book, err := grpcclients.GetBookInfo(req.BookId)
    if err != nil {
        return nil, fmt.Errorf("book not found")
    }

    log.Printf("Placing order for user %s and book %s\n", user.Name, book.Title)

    var orderID int
    err = s.DB.QueryRow(
        "INSERT INTO orders (user_id, book_id, quantity) VALUES ($1, $2, $3) RETURNING id",
        req.UserId, req.BookId, req.Quantity,
    ).Scan(&orderID)
    if err != nil {
        log.Println("Error placing order:", err)
        return nil, err
    }

    conn, err := nats.Connect("nats://nats:4222")
    if err == nil {
        notification := notificationpb.OrderNotification{
            OrderId:  int32(orderID),
            UserId:   req.UserId,
            BookId:   req.BookId,
            Quantity: req.Quantity,
        }
        data, _ := json.Marshal(notification)
        conn.Publish("order.placed", data)
        conn.Close()
    } else {
        log.Println("Failed to connect to NATS:", err)
    }

    return &pb.OrderResponse{Message: "Order placed", OrderId: int32(orderID)}, nil
}
