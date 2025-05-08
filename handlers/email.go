package handlers

import (
	"fmt"

	"github.com/thongsoi/architectures/events"
)

func HandleEmailOnUserRegistered(event events.Event) {
	userName := event.Payload.(string)
	fmt.Printf("📧 Sending welcome email to user: %s\n", userName)
}
