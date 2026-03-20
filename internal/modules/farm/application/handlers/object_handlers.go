package handlers

import (
	"context"
	"fmt"
	"log"
	"samurenkoroma/services/internal/core/domain/event"
	uow "samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/farm/domain/physicalobject"
)

func OnFieldCreated(ctx context.Context, event event.DomainEvent) error {
	UOW, ok := uow.FromContext(ctx)
	if !ok {
		return nil
	}
	fmt.Println(UOW)
	e, ok := event.(physicalobject.FieldCreated)
	if !ok {
		return nil
	}

	log.Println("Field created:", e.ID)

	return nil
}

func OnGreenhouseCreated(ctx context.Context, event event.DomainEvent) error {
	UOW, ok := uow.FromContext(ctx)
	if !ok {
		return nil
	}
	fmt.Println(UOW)
	e, ok := event.(physicalobject.GreenhouseCreated)
	if !ok {
		return nil
	}

	log.Println("Greenhouse created:", e.ID)

	return nil
}
