package repository

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ory/dockertest"
)

const (
	pgUsername = "eugen"
	pgPassword = "ur2qly1ini"
	pgDB       = "balance_db"
)

func SetupTestPgx() (*pgxpool.Pool, func(), error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, nil, fmt.Errorf("could not construct pool: %w", err)
	}
	resource, err := pool.Run("postgres", "latest", []string{
		fmt.Sprintf("POSTGRES_USER=%s", pgUsername),
		fmt.Sprintf("POSTGRESQL_PASSWORD=%s", pgPassword),
		fmt.Sprintf("POSTGRES_DB=%s", pgDB)})
	if err != nil {
		return nil, nil, fmt.Errorf("could not start resource: %w", err)
	}

	// cmd := exec.Command(
	// 	"flyway",
	// 	"-user="+pgUsername,
	// 	"-password="+pgPassword,
	// 	"-locations=filesystem:/home/yauhenishymanski/MyProject/myapp/migration",
	// 	"-url=jdbc:postgresql://localhost:5432/eugene",
	// 	"-connectRetries=10",
	// 	"-schemas=goschema",
	// 	"migrate",
	// )

	// err = cmd.Run()
	// if err != nil {
	// 	logrus.Fatalf("can't run migration: %s", err)
	// }

	dbURL := "postgres://eugen:ur2qly1ini@localhost:5432/balance_db"
	cfg, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse dbURL: %w", err)
	}
	dbpool, err := pgxpool.ConnectConfig(context.Background(), cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect pgxpool: %w", err)
	}
	cleanup := func() {
		dbpool.Close()
		pool.Purge(resource)
	}

	return dbpool, cleanup, nil
}

func TestMain(m *testing.M) {
	dbpool, cleanupPgx, err := SetupTestPgx()
	if err != nil {
		fmt.Println("Could not construct the pool: ", err)
		cleanupPgx()
		os.Exit(1)
	}
	rps = NewPsqlConnection(dbpool)

	exitVal := m.Run()
	cleanupPgx()
	os.Exit(exitVal)
}
