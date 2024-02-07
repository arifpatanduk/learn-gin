package routers

import (
	"learn-gin/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controllers.CreateCar)
	router.PUT("/cars/:carID", controllers.UpdateCar)
	router.DELETE("/cars/:carID", controllers.DeleteCar)
	router.GET("/cars/:carID", controllers.GetCar)
	router.GET("/cars/", controllers.GetAllCar)


	return router
}