package main

import (
	"Oauth/bean"
	"Oauth/constants"
	"Oauth/database"
	"Oauth/middlewares"
	"log"

	_ "Oauth/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title OAuth API
// @description OAuth API Muhammad Alif Saddid
// @schemes http https
// @BasePath  /
func main() {
	godotenv.Load()
	database.Migrate()

	postgres, err := database.GetPostgresClient()
	if err != nil {
		log.Println("ERROR | database client error")
		return
	}

	redis := database.GetRedisClient()

	r := gin.Default()
	factory := bean.NewFactory(postgres, redis)
	m := ginmetrics.GetMonitor()
	m.Use(r)

	// controllers
	authController := factory.GetOauthController()
	roleController := factory.GetRoleController()

	// oauth
	r.POST(constants.TOKEN_PATH, authController.Token)
	r.POST(constants.RESOURCE_PATH, middlewares.AuthMiddleware(factory.GetJWTService()), authController.Resource)

	// role
	r.POST(constants.ROLE_BASE_PATH, roleController.CreateRole)
	r.GET(constants.ROLE_BASE_PATH, roleController.GetAllRole)
	r.GET(constants.ROLE_WITH_ID, roleController.GetRoleById)
	r.PUT(constants.ROLE_WITH_ID, roleController.UpdateRole)
	r.DELETE(constants.ROLE_WITH_ID, roleController.DeleteRole)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.Run(":28250")
}
