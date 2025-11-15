package converter

import (
	"database/sql"
	"strconv"
	"time"

	serviceModel "github.com/SigmarWater/Avito_PR_Service/internal/models"
	repoModel "github.com/SigmarWater/Avito_PR_Service/internal/repository/models"
)

func RepoTeamToService(team *repoModel.RepoTeam) *serviceModel.Team {
	if team == nil {
		return nil
	}

	members := make([]serviceModel.TeamMember, len(team.Members))
	for i, member := range team.Members {
		members[i] = serviceModel.TeamMember{
			UserId:   intToString(member.UserId),
			Username: member.Username,
			IsActive: member.IsActive,
		}
	}

	return &serviceModel.Team{
		TeamName: team.TeamName,
		Members:  members,
	}
}

func ServiceTeamToRepo(team *serviceModel.Team) *repoModel.RepoTeam {
	if team == nil {
		return nil
	}

	members := make([]*repoModel.RepoUser, len(team.Members))
	for i, member := range team.Members {
		userId, err := stringToInt(member.UserId)
		if err != nil {
			return nil
		}
		members[i] = &repoModel.RepoUser{
			UserId:   userId,
			Username: member.Username,
			IsActive: member.IsActive,
		}
	}

	return &repoModel.RepoTeam{
		TeamName: team.TeamName,
		Members:  members,
	}
}

func RepoUserToService(user *repoModel.RepoUser) *serviceModel.User {
	if user == nil {
		return nil
	}

	return &serviceModel.User{
		UserId:   intToString(user.UserId),
		Username: user.Username,
		TeamName: "", // TeamName не хранится в RepoUser, нужно получать отдельно
		IsActive: user.IsActive,
	}
}

func ServiceUserToRepo(user *serviceModel.User) *repoModel.RepoUser {
	if user == nil {
		return nil
	}

	userId, err := stringToInt(user.UserId)
	if err != nil {
		return nil
	}
	return &repoModel.RepoUser{
		UserId:   userId,
		Username: user.Username,
		IsActive: user.IsActive,
	}
}

func RepoUserWithPullRequestsToService(user *repoModel.RepoUserWithPullRequests) *serviceModel.UserWithPullRequests {
	if user == nil {
		return nil
	}

	prs := make([]serviceModel.PullRequestShort, len(user.PullRequests))
	for i, pr := range user.PullRequests {
		prs[i] = repoPullRequestShortToService(pr)
	}

	return &serviceModel.UserWithPullRequests{
		UserId:       intToString(user.UserId),
		PullRequests: prs,
	}
}

func RepoPullRequestToService(pr *repoModel.RepoPullRequest) *serviceModel.PullRequest {
	if pr == nil {
		return nil
	}

	return &serviceModel.PullRequest{
		PullRequestId:     intToString(pr.PullRequestId),
		PullRequestName:   pr.PullRequestName,
		AuthorId:          intToString(pr.AuthorId),
		Status:            pr.Status,
		AssignedReviewers: cloneStrings(pr.AssignedReviewers),
		CreatedAt:         nullTimeToPtr(pr.CreatedAt),
		MergedAt:          nullTimeToPtr(pr.MergedAt),
	}
}

func ServicePullRequestToRepo(pr *serviceModel.PullRequest) *repoModel.RepoPullRequest {
	if pr == nil {
		return nil
	}

	pullRequestId, err := stringToInt(pr.PullRequestId)
	if err != nil {
		return nil
	}
	authorId, err := stringToInt(pr.AuthorId)
	if err != nil {
		return nil
	}
	return &repoModel.RepoPullRequest{
		PullRequestId:     pullRequestId,
		PullRequestName:   pr.PullRequestName,
		AuthorId:          authorId,
		Status:            pr.Status,
		AssignedReviewers: cloneStrings(pr.AssignedReviewers),
		CreatedAt:         timePtrToNull(pr.CreatedAt),
		MergedAt:          timePtrToNull(pr.MergedAt),
	}
}

func RepoPullRequestShortToService(pr repoModel.RepoPullRequestShort) serviceModel.PullRequestShort {
	return serviceModel.PullRequestShort{
		PullRequestId:   intToString(pr.PullRequestId),
		PullRequestName: pr.PullRequestName,
		AuthorId:        intToString(pr.AuthorId),
		Status:          pr.Status,
	}
}

func ServicePullRequestShortToRepo(pr serviceModel.PullRequestShort) repoModel.RepoPullRequestShort {
	pullRequestId, _ := stringToInt(pr.PullRequestId)
	authorId, _ := stringToInt(pr.AuthorId)
	return repoModel.RepoPullRequestShort{
		PullRequestId:   pullRequestId,
		PullRequestName: pr.PullRequestName,
		AuthorId:        authorId,
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

// intToString конвертирует int в string
func intToString(i int) string {
	return strconv.Itoa(i)
}

// stringToInt конвертирует string в int
func stringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}
