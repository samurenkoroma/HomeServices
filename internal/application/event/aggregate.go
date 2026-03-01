package event

type EventAwareAggregate interface {
	PullEvents() []DomainEvent
}

type BaseAggregate struct {
	events []DomainEvent
}

func (a *BaseAggregate) AddEvent(e DomainEvent) {
	a.events = append(a.events, e)
}

func (a *BaseAggregate) PullEvents() []DomainEvent {
	ev := a.events
	a.events = nil
	return ev
}
