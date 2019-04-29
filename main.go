package main

import (
	"net/http"
	"timesheets/timesheets"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	shift := new Shift()
	
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	http.ListenAndServe(":3333", r)
}

// SELECT s.start, s.finish, TIMESTAMPDIFF(MINUTE, s.start, s.finish) / 60 AS total from shifts s;
