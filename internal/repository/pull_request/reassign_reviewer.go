package pull_request

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
)

func (r *PostgresPullRequestsRepository) ReassignReviewer(ctx context.Context, pullRequestId int, oldUserId int, newUserId int) error {
	// Проверяем статус PR - нельзя менять ревьюверов после MERGED
	pr, err := r.GetPullRequest(ctx, pullRequestId)
	if err != nil {
		return fmt.Errorf("failed to get pull request: %w", err)
	}

	if pr.Status == "MERGED" {
		return fmt.Errorf("cannot reassign reviewer on merged PR")
	}

	// Проверяем, что oldUserId является ревьювером
	checkOldQuery := sq.Select("reviewer_id").
		From("pull_requests_users").
		Where(sq.And{
			sq.Eq{"pull_request_id": pullRequestId},
			sq.Eq{"reviewer_id": oldUserId},
		}).
		PlaceholderFormat(sq.Dollar)

	checkOldSQL, checkOldArgs, err := checkOldQuery.ToSql()
	if err != nil {
		log.Printf("failed to create check query for old reviewer: %v\n", err)
		return err
	}

	var oldReviewerId int
	err = r.pool.QueryRow(ctx, checkOldSQL, checkOldArgs...).Scan(&oldReviewerId)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user %d is not assigned as reviewer to PR %d", oldUserId, pullRequestId)
		}
		return fmt.Errorf("failed to check if old user is reviewer: %w", err)
	}

	// Проверяем, что newUserId еще не является ревьювером
	checkNewQuery := sq.Select("reviewer_id").
		From("pull_requests_users").
		Where(sq.And{
			sq.Eq{"pull_request_id": pullRequestId},
			sq.Eq{"reviewer_id": newUserId},
		}).
		PlaceholderFormat(sq.Dollar)

	checkNewSQL, checkNewArgs, err := checkNewQuery.ToSql()
	if err != nil {
		log.Printf("failed to create check query for new reviewer: %v\n", err)
		return err
	}

	var newReviewerId int
	err = r.pool.QueryRow(ctx, checkNewSQL, checkNewArgs...).Scan(&newReviewerId)
	if err == nil {
		// Запись найдена - пользователь уже является ревьювером
		return fmt.Errorf("user %d is already assigned as reviewer to PR %d", newUserId, pullRequestId)
	}
	if err != sql.ErrNoRows {
		return fmt.Errorf("failed to check if new user is reviewer: %w", err)
	}

	// Удаляем старую запись
	builderDelete := sq.Delete("pull_requests_users").
		Where(sq.And{
			sq.Eq{"pull_request_id": pullRequestId},
			sq.Eq{"reviewer_id": oldUserId},
		}).
		PlaceholderFormat(sq.Dollar)

	deleteQuery, deleteArgs, err := builderDelete.ToSql()
	if err != nil {
		log.Printf("failed to create delete query for ReassignReviewer: %v\n", err)
		return err
	}

	deleteRes, err := r.pool.Exec(ctx, deleteQuery, deleteArgs...)
	if err != nil {
		log.Printf("failed to delete old reviewer: %v\n", err)
		return err
	}

	if deleteRes.RowsAffected() == 0 {
		return fmt.Errorf("old reviewer %d not found for PR %d", oldUserId, pullRequestId)
	}

	// Вставляем новую запись
	builderInsert := sq.Insert("pull_requests_users").
		Columns("pull_request_id", "reviewer_id").
		Values(pullRequestId, newUserId).
		PlaceholderFormat(sq.Dollar)

	insertQuery, insertArgs, err := builderInsert.ToSql()
	if err != nil {
		log.Printf("failed to create insert query for ReassignReviewer: %v\n", err)
		return err
	}

	_, err = r.pool.Exec(ctx, insertQuery, insertArgs...)
	if err != nil {
		log.Printf("failed to insert new reviewer: %v\n", err)
		return err
	}

	return nil
}
