package handlers

import "github.com/gin-gonic/gin"

type CarInput struct {
	Model        string `json:"model" binding:"required" maxLength:"50"`
	Year         int    `json:"year" binding:"required"`
	Color        string `json:"color" binding:"required" maxLength:"20"`
	LicensePlate string `json:"license_plate" binding:"required" maxLength:"10"`
}

type CarOutput struct {
	Id           int    `json:"id"`
	Model        string `json:"model"`
	Year         int    `json:"year"`
	Color        string `json:"color"`
	LicensePlate string `json:"license_plate"`
}

func GetCars(c *gin.Context) []CarOutput {
	c.JSON(200, gin.H{
		"message": "GetCars",
	})
	return nil
}

func CreateCar(c *gin.Context, input *CarInput) {
	c.JSON(201, gin.H{
		"message": "CreateCar",
	})
}

func DeleteCar(c *gin.Context, id int) {
	c.JSON(204, gin.H{
		"message": "DeleteCar",
	})
}
