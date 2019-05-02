package main

import (
	"database/sql"
	"log"
	"net/http"
	"timesheets/internal/shifts"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	_ "github.com/go-sql-driver/mysql"
)

// Routes sets up all the routes for this api.
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set the content type to application/json.
		middleware.Logger,          // Loog all the API calls.
		middleware.DefaultCompress, // Add some gzip to that sweet sweet json.
		middleware.RedirectSlashes, // Tralling slashes are anoying, just make it work.
		middleware.Recoverer,       // Recover from panics without crashing the server.
	)

	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/shifts", shifts.Routes())
	})

	return router
}

func main() {

	// Setup the database and test it's connection.
	db, err := sql.Open("mysql", "timesheets:uERv3dKu5643qjUN@/timesheets?parseTime=true")

	if err != nil {
		log.Panicf("Unable to open the database, err: %s\n", err.Error())
	}
	defer db.Close() // Close the database when the main function ends.

	err = db.Ping() // Ping the database to test the username and password.
	if err != nil {
		log.Panicf("Unable to establish a connection to the database, err: %s\n", err.Error())
	}

	log.Print("Connected to the database successfully")

	shifts.Init(db)

	// Setup the http API.
	router := Routes()

	// Dump each route into the console.
	// TODO: Put this under a flag or debug level.
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Unable to walk the api, err: %s\n", err.Error())
	}

	log.Fatal(http.ListenAndServe(":8081", router)) // Start the http server.
}
