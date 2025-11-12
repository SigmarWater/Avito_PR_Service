-- +goose Up
-- +goose StatementBegin
-- DROP TABLE IF EXISTS pull_requests_users;
-- DROP TABLE IF EXISTS pull_requests;
-- DROP TABLE IF EXISTS team_users;
-- DROP TABLE IF EXISTS teams;
-- DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users(
  user_id SERIAL PRIMARY KEY,
  username TEXT not null,
  is_active BOOLEAN not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DROP TABLE IF EXISTS pull_requests_users;
-- DROP TABLE IF EXISTS pull_requests;
-- DROP TABLE IF EXISTS team_users;
-- DROP TABLE IF EXISTS teams;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
