package notification

import(
	"errors"
	"fmt"
	"notificationhandler/pkg/pb/notifier"
	"sync"
	"time"
)



func HandleNotification(notification notifier.Notification) {
	//TODO: persists data to DB with status in-progress
	wg := &sync.WaitGroup{}
	channelTypeToErrChanMap := make(map[notifier.ChannelType]chan error)
	channelToErrMap := make(map[notifier.ChannelType]error)
	for _, channel := range notification.Channels {
		errChan := make(chan error)
		channelTypeToErrChanMap[channel.Type] = errChan
		handler := getHandler(channel.Type)
		if handler == nil {
			channelToErrMap[channel.Type] = errors.New("channel not available")
			continue
		}
		wg.Add(1)
		go handler.Send(notification, errChan, wg)
	}

	for channel, errChan := range channelTypeToErrChanMap {
		select {
			case err := <- errChan:
				channelToErrMap[channel] = err
			case <-time.After(2 *time.Second):
				channelToErrMap[channel] = errors.New(fmt.Sprintf("channel %s timeout", channel))
		}
	}
	//TODO: persists the error data to DB for retries
	wg.Wait()
}
