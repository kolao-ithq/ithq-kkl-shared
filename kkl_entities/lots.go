package kkl_entities

import (
	"time"

	"gorm.io/gorm"
)

type Lots struct {
	gorm.Model
	Name          string    `gorm:"type:varchar(100)"`
	StartDatetime time.Time `gorm:"type:datetime"`
	EndDatetime   time.Time `gorm:"type:datetime"`
	LotID         string    `gorm:"type:varchar(25)"`
	LotDate       string    `gorm:"type:varchar(25)"`
	Status        int8      `gorm:"type:tinyint"`
}

type LotResults struct {
	gorm.Model
	Digit  int8   `gorm:"type:tinyint"`
	Result string `gorm:"type:varchar(10)"`
	LotID  string `gorm:"type:varchar(25)"`
}
