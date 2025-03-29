package repository

import (
	"github.com/asyncnavi/raateo/internal/models"
)

func (repo *Repository) SaveFeature(feature *models.Feature) error {
	return repo.db.Save(feature).Error
}

func (repo *Repository) GetFeaturesByProduct(productID uint) ([]models.Feature, error) {
	var features []models.Feature
	if err := repo.db.Where("product_id = ?", productID).Find(&features).Error; err != nil {
		return nil, err
	}
	return features, nil
}

func (repo *Repository) GetFeature(productId uint, organizationId uint, featureId uint) (*models.Feature, error) {
	var feature models.Feature
	result := repo.db.Where("product_id = ? AND organization_id = ? AND id = ?", productId, organizationId, featureId).First(&feature)
	return &feature, result.Error
}

func (repo *Repository) GetSingleFeature(productID, featureID uint) (*models.Feature, error) {
	var feature models.Feature
	if err := repo.db.Where("id = ? AND product_id = ?", featureID, productID).First(&feature).Error; err != nil {
		return nil, err
	}
	return &feature, nil
}
