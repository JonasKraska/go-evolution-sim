package bus

type envelope struct {
	name  EventName
	event IEvent
}
