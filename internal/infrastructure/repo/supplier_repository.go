package repo

import (
	"samurenkoroma/services/internal/domain/accountant"
	"samurenkoroma/services/pkg/db"
)

func NewSupplierRepo(database *db.Db) CRUDRepository[accountant.Supplier] {
	return CRUDRepository[accountant.Supplier]{Database: database}
}
