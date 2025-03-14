package database

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	OrganizationID uint   `gorm:"column:organization_id;not null;index;"`
	ProductID      uint   `gorm:"column:organization_id;not null;index;"`
	RatingValue    uint   `gorm:"column:rating_value;not null;"`
	UserID         uint   `gorm:"column:user_id;not null;index;"`
	Comment        string `gorm:"column:comment;type:text;"`
}

func (*Rating) TableName() string {
	return "ratings"
}
