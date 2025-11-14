package models

// CreatePullRequestRequest - запрос на создание PR
type CreatePullRequestRequest struct {
	PullRequestId   int    `json:"pull_request_id"`
	PullRequestName string `json:"pull_request_name"`
	AuthorId        int    `json:"author_id"`
}

// MergePullRequestRequest - запрос на merge PR
type MergePullRequestRequest struct {
	PullRequestId string `json:"pull_request_id"`
}

// ReassignRequest - запрос на переназначение ревьювера
type ReassignRequest struct {
	PullRequestId string `json:"pull_request_id"`
	OldUserId     string `json:"old_user_id"`
}

// SetIsActiveRequest - запрос на изменение активности пользователя
type SetIsActiveRequest struct {
	UserId   string `json:"user_id"`
	IsActive bool   `json:"is_active"`
}

// ReassignResponse - ответ на переназначение ревьювера
type ReassignResponse struct {
	PR         *PullRequest `json:"pr"`
	ReplacedBy string       `json:"replaced_by"`
}
