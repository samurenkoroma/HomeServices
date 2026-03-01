package events

import (
	"context"
	"fmt"
	"log"
	"samurenkoroma/services/internal/common/application/uow"
	"samurenkoroma/services/internal/common/domain"
	"samurenkoroma/services/internal/growing/domain/facility"
)

func OnFacilityCreated(
	ctx context.Context,
	event domain.DomainEvent,
) error {
	UOW, ok := uow.FromContext(ctx)
	if !ok {
		return nil
	}
	fmt.Println(UOW)
	e, ok := event.(facility.FacilityCreatedEvent)
	if !ok {
		return nil
	}

	log.Println("Facility created:", e.FacilityID)

	return nil
}
