// Sources for https://watermill.io/docs/getting-started/
package main

import (
	"context"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/infrastructure/googlecloud"
)

const projectID = "theta-disk-241906"
const topicName = "events"

func main() {
	subscriber, err := googlecloud.NewSubscriber(
		context.Background(),
		googlecloud.SubscriberConfig{
			GenerateSubscriptionName: func(topic string) string {
				return topic + "-evesdropping"
			},
			ProjectID: projectID,
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	messages, err := subscriber.Subscribe(context.Background(), topicName)
	if err != nil {
		panic(err)
	}

	process(messages)
}

func process(messages <-chan *message.Message) {
	for msg := range messages {
		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))
		msg.Ack()
	}
}
