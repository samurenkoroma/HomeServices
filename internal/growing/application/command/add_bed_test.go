package command

import (
	"context"
	"samurenkoroma/services/internal/common/application/uow"
	"samurenkoroma/services/internal/common/domain"
	crop "samurenkoroma/services/internal/crop/domain"
	growlingDomain "samurenkoroma/services/internal/growing/domain"
	"samurenkoroma/services/internal/growing/domain/facility"
	"samurenkoroma/services/internal/growing/domain/valueobject"
	"testing"

	"github.com/google/uuid"
)

// Мок для UnitOfWork
type mockUOW struct {
	committed     bool
	rolledBack    bool
	savedFacility *facility.GrowingFacility
}

func (m *mockUOW) CropCycles() growlingDomain.CropCycleRepository {
	//TODO implement me
	panic("implement me")
}

func (m *mockUOW) CropTemplates() growlingDomain.CropTemplateRepository {
	//TODO implement me
	panic("implement me")
}

func (m *mockUOW) CropPlans() crop.CropPlanRepository {
	//TODO implement me
	panic("implement me")
}

func (m *mockUOW) GrowingFacilities() growlingDomain.GrowingFacilitiesRepository {
	return &mockFacilityRepo{uow: m}
}

// func (m *mockUOW) CropPlans() cropplan.Repository         { return nil }
func (m *mockUOW) RegisterAggregate(agg domain.Aggregate) {}
func (m *mockUOW) Commit() error                          { m.committed = true; return nil }
func (m *mockUOW) Rollback() error                        { m.rolledBack = true; return nil }
func (m *mockUOW) EventBus() domain.EventBus              { return nil }

// Мок для репозитория
type mockFacilityRepo struct {
	uow *mockUOW
}

func (r *mockFacilityRepo) Get(id facility.FacilityID) (*facility.GrowingFacility, error) {
	// Создаем тестовую теплицу
	dim, _ := valueobject.NewDimension(100, 50)
	return facility.NewGreenhouseFacility(id, "Test Greenhouse", dim), nil
}

func (r *mockFacilityRepo) Save(unit *facility.GrowingFacility) error {
	r.uow.savedFacility = unit
	return nil
}

// Мок для фабрики UOW
type mockUOWFactory struct{}

func (f *mockUOWFactory) Begin(ctx context.Context) (uow.UnitOfWork, error) {
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
