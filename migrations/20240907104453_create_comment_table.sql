-- +goose Up
-- +goose StatementBegin
-- Authors table
CREATE TABLE IF NOT EXISTS authors (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100),
    middle_name VARCHAR(50),
    last_name VARCHAR(100),
    email VARCHAR(255),
    avatar VARCHAR(255)
);
-- Comments table
CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,
    content TEXT,
    post_id INTEGER NULL,
    parent_id INTEGER NULL,
    author_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (author_id) REFERENCES authors(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
BEGIN;
DROP TABLE IF EXISTS authors;
DROP TABLE IF EXISTS comments;
COMMIT;
-- +goose StatementEnd
