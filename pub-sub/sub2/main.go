package main

import (
	"context"
	"log"
	"net/http"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
)

// var pubsubName = os.Getenv("DAPR_PUBSUB_NAME")

const (
	pubsubName = "pubsub"
	topic      = "testTopic"
)

func main() {
	sub := &common.Subscription{
		PubsubName: pubsubName,
		Topic:      topic,
		Route:      "gadget",
	}
	s := daprd.NewService(":18889")
	s.AddBindingInvocationHandler("gadget", func(ctx context.Context, e *common.BindingEvent) (out []byte, err error) {
		log.Printf("gadget event - Data: %s\n", string(e.Data))
		return nil, nil
	})
	if err := s.AddTopicEventHandler(sub, func(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
		log.Printf("orders event - PubsubName: %s, Topic: %s, ID: %s, Data: %s\n", e.PubsubName, e.Topic, e.ID, e.Data)
		return false, nil
	}); err != nil {
		log.Fatalln("orders", err)
	}

	if err := s.Start(); err != http.ErrServerClosed {
		log.Fatalln("orders", err)
	}
}
