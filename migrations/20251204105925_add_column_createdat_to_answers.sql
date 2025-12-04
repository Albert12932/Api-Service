-- +goose Up
-- +goose StatementBegin
ALTER TABLE answers ADD COLUMN created_at timestamp DEFAULT now();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE answers DROP COLUMN IF EXISTS created_at;
-- +goose StatementEnd
