package pull_request

import (
	"context"
	"fmt"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	repoModel "github.com/SigmarWater/Avito_PR_Service/internal/repository/models"
)

func (r *PostgresPullRequestsRepository) MergePullRequest(ctx context.Context, pullRequestId int) (*repoModel.RepoPullRequest, error) {
	updateBuilder := sq.Update("pull_requests").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"pull_request_id": pullRequestId}).
		Set("status", "MERGED").
		Set("merged_at", time.Now())

	query, args, err := updateBuilder.ToSql()
	if err != nil {
		log.Printf("failed to create update_query for MergePullRequest: %v\n", err)
		return nil, err
	}

	result, err := r.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed to set merge for pull_request: %v\n", err)
		return nil, err
	}

	if result.RowsAffected() == 0 {
		return nil, fmt.Errorf("not found pull_request with pull_request_id: %v", pullRequestId)
	}

	return r.GetPullRequest(ctx, pullRequestId)
}
