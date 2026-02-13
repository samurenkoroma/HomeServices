package link

import (
	"samurenkoroma/services/pkg/db"

	"gorm.io/gorm/clause"
)

type LinkRepository struct {
	database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		database: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.database.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) Update(link *Link) (*Link, error) {

	result := repo.database.DB.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) Delete(id uint) error {

	result := repo.database.DB.Delete(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repo.database.DB.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *LinkRepository) Exist(id uint) error {
	result := repo.database.DB.First(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

type LinkQueryParams struct {
	Limit  int
	Cursor int
}

func (repo *LinkRepository) Count() int64 {
	var count int64
	repo.database.
		Table("links").
		Where("deleted_at is null").
		Count(&count)
	return count

}

func (repo *LinkRepository) GetList(params LinkQueryParams) []Link {
	var links []Link
	repo.database.
		Table("links").
		Where("deleted_at is null and id > ?", params.Cursor).
		Order("id asc").
		Limit(params.Limit).
		Scan(&links)

	return links
}
