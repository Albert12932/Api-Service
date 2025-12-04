-- +goose Up
-- +goose StatementBegin
ALTER TABLE questions ADD CONSTRAINT check_question_length CHECK(length(text) > 2);
ALTER TABLE answers ADD COLUMN text text;
UPDATE answers set text = 'default text' where text IS null;
ALTER TABLE answers alter column text set NOT NULL;
ALTER TABLE answers ADD CONSTRAINT check_answer_length CHECK(length(text) > 2);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE answers DROP CONSTRAINT IF EXISTS answers_text_length_check;
ALTER TABLE answers ALTER COLUMN text DROP NOT NULL;
ALTER TABLE answers DROP COLUMN IF EXISTS text;
ALTER TABLE questions DROP CONSTRAINT IF EXISTS questions_text_length_check;
-- +goose StatementEnd
