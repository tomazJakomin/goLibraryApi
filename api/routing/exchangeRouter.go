package routing

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/tomazJakomin/go-base-app/repositories"
	"github.com/tomazJakomin/go-base-app/services/bookExchange"
	"github.com/unrolled/render"
	"gorm.io/gorm"
	"net/http"
)

type exchangeRouter struct {
	db              gorm.DB
	render          *render.Render
	loanRepo        *repositories.LoanRepository
	exchangeService *bookExchange.ExchangeService
}

type loanRequest struct {
	UserId int
	BookId int
}

func RegisterExchangeRoutes(db *gorm.DB, router chi.Router) {
	exchangeRouter := exchangeRouter{
		db:              *db,
		render:          render.New(),
		loanRepo:        repositories.NewLoanRepository(db),
		exchangeService: bookExchange.NewExchangeService(db),
	}

	router.Route("/exchanges", func(r chi.Router) {
		r.Post("/return", exchangeRouter.returnBookForUser)
		r.Post("/loan", exchangeRouter.loanBook)
	})
}

func (router *exchangeRouter) returnBookForUser(w http.ResponseWriter, r *http.Request) {
	request := loanRequest{}
	error := json.NewDecoder(r.Body).Decode(&request)

	if error != nil {
		http.Error(w, error.Error(), 422)
		return
	}

	if request.UserId == 0 {
		http.Error(w, "UserId is not provided", http.StatusBadRequest)
		return
	}

	if request.BookId == 0 {
		http.Error(w, "BookId is not provided", http.StatusBadRequest)
		return
	}

	err := router.exchangeService.ReturnService.ReturnBookForUser(request.BookId, request.UserId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	router.render.JSON(w, http.StatusOK, "Book returned")
}

func (router *exchangeRouter) loanBook(w http.ResponseWriter, r *http.Request) {

	request := loanRequest{}
	error := json.NewDecoder(r.Body).Decode(&request)

	if error != nil {
		http.Error(w, error.Error(), 422)
		return
	}

	if request.UserId == 0 {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), 422)
		return
	}

	if request.BookId == 0 {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	err := router.exchangeService.LoanService.ExchangeBookForUser(request.BookId, request.UserId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	router.render.JSON(w, http.StatusOK, "Book lend to user")
}
