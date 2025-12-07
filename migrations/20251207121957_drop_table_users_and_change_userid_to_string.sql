-- +goose Up
-- +goose StatementBegin
ALTER TABLE answers DROP CONSTRAINT IF EXISTS answers_user_id_fkey;
ALTER TABLE answers ALTER COLUMN user_id TYPE text USING user_id::text;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name text NOT NULL
);

ALTER TABLE answers
    ALTER COLUMN user_id TYPE int
        USING user_id::integer;

ALTER TABLE answers
    ADD CONSTRAINT answers_user_id_fkey
        FOREIGN KEY (user_id)
            REFERENCES users(id)
            ON DELETE CASCADE;
-- +goose StatementEnd
