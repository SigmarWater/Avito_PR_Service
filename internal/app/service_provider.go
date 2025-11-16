package app

import (
	"context"
	"github.com/SigmarWater/Avito_PR_Service/internal/api/pull_request"
	"github.com/SigmarWater/Avito_PR_Service/internal/closer"
	"github.com/SigmarWater/Avito_PR_Service/internal/config"
	"github.com/SigmarWater/Avito_PR_Service/internal/config/env"
	"github.com/SigmarWater/Avito_PR_Service/internal/repository"
	repoPullRequest "github.com/SigmarWater/Avito_PR_Service/internal/repository/pull_request"
	"github.com/SigmarWater/Avito_PR_Service/internal/service"
	servicePullRequest "github.com/SigmarWater/Avito_PR_Service/internal/service/pull_request"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type serviceProvider struct {
	pgConfig   config.PostgresConfig
	httpConfig config.HTTPConfig

	pgPool                *pgxpool.Pool
	pullRequestRepository repository.PullRequestRepository

	servicePullRequest service.PullRequestService

	pullRequestImpl *pull_request.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PostgresConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewRepositoryConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s\n", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		httpCfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %v\n", err.Error())
		}

		s.httpConfig = httpCfg
	}
	return s.httpConfig
}

func (s *serviceProvider) PgPool(ctx context.Context) *pgxpool.Pool {
	if s.pgPool == nil {
		pool, err := pgxpool.Connect(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to db: %v\n", err)
		}

		if err := pool.Ping(ctx); err != nil {
			log.Fatalf("failed to ping to db: %v\n", err)
		}

		closer.Add(func() error {
			pool.Close()
			return nil
		})
		s.pgPool = pool
	}

	return s.pgPool
}

func (s *serviceProvider) PullRequestRepository(ctx context.Context) repository.PullRequestRepository {
	if s.pullRequestRepository == nil {
		s.pullRequestRepository = repoPullRequest.NewPostgresPullRequestsRepository(s.PgPool(ctx))
	}
	return s.pullRequestRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.PullRequestService {
	if s.servicePullRequest == nil {
		s.servicePullRequest = servicePullRequest.NewService(s.PullRequestRepository(ctx))
	}

	return s.servicePullRequest
}

func (s *serviceProvider) Impl(ctx context.Context) *pull_request.Implementation {
	if s.pullRequestImpl == nil {
		s.pullRequestImpl = pull_request.NewImplementation(s.UserService(ctx))
	}
	return s.pullRequestImpl
}
