package pull_request

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	serviceModel "github.com/SigmarWater/Avito_PR_Service/internal/models"
	repoModel "github.com/SigmarWater/Avito_PR_Service/internal/repository/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type PostgresPullRequestsRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresPullRequestsRepository(pool *pgxpool.Pool) *PostgresPullRequestsRepository {
	return &PostgresPullRequestsRepository{pool: pool}
}

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

	var member *repoModel.RepoUser
	if err := r.pool.QueryRow(ctx, query, args...).Scan(member); err != nil {
		log.Printf("failed to insert into table teams: %v\n", err)
		return nil, err
	}
	return member, nil
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

func (r *PostgresPullRequestsRepository) GetIdUsersFromTeam(ctx context.Context, teamId int) ([]int, error) {
	membersId := make([]int, 0)
	builderSelect := sq.Select("user_id").
		PlaceholderFormat(sq.Dollar).
		From("team_users").
		Where(sq.Eq{"team_id": teamId})

	query, args, err := builderSelect.ToSql()
	if err != nil {
		log.Printf("failed to create select_query from table team_users: %v\n", err)
		return nil, err
	}

	rows, err := r.pool.Query(ctx, query, args)
	if err != nil {
		log.Printf("failed to select from table team_users: %v\n", err)
		return nil, err
	}

	for rows.Next() {
		var idUser int
		if err := rows.Scan(&idUser); err != nil {
			return nil, err
		}

		membersId = append(membersId, idUser)
	}

	return membersId, nil
}

func (r *PostgresPullRequestsRepository) GetUser(ctx context.Context, userId int, teamName string) (*repoModel.RepoUser, error) {
	builderSelect := sq.Select("user_id", "username", "is_active").
		PlaceholderFormat(sq.Dollar).
		From("users").
		Where(sq.Eq{"user_id": userId})

	query, args, err := builderSelect.ToSql()
	if err != nil {
		log.Printf("failed to create select_query from table users: %v\n", err)
		return nil, err
	}

	var user *repoModel.RepoUser

	if err := r.pool.QueryRow(ctx, query, args...).Scan(user); err != nil {
		log.Printf("failed to select from table users")
		return nil, err
	}

	return user, nil
}

func (r *PostgresPullRequestsRepository) GetTeamWithMembers(ctx context.Context, teamName string) (*repoModel.RepoTeam, error) {
	builderSelect := sq.Select("team_id", "team_name").
		PlaceholderFormat(sq.Dollar).
		From("teams").
		Where(sq.Eq{"team_name": teamName})

	query, args, err := builderSelect.ToSql()
	if err != nil {
		log.Printf("failed to create select_query from table teams: %v\n", err)
		return nil, err
	}

	var teamId int

	if err := r.pool.QueryRow(ctx, query, args...).Scan(&teamId); err != nil {
		log.Printf("failed to select from table teams")
		return nil, err
	}

	members := make([]*repoModel.RepoUser, 0)

	membersId, err := r.GetIdUsersFromTeam(ctx, teamId)
	if err != nil {
		return nil, err
	}

	for _, memberId := range membersId {
		user, err := r.GetUser(ctx, memberId, teamName)
		if err != nil {
			return nil, err
		}

		members = append(members, user)
	}

	return &repoModel.RepoTeam{
		TeamId:   teamId,
		TeamName: teamName,
		Members:  members,
	}, nil
}
