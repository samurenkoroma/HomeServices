package handlers

import (
	"context"
	"fmt"
	"log"
	"samurenkoroma/services/internal/core/domain/event"
	uow "samurenkoroma/services/internal/core/port/repository"
	"samurenkoroma/services/internal/modules/farm/field/domain"
)

func OnFieldCreated(ctx context.Context, event event.DomainEvent) error {
	UOW, ok := uow.FromContext(ctx)
	if !ok {
		return nil
	}
	fmt.Println(UOW)
	e, ok := event.(domain.FieldCreated)
	if !ok {
		return nil
	}

	log.Println("Field created:", e.FieldID)

	return nil
}
