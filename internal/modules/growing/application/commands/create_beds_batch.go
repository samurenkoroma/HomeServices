package commands

//
//import (
//	"context"
//	"errors"
//	"fmt"
//	"log"
//	"samurenkoroma/services/internal/application/command"
//	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/postgres"
//
//	"samurenkoroma/services/internal/core/domain/repository"
//	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"
//)
//
//// CreateBedsBatchCommand — команда массового создания грядок
//type CreateBedsBatchCommand struct {
//	GreenhouseID string     `json:"greenhouse_id" validate:"required"`
//	SeasonID     string     `json:"season_id" validate:"required"`
//	Beds         []BedInput `json:"beds" validate:"required,min=1"`
//}
//
//// BedInput — входные данные для одной грядки
//type BedInput struct {
//	Name       string  `json:"name" validate:"required"`
//	Width      float64 `json:"width" validate:"required,gt=0"`
//	Length     float64 `json:"length" validate:"required,gt=0"`
//	PositionX  float64 `json:"position_x" validate:"required,min=0,max=100"`
//	PositionY  float64 `json:"position_y" validate:"required,min=0,max=100"`
//	CropPlanID *string `json:"crop_plan_id"`
//}
//
//// CreateBedsBatchHandler — обработчик массового создания грядок
//type createBedsBatchHandler struct {
//	uowFactory repository.Factory
//}
//
//func (h *createBedsBatchHandler) Name() string {
//	return "CreateBedsBatch"
//}
//
//func NewCreateBedsBatchHandler(uowFactory repository.Factory) command.Handler {
//	return &createBedsBatchHandler{
//		uowFactory: uowFactory,
//	}
//}
//
//func (h *createBedsBatchHandler) Handle(ctx context.Context, command any) error {
//	cmd, ok := command.(*CreateBedsBatchCommand)
//	if !ok {
//		return errors.New("invalid command")
//	}
//	log.Printf("=== CREATE BEDS BATCH === GreenhouseID: %s, Beds count: %d", cmd.GreenhouseID, len(cmd.Beds))
//
//	uow, err := h.uowFactory.Begin(ctx)
//	if err != nil {
//		return fmt.Errorf("failed to begin transaction: %w", err)
//	}
//
//	var createdBeds []*cultivationarea.Bed
//
//	err = uow.Execute(ctx, postgres.NewGrowingProvider, func(provider repository.RepositoryProvider) error {
//		growingProvider, ok := provider.(*postgres.GrowingProvider)
//		if !ok {
//			return fmt.Errorf("invalid provider type")
//		}
//
//		// 1. Проверяем существование теплицы
//		greenhouse, err := growingProvider.CultivationAreas().FindByID(ctx, cmd.GreenhouseID)
//		if err != nil {
//			return fmt.Errorf("failed to find greenhouse: %w", err)
//		}
//		if greenhouse == nil {
//			return cultivationarea.ErrAreaNotFound
//		}
//
//		// 2. Проверяем, что это теплица
//		if greenhouse.GetType() != cultivationarea.AreaTypeGreenhouse {
//			return fmt.Errorf("area is not a greenhouse: %s", greenhouse.GetType())
//		}
//
//		// 3. Создаём грядки
//		for _, bedInput := range cmd.Beds {
//
//			// Создаём грядку с точкой
//			bed := cultivationarea.NewBed(
//				cmd.GreenhouseID,
//				bedInput.Name,
//				greenhouse.GetGeometry(),
//			)
//
//			// Устанавливаем атрибуты (размеры и позиция в JSONB)
//			bed.SetAttributes(bedInput.Width, bedInput.Length, bedInput.PositionX, bedInput.PositionY)
//			bed.SetFarmRefID(greenhouse.GetFarmRefID())
//
//			// Настраиваем на сезон
//			config := cultivationarea.AreaConfig{
//				Name:       bedInput.Name,
//				Geometry:   greenhouse.GetGeometry(),
//				CropPlanID: bedInput.CropPlanID,
//				Metadata: map[string]interface{}{
//					"width":      bedInput.Width,
//					"length":     bedInput.Length,
//					"position_x": bedInput.PositionX,
//					"position_y": bedInput.PositionY,
//				},
//			}
//
//			if err := bed.ConfigureForSeason(cmd.SeasonID, config); err != nil {
//				return fmt.Errorf("failed to configure bed %s: %w", bedInput.Name, err)
//			}
//
//			// Сохраняем грядку
//			if err := growingProvider.CultivationAreas().Save(ctx, bed); err != nil {
//				return fmt.Errorf("failed to save bed: %w", err)
//			}
//
//			// Сохраняем конфигурацию на сезон
//			seasonConfig, err := bed.GetSeasonConfig(cmd.SeasonID)
//			if err != nil {
//				return err
//			}
//
//			if err := growingProvider.CultivationAreas().SaveSeasonConfig(ctx, bed.GetID(), *seasonConfig); err != nil {
//				return fmt.Errorf("failed to save season config for bed: %w", err)
//			}
//
//			// Добавляем грядку в теплицу
//			if greenhouseArea, ok := greenhouse.(*cultivationarea.GreenhouseArea); ok {
//				if err := greenhouseArea.AddBed(cmd.SeasonID, bed.GetID()); err != nil {
//					return err
//				}
//			}
//
//			createdBeds = append(createdBeds, bed)
//			uow.RegisterAggregate(bed)
//		}
//
//		uow.RegisterAggregate(greenhouse)
//		return nil
//	})
//
//	if err != nil {
//		return err
//	}
//
//	log.Printf("Successfully created %d beds", len(createdBeds))
//	return nil
//}
