package models

type RepoTeam struct {
	TeamId   int
	TeamName string
	Members  []*RepoUser
}
