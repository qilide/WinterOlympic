package models

type CountryMedals struct {
	ID      int    `json:"id" gorm:"column:id"`
	Country string `json:"country" gorm:"column:country"`
	Noc     string `json:"noc" gorm:"column:noc"`
	Gold    int    `json:"gold" gorm:"column:gold"`
	Silver  int    `json:"silver" gorm:"column:silver"`
	Bronze  int    `json:"bronze" gorm:"column:bronze"`
	Total   int    `json:"total" gorm:"column:total"`
	Year    int    `json:"year" gorm:"column:year"`
}
