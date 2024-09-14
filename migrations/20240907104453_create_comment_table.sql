-- +goose Up
-- +goose StatementBegin
-- Authors table
CREATE TABLE IF NOT EXISTS authors (
    id VARCHAR(255) PRIMARY KEY,
    first_name VARCHAR(100),
    middle_name VARCHAR(50),
    last_name VARCHAR(100),
    email VARCHAR(255),
    avatar VARCHAR(255)
);
-- Comments table
CREATE TABLE IF NOT EXISTS comments (
    id VARCHAR(255) PRIMARY KEY,
    content TEXT,
    post_id VARCHAR(255),
    parent_id VARCHAR(255),
    comment_id VARCHAR(255),
    author_id VARCHAR(255),
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
