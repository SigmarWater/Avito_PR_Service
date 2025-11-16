package pull_request

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresPullRequestsRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresPullRequestsRepository(pool *pgxpool.Pool) *PostgresPullRequestsRepository {
	return &PostgresPullRequestsRepository{pool: pool}
}
