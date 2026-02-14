-- +goose Up
CREATE TABLE users (
    id          UUID Primary Key,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    name        TEXT NOT NULL
);

CREATE TABLE feeds (
    id          UUID Primary Key,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    name        TEXT NOT NULL,
    url         TEXT NOT NULL UNIQUE,
    user_id     UUID NOT NULL
        REFERENCES users(id)
        ON DELETE CASCADE
);

CREATE TABLE feed_follows (
    id          UUID Primary Key,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    user_id     UUID NOT NULL
        REFERENCES users(id)
        ON DELETE CASCADE,
    feed_id     UUID NOT NULL
        REFERENCES feeds(id)
        ON DELETE CASCADE,
    CONSTRAINT follow UNIQUE (user_id,feed_id)
);

-- +goose Down
DROP TABLE feeds;
DROP TABLE users;
DROP TABLE feed_follows;
