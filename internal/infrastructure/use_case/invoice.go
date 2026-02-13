package use_case

import (
	"samurenkoroma/services/internal/domain/accountant"
	"samurenkoroma/services/internal/infrastructure/repo"
)

type AccountantService struct {
	suppliers repo.CRUDRepository[accountant.Supplier]
}
