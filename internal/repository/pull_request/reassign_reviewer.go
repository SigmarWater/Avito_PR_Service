package pull_request

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
)

func (r *PostgresPullRequestsRepository) ReassignReviewer(ctx context.Context, pullRequestId int, oldUserId int, newUserId int) error {
	builderUpdate := sq.Update("pull_requests_users").
		Set("reviewer_id", newUserId).
		Where(sq.And{
			sq.Eq{"pull_request_id": pullRequestId},
			sq.Eq{"reviewer_id": oldUserId},
		}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		log.Printf("failed to create update_query for ReassignReviewer: %v\n", err)
		return err
	}

	res, err := r.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed to update rewiewer: %v\n", err)
		return err
	}

	if res.RowsAffected() == 0 {
		log.Printf("not found pullRequest with id:%v\n and rewiever with id: %v\n", pullRequestId, oldUserId)
		return nil
	}

	return nil
}
