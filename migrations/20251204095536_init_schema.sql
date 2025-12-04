-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE Questions (
    id SERIAL primary key,
    text text not null,
    created_at timestamp DEFAULT now()
);
CREATE TABLE Users (
    id SERIAL primary key,
    name text not null
);
CREATE TABLE Answers (
    id SERIAL primary key,
    question_id int references Questions(id) ON DELETE CASCADE ,
    user_id int references Users(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS Answers CASCADE ;
DROP TABLE IF EXISTS Users CASCADE;
DROP TABLE IF EXISTS Questions CASCADE;
-- +goose StatementEnd
