package pull_request

import (
	"context"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	repoModel "github.com/SigmarWater/Avito_PR_Service/internal/repository/models"
)

func (r *PostgresPullRequestsRepository) GetPullRequest(ctx context.Context, pullRequestId int) (*repoModel.RepoPullRequest, error) {
	builderSelect := sq.Select("pull_request_id",
		"pull_request_name",
		"author_id",
		"status",
		"merged_at",
		"create_at").
		From("pull_requests").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"pull_request_id": pullRequestId})

	query, args, err := builderSelect.ToSql()
	if err != nil {
		log.Printf("failed to create get_pull_request query: %v\n", err)
		return nil, err
	}

	row := r.pool.QueryRow(ctx, query, args...)

	pullRequest := &repoModel.RepoPullRequest{}

	err = row.Scan(&pullRequest.PullRequestId,
		&pullRequest.PullRequestName,
		&pullRequest.AuthorId,
		&pullRequest.Status,
		&pullRequest.MergedAt,
		&pullRequest.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed scan pullRequest: %w", err)
	}
	return pullRequest, nil
}
