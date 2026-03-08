package handlers

import (
	"context"
	"fmt"
	"log"
	"samurenkoroma/services/internal/core/domain/event"
	uow "samurenkoroma/services/internal/core/port/repository"
	"samurenkoroma/services/internal/modules/farm/domain/field"
)

func OnGreenhouseCreated(ctx context.Context, event event.DomainEvent) error {
	UOW, ok := uow.FromContext(ctx)
	if !ok {
		return nil
	}
	fmt.Println(UOW)
	e, ok := event.(field.FieldCreated)
	if !ok {
		return nil
	}

	log.Println("Field created:", e.FieldID)

	return nil
}
