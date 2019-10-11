package sender

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	. "github.com/kamilkoduo/diginavis/src/kafka-playground"
	"github.com/kamilkoduo/diginavis/src/kafka-playground/pb"
	"github.com/segmentio/kafka-go"
	"os"
)
var writer *kafka.Writer = NewWriter(WriterConfig{
	Topic: TOPIC_HELLO,
})
func Run() {
	fmt.Println("Starting on address " + os.Getenv("KAFKA_BROKER_HOST"))
	for i := 0; i < 10; i++ {
		out, err := proto.Marshal(&pb.HelloMessage{
			Uuid:                 uuid.New().String(),
			Content:              "hello from sender",
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		})
		if err != nil {
			fmt.Print("Error unmarshalling", err.Error())
			break
		}
		err = writer.WriteMessages(context.Background(), kafka.Message{
			Key:   nil,
			Value: out,
		})
		if err != nil {
			fmt.Print("Could not write messages", err.Error())
		}
	}
	_ = writer.Close()
}
