package sql

import (
	"context"
	"database/sql"
	"errors"
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/domain/event"
	"samurenkoroma/services/internal/core/port/messaging"
	"samurenkoroma/services/internal/core/port/repository"
	crop "samurenkoroma/services/internal/modules/crop/domain"
	"samurenkoroma/services/internal/modules/crop/infrastructure/persistence"
	domain2 "samurenkoroma/services/internal/modules/farm/field/domain"
	"samurenkoroma/services/internal/modules/growing/domain"
	"samurenkoroma/services/internal/modules/growing/infrastructure/postgres"
	"sync"
)

var (
	ErrAlreadyCommitted  = errors.New("unit of work already committed")
	ErrAlreadyRolledBack = errors.New("unit of work already rolled back")
)

func newSQLUnitOfWork(
	ctx context.Context,
	tx *sql.Tx,
	bus messaging.EventBus,
) repository.UnitOfWork {
	return &sqlUnitOfWork{
		ctx:            ctx,
		tx:             tx,
		bus:            bus,
		facilitiesRepo: postgres.NewGrowingFacilitiesRepository(tx),
		cropRepo:       persistence.NewCropRepo(tx),
	}
}

type sqlUnitOfWork struct {
	ctx context.Context
	tx  *sql.Tx

	bus messaging.EventBus

	mu         sync.Mutex
	committed  bool
	rolledBack bool

	aggregates []aggregate.Aggregate

	facilitiesRepo domain2.FieldRepository
	cropRepo       crop.CropPlanRepository
}

func (u *sqlUnitOfWork) CropCycles() domain.CropCycleRepository {
	//TODO implement me
	panic("implement me")
}

func (u *sqlUnitOfWork) CropTemplates() domain.CropTemplateRepository {
	//TODO implement me
	panic("implement me")
}

func (u *sqlUnitOfWork) GrowingFacilities() domain2.FieldRepository {
	return u.facilitiesRepo
}
func (u *sqlUnitOfWork) CropPlans() crop.CropPlanRepository { return u.cropRepo }

func (u *sqlUnitOfWork) RegisterAggregate(agg aggregate.Aggregate) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.aggregates = append(u.aggregates, agg)
}

func (u *sqlUnitOfWork) Commit() error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if u.committed {
		return ErrAlreadyCommitted
	}
	if u.rolledBack {
		return ErrAlreadyRolledBack
	}

	if err := u.tx.Commit(); err != nil {
		return err
	}

	return u.dispatchEvents()
}

func (u *sqlUnitOfWork) Rollback() error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if u.committed {
		return ErrAlreadyCommitted
	}
	if u.rolledBack {
		return ErrAlreadyRolledBack
	}
	u.rolledBack = true
	return u.tx.Rollback()
}

func (u *sqlUnitOfWork) EventBus() messaging.EventBus {
	return u.bus
}
func (u *sqlUnitOfWork) dispatchEvents() error {
	var allEvents []event.DomainEvent

	for _, agg := range u.aggregates {
		allEvents = append(allEvents, agg.PullEvents()...)
	}

	if len(allEvents) == 0 {
		return nil
	}

	return u.bus.Publish(repository.WithUnitOfWork(u.ctx, u), allEvents)
}
