package commands

import (
	"context"
	"samurenkoroma/services/internal/core/eventbus"
	"samurenkoroma/services/internal/core/port/repository"
	crop "samurenkoroma/services/internal/modules/crop/domain"
	"samurenkoroma/services/internal/modules/farm"
	domain3 "samurenkoroma/services/internal/modules/farm/field/domain"
	"samurenkoroma/services/internal/modules/farm/valueobject"
	"samurenkoroma/services/internal/modules/growing/domain"
	"testing"

	"github.com/google/uuid"
)

// Мок для UnitOfWork
type mockUOW struct {
	committed     bool
	rolledBack    bool
	savedFacility *domain3.Field
}

func (m *mockUOW) CropCycles() domain.CropCycleRepository {
	//TODO implement me
	panic("implement me")
}

func (m *mockUOW) CropTemplates() domain.CropTemplateRepository {
	//TODO implement me
	panic("implement me")
}

func (m *mockUOW) CropPlans() crop.CropPlanRepository {
	//TODO implement me
	panic("implement me")
}

func (m *mockUOW) GrowingFacilities() domain3.FieldRepository {
	return &mockFacilityRepo{uow: m}
}

// func (m *mockUOW) CropPlans() cropplan.Repository         { return nil }
func (m *mockUOW) RegisterAggregate(agg eventbus.Aggregate) {}
func (m *mockUOW) Commit() error                            { m.committed = true; return nil }
func (m *mockUOW) Rollback() error                          { m.rolledBack = true; return nil }
func (m *mockUOW) EventBus() eventbus.EventBus              { return nil }

// Мок для репозитория
type mockFacilityRepo struct {
	uow *mockUOW
}

func (r *mockFacilityRepo) Get(id farm.FacilityID) (*domain3.Field, error) {
	// Создаем тестовую теплицу
	dim, _ := valueobject.NewDimension(100, 50)
	return domain3.NewGreenhouseFacility(id, "Test Greenhouse", dim), nil
}

func (r *mockFacilityRepo) Save(unit *domain3.Field) error {
	r.uow.savedFacility = unit
	return nil
}

// Мок для фабрики UOW
type mockUOWFactory struct{}

func (f *mockUOWFactory) Begin(ctx context.Context) (repository.UnitOfWork, error) {
	return &mockUOW{}, nil
}

func TestAddBedHandler(t *testing.T) {
	// Создаем хендлер с мок-фабрикой
	handler := AddBedHandler{
		UowFactory: &mockUOWFactory{},
	}

	// Тест 1: Добавление грядки в теплицу (без block_id)
	t.Run("add bed to greenhouse", func(t *testing.T) {
		facilityID := uuid.New().String()

		cmd := AddBedCmd{
			FacilityID: facilityID,
			Name:       "Грядка A1",
			Length:     10,
			Width:      1.2,
		}

		err := handler.Handle(context.Background(), cmd)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	// Тест 2: Добавление грядки в секцию (с block_id)
	t.Run("add bed to block", func(t *testing.T) {
		facilityID := uuid.New().String()
		blockID := uuid.New().String()

		cmd := AddBedCmd{
			FacilityID: facilityID,
			BlockID:    &blockID,
			Name:       "Грядка в секции",
			Length:     5,
			Width:      1,
		}

		err := handler.Handle(context.Background(), cmd)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	// Тест 3: Ошибка валидации
	t.Run("validation error", func(t *testing.T) {
		_, err := DecodeAddBed([]byte(`{"name":"Грядка","length":10,"width":1}`))
		if err == nil {
			t.Error("Expected validation error, got nil")
		}
	})
}
