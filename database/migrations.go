package database

import (
	"context"
	"fmt"

	"github.com/asyncnavi/raateo/internal/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var initState = "00000-Init"

func Migrations(ctx context.Context) []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: initState,
			Migrate: func(tx *gorm.DB) error {
				// Check if extension exists before creating
				var count int64
				tx.Raw("SELECT count(*) FROM pg_extension WHERE extname = 'hstore'").Scan(&count)
				if count == 0 {
					return tx.Exec("CREATE EXTENSION IF NOT EXISTS hstore").Error
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return nil
			},
		},
		{
			ID: "00001-CreateUsers",
			Migrate: func(tx *gorm.DB) error {
				if !tx.Migrator().HasTable(&models.User{}) {
					return tx.Migrator().CreateTable(&models.User{})
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Migrator().HasTable(&models.User{}) {
					return tx.Migrator().DropTable((&models.User{}).TableName())
				}
				return nil
			},
		},
		{
			ID: "00002-CreateOrganizations",
			Migrate: func(tx *gorm.DB) error {
				if !tx.Migrator().HasTable(&models.Organization{}) {
					return tx.Migrator().CreateTable(&models.Organization{})
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Migrator().HasTable(&models.Organization{}) {
					return tx.Migrator().DropTable((&models.Organization{}).TableName())
				}
				return nil
			},
		},
		{
			ID: "00003-CreateProducts",
			Migrate: func(tx *gorm.DB) error {
				if !tx.Migrator().HasTable(&models.Product{}) {
					return tx.Migrator().CreateTable(&models.Product{})
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Migrator().HasTable(&models.Product{}) {
					return tx.Migrator().DropTable((&models.Product{}).TableName())
				}
				return nil
			},
		},
		{
			ID: "00004-CreateFeatures",
			Migrate: func(tx *gorm.DB) error {
				if !tx.Migrator().HasTable(&models.Feature{}) {
					return tx.Migrator().CreateTable(&models.Feature{})
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Migrator().HasTable(&models.Feature{}) {
					return tx.Migrator().DropTable((&models.Feature{}).TableName())
				}
				return nil
			},
		},
		{
			ID: "00005-AddProductFields",
			Migrate: func(tx *gorm.DB) error {
				if !tx.Migrator().HasColumn(&models.Product{}, "Description") {
					if err := tx.Migrator().AddColumn(&models.Product{}, "Description"); err != nil {
						return err
					}
				}
				if !tx.Migrator().HasColumn(&models.Product{}, "LogoURL") {
					if err := tx.Migrator().AddColumn(&models.Product{}, "LogoURL"); err != nil {
						return err
					}
				}
				if !tx.Migrator().HasColumn(&models.Product{}, "ThumbnailURL") {
					if err := tx.Migrator().AddColumn(&models.Product{}, "ThumbnailURL"); err != nil {
						return err
					}
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Migrator().HasColumn(&models.Product{}, "Description") {
					if err := tx.Migrator().DropColumn(&models.Product{}, "Description"); err != nil {
						return err
					}
				}
				if tx.Migrator().HasColumn(&models.Product{}, "LogoURL") {
					if err := tx.Migrator().DropColumn(&models.Product{}, "LogoURL"); err != nil {
						return err
					}
				}
				if tx.Migrator().HasColumn(&models.Product{}, "ThumbnailURL") {
					if err := tx.Migrator().DropColumn(&models.Product{}, "ThumbnailURL"); err != nil {
						return err
					}
				}
				return nil
			},
		},
		{
			ID: "00006-AddWebsiteURLFieldToProduct",
			Migrate: func(tx *gorm.DB) error {
				if !tx.Migrator().HasColumn(&models.Product{}, "WebsiteURL") {
					if err := tx.Migrator().AddColumn(&models.Product{}, "WebsiteURL"); err != nil {
						return err
					}
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Migrator().HasColumn(&models.Product{}, "WebsiteURL") {
					if err := tx.Migrator().DropColumn(&models.Product{}, "WebsiteURL"); err != nil {
						return err
					}
				}
				return nil
			},
		},
		{
			ID: "00006-AddStatusFieldToFeature",
			Migrate: func(tx *gorm.DB) error {
				if !tx.Migrator().HasColumn(&models.Feature{}, "Status") {
					if err := tx.Migrator().AddColumn(&models.Feature{}, "Status"); err != nil {
						return err
					}
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Migrator().HasColumn(&models.Feature{}, "Status") {
					if err := tx.Migrator().DropColumn(&models.Feature{}, "Status"); err != nil {
						return err
					}
				}
				return nil
			},
		},
		{
			ID: "00007-AddLogoURLFieldToOrganization",
			Migrate: func(tx *gorm.DB) error {
				if !tx.Migrator().HasColumn(&models.Organization{}, "LogoURL") {
					if err := tx.Migrator().AddColumn(&models.Organization{}, "LogoURL"); err != nil {
						return err
					}
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Migrator().HasColumn(&models.Organization{}, "LogoURL") {
					if err := tx.Migrator().DropColumn(&models.Organization{}, "LogoURL"); err != nil {
						return err
					}
				}
				return nil
			},
		},
		{
			ID: "00008-CreateRatings",
			Migrate: func(tx *gorm.DB) error {
				if !tx.Migrator().HasTable(&models.Rating{}) {
					return tx.Migrator().CreateTable(&models.Rating{})
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Migrator().HasTable(&models.Rating{}) {
					return tx.Migrator().DropTable((&models.Rating{}).TableName())
				}
				return nil
			},
		},
		{
			ID: "00009-AddFeatureIDFieldToRating",
			Migrate: func(tx *gorm.DB) error {
				if !tx.Migrator().HasColumn(&models.Rating{}, "FeatureID") {
					if err := tx.Migrator().AddColumn(&models.Rating{}, "FeatureID"); err != nil {
						return err
					}
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Migrator().HasColumn(&models.Rating{}, "FeatureID") {
					if err := tx.Migrator().DropColumn(&models.Rating{}, "FeatureID"); err != nil {
						return err
					}
				}
				return nil
			},
		},
		{
			ID: "00010-AddOrganizationDescriptionField",
			Migrate: func(tx *gorm.DB) error {
				if !tx.Migrator().HasColumn(&models.Organization{}, "Description") {
					if err := tx.Migrator().AddColumn(&models.Organization{}, "Description"); err != nil {
						return err
					}
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Migrator().HasColumn(&models.Organization{}, "Description") {
					if err := tx.Migrator().DropColumn(&models.Organization{}, "Description"); err != nil {
						return err
					}
				}
				return nil
			},
		},
		{
			ID: "00011-AddProductIDToRatings",
			Migrate: func(tx *gorm.DB) error {
				if !tx.Migrator().HasColumn(&models.Rating{}, "ProductID") {
					if err := tx.Migrator().AddColumn(&models.Rating{}, "ProductID"); err != nil {
						return err
					}
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if tx.Migrator().HasColumn(&models.Rating{}, "ProductID") {
					if err := tx.Migrator().DropColumn(&models.Rating{}, "ProductID"); err != nil {
						return err
					}
				}
				return nil
			},
		},
	}
}

func Migrate(ctx context.Context, db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, Migrations(ctx))

	if err := m.Migrate(); err != nil {
		return fmt.Errorf("failed to migrate: %w", err)
	}

	return nil
}
