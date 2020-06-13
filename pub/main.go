// Sources for https://watermill.io/docs/getting-started/
package main

import (
	"flag"
	"log"

	stan "github.com/nats-io/stan.go"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-nats/pkg/nats"
	"github.com/ThreeDotsLabs/watermill/message"
)

var (
	clientID string
)

func main() {
	flag.StringVar(&clientID, "id", "stan-pub", "The NATS Streaming client ID to connect with")

	log.SetFlags(0)
	flag.Parse()
	log.Println(clientID)
	publisher, err := nats.NewStreamingPublisher(
		nats.StreamingPublisherConfig{
			ClusterID: "nats-streaming",
			ClientID:  clientID,
			StanOptions: []stan.Option{
				stan.NatsURL("nats://nats-streaming:4222"),
			},
			Marshaler: nats.GobMarshaler{},
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}
	publishMessages(publisher)

}

func publishMessages(publisher message.Publisher) {
	for {
		msg := message.NewMessage(watermill.NewUUID(), []byte("Hello, world!"))
		log.Printf("send message: %s, payload: %s", msg.UUID, string(msg.Payload))
		if err := publisher.Publish("example.topic", msg); err != nil {
			panic(err)
		}

		//time.Sleep(time.Second)
	}
}
