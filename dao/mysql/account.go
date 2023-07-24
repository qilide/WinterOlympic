package mysql

import (
	"WinterOlympic/models"
	"time"
)

// 登录参数数据库查询

func LoginVerity(email string) models.User {
	var user models.User
	db.Table("user").Where("email = ?", email).Find(&user)
	return user
}

// 注册参数数据库查询

func RegisterVerity(email string) models.User {
	var user models.User
	db.Table("user").Where("email = ?", email).Find(&user)
	return user
}

// 创建数据库数据

func CreateUser(id int, email string, password string, createTime time.Time) models.User {
	user := models.User{
		ID:         id,
		Email:      email,
		Password:   password,
		CreateTime: createTime,
		IsValid:    "1",
	}
	db.Table("user").Create(&user)
	return user
}
