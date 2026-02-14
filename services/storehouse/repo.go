package storehouse

import (
	"samurenkoroma/services/internal/infrastructure/repo"
	"samurenkoroma/services/pkg/db"
)

type SeedsRepository struct {
	Crud *repo.CRUDRepository[Seed]
}

func NewSeedsRepository(database *db.Db) *SeedsRepository {
	return &SeedsRepository{
		Crud: repo.NewCrudRepo[Seed](database),
	}
}

type VendorRepository struct {
	Crud *repo.CRUDRepository[Vendor]
}

func NewVendorRepository(database *db.Db) *VendorRepository {
	return &VendorRepository{
		Crud: repo.NewCrudRepo[Vendor](database),
	}
}
