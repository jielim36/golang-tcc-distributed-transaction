package tcc_event

import (
	"context"
	pb "tcc-based-microservice-transaction/transaction-service/rpc/tcc_rpc/proto"
)

type DepositTCC struct {
}

func NewDepositTCC() *DepositTCC {
	return &DepositTCC{}
}

func (s *DepositTCC) Try(ctx context.Context, req *pb.TryRequest) (*pb.TryResponse, error) {
	// Business logic for "Try"
	tryResp := &pb.TryResponse{
		Success: true,
		Message: "Upgrade subscription plan transaction created",
	}
	return tryResp, nil
}

func (s *DepositTCC) Confirm(ctx context.Context, req *pb.ConfirmRequest) (*pb.ConfirmResponse, error) {
	// Business logic for "Confirm"
	return nil, nil
}

func (s *DepositTCC) Cancel(ctx context.Context, req *pb.CancelRequest) (*pb.CancelResponse, error) {
	// Business logic for "Cancel"
	return nil, nil
}
