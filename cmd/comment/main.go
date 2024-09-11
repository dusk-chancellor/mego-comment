package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	adapter "github.com/dusk-chancellor/mego-comment/internal/adapters/grpc"
	"github.com/dusk-chancellor/mego-comment/internal/config"
	"github.com/dusk-chancellor/mego-comment/internal/database"
	"github.com/dusk-chancellor/mego-comment/internal/repositories"
	"github.com/dusk-chancellor/mego-comment/internal/services"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: "",
		DB:       0,
	})

	err = rdb.Ping(context.Background()).Err()
	if err != nil {
		log.Fatalf("Failed to connect to redis: %v", err)
	}
	log.Println("Connected to redis")

	commentRepo := repositories.NewCommentRepository(db, rdb)
	commentService := services.NewCommentService(commentRepo)

	l, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	adapter.RegisterGRPC(grpcServer, commentService, cfg)
	go log.Fatal(grpcServer.Serve(l))

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	grpcServer.GracefulStop()
}
