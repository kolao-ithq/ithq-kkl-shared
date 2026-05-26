package entities

import "gorm.io/gorm"

type SellRatio struct {
	gorm.Model
	DigitType uint    `gorm:"uniqueIndex"`
	Ratio     float64
}
