package pull_request

import (
	"context"
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	serviceModel "github.com/SigmarWater/Avito_PR_Service/internal/models"
	repoModel "github.com/SigmarWater/Avito_PR_Service/internal/repository/models"
)

func (r *PostgresPullRequestsRepository) CreatePullRequest(ctx context.Context, req *serviceModel.CreatePullRequestRequest) (*repoModel.RepoPullRequest, error) {
	builderInsert := sq.Insert("pull_requests").
		PlaceholderFormat(sq.Dollar).
		Columns("pull_request_name", "author_id").
		Values(req.PullRequestName, req.AuthorId).
		Suffix("RETURNING pull_request_id, status")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Printf("failed to create insert_query into table pull_requests: %v\n", err)
		return nil, err
	}

	var idPullRequest int
	var status string

	if err := r.pool.QueryRow(ctx, query, args...).Scan(&idPullRequest, &status); err != nil {
		log.Printf("failed to insert into table pull_requests: %v\n", err)
		return nil, err
	}

	return &repoModel.RepoPullRequest{
		PullRequestId:     idPullRequest,
		PullRequestName:   req.PullRequestName,
		AuthorId:          req.AuthorId,
		Status:            status,
		AssignedReviewers: []string{},
		CreatedAt:         sql.NullTime{Time: time.Now(), Valid: true},
		MergedAt:          sql.NullTime{},
	}, nil
}
