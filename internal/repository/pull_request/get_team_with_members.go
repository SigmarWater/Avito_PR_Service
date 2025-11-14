package pull_request

import (
	"context"
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	repoModel "github.com/SigmarWater/Avito_PR_Service/internal/repository/models"
)

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

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		log.Printf("failed to select from table team_users: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var idUser int
		if err := rows.Scan(&idUser); err != nil {
			return nil, err
		}

		membersId = append(membersId, idUser)
	}

	return membersId, nil
}

func (r *PostgresPullRequestsRepository) GetUser(ctx context.Context, userId int) (*repoModel.RepoUser, error) {
	builderSelect := sq.Select("user_id", "username", "is_active").
		PlaceholderFormat(sq.Dollar).
		From("users").
		Where(sq.Eq{"user_id": userId})

	query, args, err := builderSelect.ToSql()
	if err != nil {
		log.Printf("failed to create select_query from table users: %v\n", err)
		return nil, err
	}

	var user repoModel.RepoUser

	if err := r.pool.QueryRow(ctx, query, args...).Scan(&user.UserId, &user.Username, &user.IsActive); err != nil {
		log.Printf("failed to select from table users: %v\n", err)
		return nil, err
	}

	return &user, nil
}

func (r *PostgresPullRequestsRepository) GetTeamWithMembers(ctx context.Context, teamName string) (*repoModel.RepoTeam, error) {
	// Используем JOIN для получения команды и всех пользователей одним запросом
	builderSelect := sq.Select("t.team_id", "t.team_name", "u.user_id", "u.username", "u.is_active").
		PlaceholderFormat(sq.Dollar).
		From("teams t").
		LeftJoin("team_users tu ON t.team_id = tu.team_id").
		LeftJoin("users u ON tu.user_id = u.user_id").
		Where(sq.Eq{"t.team_name": teamName})

	query, args, err := builderSelect.ToSql()
	if err != nil {
		log.Printf("failed to create select_query from table teams: %v\n", err)
		return nil, err
	}

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		log.Printf("failed to select from table teams: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var teamId int
	var teamNameFromDB string
	members := make([]*repoModel.RepoUser, 0)
	teamFound := false

	for rows.Next() {
		var userID sql.NullInt32
		var username sql.NullString
		var isActive sql.NullBool

		if err := rows.Scan(&teamId, &teamNameFromDB, &userID, &username, &isActive); err != nil {
			log.Printf("failed to scan row: %v\n", err)
			return nil, err
		}

		teamFound = true

		// Если пользователь существует (не NULL), добавляем его в список
		if userID.Valid {
			members = append(members, &repoModel.RepoUser{
				UserId:   int(userID.Int32),
				Username: username.String,
				IsActive: isActive.Bool,
			})
		}
	}

	if !teamFound {
		// Команда не найдена
		return nil, nil
	}

	return &repoModel.RepoTeam{
		TeamId:   teamId,
		TeamName: teamNameFromDB,
		Members:  members,
	}, nil
}
