package postgresrepo

import (
	"OTP-JWT-sample/model"
	"OTP-JWT-sample/repository"
	"gorm.io/gorm"
)

type postgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) repository.UserRepository {
	return &postgresUserRepository{db: db}
}

func (r *postgresUserRepository) FindByMobile(mobile string) (*model.User, error) {
	var user model.User
	err := r.db.Where("mobile = ?", mobile).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *postgresUserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}
