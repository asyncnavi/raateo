package repository

import (
	"errors"
	"github.com/asyncnavi/raateo/internal/models"
	"gorm.io/gorm"
)

func (repo *Repository) SaveUser(user *models.User) error {
	return repo.db.Save(user).Error
}
func (repo *Repository) FindByClerkID(id string) (*models.User, error) {
	var user models.User
	result := repo.db.Where("clerk_id = ?", id).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}

	return &user, nil
}
