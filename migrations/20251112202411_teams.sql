-- +goose Up
-- +goose StatementBegin
-- DROP TABLE IF EXISTS pull_requests_user;
-- DROP TABLE IF EXISTS pull_requests;
-- DROP TABLE IF EXISTS team_users;
-- DROP TABLE IF EXISTS teams;

CREATE TABLE IF NOT EXISTS teams (
   team_id SERIAL PRIMARY KEY,
   team_name TEXT NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DROP TABLE IF EXISTS pull_requests_user;
-- DROP TABLE IF EXISTS pull_requests;
-- DROP TABLE IF EXISTS team_users;
DROP TABLE IF EXISTS teams;
-- +goose StatementEnd
