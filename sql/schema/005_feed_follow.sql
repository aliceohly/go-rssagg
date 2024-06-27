-- +goose Up
ALTER TABLE feeds_follow
    RENAME TO feed_follow;
-- +goose Down
ALTER TABLE feed_follow
    RENAME TO feeds_follow;