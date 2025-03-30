package repository

import (
	"errors"
	"github.com/asyncnavi/raateo/internal/models"
	"gorm.io/gorm"
)

func (repo *Repository) SaveOrganization(org *models.Organization) error {
	return repo.db.Save(org).Error
}
func (repo *Repository) GetOrganizationFromId(id uint) (*models.Organization, error) {
	var org models.Organization
	result := repo.db.First(&org, "id = ?", id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}

	return &org, nil
}
func (repo *Repository) GetProductsFromOrganizationId(id uint) ([]models.Product, error) {
	var products []models.Product
	result := repo.db.Find(&products, "organization_id = ?", id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return []models.Product{}, nil
		}
		return nil, result.Error
	}

	return products, nil
}
func (repo *Repository) OrganizationFromUser(id uint) (*models.Organization, error) {
	var org models.Organization
	result := repo.db.First(&org, "user_id = ?", id)

	return &org, result.Error
}
