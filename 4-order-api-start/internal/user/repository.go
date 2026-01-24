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

func (repo *UserRepository) FindIdByPhone(phone string) (uint, error) {
	var user User
	if err := repo.Database.Where("phone = ? ", phone).First(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}
