package pull_request

import (
	"context"
	"fmt"
	serviceModel "github.com/SigmarWater/Avito_PR_Service/internal/models"
	"github.com/SigmarWater/Avito_PR_Service/internal/repository/converter"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
)

func (r *PostgresPullRequestsRepository) MergePullRequest(ctx context.Context, pullRequestId int) (*serviceModel.PullRequest, error) {
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

	pullRequest, err := r.GetPullRequest(ctx, pullRequestId)

	return converter.RepoPullRequestToService(pullRequest), err
}
