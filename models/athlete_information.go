package models

type AthleteInformation struct {
	ID          int    `json:"id" gorm:"column:id"`
	Name        string `json:"name" gorm:"column:name"`
	Information string `json:"information" gorm:"column:information"`
	Prize       string `json:"prize" gorm:"column:prize"`
	Picture     string `json:"picture" gorm:"column:picture"`
}
