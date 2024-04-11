package usecases

import "github.com/jpcairesf/kafka-saga/car-rental/internal/entity"

func GetRents() []entity.Rent {
	// Get all rents
	return nil
}

func CreateRent(rent entity.Rent) entity.Rent {
	// Create a rent
	return rent
}

func DeleteRent(id int) {
}

// Application validations
