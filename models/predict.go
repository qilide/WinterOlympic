package models

type Predict struct {
	ID            int     `json:"id" gorm:"column:id"`
	Name          string  `json:"name" gorm:"column:name"`
	Noc           string  `json:"noc" gorm:"column:noc"`
	Men           int     `json:"men" gorm:"column:men"`
	Women         int     `json:"women" gorm:"column:women"`
	IsHome        int     `json:"is_home" gorm:"column:is_home"`
	HumanGdp      float64 `json:"human_gdp" gorm:"column:human_gdp"`
	Gdp           float64 `json:"gdp" gorm:"column:gdp"`
	HumanTotal    int     `json:"human_total" gorm:"column:human_total"`
	SocialSystem  int     `json:"social_system" gorm:"column:social_system"`
	GoldRank      int     `json:"gold_rank" gorm:"column:gold_rank"`
	SilverRank    int     `json:"silver_rank" gorm:"column:silver_rank"`
	BronzeRank    int     `json:"bronze_rank" gorm:"column:bronze_rank"`
	GoldPercent   float64 `json:"gold_percent" gorm:"column:gold_percent"`
	SilverPercent float64 `json:"silver_percent" gorm:"column:silver_percent"`
	BronzePercent float64 `json:"bronze_percent" gorm:"column:bronze_percent"`
	GoldCount     int     `json:"gold_count" gorm:"column:gold_count"`
	SilverCount   int     `json:"silver_count" gorm:"column:silver_count"`
	BronzeCount   int     `json:"bronze_count" gorm:"column:bronze_count"`
	Year          int     `json:"year" gorm:"column:year"`
}
