package book

import (
	"dz_oksp/internal/handlers"
	"fmt"
	"net/http"
)

type bookHandler struct {
}

func NewBookHandler() handlers.Handler {
	return &bookHandler{}
}

func addHandler(router *http.ServeMux, method string, path string, handler http.Handler) {
	router.Handle(fmt.Sprintf("%s %s", method, path), handler)
}

func (h *bookHandler) Register(router *http.ServeMux) {
	addHandler(router, http.MethodGet, "/book/{id}", http.HandlerFunc(h.GetBook))
}

func (h *bookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(r.PathValue("uid")))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}
