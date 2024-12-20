package audioBook

import (
	"dz_oksp/internal/handlers"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type audioBookHandler struct {
	repository Repo
}

func NewABHandler(repository Repo) handlers.Handler {
	return &audioBookHandler{repository: repository}
}

func addHandler(router *http.ServeMux, method string, path string, handler http.Handler) {
	router.Handle(fmt.Sprintf("%s %s", method, path), handler)
}

func (h *audioBookHandler) Register(router *http.ServeMux) {
	addHandler(router, http.MethodGet, "/a_book/{id}", http.HandlerFunc(h.GetABook))
	addHandler(router, http.MethodPost, "/a_book/add", http.HandlerFunc(h.AddABook))
	addHandler(router, http.MethodPatch, "/a_book/update", http.HandlerFunc(h.UpdateABook))
	addHandler(router, http.MethodDelete, "/a_book/delete/{id}", http.HandlerFunc(h.DeleteABook))

}

func (h *audioBookHandler) GetABook(w http.ResponseWriter, r *http.Request) {
	matched, _ := regexp.MatchString(`\d`, r.PathValue("id"))
	if !matched {
		w.Header().Set("Error", "id must be int")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	aBook, err := h.repository.GetAB(r.Context(), r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if (AudioBook{}) == aBook {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	json.NewEncoder(w).Encode(aBook)
}

func (h *audioBookHandler) AddABook(w http.ResponseWriter, r *http.Request) {
	var aBook AudioBook
	err := json.NewDecoder(r.Body).Decode(&aBook)
	if err != nil {
		w.Header().Set("Error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := h.repository.Create(r.Context(), aBook)
	if err != nil {
		w.Header().Set("Error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Id", id)
	w.WriteHeader(http.StatusCreated)
}

func (h *audioBookHandler) UpdateABook(w http.ResponseWriter, r *http.Request) {
	var aBook AudioBook
	err := json.NewDecoder(r.Body).Decode(&aBook)
	if err != nil {
		w.Header().Set("Error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.repository.Update(r.Context(), aBook)
	if err != nil {
		w.Header().Set("Error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *audioBookHandler) DeleteABook(w http.ResponseWriter, r *http.Request) {
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
