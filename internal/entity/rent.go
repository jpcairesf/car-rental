package entity

import "time"

type Rent struct {
	Id         int       `gorm: "type:int; primary_key"`
	CarId      int       `gorm: "type:int; foreign_key"`
	UserCode   string    `gorm: "type:varchar(20)"`
	StartDate  time.Time `gorm: "type:datetime"`
	EndDate    time.Time `gorm: "type:datetime"`
	DailyValue float64   `gorm: "type:float"`
}

// Business validations
