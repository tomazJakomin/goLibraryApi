package routing

import (
	"github.com/go-chi/chi"
	routing "github.com/tomazJakomin/go-base-app/api/controllers"
	"gorm.io/gorm"
)

func RegisterUserRoutes(db *gorm.DB, router chi.Router) {
	controller := routing.NewUserController(db)

	router.Route("/users", func(r chi.Router) {
		r.Post("/", controller.CreateUser)
		r.Get("/", controller.GetUsers)
	})
}
