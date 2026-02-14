package repo

import (
	"context"
	"log"
	"samurenkoroma/services/pkg/db"

	"gorm.io/gorm"
)

type Repository[T any] interface {
	Save(*T) error
	Get(string) (*T, error)
	List(string) ([]*T, error)
	Update(*T) (bool, error)
	Delete(string) error
}

type CRUDRepository[T any] struct {
	Database *db.Db
}

func NewCrudRepo[T any](database *db.Db) *CRUDRepository[T] {
	return &CRUDRepository[T]{
		Database: database,
	}
}

func (repo CRUDRepository[T]) Save(entity *T) error {
	if err := gorm.G[T](repo.Database.DB).Create(context.Background(), entity); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (repo CRUDRepository[T]) Get(id string) (*T, error) {
	result, err := gorm.G[T](repo.Database.DB).Where("id = ?", id).First(context.Background())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (repo CRUDRepository[T]) List(filter string) ([]*T, error) {
	var entities []*T
	result, err := gorm.G[T](repo.Database.DB).Find(context.Background())
	if err != nil {
		return nil, err
	}
	for i := range result {
		entities = append(entities, &result[i])
	}
	return entities, nil
}

func (repo CRUDRepository[T]) Update(entity *T) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (repo CRUDRepository[T]) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
