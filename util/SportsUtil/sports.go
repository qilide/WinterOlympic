package SportsUtil

// 切片排序
// 按照 Sports.Gold 从大到小排序

type GoldTotalSort struct {
	Sport interface{} `json:"sport"`
	Gold  int         `json:"gold"`
	Total int         `json:"total"`
}

type GoldSortSlice []GoldTotalSort

func (s GoldSortSlice) Len() int           { return len(s) }
func (s GoldSortSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s GoldSortSlice) Less(i, j int) bool { return s[i].Gold > s[j].Gold }

type TotalSortSlice []GoldTotalSort

func (s TotalSortSlice) Len() int           { return len(s) }
func (s TotalSortSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s TotalSortSlice) Less(i, j int) bool { return s[i].Total > s[j].Total }

type AllMedalSort struct {
	Sport  interface{} `json:"sport"`
	Gold   int         `json:"gold"`
	Silver int         `json:"silver"`
	Bronze int         `json:"bronze"`
	Total  int         `json:"total"`
}

type AllMedalSortSlice []AllMedalSort

func (s AllMedalSortSlice) Len() int           { return len(s) }
func (s AllMedalSortSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s AllMedalSortSlice) Less(i, j int) bool { return s[i].Gold > s[j].Gold }

type AllAthleteSort struct {
	Athlete interface{} `json:"athlete"`
	Gold    int         `json:"gold"`
	Silver  int         `json:"silver"`
	Bronze  int         `json:"bronze"`
	Total   int         `json:"total"`
}

type AllAthleteSortSlice []AllAthleteSort

func (s AllAthleteSortSlice) Len() int           { return len(s) }
func (s AllAthleteSortSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s AllAthleteSortSlice) Less(i, j int) bool { return s[i].Gold > s[j].Gold }

type AllAthleteTotalSlice []AllAthleteSort

func (s AllAthleteTotalSlice) Len() int      { return len(s) }
func (s AllAthleteTotalSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s AllAthleteTotalSlice) Less(i, j int) bool {
	if s[i].Total > s[j].Total {
		return s[i].Total > s[j].Total
	} else {
		return s[i].Gold > s[j].Gold
	}
}
