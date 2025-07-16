package models

type Product struct {
	Name        string  `json:"name" gorm:"not null"`
	Price       float64 `json:"price" gorm:"not null"`
	Description string  `json:"description"`
}