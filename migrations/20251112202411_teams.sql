-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS teams (
   team_id SERIAL PRIMARY KEY,
   team_name TEXT NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS teams;
-- +goose StatementEnd
