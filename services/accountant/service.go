package accountant

import (
	"log"
	"samurenkoroma/services/pkg/di"
	"samurenkoroma/services/services/accountant/entity"
)

type Service struct {
	repo di.CRUDRepository[entity.Supplier]
}

func NewSupplierService(repo di.CRUDRepository[entity.Supplier]) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(dto *CreateSupplierRequest) (CreateSupplierResponse, error) {
	supplier, err := entity.NewSupplier(dto.Name, dto.Site)
	if err != nil {
		log.Println(err)
		return CreateSupplierResponse{}, err
	}
	if err := s.repo.Save(supplier); err != nil {
		log.Println(err)
		return CreateSupplierResponse{}, err
	}
	return CreateSupplierResponse{
		ID:     supplier.ID,
		Name:   supplier.Name,
		Rating: supplier.Rating,
	}, nil
}

func (s *Service) Get(id uint) (CreateSupplierResponse, error) {
	supplier, err := s.repo.Get(id)
	if err != nil {
		return CreateSupplierResponse{}, err
	}
	return CreateSupplierResponse{
		ID:     supplier.ID,
		Name:   supplier.Name,
		Rating: supplier.Rating,
	}, nil

}

func (s *Service) List() ([]CreateSupplierResponse, error) {

	suppliers, err := s.repo.List("")
	var data []CreateSupplierResponse

	if err != nil {
		return nil, err
	}
	for _, supplier := range suppliers {
		data = append(data, CreateSupplierResponse{
			ID:   supplier.ID,
			Name: supplier.Name,
		})
	}
	return data, nil
}

func (s *Service) Update(id uint, supplier *entity.Supplier) (bool, error) {
	return s.repo.Update(id, supplier)
}
func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}
