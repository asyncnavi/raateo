package repository

import "github.com/asyncnavi/raateo/internal/models"

func (repo *Repository) SaveOrganization(org *models.Organization) error {
	return repo.db.Save(org).Error
}

func (repo *Repository) GetOrganizationFromId(id uint) (*models.Organization, error) {
	var org models.Organization
	if err := repo.db.First(&org, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &org, nil
}

func (repo *Repository) GetProductsFromOrganizationId(id uint) ([]models.Product, error) {
	var products []models.Product
	if err := repo.db.Find(&products, "organization_id = ?", id).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (repo *Repository) OrganizationFromUser(id uint) (*models.Organization, error) {
	var org models.Organization
	if err := repo.db.First(&org, "user_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &org, nil
}
