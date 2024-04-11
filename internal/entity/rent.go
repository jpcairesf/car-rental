package entity

import "time"

type Rent struct {
	Id         int
	CarId      int
	UserCode   string
	StartDate  time.Time
	EndDate    time.Time
	DailyValue float64
}

// Business validations
