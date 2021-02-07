package main

import (
	"log"
	"notificationhandler/pkg/config"
	"notificationhandler/pkg/kafka"
	"notificationhandler/pkg/notification"
	"os"
	"os/signal"
)

func main() {
	cfg := config.GetInstance()
	if err := notification.InitializeHandlers(cfg); err != nil {
		log.Fatal("failed to initialize channels", err)
	}

	go shutdownHook()
	kafka.Producer(cfg.KafkaProducers, cfg.KafkaTopic)
}


func shutdownHook() {
	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	<-signals

	notification.CloseHandlers()
}
