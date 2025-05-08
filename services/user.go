package services

import (
	"fmt"

	"github.com/thongsoi/architectures/events"
)

// Simulate a channel as an event bus
var EventBus = make(chan events.Event, 5)

func RegisterUser(username string) {
	fmt.Printf("✅ User registered: %s\n", username)

	// Emit event
	eventBus <- events.Event{
		Type:    events.UserRegistered,
		Payload: username,
	}
}
