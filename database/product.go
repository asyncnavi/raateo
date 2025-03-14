package database

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	OrganizationID uint   `gorm:"column:organization_id;not null;index;"`
	Name           string `gorm:"column:name;type:text;not null;"`
	Description    string `gorm:"column:description; type:text"`
	LogoURL        string `gorm:"column:logo_url;type:text"`
	ThumbnailURL   string `gorm:"column:thumbnail_url;type:text"`
	WebsiteURL     string `gorm:"column:website_url;type:text"`
	Organization   Organization
}

func (*Product) TableName() string {
	return "products"
}

func (db *Database) SaveProduct(prod *Product) error {
	return db.db.Save(prod).Error
}

func (db *Database) GetProduct(id uint) (Product, error) {
	var prod Product
	result := db.db.Where("id = ?", id).First(&prod)
	if result.Error != nil {
		return Product{}, result.Error
	}
	return prod, nil
}

func (db *Database) GetProductsByOrganization(id uint) ([]Product, error) {
	var products []Product
	result := db.db.Where("organization_id = ?", id).Find(&products)
	return products, result.Error
}

// GetAllProducts returns all products
func (db *Database) GetAllProducts() ([]Product, error) {
	var products []Product
	result := db.db.Find(&products)
	return products, result.Error
}
