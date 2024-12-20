package main

import (
	"context"
	"dz_oksp/internal/audioBook"
	"dz_oksp/internal/book"
	"dz_oksp/internal/user"
	"dz_oksp/pkg/pgsql"
	"log"
	"net/http"
	"time"
)

func main() {
	router := http.NewServeMux()
	sc := pgsql.StorageCfg{
		Uname:    "postgres",
		Password: "anibliss",
		Host:     "localhost",
		Port:     "5432",
		DbName:   "postgres",
	}
	pgConn, err := pgsql.NewConn(context.TODO(), sc)
	if err != nil {
		log.Fatal(err)
	}
	userRepo := user.NewRepo(pgConn)
	bookRepo := book.NewRepo(pgConn)
	aBookRepo := audioBook.NewRepo(pgConn)

	userHandler := user.NewUserHandler(userRepo)
	bookHandler := book.NewBookHandler(bookRepo)
	ABHandler := audioBook.NewABHandler(aBookRepo)

	bookHandler.Register(router)
	ABHandler.Register(router)
	userHandler.Register(router)

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		server.Close()
	}
}
