package models

type CountryHuman struct {
	ID         int    `json:"id" gorm:"column:id"`
	Noc        string `json:"noc" gorm:"column:noc"`
	Country    string `json:"country" gorm:"column:country"`
	CountryMap string `json:"country_map" gorm:"column:country_map"`
	Men        int    `json:"men" gorm:"column:men"`
	Women      int    `json:"women" gorm:"column:women"`
	Total      int    `json:"total" gorm:"column:total"`
	Year       int    `json:"year" gorm:"column:year"`
}
