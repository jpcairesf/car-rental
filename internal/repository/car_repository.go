package repository

import "github.com/jpcairesf/car-rental/internal/entity"

type CarRepository interface {
	Save(car entity.Car) (entity.Car, error)
	FindAll() ([]entity.Car, error)
	Delete(id int) error
}
