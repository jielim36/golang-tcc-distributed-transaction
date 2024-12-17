package services

import (
	"log"
	"tcc-based-microservice-transaction/user-service/models"
	"tcc-based-microservice-transaction/user-service/repositories"
	rpc_client "tcc-based-microservice-transaction/user-service/services/rpc"
)

type UserService struct {
	userRepo       *repositories.UserRepository
	transactionRPC *rpc_client.TransactionRPCClient
}

func NewUserService(
	userRepo *repositories.UserRepository,
	transactionRPC *rpc_client.TransactionRPCClient,
) *UserService {
	return &UserService{
		userRepo:       userRepo,
		transactionRPC: transactionRPC,
	}
}

func (srv *UserService) CreateUser(user *models.User) (*models.User, error) {
	if err := srv.userRepo.Create(user); err != nil {
		return nil, err
	}

	// Create a transaction
	resp, err := srv.transactionRPC.TryCreateWallet(string(user.UserId))
	if err != nil {
		return nil, err
	}
	log.Printf("Transaction response: %v", resp.Message)

	user.AccountId = uint(resp.Response.Wallet.Id)
	srv.userRepo.Update(user)

	return user, nil

}

func (srv *UserService) GetUser(id uint) (*models.User, error) {
	user, err := srv.userRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (srv *UserService) UpdateUser(user *models.User) (*models.User, error) {
	if err := srv.userRepo.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (srv *UserService) DeleteUser(userId uint) error {
	return srv.userRepo.DeleteById(userId)
}
