package pull_request

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	serviceModel "github.com/SigmarWater/Avito_PR_Service/internal/models"
	repoModel "github.com/SigmarWater/Avito_PR_Service/internal/repository/models"
)

func (r *PostgresPullRequestsRepository) InsertUser(ctx context.Context, teamMember serviceModel.TeamMember) (*repoModel.RepoUser, error) {
	builderInsert := sq.Insert("users").
		PlaceholderFormat(sq.Dollar).
		Columns("username", "is_active").
		Values(teamMember.Username, teamMember.IsActive).
		Suffix("RETURNING user_id, username, is_active")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Printf("failed to create insert_query into table users: %v\n", err)
		return nil, err
	}

	var member repoModel.RepoUser
	if err := r.pool.QueryRow(ctx, query, args...).Scan(&member.UserId, &member.Username, &member.IsActive); err != nil {
		log.Printf("failed to insert into table users: %v\n", err)
		return nil, err
	}
	return &member, nil
}

func (r *PostgresPullRequestsRepository) InsertTeamIdWithUserId(ctx context.Context, teamId int, userId int) error {
	builderInsert := sq.Insert("team_users").
		PlaceholderFormat(sq.Dollar).
		Columns("team_id", "user_id").
		Values(teamId, userId)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Printf("failed to create insert_query into table team_users: %v\n", err)
		return err
	}

	if _, err := r.pool.Exec(ctx, query, args...); err != nil {
		log.Printf("failed to insert into table teams: %v\n", err)
		return err
	}
	return nil
}

func (r *PostgresPullRequestsRepository) CreateTeamWithMembers(ctx context.Context, team *serviceModel.Team) (*repoModel.RepoTeam, error) {
	builderInsert := sq.Insert("teams").
		PlaceholderFormat(sq.Dollar).
		Columns("team_name").
		Values(team.TeamName).
		Suffix("RETURNING team_id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Printf("failed to create insert_query into table teams: %v\n", err)
		return nil, err
	}

	var idTeam int
	if err := r.pool.QueryRow(ctx, query, args...).Scan(&idTeam); err != nil {
		log.Printf("failed to insert into table teams: %v\n", err)
		return nil, err
	}

	repoMembers := make([]*repoModel.RepoUser, len(team.Members))

	for idx, member := range team.Members {
		repoMember, err := r.InsertUser(ctx, member)
		if err != nil {
			return nil, err
		}

		repoMembers[idx] = repoMember

		err = r.InsertTeamIdWithUserId(ctx, idTeam, repoMember.UserId)
		if err != nil {
			return nil, err
		}
	}

	return &repoModel.RepoTeam{
		TeamId:   idTeam,
		TeamName: team.TeamName,
		Members:  repoMembers,
	}, nil
}
