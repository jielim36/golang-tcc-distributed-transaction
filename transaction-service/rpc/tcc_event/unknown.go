package tcc_event

import (
	"context"
	pb "tcc-based-microservice-transaction/transaction-service/rpc/tcc_rpc/proto"
)

type UnknownTCC struct {
}

func NewUnknownTCC() *UnknownTCC {
	return &UnknownTCC{}
}

func (s *UnknownTCC) Try(ctx context.Context, req *pb.TryRequest) (*pb.TryResponse, error) {
	UnknownTCCResponse := &pb.TryResponse{
		Success: false,
		Message: "Unknown transaction type",
	}
	return UnknownTCCResponse, nil
}

func (s *UnknownTCC) Confirm(ctx context.Context, req *pb.ConfirmRequest) (*pb.ConfirmResponse, error) {
	UnknownTCCResponse := &pb.ConfirmResponse{
		Success: false,
		Message: "Unknown transaction type",
	}
	return UnknownTCCResponse, nil
}

func (s *UnknownTCC) Cancel(ctx context.Context, req *pb.CancelRequest) (*pb.CancelResponse, error) {
	UnknownTCCResponse := &pb.CancelResponse{
		Success: false,
		Message: "Unknown transaction type",
	}
	return UnknownTCCResponse, nil
}
