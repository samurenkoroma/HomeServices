package domain

type EventHandler interface {
	Handle(event DomainEvent) error
}

type EventBus interface {
	Publish(events []DomainEvent) error
	Register(eventName string, handler EventHandler)
}
