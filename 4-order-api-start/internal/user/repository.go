package user

import (
	"order-api-start/pkg/db"
)

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{database}
}

func (repo *UserRepository) FindOrCreate(user *User) (*User, error) {
	result := repo.Database.Where("phone = ?", user.Phone).FirstOrCreate(user)
	return user, result.Error
}

func (repo *UserRepository) FindById(id int) (*User, error) {
	var user User
	result := repo.Database.Where("id = ?", id).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
