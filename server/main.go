package main

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "dapur-cokelat", 0)
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))

	conn.WriteMessages(kafka.Message{Value: []byte("hello user 1")})
}
