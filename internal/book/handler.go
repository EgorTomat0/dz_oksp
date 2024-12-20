package book

import (
	"dz_oksp/internal/handlers"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type bookHandler struct {
	repository Repo
}

func NewBookHandler(repository Repo) handlers.Handler {
	return &bookHandler{repository: repository}
}

func addHandler(router *http.ServeMux, method string, path string, handler http.Handler) {
	router.Handle(fmt.Sprintf("%s %s", method, path), handler)
}

func (h *bookHandler) Register(router *http.ServeMux) {
	addHandler(router, http.MethodGet, "/book/{id}", http.HandlerFunc(h.GetBook))
	addHandler(router, http.MethodPost, "/book/add", http.HandlerFunc(h.AddBook))
	addHandler(router, http.MethodPatch, "/book/update", http.HandlerFunc(h.UpdateBook))
	addHandler(router, http.MethodDelete, "/book/delete/{id}", http.HandlerFunc(h.DeleteBook))
}

func (h *bookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	matched, _ := regexp.MatchString(`\d`, r.PathValue("id"))
	if !matched {
		w.Header().Set("Error", "id must be int")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	book, err := h.repository.GetBook(r.Context(), r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if (Book{}) == book {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func (h *bookHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.Header().Set("Error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := h.repository.Create(r.Context(), book)
	if err != nil {
		w.Header().Set("Error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Id", id)
	w.WriteHeader(http.StatusCreated)
}
func (h *bookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.Header().Set("Error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.repository.Update(r.Context(), book)
	if err != nil {
		w.Header().Set("Error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *bookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	matched, _ := regexp.MatchString(`\d`, r.PathValue("id"))
	if !matched {
		w.Header().Set("Error", "id must be int")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := h.repository.Delete(r.Context(), r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
