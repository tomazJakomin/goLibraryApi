package routing

import (
	"github.com/go-chi/chi"
	routing "github.com/tomazJakomin/go-base-app/api/controllers"
	"gorm.io/gorm"
)

func RegisterExchangeRoutes(db *gorm.DB, router chi.Router) {
	exchangeController := routing.NewExchangeController(db)

	router.Route("/exchanges", func(r chi.Router) {
		r.Post("/return", exchangeController.ReturnBookForUser)
		r.Post("/loan", exchangeController.LoanBook)
	})
}
