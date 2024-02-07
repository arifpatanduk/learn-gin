package controllers

import (
	"fmt"
	"learn-gin/database"
	"learn-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCar godoc
// @Summary Create new car
// @Description Create new car
// @Tag cars
// @Accept json
// @Produce json
// @Param models.Car body true "create car"
// @Success 200 {object} models.Car
// @Router /cars [post]
func CreateCar(ctx *gin.Context) {
	db := database.GetDB()

	var input models.Car

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	carInput := models.Car{Brand: input.Brand, Price: input.Price, Model: input.Model}
	db.Create(&carInput)

	ctx.JSON(http.StatusCreated, gin.H{
		"car": carInput,
	})
}

// UpdateCars godoc
// @Summary Update car indetified by the given Id
// @Description Update the car corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param models.Car body true "update car"
// @Success 200 {object} models.Car
// @Router /cars/{carID} [put]
func UpdateCar(ctx *gin.Context) {
    carID := ctx.Param("carID")

    db := database.GetDB()
    var carUpdate models.Car

    if err := db.First(&carUpdate, carID).Error; err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
        return
    }

    if err := ctx.ShouldBindJSON(&carUpdate); err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

    if err := db.Save(&carUpdate).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update car"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "message": fmt.Sprintf("Car with id %v has been successfully updated", carID),
    })
}

// GetAllCars godoc
// @Summary Get all car
// @Description Get details of all car
// @Tags cars
// @Accept json
// @Produce json
// @Success 200 {object} models.Car
// @Router /cars [get]
func GetAllCar(ctx *gin.Context) {
	db := database.GetDB()

	var cars []models.Car
	err := db.Find(&cars).Error

	if err != nil {
		fmt.Println("Error getting car datas:", err.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{"data": cars})
}

// GetOneCar godoc
// @Summary Get details for a given Id
// @Description Get details of car corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param Id path int true "ID of the car"
// @Success 200 {object} models.Car
// @Router /cars/{carID} [get]
func GetCar (ctx *gin.Context) {
	db := database.GetDB()

	var carOne models.Car
	carID := ctx.Param("carID")

	err := db.First(&carOne, "Id = ?", carID).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car": carOne,
	})
}

// DeleteCars godoc
// @Summary Delete car identified by the given Id
// @Description Delete the order corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param Id path int true "ID of the car to be deleted"
// @Success 204 "No Content"
// @Router /cars/{carID} [delete]
func DeleteCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	
	db:= database.GetDB()
	var carDelete models.Car

	err := db.First(&carDelete, "Id = ?", carID).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&carDelete)

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been seccessfully deleted", carID),
	})
}