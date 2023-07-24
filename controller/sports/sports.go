package sports

import (
	"WinterOlympic/controller/response"
	"WinterOlympic/logic/sports"
	"fmt"
	"github.com/gin-gonic/gin"
)

// CountryMedal 国家奖牌数量排行
// @Summary 国家奖牌数量排行
// @Description 国家奖牌数量排行
// @Tags 国家奖牌数量排行
// @Accept application/json
// @Produce application/json
// @Param year query string false "年份"
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "成功"
// @failure 401 {object}  response.Information "数据不存在"
// @Router /sports/country_medal [GET]
func CountryMedal(c *gin.Context) {
	year := c.Query("year")
	if year == "" || len(year) == 0 {
		year = "2022"
	}
	var sp sports.Sports
	data, err := sp.Country(year)
	if err == sports.NotData {
		response.Json(c, 401, "数据不存在", 0)
		return
	} else {
		response.Json(c, 200, "成功", data)
		return
	}
}

// CountryMedalDetail 国家奖牌数量排行(详细)
// @Summary 国家奖牌数量排行(详细)
// @Description 国家奖牌数量排行(详细)
// @Tags 国家奖牌数量排行(详细)
// @Accept application/json
// @Produce application/json
// @Param year query string false "年份"
// @Param page query string false "页数"
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "成功"
// @failure 401 {object}  response.Information "数据不存在"
// @failure 402 {object}  response.Information "页数不存在"
// @Router /sports/country_medal_detail [GET]
func CountryMedalDetail(c *gin.Context) {
	year := c.Query("year")
	page := c.Query("page")
	if year == "" || len(year) == 0 {
		year = "2022"
	}
	var sp sports.Sports
	data, err := sp.CountryDetail(year, page)
	fmt.Println(data)
	fmt.Println(err)
	if err == sports.NotData {
		response.Json(c, 401, "数据不存在", 0)
		return
	} else if err == sports.NotPage {
		response.Json(c, 402, "页数不存在", 0)
		return
	} else {
		response.Json(c, 200, "成功", data)
		return
	}
}

// SportsMedal 项目奖牌数量排行
// @Summary 项目奖牌数量排行
// @Description 项目奖牌数量排行
// @Tags 项目奖牌数量排行
// @Accept application/json
// @Produce application/json
// @Param year query string false "年份"
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "成功"
// @failure 401 {object}  response.Information "数据不存在"
// @Router /sports/sports_medal [GET]
func SportsMedal(c *gin.Context) {
	year := c.Query("year")
	if year == "" || len(year) == 0 {
		year = "2022"
	}
	var sp sports.Sports
	data, err := sp.SportMedal(year)
	if err == sports.NotData {
		response.Json(c, 401, "数据不存在", 0)
		return
	} else {
		response.Json(c, 200, "成功", data)
		return
	}
}

// SportsMedalDetail 项目奖牌数量排行(详细)
// @Summary 项目奖牌数量排行(详细)
// @Description 项目奖牌数量排行(详细)
// @Tags 项目奖牌数量排行(详细)
// @Accept application/json
// @Produce application/json
// @Param year query string false "年份"
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "成功"
// @failure 401 {object}  response.Information "数据不存在"
// @Router /sports/sports_medal_detail [GET]
func SportsMedalDetail(c *gin.Context) {
	year := c.Query("year")
	if year == "" || len(year) == 0 {
		year = "2022"
	}
	var sp sports.Sports
	data, err := sp.SportMedalDetail(year)
	if err == sports.NotData {
		response.Json(c, 401, "数据不存在", 0)
		return
	} else {
		response.Json(c, 200, "成功", data)
		return
	}
}

// CountryMessage 国家信息
// @Summary 国家信息
// @Description 国家信息
// @Tags 国家信息
// @Accept application/json
// @Produce application/json
// @Param year query string false "年份"
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "成功"
// @failure 401 {object}  response.Information "数据不存在"
// @Router /sports/country_message [GET]
func CountryMessage(c *gin.Context) {
	year := c.Query("year")
	if year == "" || len(year) == 0 {
		year = "2022"
	}
	var sp sports.Sports
	data, err := sp.CountryMsg(year)
	if err == sports.NotData {
		response.Json(c, 401, "数据不存在", 0)
		return
	} else {
		response.Json(c, 200, "成功", data)
		return
	}
}

