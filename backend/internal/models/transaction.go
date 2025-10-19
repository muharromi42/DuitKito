package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"` // "income" atau "expense"
}
