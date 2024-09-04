package birthday

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type BirthdayPostgresRepository struct {
	conn *pgxpool.Pool
}

func NewBirthdayPostgresRepository(
	conn *pgxpool.Pool,
) BirthdayPostgresRepository {
	return BirthdayPostgresRepository{
		conn: conn,
	}
}
