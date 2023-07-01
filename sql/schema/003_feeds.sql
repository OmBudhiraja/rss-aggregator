-- +goose Up
CREATE TABLE feed (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) on DELETE CASCADE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE feed;
