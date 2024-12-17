package main

import (
	"fmt"
	"log"
	pb "tcc-based-microservice-transaction/transaction-service/rpc/tcc_rpc/proto"
	"tcc-based-microservice-transaction/user-service/controller"
	"tcc-based-microservice-transaction/user-service/models"
	"tcc-based-microservice-transaction/user-service/repositories"
	"tcc-based-microservice-transaction/user-service/routes"
	"tcc-based-microservice-transaction/user-service/services"
	rpc_client "tcc-based-microservice-transaction/user-service/services/rpc"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// postgresql db
	dsn := "user=postgres password=123456 dbname=golang_tcc_microservices host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	} else {
		fmt.Println("Connected to the PostgreSQL database")
	}

	// grpc connection
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to the gRPC server: %v", err)
	} else {
		fmt.Println("Connected to the gRPC server")
	}

	// Initialize gRPC client
	transactionRPCClient := rpc_client.NewTransactionRPCClient(pb.NewTransactionServiceClient(conn))

	// Migrate the schema
	db.AutoMigrate(&models.User{})

	// Initialize repository
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo, transactionRPCClient)
	userController := controller.NewUserController(userService)

	transactionController := controller.NewTransactionController(userService, transactionRPCClient)

	// Create Gin router
	r := gin.Default()
	appRoutes := routes.NewAppRoutes(
		r,
		userController,
		transactionController,
	)
	appRoutes.RegisterRoutes()

	// Start server
	log.Println("Starting web server on :9990")
	r.Run(":9990")
}
