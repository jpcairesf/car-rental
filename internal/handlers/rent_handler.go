package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
)

type RentInput struct {
	CarId      int       `json:"car_id" binding:"required"`
	UserCode   string    `json:"user_code" binding:"required" maxLength:"20"`
	StartDate  time.Time `json:"start_date" binding:"required"`
	EndDate    time.Time `json:"end_date" binding:"required"`
	DailyValue float64   `json:"daily_value" binding:"required"`
}

type RentOutput struct {
	Id         int       `json:"id"`
	CarId      int       `json:"car_id"`
	UserCode   string    `json:"user_code"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	DailyValue float64   `json:"daily_value"`
}

func GetRents(c *gin.Context) []RentOutput {
	c.JSON(200, gin.H{
		"message": "GetRents",
	})
	return nil
}

func CreateRent(c *gin.Context, input *RentInput) {
	c.JSON(201, gin.H{
		"message": "CreateRent",
	})
}

func DeleteRent(c *gin.Context, id int) {
	c.JSON(204, gin.H{
		"message": "DeleteRent",
	})
}
