package uow

import (
	"context"
	"database/sql"
	"samurenkoroma/services/internal/growing/application"
	"samurenkoroma/services/internal/growing/infrastructure/postgres"
	"samurenkoroma/services/pkg/event"
	"sync"
)

type UnitOfWork interface {
	GrowingFacilities() application.GrowingFacilitiesRepository
	CropPlans() application.CropPlanRepository
	Commit() error
	Rollback() error
}

type UnitOfWorkFactory interface {
	New(ctx context.Context) (Uow, error)
}

type Uow struct {
	tx       *sql.Tx
	eventBus event.EventBus

	facilitiesRepo application.GrowingFacilitiesRepository
	cropRepo       application.CropPlanRepository

	aggregates []domainShared.EventAwareAggregate

	closed bool
	mu     sync.Mutex
}

func NewUnitOfWork(tx *sql.Tx, eventBus shared.EventBus) *Uow {
	uow := &Uow{
		tx:       tx,
		eventBus: eventBus,
	}

	uow.facilitiesRepo = postgres.NewGrowingFacilitiesRepository(tx)

	return uow
}

func (u *Uow) GrowingFacilities() application.GrowingFacilitiesRepository {
	return u.facilitiesRepo
}

func (u *Uow) CropPlans() application.CropPlanRepository {
	return u.cropRepo
}

func (u *Uow) Commit() error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if u.closed {
		return nil
	}

	if err := u.tx.Commit(); err != nil {
		return err
	}

	u.closed = true

	return u.dispatchEvents()
}

func (u *Uow) Rollback() error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if u.closed {
		return nil
	}

	u.closed = true
	return u.tx.Rollback()
}

func (u *Uow) dispatchEvents() error {
	var allEvents []domainShared.DomainEvent

	for _, agg := range u.aggregates {
		allEvents = append(allEvents, agg.PullEvents()...)
	}

	if len(allEvents) == 0 {
		return nil
	}

	return u.eventBus.Publish(allEvents)
}

func (u *Uow) RegisterAggregate(a domainShared.EventAwareAggregate) {
	u.aggregates = append(u.aggregates, a)
}
