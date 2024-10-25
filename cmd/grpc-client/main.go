package main

import (
	"context"
	"log"
	"time"

	desc "github.com/Mobo140/microservices/auth/pkg/user_v1"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:8080"
	userID  = 1
)

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	c := desc.NewUserV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &desc.GetRequest{Id: userID})
	if err != nil {
		log.Fatalf("failed to get user by id: %v", err)
	}

	log.Printf(color.RedString("User info:\n"), color.GreenString("%+v", r.GetInfo()))
}
