package controller

import (
	"tcc-based-microservice-transaction/user-service/services"
	rpc_client "tcc-based-microservice-transaction/user-service/services/rpc"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	userService    *services.UserService
	transactionRPC *rpc_client.TransactionRPCClient
}

func NewTransactionController(
	userService *services.UserService,
	transactionRPC *rpc_client.TransactionRPCClient,
) *TransactionController {
	return &TransactionController{
		userService:    userService,
		transactionRPC: transactionRPC,
	}
}

func (t *TransactionController) TryCreateWallet(ctx *gin.Context) {
	userId := ctx.Param("user_id")
	tryResp, err := t.transactionRPC.TryCreateWallet(userId)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, tryResp)
}

func (t *TransactionController) ConfirmCreateWallet(ctx *gin.Context) {
	// userId := ctx.Param("user_id")
	// confirmResp, err := t.transactionRPC.ConfirmCreateWallet(userId)
	// if err != nil {
	// 	ctx.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }

	ctx.JSON(200, "confirmResp")
}
