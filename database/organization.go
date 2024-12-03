package database

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name   string `gorm:"column:name;type:varchar(255);not null;unique"`
	UserID uint   `gorm:"column:user_id;not null;index"`

	User User
}

func (*Organization) TableName() string {
	return "organizations"
}

func (db *Database) SaveOrganization(org *Organization) error {
	return db.db.Save(org).Error
}

func (db *Database) FindOrganization(id int) (*Organization, error) {
	var org Organization
	if err := db.db.First(&org, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &org, nil
}

func (db *Database) FindOrganizationByUser(id int) (*Organization, error) {
	var org Organization
	if err := db.db.First(&org, "user_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &org, nil
}
