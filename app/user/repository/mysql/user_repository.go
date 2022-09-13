package repository

import (
	"gorm.io/gorm"

	"github.com/t0239184/GoChat/app/domain"
	"github.com/t0239184/GoChat/app/tool"
)

type UserRepository struct {
	db *gorm.DB
}


func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user *domain.User, salt *domain.Salt) (*int64, error) {
	var err error
	r.db.Transaction(func (tx *gorm.DB) error {
		tool.Logger.Info("Save Salt into database.")
		if err = tx.Create(salt).Error; err != nil {
			return err
		}

		tool.Logger.Info("Save User into database.")
		user.SaltId = salt.Id
		if err = tx.Create(user).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return &user.Id, nil

}