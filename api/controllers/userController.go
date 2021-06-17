package routing

import (
	"encoding/json"
	"github.com/tomazJakomin/go-base-app/models"
	"github.com/tomazJakomin/go-base-app/repositories"
	"github.com/unrolled/render"
	"gorm.io/gorm"
	"net/http"
)

type userController struct {
	db         gorm.DB
	repository repositories.UserRepository
	render     *render.Render
}

func NewUserController(db *gorm.DB) *userController {
	return &userController{
		db:         *db,
		repository: repositories.NewUserRepository(*db),
		render:     render.New(),
	}
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

func (router *userController) CreateUser(w http.ResponseWriter, r *http.Request) {
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

func (router *userController) GetUsers(w http.ResponseWriter, r *http.Request) {
	result, err := router.repository.GetAllUsers()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	router.render.JSON(w, http.StatusOK, result)
}
