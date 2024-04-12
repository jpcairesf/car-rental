package entity

import (
	"errors"
	"time"
)

type Rent struct {
	Id         int       `gorm: "type:int; primaryKey"`
	CarId      int       `gorm: "type:int; foreignKey"`
	UserCode   string    `gorm: "type:varchar(20); notNull"`
	StartDate  time.Time `gorm: "type:datetime; notNull"`
	EndDate    time.Time `gorm: "type:datetime; notNull"`
	DailyValue float64   `gorm: "type:float; notNull"`
}

// Business validations
func (r *Rent) Validate() error {
	if r.EndDate.Before(r.StartDate) {
		return errors.New("End date must be after start date")
	}
	return nil
}
