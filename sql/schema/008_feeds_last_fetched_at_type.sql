-- +goose Up
ALTER TABLE feeds
ALTER COLUMN last_fetched_at TYPE TIMESTAMP;
-- +goose Down
ALTER TABLE feeds
ALTER COLUMN last_fetched_at TYPE TIMESTAMP WITH TIME ZONE NOT NULL;