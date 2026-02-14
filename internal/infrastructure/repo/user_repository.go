package repo

import (
	"samurenkoroma/services/pkg/db"
	"samurenkoroma/services/services/account"
)

type UserRepository struct {
	database *db.Db
}

func NewUserRepo(database *db.Db) *UserRepository {
	return &UserRepository{
		database: database,
	}
}

func (repo *UserRepository) Update(email string, user *account.User) error {
	result := repo.database.Model(&account.User{}).Where("email = ?", email).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *UserRepository) Create(user *account.User) (*account.User, error) {
	result := repo.database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*account.User, error) {
	var user account.User
	result := repo.database.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepository) FindByRefresh(refresh string) (*account.User, error) {
	var user account.User
	result := repo.database.DB.First(&user, "refresh_token = ?", refresh)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
