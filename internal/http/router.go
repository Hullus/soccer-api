package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	customMiddleware "soccer-api/internal/http/middleware"
)

func CreateRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.CleanPath)

	//Auth
	r.Group(func(r chi.Router) {
		r.Post("/v1/auth/signup", nil)
		r.Post("/v1/auth/login", nil)
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
