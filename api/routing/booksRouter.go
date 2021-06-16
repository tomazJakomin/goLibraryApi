package routing

import (
	"github.com/go-chi/chi"
	"github.com/tomazJakomin/go-base-app/repositories"
	"github.com/unrolled/render"
	"gorm.io/gorm"
	"net/http"
)

type booksRouter struct {
	db         gorm.DB
	repository repositories.BookRepository
	render     *render.Render
}

func RegisterBookRoutes(db *gorm.DB, router chi.Router) {
	handler := NewBooksRouter(db)

	router.Route("/books", func(r chi.Router) {
		r.Get("/", handler.GetBooks)
	})
}

func NewBooksRouter(db *gorm.DB) booksRouter {
	return booksRouter{
		db:         *db,
		repository: repositories.NewBookRepository(*db),
		render:     render.New(),
	}
}

func (router booksRouter) GetBooks(w http.ResponseWriter, r *http.Request) {
	if book, err := router.repository.GetAvailableBooks(); err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		router.render.JSON(w, http.StatusOK, book)
	}
}
