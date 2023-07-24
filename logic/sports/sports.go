package sports

import (
	"WinterOlympic/dao/mysql"
	"WinterOlympic/util/SportsUtil"
	"errors"
	"sort"
	"strconv"
)

type Sports struct {
}

var (
	NotData = errors.New("NotData") // 数据不存在
	NotPage = errors.New("NotPage") // 页数不存在
)

//国家奖牌数量排行

func (s Sports) Country(year string) (interface{}, error) {
	data1, data2 := mysql.CountryMedalVerity(year)
	TotalData := make([]interface{}, 0)
	for i := 0; i < len(data1); i++ {
		data := map[string]interface{}{
			"country1": data1[i].Country,
			"gold":     data1[i].Gold,
			"country2": data2[i].Country,
			"total":    data2[i].Total,
		}
		TotalData = append(TotalData, data)
	}
	if len(TotalData) == 0 {
		return 0, NotData
	} else {
		return TotalData, nil
	}
}

//国家奖牌数量排行(详细)

func (s Sports) CountryDetail(year string, page string) (interface{}, error) {
	data1 := mysql.CountryMedalDetailVerity(year)
	TotalData := make([]interface{}, 0)
	for _, data2 := range data1 {
		data := map[string]interface{}{
			"country": data2.Country,
			"gold":    data2.Gold,
			"silver":  data2.Silver,
			"bronze":  data2.Bronze,
			"total":   data2.Total,
		}
		TotalData = append(TotalData, data)
	}
	if len(TotalData) == 0 {
		return 0, NotData
	}
	page1, err := strconv.Atoi(page)
	if err != nil || len(TotalData) < page1*10-10 {
		return 0, NotPage
	}
	TotalData = TotalData[(page1-1)*10 : page1*10]
	return TotalData, nil
}

//项目奖牌数量排行

func (s Sports) SportMedal(year string) (interface{}, error) {
	data1 := mysql.SportsMedalVerity(year)
	TotalData := make([]interface{}, 0)
	SportList1 := make([]interface{}, 0)
	GoldSort := SportsUtil.GoldSortSlice{}
	TotalSort := SportsUtil.TotalSortSlice{}
	for _, sport1 := range data1 {
		SportList1 = append(SportList1, sport1.Sport)
	}
	processed := make(map[interface{}]struct{})
	SportList := make([]interface{}, 0)
	for _, uid := range SportList1 {
		if _, ok := processed[uid]; ok {
			continue
		}
		SportList = append(SportList, uid)
		processed[uid] = struct{}{}
	}
	for _, sport1 := range SportList {
		GoldCount := 0
		TotalCount := 0
		for _, sport2 := range data1 {
			if sport1 == sport2.Sport {
				GoldCount += sport2.Gold
				TotalCount += sport2.Total
			}
		}
		GOldSort1 := SportsUtil.GoldTotalSort{
			Sport: sport1,
			Gold:  GoldCount,
			Total: TotalCount,
		}
		TotalSort1 := SportsUtil.GoldTotalSort{
			Sport: sport1,
			Gold:  GoldCount,
			Total: TotalCount,
		}
		GoldSort = append(GoldSort, GOldSort1)
		TotalSort = append(TotalSort, TotalSort1)
	}
	sort.Sort(GoldSort)
	sort.Sort(TotalSort)
	for i := 0; i < len(GoldSort); i++ {
		data := map[string]interface{}{
			"sport1":      TotalSort[i].Sport,
			"total_count": TotalSort[i].Total,
			"sport2":      GoldSort[i].Sport,
			"gold_count":  GoldSort[i].Gold,
		}
		TotalData = append(TotalData, data)
	}
	if len(TotalData) == 0 {
		return 0, NotData
	} else {
		return TotalData[:5], nil
	}
}

//项目奖牌数量排行(详细)

func (s Sports) SportMedalDetail(year string) (interface{}, error) {
	data1 := mysql.SportsMedalVerity(year)
	TotalData := make([]interface{}, 0)
	SportList1 := make([]interface{}, 0)
	AllMedalSort := SportsUtil.AllMedalSortSlice{}
	for _, sport1 := range data1 {
		SportList1 = append(SportList1, sport1.Sport)
	}
	processed := make(map[interface{}]struct{})
	SportList := make([]interface{}, 0)
	for _, uid := range SportList1 {
		if _, ok := processed[uid]; ok {
			continue
		}
		SportList = append(SportList, uid)
		processed[uid] = struct{}{}
	}
	for _, sport1 := range SportList {
		GoldCount := 0
		SilverCount := 0
		BronzeCount := 0
		TotalCount := 0
		for _, sport2 := range data1 {
			if sport1 == sport2.Sport {
				GoldCount += sport2.Gold
				SilverCount += sport2.Silver
				BronzeCount += sport2.Bronze
				TotalCount += sport2.Total
			}
		}
		GOldSort1 := SportsUtil.AllMedalSort{
			Sport:  sport1,
			Gold:   GoldCount,
			Silver: SilverCount,
			Bronze: BronzeCount,
			Total:  TotalCount,
		}
		AllMedalSort = append(AllMedalSort, GOldSort1)
	}
	sort.Sort(AllMedalSort)
	for i := 0; i < len(AllMedalSort); i++ {
		data := map[string]interface{}{
			"sport":        AllMedalSort[i].Sport,
			"gold_count":   AllMedalSort[i].Gold,
			"silver_count": AllMedalSort[i].Silver,
			"bronze_count": AllMedalSort[i].Bronze,
			"total_count":  AllMedalSort[i].Total,
		}
		TotalData = append(TotalData, data)
	}
	if len(TotalData) == 0 {
		return 0, NotData
	} else {
		return TotalData, nil
	}
}

