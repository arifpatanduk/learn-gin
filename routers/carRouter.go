package routers

import (
	"learn-gin/controllers"

	"github.com/gin-gonic/gin"

	_ "learn-gin/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Car API
// @version 1.0
// @description This is a sample service for REST API in Golang
// @termOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email arifpatanduk1@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controllers.CreateCar)
	router.PUT("/cars/:carID", controllers.UpdateCar)
	router.DELETE("/cars/:carID", controllers.DeleteCar)
	router.GET("/cars/:carID", controllers.GetCar)
	router.GET("/cars/", controllers.GetAllCar)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	return router
}