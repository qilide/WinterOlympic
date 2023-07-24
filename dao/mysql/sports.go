package mysql

import (
	"WinterOlympic/models"
)

func CountryMedalVerity(year string) ([]models.CountryMedals, []models.CountryMedals) {
	var data1 []models.CountryMedals
	var data2 []models.CountryMedals
	db.Table("country_medals").Where("year = ?", year).Order("gold desc").
		Order("silver desc").Order("bronze desc").Order("total desc").Limit(5).Find(&data1)
	db.Table("country_medals").Where("year = ?", year).Order("total desc").
		Order("gold desc").Order("silver desc").Order("bronze desc").Limit(5).Find(&data2)
	return data1, data2
}

func CountryMedalDetailVerity(year string) []models.CountryMedals {
	var data []models.CountryMedals
	db.Table("country_medals").Where("year = ?", year).Order("gold desc").
		Order("silver desc").Order("bronze desc").Order("total desc").Find(&data)
	return data
}

func SportsMedalVerity(year string) []models.CountryEventMedals {
	var data []models.CountryEventMedals
	db.Table("country_event_medals").Where("year = ?", year).Find(&data)
	return data
}

func CountryHumanVerity(year string) []models.CountryHuman {
	var data []models.CountryHuman
	db.Table("country_human").Where("year = ?", year).Order("total desc").
		Order("men desc").Order("women desc").Order("country desc").Find(&data)
	return data
}

func AthleteMedalVerity(year string) []models.AthletesMedals {
	var data []models.AthletesMedals
	db.Table("athletes_medals").Where("year = ?", year).Find(&data)
	return data
}

func ChinaMedalVerity(year string) []models.CountryEventMedals {
	var data []models.CountryEventMedals
	db.Table("country_event_medals").Where("year = ?", year).Where("country = ?", "中国").Find(&data)
	return data
}

func MomentVerity(SportName string) []models.SportVideo {
	var data []models.SportVideo
	db.Table("sport_video").Where("name = ?", SportName).Find(&data)
	return data
}
