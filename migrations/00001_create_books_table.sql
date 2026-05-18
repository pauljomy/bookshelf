-- +goose Up
CREATE TABLE IF NOT EXISTS books (
    id      SERIAL PRIMARY KEY,
    title   TEXT NOT NULL,
    author  TEXT NOT NULL,
    isbn13  TEXT,
    shelf   TEXT
);

-- +goose Down
DROP TABLE IF EXISTS books;
