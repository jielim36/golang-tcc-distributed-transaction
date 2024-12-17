package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type WalletStatus string

const (
	WalletStatusActive   WalletStatus = "active"
	WalletStatusInactive WalletStatus = "inactive"
)

// Wallet GORM Model
type Wallet struct {
	gorm.Model
	Balance       decimal.Decimal `gorm:"type:decimal(20,4);not null;default:0"`
	FrozenBalance decimal.Decimal `gorm:"type:decimal(20,4);not null;default:0"`
	Status        string          `gorm:"type:varchar(50);default:'inactive'"`
}
