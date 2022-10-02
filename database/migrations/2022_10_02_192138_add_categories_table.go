package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"
	"gorm.io/gorm"
)

func init() {

	type Category struct {
		models.BaseModel

		Name        string `gorm:"type:varchar(255);not null;index"`
		Description string `gorm:"type:varchar(255);default:null"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.AutoMigrate(&Category{})
		if err != nil {
			return
		}
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.DropTable(&Category{})
		if err != nil {
			return
		}
	}

	migrate.Add("2022_10_02_192138_add_categories_table", up, down)
}
