package eventhandlers

import (
	"context"
	"fmt"
	"log"
	"samurenkoroma/services/internal/common/application/uow"
	"samurenkoroma/services/internal/common/domain"
	"samurenkoroma/services/internal/shared/events"
)

func OnFacilityCreated(ctx context.Context, event domain.DomainEvent) error {
	UOW, ok := uow.FromContext(ctx)
	if !ok {
		return nil
	}
	fmt.Println(UOW)
	e, ok := event.(events.FacilityCreated)
	if !ok {
		return nil
	}

	log.Println("Facility created:", e.FacilityID)

	return nil
}
