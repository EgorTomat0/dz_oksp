package audioBook

import (
	"dz_oksp/internal/handlers"
	"fmt"
	"net/http"
)

type audioBookHandler struct {
}

func NewABHandler() handlers.Handler {
	return &audioBookHandler{}
}

func addHandler(router *http.ServeMux, method string, path string, handler http.Handler) {
	router.Handle(fmt.Sprintf("%s %s", method, path), handler)
}

func (h *audioBookHandler) Register(router *http.ServeMux) {
	addHandler(router, http.MethodGet, "/a_book/{id}", http.HandlerFunc(h.GetABook))
}

func (h *audioBookHandler) GetABook(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(r.PathValue("uid")))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}