//国家信息

func (s Sports) CountryMsg(year string) (interface{}, error) {
	data1 := mysql.CountryHumanVerity(year)
	TotalData := make([]interface{}, 0)
	for _, data2 := range data1 {
		data := map[string]interface{}{
			"id":          data2.ID,
			"noc":         data2.Noc,
			"country":     data2.Country,
			"country_map": data2.CountryMap,
			"total":       data2.Total,
		}
		TotalData = append(TotalData, data)
	}
	if len(TotalData) == 0 {
		return 0, NotData
	} else {
		return TotalData, nil
	}
}

//运动员奖牌数量排行

func (s Sports) AthleteMsg(year string) (interface{}, error) {
	data1 := mysql.AthleteMedalVerity(year)
	TotalData := make([]interface{}, 0)
	GoldSort := SportsUtil.AllAthleteSortSlice{}
	TotalSort := SportsUtil.AllAthleteTotalSlice{}
	AthleteList1 := make([]interface{}, 0)
	for _, athlete1 := range data1 {
		AthleteList1 = append(AthleteList1, athlete1.Gold)
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
		tip := 1
		for _, nor := range SportsUtil.NormalList {
			if nor == athlete1 {
				tip = 0
			}
		}
		if tip == 0 {
			continue
		}
		GoldCount := 0
		TotalCount := 0
		for _, data2 := range data1 {
			if data2.Gold == athlete1 {
				GoldCount += 1
				TotalCount += 1
			}
			if data2.Silver == athlete1 {
				TotalCount += 1
			}
			if data2.Bronze == athlete1 {
				TotalCount += 1
			}
		}
		GOldSort1 := SportsUtil.AllAthleteSort{
			Athlete: athlete1,
			Gold:    GoldCount,
		}
		TotalSort1 := SportsUtil.AllAthleteSort{
			Athlete: athlete1,
			Total:   TotalCount,
		}
		GoldSort = append(GoldSort, GOldSort1)
		TotalSort = append(TotalSort, TotalSort1)
	}
	sort.Sort(GoldSort)
	sort.Sort(TotalSort)
	for i := 0; i < len(GoldSort); i++ {
		data := map[string]interface{}{
			"athlete1":    GoldSort[i].Athlete,
			"gold_count":  GoldSort[i].Gold,
			"athlete2":    TotalSort[i].Athlete,
			"total_count": TotalSort[i].Total,
		}
		TotalData = append(TotalData, data)
	}
	if len(TotalData) == 0 {
		return 0, NotData
	} else {
		return TotalData[:5], nil
	}
}

//运动员奖牌数量排行(详细)

func (s Sports) AthleteDetailMsg(year string, page string) (interface{}, error) {
	data1 := mysql.AthleteMedalVerity(year)
	TotalData := make([]interface{}, 0)
	GoldSort := SportsUtil.AllAthleteSortSlice{}
	AthleteList1 := make([]interface{}, 0)
	for _, athlete1 := range data1 {
		AthleteList1 = append(AthleteList1, athlete1.Gold)
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
		tip := 1
		for _, nor := range SportsUtil.NormalList {
			if nor == athlete1 {
				tip = 0
			}
		}
		if tip == 0 {
			continue
		}
		GoldCount := 0
		SilverCount := 0
		BronzeCount := 0
		TotalCount := 0
		for _, data2 := range data1 {
			if data2.Gold == athlete1 {
				GoldCount += 1
				TotalCount += 1
			}
			if data2.Silver == athlete1 {
				SilverCount += 1
				TotalCount += 1
			}
			if data2.Bronze == athlete1 {
				BronzeCount += 1
				TotalCount += 1
			}
		}
		GOldSort1 := SportsUtil.AllAthleteSort{
			Athlete: athlete1,
			Gold:    GoldCount,
			Silver:  SilverCount,
			Bronze:  BronzeCount,
			Total:   TotalCount,
		}
		GoldSort = append(GoldSort, GOldSort1)
	}
	sort.Sort(GoldSort)
	for i := 0; i < len(GoldSort); i++ {
		data := map[string]interface{}{
			"athlete":      GoldSort[i].Athlete,
			"gold_count":   GoldSort[i].Gold,
			"silver_count": GoldSort[i].Silver,
			"bronze_count": GoldSort[i].Bronze,
			"total_count":  GoldSort[i].Total,
		}
		TotalData = append(TotalData, data)
	}
	if len(TotalData) == 0 {
		return 0, NotData
	}
	page1, err := strconv.Atoi(page)
	if err != nil || len(TotalData) < page1*10-10 {
		return 0, NotPage
	}
	TotalData = TotalData[(page1-1)*10 : page1*10]
	return TotalData, nil
}

