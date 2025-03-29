package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Email     string `gorm:"type:varchar(320);index;unique;not null;"`
	FirstName string `gorm:"type:varchar(255);not null"`
	LastName  string `gorm:"type:varchar(255);"`

	ClerkID string `gorm:"type:text;index;unique;not null;"`
}

func (*User) TableName() string {
	return "users"
}
