package models

import (
	"database/sql"
)

type RepoPullRequest struct {
	PullRequestId     string
	PullRequestName   string
	AuthorId          string
	Status            string
	AssignedReviewers []string
	CreatedAt         sql.NullTime
	MergedAt          sql.NullTime
}

type RepoPullRequestShort struct {
	PullRequestId   string
	PullRequestName string
	AuthorId        string
	Status          string
}
