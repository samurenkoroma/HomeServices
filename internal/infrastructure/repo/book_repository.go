package repo

import (
	"samurenkoroma/services/internal/domain"
	"samurenkoroma/services/pkg/db"
)

type BookRepository interface {
	Create(b *domain.Book) (*domain.Book, error)
	GetList(params *BookQueryParams) (books []domain.Book)
	GetById(id uint) (*domain.Book, error)
	GetResourceById(id uint) (*domain.Resource, error)
}

func NewBookRepo(database *db.Db) BookRepository {
	return &BookRepositoryImpl{
		database: database,
	}
}

type BookRepositoryImpl struct {
	database *db.Db
}

// GetResourceById implements [BookRepository].
func (repo *BookRepositoryImpl) GetResourceById(id uint) (resource *domain.Resource, err error) {
	result := repo.database.DB.First(&resource, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return resource, nil
}

// CreateAuthor implements [BookRepository].
// func (repo *BookRepositoryImpl) CreateAuthor(a *Author) (*Author, error) {
// 	result := repo.repo.Save(a)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return a, nil
// }

func (repo *BookRepositoryImpl) GetById(id uint) (book *domain.Book, err error) {
	result := repo.database.DB.Preload("Resources").Preload("Authors").First(&book, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return book, nil
}

func (repo *BookRepositoryImpl) Create(book *domain.Book) (*domain.Book, error) {
	result := repo.database.Create(book)
	if result.Error != nil {

		return nil, result.Error
	}
	// resources := []entities.Resource
	// if (len(book.Resources) > 0){
	// 	resources
	// }
	// repo.repo.Model(&book).
	// 	Association("Resources").
	// 	Append(entities.Resource{
	// 		Type: entities.DocumentType,
	// 		Meta: "часть 1",
	// 		File: book.Resources.,
	// 	})

	return book, nil
}

type BookQueryParams struct {
	Limit  int `query:"limit"`
	Cursor int `query:"cursor"`
}

func NewBookQueryParams() *BookQueryParams {
	return &BookQueryParams{
		Limit: 20,
	}
}

func (repo *BookRepositoryImpl) GetList(params *BookQueryParams) (books []domain.Book) {
	repo.database.
		Table("books").
		Preload("Resources").
		Preload("Authors").
		Where("deleted_at is null and id >= ?", params.Cursor).
		Order("id asc").
		Limit(params.Limit).
		Find(&books)

	return books
}
