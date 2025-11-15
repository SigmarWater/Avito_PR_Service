package pull_request

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	repoModel "github.com/SigmarWater/Avito_PR_Service/internal/repository/models"
)

func (r *PostgresPullRequestsRepository) GetActiveTeamMembers(ctx context.Context, teamName string, excludeUserId int) ([]*repoModel.RepoUser, error) {

	var teamId int
	checkTeamQuery := sq.Select("team_id").
		From("teams").
		Where(sq.Eq{"team_name": teamName}).
		PlaceholderFormat(sq.Dollar)

	checkQuery, checkArgs, err := checkTeamQuery.ToSql()
	if err != nil {
		log.Printf("failed to create check team query: %v\n", err)
		return nil, err
	}

	err = r.pool.QueryRow(ctx, checkQuery, checkArgs...).Scan(&teamId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("team %v is not found", teamName)
		}
		log.Printf("failed to check team existence: %v\n", err)
		return nil, err
	}

	builderSelect := sq.Select(
		"u.user_id",
		"u.username",
		"u.is_active").
		From("users AS u").
		InnerJoin("team_users AS tu ON u.user_id = tu.user_id").
		Where(sq.And{
			sq.Eq{"tu.team_id": teamId},
			sq.Eq{"u.is_active": true},
			sq.NotEq{"u.user_id": excludeUserId},
		}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builderSelect.ToSql()
	if err != nil {
		log.Printf("failed to create select_query for GetActiveTeamMembers: %v\n", err)
		return nil, err
	}

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		log.Printf("failed to get active members: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	activeUsers := make([]*repoModel.RepoUser, 0)

	for rows.Next() {
		user := &repoModel.RepoUser{}
		if err := rows.Scan(&user.UserId, &user.Username, &user.IsActive); err != nil {
			log.Printf("failed to scan user: %v\n", err)
			return nil, err
		}

		activeUsers = append(activeUsers, user)
	}

	if err := rows.Err(); err != nil {
		log.Printf("error iterating rows: %v\n", err)
		return nil, err
	}

	return activeUsers, nil
}
