package database

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	OrganizationID uint   `gorm:"column:organization_id;not null;index;"`
	Name           string `gorm:"column:name;type:text;not null;"`

	Organization Organization
}

func (*Product) TableName() string {
	return "products"
}
