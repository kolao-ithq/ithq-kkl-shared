package kkl_entities

import "gorm.io/gorm"

type SellRatio struct {
	gorm.Model
	DigitType uint    `gorm:"uniqueIndex"`
	Ratio     float64 //core side
}
