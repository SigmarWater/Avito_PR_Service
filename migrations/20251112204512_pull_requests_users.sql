-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS pull_requests_users(
    pull_request_id INT REFERENCES pull_requests(pull_request_id),
    reviewer_id INT REFERENCES users(user_id),
    PRIMARY KEY(pull_request_id, reviewer_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pull_requests_users;
-- +goose StatementEnd
