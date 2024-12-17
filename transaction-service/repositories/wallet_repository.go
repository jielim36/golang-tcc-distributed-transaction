package repositories

import (
	"tcc-based-microservice-transaction/transaction-service/models"

	"gorm.io/gorm"
)

type WalletRepository struct {
	DB *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{
		DB: db,
	}
}

func (r *WalletRepository) Create(wallet *models.Wallet) error {
	return r.DB.Create(wallet).Error
}

func (r *WalletRepository) GetById(id uint) (*models.Wallet, error) {
	wallet := &models.Wallet{}
	err := r.DB.First(wallet, id).Error
	return wallet, err
}

func (r *WalletRepository) GetByUserId(id uint) (*models.Wallet, error) {
	wallet := &models.Wallet{}
	err := r.DB.Where("user_id = ?", id).First(wallet).Error
	return wallet, err
}
