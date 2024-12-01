package database

import (
	"gorm.io/gorm"
)

type Feature struct {
	gorm.Model
	OrganizationID uint   `gorm:"column:organization_id;not null;index;"`
	ProductID      uint   `gorm:"column:product_id;not null;index;"`
	Name           string `gorm:"column:name;type:text;not null;"`
	Description    string `gorm:"column:description;type:text"`
	VideoUrl       string `gorm:"column:video_url;type:text;"`
	ThumbnailUrl   string `gorm:"column:thumbnail_url;type:text;"`
}

func (*Feature) TableName() string {
	return "features"
}
