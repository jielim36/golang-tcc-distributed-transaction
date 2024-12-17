package tcc_event

import (
	"context"
	"log"
	pb "tcc-based-microservice-transaction/transaction-service/rpc/tcc_rpc/proto"
)

type TCC_EVENT string

const (
	TCC_EVENT_UNKNOWN              TCC_EVENT = "UNKNOWN"
	TCC_EVENT_UPGRADE_SUBSCRIPTION TCC_EVENT = "UPGRADE_SUBSCRIPTION"
	TCC_EVENT_DEPOSIT              TCC_EVENT = "DEPOSIT"
	TCC_EVENT_WITHDRAW             TCC_EVENT = "WITHDRAW"
	TCC_EVENT_CREATE_WALLET        TCC_EVENT = "CREATE_WALLET"
)

type TCC_Template interface {
	Try(ctx context.Context, req *pb.TryRequest) (*pb.TryResponse, error)
	Confirm(ctx context.Context, req *pb.ConfirmRequest) (*pb.ConfirmResponse, error)
	Cancel(ctx context.Context, req *pb.CancelRequest) (*pb.CancelResponse, error)
}

type TCCExecuter struct {
	TCCFactory *TCCFactory
}

func NewTCCExecuter(
	tccFactory *TCCFactory,
) *TCCExecuter {
	return &TCCExecuter{
		TCCFactory: tccFactory,
	}
}

var EventMapping = map[pb.TCC_EVENTS]TCC_EVENT{
	pb.TCC_EVENTS_UNKNOWN:              TCC_EVENT_UNKNOWN,
	pb.TCC_EVENTS_UPGRADE_SUBSCRIPTION: TCC_EVENT_UPGRADE_SUBSCRIPTION,
	pb.TCC_EVENTS_DEPOSIT:              TCC_EVENT_DEPOSIT,
	pb.TCC_EVENTS_WITHDRAW:             TCC_EVENT_WITHDRAW,
	pb.TCC_EVENTS_CREATE_WALLET:        TCC_EVENT_CREATE_WALLET,
}

func (e *TCCExecuter) Try(ctx context.Context, req *pb.TryRequest) (*pb.TryResponse, error) {
	event := TCC_EVENT(EventMapping[req.Event.EventType])
	log.Printf("TCC Executer Try Event: %v", event)
	return e.TCCFactory.GetTCC(event).Try(ctx, req)
}

func (e *TCCExecuter) Confirm(ctx context.Context, req *pb.ConfirmRequest) (*pb.ConfirmResponse, error) {
	event := TCC_EVENT(EventMapping[req.Event.EventType])
	log.Printf("TCC Executer Confirm Event: %v", event)
	return e.TCCFactory.GetTCC(event).Confirm(ctx, req)
}

func (e *TCCExecuter) Cancel(ctx context.Context, req *pb.CancelRequest) (*pb.CancelResponse, error) {
	event := TCC_EVENT(EventMapping[req.Event.EventType])
	log.Printf("TCC Executer Cancel Event: %v", event)
	return e.TCCFactory.GetTCC(event).Cancel(ctx, req)
}
