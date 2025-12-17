package http

import (
	"soccer-api/internal/http/handlers"
	"soccer-api/internal/repo"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"

	customMiddleware "soccer-api/internal/http/middleware"
)

func CreateRouter(pool *pgxpool.Pool) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.CleanPath)

	userRepo := repo.UserRepo{
		Pool: pool,
	}

	authHandler := &handlers.AuthHandler{
		AuthRepo: userRepo,
	}

	r.Group(func(r chi.Router) {
		//Auth
		r.Post("/v1/auth/signup", authHandler.Signup)
		r.Post("/v1/auth/login", authHandler.Login)
	})

	r.Group(func(r chi.Router) {
		r.Use(customMiddleware.Auth)
		//Team routes
		r.Get("/v1/me/team", nil)
		r.Patch("/v1/me/team", nil)
		r.Patch("/v1/me/players/{playerId}", nil)
		//Market routes
		r.Post("/v1/me/players/{playerId}/list", nil)
		r.Delete("/v1/me/players/{playerId}/list\n", nil)
		r.Get("/v1/market", nil)
		r.Post("/v1/market/{listingId}/buy", nil)
	})

	return r
}
