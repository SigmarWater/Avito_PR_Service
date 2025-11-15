package pull_request

import (
	"context"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	repoModel "github.com/SigmarWater/Avito_PR_Service/internal/repository/models"
)

func (r *PostgresPullRequestsRepository) GetPullRequestsForUser(ctx context.Context, userId int) (*repoModel.RepoUserWithPullRequests, error) {
	// Проверяем существование пользователя
	user, err := r.GetUser(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Получаем PR'ы, где пользователь является ревьювером
	builderSelect := sq.Select(
		"pr.pull_request_id",
		"pr.pull_request_name",
		"pr.author_id",
		"pr.status").
		PlaceholderFormat(sq.Dollar).
		From("pull_requests AS pr").
		InnerJoin("pull_requests_users AS pru ON pr.pull_request_id = pru.pull_request_id").
		Where(sq.Eq{"pru.reviewer_id": userId})

	query, args, err := builderSelect.ToSql()
	if err != nil {
		log.Printf("failed to create select_query for GetPullRequestsForUser: %v\n", err)
		return nil, err
	}

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		log.Printf("failed to get pull requests for user: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	pullRequests := make([]repoModel.RepoPullRequestShort, 0)

	for rows.Next() {
		var pr repoModel.RepoPullRequestShort
		if err := rows.Scan(&pr.PullRequestId, &pr.PullRequestName, &pr.AuthorId, &pr.Status); err != nil {
			log.Printf("failed to scan pull request: %v\n", err)
			return nil, err
		}

		pullRequests = append(pullRequests, pr)
	}

	if err := rows.Err(); err != nil {
		log.Printf("error iterating rows: %v\n", err)
		return nil, err
	}

	return &repoModel.RepoUserWithPullRequests{
		UserId:       user.UserId,
		PullRequests: pullRequests,
	}, nil
}
