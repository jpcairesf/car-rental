package entity

import (
	"fmt"
	"time"
)

func ExampleRent_Validate() {
	r := Rent{
		StartDate: time.Now(),
		EndDate:   time.Now().Add(-time.Hour),
	}
	err := r.Validate()
	fmt.Println(err)
	// Output: End date must be after start date
}
