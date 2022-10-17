package messaging
import (
	"context"
	"encoding/json"
	"log"
	kafka "github.com/segmentio/kafka-go"
	"server/utils"
	"server/models"
)

func ConsumeEmail(){
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{brokerAddress},
		Topic:    topic,
		MinBytes: 10e3, 
		MaxBytes: 10e6, 
	})
	defer reader.Close()
	for{
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("failed to read message:", err)
			continue
		}
		email := models.Email{}
 
		_ = json.Unmarshal(m.Value, &email)
		utils.SendEmail(email)
	}
}
