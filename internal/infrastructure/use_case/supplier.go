package use_case

import (
	"log"
	"samurenkoroma/services/internal/domain/accountant"
	"samurenkoroma/services/internal/infrastructure/payload"

	"github.com/google/uuid"
)

type Repository interface {
	Save(entity *accountant.Supplier) error
	Get(id string) (*accountant.Supplier, error)
	List(filter string) ([]*accountant.Supplier, error)
	Update(entity *accountant.Supplier) (bool, error)
	Delete(id string) error
}

type Service struct {
	repo Repository
}

func NewSupplierService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(dto *payload.CreateSupplierRequest) (payload.CreateSupplierResponse, error) {
	ID := uuid.NewString()
	supplier, err := accountant.NewSupplier(ID, dto.Name)
	if err != nil {
		log.Println(err)
		return payload.CreateSupplierResponse{}, err
	}
	if err := s.repo.Save(supplier); err != nil {
		log.Println(err)
		return payload.CreateSupplierResponse{}, err
	}
	return payload.CreateSupplierResponse{
		ID:     ID,
		Name:   supplier.Name,
		Rating: supplier.Rating,
	}, nil
}

func (s *Service) Get(id string) (payload.CreateSupplierResponse, error) {
	supplier, err := s.repo.Get(id)
	if err != nil {
		return payload.CreateSupplierResponse{}, err
	}
	return payload.CreateSupplierResponse{
		ID:     supplier.ID,
		Name:   supplier.Name,
		Rating: supplier.Rating,
	}, nil

}

func (s *Service) List() ([]payload.CreateSupplierResponse, error) {

	suppliers, err := s.repo.List("")
	var data []payload.CreateSupplierResponse

	if err != nil {
		return nil, err
	}
	for _, supplier := range suppliers {
		data = append(data, payload.CreateSupplierResponse{
			ID:   supplier.ID,
			Name: supplier.Name,
		})
	}
	return data, nil
}

func (s *Service) Update(supplier *accountant.Supplier) (bool, error) {
	return s.repo.Update(supplier)
}
func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}
