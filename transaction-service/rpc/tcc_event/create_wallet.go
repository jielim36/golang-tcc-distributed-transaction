package tcc_event

import (
	"context"
	"tcc-based-microservice-transaction/transaction-service/models"
	pb "tcc-based-microservice-transaction/transaction-service/rpc/tcc_rpc/proto"
	"tcc-based-microservice-transaction/transaction-service/services"

	"github.com/shopspring/decimal"
)

type CreateWalletTCC struct {
	walletService *services.WalletService
}

func NewCreateWalletTCC(walletService *services.WalletService) *CreateWalletTCC {
	return &CreateWalletTCC{
		walletService: walletService,
	}
}

func (s *CreateWalletTCC) Try(ctx context.Context, req *pb.TryRequest) (*pb.TryResponse, error) {
	// check if wallet already exists
	wallet := &models.Wallet{
		Balance:       decimal.NewFromInt(20),
		FrozenBalance: decimal.NewFromInt(0),
		Status:        string(models.WalletStatusInactive),
	}

	if err := s.walletService.CreateWallet(wallet); err != nil {
		tryResp := &pb.TryResponse{
			Success: false,
			Message: "Failed to create wallet",
		}
		return tryResp, err
	}

	tryResp := &pb.TryResponse{
		Success: true,
		Message: "Created wallet",
		Response: &pb.DataTemplate{
			Wallet: &pb.Wallet{
				Id:            uint64(wallet.ID),
				Balance:       wallet.Balance.InexactFloat64(),
				FrozenBalance: wallet.FrozenBalance.InexactFloat64(),
				Status:        wallet.Status,
				CreatedAt:     wallet.CreatedAt.String(),
				UpdatedAt:     wallet.UpdatedAt.String(),
			},
		},
	}
	return tryResp, nil
}

func (s *CreateWalletTCC) Confirm(ctx context.Context, req *pb.ConfirmRequest) (*pb.ConfirmResponse, error) {

	return nil, nil
}

func (s *CreateWalletTCC) Cancel(ctx context.Context, req *pb.CancelRequest) (*pb.CancelResponse, error) {
	// Business logic for "Cancel"
	return nil, nil
}
