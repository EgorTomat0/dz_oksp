package main

import (
	"dz_oksp/internal/audioBook"
	"dz_oksp/internal/book"
	"dz_oksp/internal/user"
	"net/http"
	"time"
)

func main() {
	router := http.NewServeMux()

	userHandler := user.NewUserHandler()
	bookHandler := book.NewBookHandler()
	ABHandler := audioBook.NewABHandler()

	bookHandler.Register(router)
	ABHandler.Register(router)
	userHandler.Register(router)

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		server.Close()
	}
}
