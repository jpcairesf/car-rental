package entity

type Car struct {
	Id           int    `gorm: "type:int; primaryKey"`
	Model        string `gorm: "type: varchar(20); notNull"`
	Year         int    `gorm: "type:int; notNull"`
	Color        string `gorm: "type:varchar(20); notNull"`
	LicensePlate string `gorm: "type:varchar(20); notNull; unique"`
}

// Business validations
