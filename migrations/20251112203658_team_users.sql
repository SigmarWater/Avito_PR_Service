-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS team_users (
    team_id INT REFERENCES teams(team_id),
    user_id INT REFERENCES users(user_id),
    PRIMARY KEY (team_id, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS team_users;
-- +goose StatementEnd
