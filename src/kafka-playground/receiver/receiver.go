package receiver

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	. "github.com/kamilkoduo/diginavis/src/kafka-playground"
	"github.com/kamilkoduo/diginavis/src/kafka-playground/pb"
	"github.com/segmentio/kafka-go"
	"os"
)

var reader *kafka.Reader = NewReader(ReaderConfig{
	Topic: TOPIC_HELLO,
})
func Run() {
	fmt.Println("Starting on address " + os.Getenv("KAFKA_BROKER_HOST"))
	// Read loop
	for {
		// Read the next message from Kafka
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Print(err.Error())
			break
		}
		var hw = &pb.HelloMessage{}
		err = proto.Unmarshal(m.Value, hw)
		if err != nil {
			fmt.Print("Error unmarshalling", err.Error())
			break
		}
		fmt.Printf("New message received!\n>>UUID: '%s'\n>>Contents: '%s'\n", hw.GetUuid(), hw.GetContent())
	}
	_ = reader.Close()
}

