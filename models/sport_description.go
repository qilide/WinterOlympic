package models

type SportDescription struct {
	ID          int    `json:"id" gorm:"column:id"`
	Sport       string `json:"sport" gorm:"column:sport"`
	Description string `json:"description" gorm:"column:description"`
}
