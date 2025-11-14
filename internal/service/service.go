package service

import (
	"github.com/SigmarWater/Avito_PR_Service/internal/models"
)

type PullRequestService interface {
	// CreateTeamWithMembers создаёт команду с участниками (создаёт/обновляет пользователей)
	CreateTeamWithMembers(req *models.Team) (*models.Team, error)

	// GetTeamWithMembers получает команду с участниками
	GetTeamWithMembers(teamName string) (*models.Team, error)

	// SetIsActive устанавливает флаг активности пользователя
	SetIsActive(userId string, isActive bool) (*models.User, error)

	// GetPullRequestsForUser получает PR'ы, где пользователь назначен ревьювером
	GetPullRequestsForUser(userId string) (*models.UserWithPullRequests, error)

	// CreatePullRequest создаёт PR и автоматически назначает до 2 ревьюверов из команды автора
	CreatePullRequest(req *models.CreatePullRequestRequest) (*models.PullRequest, error)

	// MergePullRequest помечает PR как MERGED (идемпотентная операция)
	MergePullRequest(pullRequestId string) (*models.PullRequest, error)

	// Reassign переназначает конкретного ревьювера на другого из его команды
	Reassign(req *models.ReassignRequest) (*models.ReassignResponse, error)
}
