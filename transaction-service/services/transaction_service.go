package services

import (
	"context"
	"tcc-based-microservice-transaction/transaction-service/models"
	"tcc-based-microservice-transaction/transaction-service/repositories"
)

type TransactionService struct {
	transactionRepo *repositories.TransactionRepository
}

func NewTransactionService(transactionRepo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
	}
}

func (srv *TransactionService) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	if err := srv.transactionRepo.Create(transaction); err != nil {
		return nil, err
	}

	return transaction, nil

}

func (srv *TransactionService) GetTransaction(id uint) (*models.Transaction, error) {
	transaction, err := srv.transactionRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (srv *TransactionService) UpdateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	if err := srv.transactionRepo.Update(transaction); err != nil {
		return nil, err
	}
	return transaction, nil
}

func (srv *TransactionService) DeleteTransaction(transactionId uint) error {
	return srv.transactionRepo.DeleteById(transactionId)
}

// FindTransactionByEventID finds a transaction by its event ID
func (r *TransactionService) FindTransactionByEventID(ctx context.Context, eventID string) (*models.Transaction, error) {
	transaction, err := r.transactionRepo.FindTransactionByEventID(ctx, eventID)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *TransactionService) UpdateTransactionStatus(
	ctx context.Context,
	eventID string,
	status models.TCCStatus,
) error {
	return r.transactionRepo.UpdateTransactionStatus(ctx, eventID, status)
}

func (r *TransactionService) UpdateWalletBalance(
	ctx context.Context,
	walletID uint,
	amount float64,
) error {
	return r.transactionRepo.UpdateWalletBalance(ctx, walletID, amount)
}
