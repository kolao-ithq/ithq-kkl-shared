package kkl_entities

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	CreatedBy         *string    `gorm:"column:created_by;type:varchar(20)"`
	UpdatedBy         *time.Time `gorm:"column:updated_by"`
	Phone             string     `gorm:"type:longtext"`
	FirstName         string     `gorm:"column:first_name"`
	LastName          string     `gorm:"column:last_name;type:varchar(150)"`
	Gender            string     `gorm:"type:varchar(20)"`
	Birthday          string     `gorm:"type:varchar(30)"`
	Address           string     `gorm:"type:varchar(100)"`
	Email             string     `gorm:"column:email;type:varchar(100)"`
	OneSignalPlayerID string     `gorm:"column:one_signal_player_id;type:longtext"`
	Image             string     `gorm:"type:longtext"`
	InviteCode        string     `gorm:"column:invite_code;type:longtext"`
	InviteCodeActive  string     `gorm:"column:invite_code_active;type:longtext"`
	ActiveStatus      bool       `gorm:"column:active_status"`
	FirstBuy          bool       `gorm:"column:first_buy"`
	FirstBuyRealMoney bool       `gorm:"column:first_buy_real_money"`
	WalletID          string     `gorm:"column:wallet_id;type:longtext"`
	FromApp           *string    `gorm:"column:from_app;type:longtext"`
	ProfileImage      *string    `gorm:"column:profile_image;type:longtext"`
}

func (Customer) TableName() string { return "customers" }
