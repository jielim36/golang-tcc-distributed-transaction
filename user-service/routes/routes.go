package routes

import (
	"tcc-based-microservice-transaction/user-service/controller"

	"github.com/gin-gonic/gin"
)

type AppRoutes struct {
	r                     *gin.Engine
	userController        *controller.UserController
	transactionController *controller.TransactionController
}

func NewAppRoutes(
	r *gin.Engine,
	userController *controller.UserController,
	transactionController *controller.TransactionController,
) *AppRoutes {
	return &AppRoutes{
		r:                     r,
		userController:        userController,
		transactionController: transactionController,
	}
}

func (a *AppRoutes) RegisterRoutes() {

	v1 := a.r.Group("/api/v1")
	{
		// User routes
		userGroup := v1.Group("/users")
		{
			userGroup.POST("", a.userController.CreateUser)
			userGroup.PUT("", a.userController.UpdateUser)
			userGroup.GET("/:id", a.userController.GetUser)
			userGroup.DELETE("/:id", a.userController.DeleteUser)
		}

		// Transaction routes
		transactionGroup := v1.Group("/transactions")
		{
			tryGroup := transactionGroup.Group("/try")
			{
				// POST /api/v1/transactions/try/create-wallet
				tryGroup.POST("/create-wallet", a.transactionController.TryCreateWallet)
			}

			confirmGroup := transactionGroup.Group("/confirm")
			{
				// POST /api/v1/transactions/confirm/create-wallet
				confirmGroup.POST("/create-wallet", a.transactionController.ConfirmCreateWallet)
			}
		}
	}

}
