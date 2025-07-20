package helpers

import (
	"context"
	"database/sql"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"testing"

	"cex-core-api/app/internal/storages/postgres/sqlc"

	_ "github.com/lib/pq"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	postgresOnce      sync.Once
	postgresContainer testcontainers.Container
	pool              *pgxpool.Pool
	queries           *sqlc.Queries
)

func GetDatabaseContainer(t *testing.T) (testcontainers.Container, *pgxpool.Pool, *sqlc.Queries) {
	t.Helper()

	postgresOnce.Do(func() {
		ctx := context.Background()

		container, err := startPostgresContainer(ctx)
		require.NoError(t, err, "failed to start container")

		postgresContainer = container

		dsn, err := buildPostgresDSN(ctx, container)
		require.NoError(t, err, "failed to get container DSN")

		pool, err = connectPgxPool(ctx, dsn)
		require.NoError(t, err, "failed to connect pgx pool")

		// Apply migrations
		require.NoError(t, applyMigrations(dsn), "failed to apply migrations")

		queries = sqlc.New(pool)
	})

	return postgresContainer, pool, queries
}

func startPostgresContainer(ctx context.Context) (testcontainers.Container, error) {

	const (
		image   = "postgres:17.2"
		port    = "5432/tcp"
		waitLog = "database system is ready to accept connections"
	)

	req := testcontainers.ContainerRequest{
		Image:        image,
		ExposedPorts: []string{port},
		Env: map[string]string{
			"POSTGRES_USER":     "test",
			"POSTGRES_PASSWORD": "test",
			"POSTGRES_DB":       "test",
		},
		WaitingFor: wait.ForAll(
			wait.ForListeningPort(port),
			wait.ForLog(waitLog),
		),
	}

	return testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}

func buildPostgresDSN(ctx context.Context, container testcontainers.Container) (string, error) {
	host, err := container.Host(ctx)
	if err != nil {
		return "", err
	}

	port, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("postgres://test:test@%s:%s/test?sslmode=disable", host, port.Port()), nil
}

func connectPgxPool(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}
	return pgxpool.NewWithConfig(ctx, cfg)
}

func applyMigrations(dsn string) error {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	_, filename, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(filename), "../../../")
	migrationsDir := filepath.Join(projectRoot, "migrations")

	goose.SetLogger(goose.NopLogger())
	return goose.Up(db, migrationsDir)
}

func TruncateAllTables(t *testing.T, db *pgxpool.Pool) {
	rows, err := db.Query(context.Background(), `
		SELECT tablename FROM pg_tables WHERE schemaname = 'public';
	`)
	require.NoError(t, err)
	defer rows.Close()

	var tableNames []string
	for rows.Next() {
		var table string
		require.NoError(t, rows.Scan(&table))
		tableNames = append(tableNames, fmt.Sprintf(`"%s"`, table))
	}

	if len(tableNames) == 0 {
		return
	}

	query := fmt.Sprintf("TRUNCATE %s RESTART IDENTITY CASCADE;", strings.Join(tableNames, ", "))
	_, err = db.Exec(context.Background(), query)
	require.NoError(t, err)
}
