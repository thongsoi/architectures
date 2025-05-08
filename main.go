package main

import (
	"github.com/thongsoi/architectures/handlers"
	"github.com/thongsoi/architectures/services"

	"github.com/thongsoi/architectures/events"
)

func init() {
	// Subscribe handlers to events
	go func() {
		for {
			event := <-services.EventBus
			switch event.Type {
			case events.UserRegistered:
				handlers.HandleEmailOnUserRegistered(event)
				handlers.HandleAnalyticsOnUserRegistered(event)
			}
		}
	}()
}

func main() {
	// Simulate user registration
	services.RegisterUser("john_doe")
	services.RegisterUser("jane_smith")

	// Wait for events to be processed
	select {}
}
