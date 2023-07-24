package country

import (
	"WinterOlympic/controller/response"
	"WinterOlympic/logic/country"
	"github.com/gin-gonic/gin"
)

// MedalInformation 国家奖牌信息展示
// @Summary 国家奖牌信息展示
// @Description 国家奖牌信息展示
// @Tags 国家奖牌信息展示
// @Accept application/json
// @Produce application/json
// @Param object body CountryBinder false "国家id/年份"
// @Security ApiKeyAuth
// @Success 200 {object}  response.SuccessLogin "成功"
// @failure 401 {object}  response.Information "该国家在此届冬奥会未获得奖牌"
// @failure 402 {object}  response.Information "请输入正确的信息"
// @Router /country/medal_information [POST]
func MedalInformation(c *gin.Context) {
	Year := c.PostForm("year")
	CountryId := c.PostForm("country_id")
	if len(Year) == 0 || len(CountryId) == 0 {
		var cb CountryBinder
		if err := c.ShouldBind(&cb); err != nil {
			response.Json(c, 402, "请输入正确的信息", 0)
			return
		} else {
			Year = cb.Year
			CountryId = cb.CountryId
		}
	}
	var la country.Country
	data, err := la.MedalInfo(Year, CountryId)
	if err == country.NotMedal {
		response.Json(c, 401, "该国家在此届冬奥会未获得奖牌", data)
		return
	} else {
		response.Json(c, 200, "成功", data)
		return
	}
}

// MedalSports 国家详细项目获得奖牌
// @Summary 国家详细项目获得奖牌
// @Description 国家详细项目获得奖牌
// @Tags 国家详细项目获得奖牌
// @Accept application/json
// @Produce application/json
// @Param object body CountryBinder false "国家id/年份"
// @Security ApiKeyAuth
// @Success 200 {object}  response.SuccessLogin "成功"
// @failure 401 {object}  response.Information "该国家在此届冬奥会未获得奖牌"
// @failure 402 {object}  response.Information "请输入正确的信息"
// @Router /country/medal_sports [POST]
func MedalSports(c *gin.Context) {
	Year := c.PostForm("year")
	CountryId := c.PostForm("country_id")
	if len(Year) == 0 || len(CountryId) == 0 {
		var cb CountryBinder
		if err := c.ShouldBind(&cb); err != nil {
			response.Json(c, 402, "请输入正确的信息", 0)
			return
		} else {
			Year = cb.Year
			CountryId = cb.CountryId
		}
	}
	var la country.Country
	data, err := la.SportsInfo(Year, CountryId)
	if err == country.NotMedal {
		response.Json(c, 401, "该国家在此届冬奥会未获得奖牌", data)
		return
	} else {
		response.Json(c, 200, "成功", data)
		return
	}
}

// MedalAthlete 国家详细运动员获得奖牌
// @Summary 国家详细运动员获得奖牌
// @Description 国家详细运动员获得奖牌
// @Tags 国家详细运动员获得奖牌
// @Accept application/json
// @Produce application/json
// @Param object body CountryBinder false "国家id/年份"
// @Security ApiKeyAuth
// @Success 200 {object}  response.SuccessLogin "成功"
// @failure 401 {object}  response.Information "该国家运动员在此届冬奥会未获得奖牌"
// @failure 402 {object}  response.Information "请输入正确的信息"
// @Router /country/medal_athlete [POST]
func MedalAthlete(c *gin.Context) {
	Year := c.PostForm("year")
	CountryId := c.PostForm("country_id")
	if len(Year) == 0 || len(CountryId) == 0 {
		var cb CountryBinder
		if err := c.ShouldBind(&cb); err != nil {
			response.Json(c, 402, "请输入正确的信息", 0)
			return
		} else {
			Year = cb.Year
			CountryId = cb.CountryId
		}
	}
	var la country.Country
	data, err := la.AthleteInfo(Year, CountryId)
	if err == country.NotMedal {
		response.Json(c, 401, "该国家运动员在此届冬奥会未获得奖牌", data)
		return
	} else {
		response.Json(c, 200, "成功", data)
		return
	}
}

// TopSports 国家强势项目
// @Summary 国家强势项目
// @Description 国家强势项目
// @Tags 国家强势项目
// @Accept application/json
// @Produce application/json
// @Param object body CountryBinder false "国家id/年份"
// @Security ApiKeyAuth
// @Success 200 {object}  response.SuccessLogin "成功"
// @failure 401 {object}  response.Information "该国家在此届冬奥会未获得奖牌"
// @failure 402 {object}  response.Information "请输入正确的信息"
// @Router /country/top_sports [POST]
func TopSports(c *gin.Context) {
	Year := c.PostForm("year")
	CountryId := c.PostForm("country_id")
	if len(Year) == 0 || len(CountryId) == 0 {
		var cb CountryBinder
		if err := c.ShouldBind(&cb); err != nil {
			response.Json(c, 402, "请输入正确的信息", 0)
			return
		} else {
			Year = cb.Year
			CountryId = cb.CountryId
		}
	}
	var la country.Country
	data, err := la.TopSportsInfo(Year, CountryId)
	if err == country.NotMedal {
		response.Json(c, 401, "该国家在此届冬奥会未获得奖牌", data)
		return
	} else {
		response.Json(c, 200, "成功", data)
		return
	}
}

// TopAthlete 国家风云人物
// @Summary 国家风云人物
// @Description 国家风云人物
// @Tags 国家风云人物
// @Accept application/json
// @Produce application/json
// @Param object body CountryIdBinder false "年份"
// @Security ApiKeyAuth
// @Success 200 {object}  response.SuccessLogin "成功"
// @failure 401 {object}  response.Information "该国家运动员在前八届冬奥会未获得奖牌"
// @failure 402 {object}  response.Information "请输入正确的信息"
// @Router /country/top_athlete [POST]
func TopAthlete(c *gin.Context) {
	CountryId := c.PostForm("country_id")
	if len(CountryId) == 0 {
		var cb CountryIdBinder
		if err := c.ShouldBind(&cb); err != nil {
			response.Json(c, 402, "请输入正确的信息", 0)
			return
		} else {
			CountryId = cb.CountryId
		}
	}
	var la country.Country
	data, err := la.TopAthleteInfo(CountryId)
	if err == country.NotMedal {
		response.Json(c, 401, "该国家运动员在前八届冬奥会未获得奖牌", data)
		return
	} else {
		response.Json(c, 200, "成功", data)
		return
	}
}

// Predict 国家奖牌预测
// @Summary 国家奖牌预测
// @Description 国家奖牌预测
// @Tags 国家奖牌预测
// @Accept application/json
// @Produce application/json
// @Param object body CountryIdBinder false "国家id"
// @Security ApiKeyAuth
// @Success 200 {object}  response.SuccessLogin "成功"
// @failure 401 {object}  response.Information "该国家在前八届冬奥会未获得奖牌,故无法预测"
// @failure 402 {object}  response.Information "请输入正确的信息"
// @Router /country/predict [POST]
func Predict(c *gin.Context) {
	CountryId := c.PostForm("country_id")
	if len(CountryId) == 0 {
		var cb CountryIdBinder
		if err := c.ShouldBind(&cb); err != nil {
			response.Json(c, 402, "请输入正确的信息", 0)
			return
		} else {
			CountryId = cb.CountryId
		}
	}
	var la country.Country
	data, err := la.PredictInfo(CountryId)
	if err == country.NotMedal {
		response.Json(c, 401, "该国家在前八届冬奥会未获得奖牌,故无法预测", data)
		return
	} else {
		response.Json(c, 200, "成功", data)
		return
	}
}
