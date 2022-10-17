package messaging
import (
	"context"
	"encoding/json"
	"time"
	"log"
	kafka "github.com/segmentio/kafka-go"
	"server/models"
)
const (
	topic			= "email"
	partition		= 0
	brokerAddress	= "kafka:9092"
)

func ProduceEmail(email models.Email){
	b, err := json.Marshal(email)
    if err != nil {
        panic(err)
    }

	conn, err := kafka.DialLeader(context.Background(), "tcp", brokerAddress, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10*time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: b},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
	
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
