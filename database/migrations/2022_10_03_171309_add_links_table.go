package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"
	"gorm.io/gorm"
)

func init() {

	type Link struct {
		models.BaseModel

		Name string `gorm:"type:varchar(255);not null"`
		URL  string `gorm:"type:varchar(255);default:null"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.AutoMigrate(&Link{})
		if err != nil {
			return
		}
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.DropTable(&Link{})
		if err != nil {
			return
		}
	}

	migrate.Add("2022_10_03_171309_add_links_table", up, down)
}
