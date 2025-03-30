package repository

import "github.com/asyncnavi/raateo/internal/models"

func (repo *Repository) SaveProduct(prod *models.Product) error {
	return repo.db.Save(prod).Error
}

func (repo *Repository) DeleteProduct(productId uint, organizationId uint) error {
	return repo.db.Where("id = ? AND organization_id = ? AND deleted_at IS NULL", productId, organizationId).Delete(&models.Product{}).Error
}

func (repo *Repository) GetProduct(productId uint, organizationId uint) (models.Product, error) {
	var prod models.Product
	result := repo.db.Where("id = ? AND organization_id = ? AND deleted_at IS NULL", productId, organizationId).First(&prod)
	if result.Error != nil {
		return models.Product{}, result.Error
	}
	return prod, nil
}

func (repo *Repository) GetProducts() ([]models.Product, error) {
	var products []models.Product
	result := repo.db.Where("deleted_at IS NULL").Find(&products)
	return products, result.Error
}

func (repo *Repository) GetPaginatedProducts(page int, limit int) ([]models.Product, error) {
	var products []models.Product
	result := repo.db.Offset((page - 1) * limit).Limit(limit).Where("deleted_at IS NULL").Find(&products)
	return products, result.Error
}

func (repo *Repository) GetFeaturesForProduct(productId uint, organizationId uint) ([]models.Feature, error) {
	var features []models.Feature
	result := repo.db.Where("product_id = ? AND organization_id = ? AND deleted_at IS NULL", productId, organizationId).Find(&features)
	return features, result.Error
}

func (repo *Repository) GetAllFeatureByProduct(id uint) ([]models.Feature, error) {
	var features []models.Feature
	result := repo.db.Where("product_id = ? AND deleted_at IS NULLs", id).Find(&features)
	return features, result.Error
}

func (repo *Repository) GetProductsByOrganization(id uint) ([]models.Product, error) {
	var products []models.Product
	result := repo.db.Where("organization_id = ? AND deleted_at IS NULL", id).Find(&products)
	return products, result.Error
}

func (repo *Repository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	result := repo.db.Where("deleted_at IS NULL").Find(&products)
	return products, result.Error
}

func (repo *Repository) GetProductWithQuery(query string) ([]models.Product, error) {
	var products []models.Product
	result := repo.db.Where("LOWER(name) LIKE LOWER(?) AND deleted_at IS NULL", "%"+query+"%").Find(&products)
	return products, result.Error
}
