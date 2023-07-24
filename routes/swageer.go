package routes

import (
	_ "WinterOlympic/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//注册Swagger文档路由

func InitSwagger(r *gin.Engine) *gin.Engine {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
