package rpc_server

import (
	"context"
	"tcc-based-microservice-transaction/transaction-service/rpc/tcc_event"
	pb "tcc-based-microservice-transaction/transaction-service/rpc/tcc_rpc/proto"
	"tcc-based-microservice-transaction/transaction-service/services"
)

type TransactionRPC struct {
	pb.UnimplementedTransactionServiceServer
	transactionService *services.TransactionService
	tccExecuter        *tcc_event.TCCExecuter
}

func NewTransactionRPCServer(
	transactionService *services.TransactionService,
	walletServices *services.WalletService,
) *TransactionRPC {
	factory := tcc_event.NewTCCFactory(walletServices)
	return &TransactionRPC{
		tccExecuter:        tcc_event.NewTCCExecuter(factory),
		transactionService: transactionService,
	}
}

// Try implements the first phase of TCC - attempting the transaction
func (rpc *TransactionRPC) Try(ctx context.Context, req *pb.TryRequest) (*pb.TryResponse, error) {
	return rpc.tccExecuter.Try(ctx, req)
}

// Confirm implements the second phase of TCC - confirming the transaction
func (rpc *TransactionRPC) Confirm(ctx context.Context, req *pb.ConfirmRequest) (*pb.ConfirmResponse, error) {
	return rpc.tccExecuter.Confirm(ctx, req)
}

// Cancel implements the third phase of TCC - cancelling the transaction
func (s *TransactionRPC) Cancel(ctx context.Context, req *pb.CancelRequest) (*pb.CancelResponse, error) {
	return s.tccExecuter.Cancel(ctx, req)
}