// AthleteMedal 运动员奖牌数量排行
// @Summary 运动员奖牌数量排行
// @Description 运动员奖牌数量排行
// @Tags 运动员奖牌数量排行
// @Accept application/json
// @Produce application/json
// @Param year query string false "年份"
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "成功"
// @failure 401 {object}  response.Information "数据不存在"
// @Router /sports/athlete_medal [GET]
func AthleteMedal(c *gin.Context) {
	year := c.Query("year")
	if year == "" || len(year) == 0 {
		year = "2022"
	}
	var sp sports.Sports
	data, err := sp.AthleteMsg(year)
	if err == sports.NotData {
		response.Json(c, 401, "数据不存在", 0)
		return
	} else {
		response.Json(c, 200, "成功", data)
		return
	}
}

// AthleteMedalDetail 运动员奖牌数量排行(详细)
// @Summary 运动员奖牌数量排行(详细)
// @Description 运动员奖牌数量排行(详细)
// @Tags 运动员奖牌数量排行(详细)
// @Accept application/json
// @Produce application/json
// @Param year query string false "年份"
// @Param page query string false "页数"
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "成功"
// @failure 401 {object}  response.Information "数据不存在"
// @failure 402 {object}  response.Information "页数不存在"
// @Router /sports/athlete_medal_detail [GET]
func AthleteMedalDetail(c *gin.Context) {
	year := c.Query("year")
	page := c.Query("page")
	if year == "" || len(year) == 0 {
		year = "2022"
	}
	var sp sports.Sports
	data, err := sp.AthleteDetailMsg(year, page)
	if err == sports.NotData {
		response.Json(c, 401, "数据不存在", 0)
		return
	} else if err == sports.NotPage {
		response.Json(c, 402, "页数不存在", 0)
		return
	} else {
		response.Json(c, 200, "成功", data)
		return
	}
}

// ChinaMedal 中国项目奖牌数量排行
// @Summary 中国项目奖牌数量排行
// @Description 中国项目奖牌数量排行
// @Tags 中国项目奖牌数量排行
// @Accept application/json
// @Produce application/json
// @Param year query string false "年份"
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "成功"
// @failure 401 {object}  response.Information "数据不存在"
// @Router /sports/china_medal [GET]
func ChinaMedal(c *gin.Context) {
	year := c.Query("year")
	if year == "" || len(year) == 0 {
		year = "2022"
	}
	var sp sports.Sports
	data, err := sp.ChinaMedal(year)
	if err == sports.NotData {
		response.Json(c, 401, "数据不存在", 0)
		return
	} else {
		response.Json(c, 200, "成功", data)
		return
	}
}

// ChinaMedalDetail 中国项目奖牌数量排行(详细)
// @Summary 中国项目奖牌数量排行(详细)
// @Description 中国项目奖牌数量排行(详细)
// @Tags 中国项目奖牌数量排行(详细)
// @Accept application/json
// @Produce application/json
// @Param year query string false "年份"
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "成功"
// @failure 401 {object}  response.Information "数据不存在"
// @Router /sports/china_medal_detail [GET]
func ChinaMedalDetail(c *gin.Context) {
	year := c.Query("year")
	if year == "" || len(year) == 0 {
		year = "2022"
	}
	var sp sports.Sports
	data, err := sp.ChinaMedalDetail(year)
	if err == sports.NotData {
		response.Json(c, 401, "数据不存在", 0)
		return
	} else {
		response.Json(c, 200, "成功", data)
		return
	}
}

// Moments 项目精彩时刻
// @Summary 项目精彩时刻
// @Description 项目精彩时刻
// @Tags 项目精彩时刻
// @Accept application/json
// @Produce application/json
// @Param object body SportsName false "项目名称"
// @Security ApiKeyAuth
// @Success 200 {object}  response.SuccessLogin "成功"
// @failure 401 {object}  response.Information "未查询到此比赛的精彩视频"
// @failure 402 {object}  response.Information "请输入比赛项目名称"
// @Router /country/moments [POST]
func Moments(c *gin.Context) {
	SportName := c.PostForm("sport_name")
	if len(SportName) == 0 {
		var rec SportsName
		if err := c.ShouldBind(&rec); err != nil {
			response.Json(c, 402, "请输入比赛项目名称", 0)
			return
		} else {
			SportName = rec.SportName
		}
	}
	var sp sports.Sports
	data, err := sp.Moment(SportName)
	if err == sports.NotData {
		response.Json(c, 401, "未查询到此比赛的精彩视频", 0)
		return
	} else {
		response.Json(c, 200, "成功", data)
		return
	}
}
