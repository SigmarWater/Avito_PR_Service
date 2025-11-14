-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS pull_requests_users;
DROP TABLE IF EXISTS pull_requests;
DROP TYPE IF EXISTS STATUS;

CREATE TYPE STATUS as ENUM('OPEN', 'MERGED');

CREATE TABLE IF NOT EXISTS pull_requests(
  pull_request_id SERIAL PRIMARY KEY,
  pull_request_name TEXT NOT NULL,
  author_id INT REFERENCES users(user_id),
  status STATUS NOT NULL DEFAULT('OPEN'),
  merged_at TIMESTAMP DEFAULT NULL,
  create_at TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pull_requests_users;
DROP TABLE IF EXISTS pull_requests;
DROP TYPE IF EXISTS STATUS;
-- +goose StatementEnd
