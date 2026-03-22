package aggregate

import (
	"samurenkoroma/services/internal/core/domain"
	"samurenkoroma/services/internal/core/domain/event"
	"time"
)

type Aggregate interface {
	AddEvent(e event.DomainEvent)
	PullEvents() []event.DomainEvent
}

type BaseAggregate struct {
	events []event.DomainEvent
}

func (a *BaseAggregate) AddEvent(e event.DomainEvent) {
	a.events = append(a.events, e)
}

func (a *BaseAggregate) PullEvents() []event.DomainEvent {
	ev := a.events
	a.events = nil
	return ev
}

type EntityDates struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func Rehydrate[T any](id T, isActive bool, dates EntityDates) *Entity[T] {
	return &Entity[T]{
		BaseAggregate: BaseAggregate{},
		Id:            id,
		IsActive:      isActive,
		EntityDates:   dates,
	}
}

type Entity[T any] struct {
	BaseAggregate
	Id       T
	IsActive bool
	EntityDates
}

func (e *Entity[T]) Deactivate() error {
	if !e.IsActive {
		return domain.ForbiddenError("entity is already inactive")
	}

	e.IsActive = false
	e.Update()

	return nil
}

// Activate активирует тип культуры
func (e *Entity[T]) Activate() error {
	if e.IsActive {
		return domain.ForbiddenError("entity is already active")
	}

	e.IsActive = true
	e.Update()

	return nil
}

func (e *Entity[T]) Update() {
	e.UpdatedAt = time.Now()
}

func NewEntity[T any](id T) Entity[T] {
	now := time.Now()
	return Entity[T]{
		BaseAggregate: BaseAggregate{
			events: make([]event.DomainEvent, 0),
		},
		Id: id,
		EntityDates: EntityDates{
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: nil,
		},
	}
}
