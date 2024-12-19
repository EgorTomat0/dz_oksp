package user

import (
	"dz_oksp/internal/handlers"
	"fmt"
	"net/http"
)

type userHandler struct {
}

func NewUserHandler() handlers.Handler {
	return &userHandler{}
}

func addHandler(router *http.ServeMux, method string, path string, handler http.Handler) {
	router.Handle(fmt.Sprintf("%s %s", method, path), handler)
}

func (h *userHandler) Register(router *http.ServeMux) {
	addHandler(router, http.MethodGet, "/user/{uid}", http.HandlerFunc(h.GetUser))
}

func (h *userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(r.PathValue("uid")))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}
