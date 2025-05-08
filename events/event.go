package events

type EventType string

const (
	UserRegistered EventType = "UserRegistered"
)

type Event struct {
	Type    EventType
	Payload any
}
