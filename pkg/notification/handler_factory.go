package notification

import (
	"notificationhandler/pkg/channels/email"
	"notificationhandler/pkg/channels/slack"
	"notificationhandler/pkg/pb/notifier"
)

func getHandler(handlerType notifier.ChannelType) handler{
	switch handlerType {
	case notifier.ChannelType_EMAIL:
		return email.GetClient()
	case notifier.ChannelType_SLACK:
		return slack.GetClient()
	default:
		return nil
	}
}