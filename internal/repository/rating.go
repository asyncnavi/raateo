package repository

import "github.com/asyncnavi/raateo/internal/models"

func (repo *Repository) SaveRating(rating *models.Rating) error {
	return repo.db.Save(rating).Error
}

func (repo *Repository) GetRatings(organizationID uint, productID uint, featureID uint) ([]models.Rating, error) {
	var ratings []models.Rating

	result := repo.db.Where("organization_id = ? AND product_id = ? AND feature_id = ?", organizationID, productID, featureID).Find(&ratings)

	return ratings, result.Error
}
