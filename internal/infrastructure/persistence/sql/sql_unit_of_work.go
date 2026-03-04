package sql

import (
	"context"
	"database/sql"
	"errors"
	"samurenkoroma/services/internal/common/application/uow"
	"samurenkoroma/services/internal/common/domain"
	crop "samurenkoroma/services/internal/crop/domain"
	"samurenkoroma/services/internal/crop/infrastructure/persistence"
	growing "samurenkoroma/services/internal/growing/domain"
	"samurenkoroma/services/internal/growing/infrastructure/postgres"
	"sync"
)

var (
	ErrAlreadyCommitted  = errors.New("unit of work already committed")
	ErrAlreadyRolledBack = errors.New("unit of work already rolled back")
)

func newSQLUnitOfWork(
	ctx context.Context,
	tx *sql.Tx,
	bus domain.EventBus,
) uow.UnitOfWork {
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

	bus domain.EventBus

	mu         sync.Mutex
	committed  bool
	rolledBack bool

	aggregates []domain.Aggregate

	facilitiesRepo growing.GrowingFacilitiesRepository
	cropRepo       crop.CropPlanRepository
}

func (u *sqlUnitOfWork) CropCycles() growing.CropCycleRepository {
	//TODO implement me
	panic("implement me")
}

func (u *sqlUnitOfWork) CropTemplates() growing.CropTemplateRepository {
	//TODO implement me
	panic("implement me")
}

func (u *sqlUnitOfWork) GrowingFacilities() growing.GrowingFacilitiesRepository {
	return u.facilitiesRepo
}
func (u *sqlUnitOfWork) CropPlans() crop.CropPlanRepository { return u.cropRepo }

func (u *sqlUnitOfWork) RegisterAggregate(agg domain.Aggregate) {
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

func (u *sqlUnitOfWork) EventBus() domain.EventBus {
	return u.bus
}
func (u *sqlUnitOfWork) dispatchEvents() error {
	var allEvents []domain.DomainEvent

	for _, agg := range u.aggregates {
		allEvents = append(allEvents, agg.PullEvents()...)
	}

	if len(allEvents) == 0 {
		return nil
	}

	return u.bus.Publish(uow.WithUnitOfWork(u.ctx, u), allEvents)
}
