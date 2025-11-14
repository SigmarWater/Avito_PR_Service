package converter

import (
	"database/sql"
	"time"

	serviceModel "github.com/SigmarWater/Avito_PR_Service/internal/models"
	repoModel "github.com/SigmarWater/Avito_PR_Service/internal/repository/models"
)

func repoTeamToService(team *repoModel.RepoTeam) *serviceModel.Team {
	if team == nil {
		return nil
	}

	members := make([]serviceModel.TeamMember, len(team.Members))
	for i, member := range team.Members {
		members[i] = serviceModel.TeamMember{
			UserId:   member.UserId,
			Username: member.Username,
			IsActive: member.IsActive,
		}
	}

	return &serviceModel.Team{
		TeamName: team.TeamName,
		Members:  members,
	}
}

func serviceTeamToRepo(team *serviceModel.Team) *repoModel.RepoTeam {
	if team == nil {
		return nil
	}

	members := make([]repoModel.RepoUser, len(team.Members))
	for i, member := range team.Members {
		members[i] = repoModel.RepoUser{
			UserId:   member.UserId,
			Username: member.Username,
			TeamName: team.TeamName,
			IsActive: member.IsActive,
		}
	}

	return &repoModel.RepoTeam{
		TeamName: team.TeamName,
		Members:  members,
	}
}

func repoUserToService(user *repoModel.RepoUser) *serviceModel.User {
	if user == nil {
		return nil
	}

	return &serviceModel.User{
		UserId:   user.UserId,
		Username: user.Username,
		TeamName: user.TeamName,
		IsActive: user.IsActive,
	}
}

func serviceUserToRepo(user *serviceModel.User) *repoModel.RepoUser {
	if user == nil {
		return nil
	}

	return &repoModel.RepoUser{
		UserId:   user.UserId,
		Username: user.Username,
		TeamName: user.TeamName,
		IsActive: user.IsActive,
	}
}

func repoUserWithPullRequestsToService(user *repoModel.RepoUserWithPullRequests) *serviceModel.UserWithPullRequests {
	if user == nil {
		return nil
	}

	prs := make([]serviceModel.PullRequestShort, len(user.PullRequests))
	for i, pr := range user.PullRequests {
		prs[i] = repoPullRequestShortToService(pr)
	}

	return &serviceModel.UserWithPullRequests{
		UserId:       user.UserId,
		PullRequests: prs,
	}
}

func repoPullRequestToService(pr *repoModel.RepoPullRequest) *serviceModel.PullRequest {
	if pr == nil {
		return nil
	}

	return &serviceModel.PullRequest{
		PullRequestId:     pr.PullRequestId,
		PullRequestName:   pr.PullRequestName,
		AuthorId:          pr.AuthorId,
		Status:            pr.Status,
		AssignedReviewers: cloneStrings(pr.AssignedReviewers),
		CreatedAt:         nullTimeToPtr(pr.CreatedAt),
		MergedAt:          nullTimeToPtr(pr.MergedAt),
	}
}

func servicePullRequestToRepo(pr *serviceModel.PullRequest) *repoModel.RepoPullRequest {
	if pr == nil {
		return nil
	}

	return &repoModel.RepoPullRequest{
		PullRequestId:     pr.PullRequestId,
		PullRequestName:   pr.PullRequestName,
		AuthorId:          pr.AuthorId,
		Status:            pr.Status,
		AssignedReviewers: cloneStrings(pr.AssignedReviewers),
		CreatedAt:         timePtrToNull(pr.CreatedAt),
		MergedAt:          timePtrToNull(pr.MergedAt),
	}
}

func repoPullRequestShortToService(pr repoModel.RepoPullRequestShort) serviceModel.PullRequestShort {
	return serviceModel.PullRequestShort{
		PullRequestId:   pr.PullRequestId,
		PullRequestName: pr.PullRequestName,
		AuthorId:        pr.AuthorId,
		Status:          pr.Status,
	}
}

func servicePullRequestShortToRepo(pr serviceModel.PullRequestShort) repoModel.RepoPullRequestShort {
	return repoModel.RepoPullRequestShort{
		PullRequestId:   pr.PullRequestId,
		PullRequestName: pr.PullRequestName,
		AuthorId:        pr.AuthorId,
		Status:          pr.Status,
	}
}

func cloneStrings(src []string) []string {
	if len(src) == 0 {
		return nil
	}

	dst := make([]string, len(src))
	copy(dst, src)
	return dst
}

func nullTimeToPtr(t sql.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}

	return nil
}

func timePtrToNull(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{}
	}

	return sql.NullTime{
		Time:  *t,
		Valid: true,
	}
}
