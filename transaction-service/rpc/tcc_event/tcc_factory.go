package tcc_event

import "tcc-based-microservice-transaction/transaction-service/services"

type TCCFactory struct {
	EventManager  map[TCC_EVENT]TCC_Template
	walletService *services.WalletService
}

func NewTCCFactory(
	walletService *services.WalletService,
) *TCCFactory {
	tccFactory := &TCCFactory{
		walletService: walletService,
	}
	tccFactory.InitEventManager()
	return tccFactory
}

func (f *TCCFactory) InitEventManager() map[TCC_EVENT]TCC_Template {
	f.EventManager = make(map[TCC_EVENT]TCC_Template)
	f.RegisterTCC(TCC_EVENT_UNKNOWN, NewUnknownTCC())
	f.RegisterTCC(TCC_EVENT_UPGRADE_SUBSCRIPTION, NewUpgradeSubscriptionPlanTCC())
	f.RegisterTCC(TCC_EVENT_DEPOSIT, NewDepositTCC())
	f.RegisterTCC(TCC_EVENT_WITHDRAW, NewCreateWalletTCC(f.walletService))
	f.RegisterTCC(TCC_EVENT_CREATE_WALLET, NewCreateWalletTCC(f.walletService))
	return f.EventManager
}

func (f *TCCFactory) RegisterTCC(TCC_EVENT TCC_EVENT, tcc TCC_Template) {
	f.EventManager[TCC_EVENT] = tcc
}

func (f *TCCFactory) GetTCC(TCC_EVENT TCC_EVENT) TCC_Template {
	tcc, ok := f.EventManager[TCC_EVENT]
	if !ok {
		return f.EventManager[TCC_EVENT_UNKNOWN]
	}
	return tcc
}
