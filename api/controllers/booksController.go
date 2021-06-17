package routing

import (
	"github.com/tomazJakomin/go-base-app/repositories"
	"github.com/unrolled/render"
	"gorm.io/gorm"
	"net/http"
)

type booksController struct {
	db         gorm.DB
	repository repositories.BookRepository
	render     *render.Render
}

func NewBooksController(db *gorm.DB) booksController {
	return booksController{
		db:         *db,
		repository: repositories.NewBookRepository(*db),
		render:     render.New(),
	}
}

func (router booksController) GetBooks(w http.ResponseWriter, r *http.Request) {
	if book, err := router.repository.GetAvailableBooks(); err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		router.render.JSON(w, http.StatusOK, book)
	}
}
