package user

import (
	"dz_oksp/internal/handlers"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type userHandler struct {
	repository Repo
}

func NewUserHandler(repository Repo) handlers.Handler {
	return &userHandler{repository: repository}
}

func addHandler(router *http.ServeMux, method string, path string, handler http.Handler) {
	router.Handle(fmt.Sprintf("%s %s", method, path), handler)
}

func (h *userHandler) Register(router *http.ServeMux) {
	addHandler(router, http.MethodGet, "/user/{uid}", http.HandlerFunc(h.GetUser))
	addHandler(router, http.MethodPost, "/user/create", http.HandlerFunc(h.CreateUser))
	addHandler(router, http.MethodPatch, "/user/update", http.HandlerFunc(h.UpdateUser))
	addHandler(router, http.MethodDelete, "/user/delete/{uid}", http.HandlerFunc(h.DeleteUser))
}

func (h *userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	matched, _ := regexp.MatchString(`\d`, r.PathValue("uid"))
	if !matched {
		w.Header().Set("Error", "uid must be int")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := h.repository.GetUser(r.Context(), r.PathValue("uid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if (User{}) == user {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := h.repository.Create(r.Context(), user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Uid", id)
	w.WriteHeader(http.StatusCreated)
}

func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Header().Set("Error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.repository.Update(r.Context(), user)
	if err != nil {
		w.Header().Set("Error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	matched, _ := regexp.MatchString(`\d`, r.PathValue("id"))
	if !matched {
		w.Header().Set("Error", "id must be int")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := h.repository.Delete(r.Context(), r.PathValue("uid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
