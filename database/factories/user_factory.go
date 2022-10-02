// Package factories 存放工厂方法
package factories

import (
	"github.com/go-faker/faker/v4"
	"gohub/app/models/user"
	"gohub/pkg/helpers"
)

func MakeUsers(times int) []user.User {

	var objs []user.User

	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		model := user.User{
			Name:     faker.Username(),
			Email:    faker.Email(),
			Phone:    helpers.RandomNumber(11),
			Password: "$2a$14$.m9vGIvVPQvzpeSlRsGNKOImjmJI1qkc3vMrXWgb00elRqVO1/kCG",
		}
		objs = append(objs, model)
	}

	return objs
}
