-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
  user_id SERIAL PRIMARY KEY,
  username TEXT not null,
  is_active BOOLEAN not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
