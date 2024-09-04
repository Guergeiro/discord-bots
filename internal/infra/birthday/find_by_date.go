package birthday

import (
	"context"
	"errors"
	"time"

	"github.com/guergeiro/discord-bots/pkg/domain/entity"
	"github.com/jackc/pgx/v5"
)

func (r BirthdayPostgresRepository) FindByDate(
	ctx context.Context,
	date time.Time,
) ([]entity.Birthday, error) {
	statement := `
		SELECT *
		FROM birthdays
		WHERE EXTRACT(MONTH FROM date)=$1
			AND EXTRACT(DAY FROM date)=$2;
	`

	rows, err := r.conn.Query(ctx, statement, date.Month(), date.Day())
	if err != nil {
		return []entity.Birthday{}, err
	}
	output, err := pgx.CollectRows(
		rows,
		func(row pgx.CollectableRow) (entity.Birthday, error) {
			var id *string
			var date *time.Time
			err := row.Scan(&id, &date)
			if err != nil {
				return entity.Birthday{}, err
			}
			if id == nil || date == nil {
				return entity.Birthday{}, errors.New("No birthday exist")
			}
			return entity.NewBirthday(*id, *date), nil
		},
	)
	if err != nil {
		return []entity.Birthday{}, err
	}
	return output, nil
}
