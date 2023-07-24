package models

type AthletesMedals struct {
	ID        int    `json:"id" gorm:"column:id"`
	Event     string `json:"event" gorm:"column:event"`
	Gold      string `json:"gold" gorm:"column:gold"`
	GoldNoc   string `json:"gold_noc" gorm:"column:gold_noc"`
	Silver    string `json:"silver" gorm:"column:silver"`
	SilverNoc string `json:"silver_noc" gorm:"column:silver_noc"`
	Bronze    string `json:"bronze" gorm:"column:bronze"`
	BronzeNoc string `json:"bronze_noc" gorm:"column:bronze_noc"`
	Sport     string `json:"sport" gorm:"column:sport"`
	Year      int    `json:"year" gorm:"column:year"`
}
