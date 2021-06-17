package routing

import (
	"github.com/go-chi/chi"
	routing "github.com/tomazJakomin/go-base-app/api/controllers"
	"gorm.io/gorm"
)

func RegisterBookRoutes(db *gorm.DB, router chi.Router) {
	controller := routing.NewBooksController(db)

	router.Route("/books", func(r chi.Router) {
		r.Get("/", controller.GetBooks)
	})
}
