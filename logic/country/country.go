package country

import (
	"WinterOlympic/dao/mysql"
	"WinterOlympic/util/SportsUtil"
	"errors"
	"sort"
)

type Country struct {
}

var (
	NotMedal = errors.New("NotMedal") //该国家在此届冬奥会未获得奖牌
)

//国家奖牌信息展示

func (lc Country) MedalInfo(year string, countryId string) (interface{}, error) {
	CountryName, data := mysql.MedalInfoVerity(year, countryId)
	TotalData := make([]interface{}, 0)
	if data.ID == 0 {
		data1 := map[string]interface{}{
			"country_name": CountryName,
		}
		TotalData = append(TotalData, data1)
		return TotalData, NotMedal
	} else {
		data1 := map[string]interface{}{
			"country_name": data.Country,
			"silver":       data.Gold,
			"gold":         data.Silver,
			"bronze":       data.Bronze,
			"total":        data.Total,
		}
		TotalData = append(TotalData, data1)
		return TotalData, nil
	}
}

//国家详细项目获得奖牌

func (lc Country) SportsInfo(year string, countryId string) (interface{}, error) {
	data := mysql.SportsInfoVerity(year, countryId)
	TotalData := make([]interface{}, 0)
	for i := 0; i < len(data); i++ {
		data1 := map[string]interface{}{
			"sport_name": data[i].Sport,
			"silver":     data[i].Gold,
			"gold":       data[i].Silver,
			"bronze":     data[i].Bronze,
			"total":      data[i].Total,
		}
		TotalData = append(TotalData, data1)
	}
	if len(TotalData) == 0 {
		return 0, NotMedal
	} else {
		return TotalData, nil
	}
}

//国家详细运动员获得奖牌

func (lc Country) AthleteInfo(year string, countryId string) (interface{}, error) {
	Noc, data := mysql.AthleteInfoVerity(year, countryId)
	TotalData := make([]interface{}, 0)
	GoldSort := SportsUtil.AllAthleteSortSlice{}
	AthleteList1 := make([]interface{}, 0)
	for _, athlete1 := range data {
		AthleteList1 = append(AthleteList1, athlete1.Gold)
		AthleteList1 = append(AthleteList1, athlete1.Silver)
		AthleteList1 = append(AthleteList1, athlete1.Bronze)
	}
	processed := make(map[interface{}]struct{})
	AthleteList := make([]interface{}, 0)
	for _, uid := range AthleteList1 {
		if _, ok := processed[uid]; ok {
			continue
		}
		AthleteList = append(AthleteList, uid)
		processed[uid] = struct{}{}
	}
	for _, athlete1 := range AthleteList {
		GoldCount := 0
		SilverCount := 0
		BronzeCount := 0
		TotalCount := 0
		for _, data1 := range data {
			if data1.GoldNoc == Noc && data1.Gold == athlete1 {
				tip := 0
				for _, data3 := range SportsUtil.NormalList {
					if data1.Gold == data3 {
						tip += 1
					}
				}
				if tip == 0 {
					GoldCount += 1
					TotalCount += 1
				}
			}
			if data1.SilverNoc == Noc && data1.Silver == athlete1 {
				tip := 0
				for _, data3 := range SportsUtil.NormalList {
					if data1.Gold == data3 {
						tip += 1
					}
				}
				if tip == 0 {
					SilverCount += 1
					TotalCount += 1
				}
			}
			if data1.BronzeNoc == Noc && data1.Bronze == athlete1 {
				tip := 0
				for _, data3 := range SportsUtil.NormalList {
					if data1.Gold == data3 {
						tip += 1
					}
				}
				if tip == 0 {
					BronzeCount += 1
					TotalCount += 1
				}
			}
		}
		if TotalCount > 0 {
			GOldSort1 := SportsUtil.AllAthleteSort{
				Athlete: athlete1,
				Gold:    GoldCount,
				Silver:  SilverCount,
				Bronze:  BronzeCount,
				Total:   TotalCount,
			}
			GoldSort = append(GoldSort, GOldSort1)
		}
	}
	sort.Sort(GoldSort)
	for i := 0; i < len(GoldSort); i++ {
		data2 := map[string]interface{}{
			"athlete_name": GoldSort[i].Athlete,
			"gold":         GoldSort[i].Gold,
			"silver":       GoldSort[i].Silver,
			"bronze":       GoldSort[i].Bronze,
			"total":        GoldSort[i].Total,
		}
		TotalData = append(TotalData, data2)
	}
	if len(TotalData) == 0 {
		return 0, NotMedal
	} else {
		return TotalData, nil
	}
}

