package bus

import (
	"errors"
	"sort"
	"sync"
)

type bus struct {
	waitGroup sync.WaitGroup
	queue     chan envelope
	handlers  map[EventName][]*Handler
}

func New(buffersize ...int) *bus {
	bus := &bus{
		waitGroup: sync.WaitGroup{},
		queue:     make(chan envelope, buffersize[0]),
		handlers:  make(map[EventName][]*Handler),
	}

	go func() {
		for envelope := range bus.queue {
			for _, handler := range bus.handlers[envelope.name] {
				handler.callback(envelope.event)
			}

			bus.waitGroup.Done()
		}
	}()

	return bus
}

func (bus *bus) Wait() {
	bus.waitGroup.Wait()
}

func (bus *bus) Close() {
	close(bus.queue)
}

func (bus *bus) WaitAndClose() {
	bus.Wait()
	bus.Close()
}

func (bus *bus) AddHandler(callback Callback, priority int8, eventNames ...EventName) (*Handler, error) {

	if callback == nil {
		return nil, errors.New("handler to be registered has no callback")
	}

	if len(eventNames) == 0 {
		return nil, errors.New("handler to be registered has no event names")
	}

	handler := Handler{
		callback:   callback,
		eventNames: eventNames,
		priority:   priority,
	}

	for _, eventName := range eventNames {
		eventHandlers := bus.handlers[eventName]
		eventHandlers = append(eventHandlers, &handler)

		sort.SliceStable(eventHandlers, func(i, j int) bool {
			return eventHandlers[i].priority > eventHandlers[j].priority
		})

		bus.handlers[eventName] = eventHandlers
	}

	return &handler, nil
}

func (bus *bus) RemoveListener(handler *Handler) error {
	for _, eventName := range handler.eventNames {

		index := 0
		for _, registeredHandler := range bus.handlers[eventName] {
			if handler == registeredHandler {
				continue
			}

			bus.handlers[eventName][index] = registeredHandler
			index++
		}

		if index == len(bus.handlers[eventName]) {
			return errors.New("handler to be removed was not registered")
		}

		// clean up to prevent memory leaks
		for i := index; i < len(bus.handlers[eventName]); i++ {
			bus.handlers[eventName][i] = nil
		}

		// resize resize slice to remove nil pointers
		bus.handlers[eventName] = bus.handlers[eventName][:index]
	}

	return nil
}

func (bus *bus) DispatchAsync(name EventName, event IEvent) {
	bus.doDispatch(name, event)
}

func (bus *bus) Dispatch(name EventName, event IEvent) {
	bus.doDispatch(name, event)
	bus.Wait()
}

func (bus *bus) doDispatch(name EventName, event IEvent) {
	bus.waitGroup.Add(1)
	bus.queue <- envelope{
		name:  name,
		event: event,
	}
}
