package services

import (
	"tcc-based-microservice-transaction/transaction-service/models"
	"tcc-based-microservice-transaction/transaction-service/repositories"
)

type WalletService struct {
	walletRepo *repositories.WalletRepository
}

func NewWalletService(walletRepo *repositories.WalletRepository) *WalletService {
	return &WalletService{
		walletRepo: walletRepo,
	}
}

func (s *WalletService) CreateWallet(wallet *models.Wallet) error {
	return s.walletRepo.Create(wallet)
}

func (s *WalletService) GetWalletByUserId(userId uint) (*models.Wallet, error) {
	return s.walletRepo.GetByUserId(userId)
}

func (s *WalletService) GetWalletById(id uint) (*models.Wallet, error) {
	return s.walletRepo.GetById(id)
}
