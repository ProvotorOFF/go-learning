package user

import "order-api-start/pkg/db"

type UserRepository struct {
	Database *db.Db
}

func (repo *UserRepository) FindOrCreate(user *User) (*User, error) {
	result := repo.Database.Where("phone = ?", user.Phone).FirstOrCreate(user)
	return user, result.Error
}
