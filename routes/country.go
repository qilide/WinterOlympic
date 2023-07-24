package routes

import (
	"WinterOlympic/controller/country"
	"WinterOlympic/controller/sports"
	"github.com/gin-gonic/gin"
)

//注册国家详细页面模块路由

func CountryGroupRoute(CountryGroup *gin.RouterGroup) {
	//国家奖牌信息展示
	CountryGroup.POST("/medal_information", country.MedalInformation)
	//国家详细项目获得奖牌
	CountryGroup.POST("/medal_sports", country.MedalSports)
	//国家详细运动员获得奖牌
	CountryGroup.POST("/medal_athlete", country.MedalAthlete)
	//国家强势项目
	CountryGroup.POST("/top_sports", country.TopSports)
	//国家风云人物
	CountryGroup.POST("/top_athlete", country.TopAthlete)
	//国家奖牌预测
	CountryGroup.POST("/predict", country.Predict)
	//项目精彩时刻
	CountryGroup.POST("/moments", sports.Moments)
}
