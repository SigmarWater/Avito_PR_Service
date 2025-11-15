package pull_request

import (
	"context"
	"errors"
	serviceModel "github.com/SigmarWater/Avito_PR_Service/internal/models"
	"github.com/SigmarWater/Avito_PR_Service/internal/repository/converter"
	"log"

	sq "github.com/Masterminds/squirrel"
)

func (r *PostgresPullRequestsRepository) SetIsActive(ctx context.Context, userId int, isActive bool) (*serviceModel.User, error) {
	builderUpdate := sq.Update("users").
		Where(sq.Eq{"user_id": userId}).
		Set("is_active", isActive)

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		log.Printf("failed to create update_query for SetIsActive: %v\n", err)
		return nil, err
	}

	result, err := r.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed to set isActive: %v\n", err)
		return nil, err
	}

	// Проверяем, что обновление действительно произошло
	if result.RowsAffected() == 0 {
		log.Printf("user with id %d not found\n", userId)
		return nil, errors.New("user not found")
	}

	user, err := r.GetUser(ctx, userId)

	return converter.RepoUserToService(user), err
}
