package models

type RepoUser struct {
	UserId   int
	Username string
	IsActive bool
}

// RepoUserWithPullRequests используется для GetPullRequestsForUser
type RepoUserWithPullRequests struct {
	UserId       string
	PullRequests []RepoPullRequestShort
}
