package repo

import (
	"context"
	"log"
	"samurenkoroma/services/pkg/db"
	"samurenkoroma/services/pkg/di"

	"gorm.io/gorm"
)

type CRUDRepositoryImpl[T any] struct {
	Database *db.Db
}

func NewCrudRepo[T any](database *db.Db) di.CRUDRepository[T] {
	return &CRUDRepositoryImpl[T]{
		Database: database,
	}
}

func (repo CRUDRepositoryImpl[T]) Save(entity *T) error {
	if err := gorm.G[T](repo.Database.DB).Create(context.Background(), entity); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (repo CRUDRepositoryImpl[T]) Get(id uint) (*T, error) {
	result, err := gorm.G[T](repo.Database.DB).Where("id = ?", id).First(context.Background())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (repo CRUDRepositoryImpl[T]) List(filter interface{}) ([]*T, error) {
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

func (repo CRUDRepositoryImpl[T]) Update(id uint, entity *T) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (repo CRUDRepositoryImpl[T]) Delete(id uint) error {
	//TODO implement me
	panic("implement me")
}
