package pull_request

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
)

func (r *PostgresPullRequestsRepository) GetUserTeam(ctx context.Context, userId int) (string, error) {
	builderSelect := sq.Select("t.team_name").
		From("users AS u").
		InnerJoin("team_users AS tu ON u.user_id = tu.user_id").
		InnerJoin("teams AS t ON t.team_id = tu.team_id").
		Where(sq.Eq{"u.user_id": userId}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builderSelect.ToSql()
	if err != nil {
		log.Printf("failed to create select_query for GetUserTeam: %v\n", err)
		return "", err
	}

	var teamName string
	err = r.pool.QueryRow(ctx, query, args...).Scan(&teamName)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("user %d is not found or not assigned to any team", userId)
		}
		log.Printf("failed to get teamName: %v\n", err)
		return "", err
	}

	return teamName, nil
}
