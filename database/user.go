package database

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

func (db *Database) SaveUser(user *User) error {
	return db.db.Save(user).Error
}

func (db *Database) FindByClerkID(id string) (*User, error) {
	var user User
	if err := db.db.First(&user, "clerk_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
