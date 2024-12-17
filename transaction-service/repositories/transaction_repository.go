package repositories

import (
	"context"
	"tcc-based-microservice-transaction/transaction-service/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db}
}

func (r *TransactionRepository) Create(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *TransactionRepository) GetById(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.First(&transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) Update(transaction *models.Transaction) error {
	return r.db.Save(transaction).Error
}

func (r *TransactionRepository) DeleteById(id uint) error {
	return r.db.Delete(&models.Transaction{}, id).Error
}

// FindTransactionByEventID finds a transaction by its event ID
func (r *TransactionRepository) FindTransactionByEventID(ctx context.Context, eventID string) (*models.Transaction, error) {
	var transaction models.Transaction
	result := r.db.WithContext(ctx).Where("event_id = ?", eventID).First(&transaction)
	return &transaction, result.Error
}

// UpdateTransactionStatus updates the status of a transaction
func (r *TransactionRepository) UpdateTransactionStatus(
	ctx context.Context,
	eventID string,
	status models.TCCStatus,
) error {
	result := r.db.WithContext(ctx).
		Model(&models.Transaction{}).
		Where("event_id = ?", eventID).
		Update("status", status)
	return result.Error
}

// UpdateWalletBalance updates the balance of a wallet
func (r *TransactionRepository) UpdateWalletBalance(
	ctx context.Context,
	walletID uint,
	amount float64,
) error {
	result := r.db.WithContext(ctx).
		Model(&models.Wallet{}).
		Where("id = ?", walletID).
		UpdateColumn("balance", gorm.Expr("balance + ?", amount))
	return result.Error
}
