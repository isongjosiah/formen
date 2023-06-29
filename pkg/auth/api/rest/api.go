package rest

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"statuarius/internal/api"
	"statuarius/internal/config"
	"statuarius/internal/values"
	"statuarius/pkg/auth/deps"
	"time"
)

type API struct {
	Server *http.Server
	Config *config.AuthConfig
	Deps   *deps.AuthDep
}

func (a *API) Serve() error {
	a.Server = &http.Server{
		Addr:           fmt.Sprintf(":%d", a.Config.Port),
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		Handler:        a.setUpServerHandler(),
		MaxHeaderBytes: 1024 * 1024,
	}

	return a.Server.ListenAndServe()
}

func (a *API) Shutdown() error {
	return a.Server.Shutdown(context.Background())
}

// setUpServerHandler sets up handlers for the service
func (a *API) setUpServerHandler() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodPost, http.MethodGet, http.MethodPatch, http.MethodPut, http.MethodDelete},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-For", values.HeaderRequestID, values.HeaderRequestSource},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Method(http.MethodPost, "/auth/register", api.Handler(a.registerAccount))
	mux.Method(http.MethodPost, "/auth/login", api.Handler(a.registerAccount))
	mux.Method(http.MethodPost, "/auth/forgot-password", api.Handler(a.registerAccount))
	mux.Method(http.MethodPost, "/auth/reset-password", api.Handler(a.registerAccount))

	return mux
}
