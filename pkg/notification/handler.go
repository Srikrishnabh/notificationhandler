package notification

import (
	"log"
	"notificationhandler/pkg/channels/email"
	"notificationhandler/pkg/channels/slack"
	"notificationhandler/pkg/config"
	"notificationhandler/pkg/pb/notifier"
	"sync"
)

type handler interface {
	Send(notification notifier.Notification, errChan chan error, wg *sync.WaitGroup)
}


func InitializeHandlers(cfg *config.Config) error {
	if _, err := email.NewClient(cfg.EmailServerPort, cfg.EmailServerAddress); err != nil {
		return err
	}

	if _, err := slack.NewClient(cfg.SlackServerPort, cfg.SlackServerAddress); err != nil {
		return err
	}

	return nil
}

func CloseHandlers() {
	if err := email.GetClient().Close(); err != nil {
		log.Print("error closing email server connection", err)
	}

	if err := slack.GetClient().Close(); err != nil {
		log.Print("error closing slack server connection", err)
	}
}