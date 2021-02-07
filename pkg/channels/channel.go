package channels


/// Notification
// handler.go -> interface
	// type handler interface {
		//send(notification, errchan)
	//}

	//func intitializeHandlers() {

	//}
// handler_factory.go
	// switch statement
		//return sepcific handler (NewClient)

// controller.go -> manager.go
	// func handleNotification(notification)
	// Args ; Notification
		// TODO: persist to DB, status -> in progress
		// WaitGroup <Size of channel >
		// map <channel type: errChan >
		// For each channel
			// wg.Add(1)
			// getHandler from factory
			// go handler.send(notification, errChan)
		// wg.Wait()
		// for each from errChan in map
		// TODO: Update DB status


/// Email
// handler.go

// Slack
// handler.go