-- +goose Up
-- +goose StatementBegin
ALTER TABLE questions ADD CONSTRAINT question_text_unique UNIQUE(text);
ALTER TABLE answers ADD CONSTRAINT answers_text_unique UNIQUE(text);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE questions DROP CONSTRAINT IF EXISTS questions_text_unique;
ALTER TABLE questions DROP CONSTRAINT IF EXISTS answers_text_unique;
-- +goose StatementEnd
