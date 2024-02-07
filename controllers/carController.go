package controllers

import (
	"Gin/database"
	"Gin/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCars godoc
// @Summary Post details for a given Id
// @Description Post details of car corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param models.Car body models.Car true "create car"
// @Success 200 {object} models.Car
// @Router /cars/ [post]
func CreateCars(c *gin.Context) {
	var db = database.GetDB()

	var input models.Car
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("slebew", input.TypeCars)

	carinput := models.Car{Pemilik: input.Pemilik, Merk: input.Merk, Harga: input.Harga, TypeCars: input.TypeCars}

	err := db.Create(&carinput).Error

	c.JSON(http.StatusOK, gin.H{"data": carinput})
	fmt.Println("gndfgn", err)

}

// UpdateCars godoc
// @Summary Update car indetified by the given Id
// @Description Update the car corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "ID of the car to be updated"
// @Success 200 {object} models.Car
// @Router /cars/{id} [patch]
func UpdateCar(c *gin.Context) {
	var db = database.GetDB()
	id := c.Param("id")

	input := models.Car{}

	// var car models.Car = c.ShouldBindJSON(&input)
	err := db.First(&input).Where("id = ?", id).Updates(models.Car{
		Merk: input.Merk,
	}).Error
	// err := db.First(&car, "Id=?", c.Param("id")).Update(models.Car{
	// 	Pemilik:  car.Pemilik,
	// 	Merk:     car.Merk,
	// 	Harga:    car.Harga,
	// 	TypeCars: car.TypeCars,
	// }).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"updatedCar": input,
	})
}

// GetOneCar godoc
// @Summary Get details for a given Id
// @Description Get details of car corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param Id path int true "ID of the car"
// @Success 200 {object} models.Car
// @Router /cars/{Id} [get]
func GetOneCar(c *gin.Context) {

	var db = database.GetDB()

	var carOne models.Car

	err := db.First(&carOne, "Id= ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data One": carOne})

}

// GetAllCars godoc
// @Summary Get details
// @Description Get details of all car
// @Tags cars
// @Accept json
// @Produce json
// @Success 200 {object} models.Car
// @Router /orders [get]
func GetAllCars(c *gin.Context) {
	var db = database.GetDB()

	var cars []models.Car
	err := db.Find(&cars).Error

	if err != nil {
		fmt.Println("Error getting car datas :", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"data": cars})
}

// DeleteCars godoc
// @Summary Delete car identified by the given Id
// @Description Delete the order corresponding to the input id
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "ID of the car to be deleted"
// @Success 204 "No Content"
// @Router /cars/{id} [delete]
func DeleteCars(c *gin.Context) {
	var db = database.GetDB()

	var carDelete models.Car
	err := db.First(&carDelete, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&carDelete)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
