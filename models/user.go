package models

import "time"

type User struct {
	ID         int       `json:"id" gorm:"column:id"`
	Email      string    `json:"email" gorm:"column:email"`
	Password   string    `json:"password" gorm:"column:password"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	IsValid    string    `json:"is_valid" gorm:"column:is_valid"`
}