//国家强势项目

func (lc Country) TopSportsInfo(year string, countryId string) (interface{}, error) {
	data := mysql.TopSportsInfoVerity(year, countryId)
	TotalData := make([]interface{}, 0)
	if data.ID == 0 {
		return 0, NotMedal
	}
	data1 := map[string]interface{}{
		"sport":       data.Sport,
		"description": data.Description,
	}
	TotalData = append(TotalData, data1)
	return TotalData, nil
}

//国家风云人物

func (lc Country) TopAthleteInfo(countryId string) (interface{}, error) {
	Noc, data := mysql.TopAthleteInfoVerity(countryId)
	if len(data) == 0 {
		return 0, NotMedal
	}
	TotalData := make([]interface{}, 0)
	GoldSort := SportsUtil.AllAthleteSortSlice{}
	AthleteList1 := make([]interface{}, 0)
	for _, athlete1 := range data {
		AthleteList1 = append(AthleteList1, athlete1.Gold)
		AthleteList1 = append(AthleteList1, athlete1.Silver)
		AthleteList1 = append(AthleteList1, athlete1.Bronze)
	}
	processed := make(map[interface{}]struct{})
	AthleteList := make([]interface{}, 0)
	for _, uid := range AthleteList1 {
		if _, ok := processed[uid]; ok {
			continue
		}
		AthleteList = append(AthleteList, uid)
		processed[uid] = struct{}{}
	}
	for _, athlete1 := range AthleteList {
		GoldCount := 0
		SilverCount := 0
		BronzeCount := 0
		TotalCount := 0
		for _, data1 := range data {
			if data1.GoldNoc == Noc && data1.Gold == athlete1 {
				tip := 0
				for _, data3 := range SportsUtil.NormalList {
					if data1.Gold == data3 {
						tip += 1
					}
				}
				if tip == 0 {
					GoldCount += 1
					TotalCount += 1
				}
			}
			if data1.SilverNoc == Noc && data1.Silver == athlete1 {
				tip := 0
				for _, data3 := range SportsUtil.NormalList {
					if data1.Gold == data3 {
						tip += 1
					}
				}
				if tip == 0 {
					SilverCount += 1
					TotalCount += 1
				}
			}
			if data1.BronzeNoc == Noc && data1.Bronze == athlete1 {
				tip := 0
				for _, data3 := range SportsUtil.NormalList {
					if data1.Gold == data3 {
						tip += 1
					}
				}
				if tip == 0 {
					BronzeCount += 1
					TotalCount += 1
				}
			}
		}
		if TotalCount > 0 {
			GOldSort1 := SportsUtil.AllAthleteSort{
				Athlete: athlete1,
				Gold:    GoldCount,
				Silver:  SilverCount,
				Bronze:  BronzeCount,
				Total:   TotalCount,
			}
			GoldSort = append(GoldSort, GOldSort1)
		}
	}
	sort.Sort(GoldSort)
	data3 := mysql.TopAthleteVerity(GoldSort[0].Athlete)
	data2 := map[string]interface{}{
		"name":        data3.Name,
		"information": data3.Information,
		"prize":       data3.Prize,
		"picture":     data3.Picture,
	}
	TotalData = append(TotalData, data2)
	if len(TotalData) == 0 {
		return 0, NotMedal
	} else {
		return TotalData, nil
	}
}

//国家奖牌预测

func (lc Country) PredictInfo(countryId string) (interface{}, error) {
	data := mysql.PredictInfoVerity(countryId)
	if data.ID == 0 {
		return 0, NotMedal
	}
	TotalData := make([]interface{}, 0)
	data1 := map[string]interface{}{
		"country_name": data.CountryName,
		"gold_count":   data.GoldCount,
		"silver_count": data.SilverCount,
		"bronze_count": data.BronzeCount,
	}
	TotalData = append(TotalData, data1)
	return TotalData, nil
}
