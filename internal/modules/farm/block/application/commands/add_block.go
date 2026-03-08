package commands

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"samurenkoroma/services/internal/core/port/repository"
	"samurenkoroma/services/internal/modules/farm"
	"samurenkoroma/services/internal/modules/farm/valueobject"

	"github.com/google/uuid"
)

type AddBlockHandler struct {
	UowFactory repository.Factory
}

func NewAddBlockHandler(uowFactory repository.Factory) *AddBlockHandler {
	return &AddBlockHandler{UowFactory: uowFactory}
}

type AddBlockCmd struct {
	FacilityID string  `json:"facility_id"` // ID теплицы или поля
	Name       string  `json:"name"`        // Название грядки
	Length     float64 `json:"length"`      // Длина в метрах
	Width      float64 `json:"width"`       // Ширина в метрах
}

func DecodeAddBlock(data []byte) (any, error) {

	var cmd AddBlockCmd
	if err := json.Unmarshal(data, &cmd); err != nil {
		return nil, fmt.Errorf("failed to decode AddBed command: %w", err)
	}

	// Базовая валидация
	if cmd.FacilityID == "" {
		return nil, errors.New("facility_id is required")
	}
	if cmd.Name == "" {
		return nil, errors.New("name is required")
	}
	if cmd.Length <= 0 {
		return nil, errors.New("length must be positive")
	}
	if cmd.Width <= 0 {
		return nil, errors.New("width must be positive")
	}

	return cmd, nil
}

// Handle выполняет команду
func (h *AddBlockHandler) Handle(ctx context.Context, cmd any) error {
	// Приводим команду к нужному типу
	c, ok := cmd.(AddBlockCmd)
	if !ok {
		return errors.New("invalid command type: expected AddBedCmd")
	}

	// Начинаем транзакцию
	uowObj, err := h.UowFactory.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin unit of work: %w", err)
	}
	defer uowObj.Rollback() // Откатим, если не будет Commit

	// Получаем репозиторий для работы с facility
	repo := uowObj.GrowingFacilities()

	// Загружаем существующее сооружение (теплицу или поле)
	facility_unit, err := repo.Get(farm.FacilityID(c.FacilityID))
	if err != nil {
		return fmt.Errorf("failed to get facility %s: %w", c.FacilityID, err)
	}

	// Создаем dimension из размеров
	dim, err := valueobject.NewDimension(c.Length, c.Width)
	if err != nil {
		return fmt.Errorf("invalid dimensions: %w", err)
	}

	// Генерируем ID для новой грядки
	blockID := farm.GrowingAreaID(uuid.New().String())

	// Добавляем напрямую (для теплицы)
	err = facility_unit.AddBlock(
		blockID,
		c.Name,
		dim,
	)
	if err != nil {
		return fmt.Errorf("failed to add bed directly: %w", err)
	}

	// Сохраняем изменения в facility
	if err := repo.Save(facility_unit); err != nil {
		return fmt.Errorf("failed to save facility: %w", err)
	}

	// Регистрируем агрегат для сбора событий
	uowObj.RegisterAggregate(facility_unit)

	// Коммитим транзакцию (автоматически опубликует события)
	return uowObj.Commit()
}
