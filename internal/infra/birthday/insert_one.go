package birthday

import (
	"context"

	"github.com/guergeiro/discord-bots/pkg/domain/entity"
)

func (r BirthdayPostgresRepository) InsertOne(
	ctx context.Context,
	input entity.Birthday,
) error {
	statement := `
		INSERT INTO birthdays (id, date) VALUES ($1, $2)
		ON CONFLICT (id, date)
		DO UPDATE SET date = EXCLUDED.date;
	`
	_, err := r.conn.Exec(ctx, statement, input.Id, input.Date)
	if err != nil {
		return err
	}
	return nil
}
