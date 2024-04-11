package usecases

import "github.com/jpcairesf/kafka-saga/car-rental/internal/entity"

func GetCars() []entity.Car {
	// Create a cars
	return nil
}

func CreateCar(car entity.Car) entity.Car {
	// Create a car
	return car
}

func DeleteCar(id int) {
}

// Application validations
