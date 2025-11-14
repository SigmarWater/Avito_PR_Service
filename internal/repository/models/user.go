package models

type RepoUser struct {
	UserId   string
	Username string
	TeamName string
	IsActive bool
}

// RepoUserWithPullRequests используется для GetPullRequestsForUser
type RepoUserWithPullRequests struct {
	UserId       string
	PullRequests []RepoPullRequestShort
}
