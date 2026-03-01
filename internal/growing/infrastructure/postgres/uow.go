package postgres

import (
	"database/sql"
	"samurenkoroma/services/internal/application/shared"
	domainShared "samurenkoroma/services/internal/domain/shared"
	"samurenkoroma/services/internal/growing/application"
	"sync"
)

type PgUow struct {
	tx       *sql.Tx
	eventBus shared.EventBus
	landRepo application.GrowingFacilitiesRepository
	cropRepo application.CropPlanRepository

	aggregates []domainShared.EventAwareAggregate

	closed bool
	mu     sync.Mutex
}

func newPgUnitOfWork(tx *sql.Tx, eventBus shared.EventBus) *PgUow {
	uow := &PgUow{
		tx:       tx,
		eventBus: eventBus,
	}

	uow.landRepo = NewLandUnitRepository(tx, uow)

	return uow
}

func (u *PgUow) GrowingFacilities() application.GrowingFacilitiesRepository {
	return u.landRepo
}

func (u *PgUow) CropPlans() application.CropPlanRepository {
	return u.cropRepo
}

func (u *PgUow) Commit() error {
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

func (u *PgUow) Rollback() error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if u.closed {
		return nil
	}

	u.closed = true
	return u.tx.Rollback()
}

func (u *PgUow) dispatchEvents() error {
	var allEvents []domainShared.DomainEvent

	for _, agg := range u.aggregates {
		allEvents = append(allEvents, agg.PullEvents()...)
	}

	if len(allEvents) == 0 {
		return nil
	}

	return u.eventBus.Publish(allEvents)
}

func (u *PgUow) registerAggregate(a domainShared.EventAwareAggregate) {
	u.aggregates = append(u.aggregates, a)
}
