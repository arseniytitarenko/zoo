package dispatcher

import "zoo/internal/domain"

type EventDispatcher struct {
	handlers map[domain.EventName][]func(domain.Event)
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[domain.EventName][]func(domain.Event)),
	}
}

func (d *EventDispatcher) Register(eventName domain.EventName, handler func(domain.Event)) {
	d.handlers[eventName] = append(d.handlers[eventName], handler)
}

func (d *EventDispatcher) Dispatch(event domain.Event) {
	for _, h := range d.handlers[event.EventName()] {
		h(event)
	}
}
