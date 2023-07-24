package mysql

import (
	"WinterOlympic/models"
)

func MedalInfoVerity(year string, countryId string) (string, models.CountryMedals) {
	var data models.CountryMedals
	var data1 models.CountryHuman
	db.Table("country_human").Where("id = ?", countryId).Find(&data1)
	db.Table("country_medals").Where("year = ?", year).Where("country = ?", data1.Country).Find(&data)
	return data1.Country, data
}

func SportsInfoVerity(year string, countryId string) []models.CountryEventMedals {
	var data []models.CountryEventMedals
	var data1 models.CountryHuman
	db.Table("country_human").Where("id = ?", countryId).Find(&data1)
	db.Table("country_event_medals").Where("year = ?", year).Where("noc = ?", data1.Noc).
		Order("gold desc").Order("silver desc").Order("bronze desc").Order("total desc").Find(&data)
	return data
}

func AthleteInfoVerity(year string, countryId string) (string, []models.AthletesMedals) {
	var data []models.AthletesMedals
	var data2 []models.AthletesMedals
	var data1 models.CountryHuman
	db.Table("country_human").Where("id = ?", countryId).First(&data1)
	db.Table("athletes_medals").Where("year = ?", year).Find(&data2)
	for _, data3 := range data2 {
		if data3.GoldNoc == data1.Noc || data3.SilverNoc == data1.Noc || data3.BronzeNoc == data1.Noc {
			data = append(data, data3)
		}
	}
	return data1.Noc, data
}

func TopSportsInfoVerity(year string, countryId string) models.SportDescription {
	var data models.SportDescription
	var data1 models.CountryHuman
	var data2 models.CountryEventMedals
	db.Table("country_human").Where("id = ?", countryId).Find(&data1)
	db.Table("country_event_medals").Where("year = ?", year).Where("country = ?", data1.Country).
		Order("gold desc").Order("silver desc").Order("bronze desc").Order("total desc").First(&data2)
	db.Table("sport_description").Where("sport = ?", data2.Sport).Find(&data)
	return data
}

func TopAthleteInfoVerity(countryId string) (string, []models.AthletesMedals) {
	var data []models.AthletesMedals
	var data2 []models.AthletesMedals
	var data1 models.CountryHuman
	db.Table("country_human").Where("id = ?", countryId).First(&data1)
	db.Table("athletes_medals").Find(&data2)
	for _, data3 := range data2 {
		if data3.GoldNoc == data1.Noc || data3.SilverNoc == data1.Noc || data3.BronzeNoc == data1.Noc {
			data = append(data, data3)
		}
	}
	return data1.Noc, data
}

func TopAthleteVerity(AthleteName interface{}) models.AthleteInformation {
	var data models.AthleteInformation
	db.Table("athlete_information").Where("name = ?", AthleteName).First(&data)
	return data
}

func PredictInfoVerity(countryId string) models.PredictResult {
	var data models.PredictResult
	var data1 models.CountryHuman
	db.Table("country_human").Where("id = ?", countryId).Find(&data1)
	db.Table("predict_result").Where("country_name = ?", data1.Country).Find(&data)
	return data
}
