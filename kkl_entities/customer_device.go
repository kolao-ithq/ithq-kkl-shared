package kkl_entities

import "gorm.io/gorm"

type CustomerDevice struct {
	gorm.Model
	CustomerID    uint   `gorm:"column:customer_id"`
	DeviceID      string `gorm:"column:device_id"`
	FirebaseToken string `gorm:"column:firebase_token"`
	DeviceName    string `gorm:"column:device_name"`
	DeviceType    string `gorm:"column:device_type"`
	LatLng        string `gorm:"column:lat_lng"`
	Active        bool   `gorm:"column:active"`
}

func (CustomerDevice) TableName() string { return "customer_devices" }
