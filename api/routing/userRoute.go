package routing

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/tomazJakomin/go-base-app/models"
	"github.com/tomazJakomin/go-base-app/repositories"
	"github.com/unrolled/render"
	"gorm.io/gorm"
	"net/http"
)

type userRouter struct {
	db         gorm.DB
	repository repositories.UserRepository
	render     *render.Render
}

func newUserRouter(db *gorm.DB) *userRouter {
	return &userRouter{
		db:         *db,
		repository: repositories.NewUserRepository(*db),
		render:     render.New(),
	}
}

func RegisterUserRoutes(db *gorm.DB, router chi.Router) {
	handler := newUserRouter(db)

	router.Route("/users", func(r chi.Router) {
		r.Post("/", handler.createUser)
		r.Get("/", handler.getUsers)
	})
}

func GetUserFromRequest(w http.ResponseWriter, r *http.Request) (*models.User, error) {
	user := &models.User{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return nil, err
	}

	if err := user.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	return user, nil
}

func (router *userRouter) createUser(w http.ResponseWriter, r *http.Request) {
	user, err := GetUserFromRequest(w, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	newUser, err := router.repository.CreateUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	router.render.JSON(w, http.StatusCreated, newUser)
}

func (router *userRouter) getUsers(w http.ResponseWriter, r *http.Request) {
	result, err := router.repository.GetAllUsers()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	router.render.JSON(w, http.StatusOK, result)
}
