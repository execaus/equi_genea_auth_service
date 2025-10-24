package main

import (
	authpb "equi_genea_auth_service/internal/pb/api/auth"

	"equi_genea_auth_service/config"
	"equi_genea_auth_service/internal/app"
	"equi_genea_auth_service/internal/service"

	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	grpcServer := grpc.NewServer()
	accountService := service.NewAuthService()
	handler := app.NewAccountHandler(accountService)

	authpb.RegisterAuthServiceServer(grpcServer, handler)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go runServer(grpcServer, cfg)

	<-stop
	fmt.Println("\nShutting down gracefully...")

	grpcServer.GracefulStop()

	fmt.Println("Server stopped")
}

func runServer(grpcServer *grpc.Server, cfg *config.Config) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
	fmt.Printf("Starting gRPC server on port %s...\n", cfg.Server.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
