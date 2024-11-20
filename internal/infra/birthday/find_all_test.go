package birthday

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/guergeiro/discord-bots/pkg/domain/entity"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestFindAll(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	const dbName string = "birthdays"
	const dbUser string = "user"
	const dbPassword string = "password"

	postgresSql, err := filepath.Abs(
		filepath.Join("..", "..", "..", "assets", "postgres-entrypoint.sql"),
	)
	assert.NoError(t, err)

	postgresContainer, err := postgres.Run(context.Background(),
		"timescale/timescaledb:latest-pg14",
		postgres.WithInitScripts(postgresSql),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		postgres.WithSQLDriver("pgx"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second),
			wait.ForListeningPort("5432/tcp"),
		))
	assert.NoError(t, err)
	defer func() {
		err := testcontainers.TerminateContainer(postgresContainer)
		assert.NoError(t, err)
	}()
	err = postgresContainer.Snapshot(context.Background())
	assert.NoError(t, err)
	connStr, err := postgresContainer.ConnectionString(
		context.Background(),
		"sslmode=disable",
		"application_name=birthdays",
	)
	assert.NoError(t, err)

	const timelayout = "2006-01-02"

	t.Run("valid find all", func(t *testing.T) {
		t.Cleanup(func() {
			err := postgresContainer.Restore(context.Background())
			assert.NoError(t, err)
		})

		conn, err := pgxpool.New(context.Background(), connStr)
		assert.NoError(t, err)
		defer conn.Close()

		statement := `
			INSERT INTO birthdays (id, date) VALUES ($1, $2)
			ON CONFLICT (id, date)
			DO UPDATE SET date = EXCLUDED.date;
		`

		id := "id"
		date, err := time.Parse(timelayout, "2020-12-12")
		assert.NoError(t, err)
		expectedBirthday := entity.NewBirthday(id, date)

		_, err = conn.Exec(
			context.Background(),
			statement,
			expectedBirthday.Id,
			expectedBirthday.Date,
		)

		assert.NoError(t, err)

		repo := NewBirthdayPostgresRepository(
			conn,
		)

		output, err := repo.FindAll(context.Background())

		assert.NoError(t, err)
		assert.Len(t, output, 1)

		actualBirthday := output[0]
		assert.Equal(t, expectedBirthday.Id, actualBirthday.Id)
		assert.Equal(t, expectedBirthday.PrettyBirthday(), actualBirthday.PrettyBirthday())
	})

	t.Run("error in query", func(t *testing.T) {
		t.Cleanup(func() {
			err := postgresContainer.Restore(context.Background())
			assert.NoError(t, err)
		})

		conn, err := pgxpool.New(context.Background(), connStr)
		assert.NoError(t, err)
		defer conn.Close()

		statement := `
			DROP TABLE IF EXISTS birthdays;
		`

		_, err = conn.Exec(context.Background(), statement)

		assert.NoError(t, err)

		repo := NewBirthdayPostgresRepository(
			conn,
		)

		output, err := repo.FindAll(context.Background())

		assert.Error(t, err)
		assert.Empty(t, output)
	})
}
