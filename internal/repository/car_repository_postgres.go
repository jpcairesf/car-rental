package repository

import (
	"github.com/jpcairesf/car-rental/internal/entity"
	"gorm.io/gorm"
)

type CarRepositoryPostgres struct {
	DB *gorm.DB
}

func NewCarRepositoryPostgres(DB *gorm.DB) CarRepository {
	return &CarRepositoryPostgres{DB: DB}
}

func (*CarRepositoryPostgres) Save(car entity.Car) (entity.Car, error) {
	return car, nil
}

func (*CarRepositoryPostgres) FindAll() ([]entity.Car, error) {
	return nil, nil
}

func (*CarRepositoryPostgres) Delete(id int) error {
	return nil
}
