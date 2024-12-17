package models

import "gorm.io/gorm"

// Transaction GORM Model
type Transaction struct {
	gorm.Model
	WalletID    uint `gorm:"index:idx_wallet_transaction"`
	Amount      float64
	Description string
	EventID     string `gorm:"unique_index"`
	Status      string
	Type        string
	Metadata    string `gorm:"type:text"`
}
