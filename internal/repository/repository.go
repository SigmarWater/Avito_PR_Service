package repository

import (
	"context"
	serviceModel "github.com/SigmarWater/Avito_PR_Service/internal/models"
	repoModel "github.com/SigmarWater/Avito_PR_Service/internal/repository/models"
)

type PullRequestRepository interface {
	// CreateTeamWithMembers создаёт команду с участниками (создаёт/обновляет пользователей)
	CreateTeamWithMembers(ctx context.Context, team *serviceModel.Team) (*repoModel.RepoTeam, error)

	// GetTeamWithMembers получает команду с участниками
	GetTeamWithMembers(teamName string) (*repoModel.RepoTeam, error)

	// SetIsActive устанавливает флаг активности пользователя
	SetIsActive(userId string, isActive bool) (*repoModel.RepoUser, error)

	// GetPullRequestsForUser получает PR'ы, где пользователь назначен ревьювером
	GetPullRequestsForUser(userId string) (*repoModel.RepoUserWithPullRequests, error)

	// CreatePullRequest создаёт PR в БД
	CreatePullRequest(req *serviceModel.CreatePullRequestRequest) (*repoModel.RepoPullRequest, error)

	// MergePullRequest помечает PR как MERGED
	MergePullRequest(pullRequestId string) (*repoModel.RepoPullRequest, error)

	// GetPullRequest получает PR по ID
	GetPullRequest(pullRequestId string) (*repoModel.RepoPullRequest, error)

	// GetActiveTeamMembers получает активных участников команды (исключая указанного пользователя)
	GetActiveTeamMembers(teamName string, excludeUserId string) ([]*repoModel.RepoUser, error)

	// GetUserTeam получает команду пользователя
	GetUserTeam(userId string) (string, error)

	// ReassignReviewer переназначает ревьювера в PR
	ReassignReviewer(pullRequestId string, oldUserId string, newUserId string) error

	// IsUserReviewer проверяет, является ли пользователь ревьювером PR
	IsUserReviewer(pullRequestId string, userId string) (bool, error)

	// GetPRReviewers получает список ревьюверов PR
	GetPRReviewers(pullRequestId string) ([]string, error)
}
