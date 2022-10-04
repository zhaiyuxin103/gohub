package migrations

import (
	"database/sql"
	"gohub/pkg/logger"
	"gohub/pkg/migrate"
	"gorm.io/gorm"
)

func init() {

	type User struct {
		City         string `gorm:"type:varchar(10);"`
		Introduction string `gorm:"type:varchar(255);"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.AutoMigrate(&User{})
		logger.LogIf(err)
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.DropColumn(&User{}, "City")
		logger.LogIf(err)
		err = migrator.DropColumn(&User{}, "Introduction")
		logger.LogIf(err)
	}

	migrate.Add("2022_10_04_093651_add_fields_to_user", up, down)
}
