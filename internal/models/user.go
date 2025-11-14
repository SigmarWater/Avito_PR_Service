package models

type User struct {
	UserId       string             `json:"user_id"`
	Username     string             `json:"username"`
	TeamName     string             `json:"team_name"`
	IsActive     bool               `json:"is_active"`
}

// UserWithPullRequests используется для ответа /users/getReview
type UserWithPullRequests struct {
	UserId       string             `json:"user_id"`
	PullRequests []PullRequestShort `json:"pull_requests"`
}
