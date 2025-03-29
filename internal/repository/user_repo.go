package repository

import "github.com/asyncnavi/raateo/internal/models"

func (repo *Repository) SaveUser(user *models.User) error {
	return repo.db.Save(user).Error
}

func (repo *Repository) FindByClerkID(id string) (*models.User, error) {
	var user models.User
	if err := repo.db.Preload("User").Where(&user, "clerk_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
