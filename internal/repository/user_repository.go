package repository

import (
	"errors"
	"gorm.io/gorm"
	"itv/internal/model"
	"itv/pkg/auth"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	result := r.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) EnsureAdminExists(username, password string) error {
	var count int64
	r.db.Model(&model.User{}).Where("role = ?", "admin").Count(&count)

	if count == 0 {
		hashedPassword, err := auth.HashPassword(password)
		if err != nil {
			return err
		}

		admin := model.User{
			Username: username,
			Password: hashedPassword,
			Role:     "admin",
		}

		return r.db.Create(&admin).Error
	}

	return nil
}
