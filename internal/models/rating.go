package models

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	OrganizationID uint   `gorm:"column:organization_id;not null;index;"`
	ProductID      uint   `gorm:"column:product_id;not null;index;"`
	FeatureID      uint   `gorm:"column:feature_id;not null;index;"`
	RatingValue    uint   `gorm:"column:rating_value;not null;"`
	UserID         uint   `gorm:"column:user_id;not null;index;"`
	Comment        string `gorm:"column:comment;type:text;"`

	User         User         `gorm:"foreignKey:UserID"`
	Organization Organization `gorm:"foreignKey:OrganizationID"`
	Product      Product      `gorm:"foreignKey:ProductID"`
	Feature      Feature      `gorm:"foreignKey:FeatureID"`
}

func (*Rating) TableName() string {
	return "ratings"
}
