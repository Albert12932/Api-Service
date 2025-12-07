-- +goose Up
-- +goose StatementBegin
ALTER TABLE answers DROP CONSTRAINT answers_text_unique;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE answers ADD CONSTRAINT answers_text_unique UNIQUE(text);
-- +goose StatementEnd
