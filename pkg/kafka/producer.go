package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"notificationhandler/pkg/pb/notifier"
)

func Producer(producers []string, topic string) {
	producer, err := sarama.NewAsyncProducer(producers, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()


	notifications := make([]*notifier.Notification, 0)
	generateNotifications(&notifications)

	var enqueued, producerErrors int
	for _, notification := range notifications {
		notificationMessage, err := json.Marshal(notification)
		if err != nil {
			fmt.Println("json marshal error", err)
			continue
		}

		select {
		case producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: sarama.StringEncoder(notification.To),
														Value: sarama.ByteEncoder(notificationMessage)}:
			enqueued++
		case err := <-producer.Errors():
			log.Println("Failed to produce message", err)
			producerErrors++
		}
	}

	log.Printf("Enqueued: %d; errors: %d\n", enqueued, producerErrors)
}


func generateNotifications(notifications *[]*notifier.Notification) {
	for i :=0; i <3; i ++ {
		notification := &notifier.Notification{
			To:                   fmt.Sprintf("%d", i),
			From:                 "bot@target.com",
			Body:                 fmt.Sprintf("order %d placed", i),
			Subject:              "order status",
			Channels:             []*notifier.Channel{
				{
					Type:     notifier.ChannelType_EMAIL,
					Metadata: map[string]string{"EmailID" : fmt.Sprintf("consumer%d@gmail.com", i)},
				},
				{
					Type:     notifier.ChannelType_SLACK,
					Metadata: map[string]string{"SlackWebhook" : fmt.Sprintf("consumer%d@slack.com", i)},
				},
			},
		}

		*notifications = append(*notifications, notification)
	}
}