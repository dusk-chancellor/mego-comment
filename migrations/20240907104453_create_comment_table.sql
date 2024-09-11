-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS comments
(

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS comments;
-- +goose StatementEnd
