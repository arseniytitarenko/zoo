package out

import "zoo/internal/domain"

type IEventDispatcher interface {
	Register(eventName domain.EventName, handler func(domain.Event))
	Dispatch(event domain.Event)
}
