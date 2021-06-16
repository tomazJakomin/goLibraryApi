package routing

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/gorm"
)

func StartRouter(db *gorm.DB) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	RegisterBookRoutes(db, router)
	RegisterUserRoutes(db, router)
	RegisterExchangeRoutes(db, router)

	return router
}
