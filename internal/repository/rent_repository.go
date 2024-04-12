package repository

import "github.com/jpcairesf/car-rental/internal/entity"

type RentRepository interface {
	Save(rent entity.Rent) (entity.Rent, error)
	FindAll() ([]entity.Rent, error)
	Delete(id int) error
}
