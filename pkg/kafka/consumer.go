package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"notificationhandler/pkg/notification"
	"notificationhandler/pkg/pb/notifier"
	"log"
)

type consumerGroupHandler struct{}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		var notificationObj notifier.Notification
		if err := json.Unmarshal(msg.Value, &notificationObj); err != nil {
			fmt.Println("error while unmarshal the message", err.Error())
			continue
		}

		notification.HandleNotification(notificationObj)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func Consumer(consumerGroups []string, topics []string, groupID string) {
	config := sarama.NewConfig()


	client, err := sarama.NewClient(consumerGroups, config)
	if err != nil {
		log.Println(err)
	}
	defer func() { _ = client.Close() }()

	group, err := sarama.NewConsumerGroupFromClient(groupID, client)
	if err != nil {
		panic(err)
	}
	defer func() { _ = group.Close() }()

	// Track errors
	go func() {
		for err := range group.Errors() {
			fmt.Println("ERROR", err)
		}
	}()

	// Iterate over consumer sessions.
	ctx := context.Background()
	for {
		handler := &consumerGroupHandler{}
		topicsn := topics

		// `Consume` should be called inside an infinite loop, when a
		// server-side rebalance happens, the consumer session will need to be
		// recreated to get the new claims
		err := group.Consume(ctx, topicsn, handler)
		if err != nil {
			panic(err)
		}

	}
}
