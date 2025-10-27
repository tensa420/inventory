package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"inventory/internal/api"
	"inventory/internal/repository/repository"
	"inventory/internal/usecase"
	in "inventory/pkg/pb"
)

func main() {
	ctx := context.Background()

	lis, err := net.Listen("tcp", ":50062")
	if err != nil {
		log.Printf("Failed to listen: %v", err)
	}
	defer func() {
		err = lis.Close()
		if err != nil {
			log.Printf("Failed to close listener: %v", err)
		}
	}()

	_ = godotenv.Load(".env")

	dbURI := os.Getenv("MONGO_URI")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Printf("failed to connect to database: %v\n", err)
		return
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("failed to ping database: %v\n", err)
		return
	}

	db := client.Database("inventory")

	inventoryRepo := repository.NewInventoryRepository(db)
	inventoryUsecase := usecase.NewInventoryUsecase(inventoryRepo)
	inventoryServer := api.NewInventoryServer(inventoryUsecase)

	s := grpc.NewServer()
	in.RegisterInventoryServiceServer(s, inventoryServer)

	reflection.Register(s)

	go func() {
		err := s.Serve(lis)
		if err != nil {
			log.Printf("Failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 0)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	s.GracefulStop()
	log.Println("Server was stopped")
}
