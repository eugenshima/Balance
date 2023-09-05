// Package main is an entry point to this microservice
package main

import (
	"context"
	"fmt"
	"net"

	cfgrtn "github.com/eugenshima/Balance/internal/config"
	"github.com/eugenshima/Balance/internal/handlers"
	"github.com/eugenshima/Balance/internal/repository"
	"github.com/eugenshima/Balance/internal/service"
	proto "github.com/eugenshima/Balance/proto"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// NewDBPsql function provides Connection with PostgreSQL database
func NewDBPsql(env string) (*pgxpool.Pool, error) {
	// Initialization a connect configuration for a PostgreSQL using pgx driver
	config, err := pgxpool.ParseConfig(env)
	if err != nil {
		return nil, fmt.Errorf("error connection to PostgreSQL: %v", err)
	}

	// Establishing a new connection to a PostgreSQL database using the pgx driver
	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("error connection to PostgreSQL: %v", err)
	}
	// Output to console
	fmt.Println("Connected to PostgreSQL!")

	return pool, nil
}

// main function of our microservice
func main() {
	cfg, err := cfgrtn.NewConfig()
	if err != nil {
		fmt.Printf("Error extracting env variables: %v", err)
		return
	}
	pool, err := NewDBPsql(cfg.PgxDBAddr)
	if err != nil {
		logrus.WithFields(logrus.Fields{"PgxDBAddr: ": cfg.PgxDBAddr}).Errorf("NewDBPsql: %v", err)
	}
	pgx := repository.NewPsqlConnection(pool)
	srv := service.NewBalanceService(pgx)
	hndl := handlers.NewBalancehandler(srv, validator.New())
	lis, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		logrus.Fatalf("cannot create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	proto.RegisterBalanceServiceServer(serverRegistrar, hndl)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		logrus.Fatalf("cannot start server: %s", err)
	}
}
