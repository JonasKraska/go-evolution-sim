package bus

type Callback func(event IEvent)

type Handler struct {
	callback   Callback
	eventNames []EventName
	priority   int8
}

func (handler Handler) Callback() Callback {
	return handler.callback
}

func (handler Handler) EventNames() []EventName {
	return handler.eventNames
}

func (handler Handler) Priority() int8 {
	return handler.priority
}
