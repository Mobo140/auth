package main

import (
	"context"
	"flag"
	"log"
	"net"

	userAPI "github.com/Mobo140/microservices/auth/internal/api/user"
	"github.com/Mobo140/microservices/auth/internal/config"
	"github.com/Mobo140/microservices/auth/internal/config/env"
	userRepository "github.com/Mobo140/microservices/auth/internal/repository/user"
	userService "github.com/Mobo140/microservices/auth/internal/service/user"
	desc "github.com/Mobo140/microservices/auth/pkg/user_v1"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

type server struct {
	desc.UnimplementedUserV1Server
	pool *pgxpool.Pool
}

func main() {
	flag.Parse()
	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to load grpc config: %v", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to load pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.New(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer pool.Close()

	userRepo := userRepository.NewRepository(pool)
	userSrv := userService.NewService(userRepo)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, userAPI.NewImplementation(userSrv))

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
