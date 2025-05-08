package handlers

import (
	"fmt"

	"github.com/thongsoi/architectures/events"
)

func HandleAnalyticsOnUserRegistered(event events.Event) {
	userName := event.Payload.(string)
	fmt.Printf("📊 Logging user registration for analytics: %s\n", userName)
}
