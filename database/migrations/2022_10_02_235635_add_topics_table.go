package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"
	"gorm.io/gorm"
)

func init() {

	type User struct {
		models.BaseModel
	}
	type Category struct {
		models.BaseModel
	}

	type Topic struct {
		models.BaseModel

		Title      string `gorm:"type:varchar(255);not null;index"`
		Body       string `gorm:"type:longtext;not null"`
		UserID     string `gorm:"type:bigint;not null;index"`
		CategoryID string `gorm:"type:bigint;not null;index"`

		// 会创建 user_id 和 category_id 外键的约束
		User     User
		Category Category

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.AutoMigrate(&Topic{})
		if err != nil {
			return
		}
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.DropTable(&Topic{})
		if err != nil {
			return
		}
	}

	migrate.Add("2022_10_02_235635_add_topics_table", up, down)
}
