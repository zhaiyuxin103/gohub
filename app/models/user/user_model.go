// Package user 存放用户 Model 相关逻辑
package user

import (
	"gohub/app/models"
	"gohub/pkg/database"
	"gohub/pkg/hash"
	"gorm.io/datatypes"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name         string         `json:"name,omitempty"`
	Gender       uint           `json:"gender,omitempty"`
	Birthday     datatypes.Date `json:"birthday,omitempty"`
	State        bool           `json:"state,omitempty"`
	City         string         `json:"city,omitempty"`
	Introduction string         `json:"introduction,omitempty"`
	Avatar       string         `json:"avatar,omitempty"`
	Email        string         `json:"-"`
	Phone        string         `json:"-"`
	Password     string         `json:"-"`
	Order        uint64         `json:"order,omitempty"`

	models.CommonTimestampsField
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Omit("Birthday", "DeletedAt").Save(&userModel)
	return result.RowsAffected
}
