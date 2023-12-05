package user

import (
	"fmt"
	"net/http"
	"rest/internal/apperror"
	"rest/internal/handlers"
	"rest/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

const (
	usersURL = "/users"
	userURL  = "/user/:uuid"
)

type handler struct { // неэкспортируемя структура
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler { //конструктор интерфейс который возвращает структуру
	return &handler{
		logger: logger,
	}
}

func (h *handler) Registr(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodGet, userURL, apperror.Middleware(h.GetUserByUUID))
	router.HandlerFunc(http.MethodPost, userURL, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodPut, userURL, apperror.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodPatch, userURL, apperror.Middleware(h.PartiallyUpdateUser))
	router.HandlerFunc(http.MethodDelete, userURL, apperror.Middleware(h.DeleteUser))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {

	return apperror.ErrNotFound
	// w.WriteHeader(200)
	// w.Write([]byte("this is user list"))
	// return nil

}
func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) error {

	return apperror.NewAppError(nil, "хуй", "гавно", "жопка")
	// w.WriteHeader(204)
	// w.Write([]byte("this is get user by uuid"))
	// return nil
}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {

	return fmt.Errorf("this is api error")
	// w.WriteHeader(200)
	// w.Write([]byte("this is create user"))
	// return nil
}
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("this is fully update user"))
	return nil
}
func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("this is partially update user"))
	return nil
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("this is delete user by uuid"))

	return nil
}
