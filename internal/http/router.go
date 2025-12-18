package http

import (
	"soccer-api/internal/http/handlers"
	"soccer-api/internal/repo"
	"soccer-api/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"

	customMiddleware "soccer-api/internal/http/middleware"
)

func CreateRouter(pool *pgxpool.Pool) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.CleanPath)
	//TODO: Remake these with static factories
	//Repos
	userRepo := repo.UserRepo{Pool: pool}
	teamRepo := repo.TeamRepo{Pool: pool}
	playerRepo := repo.PlayerRepo{Pool: pool}
	marketRepo := repo.MarketRepo{Pool: pool}

	//Services
	teamService := service.TeamService{
		TeamRepo:   teamRepo,
		PlayerRepo: playerRepo,
	}
	marketService := service.MarketService{
		MarketRepo: marketRepo,
		TeamRepo:   teamRepo,
		PlayerRepo: playerRepo,
	}

	//Handlers
	authHandler := &handlers.AuthHandler{AuthRepo: userRepo, TeamRepo: teamRepo}
	teamHandler := &handlers.TeamHandler{Service: teamService}
	playerHandler := &handlers.PlayerHandler{Service: teamService}
	marketHandler := &handlers.MarketHandler{Service: marketService}

	r.Group(func(r chi.Router) {
		//Auth
		r.Post("/v1/auth/signup", authHandler.Signup)
		r.Post("/v1/auth/login", authHandler.Login)
	})

	r.Group(func(r chi.Router) {
		r.Use(customMiddleware.Auth)
		//Team routes
		r.Get("/v1/me/team", teamHandler.GetTeamInfo)
		r.Patch("/v1/me/team", teamHandler.UpdateTeam)
		r.Patch("/v1/me/players/{playerId}", playerHandler.UpdatePlayer)
		////Market routes
		r.Post("/v1/me/players/{playerId}/list", marketHandler.ListPlayer)
		r.Delete("/v1/me/players/{playerId}/list", marketHandler.CancelListing)
		r.Get("/v1/market", marketHandler.GetMarket)
		r.Post("/v1/market/{listingId}/buy", marketHandler.BuyPlayer)
	})

	return r
}
