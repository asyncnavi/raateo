package database

import (
	"context"
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var (
	initState = "00000-Init"
)

func (db *Database) Migrations(ctx context.Context) []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: initState,
			Migrate: func(tx *gorm.DB) error {
				// Create required extensions on new DB so AutoMigrate doesn't fail.
				return tx.Exec("CREATE EXTENSION IF NOT EXISTS hstore").Error
			},
			Rollback: func(tx *gorm.DB) error {
				return nil
			},
		},
		{
			ID: "00001-CreateUsers",
			Migrate: func(tx *gorm.DB) error {
				return tx.Migrator().CreateTable(&User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable((&User{}).TableName())
			},
		},
		{
			ID: "00002-CreateOrganizations",
			Migrate: func(tx *gorm.DB) error {
				return tx.Migrator().CreateTable(&Organization{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable((&Organization{}).TableName())
			},
		},
		{
			ID: "00003-CreateProducts",
			Migrate: func(tx *gorm.DB) error {
				return tx.Migrator().CreateTable(&Product{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable((&Product{}).TableName())
			},
		},
		{
			ID: "00004-CreateFeatures",
			Migrate: func(tx *gorm.DB) error {
				return tx.Migrator().CreateTable(&Feature{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable((&Feature{}).TableName())
			},
		},
		{
			ID: "00005-AddProductFields",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Migrator().AddColumn(&Product{}, "Description"); err != nil {
					return err
				}
				if err := tx.Migrator().AddColumn(&Product{}, "LogoURL"); err != nil {
					return err
				}
				if err := tx.Migrator().AddColumn(&Product{}, "ThumbnailURL"); err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Migrator().HasColumn(&Product{}, "Description") {
					if err := tx.Migrator().AddColumn(&Product{}, "Description"); err != nil {
						return err
					}
				}
				if tx.Migrator().HasColumn(&Product{}, "LogoURL") {
					if err := tx.Migrator().AddColumn(&Product{}, "LogoURL"); err != nil {
						return err
					}
				}
				if tx.Migrator().HasColumn(&Product{}, "ThumbnailURL") {
					if err := tx.Migrator().AddColumn(&Product{}, "ThumbnailURL"); err != nil {
						return err
					}
				}
				return tx.Migrator().DropTable((&Feature{}).TableName())
			},
		},
	}
}

func (db *Database) Migrate(ctx context.Context) error {
	m := gormigrate.New(db.db, gormigrate.DefaultOptions, db.Migrations(ctx))

	if err := m.Migrate(); err != nil {
		return fmt.Errorf("failed to migrate: %w", err)
	}

	return nil
}
