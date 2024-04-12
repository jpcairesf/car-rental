package repository

import (
	"github.com/jpcairesf/car-rental/internal/entity"
	"gorm.io/gorm"
)

type RentRepositoryPostgres struct {
	DB *gorm.DB
}

func NewRentRepositoryPostgres(DB *gorm.DB) RentRepository {
	return &RentRepositoryPostgres{DB: DB}
}

func (*RentRepositoryPostgres) Save(rent entity.Rent) (entity.Rent, error) {
	return rent, nil
}

func (*RentRepositoryPostgres) FindAll() ([]entity.Rent, error) {
	return nil, nil
}

func (*RentRepositoryPostgres) Delete(id int) error {
	return nil
}
