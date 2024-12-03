package database

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	OrganizationID uint   `gorm:"column:organization_id;not null;index;"`
	Name           string `gorm:"column:name;type:text;not null;"`
	Description    string `gorm:"column:description; type:text"`
	LogoURL        string `gorm:"column:logo_url;type:text"`
	ThumbnailURL   string `gorm:"column:thumbnail_url;type:text"`
	Organization   Organization
}

func (*Product) TableName() string {
	return "products"
}
