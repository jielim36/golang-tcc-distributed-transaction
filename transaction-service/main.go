package main

import (
	"fmt"
	"log"
	"net"
	"tcc-based-microservice-transaction/transaction-service/models"
	"tcc-based-microservice-transaction/transaction-service/repositories"
	rpc_server "tcc-based-microservice-transaction/transaction-service/rpc/tcc_rpc"
	pb "tcc-based-microservice-transaction/transaction-service/rpc/tcc_rpc/proto"

	"tcc-based-microservice-transaction/transaction-service/services"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// db
	dsn := "user=postgres password=123456 dbname=golang_tcc_microservices host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	} else {
		fmt.Println("Connected to the PostgreSQL database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Transaction{}, &models.Wallet{})

	// Create gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Register repo and service
	transactionRepo := repositories.NewTransactionRepository(db)
	walletRepo := repositories.NewWalletRepository(db)

	transactionService := services.NewTransactionService(transactionRepo)
	walletService := services.NewWalletService(walletRepo)

	// Register rpc server
	transaction_rpc := rpc_server.NewTransactionRPCServer(transactionService, walletService)
	pb.RegisterTransactionServiceServer(grpcServer, transaction_rpc)

	log.Println("Starting gRPC server on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
