package models

type PredictResult struct {
	ID          int    `json:"id" gorm:"column:id"`
	CountryName string `json:"country_name" gorm:"column:country_name"`
	GoldCount   int    `json:"gold_count" gorm:"column:gold_count"`
	SilverCount int    `json:"silver_count" gorm:"column:silver_count"`
	BronzeCount int    `json:"bronze_count" gorm:"column:bronze_count"`
}
