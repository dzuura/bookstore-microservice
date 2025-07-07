package subscriber

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	notificationpb "github.com/dzuura/bookstore-microservice/notification-service/proto"
)

func SubscribeToOrderPlaced() error {
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		return err
	}

	_, err = nc.Subscribe("order.placed", func(m *nats.Msg) {
		var notif notificationpb.OrderNotification
		if err := json.Unmarshal(m.Data, &notif); err != nil {
			log.Println("Failed to unmarshal notification:", err)
			return
		}

		log.Printf("📬 Notifikasi: Order #%d oleh User #%d untuk Buku #%d (Qty: %d)",
			notif.OrderId, notif.UserId, notif.BookId, notif.Quantity)
	})
	if err != nil {
		return err
	}

	log.Println("✅ Listening on NATS topic: order.placed")
	select {} // keep service running
}
