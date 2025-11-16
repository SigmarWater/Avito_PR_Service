package app

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"net/http"

	"github.com/SigmarWater/Avito_PR_Service/internal/closer"
	"github.com/SigmarWater/Avito_PR_Service/internal/config"
)

const (
	readHeaderTimeout = 5 * time.Second
	shutdowmTimeout   = 10 * time.Second
)

type app struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
}

func NewApp(ctx context.Context) (*app, error) {
	a := &app{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *app) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	go func() {
		a.runHTTPServer()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("Started graceful shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), shutdowmTimeout)
	defer cancel()

	err := a.httpServer.Shutdown(ctx)
	if err != nil {
		log.Printf("Error shutdown: %s\n", err.Error())
		return err
	}
	log.Println("Success finished server")

	return nil
}

func (a *app) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *app) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *app) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()

	return nil
}

func (a *app) initHTTPServer(_ context.Context) error {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(10 * time.Second))
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"status":"ok","service":"pr_api"}`))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	r.Post("/team/add", a.serviceProvider.pullRequestImpl.CreateTeamWithMembers)
	r.Get("/team/get", a.serviceProvider.pullRequestImpl.GetTeamWithMembers)
	r.Post("/users/setIsActive", a.serviceProvider.pullRequestImpl.SetIsActive)
	r.Get("/users/getReview", a.serviceProvider.pullRequestImpl.GetPullRequestForUser)
	r.Post("/pullRequest/create", a.serviceProvider.pullRequestImpl.CreatePullRequest)
	r.Post("/pullRequest/merge", a.serviceProvider.pullRequestImpl.MergePullRequest)

	a.httpServer = &http.Server{
		Addr:              a.serviceProvider.HTTPConfig().Address(),
		Handler:           r,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	return nil
}

func (a *app) runHTTPServer() {

	log.Printf("Server start on %s\n", a.serviceProvider.HTTPConfig().Address())
	if err := a.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("Error start server: %s\n", err.Error())
		return
	}

	return
}
