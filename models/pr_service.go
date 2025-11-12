package models

type Team struct {
	TeamName string `json:"team_name"`
	Members  []User `json:"members"`
}

type User struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

type PullRequest struct {
	PullRequestId     string `json:"pull_request_id"`
	PullRequestName   string `json:"pull_request_name"`
	AuthorId          string `json:"author_id"`
	Status            string `json:"status"`
	AssignedReviewers []User `json:"assigned_reviewers"`
	NeedMoreReviewers bool   `json:"need_more_reviewers"`
}
