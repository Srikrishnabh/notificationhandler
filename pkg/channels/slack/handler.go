package slack

import (
	"fmt"
	"notificationhandler/pkg/pb/notifier"
	"sync"
)

type Client struct {
	port string
	address string
}

var once sync.Once
var client *Client

func NewClient(port, address string) (*Client, error){
	once.Do(func() {
		client = &Client{port: port, address: address}
	})
	return client, nil
}

func GetClient() *Client {
	return client
}

func (c *Client) Send(notification notifier.Notification, errChan chan error, wg *sync.WaitGroup) {
	for _, channel := range notification.Channels {
		if channel.Type == notifier.ChannelType_SLACK {
			fmt.Printf("\nsending message via slack to %s\n", channel.Metadata["SlackWebhook"])
		}
	}
	errChan <- nil
	wg.Done()
}


func (c *Client) Close() error {
	return nil
}
