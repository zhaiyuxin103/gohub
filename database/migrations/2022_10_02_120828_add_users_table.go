package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"
	"gorm.io/datatypes"
	"time"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		models.BaseModel

		Name            string         `gorm:"type:varchar(255);not null;index;comment:姓名"`
		Email           string         `gorm:"type:varchar(255);index;default:null;comment:邮箱"`
		EmailVerifiedAt time.Time      `gorm:"column:deleted_at;default:null" json:"email_verified_at,omitempty;comment:邮箱验证时间"`
		Phone           string         `gorm:"type:varchar(20);index;default:null;comment:手机号"`
		Gender          uint           `gorm:"type:tinyint(1);comment:性别"`
		Birthday        datatypes.Date `gorm:"type:date;comment:生日"`
		Avatar          string         `gorm:"varchar(255);comment:头像"`
		State           bool           `gorm:"type:bool;default:true;comment:状态"`
		Password        string         `gorm:"type:varchar(255);comment:密码"`
		RememberToken   string         `gorm:"type:varchar(100);comment:令牌"`
		Order           uint64         `gorm:"type:int(11);default:0;comment:排序"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.AutoMigrate(&User{})
		if err != nil {
			return
		}
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.DropTable(&User{})
		if err != nil {
			return
		}
	}

	migrate.Add("2022_10_02_120828_add_users_table", up, down)
}
