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
    url         TEXT NOT NULL,
    user_id     UUID NOT NULL,
    FOREIGN KEY (user_id) 
        REFERENCES users(id)
        ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;
DROP TABLE users;

