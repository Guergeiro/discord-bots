package birthday

import (
	"context"
)

func (r BirthdayPostgresRepository) RemoveOne(
	ctx context.Context,
	id string,
) error {
	statement := `
		DELETE FROM birthdays
		WHERE id=$1;
	`
	_, err := r.conn.Exec(ctx, statement, id)
	if err != nil {
		return err
	}
	return nil
}
