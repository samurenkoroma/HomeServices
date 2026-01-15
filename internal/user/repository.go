package user

import "samurenkoroma/services/pkg/db"

type UserRepository struct {
	database *db.Db
}

func NewUserRepo(database *db.Db) *UserRepository {
	return &UserRepository{
		database: database,
	}
}

func (repo *UserRepository) Update(email string, user *User) error {
	result := repo.database.Model(&User{}).Where("email = ?", email).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	result := repo.database.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepository) FindByRefresh(refresh string) (*User, error) {
	var user User
	result := repo.database.DB.First(&user, "refresh_token = ?", refresh)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
