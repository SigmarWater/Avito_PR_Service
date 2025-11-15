package pull_request

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
)

func (r *PostgresPullRequestsRepository) IsUserReviewer(ctx context.Context, pullRequestId int, userId int) (bool, error) {
	builderSelect := sq.Select("reviewer_id").
		PlaceholderFormat(sq.Dollar).
		From("pull_requests_users").
		Where(sq.And{
			sq.Eq{"pull_request_id": pullRequestId},
			sq.Eq{"reviewer_id": userId},
		})

	query, args, err := builderSelect.ToSql()
	if err != nil {
		log.Printf("failed to create select_query for IsUserReviewer: %v\n", err)
		return false, err
	}

	var reviewerId int
	err = r.pool.QueryRow(ctx, query, args...).Scan(&reviewerId)

	if err != nil {
		if err == sql.ErrNoRows {
			// Пользователь не является ревьювером - это не ошибка, просто false
			return false, nil
		}
		log.Printf("failed to check if user is reviewer: %v\n", err)
		return false, fmt.Errorf("failed to check if user is reviewer: %w", err)
	}

	// Если запись найдена, значит пользователь является ревьювером
	return true, nil
}
