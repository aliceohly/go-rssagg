-- +goose Up
ALTER TABLE users
    RENAME COLUMN updatable_at TO updated_at;
ALTER TABLE feeds
    RENAME COLUMN updatable_at TO updated_at;
ALTER TABLE feed_follow
    RENAME COLUMN updatable_at TO updated_at;
-- +goose Down
ALTER TABLE users
    RENAME COLUMN updated_at TO updatable_at;
ALTER TABLE feeds
    RENAME COLUMN updated_at TO updatable_at;
ALTER TABLE feed_follow
    RENAME COLUMN updated_at TO updatable_at;