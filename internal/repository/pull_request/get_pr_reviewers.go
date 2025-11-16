package pull_request

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"log"
)

func (r *PostgresPullRequestsRepository) GetPRReviewers(ctx context.Context, pullRequestId int) ([]int, error) {
	builderSelect := sq.Select("reviewer_id").
		PlaceholderFormat(sq.Dollar).
		From("pull_requests_users").
		Where(sq.Eq{"pull_request_id": pullRequestId})

	query, args, err := builderSelect.ToSql()
	if err != nil {
		log.Printf("failed to create select_query for GetPRReviewers: %v\n", err)
		return nil, err
	}

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		log.Printf("failed to get PR reviewers: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	reviewers := make([]int, 0)

	for rows.Next() {
		var reviewerId int
		if err := rows.Scan(&reviewerId); err != nil {
			log.Printf("failed to scan reviewer_id: %v\n", err)
			return nil, err
		}

		reviewers = append(reviewers, reviewerId)
	}

	if err := rows.Err(); err != nil {
		log.Printf("error iterating rows: %v\n", err)
		return nil, err
	}

	// Возвращаем пустой массив, если нет ревьюверов (это валидный случай)
	return reviewers, nil
}
