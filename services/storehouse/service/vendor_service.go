package service

import (
	"context"
	"samurenkoroma/services/internal/infrastructure/db_table"
	"samurenkoroma/services/pkg/db"
	"samurenkoroma/services/services/storehouse"

	"gorm.io/gorm"
)

type VendorService struct {
	database *db.Db
}

func NewVendorService(db *db.Db) *VendorService {
	return &VendorService{
		database: db,
	}
}
func (s *VendorService) Add(dto *storehouse.CreateVendorRequest) (storehouse.CreateVendorResponse, error) {
	vendor := &storehouse.Vendor{
		Name: dto.Name,
		URL:  dto.Url,
	}

	s.database.DB.Create(&vendor)

	return storehouse.CreateVendorResponse{
		ID:   vendor.ID,
		Name: vendor.Name,
		Url:  vendor.URL,
	}, nil

}

func (s *VendorService) List() ([]storehouse.CreateVendorResponse, error) {
	vendors, err := gorm.G[db_table.Vendor](s.database.DB).Find(context.Background())
	if err != nil {
		return nil, err
	}
	var data []storehouse.CreateVendorResponse
	for _, v := range vendors {
		data = append(data, storehouse.CreateVendorResponse{
			ID:   v.ID,
			Name: v.Name,
			Url:  v.Url,
		})
	}
	return data, nil
}
