package service

import "samurenkoroma/services/services/storehouse"

type VendorService struct {
	Repo *storehouse.VendorRepository
}

func NewVendorService(repo *storehouse.VendorRepository) *VendorService {
	return &VendorService{
		Repo: repo,
	}
}
func (s *VendorService) Add(dto *storehouse.CreateVendorRequest) (storehouse.CreateVendorResponse, error) {
	vendor := &storehouse.Vendor{
		Name: dto.Name,
		Url:  dto.Url,
	}
	err := s.Repo.Crud.Save(vendor)
	if err != nil {
		return storehouse.CreateVendorResponse{}, err
	}
	return storehouse.CreateVendorResponse{
		ID:   vendor.ID,
		Name: vendor.Name,
		Url:  vendor.Url,
	}, nil

}

func (s *VendorService) List() ([]storehouse.CreateVendorResponse, error) {
	vendors, err := s.Repo.Crud.List("")
	if err != nil {
		return nil, err
	}

	var data []storehouse.CreateVendorResponse
	for _, vendor := range vendors {
		data = append(data, storehouse.CreateVendorResponse{
			ID:   vendor.ID,
			Name: vendor.Name,
			Url:  vendor.Url,
		})
	}
	return data, nil
}
