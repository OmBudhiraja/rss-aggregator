-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    url TEXT NOT NULL UNIQUE,
    published_at TIMESTAMP NOT NULL,
    feed_id UUID NOT NULL REFERENCES feed(id) ON DELETE CASCADE, 
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE posts;