//中国项目奖牌数量排行

func (s Sports) ChinaMedal(year string) (interface{}, error) {
	data1 := mysql.ChinaMedalVerity(year)
	TotalData := make([]interface{}, 0)
	SportList1 := make([]interface{}, 0)
	GoldSort := SportsUtil.GoldSortSlice{}
	TotalSort := SportsUtil.TotalSortSlice{}
	for _, sport1 := range data1 {
		SportList1 = append(SportList1, sport1.Sport)
	}
	processed := make(map[interface{}]struct{})
	SportList := make([]interface{}, 0)
	for _, uid := range SportList1 {
		if _, ok := processed[uid]; ok {
			continue
		}
		SportList = append(SportList, uid)
		processed[uid] = struct{}{}
	}
	for _, sport1 := range SportList {
		GoldCount := 0
		TotalCount := 0
		for _, sport2 := range data1 {
			if sport1 == sport2.Sport {
				GoldCount += sport2.Gold
				TotalCount += sport2.Total
			}
		}
		GOldSort1 := SportsUtil.GoldTotalSort{
			Sport: sport1,
			Gold:  GoldCount,
			Total: TotalCount,
		}
		TotalSort1 := SportsUtil.GoldTotalSort{
			Sport: sport1,
			Gold:  GoldCount,
			Total: TotalCount,
		}
		GoldSort = append(GoldSort, GOldSort1)
		TotalSort = append(TotalSort, TotalSort1)
	}
	sort.Sort(GoldSort)
	sort.Sort(TotalSort)
	for i := 0; i < len(GoldSort); i++ {
		data := map[string]interface{}{
			"sport1":      TotalSort[i].Sport,
			"total_count": TotalSort[i].Total,
			"sport2":      GoldSort[i].Sport,
			"gold_count":  GoldSort[i].Gold,
		}
		TotalData = append(TotalData, data)
	}
	if len(TotalData) == 0 {
		return 0, NotData
	} else {
		return TotalData[:5], nil
	}
}

//中国项目奖牌数量排行(详细)

func (s Sports) ChinaMedalDetail(year string) (interface{}, error) {
	data1 := mysql.ChinaMedalVerity(year)
	TotalData := make([]interface{}, 0)
	SportList1 := make([]interface{}, 0)
	AllMedalSort := SportsUtil.AllMedalSortSlice{}
	for _, sport1 := range data1 {
		SportList1 = append(SportList1, sport1.Sport)
	}
	processed := make(map[interface{}]struct{})
	SportList := make([]interface{}, 0)
	for _, uid := range SportList1 {
		if _, ok := processed[uid]; ok {
			continue
		}
		SportList = append(SportList, uid)
		processed[uid] = struct{}{}
	}
	for _, sport1 := range SportList {
		GoldCount := 0
		SilverCount := 0
		BronzeCount := 0
		TotalCount := 0
		for _, sport2 := range data1 {
			if sport1 == sport2.Sport {
				GoldCount += sport2.Gold
				SilverCount += sport2.Silver
				BronzeCount += sport2.Bronze
				TotalCount += sport2.Total
			}
		}
		GOldSort1 := SportsUtil.AllMedalSort{
			Sport:  sport1,
			Gold:   GoldCount,
			Silver: SilverCount,
			Bronze: BronzeCount,
			Total:  TotalCount,
		}
		AllMedalSort = append(AllMedalSort, GOldSort1)
	}
	sort.Sort(AllMedalSort)
	for i := 0; i < len(AllMedalSort); i++ {
		data := map[string]interface{}{
			"sport":        AllMedalSort[i].Sport,
			"gold_count":   AllMedalSort[i].Gold,
			"silver_count": AllMedalSort[i].Silver,
			"bronze_count": AllMedalSort[i].Bronze,
			"total_count":  AllMedalSort[i].Total,
		}
		TotalData = append(TotalData, data)
	}
	if len(TotalData) == 0 {
		return 0, NotData
	} else {
		return TotalData, nil
	}
}

//项目精彩时刻

func (s Sports) Moment(SportName string) (interface{}, error) {
	data1 := mysql.MomentVerity(SportName)
	TotalData := make([]interface{}, 0)
	for i := 0; i < len(data1); i++ {
		data := map[string]interface{}{
			"name":      data1[i].Name,
			"describe1": data1[i].Describe1,
			"video1":    data1[i].Video1,
			"img1":      data1[i].Img1,
			"describe2": data1[i].Describe2,
			"video2":    data1[i].Video2,
			"img2":      data1[i].Img2,
			"describe3": data1[i].Describe3,
			"video3":    data1[i].Video3,
			"img3":      data1[i].Img3,
		}
		TotalData = append(TotalData, data)
	}
	if len(TotalData) == 0 {
		return 0, NotData
	} else {
		return TotalData, nil
	}
}
