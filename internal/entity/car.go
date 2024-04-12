package entity

type Car struct {
	Id           int    `gorm: "type:int; primary_key"`
	Model        string `gorm: "type: varchar(20)"`
	Year         int    `gorm: "type:int"`
	Color        string `gorm: "type:varchar(20)"`
	LicensePlate string `gorm: "type:varchar(20)"`
	IsRented     bool   `gorm: "type:bool"`
}

// Business validations
