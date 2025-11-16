package service

import (
	"context"
	"github.com/SigmarWater/Avito_PR_Service/internal/models"
)

type PullRequestService interface {
	// CreateTeamWithMembers создаёт команду с участниками (создаёт/обновляет пользователей)
	CreateTeamWithMembers(ctx context.Context, req *models.Team) (*models.Team, error)

	// GetTeamWithMembers получает команду с участниками
	GetTeamWithMembers(ctx context.Context, teamName string) (*models.Team, error)

	// SetIsActive устанавливает флаг активности пользователя
	SetIsActive(ctx context.Context, userId string, isActive bool) (*models.User, error)

	// GetPullRequestsForUser получает PR'ы, где пользователь назначен ревьювером
	GetPullRequestsForUser(ctx context.Context, userId string) (*models.UserWithPullRequests, error)

	// CreatePullRequest создаёт PR и автоматически назначает до 2 ревьюверов из команды автора
	CreatePullRequest(ctx context.Context, req *models.CreatePullRequestRequest) (*models.PullRequest, error)

	// MergePullRequest помечает PR как MERGED (идемпотентная операция)
	MergePullRequest(ctx context.Context, pullRequestId string) (*models.PullRequest, error)

	//// Reassign переназначает конкретного ревьювера на другого из его команды
	//Reassign(ctx context.Context, req *models.ReassignRequest) (*models.ReassignResponse, error)
}
