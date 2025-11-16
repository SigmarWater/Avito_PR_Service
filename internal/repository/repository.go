package repository

import (
	"context"

	serviceModel "github.com/SigmarWater/Avito_PR_Service/internal/models"
	repoModel "github.com/SigmarWater/Avito_PR_Service/internal/repository/models"
)

type PullRequestRepository interface {
	// CreateTeamWithMembers создаёт команду с участниками (создаёт/обновляет пользователей)
	CreateTeamWithMembers(ctx context.Context, team *serviceModel.Team) (*serviceModel.Team, error)

	// GetTeamWithMembers получает команду с участниками
	GetTeamWithMembers(ctx context.Context, teamName string) (*serviceModel.Team, error)

	// SetIsActive устанавливает флаг активности пользователя
	SetIsActive(ctx context.Context, userId int, isActive bool) (*serviceModel.User, error)

	// GetPullRequestsForUser получает PR'ы, где пользователь назначен ревьювером
	GetPullRequestsForUser(ctx context.Context, userId int) (*serviceModel.UserWithPullRequests, error)

	// CreatePullRequest создаёт PR в БД
	CreatePullRequest(ctx context.Context, req *serviceModel.CreatePullRequestRequest) (*serviceModel.PullRequest, error)

	// MergePullRequest помечает PR как MERGED
	MergePullRequest(ctx context.Context, pullRequestId int) (*serviceModel.PullRequest, error)

	// GetPullRequest получает PR по ID
	GetPullRequest(ctx context.Context, pullRequestId int) (*repoModel.RepoPullRequest, error)

	// GetActiveTeamMembers получает активных участников команды (исключая указанного пользователя)
	GetActiveTeamMembers(ctx context.Context, teamName string, excludeUserId int) ([]*repoModel.RepoUser, error)

	// GetUserTeam получает команду пользователя
	GetUserTeam(ctx context.Context, userId int) (string, error)

	//// ReassignReviewer переназначает ревьювера в PR
	//ReassignReviewer(ctx context.Context, pullRequestId int, oldUserId int, newUserId int) error

	// IsUserReviewer проверяет, является ли пользователь ревьювером PR
	IsUserReviewer(ctx context.Context, pullRequestId int, userId int) (bool, error)

	// GetPRReviewers получает список ревьюверов PR
	GetPRReviewers(ctx context.Context, pullRequestId int) ([]int, error)
}
