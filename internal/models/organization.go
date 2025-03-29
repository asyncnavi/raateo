package models

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name        string `gorm:"column:name;type:varchar(255);not null;unique"`
	Description string `gorm:"column:description;type:text"`
	UserID      uint   `gorm:"column:user_id;not null;index"`
	LogoURL     string `gorm:"column:logo_url;type:text"`

	User User
}

func (*Organization) TableName() string {
	return "organizations"
}
