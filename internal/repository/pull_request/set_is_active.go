package pull_request

import (
	"context"
	"errors"
	"log"

	sq "github.com/Masterminds/squirrel"
	repoModel "github.com/SigmarWater/Avito_PR_Service/internal/repository/models"
)

func (r *PostgresPullRequestsRepository) SetIsActive(ctx context.Context, userId int, isActive bool) (*repoModel.RepoUser, error) {
	builderUpdate := sq.Update("users").
		Where(sq.Eq{"user_id": userId}).
		Set("is_active", isActive)

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		log.Printf("failed to create update_query in users table: %v\n", err)
		return nil, err
	}

	result, err := r.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed to update in table users: %v\n", err)
		return nil, err
	}

	// Проверяем, что обновление действительно произошло
	if result.RowsAffected() == 0 {
		log.Printf("user with id %d not found\n", userId)
		return nil, errors.New("user not found")
	}

	return r.GetUser(ctx, userId)
}
