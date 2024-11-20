CREATE TABLE IF NOT EXISTS birthdays (
		id TEXT NOT NULL,
		date TIMESTAMPTZ NOT NULL,
		PRIMARY KEY(id, date)
);
SELECT create_hypertable('birthdays', by_range('date'));
