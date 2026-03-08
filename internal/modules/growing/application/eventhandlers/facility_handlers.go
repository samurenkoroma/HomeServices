package eventhandlers

import (
	"context"
	"fmt"
	"log"
	"samurenkoroma/services/internal/core/domain/event"
	"samurenkoroma/services/internal/core/port/repository"
	"samurenkoroma/services/internal/modules/growing/domain/facility"
)

func OnFacilityCreated(ctx context.Context, event event.DomainEvent) error {
	UOW, ok := repository.FromContext(ctx)
	if !ok {
		return nil
	}
	fmt.Println(UOW)
	e, ok := event.(facility.FacilityCreated)
	if !ok {
		return nil
	}

	log.Println("Facility created:", e.FacilityID)

	return nil
}
