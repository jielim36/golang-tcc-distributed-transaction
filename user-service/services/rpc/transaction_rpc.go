package rpc_client

import (
	"context"
	pb "tcc-based-microservice-transaction/transaction-service/rpc/tcc_rpc/proto"
)

type TransactionRPCClient struct {
	rpc pb.TransactionServiceClient
}

func NewTransactionRPCClient(
	rpc pb.TransactionServiceClient,
) *TransactionRPCClient {
	return &TransactionRPCClient{
		rpc: rpc,
	}
}

func (t *TransactionRPCClient) TryCreateWallet(userId string) (*pb.TryResponse, error) {
	// ensure transaction status is try
	event := &pb.Event{
		EventType: pb.TCC_EVENTS_CREATE_WALLET,
		EventId:   userId,
	}

	tryRequest := &pb.TryRequest{
		Event: event,
		Data:  nil,
	}

	return t.rpc.Try(context.Background(), tryRequest)
}
